# เฉลย EP.20

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep20-pointers`

## ข้อ 1

ส่ง pointer ของตัวแปรที่มีอยู่

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
19.5
```

## ข้อ 2

p ชี้ไปที่ n

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
6
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

เครื่องหมาย * ใน *float64 และ *value ทำหน้าที่เหมือนกันหรือไม่?

ไม่เหมือน: *float64 ใช้ระบุชนิด pointer ส่วน *value ใช้อ่านหรือเขียนค่าที่ pointer ชี้ในนิพจน์

</details>

[กลับบทเรียน](../README.md)
