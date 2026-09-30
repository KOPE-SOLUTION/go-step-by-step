# EP.13 — ทดสอบกฎและค่าขอบเขตด้วย unit test

**เป้าหมาย:** เขียน test ที่จับความผิดพลาดเมื่อค่าเท่ากับเกณฑ์ได้ และอ่านว่าผลจริงต่างจากที่ต้องการอย่างไร

**ก่อนเริ่ม:** EP.5, EP.7, EP.9 และ EP.12 — ฟังก์ชัน, slice, struct และ module

## ทำความเข้าใจ

**unit test** คือโค้ดตรวจพฤติกรรมของส่วนเล็ก ๆ เช่นฟังก์ชัน `status` เรากำหนดค่าที่ส่งให้ฟังก์ชันและผลที่ต้องการ แล้วให้เครื่องเปรียบเทียบแทนการรันดูเองทุกครั้ง

ในบทนี้ใช้กฎเดิม: ตั้งแต่เกณฑ์ขึ้นไปเป็น WARNING ก่อนเขียน test ให้ตอบผลของ 29.9, 30 และ 30.1 ที่เกณฑ์ 30

## ลงมือทำทีละขั้น

ใช้ `practics` และ `go.mod` เดิมจาก EP.12 บทนี้เริ่มจาก `main.go` แล้วค่อยสร้าง `main_test.go` ข้างกัน เก็บ `sensor/reading.go` เดิมไว้ได้เพราะบทนี้ไม่ได้ import package นั้น

### 1. ตรวจฟังก์ชันด้วยการรันโปรแกรมก่อน

แทน `practics/main.go` ด้วยกฎนี้ แล้วรัน `go run .` จาก `practics`:

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
	fmt.Println("29.9 C:", status(29.9, 30))
	fmt.Println("30.0 C:", status(30, 30))
}
```

เรายังอ่านผลเองอยู่ ขั้นต่อไปจะเขียนโค้ดทดสอบให้ตรวจผลแทน

**ลองคิดก่อนรัน:** 29.9 กับ 30 ที่เกณฑ์ 30 ควรได้สถานะอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
29.9 C: OK
30.0 C: WARNING
```

</details>

### 2. เขียน test เพียงกรณีเดียว

สร้าง `practics/main_test.go` แล้วใส่โค้ดนี้ หากมีไฟล์เดิมให้เก็บสำเนาก่อนแทน:

```go
package main

import "testing"

func TestStatusEqual(t *testing.T) {
	got := status(30, 30)
	want := "WARNING"
	if got != want {
		t.Errorf("status(30, 30) = %q; want %q", got, want)
	}
}
```

รันจาก terminal ที่ `practics`:

```shell
go test -v .
```

ไฟล์ลงท้าย `_test.go` และฟังก์ชัน `TestStatusEqual` รับ `*testing.T` เพื่อรายงานผล `got` คือผลจริง ส่วน `want` คือผลที่ต้องการ `t.Errorf` ทำให้ test ไม่ผ่านเมื่อสองค่าไม่ตรงกัน

**ลองคิดก่อนรัน:** กรณีเท่ากับเกณฑ์นี้ควรผ่านหรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ควรเห็น `PASS: TestStatusEqual` และ `PASS` ส่วนเวลาและชื่อ package อาจต่างกัน

</details>

### 3. จงใจทำกฎผิดเพื่อดูว่า test ตรวจพบไหม

ใน `main.go` เปลี่ยนเฉพาะบรรทัดเงื่อนไขของ `status` เป็น:

```go
if celsius > threshold {
```

บันทึกแล้วรัน test จาก `practics` อีกครั้ง:

```shell
go test -v .
```

นี่เป็นการทดลองให้ test ไม่ผ่าน ลองอ่านผลจริงเทียบกับผลที่ต้องการ ไม่แก้ `want` เพียงเพื่อให้ผ่าน

**ลองคิดก่อนรัน:** ค่า 30 เท่ากับเกณฑ์ แต่เงื่อนไขใหม่จะคืนอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ต้องเห็น `FAIL: TestStatusEqual` พร้อมข้อความ `status(30, 30) = "OK"; want "WARNING"`

</details>

### 4. แก้กฎกลับและตรวจซ้ำ

แก้บรรทัดเดิมใน `status` กลับเป็น:

```go
if celsius >= threshold {
```

บันทึกแล้วรัน test ซ้ำ:

```shell
go test -v .
```

เมื่อกฎยังเป็น “ตั้งแต่เกณฑ์ขึ้นไป” เราแก้ตัวโปรแกรมให้ตรงกฎ โดยคงผลที่ต้องการใน test ไว้

**ลองคิดก่อนรัน:** เหตุใดจึงควรเห็น test ผ่านหลังแก้?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ควรกลับมาเห็น `PASS: TestStatusEqual` และ `PASS`

</details>

### 5. เพิ่มหลายกรณีด้วย slice และลูป

แทน `main_test.go` ทั้งไฟล์ด้วยตารางข้อมูลทดสอบ แล้วรัน `go test -v .`:

```go
package main

import "testing"

func TestStatus(t *testing.T) {
	cases := []struct {
		name      string
		celsius   float64
		threshold float64
		want      string
	}{
		{"below", 29.9, 30, "OK"},
		{"equal", 30, 30, "WARNING"},
		{"above", 30.1, 30, "WARNING"},
		{"custom threshold", 30, 35, "OK"},
	}
	for _, tc := range cases {
		got := status(tc.celsius, tc.threshold)
		if got != tc.want {
			t.Errorf("%s: got %q; want %q", tc.name, got, tc.want)
		}
	}
}
```

`cases` เป็น slice ของ struct ที่ประกาศชนิดไว้ตรงนั้น แต่ละรายการเก็บชื่อ ข้อมูลเข้า และผลที่ต้องการ ใช้ `range` ทดสอบทีละรายการด้วยขั้นตอนเดิม

**ลองคิดก่อนรัน:** กรณีใดตรวจค่าขอบเขต และกรณีใดตรวจว่าเลือกเกณฑ์อื่นได้?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ควรเห็น `PASS: TestStatus` หลังตรวจครบสี่กรณี

</details>

### 6. ตั้งชื่อผลทดสอบแต่ละกรณีด้วย t.Run

`t.Run` รับชื่อกรณีกับฟังก์ชันที่จะทดสอบ `func(t *testing.T) { ... }` คือฟังก์ชันที่ไม่ตั้งชื่อและส่งให้ `t.Run` เรียก แทน `main_test.go` ทั้งไฟล์ด้วยเวอร์ชันนี้ แล้วรัน `go test -v .`:

```go
package main

import "testing"

func TestStatus(t *testing.T) {
	cases := []struct {
		name      string
		celsius   float64
		threshold float64
		want      string
	}{
		{"below", 29.9, 30, "OK"},
		{"equal", 30, 30, "WARNING"},
		{"above", 30.1, 30, "WARNING"},
		{"custom threshold", 30, 35, "OK"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := status(tc.celsius, tc.threshold)
			if got != tc.want {
				t.Errorf("status(%v, %v) = %q; want %q",
					tc.celsius, tc.threshold, got, tc.want)
			}
		})
	}
}
```

ข้อมูลทดสอบยังเป็นสี่กรณีเดิม แต่ผลแยกชื่อ `below`, `equal`, `above` และ `custom_threshold` ชัดขึ้น เมื่อมีกรณีผิดจึงหาได้ง่าย

**ลองคิดก่อนรัน:** ตอนจงใจเปลี่ยน >= เป็น > กรณีชื่อใดควรตรวจพบ?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

ควรเห็น `PASS` ของ `TestStatus` และกรณีย่อยทั้งสี่ ชื่อที่มีช่องว่างจะแสดงเป็น `_`

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ [main.go](main.go):

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
	fmt.Println("29.9 C:", status(29.9, 30))
	fmt.Println("30.0 C:", status(30, 30))
}
```

<details>
<summary>โค้ดทดสอบใน main_test.go — เปิดเมื่อถึงขั้นทดสอบ</summary>

```go
package main

import "testing"

func TestStatus(t *testing.T) {
	cases := []struct {
		name      string
		celsius   float64
		threshold float64
		want      string
	}{
		{"below", 29.9, 30, "OK"},
		{"equal", 30, 30, "WARNING"},
		{"above", 30.1, 30, "WARNING"},
		{"custom threshold", 30, 35, "OK"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := status(tc.celsius, tc.threshold)
			if got != tc.want {
				t.Errorf("status(%v, %v) = %q; want %q",
					tc.celsius, tc.threshold, got, tc.want)
			}
		})
	}
}
```

</details>

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run .
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep13-unit-tests` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

รันทดสอบจากโฟลเดอร์เดียวกันด้วย `go test -v .` เมื่อทุกกรณีผ่านจะมี PASS และ ok เวลาและชื่อ package ที่แสดงอาจต่างกัน

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าอุณหภูมิ 30 ที่เกณฑ์ 30 ต้องได้สถานะใด และ test กรณีใดควรไม่ผ่านถ้าเปลี่ยน `>=` เป็น `>`

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
29.9 C: OK
30.0 C: WARNING
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- ตั้งชื่อไฟล์ให้ลงท้ายด้วย `_test.go` และชื่อฟังก์ชันขึ้นต้นด้วย `Test` ตามด้วยชื่อที่ขึ้นต้นตัวใหญ่ เช่น `TestStatus` โดยรับ `*testing.T` เพื่อรายงานผล
- `cases` เป็น slice ของ struct แต่ละรายการเก็บค่าที่ส่งให้ฟังก์ชันและผลที่ต้องการ จึงทดสอบหลายกรณีด้วยลูปเดียวได้
- `t.Run` ตั้งชื่อกรณีย่อย ส่วน `t.Errorf` รายงานรายละเอียดและทำให้ test ไม่ผ่าน `got` คือผลจริง `want` คือผลที่ต้องการ
- ใช้ค่าต่ำกว่า เท่ากับ และสูงกว่าเกณฑ์ เพราะค่าขอบเขตช่วยตรวจพบการใช้ `>` กับ `>=` สลับกัน
- ต้องรันด้วย `go test` เพราะ `go run` ไม่รันไฟล์ทดสอบ `PASS` หมายถึงกรณีที่เขียนไว้ผ่าน ไม่ได้ยืนยันว่าปราศจากข้อผิดพลาดทั้งหมด
- `func(t *testing.T) { ... }` คือฟังก์ชันที่ไม่ได้ตั้งชื่อ เราส่งให้ `t.Run` เรียกเพื่อทดสอบแต่ละกรณี

**ข้อผิดพลาดที่พบบ่อย:** ถ้าขึ้น `[no test files]` ให้ตรวจชื่อไฟล์และตำแหน่ง terminal ถ้าเปลี่ยน `>=` เป็น `>` แล้ว test ยังผ่าน ให้ตรวจว่ามีกรณี `equal` และกำลังรันในโฟลเดอร์ที่แก้โค้ดจริง ให้แก้โค้ดที่ผิด ไม่เปลี่ยน `want` เพียงเพื่อให้ test ผ่าน

</details>

## ฝึกเอง

ฝึกใน `practics` เดิม ก่อนทำแต่ละข้อให้เริ่มจากตัวอย่างทั้ง `main.go` และ `main_test.go` ของบทนี้ หากต้องการเก็บคำตอบก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่มกรณีทดสอบใน `main_test.go` โดยใช้เกณฑ์ 0 กับค่า -0.1 และ 0 แล้วกำหนดผลที่ต้องการตามกฎเดิม ไม่ต้องแก้ `main.go`
2. สมมติเปลี่ยนกฎให้เตือนเฉพาะค่าที่มากกว่าเกณฑ์ เริ่มจากแก้ผลที่ต้องการของ test กรณี `equal` เป็น `OK` แล้วรันให้เห็นว่าไม่ผ่าน จากนั้นแก้ `status` ให้ตรงกฎใหม่และรัน test ซ้ำ

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ทำไม test ตามกฎเดิมต้องไม่ผ่านเมื่อจงใจเปลี่ยน `>=` เป็น `>`?

<details>
<summary>แนวคำตอบ</summary>

เพื่อยืนยันว่า test ตรวจพบผลลัพธ์ที่ผิดเมื่อค่าเท่ากับเกณฑ์ได้จริง ไม่ได้เพียงรันผ่านโดยไม่ครอบคลุมกรณีสำคัญ

</details>

**นำไปใช้ต่อ:** ตรวจว่าการแก้โค้ดยังคงกฎเดิม และใช้กับโปรเจกต์สรุป

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/add-a-test) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep13)

</details>

[EP.12](../ep12-packages-modules/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.14](../ep14-measurement-report/README.md)
