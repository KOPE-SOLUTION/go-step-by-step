# เฉลย EP.6

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่ม show("0") และ show("100") ต่อท้าย เพื่อยืนยันว่าปลายช่วงทั้งสองรับได้

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


เหตุผล: เงื่อนไขเดิมใช้ < 0 และ > 100 จึงยังยอมรับค่าที่เท่ากับขอบทั้งสอง

## ข้อ 2

ก่อน ParseFloat ให้ตรวจข้อความว่าง และคืน error ว่า `temperature is required` จากนั้นเปลี่ยน warm ใน main เป็นข้อความว่าง

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep06-errors/solutions/02` ด้วย `go run .`

```text
accepted: 27.5 C
ERROR: temperature is required
ERROR: temperature outside simulated range: 101.0
accepted: 30.0 C
```


เหตุผล: ตรวจข้อความว่างก่อนแปลง ทำให้ได้ error ที่ตรงเหตุ และ return จบเฉพาะรายการนั้น


[กลับบทเรียน](../README.md)
