package main

import (
	"fmt"
	"slices"
)

/*
	แพ็กเกจ slices ของ Go มีการทำงานในการจัดเรียงลำดับ

ข้อมูลสำหรับประเภท built-in และประเภทที่ผู้ใช้กำหนดเอง
เราจะเริ่มด้วยการจัดเรียงสำหรับประเภท built-in ก่อน
ฟังก์ชันการจัดเรียงเหล่านี้เป็นแบบ generic และสามารถทำงาน
ได้กับประเภท built-in ที่มีลำดับ สำหรับรายการของประเภทที่มี
ลำดับ ดูที่ cmp.Ordered
ตัวอย่างของการจัดเรียงตัวเลขประเภท int
เรายังสามารถใช้แพ็กเกจ slices เพื่อตรวจสอบว่า slice ได้ถูก
จัดเรียงเรียบร้อยแล้วหรือไม่
*/
func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:  ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

/*โค้ดนี้ใช้แพ็กเกจ slices ในการจัดเรียงและตรวจสอบการเรียง
ลำดับของ slice ทั้งที่เป็น string และ int ในภาษา Go
โดยมีรายละเอียดดังนี้:
1.str := []string{"c", "a", "b"}
สร้าง slice ที่มีค่าเป็น {"c", "a", "b"}

2.เรียกลำดับ slice ของ string
slices.Sort(strs) จะแสดงผลเป็น Strings: [a b c]

3.พิมพ์ผลลัพธ์
fmt.Println("Strings:", strs)
จะแสดงผลเป็น Strings: [a b c]
ประกาศ slice ของ intประกาศ slicebu ของ int

4.ประกาศ slice ของ int
ints := []int{7, 2, 4}
สร้าง slice ที่มีค่าเป็น {"7", "2", "4"}

5.เรียงลำดับ slice ของ int
slices.Sort(ints)
ใช้ฟังก์ชัน Sort เพื่อจัดเรียงตัวเลขใน slice ints จะได้
{"2", "4", "7"}

6.พิมพ์ผลลัพธ์
fmt.Println("Ints: ", ints)
จะแสดงผลเป็น Ints: [2 4 7]

7.ตรวจสอบว่า slice ถูกจัดเรียงแล้วหรือไม่
s := slices.IsSorted(ints)
ใช้ฟังก์ชัน IsSorted เพื่อตรวจสอบว่า slice ints ถูก
จัดเรียงแล้วหรือไม่ ในกรณีนี้จะได้ค่า true เพราะ ints ถูก
จัดเรียงแล้ว

8.พิมพ์ผลลัพธ์การตรวจสอบ
fmt.Println("Sorted: ", s)
จะแสดงผลเป็น Sorted: true
*/
/* สรุป
โค้ดนี้ใช้ฟังก์ชันจากแพ็กเกจ slices เพื่อจัดเรียง string
และ int ใน slice และตรวจสอบว่า slice ของ int ถูกจัดเรียง
เรียบร้อยหรือไม่.
*/
