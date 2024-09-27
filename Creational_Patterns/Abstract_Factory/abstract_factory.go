package main

import "fmt"

/* Abstract Factory Pattern */
/* ให้ส่วนติดต่อสำหรับการสร้างครอบครัวของอ็อบเจ็กต์ที่เกี่ยวข้อง */
/* โดยไม่ต้องระบุคลาสที่เป็นรูปธรรม */

// การประกาศ Interface
type ( // จัดกลุ่ม interface
	// interface Color
	Color interface {
		Fill() // มีฟังก์ชัน Fill() สำหรับการเติมสี
	}
	// interface Shape
	Shape interface {
		Draw() // มีฟังก์ชัน Draw() สำหรับการวาดรูป
	}
	// interface AbstractFactory
	AbstractFactory interface {
		// Method GetShape()
		GetShape(shapeType string) Shape // สร้างอ็อบเจ็กต์ของรูปทรงตามประเภทที่ระบุ
		// Method GetColor()
		GetColor(colorType string) Color // สร้างอ็อบเจ็กต์ของสีตามประเภทที่ระบุ
	}
)

// การประกาศ Struct
type ( // จัดกลุ่ม struct
	// Struct Rectangle และ Circle มีคุณสมบัติ
	// แสดงถึงรูปทรง และมีฟิลด์คือ Color ที่เป็น interface ของ Color
	Rectangle struct {
		Color Color
	}
	Circle struct {
		Color Color
	}
	// Struct Red และ Blue แสดงถึงสีต่างๆ
	Red  struct{}
	Blue struct{}
	// ใช้สร้างอ็อบเจ็กต์ของ Shape และ Color
	ShapeFactory struct{}
	ColorFactory struct{}
)

// การกำหนด Method
// Method Fill() ของ struct Red จะพิมพ์ข้อความว่า "Filling with red color"
func (r *Red) Fill() {
	fmt.Println("Filling with red color")
}

// Method Fill() ของ struct Blue จะพิมพ์ข้อความว่า "Filling with blue color"
func (b *Blue) Fill() {
	fmt.Println("Filling with blue color")
}

// Rectangle จะเรียกใช้ Color.Fill() และพิมพ์ "Drawing a Rectangle"
func (r *Rectangle) Draw() {
	r.Color.Fill()
	fmt.Println("Drawing a Rectangle")
}

// Circle จะเรียกใช้ Color.Fill() และพิมพ์ "Drawing a Circle"
func (c *Circle) Draw() {
	c.Color.Fill()
	fmt.Println("Drawing a Circle")
}

// Shape Factory
var shape = map[string]func() Shape{
	"rectangle": func() Shape { return &Rectangle{} },
	"circle":    func() Shape { return &Circle{} },
}

// Color Factory
var color = map[string]func() Color{
	"red":  func() Color { return &Red{} },
	"blue": func() Color { return &Blue{} },
}

// Abstract Factory Interface
// สร้าง Rectangle หรือ Circle ตามประเภทที่กำหนด
func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	if shapeFunc, ok := shape[shapeType]; ok {
		return shapeFunc()
	}
	return nil
}

// สร้าง Red หรือ Blue ตามประเภทที่กำหนด
func (cf *ColorFactory) GetColor(colorType string) Color {
	if colorFunc, ok := color[colorType]; ok {
		return colorFunc()
	}
	return nil
}

// ตัวแปรที่ใช้ใน main
var (
	// สร้างอ็อบเจ็กต์ rectangle และ circle จาก ShapeFactory
	shapeFactory = &ShapeFactory{}
	// สร้างอ็อบเจ็กต์ red และ blue จาก ColorFactory
	colorFactory = &ColorFactory{}

	// เรียกใช้ GetColor จาก colorFactory โดยส่งพารามิเตอร์เป็น "red"
	// เรียกใช้ GetColor จาก colorFactory โดยส่งพารามิเตอร์เป็น "blue"
	red  = colorFactory.GetColor("red")
	blue = colorFactory.GetColor("blue")

	// เรียกใช้ GetShape จาก shapeFactory โดยส่งพารามิเตอร์เป็น "rectangle"
	// เรียกใช้ GetShape จาก shapeFactory โดยส่งพารามิเตอร์เป็น "circle"
	rectangle = shapeFactory.GetShape("rectangle").(*Rectangle)
	circle    = shapeFactory.GetShape("circle").(*Circle)
	/* เราทำการแปลงผลลัพธ์ทั้งสองเป็น
	*Rectangle (Pointer to Rectangle), *Circle (Pointer to Circle)
	เพื่อให้สามารถเข้าถึงฟิลด์ของมันได้ */
)

func main() {
	// กำหนดสีของ rectangle เป็น red และเรียก Method Draw()
	rectangle.Color = red
	rectangle.Draw()

	// กำหนดสีของ circle เป็น blue และเรียก Method Draw()
	circle.Color = blue
	circle.Draw()
}

/* โค้ดนี้แสดงถึงการสร้างอ็อบเจ็กต์ในครอบครัวของอ็อบเจ็กต์ที่เกี่ยวข้อง (Shapes และ Colors) */
/* โดยไม่ต้องระบุคลาสที่เป็นรูปธรรม ซึ่งทำให้การ สร้างอ็อบเจ็กต์มีความยืดหยุ่นและ
ง่ายต่อการปรับแต่งในอนาคต */
