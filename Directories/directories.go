package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// WarpError รับข้อผิดพลาด e และข้อความ msg
// ถ้า e เป็น nil จะคืนค่า nil แต่หาก e ไม่เป็น nil
// จะคืนค่า error ใหม่ที่ประกอบด้วยข้อความ msg
// และข้อผิดพลาด e นั้นๆ ที่ถูกห่อซ้อนกัน
// โดยใช้ fmt.Errorf เพื่อสร้างข้อผิดพลาดใหม่
func WarpError(e error, msg string) error {
	if e != nil {
		return fmt.Errorf("%s: %w", msg, e)
	}
	return nil
}

// visit เป็นฟังก์ชันที่ใช้ในการเยี่ยมชมไฟล์หรือ
// ไดเร็กทอรีโดยใช้ filepath.WalkDir
// ฟังก์ชันนี้จะรับ path string, d fs.DirEntry
// และ err error เป็นอาร์กิวเมนต์
// และคืนค่า error หากมีข้อผิดพลาด
// โดยใช้ WarpError เพื่อห่อซ้อนข้อผิดพลาด
// และเพิ่มบริบทให้กับข้อผิดพลาด
func visit(path string, d fs.DirEntry, err error) error {
	// หากมีข้อผิดพลาดระหว่างการอ่านไฟล์หรือ
	// ไดเร็กทอรี ให้ใช้ WarpError เพื่อห่อซ้อน
	// ข้อผิดพลาดและเพิ่มบริบทให้กับข้อผิดพลาด
	if err := WarpError(err, fmt.Sprintf("Error reading %s", path)); err != nil {
		return err
	}

	// แสดงผลข้อมูลของไฟล์หรือไดเร็กทอรีที่อ่านได้
	// โดยใช้ fmt.Println และแสดงผล path และชื่อไฟล์/ไดเร็กทอรี
	fmt.Println("  ", path, d.Name())
	return nil
}

func main() {
	// สร้างไดเร็กทอรีชื่อ subdir และถ้ามีข้อผิดพลาดจะส่ง error ไปที่ WarpError
	err := os.Mkdir("subdir", 0755)
	WarpError(err, "Error creating directory")
	// ลบไดเร็กทอรี subdir เมื่อโปรแกรมทำงานเสร็จ
	defer os.RemoveAll("subdir")

	// สร้างไฟล์ว่างโดยการเขียนข้อมูลว่างลงในไฟล์ที่กำหนด
	createEmptyFile := func(name string) {
		d := []byte("")
		// ตรวจสอบข้อผิดพลาด
		WarpError(os.WriteFile(name, d, 0644), "Error writing file")
	}
	// สร้างไฟล์ว่างให้กับไดเร็กทอรีที่ชื่อว่า file1
	createEmptyFile("subdir/file1")
	// สร้างโครงสร้างไดเร็กทอรีลำดับชั้น subdir/parent/child และไฟล์ในแต่ละระดับ
	err = os.MkdirAll("subdir/parent/child", 0755)
	WarpError(err, "Error creating directory")

	createEmptyFile("subdir/parent/file1")
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/child/file1")

	// อ่านเนื้อหาภายในไดเร็กทอรี subdir/parent และแสดงรายการชื่อไฟล์และไดเร็กทอรี
	c, err := os.ReadDir("subdir/parent")
	WarpError(err, "Error reading directory")

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// เปลี่ยนไดเร็กทอรีปัจจุบันไปที่ subdir/parent/child และอ่านเนื้อหา
	err = os.Chdir("subdir/parent/child")
	WarpError(err, "Error changing directory")

	c, err = os.ReadDir(".")
	WarpError(err, "Error reading directory")
	// แสดงรายการเนื้อหาภายในไดเร็กทอรี subdir/parent/child
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// เปลี่ยนไดเร็กทอรีปัจจุบันไปที่ subdir
	err = os.Chdir("../../..")
	WarpError(err, "Error changing directory")

	fmt.Println("Visiting subdir") // แสดงผลการเรียกฟังก์ชัน visit
	// ใช้ WalkDir เพื่อเดินทางผ่านทุกไฟล์และไดเร็กทอรีใน subdir
	// และเรียกฟังก์ชัน Visit เพื่อจัดการแต่ละไฟล์
	err = filepath.WalkDir("subdir", visit)
}

/* สรุป
โค้ดนี้เป็นตัวอย่างที่ดีของการทำงานกับระบบไฟล์ใน Go โดยมีการสร้างไดเร็กทอรีและไฟล์
อ่านเนื้อหา, เปลี่ยนไดเร็กทอรี, และตรวจสอบข้อผิดพลาดอย่างละเอียด
*/
/* คำสั่ง err = os.Chdir("../../..")
ใน Go ใช้เพื่อเปลี่ยนไดเร็กทอรีปัจจุบัน (current working directory) ของโปรแกรม
ไปยังตำแหน่งที่สูงขึ้นในโครงสร้างไดเร็กทอรี โดย:
os.Chdir() เป็นฟังก์ชันที่ใช้เพื่อเปลี่ยนไดเร็กทอรีที่โปรแกรมกำลังทำงานอยู่
"../../.." เป็นเส้นทางสัมพัทธ์ (relative path) ที่ระบุให้
ย้อนกลับไปสามระดับจากตำแหน่งไดเร็กทอรีปัจจุบัน:
... หมายถึงไดเร็กทอรีแม่ (parent directory) หรือระดับบนของโครงสร้างไดเร็กทอรี
"../../.." หมายถึงย้อนกลับไปสามระดับขึ้นไปในโครงสร้างไดเร็กทอรี

ตัวอย่าง
ถ้าไดเร็กทอรีปัจจุบันคือ subdir/parent/child คำสั่ง os.Chdir("../../..")
จะพาโปรแกรมกลับไปที่ "." ซึ่งคือไดเร็กทอรีรากของโปรแกรม (root directory)
หรือระดับเดียวกับ subdir
*/
