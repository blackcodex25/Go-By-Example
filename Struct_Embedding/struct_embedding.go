package main

import (
	f "fmt"
)

/*
การใช้คุณสมบัติของ Go ในการฝัง
(embedding) โครงสร้างข้อมูล (structs) และอินเทอร์เฟซ
(interfaces) เพื่อสร้างการจัดองค์ประกอบของชนิดข้อมูล (types)
ที่มีความราบรื่นมากขึ้น คำว่า embedding ในที่นี้ไม่ควรสับสนกับ
คำสั่ง //go:embed ซึ่งถูกแนะนำใน Go เวอร์ชัน 1.16+ ที่ใช้ใน
การฝังไฟล์และโฟลเดอร์เข้าไปในไฟล์ไบนารีของแอปพลิเคชัน

ในการฝังโครงสร้างข้อมูล โครงสร้างหลัก (container) จะฝัง
โครงสร้างพื้นฐาน (base) การฝังนี้จะปรากฏเหมือนฟิลด์ที่ไม่มีชื่อ
(anonymous field) เมื่อต้องการสร้างโครงสร้างข้อมูลด้วยค่าคงที่
(literals) เราจะต้องกำหนดค่าให้กับฟิลด์ที่ถูกฝังไว้โดยตรง
ซึ่งฟิลด์นี้จะใช้ชื่อของชนิดข้อมูลที่ถูกฝังเป็นชื่อฟิลด์

เราสามารถเข้าถึงฟิลด์ของโครงสร้างพื้นฐาน (base) ได้โดยตรง
ผ่านโครงสร้างหลัก (container) ตัวอย่างเช่น co.num หรือเรา
สามารถระบุชื่อชนิดข้อมูลที่ถูกฝังไว้เต็มๆ ได้ เช่น co.base.num

เนื่องจากโครงสร้างหลัก (container) ฝังโครงสร้างพื้นฐาน (base)
ฟังก์ชัน (methods) ที่อยู่ในโครงสร้างพื้นฐานก็จะกลายเป็น
ฟังก์ชันของโครงสร้างหลักด้วย ดังนั้นเราจึงสามารถเรียกใช้งาน
ฟังก์ชันที่ฝังมาจากโครงสร้างพื้นฐานได้โดยตรงผ่านโครงสร้างหลัก (container)

การฝังโครงสร้างข้อมูลที่มีฟังก์ชันสามารถใช้ในการมอบการใช้
งานอินเทอร์เฟซ (interface implementations) ให้กับโครงสร้าง
ข้อมูลอื่นได้ เช่นในตัวอย่างนี้ โครงสร้างหลัก (container)
สามารถใช้งานอินเทอร์เฟซ describer ได้เพราะฝังโครงสร้างพื้น
ฐาน (base) ไว้นั่นเอง
*/
/* 1. การกำหนดโครงสร้าง (Structs)
โครงสร้าง base มีฟิลด์ num ซึ่งเป็นตัวเลขชนิด int
และมีฟังก์ชัน describe ที่อธิบายค่า num ในรูปแบบของสตริง
*/
type base struct {
	num, enum int
}

/*
ฟังก์ชัน describe รับค่าพารามิเตอร์ชนิด base
และคืนค่าเป็นสตริงที่แสดงค่า num ในโครงสร้าง base
*/
func (b base) describer() string {
	return f.Sprintf("base with num=%v, %v", b.num, b.enum)
}

/*
โครงสร้าง container ฝังโครงสร้าง base
ไว้โดยไม่ระบุชื่อฟิลด์ (anonymous field)
และมีฟิลด์เพิ่มเติมคือ str ซึ่งเป็นชนิด string
*/
type container struct {
	base
	str string
}

func main() {
	/*การสร้างตัวแปร co ซึ่งเป็นชนิด container โดยกำหนดค่า num
	ใน base เป็น 1 และกำหนดค่า str เป็น "some name"
	*/
	co := container{
		base: base{
			num:  1,
			enum: 2,
		},
		str: "some name",
	}

	/*แสดงผลค่า num และ str ของ co โดยสามารถเข้าถึง num
	ได้โดยตรงผ่านตัวแปร co แม้ว่า num จะอยู่ในโครงสร้าง base
	ที่ถูกฝังไว้
	*/
	// fmt.Printf call needs 2 args but has 3 args เพิ่มการจัดรูปแบบ (Formatting Verbs)
	// %v ค่าเริ่มต้น และเรียกใช้ co.enum
	f.Printf("co={num: %v, %v, str: %v}\n", co.num, co.enum, co.str)

	/*การเข้าถึงฟิลด์ num ใน base สามารถทำได้ผ่าน co.base.num
	เช่นกัน
	*/
	f.Println("also num:", co.base.num, co.base.enum)

	/* เรียกใช้งานฟังก์ชัน describe ที่ฝังมาจาก base ผ่าน co โดยตรง
	ฟังก์ชันนี้จะส่งกลับสตริงที่อธิบายค่า num ใน base
	*/
	f.Println("describe:", co.describer())

	/*สร้างอินเทอร์เฟซ describer ที่มีฟังก์ชัน describe
	และกำหนดตัวแปร d ชนิด describer ให้เป็น co
	*/
	type describer interface {
		describer() string
	}

	/*เนื่องจาก co ฝัง base ซึ่งมีฟังก์ชัน describe ทำให้ co
	สามารถใช้เป็น describer ได้
	เมื่อเรียก d.describe() จะได้ผลลัพธ์เป็นสตริงที่อธิบายค่า num
	ใน base
	*/
	var d describer = co
	f.Println("describe:", d.describer())
}

/* Logic!
การฝังโครงสร้าง:
การฝังโครงสร้าง base ใน container
ทำให้สามารถเข้าถึงฟิลด์และฟังก์ชันของ base
ผ่านตัวแปร container ได้โดยตรง

การใช้งานอินเทอร์เฟซ:
ด้วยการฝังโครงสร้างที่มีฟังก์ชันที่ตรงกับฟังก์ชันในอินเทอร์เฟซ
โครงสร้างหลักสามารถทำงานกับอินเทอร์เฟซได้ทันที

โค้ดนี้แสดงให้เห็นถึงความยืดหยุ่นและความสะดวกในการใช้งาน
โครงสร้างข้อมูลและอินเทอร์เฟซใน Go เพื่อสร้างโค้ดที่มีความกระชับ
และมีประสิทธิภาพ
*/
