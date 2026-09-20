# เฉลย EP.8

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep08-format-output`

## ข้อ 1

ใช้ %.2f

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
sensor-02: 28.25 C
```

## ข้อ 2

ใช้ %q ให้เห็น string ว่างและ %d สำหรับ int

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
name="" count=3
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

Printf %.1f ทำให้ค่าตัวแปรเหลือทศนิยมหนึ่งตำแหน่งด้วยหรือไม่?

ไม่ ทำเพียงจัดรูปแบบผลที่พิมพ์ ตัวแปรยังเก็บค่าเดิม

</details>

[กลับบทเรียน](../README.md)
