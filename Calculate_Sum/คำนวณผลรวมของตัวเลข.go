// Calculates sum of all multiple of 3 and 5 less than MAX value.
// See https://projecteuler.net/problem=1

/*
	โค้ดนี้คำนวณผลรวมของตัวเลขที่เป็นผลคูณของ 3 และ 5 ทั้งหมด

ที่น้อยกว่า 1000 โดยใช้หลักการทำงานแบบ concurrent ผ่าน
goroutine และ channel เพื่อเพิ่มประสิทธิภาพในการคำนวณ
*/
package main

import (
	"fmt"
)

// กำหนดค่า MAX ที่เป็นตัวเลขสูงสุดที่เราจะใช้ตรวจสอบผลคูณ
const MAX = 1000

func main() {
	// สร้าง channel 'work' เพื่อเก็บผลคูณของ 3 และ 5
	work := make(chan int, MAX)
	// สร้าง channel 'result' เพื่อเก็บผลลัพธ์ของผลรวม
	result := make(chan int)

	// 1. ใช้ goroutine เพื่อหาผลคูณของ 3 และ 5
	go func() {
		// วนซ้ำตั้งแต่ 1 ถึงน้อยกว่า MAX
		for i := 1; i < MAX; i++ {
			// fmt.Printf("%v\n", i) ดูการวนซ้ำใน Loop แต่ระรอบ
			// ตรวจสอบว่า i เป็นผลคูณของ 3 หรือ 5
			if (i%3) == 0 || (i%5) == 0 {
				work <- i // ส่งค่า i ไปที่ Channel work ถ้าเป็นผลคูณ
			}
		}
		// ปิด Channel work เมื่อหาค่าผลคูณเสร็จ
		close(work)
	}()

	// 2. ใช้ goroutine อีกตัวเพื่อคำนวณผลรวมจาก Channel work
	go func() {
		r := 0 // กำหนดตัวแปร r สำหรับเก็บผลรวม
		// รับค่าจาก Channel work จนกว่าจะไม่มีค่าเหลือ
		for i := range work {
			r = r + i // บวกค่า i ที่ได้รับจาก Channel work เข้าไปใน r
		}
		// ส่งค่าผลรวมไปที่ Channel result
		result <- r
	}()

	// 3. รอรับผลลัพธ์จาก Channel result แล้วพิมพ์ผลลัพธ์ออกมา
	fmt.Println("Total:", <-result)
}
