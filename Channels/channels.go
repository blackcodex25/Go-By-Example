package main

import (
	"fmt"
	"sync"
)

/*
	การใช้ Channels ในภาษา Go (Golang) เพื่อเชื่อม

โยง goroutines ที่ทำงานพร้อมกัน
คำอธิบายบทความ:
Channels คืออะไร?
*Channels คือ "ท่อ" ที่เชื่อม goroutines ที่ทำงานพร้อมกัน
(concurrent) เราสามารถส่งค่าไปยัง channel จาก
goroutine หนึ่งและรับค่าจาก channel นั้นในอีก goroutine
หนึ่งได้

การสร้าง Channel
*เราสามารถสร้าง channel ใหม่ด้วยการใช้ make(chan val-type)
*โดยที่ val-type คือประเภทของค่าที่
*channel จะส่งผ่าน

การส่งค่าไปยัง Channel
*ใช้ channel <- value เพื่อส่งค่าไปยัง channel ซึ่งใน
กรณีนี้เราส่งค่า "ping" ไปยัง channel ที่ชื่อ messages
จาก goroutine ใหม่

การรับค่าออกจาก Channel
*ใช้ <-channel เพื่อรับค่าจาก channel ซึ่งในกรณีนี้เราจะ
รับค่าข้อความ "ping" ที่ส่งไปและพิมพ์มันออกมา

การบล็อก (Blocking)
โดยค่าเริ่มต้น การส่งและการรับค่าจะบล็อก (block) จนกว่า
ทั้งผู้ส่งและผู้รับจะพร้อม ซึ่งหมายความว่าโปรแกรมจะรอ
จนกว่าจะมีการรับค่าจาก channel ก่อนจะดำเนินการต่อไป
*/
func main() {
	var wg sync.WaitGroup
	// สร้าง channels ชื่อ messages ที่สามารถส่งค่าได้เป็นประเภท string
	messages := make(chan string)
	wg.Add(1) // หยุดรอ Goroutine ซึ่งคือ 1
	// สร้าง goroutine ใหม่ ที่ส่งค่า "ping" ไปยัง channel messages
	go func() {
		defer wg.Done() // เลื่อนฟังก์ชัน anonymous functions
		// ออกไปจนกว่าฟังก์ชันหลักจะทำงานเสร็จ
		messages <- "ping"
	}()

	// รับค่าจาก channels messages และเก็บในตัวแปร msg
	msg := <-messages
	wg.Wait() // หยุดรอ goroutines ทั้งหมดทำงานเสร็จ
	// พิมพ์ค่าที่รับจาก channels
	fmt.Println(msg)
}

/* Logic ของโค้ด
สร้าง Channel:
messages := make(chan string)
สร้าง channel ที่รับค่าเป็นประเภท string

ส่งค่าไปยัง Channel
go func() { messages <- "ping" }()
สร้าง goroutine ใหม่ที่ส่งค่า "ping" ไปยัง channel
messages

รับค่าออกจาก Channel
msg := <-messages
รับค่าจาก channel messages และ เก็บค่าที่รับในตัวแปร msg

พิมพ์ค่า:
fmt.Println(msg)
พิมพ์ค่าที่ได้รับ (ซึ่งก็คือ "ping")
*/
/* การทำงานของโค้ด
Goroutine ใหม่จะถูกสร้างขึ้นและทำการส่งข้อความ "ping"
ไปยัง channel messages
main goroutine จะรอจนกว่าจะมีการส่งข้อความไปยัง
messages ก่อนที่มันจะรับค่าจาก channel และพิมพ์ข้อความนั้นออกมา

การทำงานนี้จะถูกบล็อกไปจนกว่าข้อความจะถูกส่งไปยัง channel และ
รับค่าจาก channel สำเร็จ ซึ่งช่วยให้การประสานงานระหว่าง
goroutines ทำได้อย่างมีประสิทธิภาพ
*/
