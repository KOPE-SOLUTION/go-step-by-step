# เฉลย EP.23

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep23-first-test`

กรอบผลลัพธ์เป็นของ go run ส่วน go test จะแสดง ok เมื่อผ่าน

## ข้อ 1

เก็บ test ที่ 30 เดิมไว้ และเพิ่ม TestIsWarningBelow สำหรับค่าต่ำกว่าเกณฑ์ แต่ละ test ต้องใช้ชื่อฟังก์ชันไม่ซ้ำกัน

ดู [main.go](01/main.go) และ [threshold.go](01/threshold.go) และ [threshold_test.go](01/threshold_test.go)

```powershell
go run ./solutions/01
go test ./solutions/01
```

ผลของ go run:

```text
false
```

## ข้อ 2

เฉลยเก็บ test ที่ 30 และเพิ่ม TestIsWarningAbove สำหรับ 30.1 เมื่อเปลี่ยน >= เป็น > ตัวที่ 30 จับข้อผิดพลาดได้ แต่ตัวที่ 30.1 อย่างเดียวจับไม่ได้

ดู [main.go](02/main.go) และ [threshold.go](02/threshold.go) และ [threshold_test.go](02/threshold_test.go)

```powershell
go run ./solutions/02
go test ./solutions/02
```

ผลของ go run:

```text
true
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

test ผ่านหนึ่งกรณีพิสูจน์ว่าฟังก์ชันถูกทุกค่าหรือไม่?

ไม่ใช่ พิสูจน์เฉพาะกรณีที่ตรวจ ต้องเลือกกรณีตามข้อกำหนดและขอบเขตเพิ่ม

</details>

[กลับบทเรียน](../README.md)
