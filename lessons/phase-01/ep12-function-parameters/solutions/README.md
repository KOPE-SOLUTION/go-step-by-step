# เฉลย EP.12

หยุดก่อนถ้ายังไม่ได้ลอง [แบบฝึกหัด](../exercises/README.md) คำตอบไม่จำเป็นต้องเขียนเหมือนกันทุกตัวอักษร แต่ต้องตรงข้อกำหนด

คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep12-function-parameters` ไม่ใช่โฟลเดอร์ practics ซึ่งไม่มีเฉลย

## ข้อ 1

เปลี่ยน argument ที่จุดเรียก

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
Device: gateway-01
```

## ข้อ 2

parameter กำหนดชนิด int

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
Count: 3
```

## คำถามทบทวน

name ใน showDevice ใช้ตรง ๆ ใน main ได้หรือไม่?

ไม่ได้ name เป็นตัวแปรภายในฟังก์ชัน showDevice ขอบเขตของตัวแปรเรียกว่า scope

## กลับไปลองอีกครั้ง

ปิดเฉลย เปลี่ยนค่าหนึ่งจุดในพื้นที่ฝึก แล้วอธิบายผลด้วยตนเอง [กลับบทเรียน](../README.md)
