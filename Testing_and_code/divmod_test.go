package testingandcode

import "testing"

func TestDivMod(t *testing.T) {
	dvnd := 40                              // ตัวตั้ง (dividend)
	for dvsor := 1; dvsor < dvnd; dvsor++ { // ตัวหาร (divisor) เริ่มจาก 1 ถึง 39
		q, r := DivMod(dvnd, dvsor) // เรียกใช้ฟังก์ชัน Divmod
		if (dvnd % dvsor) != r {    // ตรวจสอบว่าเศษที่เหลือถูกต้องหรือไม่
			// ถ้าไม่ถูกต้อง ให้แสดงความผิดพลาด
			t.Fatalf("%d%d q=%d, r=%d, bad remainder.", dvnd, dvsor, q, r)
		}
	}
}
