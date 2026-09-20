# EP.16 — ใช้ slice กับรายการที่เพิ่มได้

**เป้าหมาย:** เพิ่มสมาชิกและเข้าใจการใช้พื้นที่ข้อมูลร่วมกัน

## 1. อ่านโค้ด

Slice อ้างถึงช่วงข้อมูลของ array และใช้ `[]` โดยไม่ระบุจำนวนสมาชิก `append` คืน slice หลังเพิ่มข้อมูล จึงต้องรับค่ากลับ

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func main() {
	temperatures := []float64{27.5, 28}
	temperatures = append(temperatures, 30)
	fmt.Println(temperatures)
	fmt.Println("count:", len(temperatures))
}
```

## 2. ลองรัน

**ก่อนรัน:** append แล้ว len เปลี่ยนเท่าไร และทำไมต้องรับค่าคืนกลับ?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep16-slices'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
[27.5 28 30]
count: 3
```

การกำหนด slice ให้อีกตัวไม่ได้คัดลอกสมาชิกเป็นชุดใหม่ ลองสังเกตเรื่องนี้ในแบบฝึกหัด

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: กำหนด b := a เมื่อ a เป็น slice ทำให้สมาชิกเป็นสำเนาอิสระหรือไม่?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

slice เหมือนหน้าต่างที่มองสมาชิกส่วนหนึ่งของ array เราใช้ [] โดยไม่ระบุจำนวนเพื่อทำงานกับรายการที่ยืดหยุ่น

slice คือข้อมูลที่อ้างถึงช่วงของ array; append คืน slice หลังเพิ่มสมาชิก; a[:2] เลือกตั้งแต่ต้นถึงก่อน index 2; slice ไม่ใช่การคัดลอกสมาชิกทั้งหมด

- []float64 ต่างจาก [3]float64 ตรงที่ไม่กำหนดความยาวไว้ในชนิด
- append อาจใช้ array เดิมหรือจัดสรรใหม่ ต้องเก็บ slice ที่คืนมา
- slice ที่เกิดจากการตัดช่วงหรือกำหนดต่อกันอาจเห็นสมาชิกชุดเดียวกัน อย่าสมมติว่าเป็นสำเนาอิสระ

**ข้อผิดพลาดที่พบบ่อย**

- เขียน append(...) แล้วไม่เก็บค่าคืน: ต้องใช้ temperatures = append(...)
- อ่าน slice ว่างที่ index 0: ตรวจ len ก่อนเข้าถึง

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/moretypes/7)

</details>

[ตอนก่อนหน้า](../ep15-arrays/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep17-maps/README.md)
