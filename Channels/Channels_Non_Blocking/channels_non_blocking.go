package main

import (
	"fmt"
)

/*อธิบายเกี่ยวกับการดำเนินการกับ channels ในภาษา Go (Golang)
แบบ non-blocking ซึ่งปกติการส่งและรับค่าบน
channels ใน Go จะเป็นแบบ blocking นั่นหมายความว่าถ้าเรา
ส่งค่าผ่าน channel โดยที่ไม่มีตัวรับ หรือถ้าเรารับค่าจากบน
channels ใน Go จะเป็นแบบ blocking นั่นหมายความว่าถ้าเรา
ส่งค่าผ่าน channel โดยที่ไม่มีตัวรับ หรือถ้าเรารับค่าจาก
channel ที่ไม่มีการส่งค่าเข้ามา การทำงานจะหยุดรออยู่ที่จุดนั้น
จนกว่าจะมีตัวรับหรือตัวส่งค่าเข้ามา

คำอธิบายบทความ
การรับค่าบน Channel แบบ Non-Blocking:
ตัวอย่างแรกแสดงการรับค่าบน channel แบบ non-
blocking โดยใช้ select ถ้ามีค่าใน channel
messages select จะเลือก case ที่มี <-
messages เพื่อรับค่า ถ้าไม่มีค่าที่จะรับ select จะ
เลือก default case ทันที

การส่งค่าบน Channel แบบ Non-Blocking:
ตัวอย่างที่สองแสดงการส่งค่าบน channel แบบ non-blocking
ถ้า channel ไม่มี buffer และไม่มีตัวรับ
select จะเลือก default case ทันทีโดยไม่ทำการส่งค่า

การทำงานแบบ Multi-Way Non-Blocking Select:
เราสามารถใช้หลายๆ case พร้อมกันกับ default
clause เพื่อสร้างการทำงานแบบ non-blocking หลาย
ทิศทาง ตัวอย่างแสดงการพยายามรับค่าจากทั้ง
messages และ signals channels แบบ non-blocking
*/
/* คำอธิบายโค้ด

 */
func main() {
	messages := make(chan string, 1) // สร้าง channel สำหรับ string
	signals := make(chan bool)       // สร้าง channel สำหรับ bool

	messages <- "Raw Strings"

	// Non-blocking receive: ถ้ามีค่าใน channel messages จะเลือก case นี้
	select {
	case msg := <-messages:
		//time.Sleep(1 * time.Second)
		fmt.Println("Received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking send: ถ้ามีที่ว่างใน channel messages จะเลือก case นี้
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	// Multi-Way Non-blocking select: ตรวจสอบ channels หลายตัวพร้อมกัน
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

/* Logic ของโค้ด
การรับค่าแบบ Non-Blocking:
โค้ดนี้สร้าง channel ชื่อ messages สำหรับส่งข้อมูล
ชนิด string และใช้ select เพื่อรอรับค่าแบบ non-blocking

ถ้าไม่มีค่าให้รับ select จะเข้าสู่ default case
และพิมพ์ข้อความ "no message received"
*/
