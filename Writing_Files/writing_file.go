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
	err := os.WriteFile("/defer1.txt", d1, 0644)
	// ตรวจสอบความผิดพลาด
	check(err, "Error writing file")

	f, err := os.Create("/defer1.txt") // สร้างไฟล์ /defer.txt
	check(err, "Error creating file")  //  ตรวจสอบความผิดพลาด

	defer f.Close() // ปิดไฟล์เมื่อจบการทำงาน
	// เขียนชนิดข้อมูล byte
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err, "Error writing to file")
	p("worte %d bytes", n2)

	n3, err := f.WriteString("writes\n")
	check(err, "Error writing string to file")
	p("wrote %d bytes", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err, "Error writing string to buffer")
	p("wrote %d bytes", n4)
	w.Flush()
}

/* Note: 0644 หมายความว่าเจ้าของไฟล์สามารถอ่านและเขียนไฟล์ได้
ขณะที่กลุ่มและผู้ใช้อื่นๆ สามารถอ่านไฟล์ได้เท่านั้น
*/
