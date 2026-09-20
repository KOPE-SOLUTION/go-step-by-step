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

รันจากโฟลเดอร์ `lessons/phase-01/ep23-first-test`:

```shell
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

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** และ `go.mod` จาก EP.21 เขียนทับ `main.go` ด้วยโค้ดตั้งต้นที่รวมไว้ด้านล่าง แล้วทำโจทย์

<details>
<summary>โค้ดตั้งต้นสำหรับ main.go</summary>

```go
package main

import "fmt"

func main() {
	fmt.Println(isWarning(30))
}

func isWarning(celsius float64) bool {
	return celsius >= 30
}
```

</details>

สร้าง `practics/main_test.go` แล้วใส่โค้ดทั้งไฟล์จาก [ตัวอย่าง test](examples/threshold_test.go) ไฟล์ทดสอบต้องลงท้าย `_test.go` จึงแยกจาก `main.go`

**แก้ `main.go` และ `main_test.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. เพิ่ม test ใน `main_test.go` สำหรับค่า `29.9` ซึ่งต้องได้ `false` โดยเก็บ test เดิมไว้
2. เพิ่ม test สำหรับค่า `30.1` ซึ่งต้องได้ `true` จากนั้นลองเปลี่ยน `>=` เป็น `>` ใน `main.go` ดูว่า test ใดไม่ผ่าน แล้วคืนโค้ดเดิม

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run .
go test .
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

test ผ่านหนึ่งกรณีพิสูจน์ว่าฟังก์ชันถูกทุกค่าหรือไม่?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

แทนการกดเครื่องคิดเลขแล้วดูด้วยตาทุกครั้ง เราเขียนข้อสอบพร้อมคำตอบที่คาดไว้ ถ้าโปรแกรมเปลี่ยนจนตอบผิด การทดสอบจะแจ้งให้เห็น

unit test ตรวจหน่วยงานเล็ก เช่น ฟังก์ชัน; ไฟล์ _test.go ใช้เก็บ test; testing เป็น package มาตรฐาน; Test... เป็นชื่อฟังก์ชันทดสอบ; t *testing.T เป็นตัวช่วยรายงานผลที่ Go ส่งให้; Fatalf รายงานและหยุด test นั้น; %v แสดงค่าทั่วไป

- go test ./... ค้น package ภายใต้ module แล้วรันฟังก์ชันทดสอบ ไม่เรียก main เพื่อพิมพ์ตัวอย่าง
- got คือผลจริง want คือผลที่ข้อกำหนดต้องการ ตั้งชื่อตามแนวทางให้อ่านง่าย
- ไฟล์ .go ใน package เดียวกันใช้ฟังก์ชันร่วมกันได้ ไม่ต้อง import กันเอง; รัน go run ./examples เพื่อรวม main กับ threshold.go

**ข้อผิดพลาดที่พบบ่อย**

- ตั้งชื่อไฟล์ threshold-test.go: ต้องลงท้าย _test.go
- ในตัวอย่าง `lessons/` การรันเฉพาะ main.go จะไม่รวม threshold.go ให้ใช้ `go run ./examples` ส่วนโค้ดฝึกได้รวมฟังก์ชันไว้ใน `main.go` แล้ว

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/add-a-test)

</details>

[ตอนก่อนหน้า](../ep22-modules/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep24-boundary-tests/README.md)
