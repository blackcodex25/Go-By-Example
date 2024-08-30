package main

import (
	"fmt"
	"time"
)

/*การใช้ Ticker ในภาษา Go ซึ่งเป็นฟีเจอร์ที่ช่วย
ให้เราสามารถทำงานซ้ำ ๆ ในช่วงเวลาที่กำหนดได้
แตกต่างจาก Timer ที่ใช้สำหรับการทำงานเพียงครั้งเดียวในอนาคต
*/
/* สรุปบทความ
Timer ใช้สำหรับการทำงานเพียงครั้งเดียวในอนาคต ในขณะ
ที่ Ticker ใช้สำหรับการทำงานซ้ำๆ ที่ระยะเวลาที่กำหนด
ตัวอย่างในบทความแสดงให้เห็นการใช้ Ticker ที่ทำงานทุกๆ
500 มิลลิวินาทีจนกว่าจะถูกหยุด
Ticker จะส่งค่าวันที่และเวลาผ่าน channel ที่เกี่ยวข้องเมื่อ
ถึงเวลาที่กำหนด
เราสามารถหยุด Ticker ได้โดยใช้คำสั่ง Stop และเมื่อ
หยุดแล้ว Ticker จะไม่ส่งค่าใด ๆ อีกต่อไป
ในตัวอย่างนี้ Ticker จะทำงาน 3 ครั้งก่อนที่จะถูกหยุดหลัง
จาก 1600 มิลลิวินาที
*/
func main() {
	// สร้าง Ticker ที่จะ tick ทุกๆ 500 มิลลิวินาที
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// เริ่ม goroutine เพื่อรอรับค่า tick จาก ticker
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// รอ 1600 มิลลิวินาที ก่อนหยุด Ticker
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

/* อธิบายการทำงานของโค้ด
การสร้าง Ticker:
ticker := time.NewTicker(500 * time.Millisecond):
สร้าง Ticker ที่จะส่งค่า (tick) ทุกๆ 500 มิลลิ
วินาทีผ่าน channel ticker.C

การเริ่ม goroutine สำหรับรับค่า tick:
go func(){ ... }()
สร้าง goroutine เพื่อทำงานใน background ที่จะ
รอรับค่าจาก channel ของ Ticker

select ใช้สำหรับเลือกทำงานระหว่าง channel ต่างๆ
ที่ใช้ใน goroutine นี้:
case <-done:: ถ้าได้รับค่าจาก channel
done (ซึ่งใช้เพื่อหยุดการทำงานของ goroutine)
จะหยุดการทำงานและออกจาก loop

case t := <-ticker.C:
ถ้าได้รับค่าจาก channel ticker.Cถ้าได้รับค่าจาก channel ticker.C
ซึ่งเป็นเวลาปัจจุบันที่ Ticker ส่งมา จะพิมพ์ค่าเวลานั้นออกมา

การหยุด Ticker:
รอ 1600 มิลลิวินาทีเพื่อให้ Ticker ทำงาน 3 ครั้ง
ticker.Stop()
หยุด Ticker จากการส่งค่า tick อีก
done <- true
ส่งค่า true ไปที่ channel done เพื่อบอกให้
goroutine หยุดทำงาน

ผลลัพธ์ที่คาดหวัง:
เมื่อรันโปรแกรมนี้ Ticker จะทำงานและส่งค่า tick 3 ครั้ง
ก่อนที่จะถูกหยุดหลังจาก 1600 มิลลิวินาที
โปรแกรมจะแสดงข้อความ "Ticker stopped" หลังจาก
หยุด Ticker
*/

/* สรุป
โค้ดนี้แสดงให้เห็นวิธีการใช้ Ticker สำหรับการทำงานซ้ำในช่วง
เวลาที่กำหนด และการหยุด Ticker เมื่อไม่ต้องการให้มันทำงาน
อีกต่อไป
*/
