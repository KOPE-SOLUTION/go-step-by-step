# EP.27 — โปรเจกต์สรุป ตอนที่ 3 — รายงานหลายรายการ

**เป้าหมาย:** ให้รายการผิดหนึ่งรายการไม่หยุดการประมวลผลรายการถัดไป

## 1. อ่านโค้ด

`fmt.Sprintf` จัดรูปแบบแล้วคืนข้อความโดยไม่พิมพ์ ส่วน `continue` ข้ามรอบที่ข้อมูลผิดแล้วทำรายการถัดไป

ส่วนที่เพิ่มใน [examples/model.go](examples/model.go)

```go
func buildReport(readings []Reading) []string {
	lines := []string{}
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		line := fmt.Sprintf("%s: %.1f C [%s]", reading.DeviceID, reading.Celsius, status(reading))
		lines = append(lines, line)
	}
	return lines
}
```

ไฟล์ที่ใช้ร่วมกัน: [examples/main.go](examples/main.go)

## 2. ลองรัน

**ก่อนรัน:** หลังรายการ sensor-03 ผิด sensor-04 จะยังได้รายงานหรือไม่?

รันจากโฟลเดอร์ `lessons/phase-01/ep27-project-report`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
```

buildReport คืนบรรทัดรายงานให้ main พิมพ์ จึงตรวจข้อความได้โดยไม่ต้องอ่านข้อความจากหน้าจอ

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
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: -1},
		{DeviceID: "sensor-04", Celsius: 28},
	}
	for _, line := range buildReport(readings) {
		fmt.Println(line)
	}
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

func buildReport(readings []Reading) []string {
	lines := []string{}
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		line := fmt.Sprintf("%s: %.1f C [%s]", reading.DeviceID, reading.Celsius, status(reading))
		lines = append(lines, line)
	}
	return lines
}
```

</details>

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. เพิ่ม sensor-05 ค่า 31 ต่อท้ายและดูรายงานครบห้าบรรทัด
2. ส่ง []Reading{} แล้วตรวจว่าไม่มีบรรทัดรายงานและไม่ล้ม

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run .
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

continue ต่างจาก return ใน buildReport อย่างไร?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เรามีแบบฟอร์มหลายใบ ถ้าใบหนึ่งผิด ให้บันทึกปัญหาของใบนั้นแล้วตรวจใบถัดไป รายงานยังรักษาลำดับข้อมูลเข้า

continue ข้ามส่วนที่เหลือของรอบปัจจุบันแล้วไปตรวจรอบต่อไป; fmt.Sprintf จัดรูปแบบแล้วคืน string โดยไม่พิมพ์; buildReport คืน slice ของบรรทัดรายงาน; _ ทิ้ง index ที่ไม่ได้ใช้

- buildReport รับ []Reading และสะสม []string ส่วน main รับผิดชอบพิมพ์ จึงทดสอบข้อความได้โดยตรง
- เมื่อ validate ไม่ผ่าน เพิ่มบรรทัด ERROR แล้ว continue เพื่อไม่สร้างบรรทัดสำเร็จของรายการเดียวกัน
- โค้ดทำทีละรายการตามลำดับ ไม่ใช่งานพร้อมกัน; นี่เป็นเพียงการจัดการข้อมูลจำลองในหน่วยความจำ

**ข้อผิดพลาดที่พบบ่อย**

- ใช้ return ในกรณีรายการผิด: จะจบทั้งฟังก์ชันและไม่ทำรายการถัดไป
- ลืม continue: รายการเสียจะมีทั้ง ERROR และบรรทัดสถานะที่ไม่ควรมี

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/ref/spec#Continue_statements)

</details>

[ตอนก่อนหน้า](../ep26-project-validation/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep28-project-tests/README.md)
