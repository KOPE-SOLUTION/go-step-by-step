# เฉลย EP.5

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep05-variables`

## ข้อ 1

ใช้ := ครั้งแรกและ = ครั้งถัดไป

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
sensor-03
sensor-04
```

## ข้อ 2

บรรทัดหนึ่งเป็น string literal อีกบรรทัดอ่านตัวแปร

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
deviceName
sensor-01
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ทำไมเปลี่ยนค่าครั้งที่สองจึงใช้ = แทน :=?

เรากำลังกำหนดค่าใหม่ให้ตัวแปรที่ประกาศแล้ว ไม่ได้สร้างชื่อใหม่

</details>

[กลับบทเรียน](../README.md)
