package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ลดตัวนับเมื่อ goroutine นี้เสร็จสิ้น
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // จำลองงานที่ใช้เวลา
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // เพิ่มตัวนับสำหรับแต่ละ goroutine
		go worker(i, &wg)
	}

	wg.Wait() // รอให้ goroutine ทั้งหมดเสร็จสิ้น
	fmt.Println("All workers done")
	/*ในตัวอย่างนี้, เราใช้ WaitGroup เพื่อรอให้ goroutine
	ทั้งหมดเสร็จสิ้นก่อนที่จะพิมพ์ข้อความว่า "All workers done".
	*/
}
