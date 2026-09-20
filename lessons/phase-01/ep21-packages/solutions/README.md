# เฉลย EP.21

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep21-packages`

## ข้อ 1

เรียก package เดิมด้วย input ใหม่

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
false
```

## ข้อ 2

ทั้งสองค่าถึงเกณฑ์

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
true
true
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ถ้าเปลี่ยนชื่อฟังก์ชันเป็น isWarning ทำไมอีก package เรียกไม่ได้?

ชื่อที่ขึ้นต้นตัวเล็กไม่ถูก export ข้าม package; ใช้ IsWarning และเรียก sensor.IsWarning ตามชื่อ package

</details>

[กลับบทเรียน](../README.md)
