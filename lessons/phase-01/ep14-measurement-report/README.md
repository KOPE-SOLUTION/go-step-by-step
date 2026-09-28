# EP.14 — โปรเจกต์สรุป: รายงานข้อมูลการวัด

**เป้าหมาย:** ประกอบความรู้พื้นฐานเป็นรายงานที่ตรวจข้อมูล ทำงานต่อเมื่อพบรายการที่ไม่ผ่านการตรวจ และมี test ยืนยัน

**ก่อนเริ่ม:** EP.1–13 โดยใช้ struct, slice, function, error และ unit test เป็นหลัก

## ทำความเข้าใจ

โปรเจกต์ท้าย Phase 1 จะตรวจข้อมูลการวัดทีละรายการ แล้วสร้างรายงานพร้อมนับจำนวนรายการที่ผ่านการตรวจ โดยกำหนดให้ชื่ออุปกรณ์ต้องไม่ว่าง อุณหภูมิอยู่ในช่วง 0–100 C และค่าตั้งแต่เกณฑ์เตือนขึ้นไปมีสถานะ `WARNING`

บทนี้ประกอบส่วนที่เคยฝึกให้ทำงานร่วมกัน `main` ทำหน้าที่พิมพ์ข้อความ ส่วน `buildReport` สร้างและคืนข้อมูลรายงาน เพื่อนำไปใช้ต่อหรือทดสอบ

## ลงมือทำ

1. เริ่มจากชนิด `Reading` และ slice ที่เก็บข้อมูล 4 รายการ เขียนลูปพิมพ์ชื่อกับอุณหภูมิให้ได้ก่อน
2. เพิ่ม `validate`, `status` และ `buildReport` ตามตัวอย่าง ก่อนรันให้คาดเดาว่าเมื่อพบค่า 101 C โปรแกรมจะยังประมวลผล `sensor-04` หรือไม่
3. **เปลี่ยนทั้ง `main.go` และ `main_test.go`** ใน `practics` เป็นตัวอย่างของบทนี้ เพราะ test จาก EP.13 อาจยังใช้กฎที่เปลี่ยนไว้ในแบบฝึกหัด รันโปรแกรมและ `go test -v .` ตรวจทั้งผลปกติ ค่าขอบเขต และข้อมูลว่าง

### ตัวอย่างเมื่อทำครบ

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

<details>
<summary>โค้ดทดสอบใน main_test.go — เปิดเมื่อถึงขั้นทดสอบ</summary>

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
