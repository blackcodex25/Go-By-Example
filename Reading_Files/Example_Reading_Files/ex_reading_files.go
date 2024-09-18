package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const fileName = "./defer.txt" // ชื่อไฟล์ที่จะอ่าน

// ฟังก์ชันสำหรับตรวจสอบและจัดการข้อผิดพลาด
// หาก e เป็น nil จะคืนค่า nil แต่หาก e ไม่เป็น nil จะคืนค่า error
// ที่ประกอบด้วยข้อความ msg และข้อผิดพลาด e นั้นๆ
// โดยใช้ fmt.Errorf เพื่อสร้างข้อผิดพลาดใหม่
func check(e error, msg string) error {
	if e != nil {
		return fmt.Errorf("%s: %w", msg, e)
	}
	return nil
}

func p(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	// อ่านข้อมูลจากไฟล์ทั้งหมดในครั้งเดียวและแสดงผล
	dat, err := os.ReadFile(fileName)
	check(err, "Error reading files")
	p(string(dat))

	// เปิดไฟล์เพื่อใช้ในการอ่านทีละส่วน
	f, err := os.Open(fileName)
	check(err, "Error can't open read files")
	defer f.Close() // ปิดไฟล์เมื่อจบการทำงาน

	// อ่าน 5 ไบต์แรกจากไฟล์
	b1 := make([]byte, 5)
	n1, err := io.ReadAtLeast(f, b1, 5)
	check(err, "Error can't read bytes files")
	p("%d bytes: %s\n", n1, string(b1[:n1]))

	// ใช้ Seek เพื่อย้ายตำแหน่งไปที่ไบต์ที่ 6 และ อ่าน 2 ไบต์ถัดไป
	o2, err := f.Seek(6, io.SeekStart)
	check(err, "Error seeking file")
	b2 := make([]byte, 2)
	n2, err := io.ReadAtLeast(f, b2, 2)
	check(err, "Error reading file contents")
	p("%d bytes @ %d: %v\n", n2, o2, string(b2[:n2]))

	// กระโดดไปที่ตำแหน่งใหม่ในไฟล์
	_, err = f.Seek(4, io.SeekCurrent)
	check(err, "Error seeking 4 bytes forward")
	_, err = f.Seek(-10, io.SeekEnd)
	check(err, "Error reading remaining file contents")

	// อ่านข้อมูลอย่างน้อย 2 ไบต์จากตำแหน่งที่กำหนด
	o3, err := f.Seek(6, io.SeekStart)
	check(err, "Error seeking to beginning of file")
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err, "Error reading file contents")
	p("%d bytes @ %d: %v\n", n3, o3, string(b3))

	// ย้าย cursor กลับไปที่จุดเริ่มต้นของไฟล์
	_, err = f.Seek(0, io.SeekStart)
	check(err, "Error seeking to beginning of file")

	// ใช้ bufio เพื่ออ่านข้อมูลแบบมีบัฟเฟอร์
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(4)
	check(err, "Error reading file stats")
	p("4 bytes: %s\n", string(b4))
}

/* อธิบายหลักการทำงานของโค้ดแต่ละบรรทัด:
1.การนำเข้าของแพ็กเกจ: นำเข้าแพ็กเกจ bufio, fmt, io, และ os
เพื่อใช้งานในการอ่านไฟล์และจัดการอินพุต-เอาต์พุต

2.ฟังก์ชัน check: ใช้เพื่อตรวจสอบข้อผิดพลาด ถ้ามีข้อผิด
พลาดจะหยุดโปรแกรมทันทีโดยใช้ panic

3.อ่านไฟล์ทั้งหมดในครั้งเดียว:
os.ReadFile("/tmp/dat") อ่านเนื้อหาทั้งหมดของไฟล์
/tmp/dat และเก็บไว้ในตัวแปร dat

4.เปิดไฟล์: os.Open("/tmp/dat") เพื่อเปิดไฟล์และเก็บ
ตัวแปรที่เป็นชนิด os.File ไว้ใน f

5.การอ่านข้อมูล 5 ไบต์แรก: f.Read(b1) อ่าน 5 ไบต์แรก
จากไฟล์และเก็บใน b1, แสดงจำนวนไบต์ที่อ่านได้และ
เนื้อหาที่อ่านออกมา

6.การใช้ Seek เพื่อเปลี่ยนตำแหน่ง cursor: f.Seek(6, io.SeekStart)
ย้าย cursor ไปที่ไบต์ที่ 6 จากจุดเริ่มต้นของไฟล์

7.การอ่านข้อมูล 2 ไบต์ถัดไป: f.Read(b2) อ่าน 2 ไบต์
จากตำแหน่งที่ Seek ไว้

8.การกระโดด cursor ไปยังตำแหน่งใหม่: f.Seek(4, io.SeekCurrent)
และ f.Seek(-10, io.SeekEnd) เพื่อย้าย cursor ไปยังตำแหน่งต่างๆ
ตามที่ต้องการ

9.การอ่านข้อมูลด้วย ReadAtLeast: io.ReadAtLeast(f, b3, 2)
 อ่านข้อมูลจากไฟล์ให้ได้อย่างน้อย 2 ไบต์

10.การใช้ bufio เพื่ออ่านข้อมูล: ใช้ bufio.NewReader(f)
สร้าง buffered reader เพื่ออ่านข้อมูลจากไฟล์อย่างมีประสิทธิภาพ
จากนั้นใช้ Peek อ่าน 5 ไบต์แรก

11.ปิดไฟล์: หลังจากเสร็จสิ้นการใช้งานไฟล์, f.Close()
จะถูกเรียกเพื่อปิดไฟล์
*/
