# เฉลย EP.25

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep25-project-reading`

## ข้อ 1

ทดสอบเกณฑ์พอดี

ดู [main.go](01/main.go) และ [model.go](01/model.go)

```powershell
go run ./solutions/01
```

ผล:

```text
sensor-02: 30.0 C [WARNING]
```

## ข้อ 2

ปัดเฉพาะตอนแสดง ส่วน 29.96 ยังต่ำกว่า 30

ดู [main.go](02/main.go) และ [model.go](02/model.go)

```powershell
go run ./solutions/02
```

ผล:

```text
sensor-03: 30.0 C [OK]
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

สถานะควรคำนวณจากค่าจริงหรือข้อความที่จัดรูปแบบแล้ว?

ค่าตัวเลขเดิมตามข้อกำหนด; ถ้าต้องการลดความสับสนในงานจริงต้องกำหนดนโยบายทศนิยมกับเกณฑ์ให้ชัด บทนี้รักษาค่าจริงไว้

</details>

[กลับบทเรียน](../README.md)
