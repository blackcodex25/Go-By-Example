package main

import "fmt"

/*ในตัวอย่างก่อนหน้านี้เราเห็นว่า for และ range ใช้
ในการวนลูปผ่านโครงสร้างข้อมูลพื้นฐานต่างๆ ได้อย่างไร
เราสามารถใช้ไวยากรณ์นี้เพื่อวนลูปผ่านค่าที่ได้รับจาก
channel ได้เช่นกัน

เราจะวนลูปผ่าน 2 ค่าใน channel ชื่อ queue โดยใช้
range ซึ่งจะวนลูปผ่านแต่ละองค์ประกอบเมื่อมันถูกดึง
ออกมาจาก queue เนื่องจากเราได้ Closing channel
ไปแล้วการวนลูปจะสิ้นสุดหลังจากที่ได้รับค่าทั้ง 2 ตัว

ตัวอย่างนี้ยังแสดงให้เห็นว่าเป็นไปได้ที่จะ Closing channel
ที่ยังมีค่าคงเหลืออยู่ แต่ยังสามารถรับค่าที่เหลืออยู่ใน
channel ได้
*/
func main() {
	// สร้าง channel พร้อม buffer ขนาด 2
	queue := make(chan string, 2)

	// ส่งค่าลงใน channel
	queue <- "one"
	queue <- "two"

	// close channels หลังจากส่งค่าครบแล้ว
	close(queue)

	// ใช้ range ในการวนลูปผ่านค่าที่ได้รับจาก channels
	for elem := range queue {
		fmt.Println(elem)
	}
}

/* อธิบายการทำงาน
การสร้าง Channel:
queue := make(chan string, 2)
สร้าง channel ชื่อ queue ที่สามารถเก็บ
ค่าประเภท string ได้สูงสุด 2 ค่า
(buffered channel)

การส่งค่าไปยัง Channel:
queue <- "one"
ส่งค่าข้อความ "one" ไปยัง channel queue
queue <- "two"
ส่งค่าข้อความ "two" ไปยัง channel queue

closing channel
close(queue)
ปิด channel queue เพื่อบอกว่าไม่มีค่าที่
จะส่งไปยัง channel นี้อีกต่อไป

การวนลูปผ่าน Channel:
for elem := range queue
ใช้ for และ range ในการวนลูปผ่านค่าที่ได้รับจาก channel queue
เนื่องจาก channel ถูกปิดแล้ว การวนลูปจะ
ทำงานจนกว่าจะได้รับค่าทั้งหมดใน
channel
เมื่อ channels ถูกปิดและค่าทั้งหมดถูกดึง
ออกมาแล้ว range จะหยุดการวนลูป

การพิมพ์ค่าที่รับจาก Channel:
fmt.Println(elem)
พิมพ์ค่าที่ได้รับจาก channel ออกมา

โดยแต่ละค่าที่ได้รับจาก channel queue จะถูกพิมพ์
ออกมาตามลำดับที่มันถูกส่งเข้าไปใน channel ก่อนหน้านี้
*/

/* สรุป
โค้ดนี้แสดงถึงการใช้งาน channel ใน Go สำหรับการส่ง
ค่าและการวนลูปผ่านค่าที่รับจาก channel ที่ถูกปิดแล้ว
โดยแสดงให้เห็นถึงการใช้ for และ range เพื่อเข้า
ถึงค่าทั้งหมดที่มีใน channel จนกว่าจะไม่มีค่าคงเหลือ
และ channel ถูกปิดแล้ว
*/
