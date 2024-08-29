package main

import (
	"fmt"
)

/*
โค้ดนี้แสดงวิธีการสื่อสารและประสานงานระหว่าง goroutines
โดยใช้ channels ในภาษา Go โดยโค้ดจะส่งงาน (jobs) ไปยัง
channel และรอจนกว่างานทั้งหมดจะถูกประมวลผล

การทำงานของโค้ด
การสร้าง Channels:
โค้ดเริ่มต้นโดยการสร้างสอง channels
jobs เป็น buffered channel ที่สามารถเก็บได้ถึง 5 งาน (chan int, 5)
การสร้าง Goroutine:
สร้าง goroutine ขึ้นมาหนึ่งตัว ซึ่งจะทำหน้าที่
รับงานจาก channel jobs ทีละงาน
ใช้ j, more := <-jobs เพื่อรับค่าจาก
channel และตรวจสอบว่า channel ยังมีงานที่จะส่งเข้ามาหรือไม่
j คือค่างานที่ถูกส่งเข้ามา
more เป็นตัวบ่งชี้ว่า channel ยังเปิดอยู่หรือไม่ (true หมายความว่ายังมีงานส่งเข้ามาได้)
ถ้า more เป็น true จะพิมพ์ข้อความว่าได้รับงาน ("received job", j)
"received all jobs" และส่งสัญญาณผ่าน channel done เพื่อบอกว่า goroutine
ทำงานเสร็จแล้ว จากนั้น return ออกจาก goroutine

การส่งงานไปยัง Channel:
ใน main goroutine (ฟังก์ชัน main) มีการ
ส่งงาน (เลข 1, 2, 3) ไปยัง channel jobs ที
ละงาน โดยใช้ jobs <- j และพิมพ์ข้อความ
"sent job", j หลังจากส่งงานแต่ละชิ้น
หลังจากส่งงานทั้งหมดแล้ว จะปิด channel
jobs ด้วย close(jobs) และพิมพ์ข้อความ

การรอให้ Goroutine ทำงานเสร็จ:
main goroutine รอรับสัญญาณจาก channel
วโดยลองรับค่าจาก channel อีกครั้ง
เนื่องจาก channel ถูกปิดแล้ว จะไม่สามารถรับ
ค่าได้อีก ดังนั้น ok จะเป็น false และพิมพ์
ข้อความ "received more jobs:", ok
*/
func main() {
	// สร้าง buffered channel สำหรับส่งงานชนิด int สามารถเก็บได้สูงสุด 5 งาน
	jobs := make(chan int, 5)
	// สร้าง unbuffered channel สำหรับส่งสัญญาณเมื่อประมวลผลงานเสร็จสิ้น
	done := make(chan bool)

	// สร้าง goroutine สำหรับรับงานจาก channels 'jobs'
	go func() {
		for {
			// รับงานจาก channels 'jobs' พร้อมตรวจสอบว่ามีงานเข้ามาอีกหรือไม่
			j, more := <-jobs
			if more {
				// ถ้ายังมีงานเข้ามา พิมพ์ข้อความว่าได้รับงานแล้ว
				fmt.Println("received job", j)
			} else {
				// ถ้ายังไม่มีงานเข้ามาอีก (channels ถูกปิด) พิมพ์ข้อความและส่งสัญญาณว่าเสร็จสิ้น
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// ส่งงานไปยัง channels 'jobs' ทั้งหมด 3 งาน (เลข 1 ถึง 3)
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// ปิด channels 'jobs' หลังจากส่งงานเสร็จ
	close(jobs)
	fmt.Println("sent all jobs")

	// รอรับสัญญาณจาก goroutine ว่าประมวลผลงานเสร็จสิ้นแล้ว
	<-done

	// ลองรับค่าจาก channel 'jobs' อีกครั้งเพื่อตรวจสอบว่า channels ปิดแล้ว
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}

/*โปรแกรมที่จำลองการส่งงาน (jobs) ไปยัง worker goroutine
เพื่อทำงานประมวลผลและส่งสัญญาณกลับเมื่อทำงานเสร็จสิ้นทั้งหมด
โค้ดนี้แสดงการทำงานร่วมกันระหว่าง goroutine
มีการใช้ channel เพื่อส่งงานและตรวจสอบว่าทุกงาน
ได้รับการประมวลผลแล้วหรือยัง
หลังจากส่งงานเสร็จแล้ว จะปิด channel เพื่อบอกว่า
ไม่มีงานเพิ่มเติมที่จะส่งเข้ามา
Goroutine จะทำงานต่อไปจนกว่าจะได้รับงาน
ทั้งหมด และส่งสัญญาณให้ goroutine หลักทราบ
เมื่อทำงานเสร็จแล้ว
*/
