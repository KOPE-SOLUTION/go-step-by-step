# EP.5 — แยกงานด้วยฟังก์ชัน รับค่าและคืนค่า

**เป้าหมาย:** แยกการแปลงหน่วยและการตัดสินสถานะออกจากการแสดงผล แล้วเรียกซ้ำได้

**ก่อนเริ่ม:** EP.3–4 — เงื่อนไขและลูป

## ทำความเข้าใจ

**ฟังก์ชัน** คือชุดคำสั่งที่ตั้งชื่อเพื่อเรียกใช้ **parameter** คือตัวแปรที่ประกาศไว้เพื่อรับค่าเมื่อเรียกฟังก์ชัน และ **return value** คือผลที่ส่งกลับไปยังผู้เรียก เช่น รับ Celsius แล้วคืน Fahrenheit

## ลงมือทำทีละขั้น

ใช้ `practics/main.go` ไฟล์เดิม เริ่มด้วยโค้ดขั้นที่ 1 แล้วแก้ต่อทีละขั้น บันทึกและรัน `go run main.go` จาก terminal ที่ `practics` ทุกครั้ง ก่อนดูผล ให้ลองคาดเดาสิ่งที่จะพิมพ์

### 1. ทบทวนสูตรที่เขียนใน main

เริ่ม `main.go` ด้วยสูตรจาก EP.3:

```go
package main

import "fmt"

func main() {
	celsius := 30.0
	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.1f C = %.1f F\n", celsius, fahrenheit)
}
```

ตอนนี้การคำนวณกับการแสดงผลอยู่ใน `main` ถ้าต้องการใช้สูตรกับหลายค่า เราจะแยกสูตรเป็นฟังก์ชันในขั้นถัดไป

**ลองคิดก่อนรัน:** 30 C ควรได้กี่ F?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
30.0 C = 86.0 F
```

</details>

### 2. แยกฟังก์ชันรับค่าและคืนค่า

เพิ่มฟังก์ชันนี้เหนือ `func main()` ไม่ใช่ภายใน `main`:

```go
func toFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}
```

แล้วแทน `main` เพื่อทดลองเรียกสองครั้ง:

```go
func main() {
	fmt.Println(toFahrenheit(0))
	fmt.Println(toFahrenheit(100))
}
```

`celsius float64` เป็น parameter ที่รับค่า ส่วน `float64` หลังวงเล็บคือชนิดผลที่คืน `return` ส่งผลกลับให้ผู้เรียก แล้วจบการทำงานของฟังก์ชันครั้งนั้น

**ลองคิดก่อนรัน:** เรียกด้วย 0 และ 100 จะได้ผลต่างกันอย่างไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
32
212
```

</details>

### 3. เพิ่มฟังก์ชันตรวจสถานะ

เพิ่ม `status` เหนือ `main` โดยเก็บ `toFahrenheit` ไว้:

```go
func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}
```

แทน `main` เพื่อทดลองเกณฑ์สองค่า:

```go
func main() {
	celsius := 30.0
	fmt.Println(status(celsius, 30))
	fmt.Println(status(celsius, 35))
}
```

`celsius, threshold float64` รับสองค่าชนิดเดียวกัน `status` คืนข้อความให้ผู้เรียกใช้ต่อ โดยยังไม่พิมพ์เอง

**ลองคิดก่อนรัน:** อุณหภูมิเดียวกัน แต่เกณฑ์ 30 กับ 35 จะได้สถานะเหมือนกันหรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
WARNING
OK
```

</details>

### 4. นำผลจากฟังก์ชันมาประกอบรายงาน

แทนเฉพาะ `main` ด้วยรายงานนี้ โดยเก็บฟังก์ชันทั้งสองไว้:

```go
func main() {
	celsius := 30.0
	fmt.Printf("%.1f C = %.1f F [%s]\n",
		celsius, toFahrenheit(celsius), status(celsius, 30))
	fmt.Println("at threshold 35:", status(celsius, 35))
}
```

เมื่อเรียกฟังก์ชันใน `Printf` Go จะคำนวณค่าที่คืนมาก่อนนำไปจัดรูปแบบ การแยกคำนวณออกจากการพิมพ์เป็นแนวทางออกแบบเพื่อให้ใช้ผลซ้ำได้

**ลองคิดก่อนรัน:** ส่วนใดคำนวณ ส่วนใดตัดสินสถานะ และส่วนใดพิมพ์ข้อความ?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
30.0 C = 86.0 F [WARNING]
at threshold 35: OK
```

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

func toFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func main() {
	celsius := 30.0
	fmt.Printf("%.1f C = %.1f F [%s]\n",
		celsius, toFahrenheit(celsius), status(celsius, 30))
	fmt.Println("at threshold 35:", status(celsius, 35))
}
```

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep05-functions` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าอุณหภูมิเดียวกันจะได้สถานะต่างกันอย่างไร เมื่อใช้เกณฑ์เตือน 30 และ 35

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
30.0 C = 86.0 F [WARNING]
at threshold 35: OK
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `func toFahrenheit(celsius float64) float64` ระบุชื่อฟังก์ชัน ตัวแปรรับค่า และชนิดของค่าที่คืนตามลำดับ
- `return` ส่งผลกลับและจบการเรียกฟังก์ชันครั้งนั้น ส่วน `Println` แสดงข้อความ จึงทำหน้าที่ต่างกัน
- `celsius, threshold float64` คือ parameter สองตัวชนิดเดียวกัน
- ชื่อภายในฟังก์ชันใช้ได้ในขอบเขตของมัน เรียกว่า **scope** ตัวแปร `celsius` ใน `main` กับ parameter ที่ชื่อเหมือนกันอยู่คนละขอบเขต
- การแยกคำนวณออกจากการพิมพ์เป็นแนวทางออกแบบ ช่วยให้ใช้กับรายงานหรือ API และทดสอบได้ง่ายขึ้น ไม่ใช่ข้อบังคับของภาษา

**ข้อผิดพลาดที่พบบ่อย:** ถ้า `status` คืนค่าเฉพาะกรณีเตือน แต่ไม่มี `return` สำหรับกรณีปกติ จะคอมไพล์ไม่ผ่าน ให้ตรวจว่าทั้งสองกรณีคืนค่าแล้ว นอกจากนี้ จำนวนและชนิดของค่าที่ส่งตอนเรียกต้องตรงกับที่ฟังก์ชันประกาศไว้

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่มฟังก์ชัน `toCelsius(fahrenheit float64) float64` ใช้สูตร `(fahrenheit - 32) * 5 / 9` แล้วพิมพ์ผลของ 86 F ต่อท้าย
2. เขียนลูป `for` ภายใน `main` เพื่อเรียก `status` ด้วยค่า 29, 30 และ 31 โดยใช้เกณฑ์ 30 แล้วพิมพ์สถานะของแต่ละค่าแทนรายงานเดิม

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้า `status` พิมพ์ `WARNING` เอง แต่ไม่คืนค่า จะนำข้อความนั้นไปประกอบรายงานได้สะดวกเหมือนเดิมหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ ถ้า `status` พิมพ์ข้อความเองอย่างเดียว ผู้เรียกจะไม่ได้รับข้อความกลับมาประกอบรายงาน การคืนค่าทำให้ผู้เรียกเลือกได้ว่าจะพิมพ์ เก็บ หรือใช้ข้อความต่อ

</details>

**นำไปใช้ต่อ:** แยกกฎที่ใช้ซ้ำและเตรียมทดสอบโดยไม่ต้องอ่านข้อความจาก terminal

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Function_declarations) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep5)

</details>

[EP.4](../ep04-loops/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.6](../ep06-errors/README.md)
