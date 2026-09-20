# เฉลย EP.19

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep19-methods`

## ข้อ 1

ต่ำกว่าเกณฑ์จึง false

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
false
```

## ข้อ 2

method คืนตัวเลขได้เช่นเดียวกับฟังก์ชัน

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
77.0 F
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ทำไม IsWarning จึงเหมาะกับ value receiver ในตัวอย่างนี้?

มันอ่าน struct ขนาดเล็กที่มีแต่ float64 และไม่ต้องเปลี่ยนต้นฉบับ จึงใช้สำเนาได้ตรงตามงาน

</details>

[กลับบทเรียน](../README.md)
