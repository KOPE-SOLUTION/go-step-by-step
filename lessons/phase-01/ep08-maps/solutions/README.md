# เฉลย EP.8

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

ก่อนค้น sensor-04 ให้เพิ่มชื่อนี้ด้วยค่า 25 แล้วดูว่าจำนวนอุปกรณ์หลัง delete เป็นเท่าไร

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep08-maps/solutions/01` ด้วย `go run .`

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: 25.0 C
devices: 3
```


เหตุผล: เพิ่ม key ใหม่ก่อนค้นทำให้ ok เป็น true หลังลบ sensor-03 เหลือ 3 key

## ข้อ 2

หลัง delete sensor-03 ให้ลองค้นชื่อนี้อีกครั้ง ต้องรายงาน not found

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep08-maps/solutions/02` ด้วย `go run .`

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: not found
sensor-03: not found
devices: 2
```


เหตุผล: delete เอา key ออก การอ่านครั้งถัดไปจึงได้ ok=false


[กลับบทเรียน](../README.md)
