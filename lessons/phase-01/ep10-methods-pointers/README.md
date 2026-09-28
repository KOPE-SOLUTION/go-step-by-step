# EP.10 — เพิ่ม method และแก้ต้นฉบับด้วย pointer

**เป้าหมาย:** เลือกได้ว่างานใดอ่านสำเนา และงานใดตั้งใจแก้ข้อมูลต้นฉบับ

**ก่อนเริ่ม:** EP.9 — struct และการคัดลอกค่า

## ทำความเข้าใจ

**method** คือฟังก์ชันที่ผูกกับชนิดข้อมูล ส่วน **receiver** คือค่าที่ method ทำงานด้วย เช่น `reading.Status()`

**pointer** เป็นค่าที่อ้างถึงตำแหน่งข้อมูล `&reading` ให้ค่า pointer ที่ชี้ไปยังตัวแปร `reading` ส่วน `*Reading` คือชนิดของ pointer ที่ชี้ไปยังข้อมูลชนิด `Reading`

## ลงมือทำ

1. เปลี่ยนฟังก์ชัน `status` ของบทก่อนเป็น method `Status` แล้วเรียกด้วย `reading.Status()`
2. ลองเรียก `adjustCopy` ซึ่งรับสำเนาของ `Reading` แล้วดูว่าค่าใน `reading` ต้นฉบับเปลี่ยนหรือไม่
3. เพิ่ม method `Adjust` ที่ใช้ receiver ชนิด `*Reading` ตามตัวอย่าง แล้วลองเปลี่ยน `pointer.Adjust(2)` เป็น `reading.Adjust(2)` เพื่อเปรียบเทียบผล

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

ก่อนเปิดผลลัพธ์ ลองคาดเดาค่า `reading.Celsius` หลังเรียก `adjustCopy` และหลังเรียก `Adjust`

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
after copy: 29.0 C [OK]
after pointer: 31.0 C [WARNING]
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `(r Reading)` รับสำเนา เหมาะกับตัวอย่าง `Status` ที่อ่านข้อมูลอย่างเดียว
- `(r *Reading)` รับ pointer จึงแก้ `Celsius` ของต้นฉบับที่ชี้ถึงได้ ตัว pointer เองก็ถูกส่งเป็นสำเนาค่าเช่นกัน
- `*pointer` หมายถึงข้อมูลที่ pointer ชี้อยู่ เรียกว่า dereference แต่ Go ย่อการเข้าถึง field ผ่าน pointer ให้เขียน `r.Celsius` ได้
- ในตัวอย่างนี้ `reading` เป็นตัวแปรที่อ้างถึงตำแหน่งได้ Go จึงยอมให้เขียน `reading.Adjust(2)` แทน `(&reading).Adjust(2)`
- การตั้งชื่อ `Adjust` เป็นทางเลือกในการออกแบบ ส่วนความต่างระหว่างการรับสำเนา `Reading` กับ pointer เป็นหลักของภาษา ไม่ต้องใช้ pointer กับทุกค่า

**ข้อผิดพลาดที่พบบ่อย:** pointer ที่ไม่ได้ชี้ไปยังข้อมูลมีค่า `nil` หากเรียก `Adjust` ในตัวอย่างนี้ผ่าน nil pointer จะเกิดข้อผิดพลาดขณะรัน เพราะ method พยายามแก้ `Celsius` ตัวอย่างจึงสร้าง `reading` ก่อน แล้วใช้ `&reading` เพื่อให้ pointer ชี้ไปยังข้อมูลนั้น

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เปลี่ยนเฉพาะ `pointer.Adjust(2)` เป็น `pointer.Adjust(-1.5)` แล้วตรวจอุณหภูมิและสถานะของ `reading` หลังเรียก method
2. เปลี่ยน `adjustCopy` ให้รับ `*Reading` และเรียกด้วย `&reading` โดยคงการปรับทั้งสองครั้งไว้ แล้วคาดเดาผลใหม่

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** เมื่อส่ง pointer ให้ฟังก์ชันหรือ method Go ยังคัดลอกค่าที่ส่งอยู่หรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ยังคัดลอกค่า pointer อยู่ แต่ pointer ทั้งสองชี้ไปยังข้อมูลเดียวกัน จึงใช้สำเนาของ pointer แก้ข้อมูลต้นฉบับได้

</details>

**นำไปใช้ต่อ:** เข้าใจการเปลี่ยนสถานะของข้อมูล และเตรียมใช้ method ผ่าน interface

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Method_declarations) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep10)

</details>

[EP.9](../ep09-structs/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.11](../ep11-interfaces/README.md)
