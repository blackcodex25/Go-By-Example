package main

import (
	"fmt"
	"os"
)

/*
String Formatting คือการจัดรูปแบบสตริงใน Go โดยใช้รูปแบบ
การพิมพ์แบบ printf ซึ่ง Go มีการรองรับการจัดการสตริงได้อย่างยอดเยี่ยม
โดยมีการใช้ "verbs" หรือรูปแบบเฉพาะในการจัดรูปแบบค่าใน Go

ตัวอย่างการใช้งาน:
Printf: พิมพ์สตริงไปยัง os.Stdout
Sprintf: คืนค่าการจัดรูปแบบสตริงโดยไม่พิมพ์ออกมา
Fprintf: จัดรูปแบบและพิมพ์ไปยัง io.Writer อื่นๆ เช่น ไฟล์

สาระสำคัญ:
Go มีการจัดการสตริงอย่างยืดหยุ่นและมีความสามารถในการจัดรูป
แบบข้อมูลประเภทต่างๆ โดยใช้รูปแบบ (verbs) ซึ่งช่วยให้การ
แสดงผลข้อมูลในรูปแบบที่ต้องการเป็นไปอย่างง่ายดาย
*/
type point struct {
	X, Y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	fmt.Printf("struct2: %+v\n", p)

	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)

	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}

/* หลักการทำงานของโค้ด
1.Struct Formatting:
p := point{1, 2}
fmt.Printf("struct1: %v\n", p)
fmt.Printf("struct2: %+v\n", p)
fmt.Printf("struct3: %#v\n", p)
%v: พิมพ์ค่าเริ่มต้นของ struct point ซึ่งจะเป็น {1 2}
%+v: พิมพ์ค่า struct พร้อมชื่อฟิลด์ จะเป็น {x:1 y:2}
%#v: พิมพ์ค่า struct ในรูปแบบ Go syntax จะเป็น main.point{x:1, y:2}

2.Type and Boolean:
fmt.Printf("type: %T\n", p)
fmt.Printf("bool: %t\n", true)
%T: แสดงชนิดของตัวแปร p ซึ่งเป็น main.point
%t: แสดงค่าตัวแปร Boolean

3.Integer, Binary, Character, Hexadecimal:
fmt.Printf("int: %d\n", 123)
fmt.Printf("bin: %b\n", 14)
fmt.Printf("char: %c\n", 33)
fmt.Printf("hex: %x\n", 456)
%d: แสดงตัวเลขฐาน 10
%b: แสดงตัวเลขฐาน 2 (binary)
%c: แสดงตัวอักษรที่ตรงกับ ASCII code เช่น 33 คือ !
%x: แสดงตัวเลขในรูปแบบฐาน 16 (hexadecimal)

4.Floating Point and Scientific Notation:
fmt.Printf("float1: %f\n", 78.9)
fmt.Printf("float2: %e\n", 123400000.0)
fmt.Printf("float3: %E\n", 123400000.0)
%f: แสดงค่า float ในรูปแบบทศนิยม
%e: แสดงค่า float ในรูปแบบ scientific notation (lowercase e)
%E: แสดงค่า float ในรูปแบบ scientific notation (uppercase E)

5.String Formatting:
fmt.Printf("str1: %s\n", "\"string\"")
fmt.Printf("str2: %q\n", "\"string\"")
fmt.Printf("str3: %x\n", "hex this")
%s: แสดงสตริงปกติ
%q: แสดงสตริงพร้อมอัญประกาศคู่
%x: แสดงสตริงในรูปแบบตัวเลขฐาน 16 (hex)

6.Pointer:
fmt.Printf("pointer: %p\n", &p)
%p: แสดง pointer ของตัวแปร p

7.Width and Precision:
fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
ตัวเลขหลัง % คือกำหนดความกว้าง เช่น %6d คือความกว้าง 6 หลัก
%-6: จัดให้เป็น left-justify
%.2f: กำหนดทศนิยม 2 ตำแหน่ง

8.Sprintf and Fprintf:
s := fmt.Sprintf("sprintf: a %s", "string")
fmt.Println(s)
fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
Sprintf: จัดรูปแบบและส่งคืนค่าเป็นสตริง โดยไม่พิมพ์ออกมา
Fprintf: จัดรูปแบบและพิมพ์ออกไปยัง io.Writer อื่น
เช่น ในตัวอย่างนี้คือ os.Stderr (พิมพ์ไปที่ standard error)
*/
/* สรุป
โค้ดนี้เป็นตัวอย่างการใช้งานฟังก์ชัน fmt สำหรับการจัดรูปแบบค่าต่างๆ
เช่นตัวเลข, สตริง, struct และ pointer
โดยมีความยืดหยุ่นในการควบคุมการแสดงผล
*/
