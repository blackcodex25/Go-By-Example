package main

import (
	"fmt"
)

/* การใช้ประเภทข้อมูลที่กำหนดเอง (custom
types) ในการสร้างข้อผิดพลาด (errors)ในภาษา Go โดยการ
ทำให้ประเภทนั้น ๆ รองรับ error interface ผ่านการสร้างเมธอด
Error() บนประเภทข้อมูลนั้น ตัวอย่างที่แสดงนี้เป็นการใช้
ประเภทข้อมูลที่กำหนดเองเพื่อระบุข้อผิดพลาดที่เกิดจากการรับค่า
อาร์กิวเมนต์ที่ไม่ถูกต้อง
*/
// กำหนด custom error type
type argError struct {
	arg     int
	message string
}

// ทำให้ argError รองรับ error interface ด้วยการสร้าง Error() method
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// ฟังก์ชัน f คืนค่า error ถ้า argument เป็น 42
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// เรียกใช้ฟังก์ชัน f ด้วยค่า argument 42
	_, err := f(42)

	// ตรวจสอบว่า error ที่ได้รับเป็นประเภท argError หรือไม่
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

/* logic และ โค้ด
กำหนดประเภทข้อมูลที่กำหนดเอง (argError):
argError เป็นโครงสร้าง (struct) ที่มีสองฟิลด์ คือ
arg และ message ที่เก็บข้อมูลเกี่ยวกับอาร์กิวเมนต์
ที่ทำให้เกิดข้อผิดพลาดและข้อความแสดงข้อผิดพลาด

สร้างเมธอด Error()
เมธอด Error() เป็นเมธอดที่จำเป็นต้องมีเพื่อทำให้
argError รองรับ error interface ได้
เมธอดนี้จะคืนค่าเป็นสตริงที่รวมข้อมูลจากฟิลด์ arg
และ message เพื่อบอกให้ทราบถึงรายละเอียดของข้อ
ผิดพลาด

ฟังก์ชัน f()
ฟังก์ชันนี้จะคืนค่าข้อผิดพลาดแบบ argError ถ้า
อาร์กิวเมนต์ที่รับเข้ามามีค่าเท่ากับ 42 โดยข้อผิดพลาดนี้
ถูกสร้างขึ้นด้วยการกำหนดค่า arg และ message ให้
กับ argError

ตรวจสอบข้อผิดพลาดใน main()
ในฟังก์ชัน main() มีการเรียกใช้ f(42) ซึ่งจะทำให้เกิดข้อผิดพลาด
ใช้ errors.As เพื่อตรวจสอบว่าข้อผิดพลาดที่คืนค่า
จาก f() ตรงกับประเภท argError หรือไม่
ถ้าใช่ จะพิมพ์ค่า arg และ message ของข้อผิดพลาดนั้น
ถ้าไม่ใช่ จะพิมพ์ข้อความบอกว่าไม่พบข้อผิดพลาดที่ตรงกับ argError
*/
/* สรุป
การสร้างข้อผิดพลาดที่กำหนดเอง
ใน Go เราสามารถสร้างประเภทข้อมูลที่กำหนดเอง (custom types) และทำให้
มันรองรับ error interface ได้โดยการเพิ่มเมธอด
Error() ซึ่งจะทำให้เราสามารถสร้างข้อผิดพลาดที่มีบริบทและรายละเอียดที่เฉพาะเจาะจง
การตรวจสอบข้อผิดพลาดด้วย errors.As
errors.As ช่วยให้เราสามารถตรวจสอบว่าข้อผิดพลาดใน
โซ่ของข้อผิดพลาดนั้นตรงกับประเภทข้อผิดพลาดที่กำหนด
เองหรือไม่ และสามารถแปลงข้อผิดพลาดนั้นให้เป็นประเภท
ที่เรากำหนดได้
*/
