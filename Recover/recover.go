package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}

/*
1.ฟังก์ชัน mayPanic:
ฟังก์ชันนี้ทำให้เกิด panic โดยการเรียก panic("a problem")

2.การเรียก defer:
ในฟังก์ชัน main มีการเรียกใช้ defer ที่ครอบคลุม
ฟังก์ชันนิรนาม ซึ่งภายในนั้นจะใช้ recover เพื่อจับ panic ที่เกิดขึ้น

3.การทำงานของ recover:
เมื่อ mayPanic ถูกเรียกและทำให้เกิด panic ฟังก์ชัน defer จะทำงาน
และ recover จะดักจับ panic จากนั้นจะพิมพ์ข้อความว่า
"Recovered. Error: a problem"

4.ข้อสังเกต:
ถ้าไม่มี recover โปรแกรมจะหยุดทำงานทันทีที่เกิด panic
แต่เนื่องจากมีการใช้ recover โปรแกรมจะไม่หยุดและสามารถ
แสดงข้อความบ่งบอกว่า panic ถูกจับได้
*/
/* สรุป
การใช้ recover มีประโยชน์ในสถานการณ์ที่เราไม่ต้องการให้
โปรแกรมหยุดทำงานทันทีเมื่อเกิด panic โดยสามารถใช้ recover
เพื่อดักจับและจัดการ panic ได้อย่างปลอดภัย
*/
