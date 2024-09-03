package main

import (
	"fmt"
	"time"
)

/* การทำ Rate Limiting (การจำกัดอัตราการทำงาน)
ซึ่งเป็นเทคนิคที่สำคัญในการควบคุมการใช้งานทรัพยากรและรักษาคุณภาพการ
ให้บริการ (Quality of Service) โดยใช้ Goroutines, Channels
และ Tickers ในภาษา Go
*/
/* Logic ของ Rate Limiting
การจำกัดอัตราแบบพื้นฐาน
เมื่อเราต้องการจำกัดการจัดการคำขอ (requests) ที่เข้ามา
เราจะใช้ Channels สำหรับคำขอเหล่านี้

เราสร้าง Channel ที่เรียกว่า limiter ซึ่งจะได้รับค่าทุกๆ 200 มิลลิวินาที
ซึ่งทำหน้าที่เป็นตัวควบคุมในแผนการจำกัดอัตตราของเรา

ก่อนที่จะจัดการคำขอแต่ละรายการ เราจะรอรับค่าจาก limiter เพื่อจำกัดให้การจัดการ
คำขอแต่ละครั้งเกิดขึ้นทุกๆ 200 มิลลิวินาที

การอนุญาตให้เกิดการ Bursting ของคำขอ
บางครั้งเราต้องการให้เกิดการ Bursting ของคำขอในเวลา
สั้นๆ แต่ยังคงรักษาการจำกัดอัตราโดยรวมเอาไว้ ซึ่ง
สามารถทำได้โดยการทำ Buffer ให้กับ limiter

เราสร้าง channels ที่เรียกว่า burstyLimiter ซึ่ง
อนุญาตให้เกิดการ Bursting ของเหตุการณ์สูงสุด 3 ครั้ง

เราจะเติมค่าลงใน Channels นี้เพื่อให้เกิดการระเบิด

การเติมค่าให้กับ burstyLimiter
ทุกๆ 200 มิลลิวินาที เราจะพยายามเติมค่าใหม่ลงใน
burstyLimiter จนเต็มที่จำนวน 3 ค่า

จากนั้นเราจำลองการร้องขอเข้ามา 5 ครั้ง ซึ่ง 3 ครั้ง
แรกจะได้รับประโยชน์จากความสามารถในการ Bursting
ของ burstyLimiter
*/
func main() {
	// สร้าง channel สำหรับคำขอ 5 รายการ
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// สร้าง time.Tick เพื่อรับค่าทุกๆ 200 มิลลิวินาที
	limiter := time.Tick(200 * time.Millisecond)

	// จัดการคำขอแต่ละรายการ โดยรอให้ limiter ส่งค่ามาให้ก่อน
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// สร้าง burstyLimiter ที่มีขนาด buffer 3
	burstyLimiter := make(chan time.Time, 3)

	// ฟังก์ชันเติมค่าลงใน burstyLimiter ทุกๆ 200 มิลลิวินาที
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// สร้าง channel สำหรับคำขอแบบ bursty 5 รายการ
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// จัดการคำขอแต่ละรายการ โดยใช้ burstyLimiter
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}

/* โค้ดนี้แสดงให้เห็นถึงการใช้ Go เพื่อจำกัดอัตราการจัดการ
คำขอ (Rate Limiting) โดยใช้ Channels และ Tickers

ในส่วนแรกของโค้ด จะมีการจัดการคำขอทีละรายการด้วย
อัตราการจัดการ 1 รายการต่อ 200 มิลลิวินาที

ในส่วนที่สอง จะมีการจัดการคำขอแบบ Bursting 3 รายการ
ทันทีที่มีคำขอเข้ามา แล้วจัดการคำขอที่เหลือตามอัตราการ
จัดการที่กำหนดไว้

โค้ดนี้แสดงให้เห็นถึงวิธีการใช้งาน Rate Limiting ใน Go ที่
สามารถควบคุมการจัดการคำขอได้อย่างมีประสิทธิภาพและยืดหยุ่นตามสถานการณ์ที่ต้องการ
*/
