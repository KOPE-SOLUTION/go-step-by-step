# เฉลย EP.12

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep12-function-parameters`

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

<details>
<summary>คำตอบคำถามทบทวน</summary>

name ใน showDevice ใช้ตรง ๆ ใน main ได้หรือไม่?

ไม่ได้ name เป็นตัวแปรภายในฟังก์ชัน showDevice ขอบเขตของตัวแปรเรียกว่า scope

</details>

[กลับบทเรียน](../README.md)
