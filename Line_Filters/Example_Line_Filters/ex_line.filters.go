package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
โค้ดนี้เป็นตัวอย่างของ Line Filter ที่แปลงข้อความในแต่ละ
บรรทัดจาก stdin ให้เป็นตัวพิมพ์ใหญ่และแสดงผลทาง stdout
*/
func main() {
	// สแกนข้อมูลจาก stdin ทีละบรรทัด
	scanner := bufio.NewScanner(os.Stdin)

	// ลูปเพื่ออ่านและประมวลผลข้อมูลบรรทัดต่อบรรทัด
	for scanner.Scan() {
		// แปลงข้อมูลในแต่ละบรรทัดให้เป็นตัวพิมพ์ใหญ่
		ucl := strings.ToUpper(scanner.Text())
		// ส่งผลลัพธ์ออกทาง stdout
		fmt.Println(ucl)
	}
	// ตรวจสอบข้อผิดพลาด
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error:", err)
		os.Exit(1)
	}
}
