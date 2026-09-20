# EP.15 — เก็บข้อมูลจำนวนคงที่ด้วย array

**เป้าหมาย:** อ่านสมาชิกและวนดูข้อมูลใน array ได้

## 1. อ่านโค้ด

Array เก็บข้อมูลชนิดเดียวกันตามจำนวนที่กำหนด index เริ่มที่ 0, `len` บอกจำนวนสมาชิก และ `range` ให้ตำแหน่งกับค่าทีละรายการ

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func main() {
	temperatures := [3]float64{27.5, 28, 30}
	fmt.Println("count:", len(temperatures))
	for index, value := range temperatures {
		fmt.Println(index, value)
	}
}
```

## 2. ลองรัน

**ก่อนรัน:** array สามช่องมี index สูงสุดเท่าไร?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep15-arrays'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
count: 3
0 27.5
1 28
2 30
```

Array สามสมาชิกมีตำแหน่ง 0, 1 และ 2

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: range เริ่มจากตำแหน่ง 1 หรือ 0 และสิ้นสุดตรงไหน?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

array เหมือนกล่องไข่ที่กำหนดจำนวนช่องไว้ตั้งแต่แรก ทุกช่องเก็บข้อมูลชนิดเดียวกัน

array คือข้อมูลเรียงลำดับที่ความยาวเป็นส่วนหนึ่งของชนิด; index คือตำแหน่ง เริ่มที่ 0; len ให้จำนวนสมาชิก; range วนอ่าน index และค่าทีละรายการ

- [3]float64 กำหนดสามช่อง ไม่ใช่ช่องหมายเลข 0 ถึง 3
- range ให้สองค่า; ใช้ index แสดงตำแหน่ง และ value แสดงข้อมูลในรอบนั้น
- array เป็นค่า การกำหนดให้ตัวแปรอีกตัวจะคัดลอก array; บท slice จะเห็นพฤติกรรมต่างออกไป

**ข้อผิดพลาดที่พบบ่อย**

- ใช้ index 3 กับ array สามช่อง: ตำแหน่งเกินขอบเขต
- คิดว่า [2]float64 และ [3]float64 เป็นชนิดเดียวกัน: ความยาวเป็นส่วนหนึ่งของชนิด

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/moretypes/6)

</details>

[ตอนก่อนหน้า](../ep14-errors/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep16-slices/README.md)
