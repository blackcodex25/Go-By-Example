package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ตัวอย่างการใช้งาน Atomic Counters
func main() {
	var counter int32
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

/* Race Condition Best Practices
1.ใช้ sync.Mutex 2.ใช้ Atomic Operation  */

/* การทำงานของโค้ด
ตัวแปร counter: ตัวแปรนี้ถูกประกาศเป็น int32 ซึ่งจะ
ถูกใช้เป็น counter ที่ปรับค่าได้แบบ atomic

การเพิ่มค่า counter: ในแต่ละ goroutine เราใช้ฟังก์ชัน atomic.AddInt32
เพื่อเพิ่มค่า counter โดยระบุที่อยู่ของตัวแปร (&counter) และเพิ่มค่า 1 ให้กับมัน
การดำเนินการนี้จะเกิดขึ้นในลักษณะ atomic ซึ่งหมายความว่าค่า counter
จะถูกเพิ่มในแต่ละ goroutine โดยไม่เกิดปัญหา race condition

การรอให้ goroutines เสร็จสิ้น
เราใช้ sync.WaitGroup เพื่อรอให้ goroutines ทั้งหมดทำงานเสร็จก่อน
ที่จะแสดงค่าของ counter สุดท้าย
*/
/* สรุปเนื้อหาสำคัญ
Atomic Counters ช่วยให้สามารถนับค่าที่ถูกต้อง
และปลอดภัยในสภาวะที่มีหลายๆ goroutine เข้าถึงตัวแปรเดียวกันพร้อมกัน
ได้โดยไม่เกิดปัญหา race condition

แพ็กเกจ sync/atomic ใน Go มีฟังก์ชันสำหรับการ
ดำเนินการบนตัวแปรแบบ atomic ซึ่งสามารถใช้กับตัวแปร
ประเภท int32, int64, uint32, uint64, และตัวชี้ (pointer) ได้

การใช้งาน WaitGroup ร่วมกับ atomic counters
ทำให้เราสามารถจัดการการทำงานพร้อมกันใน Go ได้อย่างมีประสิทธิภาพและปลอดภัย
*/
