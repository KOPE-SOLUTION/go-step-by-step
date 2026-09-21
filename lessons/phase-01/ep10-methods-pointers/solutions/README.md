# เฉลย EP.10

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เปลี่ยนเฉพาะการเรียก Adjust ให้ปรับ -1.5 แล้วดูค่าต้นฉบับและสถานะหลัง pointer

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep10-methods-pointers/solutions/01` ด้วย `go run .`

```text
after copy: 29.0 C [OK]
after pointer: 27.5 C [OK]
```


เหตุผล: pointer ชี้ reading เดิม การเพิ่มค่าลบจึงลดอุณหภูมิต้นฉบับเหลือ 27.5

## ข้อ 2

เปลี่ยน adjustCopy ให้รับ `*Reading` และเรียกด้วย `&reading` โดยคงการปรับทั้งสองครั้งไว้ แล้วคาดเดาผลใหม่

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep10-methods-pointers/solutions/02` ด้วย `go run .`

```text
after function: 31.0 C [WARNING]
after pointer: 33.0 C [WARNING]
```


เหตุผล: ทั้งฟังก์ชันและ method แก้ต้นฉบับ ครั้งแรกเป็น 31 และครั้งต่อมาเป็น 33


[กลับบทเรียน](../README.md)
