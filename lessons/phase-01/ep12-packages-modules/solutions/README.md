# เฉลย EP.12

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่มฟังก์ชัน `ToFahrenheit(celsius float64) float64` ใน package sensor แล้วให้ main พิมพ์ค่า 30 C เป็น Fahrenheit ต่อท้าย

ดู [main.go](01/main.go) และ [sensor/reading.go](01/sensor/reading.go) ทั้ง main กับ sensor ในข้อนี้เป็นตัวอย่างแยกกัน เมื่อนำไปฝึกใช้ import ตาม go.mod ของ practics

รันจาก `lessons/phase-01/ep12-packages-modules/solutions/01` ด้วย `go run .`

```text
sensor-01: WARNING
sensor-02: OK
30 C = 86.0 F
```


เหตุผล: ชื่อ ToFahrenheit ขึ้นต้นตัวใหญ่จึงเรียกผ่าน sensor จาก main ได้

## ข้อ 2

ให้ main ใช้เกณฑ์ 35 ทั้งสองรายการ โดยไม่แก้ package sensor แล้วคาดเดาสถานะใหม่

ดู [main.go](02/main.go) และ [sensor/reading.go](02/sensor/reading.go) ทั้ง main กับ sensor ในข้อนี้เป็นตัวอย่างแยกกัน เมื่อนำไปฝึกใช้ import ตาม go.mod ของ practics

รันจาก `lessons/phase-01/ep12-packages-modules/solutions/02` ด้วย `go run .`

```text
sensor-01: OK
sensor-02: OK
```


เหตุผล: เกณฑ์เป็น parameter ของ Status ผู้เรียกจึงเปลี่ยนกฎรอบนี้ได้โดยไม่แก้ตัวฟังก์ชัน


[กลับบทเรียน](../README.md)
