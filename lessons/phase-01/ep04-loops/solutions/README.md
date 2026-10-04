# เฉลย EP.4

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน ข้อ 1–2 เริ่มจากตัวอย่างเต็ม ส่วนข้อ 3 เริ่มจากลูปแบบมีเฉพาะเงื่อนไขในขั้นที่ 2

## ข้อ 1

เปลี่ยนเงื่อนไขลูปจาก `round <= 5` เป็น `round <= 3` โดยยังข้ามรอบ 3 ค่าเฉลี่ยต้องคำนวณจากสองรอบแรก

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep04-loops/solutions/01` ด้วย `go run .`

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: skipped
average: 26.50 C (2 readings)
```

เหตุผล: รอบ 3 ถูกข้าม จึงมีผลรวม 53 และ count=2 ค่าเฉลี่ยเท่ากับ 26.5

## ข้อ 2

ให้หยุดลูปเมื่อถึงรอบ 3 โดยเปลี่ยน `continue` เป็น `break` และเปลี่ยนข้อความเป็น `round 3: stopped`

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep04-loops/solutions/02` ด้วย `go run .`

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: stopped
average: 26.50 C (2 readings)
```

เหตุผล: `break` ออกจากลูปทันที รอบ 4–5 จึงไม่ทำงาน

## ข้อ 3

เริ่มจากโค้ดขั้นที่ 2 ใช้ `celsius -= 2` แทน `celsius -= 1`

```go
package main

import "fmt"

func main() {
	celsius := 33.0
	for celsius > 30 {
		fmt.Printf("cooling: %.1f C\n", celsius)
		celsius -= 2
	}
	fmt.Printf("ready: %.1f C\n", celsius)
}
```

รันจาก `practics` ด้วย `go run main.go`:

```text
cooling: 33.0 C
cooling: 31.0 C
ready: 29.0 C
```

เหตุผล: ค่าเปลี่ยนจาก 33 เป็น 31 แล้วเป็น 29 เมื่อ 29 ไม่มากกว่า 30 ลูปจึงหยุด เงื่อนไขกำหนดว่าจะทำต่อหรือไม่ ไม่ได้บังคับให้ค่าสุดท้ายต้องเท่ากับ 30

[กลับบทเรียน](../README.md)
