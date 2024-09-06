package testingandcode

func DivMod(dvdn, dvsr int) (q, r int) {
	r = dvdn        // เริ่มต้นโดยกำหนดให้ r เท่ากับตัวตั้ง (dividend)
	for r >= dvsr { // วนลูปตราบใดที่ r ยังคงมากกว่าหรือเท่ากับตัวหาร (divisor)
		q += 1       // เพิ่มค่า q ซึ่งเป็นตัวเก็บผลหารทุกครั้งที่ r ลดลง
		r = r - dvsr // ลบค่าตัวหาร dvsr ออกจาก r จนกว่า r จะน้อยกว่า dvsr
	}
	return // คืนค่า q (quotient) และ r (remainder) ออกไป
}
