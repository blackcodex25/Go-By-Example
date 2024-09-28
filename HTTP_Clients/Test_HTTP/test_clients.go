package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const url = "https://golang.withcodeexample.com/blog/golang-clean-code-guide/"
const pathFile = "./HTTP_Client/response.txt"

func main() {
	// ส่งคำขอ HTTP GET ไปยัง URL ที่กำหนด
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error Response status: %v", err) // แสดงข้อผิดพลาดถ้ามี
	}
	defer resp.Body.Close()

	// พิมพ์สถานะการตอบสนอง HTTP
	fmt.Println("Response status:", resp.Status)

	// อ่านเนื้อหาคำตอบแบบเต็ม
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf(" Error reading response: %v", err) // แสดงข้อผิดพลาดถ้ามี
	}

	// พิมพ์เนื้อหาการตอบสนองแบบเต็มเป็นสตริง
	// fmt.Println(string(body))

	// บันทึกเนื้อหาการตอบกลับเป็นไฟล์ .txt
	err = os.WriteFile(pathFile, body, 0644)
	if err != nil {
		log.Printf("Error Saved to response: %v", err) // แสดงข้อผิดพลาดถ้ามี
		return
	}

	// พิมพ์ข้อความ success message
	fmt.Println("Response saved to response.txt")
}
