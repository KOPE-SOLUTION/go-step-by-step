# EP.14 — โปรเจกต์สรุป: รายงานข้อมูลการวัด

**เป้าหมาย:** ประกอบความรู้พื้นฐานเป็นรายงานที่ตรวจข้อมูล ทำงานต่อเมื่อพบรายการที่ไม่ผ่านการตรวจ และมี test ยืนยัน

**ก่อนเริ่ม:** EP.1–13 โดยใช้ struct, slice, function, error และ unit test เป็นหลัก

## ทำความเข้าใจ

โปรเจกต์ท้าย Phase 1 จะตรวจข้อมูลการวัดทีละรายการ แล้วสร้างรายงานพร้อมนับจำนวนรายการที่ผ่านการตรวจ โดยกำหนดให้ชื่ออุปกรณ์ต้องไม่ว่าง อุณหภูมิอยู่ในช่วง 0–100 C และค่าตั้งแต่เกณฑ์เตือนขึ้นไปมีสถานะ `WARNING`

บทนี้ประกอบส่วนที่เคยฝึกให้ทำงานร่วมกัน `main` ทำหน้าที่พิมพ์ข้อความ ส่วน `buildReport` สร้างและคืนข้อมูลรายงาน เพื่อนำไปใช้ต่อหรือทดสอบ

## ลงมือทำทีละขั้น

ใช้ `practics` และ `go.mod` เดิม เริ่มจาก `main.go` ตามขั้นที่ 1 แล้วเพิ่มส่วนประมวลผลทีละอย่าง รัน `go run .` จาก `practics` หลังแต่ละขั้น ส่วน `main_test.go` จะเปลี่ยนเป็นของบทนี้เมื่อถึงขั้นทดสอบ

### 1. สร้างข้อมูลและพิมพ์ทุกรายการ

เริ่ม `practics/main.go` ด้วยข้อมูลจำลองสี่รายการ:

```go
package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: 101},
		{DeviceID: "sensor-04", Celsius: 28},
	}
	for _, reading := range readings {
		fmt.Printf("%s: %.1f C\n", reading.DeviceID, reading.Celsius)
	}
}
```

ขั้นนี้ยังแสดงทุกค่าโดยไม่ตรวจช่วง เราจะใช้ผลนี้เปรียบเทียบกับขั้นที่เพิ่มการตรวจข้อมูล

**ลองคิดก่อนรัน:** รายการไหนควรถูกปฏิเสธตามช่วง 0–100 ของบทนี้?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 27.5 C
sensor-02: 30.0 C
sensor-03: 101.0 C
sensor-04: 28.0 C
```

</details>

### 2. ตรวจข้อมูลและข้ามรายการที่ผิด

เพิ่ม `validate` เหนือ `main`:

```go
func validate(r Reading) error {
	if r.DeviceID == "" {
		return fmt.Errorf("device ID is required")
	}
	if r.Celsius < 0 || r.Celsius > 100 {
		return fmt.Errorf("temperature outside simulated range")
	}
	return nil
}
```

แทน `main` เพื่อเรียก `validate` ก่อนแสดงค่าการวัด:

```go
func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: 101},
		{DeviceID: "sensor-04", Celsius: 28},
	}
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			fmt.Printf("%s: ERROR: %v\n", reading.DeviceID, err)
			continue
		}
		fmt.Printf("%s: %.1f C\n", reading.DeviceID, reading.Celsius)
	}
}
```

`validate` คืน error เมื่อชื่อว่างหรืออุณหภูมิเกินช่วงที่กำหนด `%v` ใช้แสดงข้อความ error แล้ว `continue` ข้ามไปตรวจรายการถัดไป

**ลองคิดก่อนรัน:** หลัง sensor-03 ที่ผิด จะยังเห็น sensor-04 หรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 27.5 C
sensor-02: 30.0 C
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C
```

</details>

### 3. แยกการสร้างรายงานออกจากการพิมพ์

`fmt.Sprintf` จัดรูปแบบและคืนข้อความโดยยังไม่พิมพ์ เพิ่ม `buildReport` เหนือ `main` เพื่อเก็บแต่ละบรรทัดใน slice:

```go
func buildReport(readings []Reading) ([]string, int) {
	lines := []string{}
	accepted := 0
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		lines = append(lines, fmt.Sprintf("%s: %.1f C", reading.DeviceID, reading.Celsius))
		accepted++
	}
	return lines, accepted
}
```

แทน `main` ให้รับบรรทัดรายงานและจำนวนรายการที่ผ่านการตรวจมาพิมพ์:

```go
func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: 101},
		{DeviceID: "sensor-04", Celsius: 28},
	}
	lines, accepted := buildReport(readings)
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Printf("accepted: %d/%d\n", accepted, len(readings))
}
```

`buildReport` คืน `[]string` กับ `int` โดยนับ `accepted` เฉพาะรายการที่ผ่านการตรวจ การแยกหน้าที่นี้ช่วยให้นำรายงานไปตรวจด้วย test ได้

**ลองคิดก่อนรัน:** ทำไม accepted จึงไม่เท่ากับ len(readings)?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 27.5 C
sensor-02: 30.0 C
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C
accepted: 3/4
```

</details>

### 4. เพิ่มสถานะโดยให้ผู้เรียกกำหนดเกณฑ์

เพิ่มฟังก์ชัน `status` เหนือ `main`:

```go
func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}
```

แทน `buildReport` ทั้งฟังก์ชัน เพื่อเพิ่ม parameter `threshold` และสถานะในบรรทัดที่ผ่านการตรวจ:

```go
func buildReport(readings []Reading, threshold float64) ([]string, int) {
	lines := []string{}
	accepted := 0
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		lines = append(lines, fmt.Sprintf("%s: %.1f C [%s]",
			reading.DeviceID, reading.Celsius, status(reading.Celsius, threshold)))
		accepted++
	}
	return lines, accepted
}
```

ใน `main` แทนบรรทัดเรียก `buildReport` ด้วยบรรทัดนี้:

```go
lines, accepted := buildReport(readings, 30)
```

`validate` ตัดสินว่ารับข้อมูลได้หรือไม่ ส่วน `status` ตรวจเกณฑ์เตือนหลังข้อมูลผ่านการตรวจแล้ว ทั้งสองหน้าที่จึงไม่ใช้แทนกัน

**ลองคิดก่อนรัน:** sensor-02 ควรเป็น WARNING ส่วน sensor-03 ควรเป็น ERROR เพราะอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
accepted: 3/4
```

</details>

### 5. เขียน test ว่ารายการถัดไปยังทำงาน

เก็บคำตอบเดิมเป็น `.txt` หากต้องการ แล้วแทน `practics/main_test.go` ทั้งไฟล์ด้วย test ของบทนี้:

```go
package main

import "testing"

func TestReportContinues(t *testing.T) {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 101},
		{DeviceID: "sensor-02", Celsius: 28},
	}
	lines, accepted := buildReport(readings, 30)
	if accepted != 1 || len(lines) != 2 {
		t.Fatalf("accepted=%d lines=%v; want 1 accepted and 2 lines", accepted, lines)
	}
	if lines[1] != "sensor-02: 28.0 C [OK]" {
		t.Errorf("second line = %q; want sensor-02 to continue", lines[1])
	}
}
```

รันจาก terminal ที่ `practics`:

```shell
go test -v .
```

`t.Fatalf` รายงานความผิดพลาดและหยุด test นี้ทันที ถ้าจำนวนบรรทัดไม่ถูกต้องจึงไม่ไปอ่าน `lines[1]` ต่อ ส่วน `t.Errorf` รายงานว่าข้อความจริงต่างจากที่ต้องการ

**ลองคิดก่อนรัน:** test นี้ตรวจพฤติกรรมใดที่สำคัญเมื่อบางรายการมีข้อมูลผิด?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ควรเห็น `PASS: TestReportContinues` ยืนยันกรณีตัวอย่างที่รายการแรกผิด แต่รายการที่สองยังอยู่ในรายงาน

</details>

### 6. เพิ่มกรณีขอบเขตและข้อมูลว่าง

เมื่อ test แรกผ่านแล้ว แทน `main_test.go` ทั้งไฟล์ด้วยชุดนี้ เพื่อเพิ่มกรณีชื่อว่าง ช่วงอุณหภูมิ ข้อมูลว่าง และเกณฑ์ที่เปลี่ยนได้:

<details>
<summary>โค้ดชุดทดสอบเพิ่มเติม — เปิดเมื่อพร้อมเพิ่มกรณี</summary>

```go
package main

import "testing"

func TestValidate(t *testing.T) {
	cases := []struct {
		name    string
		reading Reading
		wantErr bool
	}{
		{"normal", Reading{"sensor-01", 27.5}, false},
		{"empty ID", Reading{"", 27.5}, true},
		{"below range", Reading{"sensor-01", -0.1}, true},
		{"lower boundary", Reading{"sensor-01", 0}, false},
		{"upper boundary", Reading{"sensor-01", 100}, false},
		{"above range", Reading{"sensor-01", 100.1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate(tc.reading)
			if (err != nil) != tc.wantErr {
				t.Errorf("validate(%+v) error = %v; wantErr %t", tc.reading, err, tc.wantErr)
			}
		})
	}
}

func TestBuildReportContinuesAfterInvalid(t *testing.T) {
	input := []Reading{
		{DeviceID: "sensor-01", Celsius: 29.9},
		{DeviceID: "sensor-02", Celsius: 101},
		{DeviceID: "sensor-03", Celsius: 30},
	}
	lines, accepted := buildReport(input, 30)
	want := []string{
		"sensor-01: 29.9 C [OK]",
		"sensor-02: ERROR: temperature outside simulated range",
		"sensor-03: 30.0 C [WARNING]",
	}
	if accepted != 2 || len(lines) != len(want) {
		t.Fatalf("accepted=%d lines=%v; want 2 accepted and %d lines", accepted, lines, len(want))
	}
	for i, line := range lines {
		if line != want[i] {
			t.Errorf("line %d = %q; want %q", i, line, want[i])
		}
	}
	if input[1].Celsius != 101 {
		t.Error("report changed the original measurement")
	}
}

func TestBuildReportEmpty(t *testing.T) {
	lines, accepted := buildReport(nil, 30)
	if len(lines) != 0 || accepted != 0 {
		t.Errorf("empty input: lines=%v accepted=%d; want no lines and zero accepted", lines, accepted)
	}
}

func TestBuildReportCustomThreshold(t *testing.T) {
	lines, accepted := buildReport([]Reading{{DeviceID: "sensor-01", Celsius: 30}}, 35)
	if accepted != 1 || len(lines) != 1 || lines[0] != "sensor-01: 30.0 C [OK]" {
		t.Errorf("custom threshold: lines=%v accepted=%d", lines, accepted)
	}
}
```

</details>

รันจาก `practics` อีกครั้ง:

```shell
go test -v .
```

ชุดนี้ใช้ตารางทดสอบแบบ EP.13 ร่วมกับ test รายงาน ตรวจว่าข้อมูลที่ผิดไม่ขัดขวางรายการถัดไป และตรวจว่าค่าการวัดต้นฉบับไม่ถูกเปลี่ยน

**ลองคิดก่อนรัน:** ถ้าไม่มีค่าการวัดเลย รายงานควรมีกี่บรรทัดและ accepted ควรเป็นเท่าไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ควรเห็น `PASS` ของ `TestValidate` และ test รายงานทั้งสามกรณี ชุดนี้ยืนยันเฉพาะข้อมูลที่ระบุใน test

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func validate(r Reading) error {
	if r.DeviceID == "" {
		return fmt.Errorf("device ID is required")
	}
	if r.Celsius < 0 || r.Celsius > 100 {
		return fmt.Errorf("temperature outside simulated range")
	}
	return nil
}

func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func buildReport(readings []Reading, threshold float64) ([]string, int) {
	lines := []string{}
	accepted := 0
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		lines = append(lines, fmt.Sprintf("%s: %.1f C [%s]",
			reading.DeviceID, reading.Celsius, status(reading.Celsius, threshold)))
		accepted++
	}
	return lines, accepted
}

func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: 101},
		{DeviceID: "sensor-04", Celsius: 28},
	}
	lines, accepted := buildReport(readings, 30)
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Printf("accepted: %d/%d\n", accepted, len(readings))
}
```

ไฟล์ `main_test.go` ใช้ชุดทดสอบจากขั้นที่ 6 หรือเปิดเทียบ [ไฟล์ทดสอบฉบับเต็ม](main_test.go)

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run .
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep14-measurement-report` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

รันทดสอบจากโฟลเดอร์เดียวกันด้วย `go test -v .` เมื่อทุกกรณีผ่านจะมี PASS และ ok เวลาและชื่อ package ที่แสดงอาจต่างกัน

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าหลังปฏิเสธค่า 101 C ของ `sensor-03` แล้ว รายงานจะยังมี `sensor-04` หรือไม่ และจะรับข้อมูลได้กี่รายการ

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
accepted: 3/4
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `validate` ตรวจว่าข้อมูลผ่านเงื่อนไขหรือไม่ ส่วน `status` ตรวจว่าอุณหภูมิถึงเกณฑ์เตือนหรือยัง จึงทำหน้าที่ต่างกัน
- `fmt.Sprintf` จัดรูปแบบแล้วคืนข้อความเป็น `string` โดยยังไม่พิมพ์ จึงเก็บแต่ละบรรทัดใน slice และตรวจข้อความด้วย test ได้
- `buildReport` คืนสองค่า คือ slice ของบรรทัดรายงานและจำนวนรายการที่ผ่านการตรวจ เมื่อพบข้อมูลผิดจะเพิ่มบรรทัด error แล้วใช้ `continue` ข้ามไปตรวจรายการถัดไป
- test ตรวจว่ารายการที่ผ่านเงื่อนไขยังปรากฏหลังรายการที่ผิด และการสร้างรายงานไม่แก้ค่าการวัดต้นฉบับ
- `t.Fatalf` รายงานและหยุด test นั้นทันทีเมื่อจำนวนบรรทัดผิด ป้องกันการอ่าน index เกินขอบเขตในขั้นเปรียบเทียบ
- เลือกใช้ฟังก์ชันธรรมดาเพราะยังไม่มีหลายแหล่งข้อมูลให้สลับ ไม่จำเป็นต้องใช้ทุกเครื่องมือที่เรียนในทุกโปรแกรม

**ข้อผิดพลาดที่พบบ่อย:** ถ้ายังใช้ `main_test.go` ที่เปลี่ยนกฎไว้ใน EP.13 ผลที่ต้องการอาจไม่ตรงกับกฎโปรเจกต์ ให้ใช้ชุด test ของบทนี้ และตรวจว่า `main` รับจำนวนรายการที่ผ่านการตรวจจาก `accepted` ซึ่ง `buildReport` คืนมา ไม่ใช้ `len(readings)` ซึ่งนับรวมรายการที่ผิดด้วย

</details>

## ฝึกเอง

ฝึกใน `practics` เดิม ก่อนทำแต่ละข้อให้เริ่มจากตัวอย่างทั้ง `main.go` และ `main_test.go` ของบทนี้ หากต้องการเก็บคำตอบก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่มข้อมูล `sensor-05` อุณหภูมิ 0 ลงใน `readings` ของ `main` โดยไม่แก้กฎ รายงานต้องแสดง `accepted: 4/5` แล้วรัน test เดิมเพื่อตรวจว่าทุกกรณียังผ่าน
2. เพิ่ม test ให้ `buildReport` รับสองรายการ: รายการแรกชื่อว่าง และรายการที่สองชื่อ `sensor-02` อุณหภูมิ 28 ตรวจว่าบรรทัดแรกของรายงานมีข้อความ error บรรทัดถัดไปแสดงข้อมูล `sensor-02` และจำนวน `accepted` เป็น 1 โดยไม่แก้ `main`

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้าต่อไปต้องคืนรายงานผ่าน API การแยก `buildReport` ออกจากการพิมพ์ช่วยอย่างไร?

<details>
<summary>แนวคำตอบ</summary>

ส่วน API เรียก `buildReport` เพื่อรับข้อมูลรายงานได้ แล้วนำไปจัดรูปแบบคำตอบโดยไม่ต้องอ่านข้อความจาก terminal

</details>

**นำไปใช้ต่อ:** จบ Phase 1 ด้วยส่วนประมวลผลที่นำไปต่อยอดไฟล์ JSON, API และฐานข้อมูลได้ ส่วนการรับข้อมูลจากภายนอก ค่าพิเศษ NaN/Inf การเก็บถาวร และงานพร้อมกันยังอยู่นอกการทดสอบบทนี้

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/add-a-test) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep14)

</details>

[EP.13](../ep13-unit-tests/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md)
