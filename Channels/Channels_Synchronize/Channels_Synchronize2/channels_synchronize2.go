package main

import (
	f "fmt"
	"sync"
)

/* โค้ดนี้เป็นตัวอย่างของการใช้ sync.WaitGroup เพื่อซิงโครไนซ์การ
ทำงานของ goroutines พร้อมกับการใช้ channel เพื่อสื่อสารระหว่าง
goroutines
*/
// ฟังก์ชัน Worker ทำงานใน goroutine และ
// ส่งสัญญาณว่าเสร็จสิ้นการทำงานผ่าน channel
func Worker(done chan bool) {
	f.Println("Working...") // พิมพ์ข้อความว่าเริ่มทำงาน
	f.Println("Done")       // พิมพ์ข้อความว่าเสร็จสิ้นการทำงาน
	done <- true            // ส่งค่า true ไปยัง channel เพื่อแจ้งว่าเสร็จสิ้นการทำงาน
}

func main() {
	var wg sync.WaitGroup      // สร้าง WaitGroup เพื่อรอการทำงานของ goroutine
	done := make(chan bool, 1) // สร้าง channel ที่มี buffer ขนาด 1 ค่า
	wg.Add(1)                  // เพิ่ม goroutine หนึ่งตัวใน WaitGroup
	go func() {
		defer wg.Done() // ลดค่าของ WaitGroup เมื่อ goroutine นี้ทำงานเสร็จ
		Worker(done)    // เรียกใช้ฟังก์ชัน Worker
	}()
	wg.Wait() // รอจนกว่า goroutine ทั้งหมดใน WaitGroup จะทำงานเสร็จ
}

/*คำอธิบายโค้ด Logic ของโค้ด
ฟังก์ชัน Worker:
ฟังก์ชัน Worker(done chan bool) จะพิมพ์ข้อความ
Working..." และ "Done" ออกมา หลังจากนั้นมันจะ
ส่งค่า true ลงใน channel done เพื่อแจ้งว่าการทำงาน
เสร็จสิ้นแล้ว

การสร้าง WaitGroup:
var wg sync.WaitGroup สร้างตัวแปร wg ซึ่งเป็น
WaitGroup ที่ใช้ในการรอให้ goroutine ทั้งหมดในกลุ่มนั้น
ทำงานเสร็จ

การสร้าง Channel:
done := make(chan bool, 1) สร้าง channel ชื่อ
done ที่มี buffer ขนาด 1 ค่า ซึ่งหมายความว่าสามารถเก็บ
ค่าที่ส่งไปได้หนึ่งค่าก่อนที่จะต้องมีการรับค่าออก

การเพิ่ม Goroutine ใน WaitGroup:
wg.Add(1) เพิ่มหนึ่งใน WaitGroup wg เพื่อแจ้งให้รู้ว่า
จะมีการเพิ่ม goroutine ที่ต้องรอให้ทำงานเสร็จก่อนที่
โปรแกรมจะจบ

การเรียกใช้ Worker ใน Goroutine ใหม่:
go func() { defer wg.Done(); Worker(done) }()
เรียกฟังก์ชัน Worker ใน goroutine ใหม่ เมื่อ goroutine
นี้ทำงานเสร็จ defer wg.Done() จะถูกเรียกเพื่อแจ้งให้
WaitGroup รู้ว่า goroutine นี้เสร็จสิ้นการทำงานแล้ว

การรอให้ Goroutine เสร็จสิ้น:
wg.Wait() ทำให้ main goroutine รอจนกว่า
goroutine ที่อยู่ใน WaitGroup จะทำงานเสร็จทั้งหมด
*/
/*การทำงานของโค้ด
1.โค้ดจะเริ่มด้วยการสร้าง WaitGroup และ done channel
2.จากนั้นจะเพิ่ม goroutine หนึ่งตัวใน WaitGroup
3.Goroutine ใหม่จะถูกสร้างและเรียกใช้ฟังก์ชัน Worker ซึ่งจะ
พิมพ์ข้อความออกมาพร้อมส่งค่า true ไปยัง channel done
4.Goroutine จะเรียก wg.Done() เพื่อแจ้งว่าเสร็จงานแล้ว
5.main goroutine จะรอจนกว่า WaitGroup จะตรวจสอบว่า
goroutine ทั้งหมดเสร็จงานก่อนที่จะดำเนินการต่อไป

ในโค้ดนี้ แม้ว่าจะมีการส่งค่าผ่าน channel done แต่ค่าที่ถูกส่งนั้นไม่
ได้ถูกใช้ในที่อื่น คุณอาจจะใช้ค่าจาก channel นี้เพื่อทำงานเพิ่มเติม
หรือการตรวจสอบผลการทำงานตามที่ต้องการ
*/
