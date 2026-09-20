# เฉลย EP.24

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep24-boundary-tests`

กรอบผลลัพธ์เป็นของ go run ส่วน go test จะแสดง ok เมื่อผ่าน

## ข้อ 1

ค่าต่ำกว่าเกณฑ์ต้องไม่เตือน

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

เมื่อข้อกำหนดเปลี่ยน ขอบเขตและคำตอบต้องเปลี่ยนตาม

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

เหตุใดค่า 30 พอดีจึงสำคัญกว่าทดสอบ 31 หลายครั้ง?

ค่า 30 แยกความหมายของ > กับ >= ได้ ส่วน 31 ให้ผลเหมือนกันทั้งคู่

</details>

[กลับบทเรียน](../README.md)
