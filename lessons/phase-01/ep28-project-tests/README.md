# EP.28 — โปรเจกต์เล็ก 4 — ตรวจผลและทบทวน Phase 1

Phase 1 — Go Basic • [สารบัญ](../../../docs/playlist-01-go-basic/README.md)

## เป้าหมายและสิ่งที่ควรรู้ก่อน

ตรวจขอบเขต รายการผิด และรายการว่างด้วย test ที่ทำซ้ำได้ ก่อนเริ่มควรทำ EP.27 ได้และอธิบายตัวอย่างด้วยคำของตนเอง ไม่จำเป็นต้องจำโค้ดจากความจำ

บทนี้มีหนึ่งเรื่องหลัก อ่านเป็นช่วงได้: ทำความเข้าใจ → คาดเดา → รัน → เปลี่ยนทีละจุด ถ้าต้องกลับไปทบทวนใช้ [ตอนก่อนหน้า](../ep27-project-report/README.md)

## แนวคิดและคำใหม่

ครั้งนี้เราตรวจทั้งชิ้นส่วนและผลรายงานรวม เหมือนตรวจแบบฟอร์มรายใบแล้วดูว่ารายงานสรุปยังครบตามลำดับ

regression คือพฤติกรรมเดิมเสียหลังแก้โค้ด; regression test เก็บกรณีไว้ตรวจซ้ำ; ไม่มีศัพท์ syntax ใหม่ บทนี้ประกอบสิ่งที่ใช้ใน EP23–27

## หยุดคาดเดาก่อนเปิดผล

ถ้าเปลี่ยน continue เป็น return lines test ใดควรจับได้?

จดคำตอบสั้น ๆ ก่อนรัน ผิดได้ เก็บเหตุผลที่คิดไว้เทียบกับสิ่งที่เกิดขึ้น

## ตัวอย่างเล็กที่รันได้

### examples/main.go

```go
package main

import "fmt"

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
```

### examples/model.go

```go
package main

import (
	"errors"
	"fmt"
)

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

### examples/model_test.go

อ่านทีละฟังก์ชันทดสอบ พักและลองรันก่อนอ่านฟังก์ชันถัดไปได้

```go
package main

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		reading   Reading
		wantError bool
	}{
		{name: "lower", reading: Reading{DeviceID: "sensor-01", Celsius: 0}, wantError: false},
		{name: "upper", reading: Reading{DeviceID: "sensor-01", Celsius: 100}, wantError: false},
		{name: "below", reading: Reading{DeviceID: "sensor-01", Celsius: -0.1}, wantError: true},
		{name: "above", reading: Reading{DeviceID: "sensor-01", Celsius: 100.1}, wantError: true},
		{name: "empty ID", reading: Reading{Celsius: 25}, wantError: true},
	}
	for _, test := range tests {
		err := validate(test.reading)
		gotError := err != nil
		if gotError != test.wantError {
			t.Errorf("%s: error = %v; wantError %v", test.name, err, test.wantError)
		}
	}
}

func TestStatus(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{input: 29.9, want: "OK"},
		{input: 30, want: "WARNING"},
		{input: 30.1, want: "WARNING"},
	}
	for _, test := range tests {
		got := status(Reading{DeviceID: "sensor-01", Celsius: test.input})
		if got != test.want {
			t.Errorf("status(%v) = %q; want %q", test.input, got, test.want)
		}
	}
}

func TestBuildReport(t *testing.T) {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 30},
		{DeviceID: "broken", Celsius: -1},
		{DeviceID: "sensor-02", Celsius: 25},
	}
	want := []string{
		"sensor-01: 30.0 C [WARNING]",
		"broken: ERROR: temperature outside simulated range",
		"sensor-02: 25.0 C [OK]",
	}
	got := buildReport(readings)
	if len(got) != len(want) {
		t.Fatalf("line count = %d; want %d", len(got), len(want))
	}
	for index, line := range got {
		if line != want[index] {
			t.Errorf("line %d = %q; want %q", index, line, want[index])
		}
	}
}

func TestBuildReportEmpty(t *testing.T) {
	got := buildReport([]Reading{})
	if len(got) != 0 {
		t.Errorf("empty input: got %d lines; want 0", len(got))
	}
}
```

## รันและตรวจผล

เปิด PowerShell แล้วเลือกโฟลเดอร์บทนี้ (เครื่องอื่นให้เปลี่ยนส่วนหน้าของตำแหน่งตามที่วางหลักสูตร):

```powershell
Set-Location -LiteralPath 'C:\Users\kopes\Documents\Codex\2026-09-17\go-youtube-go-iot-edge-gateway\Go\lessons\phase-01\ep28-project-tests'
go run ./examples
```

ผลที่คาดหวัง:

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
```

ถ้าผลต่าง ให้ตรวจตำแหน่งด้วย `Get-Location` ตรวจไฟล์ที่บันทึก และอ่าน error แรกก่อนแก้ รัน test เพิ่มจากตำแหน่งเดิม:

```powershell
go test -v ./examples
```

เมื่อผ่านจะมี PASS และบรรทัด ok; เวลาและรูปแบบบรรทัดเสริมอาจต่างตามเวอร์ชัน หากตั้งใจทำให้ผิด จะเห็น FAIL พร้อมกรณีที่ไม่ผ่าน

## ทำไมโค้ดจึงทำงานแบบนี้

- TestValidate ตรวจปลายช่วงและชื่อว่าง TestStatus ตรวจใกล้เกณฑ์เตือน
- TestBuildReport ตรวจข้อความจริง ลำดับ จำนวนบรรทัด และการทำงานต่อหลังข้อมูลเสีย
- TestBuildReportEmpty ตรวจไม่มีข้อมูล; test เหล่านี้ไม่พิสูจน์เรื่อง MQTT ฐานข้อมูล restart หรือความทนทาน ซึ่งยังไม่ได้ทำ

**แยกประเภท:** syntax และชนิดข้อมูลเป็นหลักของภาษา; การแยกหน้าที่ให้ส่วนเล็กอ่านง่ายเป็นแนวทางออกแบบ; ชื่อโฟลเดอร์และการใช้หนึ่ง module ต่อ EP เป็นทางเลือกของหลักสูตร เกณฑ์อุณหภูมิที่พบเป็นข้อกำหนดจำลอง

## ลองแก้ในพื้นที่ฝึก

เปิด [วิธีสร้างพื้นที่ฝึก](../../../docs/PRACTICE.md) แล้วสร้าง EP.28 ด้วย `./scripts/new-practice.ps1 -Episode 28` จากราก Go เพียงครั้งเดียว จากนั้นเปิด `practics/phase-01/ep28-project-tests` แก้ไฟล์ใน `examples` แล้วรัน `go run ./examples` จากโฟลเดอร์ EP ที่คัดลอกแล้ว

พื้นที่ฝึกไม่รวมเฉลย และคำสั่งสร้างจะหยุดถ้าโฟลเดอร์มีอยู่แล้ว จึงไม่ทับงานที่เคยทำ

## ข้อผิดพลาดที่พบบ่อย

- แก้ want ให้เหมือนผลจริงโดยไม่อ่านข้อกำหนด: อาจซ่อนบั๊ก
- เปิดเฉลยก่อนลองระบุ test ที่ควรล้ม: ลองคาดเดาแล้วจึงทดลองใน practics

## แบบฝึกหัดและคำถามทบทวน

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ก่อนดูเฉลย คาดผลของแต่ละข้อแล้วเปลี่ยนเพียงจุดที่โจทย์ระบุ

คำถามตรวจความเข้าใจ: ก่อนบอกว่าเรียน Phase 1 จบ ควรทำอะไรได้ด้วยตนเอง?

[เฉลยและเหตุผล](solutions/README.md) แยกไว้เพื่อเปิดหลังลองเอง ถ้ายังไม่ผ่าน ให้จดสิ่งที่คาดกับผลจริงไว้ก่อนขอความช่วยเหลือ

## สิ่งที่ได้และขอบเขต

มีโปรแกรมรายงานจำลองพร้อม test และรู้ขอบเขตที่ยังไม่ได้ทำ พร้อมเลือกทบทวนก่อนขอเริ่ม Phase 2

ใช้ข้อมูลจำลองภายในโปรแกรมเท่านั้น ยังไม่ได้อ่านอุปกรณ์ ส่ง MQTT เรียก API หรือเก็บฐานข้อมูล โปรเจกต์นี้ใช้ข้อความชื่อและตัวเลขปกติที่เขียนไว้ในโค้ด ยังไม่ครอบคลุม NaN ชื่อที่มีแต่ช่องว่าง หรือข้อมูลภายนอก

ผลตรวจจริงและส่วนที่ยังไม่ได้ทดสอบ: [tests/RESULTS.md](tests/RESULTS.md)

อ้างอิงหลัก: [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/add-a-test)

เมื่ออธิบายผลและทำแบบฝึกหัดได้ ค่อยเลือก [ตอนถัดไป / สารบัญ](../../../docs/playlist-01-go-basic/README.md) ไม่มีข้อกำหนดให้เรียนหลายตอนในวันเดียว
