# เฉลย EP.9

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่ม `Reading` ของ `sensor-03` อุณหภูมิ 31.5 ลงใน slice `readings` เพื่อให้รายงานแสดงสามอุปกรณ์

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep09-structs/solutions/01` ด้วย `go run .`

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: 31.5 C [WARNING]
original: 27.5, copy: 99.0
```

เหตุผล: สมาชิกแต่ละตัวเก็บชื่อและอุณหภูมิไว้ใน `Reading` ลูปเดิมจึงอ่านอุปกรณ์ที่เพิ่มมาได้โดยไม่ต้องแก้ลูป

## ข้อ 2

ก่อนลูป `range` ให้เปลี่ยน `Celsius` ของสมาชิกแรกเป็น 32 ผ่าน `readings[0]` แล้วตรวจอุณหภูมิในรายงานและค่า `original` ที่พิมพ์ตอนท้าย

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep09-structs/solutions/02` ด้วย `go run .`

```text
sensor-01: 32.0 C [WARNING]
sensor-02: 30.0 C [WARNING]
original: 32.0, copy: 99.0
```

เหตุผล: การแก้ `readings[0].Celsius` เปลี่ยนสมาชิกใน slice โดยตรง เมื่อคัดลอกสมาชิกนั้นไปเป็น `original` ภายหลัง จึงได้ค่า 32 ด้วย

[กลับบทเรียน](../README.md)
