# EP.8 — ค้นหาและอัปเดตข้อมูลด้วย map

**เป้าหมาย:** เก็บค่าล่าสุดตามชื่ออุปกรณ์ และแยกค่าศูนย์จริงออกจากชื่อที่ไม่พบได้

**ก่อนเริ่ม:** EP.5 และ EP.7 — ฟังก์ชันและรายการข้อมูล

## ทำความเข้าใจ

**map** เก็บข้อมูลเป็นคู่ โดย **key** คือค่าที่ใช้ค้นหา และ **value** คือค่าที่เก็บไว้ เช่น ใช้ชื่อ `sensor-01` เป็น key เพื่อค้นหาอุณหภูมิล่าสุด โดยไม่ต้องรู้ตำแหน่งแบบ slice

## ลงมือทำทีละขั้น

ใช้ `practics/main.go` ไฟล์เดิม เริ่มด้วยโค้ดขั้นที่ 1 แล้วแก้ต่อทีละขั้น บันทึกและรัน `go run main.go` จาก terminal ที่ `practics` ทุกครั้ง ก่อนดูผล ให้ลองคาดเดาสิ่งที่จะพิมพ์

### 1. สร้าง map และค้นด้วยชื่อ

เริ่ม `main.go` ด้วย map ของสองอุปกรณ์:

```go
package main

import "fmt"

func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	fmt.Println("sensor-01:", latest["sensor-01"])
	fmt.Println("devices:", len(latest))
}
```

`map[string]float64` ใช้ชื่อชนิด `string` เป็น key เพื่อค้นค่าชนิด `float64` และ `len(latest)` บอกจำนวน key

**ลองคิดก่อนรัน:** ข้อมูลชุดนี้มีทั้งหมดกี่ชื่อ?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 27.5
devices: 2
```

</details>

### 2. อัปเดตชื่อเดิมและเพิ่มชื่อใหม่

แทน `main` เพื่อเพิ่มคำสั่งกำหนดค่าตาม key:

```go
func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	latest["sensor-01"] = 29
	latest["sensor-03"] = 31
	fmt.Println("sensor-01:", latest["sensor-01"])
	fmt.Println("sensor-03:", latest["sensor-03"])
	fmt.Println("devices:", len(latest))
}
```

กำหนดค่าให้ key เดิมจะอัปเดตค่านั้น ส่วน key ใหม่จะเพิ่มรายการ ไม่ต้องใช้ `append` แบบ slice

**ลองคิดก่อนรัน:** หลังสองคำสั่งนี้จะมี 3 หรือ 4 อุปกรณ์?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 29
sensor-03: 31
devices: 3
```

</details>

### 3. แยกค่าศูนย์จริงออกจากชื่อที่ไม่พบ

แทน `main` เพื่อรับทั้งค่าและผลว่าพบ key หรือไม่:

```go
func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	latest["sensor-01"] = 29
	latest["sensor-03"] = 31
	value, ok := latest["sensor-02"]
	fmt.Printf("sensor-02: value=%.1f found=%t\n", value, ok)
	value, ok = latest["sensor-04"]
	fmt.Printf("sensor-04: value=%.1f found=%t\n", value, ok)
}
```

`ok` เป็น `bool` ที่เป็น `true` เมื่อพบ key การค้นครั้งที่สองใช้ `=` เพราะประกาศ `value` กับ `ok` แล้ว ถ้าไม่พบ key จะได้ zero value ซึ่งกรณีนี้คือ 0

**ลองคิดก่อนรัน:** สองชื่ออาจให้ value เท่ากัน แต่ ok ต่างกันได้หรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-02: value=0.0 found=true
sensor-04: value=0.0 found=false
```

</details>

### 4. แยกการแสดงผลและทดลองลบข้อมูล

เพิ่มฟังก์ชัน `show` เหนือ `main` เพื่อใช้ขั้นตอนค้นหาซ้ำ:

```go
func show(latest map[string]float64, deviceID string) {
	value, ok := latest[deviceID]
	if !ok {
		fmt.Println(deviceID + ": not found")
		return
	}
	fmt.Printf("%s: %.1f C\n", deviceID, value)
}
```

แทน `main` เพื่อเรียก `show` และลบ `sensor-03` ด้วย `delete`:

```go
func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	latest["sensor-01"] = 29
	latest["sensor-03"] = 31
	show(latest, "sensor-01")
	show(latest, "sensor-02")
	show(latest, "sensor-04")
	delete(latest, "sensor-03")
	fmt.Println("devices:", len(latest))
}
```

`!ok` หมายถึงไม่พบ key จึงพิมพ์ `not found` แล้ว `return` ส่วน `delete` ลบคู่ key กับ value ของชื่อนั้น

**ลองคิดก่อนรัน:** หลังเพิ่ม sensor-03 แล้วลบออก จะเหลือกี่อุปกรณ์?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: not found
devices: 2
```

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

func show(latest map[string]float64, deviceID string) {
	value, ok := latest[deviceID]
	if !ok {
		fmt.Println(deviceID + ": not found")
		return
	}
	fmt.Printf("%s: %.1f C\n", deviceID, value)
}

func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	latest["sensor-01"] = 29
	latest["sensor-03"] = 31
	show(latest, "sensor-01")
	show(latest, "sensor-02")
	show(latest, "sensor-04")
	delete(latest, "sensor-03")
	fmt.Println("devices:", len(latest))
}
```

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep08-maps` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าผลค้นหา `sensor-02` ที่มีค่า 0 จะต่างจาก `sensor-04` ที่ไม่มีใน map อย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: not found
devices: 2
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `map[string]float64` ระบุชนิด key และ value ต้องเตรียม map ก่อนเขียน เช่น กำหนดค่าเริ่มต้นใน `map[string]float64{...}` ตามตัวอย่าง หรือใช้ `make(map[string]float64)` สร้าง map ว่าง
- อ่าน key ที่ไม่มีจะได้ zero value ของชนิดนั้น จึงใช้ bool `ok` เพื่อแยกว่าไม่พบหรือมีค่าศูนย์จริง
- `delete` ลบคู่ข้อมูลของ key นั้น ถ้าไม่มี key ก็ไม่เกิด error
- การใช้ `range` วนอ่าน map ไม่รับประกันลำดับ จึงใช้การค้นตามชื่อในตัวอย่างเพื่อให้ผลสาธิตแน่นอน
- map นี้เก็บค่าล่าสุดเท่านั้น ถ้าต้องการประวัติหลายครั้งต่ออุปกรณ์ ต้องออกแบบให้เก็บรายการค่าการวัดของแต่ละอุปกรณ์เพิ่ม ซึ่งจะเรียนในบทต่อยอด

**ข้อผิดพลาดที่พบบ่อย:** `var latest map[string]float64` ยังเป็น nil map อ่านได้ แต่ถ้าเพิ่มหรือแก้สมาชิก โปรแกรมจะหยุดด้วยข้อผิดพลาดขณะรันที่เรียกว่า **panic** ให้สร้าง map ก่อนเขียนข้อมูล ส่วนการตรวจ `value == 0` แทน `ok` จะทำให้ค่าศูนย์จริงถูกเข้าใจผิดว่าไม่พบข้อมูล

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. ก่อนเรียก `show(latest, "sensor-04")` ให้เพิ่ม `latest["sensor-04"] = 25` แล้วตรวจจำนวนอุปกรณ์ที่เหลือหลังลบ `sensor-03`
2. หลังลบ `sensor-03` ด้วย `delete` ให้เรียก `show(latest, "sensor-03")` ต้องแสดง `not found`

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้า `sensor-02` มีค่า 0 จะตรวจว่ามีอุปกรณ์นี้จริงอย่างไร?

<details>
<summary>แนวคำตอบ</summary>

รับผลด้วย `value, ok := latest["sensor-02"]` แล้วตรวจ `ok` ซึ่งเป็น `true` เมื่อมี key นี้อยู่ แม้ `value` จะเป็น 0

</details>

**นำไปใช้ต่อ:** ค้นสถานะหรือค่าล่าสุดของอุปกรณ์ตามชื่อ

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Map_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep8)

</details>

[EP.7](../ep07-arrays-slices/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.9](../ep09-structs/README.md)
