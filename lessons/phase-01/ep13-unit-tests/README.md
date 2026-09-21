# EP.13 — ทดสอบกฎและค่าขอบเขตด้วย unit test

**เป้าหมาย:** เขียน test ที่จับความผิดพลาดตรงค่าเกณฑ์ได้ และอ่านว่าผลจริงต่างจากที่ต้องการอย่างไร

**ก่อนเริ่ม:** EP.5, EP.7, EP.9 และ EP.12 — ฟังก์ชัน, slice, struct และ module

## ทำความเข้าใจ

**unit test** คือโค้ดตรวจพฤติกรรมของส่วนเล็ก ๆ เช่นฟังก์ชัน status เรากำหนด input และผลที่ต้องการ แล้วให้เครื่องเปรียบเทียบแทนการรันดูเองทุกครั้ง

ในบทนี้ใช้กฎเดิม: ตั้งแต่เกณฑ์ขึ้นไปเป็น WARNING ก่อนเขียน test ให้ตอบผลของ 29.9, 30 และ 30.1 ที่เกณฑ์ 30

## ลงมือทำ

1. แทน main.go ด้วยตัวอย่างด้านล่าง ใช้ go.mod เดิมจาก EP.12 ไฟล์ sensor/reading.go เดิมเก็บไว้ได้ แต่บทนี้ยังไม่ import
2. สร้าง `main_test.go` ข้าง main.go เริ่มจาก test กรณีเท่ากับเกณฑ์หนึ่งกรณี แล้วค่อยเพิ่มตารางตามตัวอย่าง
3. รัน `go test -v .` จาก practics จากนั้นเปลี่ยน `>=` ใน status เป็น `>` ต้องเห็น test equal ล้มเหลว อ่าน got/want แล้วแก้กลับก่อนรันซ้ำ

### ตัวอย่างเมื่อทำครบ

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

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run .
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep13-unit-tests` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

รันทดสอบจากโฟลเดอร์เดียวกันด้วย `go test -v .` เมื่อทุกกรณีผ่านจะมี PASS และ ok เวลาและชื่อ package ที่แสดงอาจต่างกัน

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
29.9 C: OK
30.0 C: WARNING
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- ไฟล์ลงท้าย `_test.go` และฟังก์ชันชื่อ Test ตามด้วยชื่อขึ้นต้นตัวใหญ่ รับ `*testing.T` เพื่อรายงานผล
- ชุด cases เป็น slice ของ struct ที่รวม input กับผลที่ต้องการ เราใช้ความรู้เดิมสร้างหลายกรณีใน test เดียว
- `t.Run` ตั้งชื่อกรณีย่อย ส่วน `t.Errorf` รายงานรายละเอียดและทำให้ test ไม่ผ่าน `got` คือผลจริง `want` คือผลที่ต้องการ
- ใช้ค่าต่ำกว่า เท่ากับ และสูงกว่าเกณฑ์ เพราะค่าขอบเขตช่วยจับการเขียน > กับ >= ผิด
- ต้องรันด้วย `go test` เพราะ go run ไม่รันไฟล์ทดสอบ PASS หมายถึงกรณีที่เขียนไว้ผ่าน ไม่ได้ยืนยันว่าปราศจากข้อผิดพลาดทั้งหมด
- `func(t *testing.T) { ... }` ภายใน t.Run คือฟังก์ชันไม่มีชื่อที่ส่งให้ t.Run เรียก ส่วนรายละเอียด callback จะใช้เพิ่มเมื่อเรียน API

**ลองตรวจเมื่อผิด:** ถ้าขึ้น `[no test files]` ให้ตรวจชื่อไฟล์และตำแหน่ง terminal ถ้า test ผ่านแม้ใส่ bug ให้ตรวจว่าเพิ่มกรณี equal แล้ว และกำลังทดสอบ package ที่แก้จริง ไม่เปลี่ยน want เพื่อให้ bug ผ่าน

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เพิ่ม test เกณฑ์ 0 กับค่า -0.1 และ 0 เพื่อตรวจว่าฟังก์ชันเปรียบเทียบตามเกณฑ์ที่ส่งมาได้ โปรแกรมหลักไม่ต้องเปลี่ยน
2. สมมติเปลี่ยนกฎเป็นเกินเกณฑ์เท่านั้นจึงเตือน ปรับ test equal ให้คาดว่า OK ก่อน รันให้ล้มเหลว แล้วปรับ status ให้ตรงกฎใหม่

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ทำไมต้องเห็น test ล้มเหลวเมื่อจงใจเปลี่ยน >= เป็น >?

<details>
<summary>แนวคำตอบ</summary>

เพื่อยืนยันว่า test ตรวจความต่างของพฤติกรรมตรงเกณฑ์ได้จริง ไม่ได้เพียงรันผ่านโดยไม่ครอบคลุมกรณีสำคัญ

</details>

**นำไปใช้ต่อ:** ตรวจว่าการแก้โค้ดยังคงกฎเดิม และใช้กับโปรเจกต์สรุป

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/add-a-test) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep13)

</details>

[EP.12](../ep12-packages-modules/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.14](../ep14-measurement-report/README.md)
