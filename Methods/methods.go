package main

import (
	f "fmt"
)

/* การใช้งาน Methods กับ Structs ในภาษา Go
ภาษา Go รองรับการกำหนด methods บน struct types
Methods เป็นฟังก์ชันที่สามารถผูกกับ struct และช่วยใน
การดำเนินการต่างๆ ที่เกี่ยวข้องกับข้อมูลใน struct นั้น
*/
/* คุณสมบัติและการใช้งาน methods
*1.การกำหนด method บน struct
ฟังก์ชันสามารถถูกกำหนดเป็น method ของ struct โดยการเพิ่ม receiver type
เข้าไปในฟังก์ชันนั้น
*receiver type สามารถเป็นได้ทั้งแบบ pointer (*rect) หรือแบบ value (rect)

*2.การใช้งาน Pointer Receiver และ Value Receiver
การใช้ pointer receiver (*rect) เหมาะสำหรับ
การเปลี่ยนแปลงค่าใน struct โดยตรงภายใน method
หรือเพื่อหลีกเลี่ยงการคัดลอกข้อมูลใน struct เมื่อเรียกใช้ method
*การใช้ value receiver (rect) จะทำงานกับสำเนาของ struct ดังนั้น
*การเปลี่ยนแปลงที่เกิดขึ้นใน method จะไม่ส่งผลต่อค่าจริงใน struct

3.การเรียกใช้ Methods
Go สามารถจัดการการแปลงระหว่างค่าและตัวชี้
(value และ pointer) เมื่อเรียกใช้ method ได้
อย่างอัตโนมัติ ซึ่งหมายความว่าเราสามารถเรียกใช้
method ของ struct ได้ไม่ว่าเราจะมีตัวแปร
struct หรือ pointer ไปยัง struct
*/
/* โครงสร้าง rect ถูกกำหนดขึ้นเพื่อเก็บข้อมูลเกี่ยวกับ
ความกว้าง (width) และความสูง (height) ของรูป
สีเหลี่ยมผืนผ้า (rectangle)
*/
type rect struct {
	width, height int
}

// ฟังก์ชันนี้กำหนดให้เป็น method ของ rect โดยมี pointer receiver คือ *rect
func (r *rect) area() int {
	// area คำนวณพื้นที่ (area) ของรูปสี่เหลี่ยมโดยการ คูณความกว้างกับความสูง
	// (r.width * r.heightt)
	return r.width * r.height
	// การใช้ pointer receiver ทำให้สามารถเข้าถึงและเปลี่ยนแปลงค่าจริงใน
	// struct ได้
}

// ฟังก์ชันนี้กำหนดให้เป็น method ของ rect โดยมี value receiver คือ rect
/* perim คำนวณเส้นรอบรูป (perimeter) ของรูปสี่เหลี่ยมผืนผ้าโดยการคูณความกว้าง
และความสูงด้วย 2 และนำมาบวกกัน (2*r.width + 2*r.high)
การใช้ value receiver ทำให้ method ทำงานกับสำเนาของ struct
ไม่ใช่ค่าจริง
*/
func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	// สร้าง rect ชื่อ r ที่มีความกว้าง 10 และความสูง 5
	r := rect{width: 10, height: 5}

	// เรียกใช้ r.area() ซึ่งเรียก method area ผ่าน
	// ตัวแปร r Go จะจัดการการแปลง r เป็น pointer อัตโนมัติ
	f.Println("area: ", r.area())

	// เรียกใช้ r.perim() ซึ่งเรียก method perim
	// ผ่านตัวแปร r เนื่องจาก perim เป็น value receiver ไม่มีการแปลงใดๆ เกิดขึ้น
	f.Println("perim: ", r.perim())

	// สร้าง pointer ไปยัง r ชื่อ rp และเรียกใช้
	// rp.area() และ rp.perim() ซึ่ง Go จะจัดการ
	// การ dereference pointer อัตโนมัติ
	rp := &r
	f.Println("area: ", rp.area())
	f.Println("perim: ", rp.perim())
}

/* โค้ดนี้แสดงให้เห็นถึงการใช้งาน methods ในภาษา Go และ
การแยกการใช้งานระหว่าง pointer receiver และ value receiver
การใช้ pointer receiver ช่วยหลีกเลี่ยงการคัดลอกข้อมูลที่ไม่จำเป็นและ
อนุญาตให้เปลี่ยนแปลงข้อมูลใน struct ได้โดยตรง ขณะที่ value receiver
จะทำงานทั้งสองรูปแบบนี้ช่วยให้เราสามารถเลือกวิธีที่เหมาะสมที่สุดสำหรับการพัฒนาโปรแกรม
ของเราในภาษา Go
*/
