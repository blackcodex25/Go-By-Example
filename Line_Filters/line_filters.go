package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
การสร้างโปรแกรมที่สามารถอ่านข้อมูลจากอินพุต (stdin), ประมวลผลข้อมูลนั้น
แล้วส่งผลลัพธ์ออกไปยังเอาต์พุต (stdout) เช่นเดียวกับโปรแกรมอย่าง grep หรือ sed
ตัวอย่างใน Go นี้จะแปลงข้อความที่รับเข้ามาทั้งหมดให้เป็นตัวพิมพ์ใหญ่ และใช้หลักการ
สแกนข้อมูลบรรทัดต่อบรรทัด
*/
func main() {
	// สร้างตัวแสกนจาก os.Stdin ซึ่งเป็นอินพุตจากผู้ใช้หรือไฟล์ที่ถูกส่งผ่าน stdin
	scanner := bufio.NewScanner(os.Stdin)

	// ลูปวนผ่านข้อมูลที่รับเข้ามาบรรทัดต่อบรรทัด
	for scanner.Scan() {
		// แปลงข้อความที่สแกนมาเป็นตัวพิมพ์ใหญ่
		ucl := strings.ToUpper(scanner.Text())
		// พิมพ์ข้อความที่ถูกแปลงแล้วออกทาง stdout
		fmt.Println(ucl)
	}

	// ตรวจสอบว่ามีข้อผิดพลาดเกิดขึ้นในระหว่างการสแกนหรือไม่
	if err := scanner.Err(); err != nil {
		// ถ้ามีข้อผิดพลาดให้พิมพ์ข้อความแจ้งเตือนและออกจากโปรแกรม
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}

/* การทำงานของโค้ด
1.bufio.NewScanner(os.Stdin): สร้างตัวสแกนสำหรับอ่านอินพุตจาก stdin
โดยใช้แพ็คเกจ bufio เพื่อให้สามารถสแกนข้อมูลทีละบรรทัดได้อย่างมีประสิทธิภาพ

2.for scanner.Scan(): ทำงานในรูปแบบลูปเพื่ออ่านข้อมูล
ทีละบรรทัดจาก stdin จนกว่าจะถึงจุดสิ้นสุดหรือมีข้อผิดพลาด

3.strings.ToUpper(scanner.Text()): แปลงข้อความในแต่ละบรรทัด
ที่อ่านมาให้เป็นตัวพิมพ์ใหญ่ทั้งหมดโดยใช้ฟังก์ชัน ToUpper

4.fmt.Println(ucl): พิมพ์ข้อความที่ถูกแปลงแล้วออกไปที่ stdout

5.scanner.Err(): เมื่อสแกนสิ้นสุด จะตรวจสอบว่ามีข้อผิดพลาดหรือไม่
หากมีข้อผิดพลาด เช่น อินพุตไม่สามารถอ่านได้ จะพิมพ์ข้อผิดพลาดและ
ออกจากโปรแกรม
*/
/* สรุป
โค้ดนี้เป็นตัวอย่างของโปรแกรม Line Filter ที่อ่านข้อมูลจากอินพุต (stdin)
ทีละบรรทัด แปลงเป็นตัวพิมพ์ใหญ่ แล้วพิมพ์ออกที่เอาต์พุต (stdout)
*/
