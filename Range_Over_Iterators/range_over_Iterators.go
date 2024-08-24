package main

import (
	f "fmt"
	"iter"
	"slices"
)

/*โค้ดนี้แสดงให้เห็นการใช้งาน iterators ในภาษา Go โดยนำ
เสนอผ่านการสร้างและใช้งานโครงสร้างข้อมูล List ที่รองรับ
Generic และการสร้างลำดับตัวเลข Fibonacci ที่ไม่มีที่สิ้นสุด
ผ่าน iterator
*/
/* List และ element เป็นโครงสร้างข้อมูลเชื่อมโยง
(linked list) ที่ใช้สำหรับเก็บค่าของชนิดข้อมูลทั่วไป (T any)
List ประกอบด้วย head และ tail ซึ่งเป็น pointer
ไปยัง element ตัวแรกและตัวสุดท้ายใน list
*/
type List[T any] struct {
	head, tail *element[T]
}

// element มี val สำหรับเก็บค่าข้อมูลและ next สำหรับชี้ไปยัง element ถัดไป
type element[T any] struct {
	next *element[T]
	val  T
}

/*
Method Push ของ List
เมธอด Push ทำหน้าที่เพิ่มค่าใหม่ (v) เข้าไปใน list
ถ้า tail เป็น nil แสดงว่า list ว่างเปล่า ดังนั้นจะสร้าง
head และ tail เป็น element ใหม่ที่มีค่า v
ถ้า tail มีค่าแล้ว จะเพิ่ม element ใหม่เข้าไปใน
ตำแหน่งถัดไปของ tail แล้วอัปเดต tail ให้ชี้ไปยัง
element ที่เพิ่มเข้ามาใหม่
*/
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

/*
Method All ของ List
เมธอด All คืนค่า iter.Seq[T] ซึ่งเป็น iterator สำหรับ List
ใน Go, iterator นี้ถูกนิยามเป็นฟังก์ชันที่รับอีกฟังก์ชันหนึ่ง
(yield) ที่จะถูกเรียกสำหรับแต่ละ element ใน List
ลูปจะทำงานไปเรื่อย ๆ ตั้งแต่ head ไปจนถึง tail โดยเรียก yield สำหรับแต่ละค่า
ถ้า yield คืนค่าเป็น false การวนลูปจะหยุด
*/
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

/*
ฟังก์ชัน genFib
genFib สร้าง iterator ที่สร้างลำดับ Fibonacci
ฟังก์ชันนี้ใช้ลูปที่ไม่มีที่สิ้นสุด (for {}) ในการสร้างตัวเลข Fibonacci
ลำดับจะสร้างไปเรื่อยๆ ตราบเท่าที่ yield คืนค่า true
*/
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	// สร้าง List และเพิ่มค่าลงไปใน List โดยใช้ Push
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// ใช้ range loop เพื่อวนลูปผ่านค่าใน List ผ่าน
	// iterator ที่คืนค่าจากเมธอด All
	for e := range lst.All() {
		f.Println(e)
	}

	// ใช้ฟังก์ชัน slices.Collect ในการรวบรวมค่าทั้งหมดใน
	// List ลงใน slice
	all := slices.Collect(lst.All())
	f.Println("all:", all)

	// ใช้ genFib เพื่อสร้างลำดับ Fibonacci และวนลูปเพื่อ
	// พิมพ์ค่าจนกว่าค่าจะเกินหรือเท่ากับ 10
	for n := range genFib() {
		if n >= 10 {
			break
		}
		f.Println(n)
	}

}

/* สรุป
โครงสร้าง List เป็นโครงสร้างข้อมูลแบบเชื่อมโยงที่รองรับชนิดข้อมูลทั่วไป
Iterators ถูกสร้างโดยการนิยามฟังก์ชันที่รับฟังก์ชัน
yield เพื่อวนลูปผ่านข้อมูลใน List
การใช้งาน iterator ใน Go ทำได้โดยใช้ range loop
ซึ่งช่วยให้การประมวลผลข้อมูลในคอลเล็กชันหรือการสร้าง
ลำดับไม่มีที่สิ้นสุดทำได้อย่างง่ายดาย
*/
