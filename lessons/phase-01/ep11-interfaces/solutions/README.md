# เฉลย EP.11

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เปลี่ยนตัวจำลองของ `sensor-02` จาก `FailedSensor` เป็น `FixedSensor` ที่คืนค่า 28.0 โดยไม่แก้ฟังก์ชัน `show`

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep11-interfaces/solutions/01` ด้วย `go run .`

```text
sensor-01: 27.5 C
sensor-02: 28.0 C
sensor-03: 30.0 C
```

เหตุผล: ทั้งสองชนิดมี method ตรงตาม `Reader` จึงเปลี่ยนตัวจำลองที่ส่งให้ `show` ได้โดยไม่แก้ฟังก์ชัน

## ข้อ 2

เพิ่มชนิด `OffsetSensor` ที่มี field `Base` และ `Offset` ชนิด `float64` ให้ method `Read` คืนผลรวมของสองค่าและ `nil` แล้วใช้กับ `sensor-03` โดยกำหนด `Base: 30` และ `Offset: -1.5`

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep11-interfaces/solutions/02` ด้วย `go run .`

```text
sensor-01: 27.5 C
sensor-02: ERROR: simulated read failure
sensor-03: 28.5 C
```

เหตุผล: `Read` ของ `OffsetSensor` ตรงตามที่ `Reader` กำหนด จึงใช้แทนตัวจำลองเดิมได้ แม้จะคำนวณอุณหภูมิด้วยวิธีต่างกัน

[กลับบทเรียน](../README.md)
