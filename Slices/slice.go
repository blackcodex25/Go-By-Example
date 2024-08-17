package main

import (
	f "fmt"
)

func main() {
	// การประกาศ Slice และการตรวจสอบค่าเริ่มต้น
	var s []string // ประกาศ slice ของ string ชื่อ s
	// พิมพ์ว่า s ยังไม่ถูกกำหนดค่า(nil) และความยาวของมันเป็น 0
	f.Println("uninit:", s, s == nil, len(s) == 0)
}
