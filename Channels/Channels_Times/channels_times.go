package main

import (
	"fmt"
	"time"
)

/* บ่อยครั้งที่เราต้องการให้โค้ด Go ทำงานในอนาคตหรือทำซ้ำใน
ช่วงเวลาที่กำหนด ฟีเจอร์ timer และ ticker ที่มีอยู่ในตัว
Go ทำให้การทำทั้งสองสิ่งนี้เป็นเรื่องง่าย เราจะเริ่มดูที่ timer
ก่อนแล้วค่อยไปที่ ticker

Timer แสดงถึงเหตุการณ์เดียวในอนาคต เราบอก timer ว่า
เราต้องการรอนานเท่าใด และมันจะให้ channel ที่
จะแจ้งเตือนเมื่อถึงเวลานั้น ตัวอย่างนี้ timer จะรอ 2 วินาที

คำสั่ง <-timer1.C จะหยุดชั่วคราวบน channel C ของ
timer จนกว่าจะส่งค่าเพื่อระบุว่า timer ได้ทำงานแล้ว
หากเราเพียงต้องการรอ เราอาจใช้ time.Sleep แทนได้
เหตุผลหนึ่งที่ timer อาจมีประโยชน์คือเราสามารถยกเลิก
timer ได้ก่อนที่จะทำงาน ตัวอย่างต่อไปนี้แสดงถึงการทำเช่นนั้น

ให้เวลา timer2 พอที่จะทำงานหากมันถูกตั้งให้ทำงาน เพื่อ
แสดงว่ามันถูกหยุดแล้วจริงๆ

timer ตัวแรกจะทำงานประมาณ 2 วินาทีหลังจากที่เราเริ่ม
โปรแกรม แต่ timer ตัวที่สองควรถูกหยุดก่อนที่จะมีโอกาสทำงาน
*/

func main() {
	// สร้าง Timer ที่รอ 2 วินาที
	timer1 := time.NewTimer(2 * time.Second)

	// รอจนกว่า Timer 1 จะส่งสัญญาณ (block ที่  timer1.c)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// สร้าง Timer ตัวที่สองที่รอ 1 สินาที
	timer2 := time.NewTimer(time.Second)

	// เริ่ม goroutine เพื่อรอการทำงานของ timer 2
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// หยุด Timer 2 ก่อนที่จะทำงาน
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// รอ 2 วินาที เพื่อให้แน่ใจว่า Timer 2 หยุดแล้วจริงๆ
	time.Sleep(2 * time.Second)
}

/* อธิบายการทำงาน
การสร้าง Timer แรก:
timer1 := time.NewTimer(2 * time.Second)
สร้าง Timer ที่ตั้งเวลาให้ทำงานหลังจาก 2 วินาที
โดยใช้ time.NewTimer
การรอ Timer แรกทำงาน
<-timer1.C
รอ (block) จนกว่า channel timer1.C จะส่ง
สัญญาณว่าการรอ 2 วินาทีสิ้นสุดลงแล้ว
fmt.Println("Timer 1 fired")
พิมพ์ข้อความว่า Timer 1 ทำงานแล้ว เมื่อได้รับ
สัญญาณจาก Timer

การสร้าง Timer ที่สองและการยกเลิก
timer2 := time.NewTimer(time.Second)
สร้าง Timer ที่สองซึ่งตั้งเวลาให้ทำงานหลังจาก 1 วินาที
go func(){ ... }()
เริ่ม goroutine เพื่อรอ Timer 2 ทำงานใน background
stop2 := timer2.Stop()
หยุด Timer 2 ก่อนที่มันจะส่งสัญญาณ โดยใช้ timer2.Stop()
ถ้า Timer ถูกหยุดก่อนทำงาน stop2 จะเป็น true
fmt.Println("Timer 2 stopped")
พิมพ์ข้อความว่า Timer 2 ถูกหยุดแล้ว

การรอเพื่อยืนยันว่า Timer 2 หยุดแล้ว
time.Sleep(2 * time.Second)
รอเพิ่มเติมอีก 2 วินาที เพื่อให้แน่ใจว่า Timer 2 หยุดการทำงานแล้ว

สรุป
โค้ดนี้แสดงให้เห็นการใช้ Timer ในการกำหนดเวลาให้โค้ด
ทำงานในอนาคตและการยกเลิก Timer ก่อนที่จะทำงาน การรอ
ด้วย <-timer.C จะเป็นการ block จนกว่า Timer จะหมดเวลา
แต่เราสามารถใช้ timer.Stop() เพื่อยกเลิก Timer นั้นได้ หาก
Timer ถูกหยุดก่อนจะหมดเวลา Goroutine ที่กำลังรอ Timer
ทำงานจะไม่ถูกเรียกใช้งาน
*/
