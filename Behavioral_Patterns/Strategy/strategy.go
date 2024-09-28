package main

import "fmt"

/* Strategy Patterns */
/* Concept: ช่วยให้เราสามารถเปล่ยนพฤติกรรมของวัตถุ (Object) */
/* ได้โดยการใช้กลยุทธ์ (strategy) ที่ต่างกัน โดยไม่ต้องเปลี่ยนโค้ดของวัตถุหลัก */

/* ขั้นตอนที่ 1: สร้าง Strategy Interface */
type Strategy interface {
	Execute(a, b int) int // method Execute มีพารามิเตอร์ 2 ตัว a, b ตามลำดับ และคืนค่า int
	// พฤติกรรมโดยละเอียดของ Method Execute ถูกปล่อยให้เป็นไปตามประเภทที่นำไปใช้
}

/* ขั้นตอนที่ 2: สร้าง Strategies */
// ประกาศโครงสร้าง Empty strcut Addstrategy เพื่อนำไปใช้งานร่วมกับ Method Execute ของ Inteface Strategy
type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a - b
}

// ขั้นตอนที่ 3: สร้าง Context
// ประกาศ Struct Context ที่เก็บข้อมูล Strategy เท่านั้น ไม่มี method ใดๆ ที่กำหนดไว้
type Context struct {
	strategy Strategy
}

//
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &Context{}

	add := &AddStrategy{}
	context.SetStrategy(add)
	fmt.Println("Addition:", context.ExecuteStrategy(3, 4)) // Addition: 7

	subtract := &SubtractStrategy{}
	context.SetStrategy(subtract)
	fmt.Println("Subtraction:", context.ExecuteStrategy(10, 4)) // Subtraction: 6

}
