# เฉลย EP.14

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep14-errors`

## ข้อ 1

ค่าติดลบอยู่นอกช่วงจำลอง

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
ERROR: temperature outside simulated range
```

## ข้อ 2

ใช้ > 100 จึงยอมรับ 100

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
Temperature: 100
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ทำไมรับค่ามาแล้วต้องตรวจ err ก่อน?

ผลตัวเลขเมื่อ err ไม่ใช่ nil อาจไม่ใช่ข้อมูลที่ใช้ได้; สัญญาของฟังก์ชันนี้กำหนดให้ใช้ผลได้เฉพาะเมื่อ err เป็น nil

</details>

[กลับบทเรียน](../README.md)
