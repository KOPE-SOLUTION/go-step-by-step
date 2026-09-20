# EP.24 — ทดสอบค่าต่ำกว่า เท่ากับ และสูงกว่าเกณฑ์

**เป้าหมาย:** ตรวจเกณฑ์ด้วยตารางข้อมูลหลายกรณี

## 1. อ่านโค้ด

Table-driven test เก็บ input และคำตอบไว้ใน slice แล้ววนตรวจแต่ละแถว `[]struct{...}` คือ slice ของ struct ที่ไม่ตั้งชื่อชนิด และ `_` ใช้ทิ้ง index ที่ไม่ได้ใช้ `Errorf` รายงานผลผิดแล้วตรวจแถวถัดไปต่อ

ไฟล์ [examples/threshold_test.go](examples/threshold_test.go)

```go
package main

import "testing"

func TestIsWarningBoundaries(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  bool
	}{
		{name: "below", input: 29.9, want: false},
		{name: "at threshold", input: 30, want: true},
		{name: "above", input: 30.1, want: true},
	}
	for _, test := range tests {
		got := isWarning(test.input)
		if got != test.want {
			t.Errorf("%s: got %v; want %v", test.name, got, test.want)
		}
	}
}
```

ไฟล์ที่ใช้ร่วมกัน: [examples/main.go](examples/main.go) · [examples/threshold.go](examples/threshold.go)

## 2. ลองรัน

**ก่อนรัน:** ถ้าพลาดใช้ > แทน >= แถวไหนจะทำให้ test ไม่ผ่าน?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep24-boundary-tests'
go run ./examples
go test -v ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
false
```

ส่วน go test ควรแสดง PASS และ ok เมื่อผ่าน

`Errorf` รายงานว่า test ไม่ผ่านแล้วตรวจแถวถัดไปต่อ ค่า 30 พอดีช่วยแยกความหมายของ `>` กับ `>=`

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: เหตุใดค่า 30 พอดีจึงสำคัญกว่าทดสอบ 31 หลายครั้ง?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

ประตูที่เปิดเมื่ออายุถึงเกณฑ์ควรตรวจคนที่ยังไม่ถึง คนที่ถึงพอดี และคนที่เกินแล้ว เกณฑ์อุณหภูมิก็ใช้แนวคิดเดียวกัน

boundary คือขอบเขตของเงื่อนไข; table-driven test เก็บ input และคำตอบเป็น slice แล้ววนตรวจ; struct ที่ประกาศตรง ๆ โดยไม่ตั้งชื่อเรียกชนิด struct แบบไม่ตั้งชื่อ; _ คือ blank identifier ใช้ทิ้งค่าที่ไม่ต้องการ; Errorf รายงานว่า test ล้มเหลวและทำคำสั่งถัดไปได้

- ชื่อกรณีทำให้ข้อความผิดพลาดบอกได้ว่าแถวใดไม่ผ่าน
- for _, test := range tests ไม่ใช้ index จึงรับไว้ด้วย _; test คือข้อมูลของแต่ละแถว
- ตารางเป็นแนวทางจัด test ไม่ใช่ syntax เฉพาะของ Go; ตรวจค่าขอบเขตที่มีความหมายแทนการเพิ่มตัวเลขแบบสุ่ม

**ข้อผิดพลาดที่พบบ่อย**

- คัดลอกสูตรจากฟังก์ชันไปสร้าง want: อาจผิดแบบเดียวกันทั้งคู่
- เห็น PASS แล้วสรุปว่าใช้ได้ทุกสถานการณ์: ยังไม่ครอบคลุมข้อมูลทุกชนิดหรือข้อกำหนดอื่น

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/wiki/TableDrivenTests)

</details>

[ตอนก่อนหน้า](../ep23-first-test/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep25-project-reading/README.md)
