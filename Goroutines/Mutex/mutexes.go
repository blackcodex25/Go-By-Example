package main

import (
	"fmt"
	"sync"
)

/* วิธีการใช้ mutex เพื่อจัดการสถานะที่ซับซ้อนมากขึ้น
ใน Go โดยใช้ Container ซึ่งมี map ที่เก็บตัวนับหลายตัว (counters)
และใช้ mutex เพื่อซิงโครไนซ์การเข้าถึงข้อมูลจากหลายๆ goroutine พร้อมกัน
*/
/* รายละเอียดการทำงาน
Container Struct: ประกอบด้วย mu ซึ่งเป็นตัวแปร sync.Mutex
และ counters ที่เป็น map เก็บคู่ค่า key-value โดยที่ key เป็น string
และ value เป็น int

inc Method: เป็นฟังก์ชันที่ใช้เพิ่มค่าตัวนับใน counters โดยต้อง
ล็อก mutex ก่อนการเข้าถึง counters และปลดล็อกหลังการเข้าถึงเสร็จ
โดยใช้ defer เพื่อให้การปลดล็อกเกิดขึ้นทันทีเมื่อฟังก์ชันทำงานเสร็จ

doIncrement Function:
ฟังก์ชันนี้จะเพิ่มค่าตัวนับใน counters โดยวนซ้ำ n ครั้ง
และเรียกใช้ inc เพื่อเพิ่มค่าตัวนับใน counters
*/
type Container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}

func main() {
	c := Container{
		counter: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counter)
}

/* Logic และ การทำงานของโค้ด
1.Container Struct:
Container มี sync.Mutex (mu) และ map
(counters) ที่เก็บค่าตัวนับสำหรับแต่ละ string

2.Initialization:
ใน main ฟังก์ชัน สร้าง Container และกำหนดค่าเริ่มต้น
ของ counters เป็น {"a": 0, "b": 0}

3.Incrementing Counters:
ฟังก์ชัน doIncrement ถูกสร้างเพื่อเพิ่มค่าตัวนับใน
counters โดยวนซ้ำ n ครั้ง

เรียกใช้ wg.Add(3) เพื่อบอกว่าเรากำลังจะเริ่ม 3 goroutines ที่ต้องการการรอ
สร้าง 3 goroutines ที่จะเพิ่มค่าตัวนับใน counters โดย 2
goroutines จะเพิ่มค่าของ a และ 1 goroutine จะเพิ่มค่า
ของ b

4.Mutex Locking/Unlocking:
ใน inc method ของ Container จะมีการล็อก mutex (c.mu.Lock())
ก่อนที่จะเข้าถึง counters และปลดล็อก (defer c.mu.Unlock())
หลังจากการปรับปรุงเสร็จ

5.Wait for Goroutines:
ใช้ wg.Wait() เพื่อรอจนกว่า goroutines ทั้งหมดจะทำงาน
เสร็จสิ้น

6.ผลลัพธ์:
เมื่อรันโค้ด ผลลัพธ์ที่ได้จะแสดงว่าค่าตัวนับใน counters ได้
รับการอัปเดตตามที่คาดหวัง (map[a:20000 b:10000])

*อธิบายเพิ่มเติม ฟังก์ชัน doIncrement
ฟังก์ชัน doIncrement จะทำการวนลูปทั้งหมด n ครั้ง โดยในแต่ละ
รอบของลูปจะเรียกใช้ฟังก์ชัน inc เพื่อเพิ่มค่าของตัวนับ (counters)
ตามชื่อ (name) ที่ถูกระบุไว้
*ถ้าเราเรียกใช้ doIncrement("a", 10000) ลูปจะทำงานทั้งหมด
*10,000 ครั้งสำหรับตัวนับที่มีชื่อว่า "a"
*ดังนั้น ถ้า n = 10000 ฟังก์ชันนี้จะวนลูป 10,000 ครั้ง
*/

/* สรุป
การใช้ mutex ในการจัดการการเข้าถึง map ที่ใช้ร่วมกันระหว่าง
หลาย goroutines ช่วยให้มั่นใจได้ว่าข้อมูลจะไม่ถูกปรับเปลี่ยน
พร้อมๆ กัน ซึ่งอาจทำให้เกิดปัญหาข้อมูลไม่ถูกต้องหรือ race conditions

การล็อก mutex ก่อนเข้าถึงและปลดล็อกหลังจากการเข้าถึงเสร็จสิ้น
เป็นวิธีที่ปลอดภัยในการจัดการข้อมูลในสถานการณ์ที่มีการ
ประมวลผลขนาน

ฟังก์ชัน doIncrement ในโค้ดที่กล่าวถึงถูกสร้างขึ้นเพื่อเพิ่มค่าของตัวนับ (counters)
ในโครงสร้าง Container โดยใช้ชื่อ (name) ของตัว
นับและจำนวนครั้ง (n) ที่ต้องการเพิ่มค่า การทำงานของ
doIncrement คือการเรียกใช้ฟังก์ชัน inc ในลูป โดย inc จะทำ
หน้าที่ล็อก Mutex เพื่อป้องกันการเข้าถึงตัวนับพร้อมกันจากหลายๆ
goroutine และเพิ่มค่าตัวนับตามที่กำหนด จากนั้นจะปลดล็อก Mutex
เมื่อการทำงานเสร็จสิ้น

ในโปรแกรมนี้ จะมีการเรียกใช้ doIncrement จากหลายๆ goroutine
เพื่อเพิ่มค่าตัวนับ ซึ่ง Mutex จะช่วยให้การเพิ่มค่านั้นปลอดภัยจาก
ปัญหาการเข้าถึงทรัพยากรพร้อมกันโดยหลาย goroutine ซึ่งอาจทำให้
เกิดข้อผิดพลาดหรือข้อมูลที่ไม่ถูกต้อง

ประเด็นเพิ่มเติมเกี่ยวกับการทำงานของ Mutex
การใช้ Mutex เพื่อป้องกันข้อผิดพลาด:
โค้ดนี้ใช้ Mutex เพื่อป้องกันการเกิด race condition
ซึ่งจะเกิดขึ้นเมื่อหลาย goroutine เข้าถึงและปรับปรุงค่าของตัวนับพร้อมกัน
โดยไม่ใช้กลไกในการควบคุมการเข้าถึง

การใช้ defer เพื่อปลดล็อก Mutex:
การใช้ defer เพื่อปลดล็อก Mutex ทันทีที่งานในฟังก์ชันเสร็จสิ้น ช่วยให้มั่นใจได้ว่า
Mutex จะถูกปลดล็อกอย่างถูกต้อง แม้ว่าจะมีข้อผิดพลาดเกิดขึ้น ภายในฟังก์ชัน

การใช้งาน WaitGroup:
WaitGroup ถูกใช้เพื่อรอให้ทุก goroutine เสร็จสิ้นการทำงาน
ก่อนที่โปรแกรมหลักจะดำเนินการต่อ

การใช้ Mutex แบบนี้เป็นวิธีการที่ปลอดภัยและมีประสิทธิภาพในการ
จัดการสถานะที่ซับซ้อนในการเขียนโปรแกรมแบบ concurrent ด้วย Go
โดยเฉพาะอย่างยิ่งเมื่อมีการเข้าถึงและปรับปรุงข้อมูลจากหลาย
goroutine พร้อมกัน
*/
