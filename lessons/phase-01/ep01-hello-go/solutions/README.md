# เฉลย EP.1

ลองทำ [โจทย์ในบทเรียน](../README.md#ฝึกเอง) ก่อน แต่ละข้อเริ่มจากตัวอย่างตั้งต้นของบทนี้

## ข้อ 1

เปลี่ยนชื่ออุปกรณ์เป็น `sensor-02` และสถานะเป็น `waiting` โดยยังแสดงผล 3 บรรทัด

ดู [main.go](01/main.go)

รันจาก `lessons/phase-01/ep01-hello-go/solutions/01` ด้วย `go run .`

```text
Hello, Go!
Device: sensor-02
Status: waiting
```


เหตุผล: เปลี่ยนเฉพาะข้อมูลใน string ลำดับคำสั่งเดิมจึงยังให้ผล 3 บรรทัด

## ข้อ 2

ทำให้ชื่ออุปกรณ์อยู่คนละบรรทัดกับ `Device:` โดยเปลี่ยนคำสั่งแสดงผลเพียงหนึ่งจุด

ดู [main.go](02/main.go)

รันจาก `lessons/phase-01/ep01-hello-go/solutions/02` ด้วย `go run .`

```text
Hello, Go!
Device:
sensor-01
Status: ready
```


เหตุผล: Println เติมบรรทัดใหม่ให้ Device: ก่อนพิมพ์ชื่ออุปกรณ์


[กลับบทเรียน](../README.md)
