# EP.27 — โปรเจกต์สรุป ตอนที่ 3 — รายงานหลายรายการ

Phase 1 — Go Basic • [สารบัญ](../../../docs/playlist-01-go-basic/README.md)

## เป้าหมายและสิ่งที่ควรรู้ก่อน

ให้รายการผิดหนึ่งรายการไม่หยุดการประมวลผลรายการถัดไป ก่อนเริ่มควรเรียน EP.26 และอธิบายตัวอย่างในตอนนั้นด้วยคำของตนเองได้ สามารถเปิดตัวอย่างประกอบระหว่างฝึกได้

เรียนตามลำดับ: ทำความเข้าใจ → คาดเดาผล → รันโปรแกรม → ลองแก้ทีละจุด หากต้องการทบทวน เปิด [ตอนก่อนหน้า](../ep26-project-validation/README.md)

## แนวคิดและคำใหม่

เรามีแบบฟอร์มหลายใบ ถ้าใบหนึ่งผิด ให้บันทึกปัญหาของใบนั้นแล้วตรวจใบถัดไป รายงานยังรักษาลำดับข้อมูลเข้า

continue ข้ามส่วนที่เหลือของรอบปัจจุบันแล้วไปตรวจรอบต่อไป; fmt.Sprintf จัดรูปแบบแล้วคืน string โดยไม่พิมพ์; buildReport คืน slice ของบรรทัดรายงาน; _ ทิ้ง index ที่ไม่ได้ใช้

## ลองคาดเดาผลลัพธ์ก่อนรัน

หลังรายการ sensor-03 ผิด sensor-04 จะยังได้รายงานหรือไม่?

จดผลที่คาดเดาและเหตุผลสั้น ๆ ก่อนรัน แล้วเปรียบเทียบกับผลจริง หากต่างกัน ให้ลองอธิบายว่าเกิดจากส่วนใดของโค้ด

## โค้ดตัวอย่าง

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

เปิด PowerShell แล้วเลือกโฟลเดอร์บทเรียน หากวางหลักสูตรไว้ที่ตำแหน่งอื่น ให้ปรับเส้นทางในคำสั่งให้ตรงกับเครื่องของคุณ:

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

หากผลลัพธ์ไม่ตรงกับที่คาด ให้ตรวจตำแหน่งด้วย `Get-Location` และตรวจว่าได้บันทึกไฟล์แล้ว หากมีข้อผิดพลาด ให้อ่านข้อความแรกก่อนแก้โค้ด

## อธิบายการทำงานของโค้ด

- buildReport รับ []Reading และสะสม []string ส่วน main รับผิดชอบพิมพ์ จึงทดสอบข้อความได้โดยตรง
- เมื่อ validate ไม่ผ่าน เพิ่มบรรทัด ERROR แล้ว continue เพื่อไม่สร้างบรรทัดสำเร็จของรายการเดียวกัน
- โค้ดทำทีละรายการตามลำดับ ไม่ใช่งานพร้อมกัน; นี่เป็นเพียงการจัดการข้อมูลจำลองในหน่วยความจำ

**หลักภาษาและแนวทางออกแบบ:** รูปแบบคำสั่งและชนิดข้อมูลเป็นหลักของภาษา Go การแยกโค้ดตามหน้าที่เป็นแนวทางออกแบบให้เข้าใจง่าย ส่วนชื่อโฟลเดอร์และการใช้หนึ่ง module ต่อ EP เป็นทางเลือกในการจัดหลักสูตร

## ฝึกแก้ไขโค้ดด้วยตนเอง

เปิด [วิธีสร้างพื้นที่ฝึก](../../../docs/PRACTICE.md) แล้วสร้างพื้นที่ฝึกสำหรับ EP.27 ด้วย `./scripts/new-practice.ps1 -Episode 27` จากราก Go เพียงครั้งเดียว จากนั้นเปิด `practics/phase-01/ep27-project-report` แก้ไฟล์ใน `examples` แล้วรัน `go run ./examples` จากโฟลเดอร์ EP ที่คัดลอกแล้ว

พื้นที่ฝึกไม่รวมเฉลย และคำสั่งสร้างจะหยุดถ้าโฟลเดอร์มีอยู่แล้ว จึงไม่ทับงานที่เคยทำ

## ข้อผิดพลาดที่พบบ่อย

- ใช้ return ในกรณีรายการผิด: จะจบทั้งฟังก์ชันและไม่ทำรายการถัดไป
- ลืม continue: รายการเสียจะมีทั้ง ERROR และบรรทัดสถานะที่ไม่ควรมี

## แบบฝึกหัดและคำถามทบทวน

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ก่อนดูเฉลย คาดเดาผลลัพธ์ของแต่ละข้อ แล้วแก้โค้ดตามที่โจทย์กำหนด

คำถามตรวจความเข้าใจ: continue ต่างจาก return ใน buildReport อย่างไร?

เปิด [เฉลยและคำอธิบาย](solutions/README.md) หลังทำแบบฝึกหัด หากยังติดปัญหา ให้จดผลที่คาด ผลจริง และข้อความผิดพลาดไว้ประกอบการตรวจสอบ

## สรุปบทเรียนและขอบเขตของตัวอย่าง

ต่อยอดสู่งานอ่านหลายอุปกรณ์ใน Phase 5 ได้ แต่ยังไม่มีการอ่านอุปกรณ์หรือทำงานพร้อมกันจริง

ใช้ข้อมูลจำลองภายในโปรแกรมเท่านั้น ยังไม่ได้อ่านอุปกรณ์ ส่ง MQTT เรียก API หรือเก็บฐานข้อมูล โปรเจกต์นี้ใช้ข้อความชื่อและตัวเลขปกติที่เขียนไว้ในโค้ด ยังไม่ครอบคลุม NaN ชื่อที่มีแต่ช่องว่าง หรือข้อมูลภายนอก

ผลตรวจจริงและส่วนที่ยังไม่ได้ทดสอบ: [tests/RESULTS.md](tests/RESULTS.md)

อ้างอิงหลัก: [เอกสาร Go ทางการ](https://go.dev/ref/spec#Continue_statements)

เมื่อทำแบบฝึกหัดและอธิบายผลได้แล้ว จึงเปิด [ตอนถัดไป / สารบัญ](../ep28-project-tests/README.md) สามารถทบทวนบทนี้ได้ตามต้องการก่อนเรียนต่อ
