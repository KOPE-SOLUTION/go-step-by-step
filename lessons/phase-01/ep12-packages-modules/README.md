# EP.12 — จัดโค้ดเป็น package และ module

**เป้าหมาย:** แยกกฎสถานะออกจาก main และอธิบายว่า import หาโค้ดจากที่ใด

**ก่อนเริ่ม:** EP.5 และ EP.9 — ฟังก์ชันและชนิดข้อมูล

## ทำความเข้าใจ

**package** รวมไฟล์ Go ในโฟลเดอร์เดียวกันเพื่อทำหน้าที่ร่วมกัน ส่วน **module** รวมหนึ่งหรือหลาย package โดยมีไฟล์ `go.mod` ในโฟลเดอร์หลักของ module เพื่อระบุชื่อ module และเวอร์ชัน Go ขั้นต่ำ

ชื่อ module เป็นส่วนต้นของ import path ไม่จำเป็นต้องเป็น repository บนอินเทอร์เน็ต ตัวอย่างใช้ `example.com` เป็นส่วนหนึ่งของชื่อสำหรับฝึก

## ลงมือทำทีละขั้น

ฝึกต่อใน `practics` เดิม ใช้ `main.go` และค่อยเพิ่มไฟล์ตามแต่ละขั้น หากมีไฟล์ที่จะเปลี่ยนและต้องการเก็บงานเดิม ให้คัดลอกเป็น `.txt` ก่อน บทนี้จะเปลี่ยนคำสั่งรันเป็น `go run .` หลังเตรียม module

### 1. เริ่มจากฟังก์ชันในไฟล์เดียว

เริ่ม `practics/main.go` ด้วยโค้ดนี้ แล้วรัน `go run main.go` จาก `practics`:

```go
package main

import "fmt"

func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func main() {
	fmt.Println("sensor-01:", status(30, 30))
}
```

ตรวจว่ากฎสถานะทำงานก่อนแยกโฟลเดอร์ จะได้เปรียบเทียบผลหลังย้ายฟังก์ชัน

**ลองคิดก่อนรัน:** ค่าเท่ากับเกณฑ์ควรเป็นสถานะอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: WARNING
```

</details>

### 2. เตรียม module ให้รู้จัก package ในโครงการ

เปิด terminal ที่ `practics` ถ้ายังไม่มี `go.mod` ให้รันครั้งเดียว:

```shell
go mod init example.com/go-practice
```

ถ้ามี `go.mod` อยู่แล้ว ให้ใช้ชื่อหลังคำว่า `module` ในไฟล์นั้น ไม่ต้อง init ซ้ำ จากนั้นรัน:

```shell
go run .
```

`go mod init` สร้าง `go.mod` ให้ ส่วนบรรทัด `go` จะอิงเวอร์ชันที่ติดตั้ง ชื่อ module เป็นส่วนต้นของ import ในขั้นถัดไป การเตรียม module ยังไม่เปลี่ยนกฎสถานะ

**ลองคิดก่อนรัน:** ก่อนแยก package ผลควรต่างจากขั้นแรกหรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: WARNING
```

</details>

### 3. ย้ายกฎไป package sensor แล้วเรียกจาก main

สร้างโฟลเดอร์ `practics/sensor` แล้วสร้าง `reading.go` ข้างใน ใส่โค้ดนี้:

```go
package sensor

func Status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}
```

แทน `practics/main.go` ทั้งไฟล์ด้วยโค้ดนี้ ถ้า module เดิมชื่ออื่น ให้เปลี่ยนส่วน `example.com/go-practice` ให้ตรงกับ `go.mod`:

```go
package main

import (
	"fmt"

	"example.com/go-practice/sensor"
)

func main() {
	fmt.Println("sensor-01:", sensor.Status(30, 30))
}
```

เมื่อย้ายแล้ว ฟังก์ชันอยู่ใน `package sensor` และเปลี่ยนชื่อเป็น `Status` ตัวใหญ่เพื่อให้เรียกข้าม package ได้ เรียกว่า exported ใช้ `sensor.Status(...)` แล้วรัน `go run .` จาก `practics`

**ลองคิดก่อนรัน:** การย้ายไฟล์เปลี่ยนผลของกฎเดิมหรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: WARNING
```

</details>

### 4. เรียกกฎเดียวกันกับอุปกรณ์อีกตัว

แทนเฉพาะ `main` เพื่อเพิ่มรายการที่สอง เก็บ import ของ module งานฝึกไว้:

```go
func main() {
	fmt.Println("sensor-01:", sensor.Status(30, 30))
	fmt.Println("sensor-02:", sensor.Status(28, 30))
}
```

รัน `go run .` จาก `practics` อีกครั้ง ทั้งสองรายการใช้ฟังก์ชันเดียวกันจาก package `sensor`

**ลองคิดก่อนรัน:** sensor-02 ที่อุณหภูมิ 28 ควรมีสถานะอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: WARNING
sensor-02: OK
```

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ `practics/main.go` สำหรับ module `example.com/go-practice`:

```go
package main

import (
	"fmt"

	"example.com/go-practice/sensor"
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

ตัวอย่างอ้างอิงใน [main.go](main.go) ใช้ module ของหลักสูตร จึงมี import path ต่างจากงานฝึก แต่เรียกกฎเดียวกัน

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run .
```

**ชื่อส่วนต้นของ import ต้องตรงกับ `module` ใน `practics/go.mod`** เช่น `example.com/go-practice/sensor`

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
