# เฉลย EP.9

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่ม sensor-03 ค่า 31.5 ใน slice เพื่อให้รายงานแสดงสามอุปกรณ์

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep09-structs/solutions/01` ด้วย `go run .`

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: 31.5 C [WARNING]
original: 27.5, copy: 99.0
```


เหตุผล: สมาชิกแต่ละตัวรวมชื่อและค่าไว้ใน Reading เดียว ลูปเดิมจึงใช้กับจำนวนที่เพิ่มได้

## ข้อ 2

ก่อน range เปลี่ยน Celsius ของสมาชิกแรกเป็น 32 ผ่าน `readings[0]` แล้วตรวจทั้งรายงานและค่าต้นฉบับตอนท้าย

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep09-structs/solutions/02` ด้วย `go run .`

```text
sensor-01: 32.0 C [WARNING]
sensor-02: 30.0 C [WARNING]
original: 32.0, copy: 99.0
```


เหตุผล: แก้ผ่าน index จึงเปลี่ยนสมาชิกจริง ก่อนคัดลอกออกมาเป็น original


[กลับบทเรียน](../README.md)
