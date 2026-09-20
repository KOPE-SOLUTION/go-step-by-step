# เฉลย EP.7

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep07-zero-values`

## ข้อ 1

เปรียบเทียบก่อน/หลังกำหนดค่า

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
[]
[sensor-01]
```

## ข้อ 2

false มาจาก zero value ส่วน true มาจากการกำหนดของเรา

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
false
true
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

เลขศูนย์อย่างเดียวบอกได้ไหมว่าอ่านอุปกรณ์มาแล้ว?

ไม่ได้ ต้องมีข้อมูลสถานะหรือ error ของการอ่าน ซึ่งจะเรียนต่อภายหลัง

</details>

[กลับบทเรียน](../README.md)
