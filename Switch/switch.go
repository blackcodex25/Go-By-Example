package main

import (
	f "fmt"
	"time"
)

/* การใช้ switch statements ในภาษา Go
ซึ่งใช้ในการตรวจสอบเงื่อนไขที่มีหลายสาขา (branches)
switch เป็นอีกวิธีหนึ่งที่ช่วยให้โค้ดดูเรียบร้อยและอ่านง่ายขึ้น
เมื่อเทียบกับการใช้ if-else หลายๆ ชั้น
*/
/* ประเภทต่างๆ ของ switch ใน Go
1.Basic Switch Statement
- ใช้ในการตรวจสอบค่าของตัวแปรกับแต่ละ case ถ้า
ตรงกับ case ใด ก็จะทำงานตามคำสั่งใน case นั้น

2.Multiple Expressions in Case
- สามารถใช้เครื่องหมาย , ใน case เพื่อแยกหลาย
ค่าที่อาจตรงกับเงื่อนไขเดียวกันได้ และสามารถใช้
default case เพื่อจัดการกับกรณีที่ไม่ตรงกับค่าใดๆ
ใน case

3.Switch Without Expression
switch แบบไม่มี Expression จะทำงานเหมือนกับ if-else หลายๆชั้น
โดยตรวจสอบเงื่อนไขต่างๆ ที่เป็น boolean

4.Type Switch
Type switch ใช้สำหรับการเปรียบเทียบชนิดข้อมูล
(type) ของค่าที่อยู่ใน interface แทนที่จะเปรียบเทียบค่าของมัน
ตัวแปรใน case จะมีชนิดข้อมูลตามที่ตรวจสอบ
*/
func main() {
	/* 1.Basic Switch
	ตัวแปร i มีค่าเป็น 2 และถูกนำไปตรวจสอบใน switch
	เมื่อ i ตรงกับ case 2 คำสั่ง
	f.Println("two") จะถูกเรียกใช้ และพิมพ์ เขียน 2 เป็นสอง
	*/
	i := 2
	f.Println("เขียน ", i, " เช่น ")

	switch i {
	case 1:
		f.Println("หนึ่ง")
	case 2:
		f.Println("สอง")
	case 3:
		f.Println("สาม")
	}

	/* 2.Multiple Expressions in Case
	ฟังก์ชัน time.Now().Weekday() จะส่งคืนค่าวันในสัปดาห์ปัจจุบัน
	(เช่น time.Saturday หรือ time.Sunday)
	ถ้าวันปัจจุบันตรงกับ Saturday หรือ Sunday จะ
	พิมพ์ มันเป็นวันหยุดสุดสัปดาห์
	ถ้าไม่ตรงกับทั้งสองวันนี้ (เช่นวันธรรมดา) จะ
	พิมพ์ วันธรรมดา
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		f.Println("วันหยุดสุดสัปดาห์")
	default:
		f.Println("วันธรรมดา")
	}

	/* 3.Switch Without Expression
	ใช้ switch แบบไม่มี expression เพื่อตรวจสอบว่า
	ปัจจุบันเป็นช่วงเช้าหรือบ่าย(t.Hour() คืนค่าเป็น
	ชั่วโมงในรูปแบบ 24 ชั่วโมง)
	ถ้าเวลาน้อยกว่า 12 จะพิมพ์ ก่อนเที่ยง
	ถ้าไม่ใช่ (เช่นช่วงบ่าย) จะพิมพ์ หลังเที่ยง
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		f.Println("ก่อนเที่ยง")
	default:
		f.Println("หลังเที่ยง")
	}

	/* 4.Type Switch
	ฟังก์ชัน whatAmI ใช้ type switch เพื่อแยกแยะ
	ชนิดข้อมูลที่รับเข้ามาผ่าน interface
	ถ้าค่าที่ส่งเข้ามาเป็น bool จะพิมพ์ ฉันคือบูลีน
	ถ้าเป็น int จะพิมพ์ ฉันคือจำนวนเต็ม
	ถ้าเป็นชนิดข้อมูลอื่นๆ จะพิมพ์ ไม่รู้ชนิดข้อมูล
	<type> โดย <type> คือชนิดข้อมูลที่ไม่รู้จัก
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			f.Println("ฉันคือบูลีน")
		case int:
			f.Println("ฉันคือจำนวนเต็ม")
		default:
			f.Printf("ไม่รู้ชนิดข้อมูล %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	// whatAmI(3.14) ทดสอบเพิ่มชนิดข้อมูล float64 ที่ไม่มีใน Switch
	// ผลลัพธ์จากการเพิ่มคือ "ไม่รู้ชนิดข้อมูล float64"
}

/* สรุป
โค้ดนี้แสดงให้เห็นการใช้งาน switch ที่หลากหลายในภาษา
Go ตั้งแต่การตรวจสอบค่าพื้นฐาน การจัดการเงื่อนไขหลายค่าใน
case เดียวกัน การใช้ switch เพื่อทดแทน if-else และ
การใช้ type switch เพื่อตรวจสอบชนิดข้อมูลของ interface
การใช้งานเหล่านี้ทำให้โค้ดมีความยืดหยุ่นและอ่านง่ายมากขึ้น
*/
