# เฉลย EP.3

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep03-output-lines`

## ข้อ 1

ย้ายคำสั่ง READY ไปก่อนสองคำสั่งของชื่ออุปกรณ์

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
READY
Device: sensor-01
```

## ข้อ 2

เขียน string ของบรรทัดแรกให้ครบก่อนเรียก Println

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
Device: sensor-01
STOPPED
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

จำนวนครั้งที่เรียกฟังก์ชันพิมพ์เท่ากับจำนวนบรรทัดเสมอไหม?

ไม่เสมอ ขึ้นกับว่าฟังก์ชันและข้อความเพิ่ม newline หรือไม่

</details>

[กลับบทเรียน](../README.md)
