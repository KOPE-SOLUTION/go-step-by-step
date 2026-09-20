# เฉลย EP.24

หยุดก่อนถ้ายังไม่ได้ลอง [แบบฝึกหัด](../exercises/README.md) คำตอบไม่จำเป็นต้องเขียนเหมือนกันทุกตัวอักษร แต่ต้องตรงข้อกำหนด

คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep24-boundary-tests` ไม่ใช่โฟลเดอร์ practics ซึ่งไม่มีเฉลย

คำสั่ง go test จะมีบรรทัด ok เมื่อผ่าน (เวลาอาจต่างกัน) ผลในกรอบด้านล่างเป็นข้อความจาก go run เท่านั้น หาก test ไม่ผ่านจะมี FAIL และรายละเอียดกรณีที่ผิด

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

## คำถามทบทวน

เหตุใดค่า 30 พอดีจึงสำคัญกว่าทดสอบ 31 หลายครั้ง?

ค่า 30 แยกความหมายของ > กับ >= ได้ ส่วน 31 ให้ผลเหมือนกันทั้งคู่

## กลับไปลองอีกครั้ง

ปิดเฉลย เปลี่ยนค่าหนึ่งจุดในพื้นที่ฝึก แล้วอธิบายผลด้วยตนเอง [กลับบทเรียน](../README.md)
