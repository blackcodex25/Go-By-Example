package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

// IsAbsPath ตรวจสอบว่าเส้นทาง path เป็น absolute path หรือไม่
// absolute path คือเส้นทางที่ระบุถึงตำแหน่งของไฟล์หรือไดเรกทอรีอย่างครบถ้วน
// เช่น /home/user/file.txt หรือ C:\Users\user\file.txt
func IsAbsPath(path string) bool {
	// filepath.IsAbs ตรวจสอบว่าเส้นทาง path เป็น absolute path หรือไม่
	return filepath.IsAbs(path)
}
func main() {
	// ตัวอย่างการตรวจสอบว่าเส้นทาง path เป็น absolute path หรือไม่
	paths := []string{
		"/dir/file",   // Linux
		"C:/dir/file", // windows
	}

	// สร้างเส้นทางไฟล์โดยใช้ Join
	p := filepath.Join("dir1", "dir2", "filename")
	// แปลง \ เป็น / ในการสร้างเส้นทาง
	p = filepath.ToSlash(p) // ใช้ ToSlash เพื่อแปลงเส้นทางให้เป็นรูปแบบ Unix
	fmt.Println("p:", p)    // ได้เส้นทางแบพกพาสำหรับระบบปฏิบัติการนั้นๆ

	// ตัวอย่างอื่นๆ ของการใช้ Join
	// ลบตัวคั่น (Separator) ที่ไม่จำเป็น
	fmt.Println(filepath.ToSlash(filepath.Join("dir//", "filename")))
	// ทำให้เส้นทางเป็นมาตรฐาน เช่นลบการเปลี่ยน Directory ที่ไม่จำเป็น
	fmt.Println(filepath.ToSlash(filepath.Join("dir1/../dir1", "filename")))

	// แยกส่วนของเส้นทางไฟล์เป็น Directory และ Base (ชื่อไฟล์)
	// คืนค่า directory
	fmt.Println("Dir(p):", filepath.Dir(p))
	// คืนค่าชื่อไฟล์
	fmt.Println("Base(p):", filepath.Base(p))

	// ตรวจสอบว่าเส้นทางใน slice paths เป็น absolute path หรือไม่
	// ถ้าเป็น absolute path จะแสดงผลว่า true
	// ถ้าไม่เป็น absolute path จะแสดงผลว่า false
	for _, path := range paths {
		fmt.Printf("Is '%s' an absolute path? %v\n", path, IsAbsPath(path))
	}

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
	fmt.Println(filepath.ToSlash(rel))

	// คืนค่าเส้นทางสัมพัทธ์ อีกชุด
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println(filepath.ToSlash(rel))

}
