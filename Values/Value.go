package main

import (
	f "fmt"
)

/*
	Go มีประเภทค่าต่างๆ รวมถึงสตริง จำนวนเต็ม จำนวนทศนิยม บูลีน

และอื่นๆ นี่เป็นตัวอย่างพื้นฐานบางส่วน
สตริง ซึ่งสามารถนำมาต่อกันได้ด้วย +
จำนวนเต็มและจำนวนทศนิยม
บูลีน พร้อมด้วยตัวดำเนินการบูลีนตามที่เราคาดหวัง
*/
func main() {
	f.Println("go" + "lang")
	f.Println("1+1 =", 1+1)
	f.Println("7.0/3.0 =", 7.0/3.0)

	f.Println(true && false)
	f.Println(true || false)
	f.Println(!true)
}

/* อธิบายโค้ด
แสดงผลการต่อสตริง go และ lang ด้วยเครื่อหมาย +
แสดงผลการบวกเลขจำนวนเต็ม 1+1
แสดงผลการหารเลขทศนิยม 7.0 ด้วย 3.0
แสดงผลการดำเนินการทางตรรกะ (logical operations) สำหรับค่าบูลีน
AND (&&) ระหว่าง true และ false
OR (||) ระหว่าง true และ false
NOT (!) ของ true
โค้ดนี้แสดงผลลัพธ์ของแต่ละการดำเนินการผ่านฟังก์ชัน fmt.Println()
*/
