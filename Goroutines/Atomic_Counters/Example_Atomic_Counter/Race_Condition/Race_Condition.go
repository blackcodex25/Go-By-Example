package main

import (
	"fmt"
	"sync"
)

// ตัวอย่าง Race Condition
func main() {
	var visitorCount int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			visitorCount++
			// atomic.AddInt32(&visitorCount, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final Visitor Count:", visitorCount)
}

/* Race Condition Best Practices
1.ใช้ sync.Mutex 2.ใช้ Atomic Operation  */

/* สมมติว่าเรามีตัวแปรหนึ่งที่เก็บค่าจำนวนการเข้าชมเว็บไซต์
(visitorCount) และทุกครั้งที่มีการเข้าชมเว็บไซต์ goroutine
จะเพิ่มค่าตัวแปรนี้ขึ้น 1

ในโค้ดนี้ เราคาดหวังว่า visitorCount จะมีค่าเป็น 1000 หลังจากการทำงานเสร็จสิ้น
แต่เนื่องจากไม่มีการควบคุมการเข้าถึงตัวแปร visitorCount อาจเกิด race condition ขึ้น

ผลกระทบของ Race Condition
เมื่อเกิด race condition ขึ้น
ผลลัพธ์ที่ได้อาจจะไม่ถูกต้องหรือแตกต่างกันไปในแต่ละครั้งที่รันโปรแกรม
ตัวแปร visitorCount อาจมีค่าเป็น 1000 หรืออาจจะเป็นค่าน้อยกว่านั้น
เนื่องจากการทำงานของ goroutine หลายๆ ตัวที่เข้าถึงและปรับค่า visitorCount พร้อมๆ กัน

*/
