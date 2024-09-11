package main

import (
	"encoding/xml" // แปลง XML ไปเป็นโครงสร้างข้อมูล
	"fmt"          // แสดงผล
	"os"           // จัดการไฟล์
)

// โครงสร้างข้อมูลของ Plant
type Plant struct {
	XMLName xml.Name `xml:"plant"`   // กําหนดชื่อ tag ให้กับโครงสร้างข้อมูล
	Id      int      `xml:"id,attr"` // กําหนดค่า attribute ให้กับโครงสร้างข้อมูล
	Name    string   `xml:"name"`    // กําหนดชื่อ tag ให้กับโครงสร้างข้อมูล
	Origin  string   `xml:"origin"`  // กําหนดชื่อ tag ให้กับโครงสร้างข้อมูล
}

// โครงสร้างที่มีการซ้อนกัน
type Garden struct {
	XMLName xml.Name `xml:"parent"`      // กําหนดชื่อ tag ให้กับโครงสร้างข้อมูล
	Plants  []Plant  `xml:"child>plant"` // กําหนดชื่อ tag ให้กับโครงสร้างข้อมูล
}

func main() {
	// สร้างข้อมูลตัวอย่าง Plant
	garden := &Garden{
		Plants: []Plant{
			{Id: 1, Name: "Rose", Origin: "Asia"},
			{Id: 2, Name: "Tulip", Origin: "Europe"},
			{Id: 3, Name: "Sunflower", Origin: "North America"},
		},
	}

	// MarshalIndent เพื่อสร้าง XML ที่อ่านง่าย
	// "  " คือช่องว่าง 2 ช่อง
	// "" คือชื่อ tag ให้กับโครงสร้างข้อมูล
	out, err := xml.MarshalIndent(garden, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err) // ถ้ามี error ให้แสดงผล
		os.Exit(1)                     //  ถ้ามี error ให้หยุดโปรแกรม
	}

	// เพิ่ม XML header ด้วยตัวเอง
	xmlHeader := []byte(xml.Header)
	xmlOutput := append(xmlHeader, out...)

	// แสดงผล XML ที่อ่านง่าย
	fmt.Println(string(xmlOutput))

	// Unmarshal ใช้สำหรับแปลง XML กลับไปเป็นโครงสร้างข้อมูล
	// โดยทำนให้มี tag ที่เป็น parent และ child
	input := []byte(`
		<parent>
			<child>
				<plant id="1">
					<name>Rose</name>
					<origin>Asia</origin>
				</plant>
				<plant id="2">
					<name>Tulip</name>
					<origin>Europe</origin>
				</plant>
				<plant id="3">
					<name>Sunflower</name>
					<origin>North America</origin>
				</plant>
				</child>
		</parent>
				`)

	var gardenData Garden                   //  กําหนดตัวแปรว่า gardenData จะเป็นโครงสร้างข้อมูล Garden
	err = xml.Unmarshal(input, &gardenData) // แปลง XML ไปเป็นโครงสร้างข้อมูล
	if err != nil {
		fmt.Printf("error: %v\n", err) // ถ้ามี error ให้แสดงผล
		os.Exit(1)                     //  ถ้ามี error ให้หยุดโปรแกรม
	}

	// แสดงผลโครงสร้างข้อมูล Go หลังจากแปลงจาก XML
	fmt.Printf("%+v\n", gardenData)

}

/* อธิบายโค้ด:
1.การใช้ Struct Tags เพื่อกำหนดการแมปกับ XML:
ฟิลด์ XMLName ใช้เพื่อกำหนดชื่อของ element ใน XML

ฟิลด์ id,attr จะกำหนดว่า Id จะเป็น attribute ใน
XML element แทนที่จะเป็น element ย่อย

ฟิลด์ Plants ถูกกำหนดด้วย tag child>plant เพื่อ
สร้างการซ้อนกันของ Plant ภายใต้ <parent><child><plant>...

2.การใช้ MarshalIndent เพื่อสร้าง XML:
ฟังก์ชัน xml.MarshalIndent ใช้ในการสร้าง XML ที่มี
การจัดย่อหน้าและเว้นวรรคให้อ่านง่าย
ส่วนหัวของ XML (xml.Header) ถูกเพิ่มด้วยตัวเองก่อนแสดงผล

3.การ Unmarshal ข้อมูล XML กลับมาเป็น Struct ใน Go:
ฟังก์ชัน xml.Unmarshal ใช้เพื่อแปลงข้อมูล XML ให้
กลับมาอยู่ในรูปแบบของโครงสร้าง Garden
ถ้า XML มีความผิดพลาดจะเกิด error ขึ้น
*/
/* แนวคิดสำคัญ:
Struct Tags ใช้ในการควบคุมวิธีการแมปข้อมูลระหว่าง Go และ XML
สามารถกำหนดให้ฟิลด์เป็น Attributes ใน XML หรือเป็น Nested Elements ได้
การตรวจสอบข้อผิดพลาด: xml.Unmarshal จะคืนค่าข้อผิดพลาด
ถ้าโครงสร้างของ XML ไม่ตรงกับ Struct ใน Go
*/
