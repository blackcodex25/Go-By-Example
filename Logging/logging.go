package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

/* การใช้เครื่องมือสำหรับการบันทึกข้อมูลหรือเหตุการณ์ต่างๆ */
/* ในภาษา Go ผ่านแพ็กเกจ log และ log/slog */
/* ซึ่งมีฟังก์ชันสำหรับการสร้างบันทึก (log) ที่เรียบง่ายและโครงสร้าง*/
/* การบันทึกที่ซับซ้อน เช่น การจัดเก็บข้อมูลในรูปแบบ JSON นอกจากนี้ */
/* ยังสามารถตั้งค่า flag เพื่อปรับรูปแบบการแสดงผล เช่น วันที่ เวลา */
/* หรือชื่อไฟล์และบรรทัดที่บันทึกถูกเรียกใช้ ผู้ใช้ยังสามารถสร้าง logger */
/* แบบกำหนดเองหรือเปลี่ยนแปลงจุดหมายของการบันทึก เช่นเขียนบันทึกลงไฟล์ */
func main() {
	// บันทึกข้อความด้วย logger มาตรฐาน
	log.Println("standard logger")

	//  ตั้งค่า flag เพื่อบันทึกเวลาในระดับ microseconds
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro") // บันทึกข้อมูลพร้อมเวลา

	// ตั้งค่า flag เพื่อบันทึกไฟล์และบรรทัด
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line") // บันทึกข้อมูลพ้อมชื่อไฟล์และบรรทัด

	// สร้าง logger ใหม่ที่ส่งออกไปยัง os.Stdout พร้อม prefix
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog") // บันทึกข้อความด้วย mylog

	// เปลี่ยน prefix ของ mylog
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog") // บันทึกข้อความด้วย prefix ใหม่

	// สร้าง buffer เพื่อเก็บ log ข้อความ
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags) // logger ที่บันทึกลง buffer

	buflog.Println("hello") // บันทึกข้อความไปยัง buffer

	// แสดงผลข้อมูลจาก buffer
	fmt.Print("from buflog:", buf.String())

	// สร้าง jsonHandler สำหรับบันทึก log แบบ JSON
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler) // สร้าง logger ที่ใช้ jsonHandler
	myslog.Info("hi there")         // บันทึกข้อมูลในรูปแบบ JSON

	// บันทึกข้อมูลเพิ่มเติมในรูปแบบ key=value
	myslog.Info("hello again", "key", "val", "age", 25)
}

/* อธิบายโค้ด */
/* Logger: ใช้เพื่อบันทึกข้อมูลที่เกิดขึ้นในโปรแกรม มีให้เลือกทั้งแบบ */
/* มาตรฐาน (standard logger) และแบบมีโครงสร้าง (structured logger) */
/* Flags: ปรับแต่งรูปแบบของข้อมูลที่บันทึก เช่น เวลา ชื่อไฟล์ และหมายเลขบรรทัด */
/* Buffer: ใช้เพื่อเก็บข้อมูลที่บันทึกโดยไม่ส่งออกทันที สามารถพิมพ์ได้ภายหลัง*/
/* JSON Logger: ใช้เพื่อบันทึกข้อมูลในรูปแบบ JSON ซึ่งสะดวก */
/* สำหรับการวิเคาะห์และการจัดการข้อมูล */
