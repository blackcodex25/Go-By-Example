package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:   ", s.Count("test", "t"))
	p("HasPrefix   ", s.HasPrefix("test", "te"))
	p("HasSuffix   ", s.HasSuffix("test", "st"))
	p("Index:   ", s.Index("test", "e"))
	p("Join:   ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:   ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:   ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	parts := []string{"https://example.com", "users", "123", "profile"}
	url := s.Join(parts, "/")
	fmt.Println(url)
	// Output: https://example.com/users/123/profile

}

/* รายละเอียดของโค้ด:
1.การ Import แพ็กเกจ:
นำเข้าแพ็กเกจ fmt สำหรับการพิมพ์ผลลัพธ์ และย่อชื่อ
แพ็กเกจ strings เป็น s เพื่อทำให้โค้ดกระชับขึ้น

2.ตัวแปร p:
ตัวแปร p เป็นการย่อคำสั่ง fmt.Println เพื่อให้เรียกใช้ได้สะดวก

3.การใช้ฟังก์ชันจาก strings:
s.Contains(str, substr):
ตรวจสอบว่าสตริง str มีสตริงย่อย substr อยู่หรือไม่
ตัวอย่าง: s.Contains("test", "es") ผลลัพธ์คือ true

s.Count(str, substr):
นับจำนวนครั้งที่สตริงย่อย substr ปรากฏในสตริง str
ตัวอย่าง: s.Count("test", "t") ผลลัพธ์คือ 2

s.HasPrefix(str, prefix):
ตรวจสอบว่าสตริง str เริ่มต้นด้วยสตริง prefix หรือไม่
ตัวอย่าง: s.HasPrefix("test", "te") ผลลัพธ์คือ true

s.HasSuffix(str, suffix):
ตรวจสอบว่าสตริง str สิ้นสุดด้วยสตริง suffix หรือไม่
ตัวอย่าง: s.HasSuffix("test", "st") ผลลัพธ์คือ true

s.Index(str, substr):
คืนค่าตำแหน่งแรกที่พบสตริงย่อย substr ในสตริง str (หากไม่พบคืนค่า -1)
ตัวอย่าง: s.Index("test", "e") ผลลัพธ์คือ 1

s.Join(slice, sep):
รวมสตริงใน slice ด้วยตัวคั่น sep
ตัวอย่าง: s.Join([]string{"a", "b"}, "-") ผลลัพธ์คือ "a-b"

s.Repeat(str, count):
ทำซ้ำสตริง str จำนวน count ครั้ง
ตัวอย่าง: s.Repeat("a", 5) ผลลัพธ์คือ "aaaaa"

s.Replace(str, old, new, n):
แทนที่สตริง old ด้วย new ในสตริง str จำนวน n ครั้ง
(หาก n เป็น -1 จะแทนที่ทั้งหมด)
ตัวอย่าง:
s.Replace("foo", "o", "0", -1) ผลลัพธ์คือ "f00"
s.Replace("foo", "o", "0", 1) ผลลัพธ์คือ "f0o"

s.Split(str, sep):
แบ่งสตริง str ออกเป็น slice โดยใช้ sep เป็นตัวคั่น
ตัวอย่าง: s.Split("a-b-c-d-e", "-") ผลลัพธ์คือ ["a", "b", "c", "d", "e"]

s.ToLower(str):
แปลงสตริง str ให้เป็นอักษรตัวพิมพ์เล็กทั้งหมด
ตัวอย่าง: s.ToLower("TEST") ผลลัพธ์คือ "test"

s.ToUpper(str):
แปลงสตริง str ให้เป็นอักษรตัวพิมพ์ใหญ่ทั้งหมด
ตัวอย่าง: s.ToUpper("test") ผลลัพธ์คือ "TEST"
*/
/* สรุปหลักการทำงาน:
โค้ดนี้ใช้ฟังก์ชันต่างๆ จากแพ็กเกจ strings ในการทำงานกับสตริง
เช่น การตรวจสอบสตริง, การแทนที่สตริง, การแปลงตัวอักษร, การแบ่งสตริง
และอื่นๆ เป็นต้น
*/
