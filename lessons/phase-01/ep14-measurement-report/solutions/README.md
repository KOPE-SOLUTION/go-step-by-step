# เฉลย EP.14

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่ม sensor-05 ค่า 0 ใน main โดยไม่แก้กฎ รายงานต้องรับค่านี้และสรุป accepted เป็น 4/5 รัน test ขอบเขตเดิมด้วย

ดู [main.go](01/main.go) และ [main_test.go](01/main_test.go)

รันจาก `lessons/phase-01/ep14-measurement-report/solutions/01` ด้วย `go run .` และ `go test -v .`

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
sensor-05: 0.0 C [OK]
accepted: 4/5
```


เหตุผล: 0 อยู่ในช่วงที่รับได้ จึงเพิ่มจำนวนสำเร็จอีกหนึ่งโดยไม่ต้องแก้ validate

## ข้อ 2

เพิ่ม test ให้ buildReport รับรายการแรกชื่อว่างและรายการถัดไปชื่อ sensor-02 ค่า 28 ต้องคืน error บรรทัดแรกและ accepted=1 โดย main ไม่ต้องเปลี่ยน

ดู [main.go](02/main.go) และ [main_test.go](02/main_test.go)

รันจาก `lessons/phase-01/ep14-measurement-report/solutions/02` ด้วย `go run .` และ `go test -v .`

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
accepted: 3/4
```


เหตุผล: test ตรวจทั้ง error ของรายการชื่อว่างและรายการดีหลังจากนั้น เพื่อยืนยันว่าไม่หยุดทั้งรายงาน


[กลับบทเรียน](../README.md)
