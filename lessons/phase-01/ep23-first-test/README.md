# EP.23 — เขียน unit test แรก

**เป้าหมาย:** ให้ Go ตรวจคำตอบของฟังก์ชันแทนการมองผลเอง

## 1. อ่านโค้ด

Unit test ตรวจฟังก์ชันด้วยผลที่คาดไว้ เก็บในไฟล์ `_test.go` และตั้งชื่อ `Test...` ส่วน `t *testing.T` ช่วยรายงานผลทดสอบ `got` คือผลจริง `want` คือผลที่คาด ส่วน `Fatalf` รายงานและหยุด test นั้น โดย `%v` ใช้แสดงค่า

ไฟล์ [examples/threshold.go](examples/threshold.go)

```go
package main

func isWarning(celsius float64) bool {
	return celsius >= 30
}
```

ไฟล์ [examples/threshold_test.go](examples/threshold_test.go)

```go
package main

import "testing"

func TestIsWarning(t *testing.T) {
	got := isWarning(30)
	want := true
	if got != want {
		t.Fatalf("isWarning(30) = %v; want %v", got, want)
	}
}
```

ไฟล์ที่ใช้ร่วมกัน: [examples/main.go](examples/main.go)

## 2. ลองรัน

**ก่อนรัน:** เปลี่ยน return celsius >= 30 เป็น > 30 แล้ว test ค่า 30 ควรผ่านหรือไม่?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep23-first-test'
go run ./examples
go test -v ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
true
```

ส่วน go test ควรแสดง PASS และ ok เมื่อผ่าน

`got` คือผลจริง `want` คือผลที่คาด `Fatalf` รายงานข้อผิดพลาดแล้วหยุด test นั้น และ `%v` ใช้แสดงค่าทั่วไป

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: test ผ่านหนึ่งกรณีพิสูจน์ว่าฟังก์ชันถูกทุกค่าหรือไม่?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

แทนการกดเครื่องคิดเลขแล้วดูด้วยตาทุกครั้ง เราเขียนข้อสอบพร้อมคำตอบที่คาดไว้ ถ้าโปรแกรมเปลี่ยนจนตอบผิด การทดสอบจะแจ้งให้เห็น

unit test ตรวจหน่วยงานเล็ก เช่น ฟังก์ชัน; ไฟล์ _test.go ใช้เก็บ test; testing เป็น package มาตรฐาน; Test... เป็นชื่อฟังก์ชันทดสอบ; t *testing.T เป็นตัวช่วยรายงานผลที่ Go ส่งให้; Fatalf รายงานและหยุด test นั้น; %v แสดงค่าทั่วไป

- go test ./... ค้น package ภายใต้ module แล้วรันฟังก์ชันทดสอบ ไม่เรียก main เพื่อพิมพ์ตัวอย่าง
- got คือผลจริง want คือผลที่ข้อกำหนดต้องการ ตั้งชื่อตามแนวทางให้อ่านง่าย
- ไฟล์ .go ใน package เดียวกันใช้ฟังก์ชันร่วมกันได้ ไม่ต้อง import กันเอง; รัน go run ./examples เพื่อรวม main กับ threshold.go

**ข้อผิดพลาดที่พบบ่อย**

- ตั้งชื่อไฟล์ threshold-test.go: ต้องลงท้าย _test.go
- ใช้ go run ./examples/main.go: จะไม่รวม threshold.go ให้รันทั้ง package ด้วย ./examples

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/add-a-test)

</details>

[ตอนก่อนหน้า](../ep22-modules/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep24-boundary-tests/README.md)
