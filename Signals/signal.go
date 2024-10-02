package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// สร้าง channel สำหรับรับสัญญาณ (Signals)
	sigs := make(chan os.Signal, 1)

	// ลงทะเบียนให้ channel นี้รับสัญญาณ SIGINT และ SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// สร้าง channel สำหรับแจ้งว่าโปรแกรมสามารถออกได้
	done := make(chan bool, 1)

	// ใช้ goroutine เพื่อรอรับสัญญาณ
	go func() {
		// รอรับสัญญาณจาก channel sigs
		sig := <-sigs
		// พิมพ์ข้อความเมื่อได้รับสัญญาณ
		fmt.Println()
		fmt.Println(sig)
		// แจ้งว่าโปรแกรมสามารจบได้
		done <- true
	}()

	// แสดงข้อความว่าโปรแกรมกำลังรอรับสัญญาณ
	fmt.Println("awaiting signal")
	// รอรับค่าจาก channel done
	<-done
	// แสดงข้อความก่อนออกจากโปรแกรม
	fmt.Println("exiting")
}

/* ลงทะเบียน Channel กับ Signal */
/* signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) */
/* ใช้ signal.Notify เพื่อลงทะเบียนให้ channel sigs รับสัญญาณ */
/* SIGINT และ SIGTERM */

/* สร้าง Channel สำหรับรับสัญญาณ */
/* sigs := make(chan os.Signal, 1) */
/* สร้าง Channel sigs เพื่อใช้รับสัญญาณจากระบบ */

/* Goroutine รอรับสัญญาณ */
/* go func() {
	sig := <-sigs
	fmt.Println(sig)
	done <- true
}() */
/* สร้าง goroutine ที่รอรับสัญญาณจาก channel sigs */
/* เมื่อได้รับสัญญาณ จะพิมพ์สัญญาณที่ได้รับออกมาและส่งค่าจริง */
/* ไปยัง channel done */

/* รอรับสัญญาณในฟังก์ชันหลัก */
/* fmt.Println("awaiting signal") */
/* <-done */
/* แสดงข้อความว่ากำลังรอรับสัญญาณ */
/* รอรับค่าจาก channel done ซึ่งจะถูกส่งเมื่อได้รับสัญญาณ */

/* การทดสอบโปรแกรม */
/* เมื่อเรารันโปรแกรมนี้ใน terminal และกด Ctrl+C (ส่ง SIGINT) หรือส่ง */
/* หรือส่งคำสั่ง SIGTERM โปรแกรมจะพิมพ์สัญญาณที่ได้รับและแสดงข้อความ */
/* "exiting" ก่อนที่จะสิ้นสุดการทำงาน */

/* การจัดการ Signal ใน Go ช่วยให้โปรแกรมของเรามีความสามารถในการ */
/* ปิดการทำงานอย่างมีระเบียบ ลดโอกาสในการสูญเสียข้อมูลหรือทำให้ระบบ */
/* อยู่ในสถานะไม่เสถียร */
