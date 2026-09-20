# เฉลย EP.2

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep02-read-program`

## ข้อ 1

แก้เฉพาะ string แล้วบันทึก ไม่ต้องเปลี่ยนชื่อฟังก์ชัน

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
Device simulator
```

## ข้อ 2

บันทึก source เป็น UTF-8 ถ้าอักษรแสดงผิดให้แยกปัญหาฟอนต์ออกจากการ compile

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
สวัสดี Go
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

บรรทัดไหนในโปรแกรมเป็นคำสั่งที่ทำให้ข้อความแสดงบนหน้าจอ?

`fmt.Println(...)` เป็นคำสั่งที่แสดงข้อความ ส่วน `import "fmt"` ระบุว่าไฟล์นี้จะใช้ package fmt

</details>

[กลับบทเรียน](../README.md)
