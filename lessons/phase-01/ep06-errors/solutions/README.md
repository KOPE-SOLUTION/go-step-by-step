# เฉลย EP.6

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่ม `show("0")` และ `show("100")` ท้ายฟังก์ชัน `main` เพื่อตรวจว่าค่า 0 และ 100 ยังอยู่ในช่วงที่รับได้

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep06-errors/solutions/01` ด้วย `go run .`

```text
accepted: 27.5 C
ERROR: temperature must be a number: "warm"
ERROR: temperature outside simulated range: 101.0
accepted: 30.0 C
accepted: 0.0 C
accepted: 100.0 C
```

เหตุผล: เงื่อนไขเดิมปฏิเสธค่าที่ `< 0` หรือ `> 100` จึงยังรับค่า 0 และ 100

## ข้อ 2

ใน `parseCelsius` ให้ตรวจข้อความว่างก่อนเรียก `strconv.ParseFloat` และคืน error ว่า `temperature is required` จากนั้นเปลี่ยน `show("warm")` ใน `main` เป็น `show("")` เพื่อทดลอง

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep06-errors/solutions/02` ด้วย `go run .`

```text
accepted: 27.5 C
ERROR: temperature is required
ERROR: temperature outside simulated range: 101.0
accepted: 30.0 C
```

เหตุผล: การตรวจข้อความว่างก่อนแปลงทำให้แจ้งสาเหตุได้ตรงจุด เมื่อ `show` พบ error จะพิมพ์ข้อความแล้วจบการเรียกครั้งนั้น ส่วน `main` ยังเรียก `show` กับข้อมูลถัดไปได้

[กลับบทเรียน](../README.md)
