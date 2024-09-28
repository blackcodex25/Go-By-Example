package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

/* การทำงานของ HTTP Client ใน Go สามารถใช้งานแพ็กเกจ net/http
สำหรับการส่ง HTTP request ได้ง่ายๆ โดยใช้ฟังก์ชัน http.Get ซึ่งเป็น shortcut
ในการสร้าง HTTP request แบบ GET โดยไม่ต้องสร้าง object ของ http.Client เอง: */
// Request: ส่งคำขอ HTTP แบบ GET ไปยังเซิร์ฟเวอร์โดยใช้ http.Get()
// Status: พิมพ์สถานะของการตอบสนอง (response.Status)
// Body: อ่านและแสดงผลบรรทัดแรกของเนื้อหาตอบสนอง (response.Body)
// โดยสามารถใช้ io.ReadAll หรืออ่านทีละบรรทัดตามความต้องการ

func main() {
	// ส่งคำขอ HTTP GET ไปยังเว็บไซต์ gobyexample.com
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		log.Printf("Error Response status: %v", err) // แสดงข้อผิดพลาดถ้ามี
		return
	}
	defer resp.Body.Close() // ปิด Body ของการตอบสนองเมื่อใช้งานเสร็จ

	// แสดงสถานะของการตอบสนอง HTTP
	fmt.Println("Response status:", resp.Status)

	// อ่านเนื้อหาจาก Body ของการตอบสนองโดยใช้ Scanner
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text()) // แสดงบรรทัดเนื้อหาทีละบรรทัด
	}

	// ตรวจสอบว่ามีข้อผิดพลาดในการอ่านหรือไม่
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading response: %v", err) // แสดงข้อผิดพลาดถ้ามี
		return
	}
}

/* หลักการทำงาน */
/* 1.ส่งคำขอ HTTP ไปยัง URL ที่ระบุโดยใช้ http.Get() */
/* 2.ตรวจสอบว่ามีข้อผิดพลาดในการ Get หรือไม่
ถ้ามีก็พิมพ์ข้อความิดพลาดแล้วออกจากฟังก์ชัน */
/* 3.ใช้ defer เพื่อให้แน่ใจว่า resp.Body จะถูกปิดหลังจากใช้งานเสร็จ */
/* 4.แสดงสถานะการตอบสนอง HTTP (resp.Status) */
/* 5.ใช้ bufio.NewScanner อ่านบรรทัดเนื้อหาจาก resp.Body
และแสดงบรรทัดแรก 5 บรรทัด */
/* 6.ตรวจสอบข้อผิดพลาดจาก scanner เมื่ออ่านเนื้อหา */
