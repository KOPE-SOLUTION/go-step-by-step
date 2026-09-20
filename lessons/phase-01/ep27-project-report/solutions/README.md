# เฉลย EP.27

หยุดก่อนถ้ายังไม่ได้ลอง [แบบฝึกหัด](../exercises/README.md) คำตอบไม่จำเป็นต้องเขียนเหมือนกันทุกตัวอักษร แต่ต้องตรงข้อกำหนด

คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep27-project-report` ไม่ใช่โฟลเดอร์ practics ซึ่งไม่มีเฉลย

## ข้อ 1

รายการใหม่ต่อท้าย slice จึงอยู่ท้ายรายงาน

ดู [main.go](01/main.go) และ [model.go](01/model.go)

```powershell
go run ./solutions/01
```

ผล:

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
sensor-05: 31.0 C [WARNING]
```

## ข้อ 2

slice ว่างทำให้ลูปไม่เข้ารอบ

ดู [main.go](02/main.go) และ [model.go](02/model.go)

```powershell
go run ./solutions/02
```

ผล: ไม่มีข้อความ โปรแกรมจบปกติ

## คำถามทบทวน

continue ต่างจาก return ใน buildReport อย่างไร?

continue ไปยังรอบถัดไปใน for ส่วน return ออกจาก buildReport ทั้งฟังก์ชัน

## กลับไปลองอีกครั้ง

ปิดเฉลย เปลี่ยนค่าหนึ่งจุดในพื้นที่ฝึก แล้วอธิบายผลด้วยตนเอง [กลับบทเรียน](../README.md)
