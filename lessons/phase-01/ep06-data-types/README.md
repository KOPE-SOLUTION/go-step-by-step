# EP.6 — ชนิดข้อมูล string, int, float64 และ bool

**เป้าหมาย:** เลือกชนิดข้อมูลพื้นฐานให้ตรงความหมาย

## 1. อ่านโค้ด

`string` เก็บข้อความ, `int` เก็บจำนวนเต็ม, `float64` เก็บตัวเลขทศนิยม และ `bool` เก็บ `true` หรือ `false`

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func main() {
	var deviceID string = "001"
	var count int = 3
	var temperature float64 = 27.5
	var online bool = true
	fmt.Println(deviceID)
	fmt.Println(count)
	fmt.Println(temperature)
	fmt.Println(online)
}
```

## 2. ลองรัน

**ก่อนรัน:** 001 ที่เก็บเป็น string จะมีศูนย์นำหน้าในผลหรือไม่?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep06-data-types'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
001
3
27.5
true
```

`"001"` เป็นข้อความ จึงเก็บเลขศูนย์ด้านหน้าไว้ได้

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: รหัส 001 กับจำนวน 1 ควรเลือกชนิดเหมือนกันเสมอไหม?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

รหัสอุปกรณ์กับจำนวนอุปกรณ์อาจมีตัวเลขเหมือนกัน แต่ทำหน้าที่ต่างกัน รหัส 001 ต้องรักษาศูนย์นำหน้า ส่วนจำนวนต้องใช้คำนวณ ชนิดข้อมูลช่วยกันใช้ค่าผิดประเภท

**string** ข้อความ; **int** จำนวนเต็ม; **float64** จำนวนที่เก็บเศษทศนิยมแบบประมาณค่า; **bool** ค่า true หรือ false; `var name type = value` ประกาศพร้อมระบุชนิด

- ระบุชนิดให้เห็นชัดก่อนกลับไปใช้ := ในตอนอื่น
- 27.5 เป็นค่าตัวเลข จึงไม่ครอบ quote แต่รหัส 001 เป็นข้อความ
- bool ใช้ true/false โดยไม่ใส่ quote การใส่ quote จะกลายเป็น string

**ข้อผิดพลาดที่พบบ่อย**

- ใส่ "true" ในตัวแปร bool: ข้อความไม่ใช่ boolean
- ใส่ 27.5 ใน int: จำนวนเต็มไม่เก็บเศษแบบนี้ ต้องเลือกชนิดให้ตรง

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/basics/11)

</details>

[ตอนก่อนหน้า](../ep05-variables/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep07-zero-values/README.md)
