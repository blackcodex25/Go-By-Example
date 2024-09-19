package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error, msg string) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, e)
}

func p(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	// สร้าง byte slice ที่มีึค่าเป็นข้อความ "hello\ngo\n"
	d1 := []byte("hello\ngo\n")
	// ใช้ os.WriteFile เพื่อเขียนข้อมูลนี้ลงในไฟล์ /defer.txt โดยตั้ง permission เป็น 0644
	err := os.WriteFile("elevoc_dnn_kernel.txt", d1, 0644)
	// ตรวจสอบความผิดพลาด
	check(err, "Error writing file")

	f, err := os.Create("elevoc_dnn_kernel.txt") // สร้างไฟล์ /defer.txt
	check(err, "Error creating file")            //  ตรวจสอบความผิดพลาด

	defer f.Close() // ปิดไฟล์เมื่อจบการทำงาน

	// เขียนชนิดข้อมูล byte slice d2 ที่ประกอบด้วยตัวเลข ASCII ซึ่งแทนคำว่า some\n ลงในไฟล์
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err, "Error writing to file")
	p("wrote %d bytes", n2)

	// ใช้ WriteString เพื่อเขียนสตริง "writes\n" ลงในไฟล์
	n3, err := f.WriteString("writes\n")
	check(err, "Error writing string to file")
	p("wrote %d bytes", n3)

	f.Sync() // บันทึกข้อมูลทั้งหมดลงดิสก์ทันที

	w := bufio.NewWriter(f) // สร้าง writer แบบบัฟเฟอร์
	// เขียนข้อมูล "buffered\n" ลงในบัฟเฟอร์ก่อนจะเรียก
	// w.Flush() เพื่อบันทึกข้อมูลทั้งหมดลงดิสก์
	n4, err := w.WriteString("buffered\n")
	check(err, "Error writing string to buffer")
	p("wrote %d bytes", n4)
	w.Flush()
}

/* Note: 0644 หมายความว่าเจ้าของไฟล์สามารถอ่านและเขียนไฟล์ได้
ขณะที่กลุ่มและผู้ใช้อื่นๆ สามารถอ่านไฟล์ได้เท่านั้น
*/
/* ผลลัพธ์ที่คาดหวัง:
โปรแกรมจะเขียนข้อมูลลงในไฟล์ทั้งสองไฟล์ /tmp/dat1 และ /tmp/dat2
โดยจะมีการแสดงจำนวนไบต์ที่ถูกเขียนลงไฟล์สำหรับแต่ละส่วน
*/
