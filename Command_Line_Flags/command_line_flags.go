package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/* การใช้งาน command-line flags ในภาษา Go
โดยใช้แพ็กเกจ flag เพื่อจัดการและประมวลผลอาร์กิวเมนต์ของโปรแกรม
ในรูปแบบของ flags (ธงหรือสวิตซ์) เช่น -word หรือ -numb เพื่อ
กำหนดค่าให้กับตัวเลือกต่างๆ ของโปรแกรม

1.การประกาศ Flags
wordPtr := flag.String("word", "foo", "a string")
numbPtr := flag.Int("numb", 42, "an int")
forkPtr := flag.Bool("fork", false, "a bool")
เราสามารถสร้าง flag โดยใช้ฟังก์ชันต่างๆ จากแพ็กเกจ flag
ซึ่งในที่นี้จะมีสาม flags:
-word ใช้เพื่อรับค่าเป็น string โดยมีค่าเริ่มต้นเป็น "foo"
-numb ใช้เพื่อรับค่าเป็น int โดยมีค่าเริ่มต้นเป็น 42
-fork ใช้เพื่อรับค่าเป็น boolean (true หรือ false)
โดยมีค่าเริ่มต้นเป็น false
ฟังก์ชันเหล่านี้จะคืนค่ากลับมาเป็น pointer ไปยังชนิดข้อมูลที่กำหนด
(string, int, bool)

2.การประกาศ Flags ที่ใช้ตัวแปรที่มีอยู่แล้ว
var svar string
flag.StringVar(&svar, "svar", "bar", "a string var")
นี่เป็นการประกาศ flag ที่เชื่อมโยงกับตัวแปรที่มีอยู่แล้ว (ในที่นี้คือ svar)
โดยใช้ flag.StringVar ซึ่งต้องส่ง pointer ของตัวแปร (&svar)
เข้าไป
flag นี้ใช้สำหรับการรับค่า string และมีค่าเริ่มต้นเป็น bar

3.การประมวลผล Flags
flag.Parse()
flag.Parse() จะทำการอ่านและประมวลผล flags ทั้งหมดที่
ถูกส่งเข้ามาผ่านทาง command-line
flags ที่ไม่ได้ถูกประกาศในโค้ดจะทำให้โปรแกรมแสดงข้อผิดพลาด
และแสดงข้อความช่วยเหลืออัตโนมัติ

4.การใช้งานและการแสดงผลค่าของ Flags
fmt.Println("word:", *wordPtr)
fmt.Println("numb:", *numbPtr)
fmt.Println("fork:", *forkPtr)
fmt.Println("svar:", svar)
fmt.Println("tail:", flag.Args())
เมื่อ flags ถูกประมวลผลแล้ว เราต้องใช้การ dereference pointer
เพื่อเข้าถึงค่าที่แท้จริงของ flags เช่น *wordPtr, *numbPtr, และ *forkPtr
ตัวแปร svar ไม่จำเป็นต้อง dereference เนื่องจากมันถูกผูกตรงกับตัวแปร
ที่ประกาศไว้แล้ว
flag.Args() จะคืนค่าเป็น slice ของอาร์กิวเมนต์ที่เหลือ (หรือที่เรียกว่า
positional arguments) ที่ไม่ได้เป็นส่วนหนึ่งของ flags
*/
/* ผลลัพธ์จะเป็น:
word: hello
numb: 7
fork: true
svar: test
tail: [arg1 arg2]
-word=hello กำหนดให้ flag word มีค่า "hello"
-numb=7 กำหนดให้ flag numb มีค่า 7
-fork เป็น boolean flag ดังนั้นแค่ระบุ flag นี้ก็จะทำให้ค่าเป็น true
-svar=test กำหนดให้ตัวแปร svar มีค่า "test"
arg1 และ arg2 เป็น positional arguments ที่เหลือจากการประมวลผล flags
*/
/* สรุป
โค้ดนี้ใช้แพ็กเกจ flag เพื่อจัดการกับ command-line flags ที่
ประกอบด้วย string, int, และ boolean flags ซึ่งสามารถประมวลผล
flags ที่ถูกส่งเข้ามา และใช้ flags เหล่านั้นในโปรแกรม
*/
