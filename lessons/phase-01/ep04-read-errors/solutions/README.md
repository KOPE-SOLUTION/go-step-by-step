# เฉลย EP.4

หยุดก่อนถ้ายังไม่ได้ลอง [แบบฝึกหัด](../exercises/README.md) คำตอบไม่จำเป็นต้องเขียนเหมือนกันทุกตัวอักษร แต่ต้องตรงข้อกำหนด

คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep04-read-errors` ไม่ใช่โฟลเดอร์ practics ซึ่งไม่มีเฉลย

## ข้อ 1

แก้ `fmt.println` กลับเป็น `fmt.Println` ไม่ลบ fmt. เพื่อเปลี่ยนไปใช้ฟังก์ชันอื่น

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
Check complete
```

## ข้อ 2

เมื่อไม่ import fmt ชื่อ package นี้ใช้ไม่ได้ในไฟล์ คืน import แล้วแก้ string

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
Fixed
```

## คำถามทบทวน

go ไม่เป็นที่รู้จัก กับ undefined: fmt.println ต่างกันที่ขั้นไหน?

แบบแรก terminal ยังเรียกเครื่องมือ Go ไม่ได้ แบบหลังเรียก Go ได้แต่ compiler ตรวจ source ไม่ผ่าน

## กลับไปลองอีกครั้ง

ปิดเฉลย เปลี่ยนค่าหนึ่งจุดในพื้นที่ฝึก แล้วอธิบายผลด้วยตนเอง [กลับบทเรียน](../README.md)
