package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// สร้างโครงข้อมูล response1 ที่มี 2 ฟิลด์
// Page: จำนวนเต็มที่แทนหมายเลขหน้า
// Fruits: Slice ของ string ที่แทนรายชื่อผลไม้
type response1 struct {
	Page   int
	Fruits []string
}

// สร้างโครงข้อมูล response2 ที่มี 2 ฟิลด์
// Page: จำนวนเต็มที่แทนหมายเลขหน้า
// Fruits: Slice ของ string ที่แทนรายชื่อผลไม้
type response2 struct {
	Page   int      `json:"page"`   // เพิ่มค่าใน tag json
	Fruits []string `json:"fruits"` // เพิ่มค่าใน tag json
}

/* ความแตกต่างระหว่าง response1 และ response2 คือ
response2 ใช้ tag json:"..." เพื่อกำหนดชื่อฟิลด์ใน
JSON ที่จะสื่อสารกับภายนอก
*/
/*
func main แสดงวิธีการใช้งาน json.Marshal และ json.Unmarshal
เพื่อแปลงโครงสร้างข้อมูลประเภทต่างๆ ให้เป็นรูปแบบ JSON
และแปลงรูปแบบ JSON กลับเป็นโครงสร้างข้อมูลประเภทต่างๆ
*/
func main() {
	// Marshaling (แปลงเป็น JSON)
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("Gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Unmarshaling (แปลงจาก JSON กลับมาเป็นโครงสร้างข้อมูล Go)
	byt := []byte(`{"num":6.33, "strs": ["a", "b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// การใช้ Struct เพื่อ Unmarshal
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// การเขียน JSON ไปที่ Output (Encode)
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

/* โค้ดนี้มีการทำงานเกี่ยวกับการ Marshaling (แปลงโครงสร้างข้อมูลเป็น JSON)
และ Unmarshaling (แปลง JSON กลับมาเป็นโครงสร้างข้อมูล)
โดยใช้แพ็กเกจ encoding/json ของภาษา Go ซึ่งโค้ดจะทำตามลำดับดังนี้:
*/

/* ส่วนประกอบหลักของโค้ด:
1.การประกาศโครงสร้าง (Struct)
type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}
มีการสร้างโครงสร้างข้อมูล response1 และ response2
ที่มีฟิลด์ Page และ Fruits เพื่อใช้ในการจัดการข้อมูล
เกี่ยวกับหน้าและรายชื่อผลไม้

ความแตกต่างระหว่าง response1 และ response2 คือ
response2 ใช้ tag json:"..." เพื่อกำหนดชื่อฟิลด์ใน JSON
ที่จะสื่อสารกับภายนอก

2.Marshaling การแปลงโครงสร้างเป็น JSON
bolB, _ := json.Marshal(true)
fmt.Println(string(bolB))
ในส่วนนี้โค้ดจะแปลงค่าต่างๆ เช่น
Boolean (true), Integer, Float, String, Slice, Map,
และ Struct ให้กลายเป็น JSON
ใช้ฟังก์ชัน json.Marshal() ซึ่งจะคืนค่าข้อมูลในรูปแบบ []byte (บิตข้อมูล)
ซึ่งเราจะแปลงเป็น string เพื่อนำมาแสดงผล

3.Unmarshaling การแปลง JSON กลับมาเป็นโครงสร้างข้อมูล
byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
var dat map[string]interface{}
if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
}
โค้ดจะใช้ฟังก์ชัน json.Unmarshal() เพื่อแปลงข้อมูล JSON
กลับมาเป็นโครงสร้างข้อมูล Go โดยข้อมูล JSON ที่แปลงจะถูกเก็บไว้ใน
map[string]interface{} เนื่องจากเราไม่ทราบประเภทข้อมูลแน่นอนล่วงหน้า

หลังจากแปลง JSON มาเป็น map แล้ว เราสามารถเข้าถึง
ค่าภายใน map โดยใช้ key เช่น dat["num"] และ dat["strs"]

4.การใช้ Struct เพื่อ Unmarshal
str := `{"page": 1, "fruits": ["apple", "peach"]}`
res := response2{}
json.Unmarshal([]byte(str), &res)
fmt.Println(res)

ในตัวอย่างนี้ JSON ถูกแปลงกลับมาเป็นโครงสร้างข้อมูล response2
โดยตรง ซึ่งจะช่วยให้การจัดการกับข้อมูล JSON
ที่เรารู้โครงสร้างแน่นอนเป็นไปได้สะดวกมากขึ้น

5.การเขียน JSON ไปที่ Output (Encode)
enc := json.NewEncoder(os.Stdout)
d := map[string]int{"apple": 5, "lettuce": 7}
enc.Encode(d)
ฟังก์ชัน json.NewEncoder() ถูกใช้เพื่อแปลงข้อมูล Go
ให้เป็น JSON แล้วเขียนออกไปที่ os.Stdout (จอแสดงผล) โดยตรง
*/
/* หลักการทำงาน
1.Marshaling:
โค้ดจะนำข้อมูลจาก Go เช่น Boolean, Number, String, Slice, Map
หรือ Struct มาแปลงเป็น JSON โดยใช้ json.Marshal()
ซึ่งจะทำให้ข้อมูล Go เป็นรูปแบบที่สามารถส่งไปยังระบบอื่นๆ
หรือจัดเก็บในไฟล์ JSON ได้

2.Unmarshaling:
โค้ดจะอ่าน JSON และแปลงข้อมูลกลับมาเป็นโครงสร้างใน Go
โดยใช้ json.Unmarshal() ซึ่งช่วยให้การ
จัดการข้อมูลที่มาจากแหล่งภายนอกทำได้สะดวกมากขึ้น
*/
/* สรุป
Marshaling คือการแปลงข้อมูลจาก Go ไปเป็น JSON
Unmarshaling คือการแปลงข้อมูลจาก JSON ไปเป็นโครงสร้างข้อมูล
โค้ดนี้แสดงการทำงานของทั้งสองกระบวนการเพื่อจััดการกับข้อมูล
ในรูปแบบต่างๆ เช่น สร้าง JSON, อ่าน JSON, แปลงข้อมูล
และนำข้อมูลมาใช้งานในโครงสร้าง Go
*/
