# EP.9 — สร้างชนิดข้อมูลการวัดด้วย struct

**เป้าหมาย:** รวมชื่อกับอุณหภูมิเป็นข้อมูลหนึ่งรายการ แล้วทำรายงานหลายรายการได้

**ก่อนเริ่ม:** EP.5 และ EP.7 — ฟังก์ชันกับ slice

## ทำความเข้าใจ

**struct** คือชนิดข้อมูลที่รวมหลายช่องไว้ด้วยกัน แต่ละช่องเรียกว่า **field** เราจะสร้างชนิด `Reading` เพื่อเก็บชื่ออุปกรณ์ `DeviceID` และอุณหภูมิ `Celsius` ไว้ด้วยกันในข้อมูลหนึ่งรายการ

## ลงมือทำ

1. สร้างชนิด `Reading` และตัวแปร `reading` สำหรับเก็บข้อมูลหนึ่งรายการ แล้วลองอ่านและแก้ `reading.Celsius`
2. เพิ่มเป็น `[]Reading` และส่งแต่ละรายการให้ `status` เพื่อแสดงรายงานตามตัวอย่าง
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

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าหลังแก้ `copyReading.Celsius` แล้ว ค่าใน `original.Celsius` จะเปลี่ยนด้วยหรือไม่

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 27.5 C [OK]
sensor-02: 30.0 C [WARNING]
original: 27.5, copy: 99.0
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `type Reading struct` ประกาศชนิดข้อมูลชื่อ `Reading` ส่วน `Reading{...}` ใช้สร้างค่าของชนิดนั้น
- การระบุชื่อ field เมื่อสร้างค่า เช่น `Reading{DeviceID: "sensor-01", Celsius: 27.5}` ช่วยให้รู้ว่าแต่ละค่าเก็บในช่องใด โดยไม่ต้องจำลำดับ field
- `[]Reading` คือ slice ที่สมาชิกแต่ละตัวเป็น `Reading` จึงเก็บทั้งชื่ออุปกรณ์และอุณหภูมิไว้ในรายการเดียวกัน
- struct ถูกคัดลอกเมื่อกำหนดค่าและส่งให้ฟังก์ชัน ตัวอย่างที่มี string กับ float64 จึงแก้สำเนาโดยไม่เปลี่ยนต้นฉบับ
- ถ้า struct มี field แบบ slice หรือ map การคัดลอก struct ไม่ได้คัดลอกข้อมูลเบื้องหลังทั้งหมด เรื่องนี้จะสำคัญเมื่อต้องแก้ข้อมูลร่วมกัน

**ข้อผิดพลาดที่พบบ่อย:** ชื่อ field แยกตัวพิมพ์เล็กใหญ่ และ `reading` คือค่าที่สร้างขึ้น ส่วน `Reading` คือชื่อชนิด ถ้าแก้ `reading.Celsius` ภายใน `range` จะเปลี่ยนเฉพาะสำเนา ถ้าต้องการแก้สมาชิกใน slice ให้ใช้ `readings[index].Celsius`

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่ม `Reading` ของ `sensor-03` อุณหภูมิ 31.5 ลงใน slice `readings` เพื่อให้รายงานแสดงสามอุปกรณ์
2. ก่อนลูป `range` ให้เปลี่ยน `Celsius` ของสมาชิกแรกเป็น 32 ผ่าน `readings[0]` แล้วตรวจอุณหภูมิในรายงานและค่า `original` ที่พิมพ์ตอนท้าย

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** เมื่อแก้ `copyReading.Celsius` ทำไม `original.Celsius` จึงไม่เปลี่ยน?

<details>
<summary>แนวคำตอบ</summary>

`copyReading := original` คัดลอกค่าของทุก field มาเก็บใน struct อีกตัว จึงแก้ field ชนิด `float64` ในสำเนาได้โดยไม่เปลี่ยนต้นฉบับ

</details>

**นำไปใช้ต่อ:** ใช้ struct กำหนดข้อมูลการวัดที่จะนำไปแปลงเป็น JSON ส่งผ่าน API หรือบันทึกในเฟสถัดไป

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Struct_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep9)

</details>

[EP.8](../ep08-maps/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.10](../ep10-methods-pointers/README.md)
