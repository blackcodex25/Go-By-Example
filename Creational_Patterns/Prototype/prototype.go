package main

import "fmt"

/* Prototype Pattern */
/* สร้างอ็อบเจ็กต์ใหม่โดยการคัดลอกอ็อบเจ็กต์ที่มีอยู่ */

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	ID int
}

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{ID: p.ID}
}

func main() {
	prototype := &ConcretePrototype{ID: 1}
	clone := prototype.Clone().(*ConcretePrototype)

	fmt.Println(prototype.ID) // 1
	fmt.Println(clone.ID)     // 1

	clone.ID = 2
	fmt.Println(prototype.ID) // 1
	fmt.Println(clone.ID)     // 2
}

/* สรุป */
// รูปแบบการสร้างเหล่านี้ช่วยให้การสร้างอ็อบเจ็กต์ใน Go มรประสิทธิภาพ
// และยืดหยุ่นมากขึ้น เราสามารถเลือกใช้รูปแบบที่เหมาะสมกับความ
// ต้องการของโปรเจกต์ของเรา
