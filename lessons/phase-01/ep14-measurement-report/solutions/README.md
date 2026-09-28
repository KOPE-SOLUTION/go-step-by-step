# เฉลย EP.14

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เพิ่มข้อมูล `sensor-05` อุณหภูมิ 0 ลงใน `readings` ของ `main` โดยไม่แก้กฎ รายงานต้องแสดง `accepted: 4/5` แล้วรัน test เดิมเพื่อตรวจว่าทุกกรณียังผ่าน

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

เหตุผล: 0 อยู่ในช่วงที่รับได้ จึงมีรายการที่ผ่านการตรวจเพิ่มอีกหนึ่ง โดยไม่ต้องแก้ `validate`

## ข้อ 2

เพิ่ม test ให้ `buildReport` รับสองรายการ: รายการแรกชื่อว่าง และรายการที่สองชื่อ `sensor-02` อุณหภูมิ 28 ตรวจว่าบรรทัดแรกของรายงานมีข้อความ error บรรทัดถัดไปแสดงข้อมูล `sensor-02` และจำนวน `accepted` เป็น 1 โดยไม่แก้ `main`

ดู [main.go](02/main.go) และ [main_test.go](02/main_test.go)

รันจาก `lessons/phase-01/ep14-measurement-report/solutions/02` ด้วย `go run .` และ `go test -v .`

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
accepted: 3/4
```

เหตุผล: test ตรวจทั้งข้อความ error ของรายการชื่อว่างและข้อมูล `sensor-02` ที่ตามมา เพื่อยืนยันว่าโปรแกรมยังสร้างรายงานต่อหลังพบข้อมูลผิด

[กลับบทเรียน](../README.md)
