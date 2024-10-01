package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/* บทความนี้อธิบายการใช้ฟังก์ชัน syscall.Exec() ในภาษา Go */
/* เพื่อแทนที่โปรเซสปัจจุบันด้วยโปรเซสใหม่ โดยตัวอย่างในบทความใช้ */
/* คำสั่ง ls เพื่อแสดงไฟล์ในระบบ */
/* 1.ใช้ exec.LookPath("ls") เพื่อค้นหา path ของ bianry ls */
/* 2.ส่งอาร์กิวเมนต์ในรูปแบบ slice เพื่อใช้กับคำสั่ง ls เช่น -a, -l, และ -h */
/* 3.ใช้ syscall.Exec() เพื่อแทนที่โปรเซสปัจจุบันด้วยคำสั่ง ls */
/* พร้อมกับ environment variables เดิม */

func main() {
	// ค้นหา path ของโปรแกรม ls ด้วย exec.LookPath
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		log.Printf("Error: %v", lookErr) // ถ้าหาไม่เจอให้เกิด panic
		return
	}

	// กำหนดอาร์กิวเมนต์ของคำสั่ง ls
	args := []string{"ls", "-a", "-l", "-h"}

	// นำ environment variables ปัจจุบันมาใช้
	env := os.Environ()

	// เรียก syscall.Exec เพื่อแทนที่โปรเซสด้วย "ls" พร้อมอาร์กิวเมนต์และ environment variables
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Printf("Error: %v", execErr) // ถ้าเกิดข้อผิดพลาดให้แสดง error
		return
	}
}

/* หลักการทำงาน */
/* 1.exec.LookPath("ls"): ค้นหา binary ของคำสั่ง ls */
/* (อาจเป็น/bin/ls) ถ้าหาไม่เจอจะเกิด panic */
/* 2.args := []string{"ls", "-a", "-l", "-h"}: กำหนด */
/* อาร์กิวเมนต์ให้คำสั่ง ls ได้แก่แสดงไฟล์ที่ซ่อน, รายละเอียดแบบ long list, */
/* และแสดงขนาดไฟล์ในรูปแบบที่เข้าใจง่าย */
/* 3.os.Environ(): ใช้ environment variables ปัจจุบัน */
/* 4.syscall.Exec(binary, args, env): แทนที่โปรเซสปัจจุบันด้วยคำสั่ง ls */
/* พร้อมอาร์กิวเมนต์และ environment */
