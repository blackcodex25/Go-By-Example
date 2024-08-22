package main

import (
	f "fmt"
)

/* การใช้งาน Enumerated Types (Enums) ในภาษา Go
ภาษา Go ไม่มี enum ในรูปแบบคุณสมบัติที่เป็นเอกลักษณ์
ของภาษาเช่นในบางภาษาอื่นๆ แต่ enum สามารถสร้างได้
อย่างง่ายดายโดยใช้วิธีการและแนวปฏิบัติที่มีอยู่ในภาษา Go
โดยทั่วไป enum คือชนิดข้อมูลที่มีค่าที่เป็นไปได้เฉพาะ ซึ้ง
แต่ละค่าจะมีชื่อที่แตกต่างกันไป

คุณสมบัติและการใช้งาน Enumerated Types
1.การประกาศ Enumerated Type
ใน Go เราสามารถสร้าง enum โดยการประกาศชนิด
ข้อมูลใหม่ที่มีฐานเป็นชนิดข้อมูลที่มีอยู่ เช่น int
ในตัวอย่างนี้ ServerState ถูกประกาศเป็น enum
โดยใช้ type ServerState int

2.การกำหนดค่าที่เป็นไปได้สำหรับ Enumerated Type
ค่าที่เป็นไปได้ของ ServerState ถูกกำหนดเป็น
constants โดยใช้คำสั่ง const
คำสั่ง iota ถูกใช้เพื่อสร้างค่าคงที่ (constant
values) แบบต่อเนื่องกันอัตโนมัติ โดยเริ่มจาก
0 ซึ่งทำให้การสร้าง enum ง่ายขึ้น

3.การแปลง enum เป็นสตริงด้วยการใช้งาน
fmt.Stringer Interfaces
การใช้งาน fmt.Stringer interface ช่วยให้ค่าของ
ServerState สามารถถูกแปลงเป็นสตริงและพิมพ์
ออกมาได้ง่าย

4.ความปลอดภัยของชนิดข้อมูลในช่วง Compile-Time
หากมีการพยายามใช้ค่าที่ไม่ใช่ชนิด ServerState
กับฟังก์ชันที่ต้องการค่าชนิด SeverState เช่น
transition คอมไพเลอร์จะทำการแจ้งเตือนว่าเกิด
ข้อผิดพลาดเกี่ยวกับชนิดข้อมูล ซึ่งช่วยให้โปรแกรมมี
ความปลอดภัยมากขึ้นในช่วงการคอมไพล์ (compile-time)

5.การเปลี่ยนสถานะ (State Transition)
ฟังก์ชัน transition จำลองการเปลี่ยนแปลง
สถานะของเซิร์ฟเวอร์ โดยรับสถานะปัจจุบัน
การเปลี่ยนแปลงสถานะ (transition) ถูกควบคุม
โดย switch statement ซึ่งจะเลือกสถานะถัดไป
ตามสถานะปัจจุบัน
*/
// การประกาศ ServerState และกำหนดค่าที่เป็นไปได้
// ServerState ถูกประกาศเป็น int ที่เป็นฐานของ enum
type ServerState int

// ค่าที่เป็นไปได้ของ Server State ถูกกำหนดโดยใช้
// const และ iota ซึ่งจะกำหนดค่าให้ StateIdle = 0
// StateConnected = 1
// StateError = 2
// และ StateRetrying = 3
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// การแปลงค่าของ ServerState เป็นสตริง
// map ชื่อ stateName ถูกใช้เพื่อเก็บชื่อของแต่ละ
// สถานะในรูปแบบของสตริง
// ฟังก์ชัน String() บน ServerState ถูกใช้เพื่อ
// แปลงสถานะให้เป็นสตริง โดยคืนค่าจาก map
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// การจำลองการเปลี่ยนแปลงสถานะในฟังก์ชัน
func transition(s ServerState) ServerState {
	// transition ใช้ switch statement เพื่อเลือก
	// สถานะถัดไปตามสถานะปัจจุบัน
	switch s {
	// ถ้าสถานะปััจจุบันเป็น StateIdle,
	case StateIdle:
		// สถานะใหม่จะเป็น StateConnected
		return StateConnected
		// ถ้าสถานะปัจจุบันเป็น StateError หรือ StateRetrying
	case StateConnected, StateRetrying:
		// สถานะใหม่จะเป็น StateIdle
		return StateIdle
		// ถ้าสถานะปัจจุบันเป็น StateError
	case StateError:
		// สถานะจะไม่เปลี่ยนแปลง
		return StateError
		// ถ้าพบสถานะที่ไม่รู้จัก (unknown state)
		// โปรแกรมจะหยุดทำงานโดยใช้ panic
	default:
		panic(f.Errorf("unknown state: %s", s))

	}
}
func main() {
	// เรียกใช้ฟังก์ชัน transition โดยส่ง StateIdle เข้าไป
	// และผลลัพธ์ของการเปลี่ยนแปลงสถานะ (transition)
	// จะถูกพิมพ์ออกมา
	ns := transition(StateIdle)
	f.Println(ns)

	// ฟังก์ชัน transition ถูกเรียกใช้ซ้ำกับสถานะใหม่ที่
	// ได้รับจากครั้งแรก และผลลัพธ์จะถูกพิมพ์ออกมาเช่นกัน
	ns2 := transition(ns)
	f.Println(ns2)
}

/* สรุป
โค้ดนี้แสดงการใช้งาน enumerated types (enums) ในภาษา
Go โดยใช้วิธีการที่มีอยู่ในภาษา Go เช่นการใช้ const และ
iota เพื่อกำหนดค่าคงที่สำหรับ enum การแปลงค่าของ
enum การใช้ enum ใน Go ช่วยให้โค้ดมีความชัดเจนและ
จัดการง่ายขึ้น โดยเฉพาะในกรณีที่มีค่าที่เป็นไปได้จำนวนมากที่
ต้องการการจัดการที่เป็นระเบียบ
*/
