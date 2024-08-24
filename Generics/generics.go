package main

import (
	f "fmt"
)

/*โค้ดนี้แสดงการใช้งานฟังก์ชันและประเภทข้อมูลทั่วไป (generic
functions and types) ในภาษา Go โดยผสมผสานการใช้ generic กับ
การจัดการข้อมูลในลิสต์เชื่อมโยง (linked list)
*/
/*ฟังก์ชันนี้รับพารามิเตอร์ s ซึ่งเป็นสไลซ์ของชนิดข้อมูล
E และ v ซึ่งเป็นค่าที่จะค้นหาในสไลซ์นั้น ฟังก์ชันจะวนลูป
ผ่านสไลซ์ s และตรวจสอบว่าค่าที่กำหนด v ตรงกับค่าที่
ตำแหน่งใด หากพบจะคืนค่าดัชนีนั้น (index) หากไม่พบจะคืนค่าเป็น -1
*/
// ข้อจำกัดของชนิดข้อมูล:
// *S ~[]E: หมายถึง S เป็นชนิดข้อมูลที่เป็นสไลซ์ของ E
// *E comparable:
// *หมายถึง E เป็นชนิดข้อมูลที่สามารถเปรียบเทียบได้ (ใช้กับ == และ !=)
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// การกำหนดประเภทข้อมูลทั่วไป (Generic Types)
/* List:
ประเภท List เป็นลิสต์เชื่อมโยงเดี่ยว (singly-linked list)
ที่เก็บค่าของชนิดข้อมูลใดก็ได้ (indicated by T any) */
type List[T any] struct {
	head, tail *element[T]
}

// element:
/*โครงสร้าง element เป็นโหนดในลิสต์เชื่อมโยงเดี่ยว
ที่มีฟิลด์ val ซึ่งเก็บค่าของชนิด T และฟิลด์ next ซึ่งชี้ไป
ยังโหนดถัดไป
*/
type element[T any] struct {
	val  T
	next *element[T]
}

/*
Logic: ฟังก์ชัน Push ใช้ในการเพิ่มค่า v เข้าไปที่ท้ายของลิสต์
หากลิสต์ยังว่าง (ไม่มี tail) จะสร้างโหนดแรกและกำหนดให้
เป็น head และ tail หากมีโหนดอยู่แล้ว จะเพิ่มโหนดใหม่ที่ท้ายของลิสต์
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
Logic: ฟังก์ชันนี้จะคืนค่าสมาชิกทั้งหมดในลิสต์เป็นสไลซ์ โดยจะ
วนลูปผ่านโหนดทั้งหมดในลิสต์แล้วเก็บค่าของโหนดในสไลซ์ elems
*/
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

/*
การใช้งาน SlicesIndex:
มีการสร้างสไลซ์ s ที่เก็บสตริงและเรียกใช้ฟังก์ชัน SlicesIndex เพื่อ
ค้นหาดัชนีของคำว่า "zoo" ในสไลซ์ โดยจะพิมพ์ค่าที่ได้ออกมา

การสร้างและใช้งาน List:
มีการสร้างลิสต์ของตัวเลข
(List[int]), เพิ่มค่า 10, 13, และ 23 ลงในลิสต์
จากนั้นเรียกฟังก์ชัน AllElements เพื่อแสดงสมาชิกทั้งหมดในลิสต์
*/
func main() {
	var s = []string{"foo", "bar", "zoo"}

	f.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// การละเว้นค่า index โดยใช้ _
	_ = SlicesIndex[[]string, string](s, "zoo")

	// หากเราต้องการเก็บค่าที่คืนกลับมาจากฟังก์ชัน sliceIndex แทน
	// การละเว้นค่า เราสามารถเก็บค่าดัชนีที่ได้ในตัวแปร เช่น index แทน
	// การใช้ละเว้นค่า _
	/*index := SlicesIndex(s, "zoo")
	f.Println("index of zoo:", index) */

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	f.Println("list:", lst.AllElements())
}

/*สรุป Logic
ฟังก์ชัน SlicesIndex: ค้นหาดัชนีของค่าที่กำหนดในสไลซ์
ประเภท List: ใช้จัดการลิสต์เชื่อมโยงเดี่ยวที่สามารถเก็บค่าของชนิดข้อมูลใดก็ได้
ฟังก์ชัน Push และ AllElements: ใช้ในการเพิ่มสมาชิกในลิสต์
และดึงข้อมูลสมาชิกทั้งหมดออกมา

โค้ดนี้เป็นตัวอย่างที่ดีในการแสดงวิธีการใช้ generic ใน Go เพื่อสร้าง
โค้ดที่ยืดหยุ่นและนำไปใช้ซ้ำได้.
*/
