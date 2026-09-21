# EP.14 — โปรเจกต์สรุป: รายงานข้อมูลการวัด

**เป้าหมาย:** ประกอบความรู้พื้นฐานเป็นรายงานที่ตรวจข้อมูล ทำต่อหลังรายการเสีย และมี test ยืนยัน

**ก่อนเริ่ม:** EP.1–13 โดยใช้ struct, slice, function, error และ unit test เป็นหลัก

## ทำความเข้าใจ

งานปลาย Phase 1 คือรับรายการการวัด ตรวจแต่ละรายการ และคืนรายงานพร้อมจำนวนที่รับได้ เงื่อนไขเดิมคือชื่อไม่ว่าง ช่วง 0–100 C และตั้งแต่เกณฑ์ขึ้นไปเป็น WARNING

บทนี้ประกอบส่วนที่เคยฝึกให้ทำงานร่วมกัน main ดูแลการพิมพ์ ส่วน buildReport คืนข้อมูลเพื่อใช้หรือทดสอบต่อได้

## ลงมือทำ

1. เริ่มจาก Reading และ slice ข้อมูล 4 รายการ เขียนลูปพิมพ์ชื่อกับอุณหภูมิให้ได้ก่อน
2. เพิ่ม validate, status และ buildReport ตามตัวอย่าง ก่อนรันให้ทายว่ารายการ 101 C จะกระทบ sensor-04 หรือไม่
3. **เปลี่ยนทั้ง main.go และ main_test.go** ใน practics เป็นของบทนี้ เพื่อแทน test กฎทดลองจาก EP.13 รันโปรแกรมและ `go test -v .` ตรวจทั้งผลปกติ ค่าขอบเขต และข้อมูลว่าง

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

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

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
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- validate ตรวจว่ารับข้อมูลได้หรือไม่ ส่วน status ตัดสินคำเตือน ทำหน้าที่ต่างกันแม้ใช้ค่าการวัดเดียวกัน
- `fmt.Sprintf` จัดรูปแบบแล้วคืน string ต่างจาก Printf ที่พิมพ์ทันที จึงสะสมบรรทัดเพื่อทดสอบได้
- buildReport คืนสองค่า คือรายการบรรทัดและจำนวนที่รับได้ ใช้ continue เพื่อไม่ให้ข้อมูลเสียเข้าเส้นทางปกติ
- test ตรวจว่ารายการดีหลังข้อมูลเสียยังปรากฏ และรายงานไม่แก้ค่าการวัดต้นฉบับ
- `t.Fatalf` รายงานและหยุด test นั้นทันทีเมื่อจำนวนบรรทัดผิด ป้องกันการอ่าน index เกินขอบเขตในขั้นเปรียบเทียบ
- เลือกใช้ฟังก์ชันธรรมดาเพราะยังไม่มีหลายแหล่งข้อมูลให้สลับ ไม่จำเป็นต้องใช้ทุกเครื่องมือที่เรียนในทุกโปรแกรม

**ลองตรวจเมื่อผิด:** ถ้าเหลือ main_test.go จากการทดลองเปลี่ยนกฎใน EP.13 ผลอาจไม่ตรงกับกฎโปรเจกต์ ให้ใช้ชุด test ของบทนี้ ตรวจว่า main รับค่า accepted จาก buildReport แทนการใช้ len เป็นจำนวนสำเร็จ

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เพิ่ม sensor-05 ค่า 0 ใน main โดยไม่แก้กฎ รายงานต้องรับค่านี้และสรุป accepted เป็น 4/5 รัน test ขอบเขตเดิมด้วย
2. เพิ่ม test ให้ buildReport รับรายการแรกชื่อว่างและรายการถัดไปชื่อ sensor-02 ค่า 28 ต้องคืน error บรรทัดแรกและ accepted=1 โดย main ไม่ต้องเปลี่ยน

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้าต่อไปต้องคืนรายงานผ่าน API ทำไมแยก buildReport ออกจากการพิมพ์จึงช่วยได้?

<details>
<summary>แนวคำตอบ</summary>

เรียกกฎและรับผลเดิมได้โดยไม่บังคับพิมพ์ใน terminal แล้วให้ส่วน API เลือกรูปแบบคำตอบเอง

</details>

**นำไปใช้ต่อ:** จบ Phase 1 ด้วยส่วนประมวลผลที่นำไปต่อยอดไฟล์ JSON, API และฐานข้อมูลได้ ส่วนการอ่าน input จริง ค่าพิเศษ NaN/Inf การเก็บถาวร และงานพร้อมกันยังอยู่นอกการทดสอบบทนี้

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/add-a-test) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep14)

</details>

[EP.13](../ep13-unit-tests/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md)
