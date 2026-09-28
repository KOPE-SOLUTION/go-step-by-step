# EP.12 — จัดโค้ดเป็น package และ module

**เป้าหมาย:** แยกกฎสถานะออกจาก main และอธิบายว่า import หาโค้ดจากที่ใด

**ก่อนเริ่ม:** EP.5 และ EP.9 — ฟังก์ชันและชนิดข้อมูล

## ทำความเข้าใจ

**package** รวมไฟล์ Go ในโฟลเดอร์เดียวกันเพื่อทำหน้าที่ร่วมกัน ส่วน **module** รวมหนึ่งหรือหลาย package โดยมีไฟล์ `go.mod` ในโฟลเดอร์หลักของ module เพื่อระบุชื่อ module และเวอร์ชัน Go ขั้นต่ำ

ชื่อ module เป็นส่วนต้นของ import path ไม่จำเป็นต้องเป็น repository บนอินเทอร์เน็ต ตัวอย่างใช้ `example.com` เป็นส่วนหนึ่งของชื่อสำหรับฝึก

## ลงมือทำ

1. ในโฟลเดอร์ `practics` เดิม ให้รัน `go mod init example.com/go-practice` **ครั้งเดียว** ถ้ามี `go.mod` แล้วให้อ่านชื่อ module เดิมก่อน ไม่ต้อง init ซ้ำ
2. สร้างโฟลเดอร์ `sensor` ข้าง `main.go` แล้วสร้างไฟล์ `sensor/reading.go` ใส่โค้ดของ package `sensor` ด้านล่าง
3. เปลี่ยน `main.go` เป็นโค้ดตัวอย่าง แล้วแก้ import ของ `sensor` ให้ใช้ชื่อ module ใน `practics/go.mod` ตามด้วย `/sensor` เช่น `example.com/go-practice/sensor` จากนั้นรัน `go run .` ที่ `practics`

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

**ถ้า `practics/go.mod` ระบุ `module example.com/go-practice` ให้ใช้ import `example.com/go-practice/sensor`** ส่วนตัวอย่างด้านบนใช้ชื่อ module ของหลักสูตร จึงมี import path ต่างกัน

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep12-packages-modules` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาสถานะของอุณหภูมิ 30 และ 28 เมื่อใช้เกณฑ์ 30 การแยกฟังก์ชันไว้ใน package `sensor` จะเปลี่ยนผลการตรวจหรือไม่

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: WARNING
sensor-02: OK
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `main.go` อยู่ใน package `main` ส่วน `sensor/reading.go` อยู่ใน package `sensor` โดยแยกคนละโฟลเดอร์
- ชื่อ `Status` ขึ้นต้นตัวใหญ่จึงเรียกจาก package อื่นได้ เรียกว่า **exported** ถ้าใช้ชื่อ `status` จะเรียกได้เฉพาะภายใน package `sensor`
- ตัวอย่างอ้างอิงใช้ `lessons/phase-01/go.mod` ร่วมกัน จึงมี import path ยาวกว่างานฝึก หลักคือชื่อ module ตามด้วยตำแหน่งโฟลเดอร์ package
- `go run .` ใช้ไฟล์ Go สำหรับโปรแกรมใน package ของโฟลเดอร์ปัจจุบัน ส่วน `go run main.go` ใช้เฉพาะไฟล์ `main.go` เป็นโค้ดของ package หลัก ทั้งสองคำสั่งยังเรียกใช้ package ที่ import ได้
- ใช้ module เดียวสำหรับ Phase 1 เพราะตัวอย่างทั้งหมดใช้ package ที่มากับ Go หรือ **standard library** ยังไม่มี package ภายนอกที่ต้องแยกจัดการเวอร์ชัน ชื่อ module ไม่ได้เปลี่ยน git remote หรือ push โค้ด

**ข้อผิดพลาดที่พบบ่อย:** ถ้าหา package ไม่พบ ให้ตรวจว่า import เริ่มด้วยชื่อ module ใน `go.mod` แล้วตามด้วยโฟลเดอร์ package แยกไฟล์ของ package `main` และ `sensor` ไว้คนละโฟลเดอร์ หากมี `go.mod` เดิม ให้ใช้ชื่อ module เดิมโดยไม่ต้อง init ซ้ำ

</details>

## ฝึกเอง

ฝึกใน `practics` เดิม ก่อนทำแต่ละข้อให้เริ่มจากตัวอย่าง `main.go` และ `sensor/reading.go` ของบทนี้ โดยแก้ import ให้ตรงกับ module งานฝึก หากต้องการเก็บคำตอบก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่มฟังก์ชัน `ToFahrenheit(celsius float64) float64` ใน `sensor/reading.go` แล้วให้ `main` เรียกฟังก์ชันเพื่อแปลง 30 C เป็น Fahrenheit และพิมพ์ต่อท้ายรายงาน
2. ใน `main` เปลี่ยนเกณฑ์ที่ส่งให้ `sensor.Status` เป็น 35 ทั้งสองรายการ โดยไม่แก้ package `sensor` แล้วคาดเดาสถานะใหม่

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้าเปลี่ยนโฟลเดอร์ที่เก็บ repository บนเครื่อง ต้องเปลี่ยน module path ทุกครั้งหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ ชื่อ module ใช้ระบุ package ใน import ไม่ใช่เส้นทางโฟลเดอร์เต็มบนเครื่อง การย้ายโฟลเดอร์ที่เก็บ repository จึงไม่ทำให้ต้องเปลี่ยนชื่อ module

</details>

**นำไปใช้ต่อ:** จัดกฎที่ใช้ร่วมกันก่อนโปรแกรมโตเป็น API หรือบริการหลายส่วน

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/create-module) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep12)

</details>

[EP.11](../ep11-interfaces/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.13](../ep13-unit-tests/README.md)
