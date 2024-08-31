package main

import (
	"fmt"
	"sync"
	"time"
)

/* วิธีการใช้ WaitGroup ในการรอให้ goroutine หลายๆ ตัวทำงานเสร็จสิ้น
วิธีนี้ช่วยจัดการกับการทำงานที่เป็นไปพร้อมกัน (concurrency) โดยรอให้ทุก
goroutine ทำงานเสร็จทั้งหมดก่อนที่จะดำเนินการต่อไป
*/

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Work %d done\n", id)

}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}

/* การทำงานของโค้ด
โค้ดนี้เป็นตัวอย่างการใช้ WaitGroup เพื่อรอให้ goroutine
ห้าตัวทำงานเสร็จ โดยแต่ละ goroutine จะทำงานในฟังก์ชัน
worker ที่จะใช้เวลาในการทำงานหนึ่งวินาที (time.Sleep(time.Second))

ฟังก์ชัน worker(id int)
ฟังก์ชันนี้รับค่า id เป็นพารามิเตอร์เพื่อระบุตัวตนของ worker
เมื่อ worker เริ่มทำงาน จะพิมพ์ข้อความ "Worker [id] starting"
จากนั้น worker จะหยุดทำงานเป็นเวลา 1 วินาทีเพื่อจำลองการทำงานที่ใช้เวลานาน
เมื่อเสร็จสิ้น จะพิมพ์ข้อความ "Worker [id] done"

ในฟังก์ชัน main()
ประกาศตัวแปร wg เป็น sync.WaitGroup เพื่อใช้จัดการ goroutine
ใช้ลูป for เพื่อสร้าง goroutine 5 ตัว (จาก i = 1 ถึง i = 5)
สำหรับแต่ละ goroutine จะเรียก wg.Add(1) เพื่อเพิ่ม counter ของ WaitGroup
บ่งบอกว่าเรากำลังเริ่ม goroutine ใหม่

สร้าง goroutine โดยใช้ go func() ซึ่งเป็นการ
ประกาศฟังก์ชันแบบนิรนาม (anonymous function)
ภายใน goroutine นี้ จะเรียกใช้ฟังก์ชัน worker(i)
โดยส่งค่า i ที่ได้จากลูปเป็นพารามิเตอร์
ใช้ defer wg.Done() เพื่อแจ้ง WaitGroup ว่า goroutine นี้เสร็จสิ้นแล้ว

หลังจากลูปสิ้นสุดลง จะเรียก wg.Wait() เพื่อบล็อก
การทำงานของ main() จนกว่าทุก goroutine จะเสร็จสิ้น
(ค่า counter ของ WaitGroup กลับไปเป็น 0)
*/
