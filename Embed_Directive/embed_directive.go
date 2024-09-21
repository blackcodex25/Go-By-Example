package main

import (
	"embed"
	"fmt"
)

var (
	fileString string   // ฝังเนื้อหาของไฟล์ในรูปแบบ string
	fileByte   []byte   // ฝังเนื้อหาของไฟล์ในรูปแบบ byte array ([]byte)
	folder     embed.FS // ฝังหลายไฟล์หรือโฟลเดอร์โดยใช้ embed.FS
)

func main() {
	print(fileString)       // แสดงผลเนื้อหาของไฟล์จากตัวแปร fileString
	print(string(fileByte)) // แสดงผลเนื้อหาของไฟล์จากตัวแปร fileByte โดยแปลงเป็น string

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1)) // อ่านและแสดงผลเนื้อหาจากไฟล์ file1.hash

	// อ่านเนื้อหาของไฟล์ file1.hash จากโฟลเดอร์ folder และแปลงเป็น string เพื่อนำไปแสดงผล
	content2, _ := folder.ReadFile("folder/file1.hash")
	print(string(content2)) // อ่านและแสดงผลเนื้อหาจากไฟล์ file1.hash ซ้ำ
}

func print(args ...interface{}) {
	fmt.Println(args...)
}

/* รายละเอียดการทำงาน
1.print(fileString): แสดงเนื้อหาของไฟล์
single_file.txt ที่ถูกฝังในรูปแบบ string โดยใช้ตัวแปร fileString

2.print(string(fileByte)): แสดงเนื้อหาไฟล์เดียวกัน
(single_file.txt) แต่ใช้ข้อมูลจากตัวแปร fileByte
ซึ่งเป็น byte array โดยแปลงเป็น string ก่อนแสดงผล

3.folder.ReadFile("folder/file1.hash"):
อ่านเนื้อหาของไฟล์ file1.hash จากโฟลเดอร์ folder
และแปลงเป็น string เพื่อนำไปแสดงผล

4.การแสดงผลเนื้อหาซ้ำ:
โค้ดบรรทัดสุดท้ายอ่านและแสดงผลเนื้อหาไฟล์ file1.hash
อีกครั้ง ซึ่งทำให้ได้ผลลัพธ์เดิม
*/
/* สรุป
โค้ดนี้ใช้ความสามารถของ //go:embed เพื่อฝังไฟล์ลงใน binary
ของโปรแกรม ซึ่งสามารถใช้งานได้ทั้งในรูปแบบข้อความ และ byte array
นอกจากนี้ยังสามารถจัดการกับไฟล์หลายไฟล์ภายในโฟล์เดอร์เดียวกันผ่าน
ตัวแปร embed.FS
*/
