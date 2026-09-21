# เฉลย EP.11

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

ให้ sensor-02 ใช้ FixedSensor ค่า 28.0 แทน FailedSensor โดยไม่แก้ show

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep11-interfaces/solutions/01` ด้วย `go run .`

```text
sensor-01: 27.5 C
sensor-02: 28.0 C
sensor-03: 30.0 C
```


เหตุผล: ทั้งสองชนิดรองรับ Reader จึงเปลี่ยนค่าที่ส่งเข้า show โดยไม่แก้ show

## ข้อ 2

เพิ่ม OffsetSensor มี field Base และ Offset ชนิด float64 ให้ Read คืนผลรวม แล้วใช้เป็น sensor-03 ด้วย Base=30, Offset=-1.5

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep11-interfaces/solutions/02` ด้วย `go run .`

```text
sensor-01: 27.5 C
sensor-02: ERROR: simulated read failure
sensor-03: 28.5 C
```


เหตุผล: Read ของชนิดใหม่มีชนิดผลลัพธ์ตรงกับ interface แม้คำนวณค่าด้วยอีกวิธี


[กลับบทเรียน](../README.md)
