# เฉลย EP.16

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep16-slices`

## ข้อ 1

เพิ่มสมาชิกใน slice ว่างได้

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
[sensor-01]
1
```

## ข้อ 2

first และ values มอง array เดียวกันในช่วงนี้

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
[9 2 3]
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

กำหนด b := a เมื่อ a เป็น slice ทำให้สมาชิกเป็นสำเนาอิสระหรือไม่?

ไม่ใช่ มีการคัดลอกตัว slice แต่ยังอ้างถึง array เดิม การแก้สมาชิกจึงอาจเห็นร่วมกัน; append อาจเปลี่ยน array ที่อ้างถึงได้

</details>

[กลับบทเรียน](../README.md)
