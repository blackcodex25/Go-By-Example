package main

import (
	f "fmt"
	"unicode/utf8"
)

/* การใช้งาน String และ Runes ในภาษา Go
ในภาษา Go string เป็น slice ของ byte ที่สามารถอ่านได้
เท่านั้น (read-only) ภาษา Go และไลบรารีมาตรฐานมีการจัดการ string
เป็นพิเศษ เนื่องจากมันเป็นตัวเก็บข้อมูลของข้อความที่ถูกเข้ารหัสในรูปแบบ UTF-8
ในภาษาการเขียนโปรแกรมอื่นๆ string อาจประกอบด้วย Characters แต่ในภาษา Go
คำว่า Characters ถูกแทนที่ด้วย rune ซึ่งเป็น integer
ที่แสดงถึง Unicode code point
*/
/* คุณสมบัติและการใช้งาน Strings และ Runes
1.การจัดการ string ในภาษา Go
string ในภาษา Go เป็น slice ของ byte ที่ถูกเข้า
รหัสเป็น UTF-8 และเป็น read-only หมายความว่าไม่
สามารถเปลี่ยนแปลงค่าใน string ได้โดยตรง
string ถูกใช้เพื่อเก็บข้อมูลข้อความ (text) โดยแต่ละ
byte แสดงถึง rune หรือ Unicode code point

2.การเข้าถึงข้อมูลใน string
เนื่องจาก string ใน Go เป็น slice ของ byte การ
เข้าถึงค่าภายใน string จะเป็นการเข้าถึง byte raw
การวนลูป (loop) ผ่าน string เพื่อเข้าถึง byte แต่ละ
ตัวสามารถทำได้โดยใช้ดัชนี (index)

3.การนับจำนวน runes ใน string
ไลบรารี utf8 มีฟังก์ชัน RuneCountInString ที่สามารถ
นับจำนวน rune ใน string ได้
การนับจำนวน rune ต้องใช้เวลาในการทำงานตามขนาด
ของ string เนื่องจากต้องทำการถอดรหัส (decode)
UTF-8 rune ทีละตัว

4.การใช้งาน rune loop กับ string
เมื่อใช้ range loop กับ string, Go จะทำการถอดรหัส (decode) rune แต่ละตัวใน
string ให้โดยอัตโนมัติ และส่งกลับทั้งค่า rune และตำแหน่ง (index)
ของมันใน string

5.การใช้ฟังก์ชัน utf8.DecodeRuneInString
ฟังก์ชัน utf8.DecodeRuneInString สามารถใช้ถอดรหัส
rune จาก string ได้โดยตรง และยังสามารถระบุ
ตำแหน่งเริ่มต้น (index) ของ rune นั้นได้

6.การเปรียบเทียบ rune
ในภาษา Go ค่าที่อยู่ในเครื่องหมาย single quotes ('') เป็น rune literal
ซึ่งสามารถเปรียบเทียบกับค่า rune ที่ได้จาก string โดยตรง
*/
// การเปรียบเทียบ rune ในฟังก์ชัน examineRune
// ฟังก์ชัน examineRune รับ rune เป็นพารามิเตอร์และ
// เปรียบเทียบกับ rune literals

func examineRune(r rune) {
	// ถ้า rune เท่ากับ t
	if r == 't' {
		f.Println("found tee") // ถ้าใช่ พิมพ์ "found tee"
	} else if r == 'ส' { // ถ้า rune เท่ากับ 'ส'
		f.Println("found so sua") // พิมพ์ "found so sua"
	}
}

func main() {
	// การประกาศและใช้งาน string
	// ประกาศ string s ซึ่งมีค่าเป็นคำว่า สวัสดี ในภาษาไทย
	const s = "สวัสดี"
	// การใช้ len(s) จะคืนค่าความยาวของ string ในหน่วย
	// byte ไม่ใช่จำนวน rune (สำหรับ string "สวัสดี" จะมีความยาวเป็น 18 bytes)
	f.Println("Len:", len(s))

	// การวนลูปผ่าน string เพื่อเข้าถึง byte raw
	// ใช้ลูป for เพื่อเข้าถึงแต่ละ byte ใน string
	for i := 0; i < len(s); i++ {
		// พิมพ์ค่าในรูปแบบ hexadecimal (%x)
		f.Printf("%x ", s[i])
		// การวนลูปนี้แสดงให้เห็นถึงค่าของ byte raw ที่ประกอบ
		// เป็น UTF-8 code points ใน string "สวััสดี"
	}
	f.Println()

	// การนับจำนวน rune ใน string
	// ฟังก์ชัน utf8.RuneCountInString(s) จะคืนค่าจำนวน
	// rune ใน string s
	f.Println("Rune count:", utf8.RuneCountInString(s))

	// การใช้ range loop เพื่อวนลูปผ่าน rune ใน string
	// ใช้ range loop เพื่อวนลูปผ่าน string และถอดรหัส (decode)
	// rune แต่ละตัวพร้อมกับตำแหน่งเริ่มต้น (index)
	for idx, runeValue := range s {
		// พิมพ์ rune ในรูปแบบ Unicode("%#U") และตำแหน่งเริ่มต้นใน string
		f.Printf("%#U starts at %d\n", runeValue, idx)
	}

	f.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		// การใช้ฟังก์ชัน utf8.DecodeRuneInString
		// ใช้ utf8.DecodeRuneInString(s[i:]) เพื่อถอดรหัส
		// rune จาก string s โดยเริ่มจากตำแหน่ง i
		// ฟังก์ชันจะคืนค่า rune และความกว้าง (width) ของ rune นั้นในหน่วย byte
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		f.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// ใช้ฟังก์ชัน examineRune เพื่อเปรียบเทียบ rune ที่ได้รับ
		// และพิมพ์ข้อความที่เกี่ยวข้องออกมา
		examineRune(runeValue)
	}
}

/* สรุป
โค้ดนี้แสดงให้เห็นถึงการจัดการ strings และ rune ในภาษา Go
รวมถึงการทำงานกับ UTF-8 strings โดยเน้นการเข้าถึงและ
ถอดรหัส rune จาก string การนับจำนวน rune และการ
เปรียบเทียบ rune กับ rune literals การเข้าใจถึงการทำงาน
ของ string ในภาษา Go มีความสำคัญมาก เนื่องจากข้อความใน
Go ถูกเข้ารหัสในรูปแบบ UTF-8 ซึ่งมีผลต่อการทำงานและการ
จัดการข้อความที่ซับซ้อน
*/
