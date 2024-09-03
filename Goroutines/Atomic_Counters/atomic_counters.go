package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/* การใช้ Atomic Counters ซึ่งเป็นวิธี
การที่ปลอดภัยและมีประสิทธิภาพในการจัดการสถานะ (state) ที่
เข้าถึงโดยหลาย Goroutine พร้อมกันในภาษา Go โดยไม่เกิดปัญหา Race Condition
บทความได้แนะนำการใช้แพ็กเกจ sync/atomic
ในการดำเนินการเกี่ยวกับค่าตัวเลขในลักษณะที่เป็น atomic
ซึ่งหมายความว่าการดำเนินการนั้นๆ จะไม่ถูกรบกวนหรือเข้ามาแทรกแซงโดย Goroutine อื่นๆ
*/
/* สรุปเนื้อหา
การจัดการสถานะใน Go: โดยปกติใน Go การจัดการ
สถานะระหว่างหลาย Gorouine มักทำผ่านการสื่อสารด้วย
Channels แต่ในบางกรณีสามารถใช้แพ็กเกจ

Atomic Counters: เป็นเคาน์เตอร์ที่ถูกควบคุมการเข้าถึง
ด้วยวิธีการที่รับประกันว่าเมื่อมีการเพิ่มหรือลดค่า จะไม่มีการแย่งกันทำงาน (Race Condition)
จาก Goroutine อื่น

การใช้งาน sync/atomic: ตัวอย่างโค้ดที่แสดงให้เห็นถึงการใช้งาน sync/atomic
เพื่อเพิ่มค่าเคาน์เตอร์จากหลาย Goroutine อย่างปลอดภัย
และการใช้ sync.WaitGroup เพื่อรอให้ Goroutine ทั้งหมดทำงานเสร็จ
*/
func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println("ops:", ops.Load())
}

/* การทำงานของโค้ด
โค้ดตัวอย่างนี้ประกอบด้วยสองส่วนหลักคือ การสร้าง Goroutine
หลายตัวเพื่อทำงานพร้อมกันและการใช้ตัวแปร atomic เพื่อ
จัดการเคาน์เตอร์ที่ถูกเพิ่มค่าโดย Goroutine เหล่านั้น

การประกาศตัวแปร ops และ wg
var ops atomic.Uint64
var wg sync.WaitGroup
ops: เป็นตัวแปรที่ใช้เก็บค่าเคาน์เตอร์โดยใช้ประเภท
atomic.Uint64 เพื่อให้สามารถเพิ่มค่าได้อย่างปลอดภัยจากหลาย Goroutine
wg: เป็น WaitGroup ที่ใช้ในการรอให้ Goroutine ทุกตัวทำงานเสร็จสมบูรณ์

การสร้าง Goroutine หลายตัวและเพิ่มค่าเคาน์เตอร์
for i := 0; i < 50; i++ {
    wg.Add(1)

    go func() {
        for c := 0; c < 1000; c++ {
            ops.Add(1)
        }
        wg.Done()
    }()
}
ลูปนี้สร้าง Goroutine ทั้งหมด 50 ตัว แต่ละ Goroutine จะเพิ่มค่า ops 1,000 ครั้ง
โดยใช้ ops.Add(1) ซึ่งเป็นการเพิ่มค่าเคาน์เตอร์แบบ atomic

wg.Add(1) จะเพิ่มจำนวนที่ต้องรอใน WaitGroup และ
wg.Done() จะลดจำนวนลงเมื่อ Goroutine ทำงานเสร็จ

รอให้ Goroutine ทั้งหมดทำงานเสร็จ
wg.Wait()
wg.Wait() จะรอจนกว่า Goroutine ทุกตัวจะส่งสัญญาณว่าเสร็จสิ้นงาน
(เมื่อ wg.Done() ถูกเรียกครบทุกครั้ง)

แสดงผลลัพธ์
fmt.Println("ops:", ops.Load())
ops.Load() ใช้ในการอ่านค่าของตัวแปร ops อย่าง
ปลอดภัยหลังจากที่ Goroutine ทุกตัวทำงานเสร็จแล้ว
ค่านี้จะถูกพิมพ์ออกมาเพื่อแสดงผลลัพธ์
*/
/* สรุปสำคัญ
โค้ดนี้แสดงถึงการใช้ Atomic Counters เพื่อจัดการ
เคาน์เตอร์ที่ถูกอัปเดตโดยหลาย Goroutine พร้อมกันอย่างปลอดภัย

sync/atomic ช่วยให้มั่นใจได้ว่าการดำเนินการเพิ่มค่า
หรือลดค่าของตัวแปรจะเกิดขึ้นอย่างถูกต้องโดยไม่มีการแย่งกันทำงาน

การใช้ sync.WaitGroup เป็นวิธีที่มีประสิทธิภาพในการรอ
ให้ทุก Goroutine ทำงานเสร็จ

Atomic Counters และการจัดการ Race Condition ถือเป็นแนว
ปฏิบัติที่สำคัญในการพัฒนาโปรแกรมที่ทำงานแบบ concurrent
ในภาษา Go โดยเฉพาะเมื่อมีการเข้าถึงหรืออัปเดตค่าจากหลาย Goroutine พร้อมกัน
*/
//adadawd
//wd
