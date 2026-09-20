# EP.19 — เพิ่ม method ให้กับชนิดข้อมูล

**เป้าหมาย:** เรียกการตรวจสถานะจาก Reading

## 1. อ่านโค้ด

Method คือฟังก์ชันที่ผูกกับชนิดข้อมูล วงเล็บ `(r Reading)` ก่อนชื่อ method เรียกว่า receiver และรับค่าจากฝั่งที่เรียก

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

type Reading struct {
	Celsius float64
}

func (r Reading) IsWarning() bool {
	return r.Celsius >= 30
}

func main() {
	reading := Reading{Celsius: 30}
	fmt.Println(reading.IsWarning())
}
```

## 2. ลองรัน

**ก่อนรัน:** r ใน method มาจากตัวแปรใดตอนเรียก reading.IsWarning()?

รันจากโฟลเดอร์ `lessons/phase-01/ep19-methods`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
true
```

ตัวอย่างนี้ใช้สำเนา Reading เพื่อตรวจสถานะ โดยไม่แก้ข้อมูลต้นฉบับ

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: ทำไม IsWarning จึงเหมาะกับ value receiver ในตัวอย่างนี้?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เมื่อมีแบบฟอร์มการวัดแล้ว เราตั้งวิธีอ่านสถานะไว้คู่กับชนิดข้อมูลนั้น ผู้เรียกจึงเขียน reading.IsWarning() ได้

method คือฟังก์ชันที่มี receiver; receiver อยู่ในวงเล็บก่อนชื่อ method ระบุชนิดที่ผูกด้วย; value receiver รับสำเนาค่าในตัวอย่างนี้

- (r Reading) คือ receiver ไม่ใช่ parameter หลังชื่อ method
- IsWarning อ่านค่าและคืน bool ไม่แก้ต้นฉบับ
- จะเขียนเป็นฟังก์ชัน isWarning(reading Reading) ก็ได้ การเลือก method เป็นการออกแบบให้พฤติกรรมอยู่กับชนิด

**ข้อผิดพลาดที่พบบ่อย**

- ใส่ receiver ไว้หลังชื่อ method: ตำแหน่งนั้นจะกลายเป็น parameter
- หวังให้การแก้ r.Celsius ใน value receiver เปลี่ยนต้นฉบับ: บทนี้ r เป็นสำเนา

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/methods/1)

</details>

[ตอนก่อนหน้า](../ep18-structs/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep20-pointers/README.md)
