package main

import (
	f "fmt"
)

/* Variadic functions ในภาษา Go คือฟังก์ชันที่สามารถรับอากิวเมนต์
(arguments) จำนวนเท่าใดก็ได้ในตำแหน่งท้ายสุด ตัวอย่างที่ใช้บ่อยของ
variadic function ก็คือ fmt.Println() ซึ่งสามารถรับอาร์กิวเมนต์
หลายๆ ตัวและพิมพ์ออกมาได้ในบรรทัดเดียว
*/
/* คุณสมบัติและการใช้งาน Variadic Functions
1.การสร้าง Variadic Function
ฟังก์ชันที่มีการกำหนด ...type ในพารามิเตอร์จะกลายเป็น
variadic function ซึ่งสามารถรับอาร์กิวเมนต์จำนวนมากเท่า
ที่ต้องการได้
ในตัวอย่างนี้ func sum(nums ...int) คือฟังก์ชันที่สามารถรับ
จำนวนตัวเลข (int) ใดๆ ก็ได้

2.การใช้งานพารามิเตอร์ Variadic ภายในฟังก์ชัน
ภายในฟังก์ชัน nums จะถูกมองว่าเป็น slice ของ int
([]int)
เราสามารถใช้ len(nums) เพื่อหาความยาวของ slice ใช้
range เพื่อวนลูปผ่าน slice และทำการดำเนินการอื่นๆ ที่ใช้กับ
slice ได้

3.การเรียกใช้ Variadic Function
เราสามารถเรียกใช้ variadic function โดยส่งอาร์กิวเมนต์
เป็นจำนวนเท่าใดก็ได้ ตัวอย่างเช่น sum(1, 2) และ sum(1,
2, 3) ซึ่งจะส่งค่าตัวเลขที่แตกต่างกันให้ฟังก์ชัน
หากเรามี slice ที่มีค่าหลายค่าอยู่แล้ว เราสามารถส่ง slice
นั้นเข้าไปใน vaiadic function ได้โดยใช้
func(slice...) เช่น sum(nums...) ซึ่งจะกระจายค่าภายใน
slice นั้นเป็นอากิวเมนต์แต่ละตัว
*/
// ฟังก์ชัน sum คือฟังก์ชันที่รับอาร์กิวเมนต์หลายตัวที่เป็นชนิด int
// nums ภายในฟังก์ชันจะถูกมองว่าเป็น slice ของ int ([]int)
func sum(nums ...int) {
	// พิมพ์ค่าของ nums ที่เป็น slice และตามด้วยช่องว่าง
	f.Println(nums, " ")
	total := 0
	// การวนลูปผ่านพารามิเตอร์ Variadic
	// ลูป for จะวนลูปผ่านแต่ละค่าใน nums
	for _, num := range nums {
		// ในแต่ละรอบของลูป, ค่าของ num
		// จะถูกเพิ่มลงในตัวแปร total
		total += num // total + num = total
	}
	// พิมพ์ผลรวมของค่าใน nums
	f.Println(total)
}
func main() {
	// การเรียกใช้ฟังก์ชัน sum
	// เรียกใช้ฟังก์ชัน sum โดยส่งอาร์กิวเมนต์ 2 ตัว คือ 1 และ 2
	sum(1, 2)
	// เรียกใช้ฟังก์ชัน sum โดยส่งอาร์กิวเมนต์ 3 ตัว คือ 1, 2, และ 3
	// ฟังก์ชัน sum จะคำนวณผลรวมและพิมพ์ผลลัพธ์ออกมา
	sum(1, 2, 3)

	// สร้าง slice ชื่อ nums ที่ประกอบด้วยค่าคือ [1, 2, 3, 4]
	nums := []int{1, 2, 3, 4}
	// เรียกใช้ฟังก์ชัน sum โดยส่ง slice nums เข้าไป โดย nums...
	// จะกระจายค่าภายใน slice นั้นเป็นอาร์กิวเมนต์แต่ละตัวให้ฟังก์ชัน sum
	sum(nums...)
}

/*สรุป
โค้ดนี้แสดงให้เห็นถึงการใช้งาน variadic function ในภาษา Go ซึ่ง
เป็นฟังก์ชันที่สามารถรับจำนวนอาร์กิวเมนต์ที่ไม่จำกัด โดยพารามิเตอร์
variadic จะถูกมองว่าเป็น slice ภายในฟังก์ชัน ทำให้สามารถดำเนินการ
ต่างๆ กับมันได้อย่างง่ายดาย นอกจากนี้ยังแสดงวิธีการส่ง slice เข้าไปใน
variadic function โดยการกระจายค่าภายใน slice นั้นเป็น
อาร์กิวเมนต์แต่ละตัว ฟีเจอร์นี้ช่วยให้ฟังก์ชันมีความยืดหยุ่นและใช้งานง่าย
ขึ้นมากในสถานการณ์ที่ต้องการรับค่าจำนวนมากๆ โดยไม่ต้องระบุจำนวน
อาร์กิวเมนต์ล่วงหน้า
*/
