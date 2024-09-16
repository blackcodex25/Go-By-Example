package main

import (
	"fmt"
	"net/url"
	"os"
)

/*
การใช้ net/url ใน Go
เราสามารถใช้แพ็คเกจ net/url เพื่อแยกส่วนของ URL และเข้าถึงแต่ละส่วนได้
*/
func main() {
	// ตัวอย่าง URL
	rawURL := "https://user:pass@www.example.com:8080/path/to/resource?key=value#section"

	// แปลง URL ให้เป็นโครงสร้างข้อมูล
	// ใช้ url.Parse เพื่อแปลง URL ให้เป็นโครงสร้างข้อมูล
	parsedURL, err := url.Parse(rawURL)
	if err != nil { // ตรวจสอบว่า error ที่ได้รับเป็น nil หรือไม่
		fmt.Printf("error: %v\n", err) // แสดงข้อผิดพลาด ถ้าไม่สามารถแปลง URL ได้
		os.Exit(1)                     // ออกจากโปรแกรม
	}

	// แสดงส่วนต่างๆของ URL
	p("Scheme:", parsedURL.Scheme)     // ดึงโครงสร้างของ scheme จาก parsedURL
	p("User:", parsedURL.User)         // ดึงข้อมูลผู้ใช้จาก parsedURL
	p("Host:", parsedURL.Host)         // ดึงโดเมนเนมจาก parsedURL
	p("Path:", parsedURL.Path)         // ดึงที่อยู่ของไฟล์จาก parsedURL
	p("Query:", parsedURL.RawQuery)    // ดึงพารามิเตอร์เพิ่มเติมจาก parsedURL
	p("Fragment:", parsedURL.Fragment) // ดึงส่วนของเอกสารจาก parsedURL

	// ดึงข้อมูลผู้ใช้
	username := parsedURL.User.Username()    // ดึงชื่อผู้ใช้จาก parsedURL.User
	password, _ := parsedURL.User.Password() // ดึงรหัสผ่านจาก parsedURL.User
	p("Username:", username)                 // แสดงชื่อผู้ใช้
	p("Password:", password)                 // แสดงรหัสผ่าน

	// ดึงข้อมูลจาก Query parameters
	queryparams := parsedURL.Query()
	p("Query parameters:", queryparams)
	p("Query 'query' value:", queryparams.Get("query"))
}

// ฟังก์ชัน p ใช้สำหรับพิมพ์ผลลัพธ์ออกมา
// โดยใช้ fmt.Println เพื่อแสดงผลลัพธ์
// และ s เป็นพารามิเตอร์ที่รับค่า argument มาหลายตัว
func p(s ...interface{}) {
	fmt.Println(s...)
}

/* จากโค้ดด้านบน เราจะแยก URL ออกมาเป็นส่วนต่างๆ ได้ดังนี้
Scheme: https
User: user:pass
Host: www.example.com:8080
Path: /path/to/resource
Query: key=value
Fragment: section
*/
/* นอกจากนี้ เรายังความสามารถเข้าถึงค่า query paremeters เช่น
query=value โดยใช้ฟังก์ชัน Query() ของ url.URL ได้อีกด้วย
*/
/* สรุป
URL Parsing คือการแยกและวิเคราะห์ URL เพื่อนำข้อมูลจากส่วน
ประกอบต่างๆ ของ URL มาใช้งาน โดยใน Go แพ็คเกจ net/url
ช่วยให้การทำงานนี้ง่ายและมีประสิทธิภาพ
*/
