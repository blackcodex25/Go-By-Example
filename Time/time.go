package main

import (
	"fmt"
	"time" // ใช้สำหรับการทำงานกับเวลาและวันที่ในภาษา Go
)

// ตัวอย่างการใช้งาน time
func main() {
	p := fmt.Println // ตัวแปร p เป็นการย่อคำสั่ง fmt.Println เพื่อให้เรียกใช้ได้สะดวก

	// รับเวลาปัจจุบัน
	now := time.Now()
	p(now) // แสดงเวลาปัจจุบัน

	// รับเวลาที่กำหนด
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then) // แสดงเวลาที่กำหนด

	// พิมพ์ประเภทของเวลา
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	//  พิมพ์วันในสัปดาห์
	p(then.Weekday())

	// เปรียบเทียบเวลา
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// คำนวณความแตกต่างของเวลา
	diff := now.Sub(then)
	p(diff) // คำนวณความแตกต่าง

	// ความแตกต่างของเวลาในการพิมพ์ในหน่วยต่างๆ
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// บวกและลบระยะเวลา
	p(then.Add(diff))  // บวก (บวกจากเวลาปัจจุบัน)
	p(then.Add(-diff)) // ลบ (ลบจากเวลาปัจจุบัน)
}

/* อธิบายทีละส่วน
ฟังก์ชันหลัก (main):
p := fmt.Println เป็นการสร้างตัวแปรย่อ p ที่ชี้ไปที่ฟังก์ชัน fmt.Println เพื่อทำให้โค้ดสั้นลง
เราสามารถใช้ p(...) แทน fmt.Println(...) ได้ในโค้ดถัดไป

การดึงเวลาปัจจุบัน:
time.Now() ใช้สำหรับดึงเวลาปัจจุบันในขณะนั้น และผลลัพธ์
จะถูกเก็บในตัวแปร now
จากนั้นใช้ p(now) เพื่อแสดงเวลาปัจจุบันนั้น

การสร้างวันที่และเวลาที่กำหนดเอง:
time.Date() สร้างเวลาและวันที่เฉพาะโดยระบุปี, เดือน, วัน, ชั่วโมง, นาที, วินาที, นาโนวินาที,
และเขตเวลา (ในกรณีนี้คือ time.UTC)
วันที่และเวลาที่สร้างขึ้นในตัวอย่างนี้คือ 17 พฤศจิกายน 2009, เวลา 20:34:58.651387237 UTC
จากนั้นใช้ p(then) เพื่อแสดงเวลาที่เรากำหนดเอง

เราสามารถดึงข้อมูลแต่ละส่วนของ then ได้ เช่น:
then.Year() ได้ปี (2009)
then.Month() ได้เดือน (พฤศจิกายน)
then.Day() ได้วัน (17)
then.Hour() ได้ชั่วโมง (20)
then.Minute() ได้นาที (34)
then.Second() ได้วินาที (58)
then.Nanosecond() ได้ค่านาโนวินาที (651387237)
then.Location() ได้เขตเวลา (UTC)

การดึงวันในสัปดาห์:
then.Weekday() ใช้เพื่อดึงชื่อวันของสัปดาห์ (ในที่นี้จะได้ Tuesday หรือวันอังคาร)

การเปรียบเทียบเวลา:
ฟังก์ชันเหล่านี้ใช้เปรียบเทียบเวลาระหว่าง then และ now:
then.Before(now) จะคืนค่า true ถ้า then อยู่ก่อน now
then.After(now) จะคืนค่า true ถ้า then อยู่หลัง now
then.Equal(now) จะคืนค่า true ถ้า then เท่ากับ now

การคำนวณระยะเวลาระหว่างสองเวลา:
now.Sub(then) คำนวณระยะเวลาที่ต่างกันระหว่าง now และ then
ผลลัพธ์ที่ได้จะเก็บอยู่ในตัวแปร diff ซึ่งเป็นค่าที่เป็นชนิด time.Duration
จากนั้น p(diff) จะแสดงความแตกต่างระหว่างสองเวลา เช่น "X ชั่วโมง X นาที X วินาที"

การแปลงระยะเวลาไปเป็นหน่วยต่างๆ:
ฟังก์ชันเหล่านี้ใช้เพื่อแสดงระยะเวลาในหน่วยต่างๆ:
diff.Hours() แปลงเป็นชั่วโมง
diff.Minutes() แปลงเป็นนาที
diff.Seconds() แปลงเป็นวินาที
diff.Nanoseconds() แปลงเป็นนาโนวินาที

การบวกและลบเวลาจาก then:
then.Add(diff) จะบวกระยะเวลา diff เข้าไปในเวลา then ซึ่งจะได้เวลาที่ตรงกับ now
then.Add(-diff) จะลบระยะเวลา diff จากเวลา then ซึ่ง
จะได้เวลาที่ห่างจาก then ย้อนกลับไปในอดีตเท่ากับ diff
*/
/* สรุป
โค้ดนี้แสดงวิธีการทำงานกับเวลาและวันที่ใน Go เช่น การดึงเวลาปัจจุบัน,
การสร้างวันที่และเวลาที่กำหนดเอง, การแยกส่วนข้อมูลจากวันที่,
การเปรียบเทียบเวลา, การคำนวณระยะเวลา, และการบวกหรือลบระยะเวลา

ฟังก์ชันในแพ็กเกจ time ช่วยให้เราสามารถจัดการกับเวลาได้
อย่างง่ายดายและมีประสิทธิภาพ
*/
