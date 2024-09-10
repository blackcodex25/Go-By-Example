package main

import (
	"fmt"
	"regexp"
)

// ตัวอย่างการใช้งาน Regular Expression ใน Go
// เราสามารถใช้ฟังก์ชัน MatchString เพื่อตรวจสอบว่า string นั้น
// ตรงกับรูปแบบที่กำหนดหรือไม่
// สำหรับรูปแบบนี้ เราจะใช้ a(b*)c ซึ่งหมายความว่าเราต้องการ
// string ที่เริ่มต้นด้วย a และลงท้ายด้วย c
// และอาจมีตัวอักษร b อยู่ระหว่างกลางจำนวนใดๆ ก็ได้
func main() {

	// รูปแบบที่กำหนดในนี้คือ a(b*)c ซึ่งเราต้องการรูปแบบที่เริ่มต้นด้วย a และลงท้ายด้วย c
	// และอาจมีตัวอักษร b อยู่ระหว่างกลางจำนวนใดๆ ก็ได้
	re := regexp.MustCompile(`a(b*)c`)

	fmt.Println(match("ac", re))    // true
	fmt.Println(match("abc", re))   // true
	fmt.Println(match("abbbc", re)) // true
	fmt.Println(match("abccc", re)) // false
}

// ตรวจสอบว่าสตริง s ตรงกับรูปแบบที่กำหนดใน re หรือไม่
// ฟังก์ชันนี้จะใช้ฟังก์ชัน MatchString ของ regexp.Regexp
// เพื่อตรวจสอบรูปแบบของสตริง
// s สตริงที่ต้องการตรวจสอบ
// re รูปแบบที่กำหนดโดยใช้ *regexp.Regexp
// คืนค่า true หากสตริง s ตรงกับรูปแบบที่กำหนดใน re
// และ false หากไม่ตรงกับรูปแบบ
func match(s string, re *regexp.Regexp) bool {
	return re.MatchString(s)
}
