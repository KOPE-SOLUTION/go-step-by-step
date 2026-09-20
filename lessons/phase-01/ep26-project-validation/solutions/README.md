# เฉลย EP.26

หยุดก่อนถ้ายังไม่ได้ลอง [แบบฝึกหัด](../exercises/README.md) คำตอบไม่จำเป็นต้องเขียนเหมือนกันทุกตัวอักษร แต่ต้องตรงข้อกำหนด

คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep26-project-validation` ไม่ใช่โฟลเดอร์ practics ซึ่งไม่มีเฉลย

## ข้อ 1

ตรวจชื่อก่อนค่า

ดู [main.go](01/main.go) และ [model.go](01/model.go)

```powershell
go run ./solutions/01
```

ผล:

```text
ERROR: device ID is empty
```

## ข้อ 2

100 ยังอยู่ในช่วงที่ยอมรับ

ดู [main.go](02/main.go) และ [model.go](02/model.go)

```powershell
go run ./solutions/02
```

ผล:

```text
sensor-01: 100.0 C [WARNING]
```

## คำถามทบทวน

ถ้าชื่อว่างและค่า -1 โปรแกรมแจ้ง error ใด เพราะอะไร?

device ID is empty เพราะตรวจและ return จากเงื่อนไขชื่อก่อน จึงไม่ถึงการตรวจช่วง

## กลับไปลองอีกครั้ง

ปิดเฉลย เปลี่ยนค่าหนึ่งจุดในพื้นที่ฝึก แล้วอธิบายผลด้วยตนเอง [กลับบทเรียน](../README.md)
