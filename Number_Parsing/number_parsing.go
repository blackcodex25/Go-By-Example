package main

import (
	"fmt"
	"strconv"
)

/*
โค้ดนี้แสดงตัวอย่างการใช้ฟังก์ชันในแพ็คเกจ strconv ของ Go
สำหรับการแปลงค่าจากสตริงเป็นตัวเลขประเภทต่างๆ
*/
func main() {
	// แปลงสตริง "1.234" เป็น flat64
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// แปลงสตริง "1234" เป็น int64 โดยใช้ฐาน 0 (มายถึงฐานอัตโนมัติที่เป็นฐาน 10)
	i, _ := strconv.ParseInt("1234", 0, 64)
	fmt.Println(i)

	// แปลงสตริง "0x1c8" เป็น int64 โดยใช้ฐาน 0 (หมายถึงฐานอัตโนมัติที่เป็นฐาน 16)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// แปลงสตริง "789" เป็น uint64 โดยใช้ฐาน 0 (หมายถึงฐานอัตโนมัติที่เป็นฐาน 10)
	u, _ := strconv.ParseInt("789", 0, 64)
	fmt.Println(u)

	// แปลงสตริง "135" เป็น int โดยใช้ Atoi ซึ่งเป็นการแปลงสตริงเป็น int
	// โดยอัติโนมัติ
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// พยายามแปลงสตริง "wat" เป็น int ซึ่งจะล้มเหลวและส่งกลับข้อผิดพลาด
	_, err := strconv.Atoi("wat")
	fmt.Println(err)
}

/* อธิบายโค้ด
1.strconv.ParseFloat:
ใช้สำหรับแปลงสตริงที่เป็นจำนวนจริง (float) เป็น float64
ตัวอย่าง: "1.234" แปลงเป็น 1.234

2.strconv.ParseInt:
ใช้สำหรับแปลงสตริงเป็นจำนวนเต็ม (integer) โดยเราต้องระบุฐาน (base)
และขนาดของบิต (bit size)
ตัวอย่าง:
"123" แปลงเป็น 123 (ฐาน 10 เนื่องจากใช้ฐาน 0)
"0x1c8" แปลงเป็น 456 (ฐาน 16 เนื่องจากมี "0x" นำหน้า)

3.strconv.ParseUint:
ใช้สำหรับแปลงสตริงเป็นจำนวนเต็มที่ไม่เป็นลบ (unsigned integer) โดย
เราต้องระบุฐาน (base) และขนาดของบิต (bit size)
ตัวอย่าง:
"789" แปลงเป็น 789 (ฐาน 10 เนื่องจากใช้ฐาน 0)

4.strconv.Atoi:
ใช้สำหรับแปลงสตริงเป็น int, ซึ่งเป็นการใช้ฐาน 10 โดยอัตโนมัติ
ตัวอย่าง: "135" แปลงเป็น 135

5.ข้อผิดพลาด:
strconv.Atoi("wat") จะเกิดข้อผิดพลาดเพราะ "wat" ไม่สามารถแปลงเป็นตัวเลขได้
e จะเก็บข้อผิดพลาดนี้และพิมพ์ออกมา
*/
/* การจัดการข้อผิดพลาด
โค้ดนี้ใช้ _ เพื่อละเว้นข้อผิดพลาดที่เกิดขึ้นจากการแปลงค่า
เราควรตรวจสอบข้อผิดพลาดอย่างจริงจังในโปรแกรมจริงเพื่อให้แน่ใจว่าโปรแกรมทำงานได้
อย่างถูกต้อง
โดยทั่วไปควรจัดการกับข้อผิดพลาดเพื่อป้องกันไม่ให้เกิดปัญหาที่อาจเกิดขึ้นจาก
การแปลงข้อมูลผิดพลาด
*/
