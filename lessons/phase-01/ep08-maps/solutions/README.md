# เฉลย EP.8

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

ก่อนเรียก `show(latest, "sensor-04")` ให้เพิ่ม `latest["sensor-04"] = 25` แล้วตรวจจำนวนอุปกรณ์ที่เหลือหลังลบ `sensor-03`

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep08-maps/solutions/01` ด้วย `go run .`

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: 25.0 C
devices: 3
```

เหตุผล: เมื่อเพิ่ม `sensor-04` ก่อนค้น จะได้ `ok` เป็น `true` หลังลบ `sensor-03` จึงเหลือ 3 อุปกรณ์

## ข้อ 2

หลังลบ `sensor-03` ด้วย `delete` ให้เรียก `show(latest, "sensor-03")` ต้องแสดง `not found`

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep08-maps/solutions/02` ด้วย `go run .`

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: not found
sensor-03: not found
devices: 2
```

เหตุผล: `delete` ลบ key ออกจาก map การค้นชื่อเดิมอีกครั้งจึงได้ `ok` เป็น `false`

[กลับบทเรียน](../README.md)
