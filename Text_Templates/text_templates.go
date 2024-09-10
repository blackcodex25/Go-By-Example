package main

import (
	"os"
	"text/template"
)

/*
โค้ดนี้ใช้แพ็กเกจ text/template
ในการสร้างและจัดการเทมเพลตเพื่อแสดงข้อมูลแบบไดนามิกใน Go โดยมีการใช้เทคนิค
การสร้างเทมเพลตที่หลากหลายและการจัดรูปแบบข้อมูลตามเงื่อนไขต่างๆ
*/
// ส่วนนี้ใช้ในการสร้างและจัดการเทมเพลต
func main() {
	// ส่วนนี้ใช้ในการสร้างเทมเพลต t1
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		// ถ้ามี error ให้ panic
		panic(err)
	}
	// ส่วนนี้ใช้ในการสร้างเทมเพลต t1 ด้วย template.Must
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	// ส่วนนี้ใช้ในการแสดงข้อมูลตามรูปแบบที่กำหนดใน t1
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// ส่วนนี้ใช้ในการสร้างฟังก์ชันช่วยเหลือสำหรับการสร้างเทมเพลต
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// ส่วนนี้ใช้ในการสร้างเทมเพลต t2
	t2 := Create("t2", "Name: {{.Name}}\n")

	// ส่วนนี้ใช้ในการแสดงข้อมูลตามรูปแบบที่กำหนดใน t2
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// ส่วนนี้ใช้ในการสร้างเทมเพลต t3
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")

	// ส่วนนี้ใช้ในการแสดงข้อมูลตามรูปแบบที่กำหนดใน t3
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// ส่วนนี้ใช้ในการสร้างเทมเพลต t4
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")

	// ส่วนนี้ใช้ในการแสดงข้อมูลตามรูปแบบที่กำหนดใน t4
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
}

/* ส่วนประกอบหลักของโค้ด
1.การสร้างและจัดการเทมเพลตพื้นฐาน
t1 := template.New("t1")
t1, err := t1.Parse("Value is {{.}}\n")
if err != nil {
    panic(err)
}

t1 = template.Must(t1.Parse("Value: {{.}}\n"))

ส่วนนี้สร้างเทมเพลต t1 โดยใช้ข้อความ "Value: {{.}}" ซึ่ง {{.}}
คือการแทนค่าที่ส่งมาทางฟังก์ชัน Execute เมื่อตัวเทมเพลตถูกใช้งาน

2.การส่งค่าให้กับเทมเพลต
t1.Execute(os.Stdout, "some text")
t1.Execute(os.Stdout, 5)
t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
ส่งค่า string, int, และ slice ของ string เพื่อให้
เทมเพลตแสดงผลตามรูปแบบที่กำหนดใน t1

3.การสร้างฟังก์ชันช่วยเหลือ (Helper Function)
Create := func(name, t string) *template.Template {
    return template.Must(template.New(name).Parse(t))
}
ฟังก์ชัน Create ถูกสร้างเพื่อช่วยในการสร้างเทมเพลตได้สะดวกขึ้น
โดยรับชื่อและเนื้อหาของเทมเพลต และใช้ template.Must เพื่อตรวจสอบข้อผิดพลาด

4.การใช้งานเทมเพลตกับ struct และ map
t2 := Create("t2", "Name: {{.Name}}\n")
t2.Execute(os.Stdout, struct { Name string }{"Jane Doe"})
t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})
เทมเพลต t2 ใช้การแทนที่ด้วยค่า Name จาก struct และ map
ผ่าน {{.Name}} ซึ่งสามารถนำค่าจากทั้ง struct และ map มาใช้ได้

5.การใช้เงื่อนไข (if/else)
t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
t3.Execute(os.Stdout, "not empty")
t3.Execute(os.Stdout, "")
ใช้การตรวจสอบเงื่อนไข if/else โดยถ้ามีค่าเป็น nil หรือค่าเริ่มต้น
เช่น สตริงว่าง จะพิมพ์ no มิฉะนั้นจะพิมพ์ yes

6.การวนซ้ำ (range)
t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
ใช้คำสั่ง range เพื่อวนซ้ำใน slice ของ string และแสดงแต่ละค่าต่อกัน
*/
/* สรุป
โค้ดนี้แสดงตัวอย่างการใช้งานเทมเพลตใน Go เพื่อแสดงข้อมูล
แบบไดนามิกผ่านการใช้ struct, map, การใช้เงื่อนไข และการวนซ้ำ
*/
