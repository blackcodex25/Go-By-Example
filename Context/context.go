package main

import (
	"fmt"
	"net/http"
	"time"
)

/* โค้ดนี้แสดงถึงการใช้ context เพื่อจัดการการยกเลิกคำขอใน HTTP server อย่างมีประสิทธิภาพ
โดยช่วยให้สามารถตอบสนองต่อการเปลี่ยนแปลงของสถานะคำขอได้อย่างเหมาะสม */

// ฟังก์ชัน handler สำหรับการตอบกลับเมื่อมีคำขอที่เส้นทาง /hello
func hello(w http.ResponseWriter, req *http.Request) {
	// รับ context จากคำขอ
	ctx := req.Context()
	fmt.Println("server: hello handler started")     // แจ้งว่า handler เริ่มทำงาน
	defer fmt.Println("server: hello handler ended") // แจ้งเมื่อ handler สิ้นสุดการทำงาน

	// เลือกว่าจะรอเวลาหรือยกเลิกคำขอ
	select {
	// รอ 10 วินาทีเพื่อส่งข้อความ "hello"
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n") // ส่งคำตอบ
		// หากมีสัญญาณจาก context ว่าคำขอถูกยกเลิก
	case <-ctx.Done():

		err := ctx.Err()                                // รับข้อผิดพลาด
		fmt.Println("server:", err)                     // แจ้งข้อผิดพลาดในเซิรฟ์เวอร์
		internalError := http.StatusInternalServerError // รหัสข้อผิดพลาด
		http.Error(w, err.Error(), internalError)       // ส่งข้อผิดพลาดไปยัง client
	}
}

func main() {
	http.HandleFunc("/hello", hello)  // ลงทะเบียน handler
	http.ListenAndServe(":8090", nil) // เริ่มเซิรฟเวอร์บนพอร์ต 8090
}

/*
อธิบายรายละเอียด
โค้ดนี้เป็นตัวอย่างการใช้งาน context.Context ใน HTTP server
เพื่อควบคุมการยกเลิกคำขอ (cancellation)
โดยมีการกำหนดเวลาในการตอบกลับให้ใช้เวลา 10 วินาทีเพื่อจำลองงานที่เซิร์ฟเวอร์กำลังทำ
หากคำขอถูกยกเลิกก่อนที่จะมีการตอบกลับ จะมีการส่งข้อความแสดงข้อผิดพลาดไปยัง client

หลักการทำงาน
ฟังก์ชัน hello:
ใช้ req.Context() เพื่อรับ context ที่เกี่ยวข้องกับคำขอปัจจุบัน
สร้าง select statement เพื่อตรวจสอบว่าจะส่งคำตอบเมื่อมีเวลาผ่านไป 10 วินาทีหรือ
เมื่อได้รับสัญญาณจาก context ว่าคำขอถูกยกเลิก
หาก ctx.Done() ถูกส่งสัญญาณ ฟังก์ชันจะพิมพ์ข้อความข้อผิดพลาดและ
ส่งกลับข้อความข้อผิดพลาด HTTP 500

ฟังก์ชัน main:
ใช้ http.HandleFunc เพื่อลงทะเบียนฟังก์ชัน handler สำหรับเส้นทาง /hello
เรียก http.ListenAndServe เพื่อเริ่มเซิร์ฟเวอร์บนพอร์ต 8090 */
