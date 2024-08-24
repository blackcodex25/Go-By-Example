package main

import (
	"errors"
	f "fmt"
)

/*บทความนี้กล่าวถึงการห่อข้อผิดพลาด (error)
เพื่อเพิ่มบริบทและการจัดการข้อผิดพลาดที่มีโครงสร้างดีขึ้น
*/
/* ในภาษา Go เราสามารถห่อข้อผิดพลาด (wrapping errors)
เพื่อเพิ่มบริบทให้กับข้อผิดพลาดนั้น ๆ ได้ โดยใช้ %w ใน
fmt.Errorf เพื่อสร้างข้อผิดพลาดที่มีการห่อซ้อนกัน ซึ่ง
สามารถสร้างโซ่ (chain) ของข้อผิดพลาดที่สามารถตรวจสอบได้
ด้วยฟังก์ชันอย่าง errors.Is และ errors.As ฟังก์ชันเหล่า
นี้ช่วยให้เราสามารถระบุประเภทของข้อผิดพลาดหรือ sentinel
errors ในโซ่ของข้อผิดพลาดได้
*/
// ฟังก์ชัน f จะคืนค่า error เมื่อ argument มีค่าเท่ากับ 42
func F(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// ฟังก์ชัน f จะคืนค่า error เมื่อ argument มีค่าเท่ากับ 42
var ErrOutoftea = f.Errorf("no more tea availble")
var ErrPower = f.Errorf("can't boil water")

// ฟังก์ชัน maketea คืนค่า error เมื่อเจอปัญหาที่กำหนด
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutoftea
	} else if arg == 4 {
		// ห่อ error ด้วยบริบทเพิ่มเติม
		return f.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	// ตัวอย่างการใช้ฟังก์ชัน f
	for _, i := range []int{7, 42} {
		if r, e := F(i); e != nil {
			f.Println("f failed:", e)
		} else {
			f.Println("f worked:", r)
		}
	}

	// ตัวอย่างการใช้ฟังก์ชัน makeTea และตรวจสอบข้อผิดพลาด
	for i := range 5 {
		if err := makeTea(i); err != nil {
			// ใช้ errors.Is เพื่อตรวจสอบว่าข้อผิดพลาดเป็นชนิดไหน
			if errors.Is(err, ErrOutoftea) {
				f.Println("We Should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				f.Println("Now it is dark")
			} else {
				f.Printf("unknown error: %s\n", err)
			}
			continue
		}
		f.Println("Tea is Ready!")
	}

}

/* สรุป
การห่อข้อผิดพลาด:
ใช้ %w ใน fmt.Errorf เพื่อเพิ่ม
บริบทให้กับข้อผิดพลาดและสร้างโซ่ของข้อผิดพลาด

การตรวจสอบข้อผิดพลาดที่ห่อซ้อนกัน:
ใช้ errors.Is เพื่อตรวจสอบว่าข้อผิดพลาดในโซ่เป็นชนิดใด
ทำให้สามารถจัดการกับข้อผิดพลาดได้อย่างยืดหยุ่นและมี
ประสิทธิภาพมากขึ้น

การใช้งาน sentinel errors:
ควรใช้ sentinel errors เพื่อ
ระบุสถานการณ์ที่เฉพาะเจาะจงและตรวจสอบข้อผิดพลาดที่
เจาะจงนั้นได้อย่างง่ายดาย

โดยรวมแล้ว เทคนิคนี้ช่วยให้โค้ดสามารถจัดการกับข้อผิดพลาด
ได้อย่างมีระบบและง่ายต่อการดูแลรักษา
*/
