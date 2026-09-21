# EP.9 — สร้างชนิดข้อมูลการวัดด้วย struct

**เป้าหมาย:** รวมชื่อกับอุณหภูมิเป็นข้อมูลหนึ่งรายการ แล้วทำรายงานหลายรายการได้

**ก่อนเริ่ม:** EP.5 และ EP.7 — ฟังก์ชันกับ slice

## ทำความเข้าใจ

**struct** คือชนิดข้อมูลที่รวมหลายช่องไว้ด้วยกัน แต่ละช่องเรียกว่า **field** เราจะสร้าง Reading ให้หนึ่งรายการมีทั้ง DeviceID และ Celsius แทนการเก็บไว้ในตัวแปรที่ไม่ผูกกัน

## ลงมือทำ

1. สร้างชนิด Reading และข้อมูลหนึ่งรายการ ทดลองอ่านและแก้ `reading.Celsius`
2. เพิ่มเป็น `[]Reading` และส่งรายการเข้า status เพื่อแสดงรายงานตามตัวอย่าง
3. ทดลองคัดลอก struct แล้วเปลี่ยนเฉพาะสำเนา เดาก่อนว่าต้นฉบับจะเปลี่ยนด้วยหรือไม่

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

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

func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
	}
	for _, reading := range readings {
		fmt.Printf("%s: %.1f C [%s]\n",
			reading.DeviceID, reading.Celsius, status(reading))
	}
	original := readings[0]
	copyReading := original
	copyReading.Celsius = 99
	fmt.Printf("original: %.1f, copy: %.1f\n", original.Celsius, copyReading.Celsius)
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep09-structs` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
original: 27.5, copy: 99.0
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- `type Reading struct` ประกาศชนิดที่ตั้งชื่อเอง ไม่ได้สร้างข้อมูลจริงจนกว่าจะใช้ literal เช่น `Reading{...}`
- ระบุชื่อ field ใน literal ช่วยให้เห็นว่าแต่ละค่าเป็นอะไร และไม่ต้องจำลำดับช่อง
- `[]Reading` คือ slice ที่แต่ละสมาชิกเป็น Reading ทำให้หนึ่งรายการเดินทางไปพร้อมชื่อและอุณหภูมิ
- struct ถูกคัดลอกเมื่อกำหนดค่าและส่งให้ฟังก์ชัน ตัวอย่างที่มี string กับ float64 จึงแก้สำเนาโดยไม่เปลี่ยนต้นฉบับ
- ถ้า struct มี field แบบ slice หรือ map การคัดลอก struct ไม่ได้คัดลอกข้อมูลเบื้องหลังทั้งหมด เรื่องนี้จะสำคัญเมื่อต้องแก้ข้อมูลร่วมกัน

**ลองตรวจเมื่อผิด:** ชื่อ field แยกตัวพิมพ์เล็กใหญ่ และ `reading` คือค่าที่สร้างขึ้น ส่วน `Reading` คือชื่อชนิด ถ้าเปลี่ยน reading ภายใน range จะเป็นการแก้สำเนา ให้แก้ผ่าน index เมื่อต้องการแก้สมาชิกจริง

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เพิ่ม sensor-03 ค่า 31.5 ใน slice เพื่อให้รายงานแสดงสามอุปกรณ์
2. ก่อน range เปลี่ยน Celsius ของสมาชิกแรกเป็น 32 ผ่าน `readings[0]` แล้วตรวจทั้งรายงานและค่าต้นฉบับตอนท้าย

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้าแก้ copyReading.Celsius ทำไม original.Celsius จึงไม่เปลี่ยน?

<details>
<summary>แนวคำตอบ</summary>

การกำหนด struct คัดลอกค่า field และ float64 ในตัวอย่างเป็นค่าคนละสำเนา

</details>

**นำไปใช้ต่อ:** เป็นรูปข้อมูลเดียวที่นำไปใช้กับ JSON, API และการจัดเก็บในเฟสถัดไป

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Struct_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep9)

</details>

[EP.8](../ep08-maps/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.10](../ep10-methods-pointers/README.md)
