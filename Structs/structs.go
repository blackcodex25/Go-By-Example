package main

import (
	f "fmt"
)

/* การใช้งาน Structs ในภาษา Go
Structs ในภาษา Go เป็นโครงสร้างข้อมูลที่ประกอบด้วยหลาย
fields ที่มีประเภท (type) กำหนดไว้อย่างชัดเจน Structs
เหมาะสำหรับการจัดกลุ่มข้อมูลเพื่อสร้างเป็นระเบียน (records)

คุณสมบัติและการใช้งาน Structs
1.การสร้าง Struct
ในภาษา Go Struct เป็นการรวมกลุ่มของ fields ที่มี
ชนิดข้อมูล (type) ต่างกันหรือเหมือนกัน
ในตัวอย่างนี้ person เป็น struct ที่มี fields คือ name และ age

2.การสร้างฟังก์ชันที่คืนค่า Struct
ฟังก์ชัน newPerson ทำหน้าที่เป็น constructor function
ที่สร้างและคืนค่า person struct ตามชื่อที่ได้รับ (name)
ภาษา Go มีการจัดการหน่วยความจำแบบ garbage
collected language ซึ่งหมายความว่าเราสามารถคืนค่าตัวชี้
(pointer) ไปยังตัวแปรที่ถูกสร้างในฟังก์ชันได้อย่าง
ปลอดภัย เพราะตัวแปรจะไม่ถูกทำลายจนกว่าจะไม่มีการ
อ้างอิงถึงมัน

3.การสร้าง Struct ใหม่
เราสามารถสร้าง strcut ใหม่ได้ด้วยการกำหนดค่า fields ทันที
ถ้า fields ใดถูกละเว้น (omitted) ค่านั้นจะถูกกำหนด
เป็นค่าเริ่มต้น (zero-valued) ตามชนิดของ field นั้น

4.การใช้งาน Pointers กับ Structs
การใช้เครื่องหมาย & หน้าชื่อ struct จะให้ตัวช้ (pointer) ไปยัง struct นั้น
การเข้าถึง fields ของ struct ผ่านตัวชี้สามารถทำได้ โดยใช้ dot notation
ซึ่งตัวชี้จะถูก dereference โดยอัตโนมัติ

5.ความสามารถในการเปลี่ยนแปลงข้อมูล (Mutability)
Structs ใน Go สามารถถูกเปลี่ยนแปลงได้ (mutable)
ซึ่งหมายความว่าเราสามารถแก้ไขค่าของ fields ใน
struct ได้หลังจากที่มันถูกสร้างขึ้นแล้ว

6.การใช้ Anonymous Structs
ถ้า struct ถูกใช้งานเพียงครั้งเดียว เราสามารถประกาศ
และกำหนดค่าให้กับ struct นั้นได้โดยไม่ต้องให้ชื่อ (anonymous struct)
เทคนิคนี้มักถูกใช้ใน table-driven tests
*/
// ประกาศ struct ชื่อ person ที่มีฟิลด์คือ name ชนิดข้อมูล string
// และ age ชนิดข้อมูล int
// struct นี้ถูกใช้เพื่อเก็บข้อมูลที่เกี่ยวข้องกับบุคคล เช่น ชื่อและอายุ
type person struct {
	name string
	age  int
}

// ประกาศฟังก์ชัน newPerson สร้างพารามิเตอร์ name ชนิดข้อมูล string คืนค่า
// pointer person กลับไปยัง struct person
func newPerson(name string) *person {
	// ฟังก์ชันนี้สร้าง person ใหม่และตั้งค่าชื่อ name ตามที่ได้รับ
	p := person{name: name}
	p.age = 42 // age ถูกตั้งค่าเป็น 42 โดยอัตโนมัติ
	return &p  // ฟังก์ชันนี้คืนค่าตัวชี้ (pointer) ไปยัง person ที่ถูกสร้างขึ้น
}

func main() {
	// การสร้างและะการพิมพ์ค่า struct
	// สร้าง person ที่มี name เป็น "Bob" และ "age" เป็น 20
	f.Println(person{"Bob", 20})

	// สร้าง person โดยระบุฟิลด์ที่ต้องการโดยตรง
	f.Println(person{name: "Alice", age: 30})

	// สร้าง person ที่มีชื่อ Fred แต่ไม่ระบุ age ดังนั้น age จะถูกตั้งค่าเป็นค่าเริ่มต้น
	// (zero value) คือ 0
	f.Println(person{name: "Fred"})

	// สร้าง pointer ไปยัง person ที่มีชื่อ Ann และอายุ 40
	f.Println(&person{name: "Ann", age: 40})

	// สร้าง person ใหม่และคืนค่าตัวชี้ไปยัง person นั้น
	f.Println(newPerson("Jon"))

	// สร้าง person และเก็บไว้ในตัวแปร s
	s := person{name: "Sean", age: 50}
	// การเข้าถึงฟิลด์ของ struct ใช้ dot notation เช่น s.name
	f.Println(s.name)

	// สร้าง pointer ไปยัง strcut s
	sp := &s
	f.Println(sp.age)

	/* การเปลี่ยนแปลงค่าของ age ผ่าน pointer จะส่งผลให้ค่าจริงใน struct
	   ถูกเปลี่ยนแปลง */
	sp.age = 51
	f.Println(sp.age)

	// สร้าง anonymous struct โดยไม่ต้องกำหนดชื่อ struct
	dog := struct {
		name   string
		isGood bool
	}{
		// dog ถูกกำหนดค่าชื่อเป็น "Rex" และ isGood เป็น true
		"Rex",
		true,
	}
	f.Println(dog) // พิมพ์ผลลัพธ์ออกจอ
	// เทคนิคนี้มักใช้ในกรณีที่ struct ถูกใช้งานเพียงครั้งเดียวเช่นในการทดสอบ
}

/* สรุป
โค้ดนี้แสดงให้เห็นถึงการใช้งาน structs ในภาษา Go ตั้งแต่การ
ประกาศและการใช้งาน struct แบบธรรมดาไปจนถึงการใช้งาน
anonymous struct, Structs ช่วยในการจัดกลุ่มข้อมูลที่
เกี่ยวข้องกันและทำให้การจัดการข้อมูลมีความยืดหยุ่นมากขึ้น การใช้
pointers กับ structs ยังช่วยให้เราสามารถเปลี่ยนแปลงข้อมูล
ใน struct ได้อย่างมีประสิทธิภาพ
*/
