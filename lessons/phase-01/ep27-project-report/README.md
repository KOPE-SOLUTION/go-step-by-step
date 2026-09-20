# EP.27 — โปรเจกต์เล็ก 3 — รายงานหลายรายการ

Phase 1 — Go Basic • [สารบัญ](../../../docs/playlist-01-go-basic/README.md)

## เป้าหมายและสิ่งที่ควรรู้ก่อน

ให้รายการผิดหนึ่งรายการไม่หยุดการประมวลผลรายการถัดไป ก่อนเริ่มควรทำ EP.26 ได้และอธิบายตัวอย่างด้วยคำของตนเอง ไม่จำเป็นต้องจำโค้ดจากความจำ

บทนี้มีหนึ่งเรื่องหลัก อ่านเป็นช่วงได้: ทำความเข้าใจ → คาดเดา → รัน → เปลี่ยนทีละจุด ถ้าต้องกลับไปทบทวนใช้ [ตอนก่อนหน้า](../ep26-project-validation/README.md)

## แนวคิดและคำใหม่

เรามีแบบฟอร์มหลายใบ ถ้าใบหนึ่งผิด ให้บันทึกปัญหาของใบนั้นแล้วตรวจใบถัดไป รายงานยังรักษาลำดับข้อมูลเข้า

continue ข้ามส่วนที่เหลือของรอบปัจจุบันแล้วไปตรวจรอบต่อไป; fmt.Sprintf จัดรูปแบบแล้วคืน string โดยไม่พิมพ์; buildReport คืน slice ของบรรทัดรายงาน; _ ทิ้ง index ที่ไม่ได้ใช้

## หยุดคาดเดาก่อนเปิดผล

หลังรายการ sensor-03 ผิด sensor-04 จะยังได้รายงานหรือไม่?

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

## รันและตรวจผล

เปิด PowerShell แล้วเลือกโฟลเดอร์บทนี้ (เครื่องอื่นให้เปลี่ยนส่วนหน้าของตำแหน่งตามที่วางหลักสูตร):

```powershell
Set-Location -LiteralPath 'C:\Users\kopes\Documents\Codex\2026-09-17\go-youtube-go-iot-edge-gateway\Go\lessons\phase-01\ep27-project-report'
go run ./examples
```

ผลที่คาดหวัง:

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
sensor-03: ERROR: temperature outside simulated range
sensor-04: 28.0 C [OK]
```

ถ้าผลต่าง ให้ตรวจตำแหน่งด้วย `Get-Location` ตรวจไฟล์ที่บันทึก และอ่าน error แรกก่อนแก้

## ทำไมโค้ดจึงทำงานแบบนี้

- buildReport รับ []Reading และสะสม []string ส่วน main รับผิดชอบพิมพ์ จึงทดสอบข้อความได้โดยตรง
- เมื่อ validate ไม่ผ่าน เพิ่มบรรทัด ERROR แล้ว continue เพื่อไม่สร้างบรรทัดสำเร็จของรายการเดียวกัน
- โค้ดทำทีละรายการตามลำดับ ไม่ใช่งานพร้อมกัน; นี่เป็นเพียงการจัดการข้อมูลจำลองในหน่วยความจำ

**แยกประเภท:** syntax และชนิดข้อมูลเป็นหลักของภาษา; การแยกหน้าที่ให้ส่วนเล็กอ่านง่ายเป็นแนวทางออกแบบ; ชื่อโฟลเดอร์และการใช้หนึ่ง module ต่อ EP เป็นทางเลือกของหลักสูตร เกณฑ์อุณหภูมิที่พบเป็นข้อกำหนดจำลอง

## ลองแก้ในพื้นที่ฝึก

เปิด [วิธีสร้างพื้นที่ฝึก](../../../docs/PRACTICE.md) แล้วสร้าง EP.27 ด้วย `./scripts/new-practice.ps1 -Episode 27` จากราก Go เพียงครั้งเดียว จากนั้นเปิด `practics/phase-01/ep27-project-report` แก้ไฟล์ใน `examples` แล้วรัน `go run ./examples` จากโฟลเดอร์ EP ที่คัดลอกแล้ว

พื้นที่ฝึกไม่รวมเฉลย และคำสั่งสร้างจะหยุดถ้าโฟลเดอร์มีอยู่แล้ว จึงไม่ทับงานที่เคยทำ

## ข้อผิดพลาดที่พบบ่อย

- ใช้ return ในกรณีรายการผิด: จะจบทั้งฟังก์ชันและไม่ทำรายการถัดไป
- ลืม continue: รายการเสียจะมีทั้ง ERROR และบรรทัดสถานะที่ไม่ควรมี

## แบบฝึกหัดและคำถามทบทวน

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ก่อนดูเฉลย คาดผลของแต่ละข้อแล้วเปลี่ยนเพียงจุดที่โจทย์ระบุ

คำถามตรวจความเข้าใจ: continue ต่างจาก return ใน buildReport อย่างไร?

[เฉลยและเหตุผล](solutions/README.md) แยกไว้เพื่อเปิดหลังลองเอง ถ้ายังไม่ผ่าน ให้จดสิ่งที่คาดกับผลจริงไว้ก่อนขอความช่วยเหลือ

## สิ่งที่ได้และขอบเขต

ต่อยอดสู่งานอ่านหลายอุปกรณ์ใน Phase 5 ได้ แต่ยังไม่มีการอ่านอุปกรณ์หรือทำงานพร้อมกันจริง

ใช้ข้อมูลจำลองภายในโปรแกรมเท่านั้น ยังไม่ได้อ่านอุปกรณ์ ส่ง MQTT เรียก API หรือเก็บฐานข้อมูล โปรเจกต์นี้ใช้ข้อความชื่อและตัวเลขปกติที่เขียนไว้ในโค้ด ยังไม่ครอบคลุม NaN ชื่อที่มีแต่ช่องว่าง หรือข้อมูลภายนอก

ผลตรวจจริงและส่วนที่ยังไม่ได้ทดสอบ: [tests/RESULTS.md](tests/RESULTS.md)

อ้างอิงหลัก: [เอกสาร Go ทางการ](https://go.dev/ref/spec#Continue_statements)

เมื่ออธิบายผลและทำแบบฝึกหัดได้ ค่อยเลือก [ตอนถัดไป / สารบัญ](../ep28-project-tests/README.md) ไม่มีข้อกำหนดให้เรียนหลายตอนในวันเดียว
