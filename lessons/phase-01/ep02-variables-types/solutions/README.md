# เฉลย EP.2

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เปลี่ยนเป็น `sensor-02` ตั้งค่าอุณหภูมิล่าสุดเป็น `26.75` และแสดงทศนิยมสองตำแหน่ง

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep02-variables-types/solutions/01` ด้วย `go run .`

```text
sensor-02: 26.75 C
connected=true retries=0
label="" failed=false
```


เหตุผล: เปลี่ยนค่าที่กำหนดล่าสุด ไม่ใช่เพียงค่าเริ่มต้น และ %.2f แสดงทศนิยมสองตำแหน่ง

## ข้อ 2

ก่อนพิมพ์รายงาน กำหนด `retries = 2`, `label = "backup"` และ `failed = true` โดยไม่ประกาศตัวแปรซ้ำ

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep02-variables-types/solutions/02` ด้วย `go run .`

```text
sensor-01: 28.2 C
connected=true retries=2
label="backup" failed=true
```


เหตุผล: ตัวแปรถูกประกาศแล้วจึงกำหนดค่าใหม่ด้วย = ส่วน %q ทำให้มองเห็นข้อความชัด


[กลับบทเรียน](../README.md)
