package main

import (
	f "fmt"
)

/* บทความ การใช้งาน range ในภาษา Go
range เป็นคำสั่งที่ใช้สำหรับการวนลูป (literate) ผ่านองค์ประกอบของ
โครงสร้างข้อมูลต่างๆ ในภาษา Go เช่น slice, array, map, และ string
การใช้ range ช่วยให้การเข้าถึงทั้งค่าขององค์ประกอบ (value) และตำแน่ง
ขององค์ประกอบ (index) ในโครงสร้างข้อมูลนั้นเป็นเรื่องง่าย
*/
/* คุณสมบัติและการใช้งาน range
1.การใช้งาน range กับ slice และ array
เมื่อใช้ range กับ slice หรือ array เราจะได้รับทั้ง index และ
value ของแต่ละองค์ประกอบใน slice หรือ array
ถ้าไม่ต้องการใช้ index เราสามารถละทิ้งมันโดยใช้เครื่องหมาย _
เพื่อเป็นตัวแทน
2.การใช้งาน range กับ map
เมื่อใช้ range กับ map เราจะได้รับทั้ง key และ value ของ
แต่ละคู่ key-value ใน map
เรายังสามารถเลือกที่จะวนลูปเฉพาะ key ของ map ได้ด้วย
3.การใช้งาน range กับ string
เมื่อใช้ range กับ string เราจะได้รับ index ของ byte แรกใน
rune และตัว rune
สิ่งนี้มีประโยชน์เมื่อทำงานกับ string ที่ประกอบด้วยหลายๆ rune
หรือ Unicode code points
*/
func main() {
	// การใช้งาน range กับ slice เพื่อหาผลรวม
	// สร้าง slice ชื่อ nums ที่ประกอบด้วยค่า [2,3,4]
	nums := []int{2, 3, 4}
	sum := 0
	// ใช้ range เพื่อวนลูปผ่านแต่ value ใน slice
	// โดยละทิ้ง index (ใช้อันเดอร์สกอร์ _ )
	for _, num := range nums {
		// ในแต่ละรอบของลูป sum + num = sum จะเพิ่มค่าของ num
		// ลงในตัวแปร sum
		sum += num
	}
	// พิมพ์ผลรวมของค่าใน slice ซึ่งในกรณีนี้คือ 9
	f.Println("sum:", sum)

	// การใช้งาน range เพื่อค้นหาตำแหน่ง (index) ของค่าใน slice
	//  ใช้ range เพื่อรับทั้ง index (i) และ value (num)
	for i, num := range nums {
		// ถ้า num เท่ากับ 3
		if num == 3 {
			// พิมพ์ตำแหน่งของมันออกมา
			f.Println("index:", i)
		}
	}

	// การใช้งาน range กับ map เพื่อแสดงคู่ key-value
	// สร้าง map ชื่อ kvs ที่มีคู่ key-value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// ใช้ range เพื่อวนลูปผ่านแต่ละคู่ key-value ใน map
	for k, v := range kvs {
		// พิมพ์ค่าของแต่ละคู่ในรูปแบบ "key -> value"
		// เช่น "a -> apple" และ "b -> banana"
		f.Printf("%s -> %s\n", k, v)
	}

	// ใช้ range เพื่อวนลูปผ่านเฉพาะ key ใน map
	for k := range kvs {
		// พิมพ์ค่า key แต่ละตัวใน map เช่น
		// "key: a" และ "key: b"
		f.Println("key:", k)
	}

	// *การใช้งาน range กับ string
	// *ใช้ range เพื่อวนลูปผ่าน string "go"
	// *i คือ index ของ byte แรกในแต่ละ rune และ c
	// *คือ rune เอง
	for i, c := range "go" {
		/* พิมพ์ index และ rune สำหรับแต่ละตัวอักษรใน
		string เช่น 0 103 และ 1 111 (103 และ 111
		คือ Unicode code points ของ g และ o ตามลำดับ) */
		f.Println(i, c)
	}
}

/* สรุป
โค้ดนี้แสดงให้เห็นถึงความยืดหยุ่นและประโยชน์ของการใช้ range ในการรวมลูปผ่านโครงสร้าง
ข้อมูลต่างๆ ในภาษา Go ไม่ว่าจะเป็น slice, map, หรือ string, range ช่วยให้เราสามารถ
้เข้าถึงข้อมูลได้อย่างสะดวกทั้งในรูปแบบของ index, value, key, และ rune, ทำให้การจัด
การข้อมูลใน Go เป็นเรื่องง่ายและตรงไปตรงมา
*/
