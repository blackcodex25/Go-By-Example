package main

import (
	b64 "encoding/base64" // การเข้ารหัสและถอดรหัสข้อมูลในรูปแบบ Base64
	"fmt"
)

// ฟังก์ชัน p ใช้สำหรับพิมพ์ผลลัพธ์ออกมา
func p(s ...interface{}) {
	fmt.Println(s...) // ใช้ fmt.Println เพื่อพิมพ์ผลลัพธ์
}

/*
Base64 Encoding เป็นวิธีการเข้ารหัสข้อมูลในรูปแบบของข้อความ ASCII
โดยแปลงข้อมูลไบนารีให้เป็นตัวอักษรที่สามารถส่งผ่านในโปรโตคอลต่างๆ ได้
เช่น HTTP, SMTP หรือ JSON
บทความนี้อธิบายว่า Go มีการรองรับการเข้ารหัสและถอดรหัสข้อมูลในรูปแบบ Base64 อยู่แล้ว
โดยใช้ package encoding/base64 ตัวอย่างโค้ดจะแสดงการเข้ารหัสและถอดรหัสข้อมูล
ทั้งในรูปแบบ Base64 มาตรฐาน และแบบที่ใช้ใน URL
*/
func main() {
	// ประกาศตัวแปร data ซึ่งเป็นข้อความที่เราต้องการเข้ารหัสและถอดรหัสด้วย Base64
	// ข้อความนี้ประกอบด้วยอักษร ตัวเลข และสัญลักษณ์พิเศษ
	data := "abc123!?$*&()'-=@~"
	// ใช้ฟังก์ชัน EncodeToString() จาก b64.StdEncoding เพื่อเข้ารหัสข้อมูลในรูปแบบ Base64 มาตรฐาน
	// ฟังก์ชันนี้รับค่าที่เป็น byte slice ดังนั้นต้องแปลง string data เป็น byte slice ก่อนด้วย []byte(data)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	// ผลลัพธ์จะถูกเก็บในตัวแปร sEnc และแสดงผลออกมา
	p(sEnc)

	// ถอดรหัสข้อมูลจาก Base64 กลับมาเป็นข้อมูลดั้งเดิมโดยใช้ ฟังก์ชัน DecodeString() จาก b64.StdEncoding
	// ค่าที่ได้จาก การถอดรหัสจะเป็น byte slice
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	// ดังนั้นต้องแปลงกลับเป็น string ด้วย string(sDec)
	p(string(sDec))
	p()

	// ส่วนนี้แสดงการเข้ารหัสและถอดรหัสโดยใช้ Base64 แบบ URL-compatible
	// ฟังก์ชัน URLEncoding เป็นการเข้ารหัสในรูปแบบที่เหมาะสำหรับการใช้ใน URL
	// (ซึ่งจะแทนที่เครื่องหมาย + ด้วย - และ / ด้วย _)
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}

/* สรุปเนื้อหาสำคัญ
Go รองรับการเข้ารหัสและถอดรหัสข้อมูลในรูปแบบ Base64 ผ่าน package encoding/base64
มีทั้งแบบมาตรฐาน (StdEncoding) และแบบ URL-compatible (URLEncoding)
การเข้ารหัสจะทำการแปลง string เป็น byte slice ก่อน จากนั้นจึ่งแปลงผลลัพธ์ Base64
กลับมาเป็น string ได้
Base64 ใช้ในหลายโปรโตคอลเพื่อส่งข้อมูลที่ปลอดภัยและสามารถใช้ได้
ในทุกสภาพแวดล้อม
*/
