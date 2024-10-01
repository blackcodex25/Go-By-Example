package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

/* โค้ดนี้แสดงการใช้งานแพ็กเกจ os/exec เพื่อเรียกใช้คำสั่งภายนอกใน */
/* Go โดยใช้คำสั่ง exec.Command() ซึ่งเป็นการสร้างและเรียกใช้โปรเซส */
/* ที่ไม่ใช่ Go เช่น date, grep, และ ls พร้อมกับจัดการ input/output */
/* ระหว่างโปรเซสเหล่านั้น */

func main() {
	// สร้างคำสั่ง date เพื่อเรียกโปรเซสสำหรับแสดงวันที่
	dateCmd := exec.Command("date")

	// รันคำสั่งและเก็บผลลัพธ์
	dateOut, err := dateCmd.Output()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut)) // แสดงผลลัพธ์ของคำสั่ง "date"

	// เรียกใช้คำสั่ง "date" พร้อม flag ที่ไม่มีอยู่จริง เพื่อแสดงข้อผิดพลาด
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		// ตรวจสอบชนิดของข้อผิดพลาด
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err) // ข้อผิดพลาดเนื่องจากการเรียกโปรเซสล้มเหลว
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode()) // ข้อผิดพลาดเนื่องจากคำสั่งสิ้นสุดด้วยสถานะผิดพลาด
		default:
			fmt.Println("unknown error:", err)
		}
	}
	// สร้างคำสั่ง "grep" และใช้สำหรับตรวจหาคำใน input
	grepCmd := exec.Command("grep", "hello")

	// จัดการ input/output pipe
	grepIn, _ := grepCmd.StdinPipe()   // สร้าง pipe สำหรับส่งข้อมูล input
	grepOut, _ := grepCmd.StdoutPipe() // สร้าง pipe สำหรับรับข้อมูล output
	grepCmd.Start()                    // เริ่มคำสัง grep

	// ส่งข้อมูลเข้าไปในคำสั่ง "grep"
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close() // ปิด pipe หลังจากส่งข้อมูลเสร็จ

	// อ่านผลลัพธ์จากคำสั่ง "grep"
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait() // รอจนคำสังสิ้นสุด

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes)) // แสดงผลลัพธ์ที่ได้จากคำสั่ง "grep"

	// ใช้คำสั่ง "ls" ผ่าน bash เพื่อแสดงรายละเอียดไฟล์
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output() // รันคำสั่งและเก็บผลลัพธ์
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut)) // แสดงผลลัพธ์ที่ได้จากคำสั่ง ls
}

/* หลักการทำงานของโค้ด */
/* 1.การเรียกใช้โปรเซสภายนอก */
/* โค้ดใช้ exec.Command() เพื่อเรียกใช้คำสั่ง shell เช่น date, grep, และ ls */
/* ผลลัพธ์จากการเรียกใช้คำสั่งจะถูกเก็บในตัวแปร output */
/* 2.การจัดการข้อผิดพลาด */
/* เมื่อคำสั่งล้มเหลว จะมีการแสดงข้อผิดพลาดผ่าน exec.Error */
/* หรือ exec.ExitError ขึ้นอยู่กับชนิดของข้อผิดพลาด */
/* 3.การใช้ pipe */
/* StdinPipe() และ StdoutPipe() ใช้สำหรับส่ง input และ รับ output */
/* ระหว่างโปรเซส เช่นการใช้ grep กับข้อความ input ที่ส่งไป */
