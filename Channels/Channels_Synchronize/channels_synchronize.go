package main

import (
	f "fmt"
	"time"
)

/* การใช้ channels ในภาษา Go (Golang) เพื่อทำการ
ซิงโครไนซ์ (synchronize) การทำงานระหว่าง goroutines โดยเฉพาะ
การใช้ channel เพื่อรอให้ goroutine หนึ่งทำงานเสร็จแล้วจึงดำเนิน
การต่อใน goroutine อื่นๆ

คำอธิบายบทความ
การซิงโครไนซ์การทำงานของ Goroutines:
เราสามารถใช้ channels เพื่อซิงโครไนซ์การทำงานระหว่าง
goroutines เช่น การรอให้ goroutine หนึ่งทำงานเสร็จแล้ว
จึงดำเนินการต่อ

ตัวอย่างการรอการทำงานของ Goroutine:
ในตัวอย่างนี้ เราจะใช้การรับค่าจาก channel เพื่อรอให้
goroutine หนึ่งทำงานเสร็จแล้วจึงดำเนินการต่อไป

การใช้ done Channel:
ฟังก์ชันที่ทำงานใน goroutine จะส่งค่าไปยัง channel ชื่อ
done เพื่อแจ้งให้ goroutine อื่นทราบว่างานเสร็จสิ้นแล้ว

การบล็อกการทำงานจนกว่าจะได้รับการแจ้งเตือน:
ในโค้ดนี้ เราจะบล็อกการทำงานของ main goroutine
จนกว่าจะได้รับค่าจาก channel done ซึ่งเป็นการยืนยันว่า
goroutine อื่นทำงานเสร็จสิ้นแล้ว

ถ้าลบ <- done ออก
ถ้าเราลบบรรทัดที่มี <- done ออก โปรแกรมจะจบการ
ทำงานก่อนที่ worker goroutine จะเริ่มทำงานเสียอีก
เพราะไม่มีการรอให้ worker ทำงานเสร็จ
*/
// ฟังก์ชัน worker ที่ทำงานใน goroutine
func worker(done chan bool) {
	f.Println("Working...")
	// หน่วงเวลา 1 วินาที
	time.Sleep(time.Second)
	f.Println("done")

	// ส่งค่าลงใน channel เพื่อบอกว่าทำงานเสร็จแล้ว
	done <- true

}
func main() {
	// สร้าง buffered channel ที่มีขนาด bufer เท่ากับ 1
	done := make(chan bool, 1)

	// เรียกใช้ฟังก์ชัน worker ใน goroutine ใหม่
	go worker(done)

	// บล็อการทำงานของ main goroutines
	// จนกว่าจะได้รับค่าใน channels done
	<-done
}

/* Logic ของโค้ด
ฟังก์ชัน worker:
worker(done chan bool) คือฟังก์ชันที่ถูกเรียกใน
goroutine ใหม่
จะพิมพ์ "working..." จากนั้นหน่วงเวลาการทำงานเป็น
เวลา 1 วินาทีด้วย time.Sleep(time.Second)
เมื่อทำงานเสร็จ จะพิมพ์ "done" และส่งค่า true
ลงใน channel done เพื่อแจ้งให้ goroutine อื่นรู้ว่างาน
เสร็จสิ้นแล้ว

สร้าง Channel:
done := make(chan bool, 1) สร้าง buffered channel
ชื่อ done ที่รับค่า bool และมี buffer ขนาด 1 ซึ่ง
หมายความว่ามันสามารถเก็บค่าหนึ่งค่าไว้ได้ก่อนที่จะต้องมี
การรับค่าออก

เรียกใช้ worker ใน Goroutine ใหม่:
go worker(done) เรียกใช้ฟังก์ชัน worker ใน
goroutine ใหม่ พร้อมกับส่ง channel done เข้าไปเพื่อให้
worker ส่งค่ากลับมาเมื่อทำงานเสร็จ

บล็อกการทำงานจนกว่า Worker จะเสร็จสิ้น:
*<-done บล็อกการทำงานของ main goroutine จนกว่า
จะมีการรับค่าจาก channel done ซึ่งจะเกิดขึ้นหลังจากที่
worker goroutine ส่งค่า true ลงใน channel เพื่อบอก
ว่าทำงานเสร็จแล้ว
*/
/*การทำงานของโค้ด
เมื่อโค้ดนี้รัน ฟังก์ชัน main จะเริ่มต้นด้วยการสร้าง channel done
จากนั้นจะเรียกฟังก์ชัน worker ใน goroutine ใหม่
ฟังก์ชัน worker จะทำงานและใช้เวลา 1 วินาทีเพื่อจำลองการ
ทำงานเสร็จ
หลังจากนั้น worker จะส่งค่า true ลงใน channel done
เพื่อบอกว่าเสร็จงานแล้ว
main goroutine จะรอจนกว่าจะได้รับค่าจาก channel done
ซึ่งจะเกิดขึ้นหลังจากที่ worker goroutine ทำงานเสร็จ ทำให้
โปรแกรมไม่จบการทำงานก่อนที่ worker จะทำงานเสร็จ
*/
