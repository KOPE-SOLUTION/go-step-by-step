# เฉลย EP.12

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่มฟังก์ชัน `ToFahrenheit(celsius float64) float64` ใน `sensor/reading.go` แล้วให้ `main` เรียกฟังก์ชันเพื่อแปลง 30 C เป็น Fahrenheit และพิมพ์ต่อท้ายรายงาน

ดู [main.go](01/main.go) และ [sensor/reading.go](01/sensor/reading.go) ไฟล์ทั้งสองทำงานร่วมกัน เพราะ `main.go` เรียกฟังก์ชันจาก package `sensor` เมื่อนำไปฝึก ให้แก้ import เป็นชื่อ module ใน `practics/go.mod` ตามด้วย `/sensor`

รันจาก `lessons/phase-01/ep12-packages-modules/solutions/01` ด้วย `go run .`

```text
sensor-01: WARNING
sensor-02: OK
30 C = 86.0 F
```

เหตุผล: ชื่อ `ToFahrenheit` ขึ้นต้นด้วยตัวใหญ่ จึงเรียก `sensor.ToFahrenheit` จาก package `main` ได้

## ข้อ 2

ใน `main` เปลี่ยนเกณฑ์ที่ส่งให้ `sensor.Status` เป็น 35 ทั้งสองรายการ โดยไม่แก้ package `sensor` แล้วคาดเดาสถานะใหม่

ดู [main.go](02/main.go) และ [sensor/reading.go](02/sensor/reading.go) ไฟล์ทั้งสองทำงานร่วมกัน เพราะ `main.go` เรียกฟังก์ชันจาก package `sensor` เมื่อนำไปฝึก ให้แก้ import เป็นชื่อ module ใน `practics/go.mod` ตามด้วย `/sensor`

รันจาก `lessons/phase-01/ep12-packages-modules/solutions/02` ด้วย `go run .`

```text
sensor-01: OK
sensor-02: OK
```

เหตุผล: `Status` รับเกณฑ์ผ่าน parameter ผู้เรียกจึงเปลี่ยนเกณฑ์ได้โดยไม่แก้โค้ดภายในฟังก์ชัน

[กลับบทเรียน](../README.md)
