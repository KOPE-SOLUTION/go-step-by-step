# EP.10 — เพิ่ม method และแก้ต้นฉบับด้วย pointer

**เป้าหมาย:** เลือกได้ว่างานใดอ่านสำเนา และงานใดตั้งใจแก้ข้อมูลต้นฉบับ

**ก่อนเริ่ม:** EP.9 — struct และการคัดลอกค่า

## ทำความเข้าใจ

**method** คือฟังก์ชันที่ผูกกับชนิดข้อมูล ส่วน **receiver** คือค่าที่ method ทำงานด้วย เช่น `reading.Status()`

**pointer** เป็นค่าที่อ้างถึงตำแหน่งข้อมูล `&reading` ใช้เอาที่อยู่ของ reading ส่วน `*Reading` หมายถึงชนิด pointer ไปยัง Reading

## ลงมือทำ

1. ย้าย status ของบทก่อนมาเป็น method `Status` แล้วเรียกจาก reading
2. ลอง adjustCopy ซึ่งรับ Reading แบบค่า แล้วดูว่าต้นฉบับเปลี่ยนหรือไม่
3. เพิ่ม Adjust ที่รับ `*Reading` ตามตัวอย่าง ลองเรียก `reading.Adjust(2)` แทน pointer.Adjust(2) แล้วเทียบผล

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func (r Reading) Status() string {
	if r.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}

func (r *Reading) Adjust(offset float64) {
	r.Celsius += offset
}

func adjustCopy(r Reading, offset float64) {
	r.Celsius += offset
}

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 29}
	adjustCopy(reading, 2)
	fmt.Printf("after copy: %.1f C [%s]\n", reading.Celsius, reading.Status())
	pointer := &reading
	pointer.Adjust(2)
	fmt.Printf("after pointer: %.1f C [%s]\n", reading.Celsius, reading.Status())
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep10-methods-pointers` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
after copy: 29.0 C [OK]
after pointer: 31.0 C [WARNING]
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- `(r Reading)` รับสำเนา เหมาะกับตัวอย่าง Status ที่อ่านข้อมูลอย่างเดียว
- `(r *Reading)` รับ pointer จึงแก้ Celsius ของต้นฉบับที่ชี้ถึงได้ ตัว pointer เองก็ถูกส่งเป็นสำเนาค่าเช่นกัน
- `*pointer` หมายถึงข้อมูลที่ pointer ชี้อยู่ เรียกว่า dereference แต่ Go ย่อการเข้าถึง field ผ่าน pointer ให้เขียน `r.Celsius` ได้
- ถ้าตัวแปร reading สามารถเอาที่อยู่ได้ Go ช่วยให้เรียก pointer method เป็น `reading.Adjust(2)` ได้
- การตั้งชื่อ Adjust เป็นทางเลือกออกแบบ ส่วนความต่างระหว่าง value/pointer เป็นหลักของภาษา ไม่ต้องใช้ pointer กับทุกค่า

**ลองตรวจเมื่อผิด:** pointer ที่ไม่ได้ชี้ข้อมูลมีค่า nil หากเรียก Adjust บน nil จะผิดพลาดตอนรัน ตัวอย่างจึงเริ่มจาก Reading ที่สร้างแล้วและใช้ &reading ให้ลองไล่ว่า pointer แต่ละตัวอ้างถึงอะไร

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เปลี่ยนเฉพาะการเรียก Adjust ให้ปรับ -1.5 แล้วดูค่าต้นฉบับและสถานะหลัง pointer
2. เปลี่ยน adjustCopy ให้รับ `*Reading` และเรียกด้วย `&reading` โดยคงการปรับทั้งสองครั้งไว้ แล้วคาดเดาผลใหม่

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** pointer receiver หมายความว่า Go หยุดส่ง argument แบบคัดลอกค่าหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ Go ยังคัดลอกค่า pointer แต่สำเนาของ pointer ชี้ข้อมูลต้นฉบับเดียวกัน จึงแก้ข้อมูลนั้นได้

</details>

**นำไปใช้ต่อ:** เข้าใจการเปลี่ยนสถานะของข้อมูล และเตรียมใช้ method ผ่าน interface

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Method_declarations) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep10)

</details>

[EP.9](../ep09-structs/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.11](../ep11-interfaces/README.md)
