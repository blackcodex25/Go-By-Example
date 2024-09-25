package main

import (
	"testing"
)

// Benchmark for inserting nodes
func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root := &TreeNode{Value: 0}
		for j := 1; j <= 10000; j++ { // Insert 1000 nodes
			insert(root, j)
		}
	}
}

// Benchmark for searching nodes
func BenchmarkSearch(b *testing.B) {
	root := &TreeNode{Value: 5000} // สร้างโหนดรากของต้นไม้ชื่อ root เริ่มต้นที่โหนดราก ที่มีค่า 5000
	for j := 1; j <= 100000; j++ { // ใช้ for เพื่อแทรกโหนดใหม่ทั้งหมด 100,000 โหนด
		insert(root, j) // โดยมีฟังก์ชัน insert เพื่อเพิ่มโหนดใหม่
	}

	b.ResetTimer()             // รีเซ็ตเวลา เพื่อวัดเฉพาะเวลาค้นหา
	for i := 0; i < b.N; i++ { // วนแต่ละรอบของ Benchmark (b.N รอบ)
		search(root, 5000) // จะเรียกฟังก์ชัน search เพื่อค้นหาค่า 5000
	}
	//  ซึ่ง b.N คือจำนวนรอบที่ testing.B กำหนดไว้สำหรับการทดสอบ
	//  โดยที่ค่าของ b.N จะถูกปรับโดยอัตโนมัติตามความเร็วของการดำเนินการเพื่อให้ได้ผลลัพธ์ที่แม่นยำ
}

func TestBenchmarkInsert(t *testing.T) {
	type args struct {
		b *testing.B
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BenchmarkInsert(tt.args.b)
		})
	}
}

func TestBenchmarkSearch(t *testing.T) {
	type args struct {
		b *testing.B
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BenchmarkSearch(tt.args.b)
		})
	}
}
