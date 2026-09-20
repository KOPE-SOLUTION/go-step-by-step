# EP.28 — โปรเจกต์สรุป ตอนที่ 4 — ตรวจผลและทบทวน Phase 1

**เป้าหมาย:** ตรวจขอบเขต รายการผิด และรายการว่างด้วย test ที่ทำซ้ำได้

## 1. อ่านโค้ด

ตรวจโปรเจกต์ทั้งกรณีปกติ ค่าขอบเขต ข้อมูลเสีย และรายการว่าง เริ่มอ่าน test สั้นนี้ก่อน แล้วค่อยเปิดกรณีอื่นในไฟล์เดียวกัน

ส่วนที่เพิ่มใน [examples/model_test.go](examples/model_test.go)

```go
func TestBuildReportEmpty(t *testing.T) {
	got := buildReport([]Reading{})
	if len(got) != 0 {
		t.Errorf("empty input: got %d lines; want 0", len(got))
	}
}
```

ไฟล์ที่ใช้ร่วมกัน: [examples/main.go](examples/main.go) · [examples/model.go](examples/model.go)

## 2. ลองรัน

**ก่อนรัน:** ถ้า buildReport คืนข้อความหนึ่งบรรทัดเมื่อไม่มีข้อมูล test นี้ควรผ่านหรือไม่?

รันจากโฟลเดอร์ `lessons/phase-01/ep28-project-tests`:

```shell
go run ./examples
go test -v ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
```

ส่วน go test ควรแสดง PASS และ ok เมื่อผ่าน

ถ้าเปลี่ยนโค้ดจนพฤติกรรมเดิมเสีย test ที่เก็บไว้จะช่วยตรวจพบ

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

ใช้ `practics/main_test.go` แล้วเขียนทับด้วยโค้ดทั้งไฟล์จาก [ตัวอย่าง test](examples/model_test.go) ไฟล์ทดสอบต้องลงท้าย `_test.go` จึงแยกจาก `main.go`

**แก้ `main_test.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. เพิ่ม test ใน `main_test.go` สำหรับข้อมูลหนึ่งรายการที่ชื่อว่าง ต้องได้ข้อความ `: ERROR: device ID is empty`
2. เพิ่ม test สำหรับข้อมูลผิดสองรายการติดกัน ตามด้วยข้อมูลถูกหนึ่งรายการ ต้องได้รายงานครบสามบรรทัด

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run .
go test .
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

ก่อนบอกว่าเรียน Phase 1 จบ ควรทำอะไรได้ด้วยตนเอง?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

ครั้งนี้เราตรวจทั้งชิ้นส่วนและผลรายงานรวม เหมือนตรวจแบบฟอร์มรายใบแล้วดูว่ารายงานสรุปยังครบตามลำดับ

regression คือพฤติกรรมเดิมเสียหลังแก้โค้ด; regression test เก็บกรณีไว้ตรวจซ้ำ; ไม่มีศัพท์ syntax ใหม่ บทนี้ประกอบสิ่งที่ใช้ใน EP23–27

- TestValidate ตรวจปลายช่วงและชื่อว่าง TestStatus ตรวจใกล้เกณฑ์เตือน
- TestBuildReport ตรวจข้อความจริง ลำดับ จำนวนบรรทัด และการทำงานต่อหลังข้อมูลเสีย
- TestBuildReportEmpty ตรวจไม่มีข้อมูล; test เหล่านี้ไม่พิสูจน์เรื่อง MQTT ฐานข้อมูล restart หรือความทนทาน ซึ่งยังไม่ได้ทำ

**ข้อผิดพลาดที่พบบ่อย**

- แก้ want ให้เหมือนผลจริงโดยไม่อ่านข้อกำหนด: อาจซ่อนบั๊ก
- เปิดเฉลยก่อนลองระบุ test ที่ควรล้ม: ลองคาดเดาแล้วจึงทดลองใน practics

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/add-a-test)

</details>

[ตอนก่อนหน้า](../ep27-project-report/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md)
