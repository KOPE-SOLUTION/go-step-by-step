# EP.12 — จัดโค้ดเป็น package และ module

**เป้าหมาย:** แยกกฎสถานะออกจาก main และอธิบายว่า import หาโค้ดจากที่ใด

**ก่อนเริ่ม:** EP.5 และ EP.9 — ฟังก์ชันและชนิดข้อมูล

## ทำความเข้าใจ

**package** รวมไฟล์ Go ในโฟลเดอร์เดียวกันเพื่อทำหน้าที่ร่วมกัน ส่วน **module** รวมหนึ่งหรือหลาย package โดยมี `go.mod` ที่ราก ระบุชื่อ module และเวอร์ชัน Go ขั้นต่ำ

ชื่อ module เป็นส่วนต้นของ import path ไม่จำเป็นต้องเป็น repository บนอินเทอร์เน็ต ตัวอย่างใช้ example.com เป็นชื่อสำหรับฝึก

## ลงมือทำ

1. ใน practics เดิม ให้รัน `go mod init example.com/go-practice` **ครั้งเดียว** ถ้ามี go.mod แล้วให้อ่านชื่อ module เดิมก่อน ไม่ต้อง init ซ้ำ
2. สร้างโฟลเดอร์ `sensor` ข้าง main.go และไฟล์ `sensor/reading.go` ใส่โค้ด sensor ด้านล่าง
3. เปลี่ยน main.go เป็นโค้ดตัวอย่าง แต่ใช้ import `example.com/go-practice/sensor` ให้ตรง go.mod ของงานฝึก แล้วรัน `go run .` จาก practics

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import (
	"fmt"

	"example.com/go-course/phase01/ep12-packages-modules/sensor"
)

func main() {
	fmt.Println("sensor-01:", sensor.Status(30, 30))
	fmt.Println("sensor-02:", sensor.Status(28, 30))
}
```

ไฟล์ `sensor/reading.go`:

```go
package sensor

func Status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run .
```

**งานฝึกต้องใช้ import `example.com/go-practice/sensor`** ส่วนโค้ดอ้างอิงด้านบนใช้ชื่อ module ของหลักสูตร ดูขั้นตอนสร้างไฟล์ในหัวข้อลงมือทำ

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep12-packages-modules` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: WARNING
sensor-02: OK
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- main อยู่ package main ส่วนไฟล์ sensor/reading.go อยู่ package sensor คนละโฟลเดอร์จึงเป็นคนละ package
- ชื่อ `Status` ขึ้นต้นตัวใหญ่จึงเรียกจาก package อื่นได้ เรียกว่า **exported** ถ้าใช้ status จะใช้ได้ภายใน package sensor เท่านั้น
- ตัวอย่างอ้างอิงใช้ go.mod ร่วมกันที่ phase-01 จึงมี import path ยาวกว่างานฝึก หลักคือชื่อ module ตามด้วยตำแหน่งโฟลเดอร์ package
- `go run .` เลือกไฟล์โปรแกรมใน package ปัจจุบัน ส่วน `go run main.go` ระบุเฉพาะไฟล์ที่ส่งไป
- ใช้ module เดียวสำหรับ Phase 1 เพราะตัวอย่างทั้งหมดใช้ standard library ไม่มี dependency ภายนอกที่ต้องแยกเวอร์ชัน ชื่อ module ไม่ได้เปลี่ยน git remote หรือ push โค้ด

**ลองตรวจเมื่อผิด:** ถ้าหา package ไม่พบ ให้ตรวจชื่อ module กับ import ให้ตรงกัน อย่าเอา package main กับ package sensor ไว้โฟลเดอร์เดียวกัน ถ้ามี go.mod อยู่แล้วต้องใช้ชื่อนั้น ไม่แก้ชื่อโดยไม่ตรวจ import ที่เกี่ยวข้อง

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เพิ่มฟังก์ชัน `ToFahrenheit(celsius float64) float64` ใน package sensor แล้วให้ main พิมพ์ค่า 30 C เป็น Fahrenheit ต่อท้าย
2. ให้ main ใช้เกณฑ์ 35 ทั้งสองรายการ โดยไม่แก้ package sensor แล้วคาดเดาสถานะใหม่

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้าเปลี่ยนโฟลเดอร์ที่เก็บ repository บนเครื่อง ต้องเปลี่ยน module path ทุกครั้งหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ module path เป็นชื่อที่ใช้กับระบบ import ไม่ใช่ absolute path บนเครื่อง การย้ายโฟลเดอร์อย่างเดียวไม่เปลี่ยนชื่อนี้

</details>

**นำไปใช้ต่อ:** จัดกฎที่ใช้ร่วมกันก่อนโปรแกรมโตเป็น API หรือบริการหลายส่วน

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/create-module) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep12)

</details>

[EP.11](../ep11-interfaces/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.13](../ep13-unit-tests/README.md)
