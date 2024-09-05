package main

import (
	"fmt"
	"sync"
)

// ตัวอย่างโค้ด Mutexes ใน Go

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock() // ล็อกเพื่อป้องกันการเข้าถึง counter พร้อมกัน
			counter++
			mu.Unlock() // ปลดล็อกหลังจากใช้งานเสร็จ
		}()
	}
	wg.Wait()

	fmt.Println("Counter:", counter)
}

/* โครงสร้างของ Mutexes
สร้างตัวแปร Mutex: var mu sync.Mutex
การล็อก: mu.Lock()
การปลดล็อก: mu.Unlock()
*/
/* การทำงาน
ในโค้ดข้างต้น:
1.มีตัวแปร counter ซึ่งถูกแชร์ระหว่าง goroutines
2.ใช้ mu.Lock() เพื่อทำการล็อกก่อนที่จะเข้าถึงตัวแปร counter
3.ทำการเพิ่มค่า counter
4.ปลดล็อกด้วย mu.Unlock() เพื่อให้ goroutine อื่นสามารถ
เข้าถึง counter ได้
5.ใช้ wg.Wait() เพื่อรอให้ goroutine ทุกตัวทำงานเสร็จสิ้น
ก่อนที่โปรแกรมจะจบการทำงาน
*/
/* สรุป
การใช้ Mutexes ใน Go เป็นวิธีการสำคัญในการจัดการการเข้าถึง
ทรัพยากรที่ใช้ร่วมกันเพื่อป้องกันปัญหาการแข่งขันของข้อมูล (race conditions)
ซึ่งจะช่วยให้โปรแกรมทำงานได้อย่างถูกต้อง และปลอดภัยในการทำงานแบบขนาน
(concurrent programming)
*/
