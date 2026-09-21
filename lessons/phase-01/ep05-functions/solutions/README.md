# เฉลย EP.5

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่มฟังก์ชัน `toCelsius(fahrenheit float64) float64` ใช้สูตร `(fahrenheit - 32) * 5 / 9` แล้วพิมพ์ผลของ 86 F ต่อท้าย

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep05-functions/solutions/01` ด้วย `go run .`

```text
30.0 C = 86.0 F [WARNING]
at threshold 35: OK
86 F = 30.0 C
```


เหตุผล: ฟังก์ชันรับค่าเป็น float64 จึงคำนวณทศนิยมและคืนค่าให้ main เลือกแสดงผล

## ข้อ 2

ใช้ for เรียก status กับค่า 29, 30 และ 31 โดยใช้เกณฑ์ 30 แล้วพิมพ์ผลแต่ละค่า แทน main เดิม

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep05-functions/solutions/02` ด้วย `go run .`

```text
29 C: OK
30 C: WARNING
31 C: WARNING
```


เหตุผล: ส่งค่าที่เปลี่ยนแต่ละรอบเข้า status โดยใช้กฎเกณฑ์เดียวกัน ไม่ต้องเขียนเงื่อนไขซ้ำใน main


[กลับบทเรียน](../README.md)
