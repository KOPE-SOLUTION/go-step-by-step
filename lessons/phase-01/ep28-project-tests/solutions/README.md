# เฉลย EP.28

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep28-project-tests`

กรอบผลลัพธ์เป็นของ go run ส่วน go test จะแสดง ok เมื่อผ่าน

## ข้อ 1

ตรวจข้อความผิดพลาดของชื่อที่ว่างด้วย

ดู [main.go](01/main.go) และ [model.go](01/model.go) และ [model_test.go](01/model_test.go)

```powershell
go run ./solutions/01
go test ./solutions/01
```

ผลของ go run:

```text
: ERROR: device ID is empty
```

## ข้อ 2

ยืนยันว่าข้ามได้มากกว่าหนึ่งรายการโดยไม่หยุด

ดู [main.go](02/main.go) และ [model.go](02/model.go) และ [model_test.go](02/model_test.go)

```powershell
go run ./solutions/02
go test ./solutions/02
```

ผลของ go run:

```text
bad-01: ERROR: temperature outside simulated range
bad-02: ERROR: temperature outside simulated range
good: 25.0 C [OK]
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ก่อนบอกว่าเรียน Phase 1 จบ ควรทำอะไรได้ด้วยตนเอง?

รันและอธิบายโปรแกรม แก้ input โดยคาดผลได้ อ่าน error จัดข้อมูลเป็น struct/slice แยกฟังก์ชัน และเขียน test จากข้อกำหนด โดยไม่ต้องจำ syntax ทุกคำ; เปิดเอกสารอ้างอิงได้

</details>

[กลับบทเรียน](../README.md)
