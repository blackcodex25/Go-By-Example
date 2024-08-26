package main

import (
	f "fmt"
)

/*
	การใช้ Channel Buffering ในภาษา Go (Golang)

ซึ่งช่วยให้เราสามารถส่งค่าผ่าน channel ได้แม้ว่าจะยังไม่มี
goroutine อื่นที่พร้อมรับค่าก็ตาม

คำอธิบายบทความ
Channel Unbuffered:
Channel เริ่มต้นเป็น unbuffered หมายความว่าการส่งค่า
( โดยใช้ chan <- ) จะถูกบล็อกจนกว่าจะมีการรับค่า ( <-chan )
ที่พร้อมรับค่าที่ถูกส่ง

Channel Buffered:
Buffered channels สามารถเก็บค่าจำนวนหนึ่งโดยไม่ต้องมี
การรับค่าตรงข้ามอยู่ในทันที Channels ที่มี buffer สามารถรับค่า
จำนวนจำกัดโดยไม่ต้องมี receiver พร้อมรับค่าทันที

การสร้าง Channel ที่มี Buffer:
ในตัวอย่างนี้ เราสร้าง channel ชนิด string ที่สามารถ
เก็บค่าได้ถึง 2 ค่า (buffer size = 2)

การส่งและรับค่าจาก Buffered Channel:
เนื่องจาก channel นี้มี buffer เราสามารถส่งค่าลงใน channel
ได้โดยไม่ต้องมี receiver พร้อมรับค่าในทันที
หลังจากนั้น เราสามารถรับค่าที่ถูกส่งเข้าไปใน channel ได้ตามปกติ
*/
func main() {
	// สร้าง channel ที่มี buffer ขนาด 2 ตัว
	messages := make(chan string, 2)

	// ส่งค่าลงใน channels
	messages <- "Buffered"
	messages <- "channels"

	// รับค่าจาก channels และพิมพ์ออกมา
	f.Println(<-messages)
	f.Println(<-messages)
}

/* Logic ของโค้ด
1.สร้าง Buffered Channel:
message := make(chan string, 2) สร้าง channel ที่
รับค่าประเภท string และมี buffer ขนาด 2 ค่า
หมายความว่าสามารถเก็บค่าที่ส่งไปได้ถึง 2 ค่า ก่อนที่ค่าจะถูกรับ

2.ส่งค่าไปยัง Channel:
messages <- "buffered" ส่งค่า "buffered" ลงใน
channel messages
messages <- "channel" ส่งค่า "channel" ลงใน
channel messages
เนื่องจาก channel มี buffer ขนาด 2 ค่า จึงสามารถส่งค่าทั้งสอง
ไปได้โดยไม่ต้องมีการรับค่าในทันที
*/
