package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	// สร้างเส้นทางไฟล์โดยใช้ Join
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p) // ได้เส้นทางแบพกพาสำหรับระบบปฏิบัติการนั้นๆ

	// ตัวอย่างอื่นๆ ของการใช้ Join
	// ลบตัวคั่น (Separator) ที่ไม่จำเป็น
	fmt.Println(filepath.Join("dir//", "filename"))
	// ทำให้เส้นทางเป็นมาตรฐาน เช่นลบการเปลี่ยน Directory ที่ไม่จำเป็น
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// แยกส่วนของเส้นทางไฟล์เป็น Directory และ Base (ชื่อไฟล์)
	// คืนค่า directory
	fmt.Println("Dir(p)", filepath.Dir(p))
	// คืนค่าชื่อไฟล์
	fmt.Println("Base(p)", filepath.Base(p))

	// ตรวจสอบว่าเส้นทางเป็น Absolute Path หรือไม่
	// false เพราะไม่ใช่ Absolute Path
	fmt.Println(filepath.IsAbs("dir/file"))
	// true เพราะมี "/" ที่จุดเริ่มต้น
	fmt.Println(filepath.IsAbs("/dir/file"))

	// จัดการกับส่วนขยายของไฟล์
	filename := "config.json"
	ext := filepath.Ext(filename) // คืนค่านามสกุลไฟล์ (.json)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext)) // ตัดนามสกุลออกจากชื่อไฟล์

	// หา relative path ระหว่างสองเส้นทาง
	// คืนค่าเส้นทางความสัมพัทธ์ (Relative Path)
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println(rel)

	// คืนค่าเส้นทางสัมพัทธ์ อีกชุด
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println(rel)

}

/* สรุปหลักการทำงาน:
โค้ดนี้แสดงให้เห็นถึงการใช้ฟังก์ชันจากแพ็คเกจ filepath ในการสร้าง, ตรวจสอบ,
และจัดการเส้นทางไฟล์อย่างเป็นระบบ
ใช้ Join เพื่อสร้างเส้นทางไฟล์ที่พกพาระหว่างระบบปฏิบัติการ
ใช้ Dir และ Base ในการแยกเส้นทางและชื่อไฟล์
ใช้ Rel เพื่อค้นหาเส้นทางสัมพัทธ์ระหว่างสองเส้นทาง
Ext ใช้เพื่อแยกส่วนขยายของไฟล์ และ TrimSuffix ใช้เพื่อตัดนามสกุลออก
*/
