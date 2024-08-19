package main

import (
	f "fmt"
	"maps"
)

func main() {
	// การสร้าง map และการกำหนดค่า key-value
	// สร้าง map ว่างชื่อ m ซึ่งมี key เป็น string และ value เป็น int
	m := make(map[string]int)

	// กำหนดค่าคู่ key-value ใน map m
	m["k1"] = 7
	m["k2"] = 13
	// พิมพ์ค่า map ที่มีคู่ key-value คือ map[k1:7 k2:13]
	f.Println("map:", m)

	// การเข้าถึงค่าใน map
	// เข้าถึงค่าใน map ที่ key เป็น "k1" และเก็บค่าในตัวแปร v1
	v1 := m["k1"]
	// พิมพ์ค่า v1 ซึ่งเป็น 7
	f.Println("v1:", v1)

	// พยายามเข้าถึงค่าใน map ที่ key เป็น k3 ซึ่งไม่มีอยู่ใน map
	// ดังนั้นจะคืนค่าเริ่มต้น (zero value) คือ 0
	v3 := m["k3"]
	// พิมพ์ค่า v3 ซึ่งเป็น 0
	f.Println("v3:", v3)

	// การหาจำนวนคู่ key-value ใน map
	// จะพิมพ์จำนวนคู่ key-value ที่มีอยู่ใน map m ซึ่งในที่นี้คือ 2
	f.Println("len:", len(m))

	// การลบคู่ key-value ใน map
	// ลบคู่ key-value ที่ key คือ "k2" ออกจาก map
	delete(m, "k2")
	// พิมพ์ค่า map ที่เหลือคู่ key-value คือ map[k1:7]
	f.Println("map:", m)

	// การลบคู่ key-value ทั้งหมดใน map
	clear(m)             // ลบคู่ key-value ทั้งหมดใน map m
	f.Println("map:", m) // พิมพ์ค่า map ว่าง map[]

	// การตรวจสอบว่า key มีอยู่ใน map หรือไม่
	// ตรวจสอบว่า key "k2" มีอยู่ใน map หรือไม่ โดยไม่สนใจค่า (_)
	_, prs := m["k2"]
	// พิมพ์ค่า prs ซึ่งในกรณีนี้จะเป็น false เพราะ "k2" ถูกลบไปแล้ว
	f.Println("prs:", prs)

	// การประกาศและกำหนดค่า map ในบรรทัดเดียว
	n := map[string]int{"foo": 1, "bar": 2}
	// พิมพ์ค่า map n ที่มีคู่ key-value คือ map[foo:1 bar:2]
	f.Println("map:", n)
	// สร้าง map n2 ที่มีค่าเหมือนกับ n
	n2 := map[string]int{"foo": 1, "bar": 2}
	// ตรวจสอบว่า map n และ n2 มีค่าเท่ากันหรือไม่
	if maps.Equal(n, n2) {
		// ถ้าเท่ากันจะพิมพ์ "n == n2"
		f.Println("n == n2")
	}

}

/* สรุป
โค้ดนี้แสดงให้เห็นถึงวิธีการใช้งาน maps ในภาษา Go ซึ่งเป็น
โครงสร้างข้อมูลที่ทรงพลังในการเก็บคู่ key-value โคดครอบคลุม
การสร้าง map การเพิ่มและลบข้อมูล การตรวจสอบ key และการ
เปรียบเทียบ maps, maps ใน Go ให้การใช้งานที่ยืดหยุ่นและมี
ฟังก์ชันเสริมจากแพ็กเกจ maps ที่ช่วยในการจัดการ maps ได้ง่าย
ขึ้น
*/
