package main

import (
	f "fmt"
	"math"
)

/* การใช้งาน Interfaces ในภาษา Go
ในภาษา Go, interfaces คือคอลเลกชันที่มีชื่อของ method
signatures ที่กำหนดไว้อย่างชัดเจน interfaces ช่วยในการ
ออกแบบโปรแกรมที่สามารถทำงานกับประเภทข้อมูลหลายๆ ชนิด
ได้อย่างยืดหยุ่น โดยไม่ต้องรู้จักประเภทข้อมูลนั้นโดยตรง
ตราบใดที่ประเภทข้อมูลนั้นมีการนำ methods ที่กำหนดไว้ใน
interface ไปใช้งาน (implement)

คุณสมบัติและการใช้งาน Interfaces
1.ประกาศ Interface
interface ใน Go ประกอบด้วยคอลเลกชันของ
method signatures ที่กำหนดไว้ โดยไม่มีการนิยาม
โครงสร้างหรือการนำไปใช้งาน (implementation)
ของ methods เหล่านั้น
ตัวอย่างเช่น geometry คือ interface ที่ประกอบ
ด้วย methods สองตัวคือ area() และ perim()

2.การนำ Interface ไปใช้งาน
เพื่อ struct สามารถใช้งาน interface ได้
struct นั้นจำเป็นต้องมี medthods ทั้งหมดที่
กำหนดไว้ใน interface
ในตัวอย่างนี้ rect และ circle เป็น struct ที่
นำ interface geometry ไปใช้งานโดยการนิยาม
methods area() และ perim() สำหรับแต่ละ struct

3.การเรียกใช้ Methods ผ่าน Interface
เมื่อมีตัวแปรที่มีประเภทเป็น interface เราสามารถ
เรียกใช้ methods ที่กำหนดไว้ใน interface นั้นได้
โดยตรง ไม่ว่าตัวแปรนั้นจะเป็นชนิดใด ตราบใดที่ชนิด
นั้นนำ interface ไปใช้งาน (implements)

4.ความยืดหยุ่นในการใช้งาน
ฟังก์ชัน measure ในตัวอย่างนี้รับ geometry
interface เป็นพารามิเตอร์ นั่นหมายความว่าฟังก์ชัน
นี้สามารถทำงานกับข้อมูลชนิดใดก็ได้ตราบใดที่ข้อมูล
นั้นนำ geometry interface ไปใช้งาน เช่น rect
และ circle
*/
// การประกาศ Interface geometry
// geometry เป็น interface ที่ประกอบด้วย method
// signatures สองตัวคือ area() และ perim()
// ทั้งสอง methods นี้คืนค่าผลลัพธ์เป็นชนิด float64
type geometry interface {
	area() float64
	perim() float64
}

// การประกาศ Struct rect และ circle
// rect เป็น struct ที่มีฟิลด์คืือ width และ height (ทั้งสองเป็น float64)
type rect struct {
	width, height float64
}

// circle เป็น struct ที่มีฟิลด์คือ radius (เป็น float64)
type circle struct {
	radius float64
}

// การนำ Interface ไปใช้งานใน rect
// rect นำ Interface geometry ไปใช้งานโดย
// การนิยาม methods area() และ perim()
// area() คำนวณพื้นที่ของสี่เหลี่ยมผืนผ้า และ perim()
// คำนวณเส้นรอบรูปของสี่เหลี่ยมผืนผ้า
func (r rect) area() float64 {
	return r.width * r.height
}
func (c rect) perim() float64 {
	return 2*c.width + 2*c.height
}

// การนำ Interface ไปใช้งานใน circle
// เช่นเดียวกับ rect, circle นำ interface geometry ไปใช้งานโดยการ
// นิยาม methods area() และ perim()
// area() คำนวณพื้นที่ของวงกลม และ perim() คำนวณเส้นรอบวงของวงกลม
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// การสร้างฟังก์ชัน measure
/* ฟังก์ชันนี้รับพารามิเตอร์ชนิด geometry interface
ดังนั้นมันสามารถทำงานกับ rect, circle หรือ
ชนิดใดๆ ที่นำ geometry ไปใช้งาน
ฟังก์ชันนี้พิมพ์ข้อมูลของ struct พื้นที่ (area)
และเส้นรอบรูป (perimeter)
*/
func measure(g geometry) {
	f.Println(g)
	f.Println(g.area())
	f.Println(g.perim())
}
func main() {
	// สร้าง rect และ circle แล้วเรียกใช้ฟังก์ชัน measure กับแต่ละตัว
	// เนื่องจาก rect และ circle นำ geometry interface ไปใช้งาน
	// ฟังก์ชัน measure จึงสามารถทำงานกับทั้งสองชนิดข้อมูลได้
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

/* สรุป
โค้ดนี้แสดงให้เห็นถึงการใช้งาน interfaces ในภาษา
Go ที่ช่วยให้การเขียนโปรแกรมมีความยืดหยุ่นและสามารถขยายตัวได้ง่าย
การใช้ interfaces ช่วยให้โปรแกรมทำงานกับข้อมูลชนิดต่างๆ ที่มี methods
เหมือนกันได้โดยไม่ต้องรู้จักชนิดข้อมูลนั้นโดยตรงเป็นการแยกการนิยามโครงสร้าง
ข้อมูล (struct) ออกจากการดำเนินการ (methods) ทำให้โปรแกรมมี
ความยืดหยุ่นและการบำรุงรักษาง่ายขึ้น
*/
