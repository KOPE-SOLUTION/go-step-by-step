# เฉลย EP.2

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เปลี่ยนค่า `deviceID` เป็น `"sensor-02"` แก้บรรทัด `celsius = 28.26` ให้เป็น `celsius = 26.75` และแสดงอุณหภูมิเป็นทศนิยมสองตำแหน่ง

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep02-variables-types/solutions/01` ด้วย `go run .`

```text
sensor-02: 26.75 C
connected=true retries=0
label="" failed=false
```

เหตุผล: บรรทัด `celsius = 26.75` กำหนดค่าใหม่แทนค่าเริ่มต้น `27.5` ส่วน `%.2f` แสดงทศนิยมสองตำแหน่ง

## ข้อ 2

ก่อนพิมพ์รายงาน กำหนด `retries = 2`, `label = "backup"` และ `failed = true` โดยไม่ประกาศตัวแปรซ้ำ

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep02-variables-types/solutions/02` ด้วย `go run .`

```text
sensor-01: 28.3 C
connected=true retries=2
label="backup" failed=true
```

เหตุผล: ตัวแปรถูกประกาศแล้วจึงกำหนดค่าใหม่ด้วย `=` ส่วน `%q` แสดงข้อความพร้อมเครื่องหมายคำพูด

[กลับบทเรียน](../README.md)
