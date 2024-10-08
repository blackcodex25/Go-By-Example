package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	p := fmt.Print         // ตัวแปร p เป็นการย่อคำสั่ง fmt.Print เพื่อให้เรียกใช้ได้สะดวก
	p(rand.IntN(100), ",") // สร้างตัวเลขสุ่มประเภท int ที่มีค่าในช่วง 0 ถึง 99
	p(rand.IntN(100))      // สร้างตัวเลขสุ่มประเภท int ที่มีค่าในช่วง 0 ถึง 99
	p()                    // เรียกใช้ p เพื่อเว้นบรรทัด

	p(rand.Float64()) // สร้างตัวเลขสุ่มประเภท float64 ที่มีค่าในช่วง 0 ถึง 1

	p((rand.Float64()*5)+5, ",") // สร้างตัวเลขสุ่มประเภท float64 ที่มีค่าในช่วง 5 ถึง 10
	p((rand.Float64() * 50) + 5) // สร้างตัวเลขสุ่มประเภท float64 ที่มีค่าในช่วง 5 ถึง 55

	// สร้างแหล่งข้อมูล (source) ของตัวเลขสุ่มแบบ PCG
	// โดยใช้ seed เป็นสองหมายเลข uint64 (42 และ 1024)
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)   // สร้างเลขสุ่มใหม่โดยใช้แหล่งข้อมูล PCG ที่กำหนด
	p(r2.IntN(100), ",") // สร้างตัวเลขสุ่มประเภท int ที่มีค่าในช่วง 0 ถึง 99
	p(r2.IntN(100))      // สร้างตัวเลขสุ่มประเภท int ที่มีค่าในช่วง 0 ถึง 99
	p()                  // เรียกใช้ p เพื่อเว้นบรรทัด

	// สร้างแหล่งข้อมูล (source) ของตัวเลขสุ่มแบบ PCG
	// โดยใช้ seed เป็นสองหมายเลข uint64 (42 และ 1024)
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)   // สร้างเลขสุ่มใหม่โดยใช้แหล่งข้อมูล PCG ที่กำหนด
	p(r3.IntN(100), ",") // สร้างตัวเลขสุ่มประเภท int ที่มีค่าในช่วง 0 ถึง 99
	p(r3.IntN(100))      // สร้างตัวเลขสุ่มประเภท int ที่มีค่าในช่วง 0 ถึง 99
	p()                  // เรียกใช้ p เพื่อเว้นบรรทัด

	/* การเปรียบเทียบผลลัพธ์:
	สร้างตัวเลขสุ่มจากแหล่งข้อมูลที่แตกต่างกันแต่ใช้ seed เดียวกัน (s2 และ s3)
	เพื่อแสดงให้เห็นว่าผลลัพธ์จะเหมือนกันหากใช้ seed เดียวกัน
	*/
}

/* รายละเอียดการทำงาน
1.การสร้างตัวเลขสุ่มพื้นฐาน:
fmt.Print(rand.IntN(100), ",")
fmt.Print(rand.IntN(100))
fmt.Println()

fmt.Println(rand.Float64())
rand.IntN(100) จะพิมพ์ตัวเลขสุ่มที่มีค่าในช่วง 0 ถึง 99
rand.Float64() จะพิมพ์ตัวเลขสุ่มที่มีค่าในช่วง 0.0 ถึง 1.0

2.การสร้างตัวเลขสุ่มในช่วงที่กำหนด:
fmt.Print((rand.Float64()*5)+5, ",")
fmt.Print((rand.Float64() * 5) + 5)
fmt.Println()
ใช้ rand.Float64() เพื่อสร้างตัวเลขสุ่มในช่วง 0.0 ถึง 1.0 และแปลงเป็นช่วง 5.0 ถึง 10.0
ด้วยการคูณและบวกค่า

3.การใช้ rand.NewPCG:
s2 := rand.NewPCG(42, 1024)
r2 := rand.New(s2)
fmt.Print(r2.IntN(100), ",")
fmt.Print(r2.IntN(100))
fmt.Println()

s3 := rand.NewPCG(42, 1024)
r3 := rand.New(s3)
fmt.Print(r3.IntN(100), ",")
fmt.Print(r3.IntN(100))
fmt.Println()
rand.NewPCG(42, 1024) สร้างแหล่งข้อมูลสุ่ม PCG ใหม่ด้วย seed 42 และ 1024
rand.New(s2) สร้างตัวสร้างเลขสุ่มที่ใช้แหล่งข้อมูล PCG s2
ผลลัพธ์ของ r2.IntN(100) และ r3.IntN(100) จะเหมือนกันถ้าใช้ seed และแหล่งข้อมูลเดียวกัน
*/
/* สรุป
โค้ดนี้แสดงให้เห็นถึงการใช้งานของ math/rand/v2 สำหรับการสร้างตัวเลขสุ่มพื้นฐาน
และการใช้งาน rand.NewPCG เพื่อสร้างตัวเลขสุ่มที่สามารถคาดเดาได้ (reproducible)
ด้วยการใช้ seed ที่ระบุ ผลลัพธ์ของตัวเลขสุ่มจะคงที่ถ้าใช้ seed เดียวกัน
โดยแสดงให้เห็นว่าการใช้ seed เดียวกันจะให้ผลลัพธ์เดียวกันในแหล่งข้อมูลเดียวกัน
*/
