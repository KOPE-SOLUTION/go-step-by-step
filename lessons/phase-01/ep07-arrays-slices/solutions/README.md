# เฉลย EP.7

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่มค่า `26.5` ด้วย append อีกหนึ่งรายการ แล้วตรวจค่าเฉลี่ยและจำนวนคำเตือน

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep07-arrays-slices/solutions/01` ด้วย `go run .`

```text
0: 27.5 C
1: 30.0 C
2: 28.5 C
3: 32.0 C
4: 26.5 C
average: 28.9 C, warnings: 2
```


เหตุผล: append เพิ่มสมาชิกที่ห้า ผลรวมเป็น 144.5 หาร 5 ได้ 28.9 โดยจำนวนที่ >=30 ยังเป็น 2

## ข้อ 2

หลังสร้าง readings ครบแล้ว ให้เปลี่ยนเป็น slice ว่าง ก่อนเริ่มคำนวณ โปรแกรมต้องพิมพ์ no readings

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep07-arrays-slices/solutions/02` ด้วย `go run .`

```text
no readings
```


เหตุผล: slice ว่างมี len=0 จึงเข้า else และไม่หารด้วยศูนย์


[กลับบทเรียน](../README.md)
