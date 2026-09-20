# เฉลย EP.17

เปิดหลังลอง [แบบฝึกหัดในบทเรียน](../README.md#practice) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep17-maps`

## ข้อ 1

สร้างด้วย literal ว่างก่อนเพิ่ม

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
28.5 true
```

## ข้อ 2

delete ทำให้ไม่พบ key

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
0 false
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

ทำไมการทดสอบไม่ควรอิงลำดับที่ range map ให้มา?

ภาษารับประกันการเข้าถึงสมาชิก แต่ไม่รับประกันลำดับการวน map จึงต้องทดสอบด้วย key หรือจัดลำดับเองเมื่อจำเป็น

</details>

[กลับบทเรียน](../README.md)
