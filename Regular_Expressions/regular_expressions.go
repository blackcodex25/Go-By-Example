package main

import (
	"bytes"  // ใช้งาน bytes แสดงผลข้อความที่เป็น byte
	"fmt"    // แสดงผลข้อความ
	"regexp" // ใช้งาน Regular Expressions
)

func main() {
	// จับคู่รูปแบบกับ string โดยใช้ MatchString
	// จับคู่รูปแบบ "p([a-z]+)ch" กับ "peach"
	// ผลลัพธ์จะเป็น true
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// จัดเก็บ Regular Expression ไว้ในตัวแปร r
	// โดยใช้ Compile จาก string "p([a-z]+)ch"
	r, _ := regexp.Compile("p([a-z]+)ch")

	// จับคู่รูปแบบกับ string โดยใช้ MatchString
	// จับคู่รูปแบบที่จัดเก็บไว้ใน r กับ "peach"
	// ผลลัพธ์จะเป็น true
	fmt.Println(r.MatchString("peach"))

	// จับคู่รูปแบบกับ string โดยใช้ MatchString
	// จับคู่รูปแบบที่จัดเก็บไว้ใน r กับ "peach punch"
	// ผลลัพธ์จะเป็น true
	fmt.Println(r.MatchString("peach punch"))

	// หาคำที่ตรงกับรูปแบบโดยใช้ FindString
	// หาคำแรกที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "peach punch"
	fmt.Println(r.FindString("peach punch"))

	// หาตำแหน่งที่ตรงกับรูปแบบโดยใช้ FindStringIndex
	// หาตำแหน่งแรกที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "peach punch"
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// หาคำที่ตรงกับรูปแบบโดยใช้ FindStringSubmatch
	// หาคำแรกที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "peach punch"
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// หาคำที่ตรงกับรูปแบบโดยใช้ FindAllString
	// หาคำทั้งหมดที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "peach punch pinch"
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// หาตำแหน่งที่ตรงกับรูปแบบโดยใช้ FindAllStringSubmatchIndex
	// หาตำแหน่งทั้งหมดที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "peach punch pinch"
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// หาคำที่ตรงกับรูปแบบโดยใช้ FindAllString
	// หาคำทั้งหมดที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "peach punch pinch" โดยจำกัดจำนวนผลลัพธ์เป็น 2
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// ตรวจสอบว่ามีการจับคู่รูปแบบกับ byte slice โดยใช้ Match
	// ตรวจสอบว่ารูปแบบที่จัดเก็บไว้ใน r ตรงกับ byte slice "peach"
	fmt.Println(r.Match([]byte("peach")))

	// จัดเก็บ Regular Expression ไว้ในตัวแปร r
	// โดยใช้ Compile จาก string "p([a-z]+)ch"
	r = regexp.MustCompile("p([a-z]+)ch")

	// แสดงผลรูปแบบที่จัดเก็บไว้ใน r
	fmt.Println("regexp:", r)

	// แทนที่คำที่ตรงกับรูปแบบด้วยสตริงอื่น
	// แทนที่คำที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน string "a peach" ด้วย "<fruit>"
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// แทนที่คำที่ตรงกับรูปแบบด้วยสตริงอื่น
	// โดยใช้ ReplaceAllFunc
	// แทนที่คำที่ตรงกับรูปแบบที่จัดเก็บไว้ใน r
	// ใน byte slice "a peach" ด้วยสตริงแปลงเป็น uppercase
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

/* ส่วนประกอบของโค้ด:
1.การทดสอบการจับคู่แบบง่าย (MatchString):
ฟังก์ชัน regexp.MatchString ใช้ตรวจสอบว่าสตริง "peach" ตรงกับ
รูปแบบที่กำหนดหรือไม่ ในที่นี้รูปแบบคือ "p([a-z]+)ch" ซึ่งหมายถึง
สตริงที่ขึ้นต้นด้วย p ตามด้วยอักขระตัวเล็กตั้งแต่ a-z อย่างน้อยหนึ่งตัว
และลงท้ายด้วย ch
match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
fmt.Println(match) // true

2.การคอมไพล์ regex (Compile):
ฟังก์ชัน regexp.Compile สร้างโครงสร้าง regex ที่สามารถใช้ซ้ำได้จาก
รูปแบบ "p([a-z]+)ch"
สามารถใช้วิธีการต่างๆ เช่น MatchString และ FindString กับโครงสร้างนี้
r, _ := regexp.Compile("p([a-z]+)ch")
fmt.Println(r.MatchString("peach")) // true

3.การค้นหาสตริงที่ตรงกับรูปแบบ (FindString):
ฟังก์ชัน FindString ค้นหาสตริงที่ตรงกับรูปแบบแรกที่พบในสตริงที่ระบุ
fmt.Println(r.FindString("peach punch")) // "peach"

4.การหาดัชนีของข้อความที่ตรงกับรูปแบบ (FindStringIndex):
ฟังก์ชัน FindStringIndex คืนค่าเป็นดัชนีเริ่มต้นและสิ้นสุดของการจับคู่ข้อความแรกที่พบ
fmt.Println("idx:", r.FindStringIndex("peach punch")) // [0, 5]

5.การหาข้อความและกลุ่มย่อย (FindStringSubmatch):
ฟังก์ชัน FindStringSubmatch คืนค่าข้อความที่ตรงกับรูปแบบทั้งหมด
และกลุ่มย่อย (เช่น ([a-z]+))
fmt.Println(r.FindStringSubmatch("peach punch")) // ["peach" "ea"]

6.การค้นหาดัชนีของข้อความและกลุ่มย่อย (FindStringSubmatchIndex):
ฟังก์ชันนี้คืนค่าดัชนีของข้อความที่จับคู่และกลุ่มย่อย
fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

7.การค้นหาทุกการจับคู่ (FindAllString):
ฟังก์ชัน FindAllString คืนค่าสตริงที่ตรงกับรูปแบบทั้งหมดในข้อความ
fmt.Println(r.FindAllString("peach punch pinch", -1)) // ["peach", "punch", "pinch"]

8.การจำกัดจำนวนการจับคู่ (FindAllString):
สามารถจำกัดจำนวนครั้งที่ต้องการจับคู่ได้โดยการระบุพารามิเตอร์ เช่น 2
fmt.Println(r.FindAllString("peach punch pinch", 2)) // ["peach", "punch"]

9.การใช้ข้อมูลแบบ byte (Match):
ฟังก์ชัน Match รองรับข้อมูลในรูปแบบ []byte แทนการใช้สตริง
fmt.Println(r.Match([]byte("peach"))) // true

10.การคอมไพล์ regex โดยใช้ MustCompile:
ฟังก์ชัน MustCompile ใช้คอมไพล์ regex และจะเกิด panic ถ้ามีข้อผิดพลาด
r = regexp.MustCompile("p([a-z]+)ch")
fmt.Println("regexp:", r) // "regexp: p([a-z]+)ch"

11.การแทนที่ข้อความ (ReplaceAllString):
ฟังก์ชัน ReplaceAllString แทนที่ข้อความที่ตรงกับรูปแบบด้วยข้อความใหม่
fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // "a <fruit>"

12.การแปลงข้อความด้วยฟังก์ชัน (ReplaceAllFunc):
ฟังก์ชัน ReplaceAllFunc แทนที่ข้อความที่จับคู่โดยใช้ฟังก์ชัน เช่น
bytes.ToUpper เพื่อแปลงเป็นตัวพิมพ์ใหญ่
in := []byte("a peach")
out := r.ReplaceAllFunc(in, bytes.ToUpper)
fmt.Println(string(out)) // "a PEACH"


โค้ดนี้สาธิตการใช้ regex ใน Go เพื่อทำการจับคู่ ค้นหา และแทนที่ข้อความตามรูป
แบบที่กำหนดอย่างมีประสิทธิภาพ
*/
