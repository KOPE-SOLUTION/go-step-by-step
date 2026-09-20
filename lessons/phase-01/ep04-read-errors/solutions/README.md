# เฉลย EP.4

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep04-read-errors`

## ข้อ 1

แก้ `fmt.println` กลับเป็น `fmt.Println` ไม่ลบ fmt. เพื่อเปลี่ยนไปใช้ฟังก์ชันอื่น

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
Check complete
```

## ข้อ 2

เมื่อไม่ import fmt ชื่อ package นี้ใช้ไม่ได้ในไฟล์ คืน import แล้วแก้ string

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
Fixed
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

go ไม่เป็นที่รู้จัก กับ undefined: fmt.println ต่างกันที่ขั้นไหน?

แบบแรก terminal ยังเรียกเครื่องมือ Go ไม่ได้ แบบหลังเรียก Go ได้แต่ compiler ตรวจ source ไม่ผ่าน

</details>

[กลับบทเรียน](../README.md)
