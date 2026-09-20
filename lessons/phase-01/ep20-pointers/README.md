# EP.20 — ใช้ pointer เมื่อต้องการแก้ค่าต้นฉบับ

**เป้าหมาย:** อธิบาย & และ * พร้อมติดตามค่าที่ถูกแก้ได้

## 1. อ่านโค้ด

Pointer ชี้ไปยังตัวแปร `&temperature` ให้ตำแหน่งของตัวแปร, `*float64` ระบุชนิด pointer และ `*value` ใช้อ่านหรือแก้ค่าที่ชี้อยู่

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func addOffset(value *float64, offset float64) {
	*value = *value + offset
}

func main() {
	temperature := 25.0
	addOffset(&temperature, 0.5)
	fmt.Println(temperature)
}
```

## 2. ลองรัน

**ก่อนรัน:** ถ้าฟังก์ชันรับ float64 ธรรมดาแล้วแก้ภายใน ค่า temperature ภายนอกจะเปลี่ยนไหม?

รันจากโฟลเดอร์ `lessons/phase-01/ep20-pointers`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
25.5
```

ตัวอย่างส่ง pointer ของตัวแปรที่มีอยู่แล้ว จึงแก้ค่าต้นฉบับได้

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. เปลี่ยนค่าเริ่ม 20 และ offset -0.5 ให้ได้ 19.5
2. สร้าง n := 5 และ p := &n แล้วเพิ่มค่าผ่าน *p อีก 1 พิมพ์ n

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

เครื่องหมาย * ใน *float64 และ *value ทำหน้าที่เหมือนกันหรือไม่?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

การส่งสำเนาแบบฟอร์มไม่เปลี่ยนใบเดิม แต่ pointer เหมือนบอกตำแหน่งของใบเดิม ผู้รับจึงแก้ข้อมูลที่ตำแหน่งนั้นได้

pointer คือค่าที่ชี้ไปยังตัวแปร; &value ให้ pointer ไปยัง value; *float64 คือชนิด pointer ไปยัง float64; *p ในคำสั่งคือค่าที่ pointer ชี้; pointer ที่ยังไม่กำหนดเป็น nil และนำมา dereference ไม่ได้

- &temperature ส่งตำแหน่งตัวแปรให้ฟังก์ชัน
- *value ทางซ้ายของ = เขียนค่าที่ตำแหน่งนั้น ทางขวาอ่านค่าเดิม
- ตัวอย่างต้องส่ง pointer ที่ไม่เป็น nil; ไม่ต้องใช้ pointer กับทุกอย่าง ใช้เมื่องานต้องแก้ต้นฉบับหรือมีเหตุผลชัดเจน

**ข้อผิดพลาดที่พบบ่อย**

- ส่ง temperature แทน &temperature: float64 ไม่ตรงกับ *float64
- ใช้ var p *float64 แล้วเขียน *p: p เป็น nil ต้องมีตัวแปรให้ชี้ก่อน

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/moretypes/1)

</details>

[ตอนก่อนหน้า](../ep19-methods/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep21-packages/README.md)
