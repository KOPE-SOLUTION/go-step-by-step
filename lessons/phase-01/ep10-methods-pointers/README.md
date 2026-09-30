# EP.10 — เพิ่ม method และแก้ต้นฉบับด้วย pointer

**เป้าหมาย:** เลือกได้ว่างานใดอ่านสำเนา และงานใดตั้งใจแก้ข้อมูลต้นฉบับ

**ก่อนเริ่ม:** EP.9 — struct และการคัดลอกค่า

## ทำความเข้าใจ

**method** คือฟังก์ชันที่ผูกกับชนิดข้อมูล ส่วน **receiver** คือค่าที่ method ทำงานด้วย เช่น `reading.Status()`

**pointer** เป็นค่าที่อ้างถึงตำแหน่งข้อมูล `&reading` ให้ค่า pointer ที่ชี้ไปยังตัวแปร `reading` ส่วน `*Reading` คือชนิดของ pointer ที่ชี้ไปยังข้อมูลชนิด `Reading`

## ลงมือทำทีละขั้น

ใช้ `practics/main.go` ไฟล์เดิม เริ่มด้วยโค้ดขั้นที่ 1 แล้วแก้ต่อทีละขั้น บันทึกและรัน `go run main.go` จาก terminal ที่ `practics` ทุกครั้ง ก่อนดูผล ให้ลองคาดเดาสิ่งที่จะพิมพ์

### 1. เปลี่ยนฟังก์ชัน status เป็น method

เริ่ม `main.go` ด้วย `Reading` และ method `Status`:

```go
package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func (r Reading) Status() string {
	if r.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 29}
	fmt.Printf("%.1f C [%s]\n", reading.Celsius, reading.Status())
}
```

`(r Reading)` คือ receiver ระบุว่า method ทำงานกับข้อมูลชนิดใด จึงเรียกผ่าน `reading.Status()` แทนการส่งให้ `status(reading)` แบบบทก่อน

**ลองคิดก่อนรัน:** 29 C จะได้สถานะอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
29.0 C [OK]
```

</details>

### 2. ลองแก้ค่าผ่านฟังก์ชันที่รับสำเนา

เพิ่มฟังก์ชัน `adjustCopy` เหนือ `main`:

```go
func adjustCopy(r Reading, offset float64) {
	r.Celsius += offset
}
```

แทน `main` เพื่อเรียกฟังก์ชันแล้วตรวจต้นฉบับ:

```go
func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 29}
	adjustCopy(reading, 2)
	fmt.Printf("after copy: %.1f C [%s]\n", reading.Celsius, reading.Status())
}
```

`adjustCopy` รับสำเนาของ `Reading` จึงเพิ่มอุณหภูมิให้เฉพาะสำเนาภายในฟังก์ชัน และไม่ได้คืนสำเนานั้นออกมา

**ลองคิดก่อนรัน:** reading ต้นฉบับจะเป็น 29 หรือ 31?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
after copy: 29.0 C [OK]
```

</details>

### 3. ใช้ pointer ชี้ไปยังต้นฉบับ

แทน `main` เพื่อเพิ่มการแก้ค่าผ่าน pointer:

```go
func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 29}
	adjustCopy(reading, 2)
	fmt.Printf("after copy: %.1f C [%s]\n", reading.Celsius, reading.Status())
	pointer := &reading
	pointer.Celsius += 2
	fmt.Printf("after pointer: %.1f C [%s]\n", reading.Celsius, reading.Status())
}
```

`&reading` ให้ pointer ที่ชี้ไปยังตัวแปร `reading` เมื่อแก้ `pointer.Celsius` จึงแก้ field ของข้อมูลต้นฉบับที่ pointer ชี้อยู่

**ลองคิดก่อนรัน:** บรรทัด after pointer จะต่างจาก after copy อย่างไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
after copy: 29.0 C [OK]
after pointer: 31.0 C [WARNING]
```

</details>

### 4. ย้ายการแก้ค่ามาไว้ใน method

เพิ่ม method `Adjust` เหนือ `main` โดยเก็บ `Status` และ `adjustCopy` ไว้:

```go
func (r *Reading) Adjust(offset float64) {
	r.Celsius += offset
}
```

ใน `main` แทน `pointer.Celsius += 2` ด้วยการเรียก method:

```go
pointer.Adjust(2)
```

`*Reading` คือชนิด pointer ที่ชี้ไปยัง `Reading` ส่วน `offset` คือจำนวนที่จะเพิ่มหรือลด method นี้ตั้งใจแก้ต้นฉบับ จึงใช้ pointer receiver

**ลองคิดก่อนรัน:** ผลลัพธ์ควรเหมือนขั้นก่อนหรือไม่? ลองอธิบายว่าต่างกันที่การจัดโค้ดหรือข้อมูลที่แก้

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
after copy: 29.0 C [OK]
after pointer: 31.0 C [WARNING]
```

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func (r Reading) Status() string {
	if r.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}

func (r *Reading) Adjust(offset float64) {
	r.Celsius += offset
}

func adjustCopy(r Reading, offset float64) {
	r.Celsius += offset
}

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 29}
	adjustCopy(reading, 2)
	fmt.Printf("after copy: %.1f C [%s]\n", reading.Celsius, reading.Status())
	pointer := &reading
	pointer.Adjust(2)
	fmt.Printf("after pointer: %.1f C [%s]\n", reading.Celsius, reading.Status())
}
```

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep10-methods-pointers` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาค่า `reading.Celsius` หลังเรียก `adjustCopy` และหลังเรียก `Adjust`

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
after copy: 29.0 C [OK]
after pointer: 31.0 C [WARNING]
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `(r Reading)` รับสำเนา เหมาะกับตัวอย่าง `Status` ที่อ่านข้อมูลอย่างเดียว
- `(r *Reading)` รับ pointer จึงแก้ `Celsius` ของต้นฉบับที่ชี้ถึงได้ ตัว pointer เองก็ถูกส่งเป็นสำเนาค่าเช่นกัน
- `*pointer` หมายถึงข้อมูลที่ pointer ชี้อยู่ เรียกว่า dereference แต่ Go ย่อการเข้าถึง field ผ่าน pointer ให้เขียน `r.Celsius` ได้
- ในตัวอย่างนี้ `reading` เป็นตัวแปรที่อ้างถึงตำแหน่งได้ Go จึงยอมให้เขียน `reading.Adjust(2)` แทน `(&reading).Adjust(2)`
- การตั้งชื่อ `Adjust` เป็นทางเลือกในการออกแบบ ส่วนความต่างระหว่างการรับสำเนา `Reading` กับ pointer เป็นหลักของภาษา ไม่ต้องใช้ pointer กับทุกค่า

**ข้อผิดพลาดที่พบบ่อย:** pointer ที่ไม่ได้ชี้ไปยังข้อมูลมีค่า `nil` หากเรียก `Adjust` ในตัวอย่างนี้ผ่าน nil pointer จะเกิดข้อผิดพลาดขณะรัน เพราะ method พยายามแก้ `Celsius` ตัวอย่างจึงสร้าง `reading` ก่อน แล้วใช้ `&reading` เพื่อให้ pointer ชี้ไปยังข้อมูลนั้น

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เปลี่ยนเฉพาะ `pointer.Adjust(2)` เป็น `pointer.Adjust(-1.5)` แล้วตรวจอุณหภูมิและสถานะของ `reading` หลังเรียก method
2. เปลี่ยน `adjustCopy` ให้รับ `*Reading` และเรียกด้วย `&reading` โดยคงการปรับทั้งสองครั้งไว้ แล้วคาดเดาผลใหม่

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** เมื่อส่ง pointer ให้ฟังก์ชันหรือ method Go ยังคัดลอกค่าที่ส่งอยู่หรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ยังคัดลอกค่า pointer อยู่ แต่ pointer ทั้งสองชี้ไปยังข้อมูลเดียวกัน จึงใช้สำเนาของ pointer แก้ข้อมูลต้นฉบับได้

</details>

**นำไปใช้ต่อ:** เข้าใจการเปลี่ยนสถานะของข้อมูล และเตรียมใช้ method ผ่าน interface

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Method_declarations) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep10)

</details>

[EP.9](../ep09-structs/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.11](../ep11-interfaces/README.md)
