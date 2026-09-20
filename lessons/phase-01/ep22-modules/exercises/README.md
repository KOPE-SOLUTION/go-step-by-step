# แบบฝึกหัด EP.22 — อ่าน go.mod และเข้าใจขอบเขต module

สร้างพื้นที่ฝึกจากรากหลักสูตรด้วย `./scripts/new-practice.ps1 -Episode 22` แล้วเปิด `practics/phase-01/ep22-modules` ([วิธีใช้](../../../../docs/PRACTICE.md))

## ข้อ 1

รัน Get-Content go.mod และ go list ./... ในพื้นที่ฝึก แล้วเขียนอธิบายว่าอันใดคือ module อันใดคือ package จากนั้นรันด้วยค่า 25

## ข้อ 2

เปลี่ยน input เป็น 35 โดยไม่แก้ go.mod แล้วรันใหม่ ระบุเหตุผลว่าแก้ข้อมูลต้องเปลี่ยน module หรือไม่

คาดเดาผลก่อนแก้ แล้วรันจากโฟลเดอร์ฝึก:

```powershell
go run ./examples
```

[เฉลยและคำอธิบาย](../solutions/README.md) · [กลับบทเรียน](../README.md)
