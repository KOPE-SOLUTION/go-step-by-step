# เฉลย EP.9

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep09-arithmetic`

## ข้อ 1

แปลงก่อนหาร

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
3.5
```

## ข้อ 2

กำหนด c เป็น float64 ก่อน สูตรนี้จึงคำนวณแบบทศนิยม

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
77.0
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ทำไมต้องแปลงก่อนหารเมื่ออยากรักษาส่วนเศษ?

เพราะเมื่อหาร int เสร็จแล้ว ส่วนเศษถูกตัดไป การแปลงทีหลังไม่ทำให้กลับคืน

</details>

[กลับบทเรียน](../README.md)
