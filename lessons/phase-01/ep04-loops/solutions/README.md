# เฉลย EP.4

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

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

[กลับบทเรียน](../README.md)
