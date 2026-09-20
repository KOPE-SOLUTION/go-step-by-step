# EP.26 — โปรเจกต์สรุป ตอนที่ 2 — ตรวจข้อมูลก่อนทำรายงาน

**เป้าหมาย:** ปฏิเสธชื่อว่างและอุณหภูมินอกช่วงจำลอง

## 1. อ่านโค้ด

Validation คือการตรวจข้อมูลก่อนใช้ บทนี้รับชื่อที่ไม่ว่างและตัวเลขจำลองตั้งแต่ 0 ถึง 100

ส่วนที่เพิ่มใน [examples/model.go](examples/model.go)

```go
func validate(reading Reading) error {
	if reading.DeviceID == "" {
		return errors.New("device ID is empty")
	}
	if reading.Celsius < 0 || reading.Celsius > 100 {
		return errors.New("temperature outside simulated range")
	}
	return nil
}
```

ไฟล์ที่ใช้ร่วมกัน: [examples/main.go](examples/main.go)

## 2. ลองรัน

**ก่อนรัน:** ข้อมูลผิดจะได้รายงานสถานะ OK ตามมาไหม?

รันจากโฟลเดอร์ `lessons/phase-01/ep26-project-validation`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
ERROR: temperature outside simulated range
```

ถ้า validate คืน error ให้แสดงปัญหาและหยุดก่อนสร้างรายงาน

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** และ `go.mod` จาก EP.21 เขียนทับ `main.go` ด้วยโค้ดตั้งต้นที่รวมไว้ด้านล่าง แล้วทำโจทย์

<details>
<summary>โค้ดตั้งต้นสำหรับ main.go</summary>

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: -1}
	err := validate(reading)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%s: %.1f C [%s]\n", reading.DeviceID, reading.Celsius, status(reading))
}

type Reading struct {
	DeviceID string
	Celsius  float64
}

func status(reading Reading) string {
	if reading.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}

func validate(reading Reading) error {
	if reading.DeviceID == "" {
		return errors.New("device ID is empty")
	}
	if reading.Celsius < 0 || reading.Celsius > 100 {
		return errors.New("temperature outside simulated range")
	}
	return nil
}
```

</details>

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. ใช้ชื่อว่างกับค่า 25 ต้องรายงาน device ID is empty
2. ใช้ sensor-01 ค่า 100 ต้องได้รายงาน WARNING

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run .
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

ถ้าชื่อว่างและค่า -1 โปรแกรมแจ้ง error ใด เพราะอะไร?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

ก่อนใส่ข้อมูลลงใบรายงาน เราตรวจว่ามีชื่อและค่าอยู่ในขอบเขตที่ตกลงกัน เหมือนตรวจแบบฟอร์มก่อนส่งต่อ

validation คือการตรวจว่าข้อมูลตรงข้อกำหนดหรือไม่; บทนี้กำหนดชื่อไม่เป็นข้อความว่าง และตัวเลขจำลองอยู่ระหว่าง 0 ถึง 100 รวมปลายทั้งสอง; ใช้ตัวเลขธรรมดาที่กำหนดในโค้ด

- validate คืน error อย่างเดียว เพราะหน้าที่คือบอกว่าผ่านหรือไม่
- ตรวจชื่อก่อนช่วงค่า ถ้าผิดทั้งสองอย่างจะรายงานข้อแรกที่พบ
- เปรียบเทียบ < 0 และ > 100 ทำให้ 0 กับ 100 ผ่าน; ยังไม่รองรับชื่อที่มีแต่ช่องว่าง ค่า NaN หรือข้อมูลจากภายนอก

**ข้อผิดพลาดที่พบบ่อย**

- เรียก status ก่อนตรวจข้อมูลจนค่าติดลบได้ OK: ตรวจ validate ก่อน
- เปลี่ยนเงื่อนไขเป็น <= 0 หรือ >= 100 ทำให้ปลายช่วงถูกปฏิเสธผิดข้อกำหนด

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/handle-errors)

</details>

[ตอนก่อนหน้า](../ep25-project-reading/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep27-project-report/README.md)
