package main

import (
	"fmt"
	"log"
	"net/http"
)

// ฟังก์ชัน hello ส่งข้อความ "hello\n" เป็นคำตอบ
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request method:", req.Method)
	fmt.Fprintf(w, "hello\n")
}

// ฟังก์ชัน headers อ่าน HTTP headers จากคำขอแล้วส่งกลับเป็นคำตอบ
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				log.Printf("Error writing headers: %v", err)
				http.Error(w, "Error writing headers", http.StatusInternalServerError)
				return
			}
		}
	}
}
func main() {
	// ใช้ http.HandleFunc เพื่อจับคู่เส้นทาง /hello และ /headers
	// กับฟังก์ชัน handler
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	// ใช้ http.ListenAndServe เพื่อเริ่มเซิร์ฟเวอร์บนพอร์ต 8090
	http.ListenAndServe(":8090", nil)
}

/* บทความนี้เกี่ยวกับการเขียน HTTP Server โดยใช้แพ็กเกจ net/http
ใน Go โดยเน้นการใช้ handlers ซึ่งเป็นฟังก์ชันที่จัดการคำขอ HTTP
โดยรับพารามิเตอร์ http.ResponseWriter สำหรับการเขียนการตอบกลับ
และ http.Request สำหรับอ่านคำขอ HTTP โค้ดนี้มีสอง handler คือ hello
ที่แสดงคำว่า "hello" และ headers ที่แสดง HTTP headers จากคำขอผู้เขียนใช้
http.HandleFunc เพื่อลงทะเบียนเส้นทาง /hello และ /headers
แล้วเรียก http.ListenAndServe(":8090", nil)
เพื่อเริ่มเซิร์ฟเวอร์บนพอร์ต 8090 */
