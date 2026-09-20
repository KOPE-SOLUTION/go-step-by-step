# เฉลย EP.18

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep18-structs`

## ข้อ 1

field ที่ละไว้ได้ค่าศูนย์

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
sensor-02 0
```

## ข้อ 2

field ตัวเลขถูกคัดลอกเป็นคนละค่า

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
25 30
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

struct ต่างจาก map[string]float64 อย่างไรในตัวอย่างการวัด?

struct กำหนดช่องและชนิดของแต่ละช่องไว้ เช่น string กับ float64; map แบบนี้ใช้ key ยืดหยุ่นแต่ value ทุกตัวเป็น float64

</details>

[กลับบทเรียน](../README.md)
