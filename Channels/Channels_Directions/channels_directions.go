package main

import (
	f "fmt"
)

/*การกำหนดทิศทางของ channels เมื่อใช้เป็น
พารามิเตอร์ในฟังก์ชันในภาษา Go (Golang) ซึ่งช่วยเพิ่มความ
ปลอดภัยของชนิดข้อมูล (type-safety) ในโปรแกรม โดยการ
กำหนดทิศทางสามารถระบุได้ว่าช่องทาง (channel) นั้นถูกใช้เพื่อ
การส่งหรือรับค่าเท่านั้น

คำอธิบายบทความ
Channel Directions (ทิศทางของ Channel):
เมื่อเราใช้ channels เป็นพารามิเตอร์ในฟังก์ชัน เรา
สามารถกำหนดได้ว่า channels นั้นจะถูกใช้สำหรับการ
ส่ง (send) หรือการรับ (receive) ค่าเท่านั้น ซึ่งจะช่วย
ให้โค้ดมีความปลอดภัยมากขึ้น (type-safety)

ฟังก์ชัน Ping:
ฟังก์ชัน ping ยอมรับ channel ที่ใช้สำหรับการส่งค่า
จากภายนอกเท่านั้น (ช่องทางสำหรับการส่งค่า chan<-)

ฟังก์ชัน Pong:
ฟังก์ชัน pong ยอมรับ channel หนึ่งสำหรับการรับค่า
(<-chan) และอีกช่องทางหนึ่งสำหรับการส่งค่า (chan<-)
*/
// ฟังก์ชัน ping ที่รับ channels สำหรับการส่งค่าเท่านั้น
func ping(pings chan<- string, msg string) {
	pings <- msg // ส่งข้อความ msg ลงใน channels ping
}

// ฟังก์ชัน pong ที่รับ channels สำหรับการรับค่า และอีก channels
// สำหรับการส่งค่า
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // รับข้อความจาก channel pings
	pongs <- msg   // ส่งข้อความที่รับไปยัง channel pongs
}

func main() {
	// สร้าง channel สองช่องทาง มี buffer ขนาด 1 ค่า
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// เรียกใช้ฟังก์ชัน ping เพื่อส่งข้อความไปยัง channel pings
	ping(pings, "passed message")

	// เรียกใช้ฟังก์ชัน pong เพื่อรับข้อความจาก pings และส่งต่อไปยัง pongs
	pong(pings, pongs)

	// รับข้อความจาก channel pongs และพิมพ์ออกมา
	f.Println(<-pongs)
}

/* Logic ของโค้ด
ฟังก์ชัน ping(pings chan<- string, msg string)
รับ channel ที่ใช้สำหรับการส่งค่า chan<- string และข้อความ msg string
ที่จะส่งลงใน channel นี้ ฟังก์ชันนี้จะทำหน้าที่ส่งข้อความลงใน
channel pings

ฟังก์ชัน pong(pings <-chan string, pongs chan<- string)
รับ channel หนึ่งที่ใช้สำหรับการรับค่า <-chan string
และอีก channel หนึ่งสำหรับ การส่งค่า chan<- string ฟังก์ชันนี้จะทำหน้าที่รับ
ข้อความจาก channel pings และส่งต่อข้อความนั้น ไปยัง channel pongs

ฟังก์ชัน main:
สร้าง channel สองช่อง คือ pings และ pongs ที่มี buffer ขนาด 1 ค่า
เรียกใช้ฟังก์ชัน ping เพื่อส่งข้อความ "passed message" ลงใน channel pings
เรียกใช้ฟังก์ชัน pong เพื่อรับข้อความจาก channel pings
และส่งต่อไปยัง channel pongs
รับข้อความจาก channel pongs และพิมพ์ข้อความนั้นออกมา
*/
/*การทำงานของโค้ด
ฟังก์ชัน ping จะส่งข้อความ "passed message" ลงใน channel pings
ฟังก์ชัน pong จะรับข้อความจาก channel pings แล้วส่งต่อไปยัง channel pongs
ใน main goroutine ข้อความที่ส่งไปจะถูกพิมพ์ออกมาจาก channel pongs

โค้ดนี้แสดงให้เห็นถึงการกำหนดทิศทางของ channels ที่เพิ่ม
ความชัดเจนและความปลอดภัยในโปรแกรมว่า channel ใดใช้
สำหรับการส่งและ channel ใดใช้สำหรับการรับ
ซึ่งช่วยลดข้อผิดพลาดที่อาจเกิดขึ้นจากการใช้งาน channels ผิดทิศทาง
*/
