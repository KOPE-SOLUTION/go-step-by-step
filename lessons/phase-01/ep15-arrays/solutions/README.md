# เฉลย EP.15

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep15-arrays`

## ข้อ 1

ตำแหน่ง 1 คือค่าที่สอง

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
25
```

## ข้อ 2

การกำหนด array คัดลอกค่า

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
[1 2]
[9 2]
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

range เริ่มจากตำแหน่ง 1 หรือ 0 และสิ้นสุดตรงไหน?

เริ่ม 0 และไปถึง len(array)-1; ถ้าไม่มีสมาชิกจะไม่เข้ารอบ

</details>

[กลับบทเรียน](../README.md)
