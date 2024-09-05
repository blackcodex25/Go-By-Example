package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"math/rand"
)

/* การใช้ Stateful Goroutines เพื่อติดตามการอ่านและเขียน
ข้อมูลใน concurrent environment โดยใช้ช่องทาง (channels) และการ
นับจำนวนการอ่านและเขียนข้อมูลโดยใช้ atomic operations เพื่อความปลอดภัย
*/
// โครงสร้าง readOp ใช้สำหรับการอ่านข้อมูลจาก state
// มีฟิลด์ key ที่เป็นค่า int และ resp ที่เป็นช่องทางรับค่าผลลัพธ์การอ่าน
type readOp struct {
	key  int
	resp chan int
}

// โครงสร้าง writeOp ใช้สำหรับการเขียนข้อมูลลงใน state
// มีฟิลด์ key และ val ที่เป็นค่า int และ resp ที่เป็นช่องทางยืนยันการเขียน
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// ตัวแปร atomic ใช้สำหรับนับจำนวนการอ่านและเขียน
	var readOps uint64
	var writeOps uint64

	// สร้างช่องทางสำหรับการส่งข้อมูลการอ่านและเขียน
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Goroutine นี้เป็นเจ้าของ state ซึ่งเป็น map
	// ใช้ select สำหรับการตอบสนองต่อคำร้องขอการอ่านและเขียนจากช่องทาง reads และ writes
	go func() {
		var state = make(map[int]int) // สร้างแผนที่สำหรับเก็บข้อมูล state
		for {
			select {
			// กรณีที่มีการอ่านจากช่องทาง reads
			case read := <-reads:
				read.resp <- state[read.key] // ส่งค่าจาก state กลับไปยังช่องทาง resp
				// กรณีที่มีการเขียนข้อมูลลงในช่องทาง writes
			case write := <-writes:
				state[write.key] = write.val // เขียนค่า val ลงใน state
				write.resp <- true           // ยืนยันการเขียนกลับไปยังช่องทาง resp
			}
		}
	}()

	// สร้าง 100 goroutine สำหรับส่งคำร้องขอการอ่านข้อมูลจาก state
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// สร้าง readOp พร้อมกับช่องทางสำหรับรับค่าผลลัพธ์
				read := readOp{
					key:  rand.Intn(5), // สุ่มค่า key ในช่วง 0-4
					resp: make(chan int)}
				reads <- read                 // ส่งคำร้องขอการอ่านผ่านช่องทาง reads
				<-read.resp                   // รอรับผลลัพธ์จาก resp
				atomic.AddUint64(&readOps, 1) // เพิ่มจำนวนการอ่าน (นับแบบ atomic)
				time.Sleep(time.Millisecond)  // หน่วงเวลา 1 มิลลิวินาที
			}
		}()
	}

	// สร้าง 10 goroutine สำหรับส่งคำร้องขอการเขียนข้อมูลลงใน state
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// สร้าง writeOp พร้อมกับช่องทางสำหรับยืนยันการเขียน
				write := writeOp{
					key:  rand.Intn(5),   // สุ่มค่า key ในช่วง 0-4
					val:  rand.Intn(100), // สุ่มค่า val ในช่วง 0-99
					resp: make(chan bool)}
				writes <- write                // ส่งคำร้องขอการเขียนผ่านช่องทาง writes
				<-write.resp                   // รอรับการยืนยันการเขียนจาก resp
				atomic.AddUint64(&writeOps, 1) // เพิ่มจำนวนการเขียน (นับแบบ atomic)
				time.Sleep(time.Millisecond)   // หน่วงเวลา 1 มิลลิวินาที
			}
		}()
	}
	// รอให้ goroutines ทำงาน 1 วินาที
	time.Sleep(time.Second)

	// โหลดค่าจำนวนการอ่านและเขียนที่ทำเสร็จแล้ว
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writes:", writeOpsFinal)
}

/* รายละเอียดของโค้ด
ประกาศและนิยาม Structs:
readOp: ใช้สำหรับการอ่านข้อมูลจาก state โดย
ประกอบด้วย key (คีย์ที่ต้องการอ่าน) และ resp (ช่อง
ทางตอบกลับสำหรับค่าที่อ่านได้)
writeOp: ใช้สำหรับการเขียนข้อมูลลงใน state โดย
ประกอบด้วย key (คีย์ที่ต้องการเขียน), val (ค่าที่จะเขียน)
และ resp (ช่องทางตอบกลับสำหรับการเขียนข้อมูล)

การสร้าง Channels:
reads: ช่องทางสำหรับการส่งคำร้องขอการอ่านข้อมูล
writes: ช่องทางสำหรับการส่งคำร้องขอการเขียนข้อมูล

Goroutine ที่ครอบครอง State:
สร้าง goroutine หนึ่งตัวที่เป็นผู้ดูแล state ซึ่งเก็บข้อมูล
ในรูปของ map[int]int
ใช้ select เพื่อรับคำร้องขอจากช่องทาง reads และ writes
ถ้ามีคำร้องขออ่าน (read) จะส่งค่าที่อ่านจาก
state ไปยังช่องทางตอบกลับ (read.resp)
ถ้ามีคำร้องขอเขียน (write) จะอัพเดต state
และส่งการตอบกลับ (write.resp)

Goroutines สำหรับการอ่านข้อมูล:
สร้าง 100 goroutines ซึ่งจะส่งคำร้องขอการอ่านข้อมูล
ไปยังช่องทาง reads และรอรับค่าที่อ่านได้จากช่องทางตอบกลับ
นับจำนวนการอ่านที่สำเร็จโดยใช้
atomic.AddUint64(&readOps, 1)

Goroutines สำหรับการเขียนข้อมูล:
สร้าง 10 goroutines ซึ่งจะส่งคำร้องขอการเขียนข้อมูล
ไปยังช่องทาง writes และรอรับการตอบกลับจากการ
เขียน
นับจำนวนการเขียนที่สำเร็จโดยใช้
atomic.AddUint64(&writeOps, 1)

การรอและการแสดงผล:
รอ 1 วินาทีเพื่อให้ goroutines ทำงานเสร็จ
ใช้ atomic.LoadUint64(&readOps) และ atomic.LoadUint64(&writeOps)
เพื่อดึงค่าที่นับได้ และแสดงผลลัพธ์ออกมาทางหน้าจอ
*/
/* การทำงานของโค้ด
การจัดการ State:
Goroutine ที่ครอบครอง state รับผิดชอบในการจัดการ
การอ่านและเขียนข้อมูลโดยตรง
การใช้ select ช่วยให้สามารถตอบสนองต่อคำร้องขอ
ทั้งสองประเภทได้อย่างต่อเนื่อง

การนับการอ่านและเขียน:
การนับการอ่านและเขียนข้อมูลถูกทำโดยการใช้ atomic operations
เพื่อหลีกเลี่ยงปัญหาการเข้าถึงข้อมูลพร้อมกันจากหลาย goroutines

การควบคุมและการรอ:
การใช้ time.Sleep เพื่อให้แน่ใจว่า goroutines จะ
ทำงานครบถ้วนก่อนที่จะรายงานผลลัพธ์

โค้ดนี้ช่วยให้เห็นถึงการจัดการกับ state อย่างมีประสิทธิภาพใน
ระบบ concurrent โดยใช้ช่องทางและการนับจำนวนการดำเนินการ
โดยการใช้ atomic operations
*/
