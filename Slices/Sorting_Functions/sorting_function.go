package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
โค้ดนี้แสดงวิธีใช้ฟังก์ชันเปรียบเทียบแบบกำหนดเองเพื่อจัดเรียง
ข้อมูลตามเกณฑ์ที่เราต้องการ
*/
func main() {
	// สร้าง slice fruits เก็บค่า string คือ peach, banana, และ kiwi
	fruits := []string{"peach", "banana", "kiwi"}

	// สร้างฟังก์ชัน lenCmp เพื่อจัดเรียง string โดยเปรียบความยาวของ string
	lenCmp := func(a, b string) int {
		// คืนค่าเป็น -1, 0, หรือ 1 ขึ้นอยู่กับว่า
		// ความยาวของ a น้อยกว่า เท่ากับ หรือมากกว่า b
		return cmp.Compare(len(a), len(b))
	}

	// ใช้ SortFunc เพื่อจัดเรียง fruits ตามฟังก์ชันเปรียบเทียบ lenCmp
	slices.SortFunc(fruits, lenCmp)

	// หลังจากจัดเรียงแล้ว จะได้ผลลัพธ์ที่เรียงลำดับตามความยาว
	// ของ string จากน้อยไปมาก
	fmt.Println(fruits)

	// สร้างโครงสร้างข้อมูล Struct Person ที่มีฟิลด์ name ชนิด string และ age ชนิด int
	type Person struct {
		name string
		age  int
	}

	// สร้าง slice ที่ชื่อว่า people ซึ่งเก็บข้อมูล Person 3 คนคือ
	// "Jax", "TJ", และ "Alex" พร้อมกับอายุของแต่ละคน
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// ใช้ slices.SortFunc เพื่อจัดเรียง people โดยใช้ฟังก์ชัน
	// เปรียบเทียบที่เปรียบเทียบฟิลด์ age ของ Person
	slices.SortFunc(people, func(a, b Person) int {
		// ฟังก์ชันเปรียบเทียบนี้จะเรียกใช้ cmp.Compare(a.age,
		// b.age) ซึ่งทำให้ people ถูกจัดเรียงตามอายุจากน้อยไปมาก
		return cmp.Compare(a.age, b.age)
	})

	// ผลลัพธ์จะแสดง slice people ที่จัดเรียงตามอายุจากน้อยไปมาก
	fmt.Println(people)
}

/* สรุปการทำงานของโค้ด
จัดเรียง string: จัดเรียง string ใน slice fruits โดยใช้
ฟังก์ชันเปรียบเทียบที่เปรียบเทียบตามความยาวของ string
จัดเรียง struct: จัดเรียง slice ของ struct Person โดยเรียง
ตามอายุของแต่ละคน
*/
