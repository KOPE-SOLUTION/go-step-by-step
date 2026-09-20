# แบบฝึกหัด EP.23 — เขียน unit test แรก

สร้างพื้นที่ฝึกจากรากหลักสูตรด้วย `./scripts/new-practice.ps1 -Episode 23` แล้วเปิด `practics/phase-01/ep23-first-test` ([วิธีใช้](../../../../docs/PRACTICE.md))

## ข้อ 1

เพิ่ม test กรณี 29.9 ต้อง false แล้วรัน go test ในพื้นที่ฝึก

## ข้อ 2

เพิ่ม test กรณี 30.1 ต้อง true แล้วทดลองเปลี่ยน >= เป็น > และดูว่า test ชุดใดจับความผิดพลาดได้ คืนโค้ดเดิมหลังทดลอง

คาดเดาผลก่อนแก้ แล้วรันจากโฟลเดอร์ฝึก:

```powershell
go run ./examples
go test ./examples
```

[เฉลยและคำอธิบาย](../solutions/README.md) · [กลับบทเรียน](../README.md)
