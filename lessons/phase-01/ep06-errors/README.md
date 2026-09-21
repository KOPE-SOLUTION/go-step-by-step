# EP.6 — รับมือข้อมูลผิดด้วย error

**เป้าหมาย:** แปลงข้อความเป็นตัวเลข ตรวจ error และให้โปรแกรมรับข้อมูลรายการถัดไปได้

**ก่อนเริ่ม:** EP.5 — ฟังก์ชันรับและคืนค่า

## ทำความเข้าใจ

**error** เป็นค่าที่บอกว่างานไม่สำเร็จ ฟังก์ชันใน Go มักคืนทั้งผลลัพธ์และ error ส่วน **nil** ใช้ในตัวอย่างนี้เพื่อบอกว่าไม่มี error ผู้เรียกต้องตรวจเองว่าจะทำอะไรต่อ

`strconv` เป็น package ที่ช่วยแปลงข้อความและตัวเลข เราจะลองข้อมูลดี ข้อความที่ไม่ใช่ตัวเลข และตัวเลขนอกช่วง

## ลงมือทำ

1. เรียก `strconv.ParseFloat("27.5", 64)` เพื่อรับค่าทศนิยมและ error แล้วลองเปลี่ยนข้อความเป็น `warm`
2. รวมการแปลงและตรวจช่วงไว้ใน `parseCelsius` ตามตัวอย่าง แต่ละทางเดินต้องคืนสองค่า
3. เรียกผ่าน show หลายครั้ง เดาก่อนว่ารายการผิดจะทำให้บรรทัด accepted สุดท้ายหายไปหรือไม่

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import (
	"fmt"
	"strconv"
)

func parseCelsius(text string) (float64, error) {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("temperature must be a number: %q", text)
	}
	if value < 0 || value > 100 {
		return 0, fmt.Errorf("temperature outside simulated range: %.1f", value)
	}
	return value, nil
}

func show(text string) {
	value, err := parseCelsius(text)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("accepted: %.1f C\n", value)
}

func main() {
	show("27.5")
	show("warm")
	show("101")
	show("30")
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep06-errors` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
accepted: 27.5 C
ERROR: temperature must be a number: "warm"
ERROR: temperature outside simulated range: 101.0
accepted: 30.0 C
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- `value, err := ...` รับผลลัพธ์สองค่า ต้องตรวจ `err != nil` ก่อนใช้ value เป็นค่าที่สำเร็จ
- `fmt.Errorf` สร้าง error พร้อมข้อความอธิบาย คืน 0 เป็นผลลัพธ์ประกอบในทางที่ล้มเหลว แต่ผู้เรียกไม่ควรนำ 0 นั้นไปใช้เมื่อ err ไม่ใช่ nil
- `return` ใน show จบเฉพาะการเรียก show ครั้งนั้น main จึงเรียกข้อมูลรายการถัดไปได้
- เราแยก error ระหว่างการแปลงกับกฎช่วงข้อมูล เพื่อให้รู้ว่าควรแก้ input อย่างไร การคืน error ให้ผู้เรียกตัดสินใจเป็นแนวทางที่ใช้ทั่วไป
- ตัวอย่างนี้ตรวจข้อความธรรมดาและช่วง 0–100 เท่านั้น ค่าพิเศษ NaN/Inf และการเก็บสาเหตุต้นทางด้วย error wrapping จะทบทวนเมื่อรับ input จริง

**ลองตรวจเมื่อผิด:** อย่าเขียน `value, _ := ...` เพื่อทิ้ง error แล้วใช้ value ต่อ เครื่องหมาย `_` คือทิ้งค่าที่รับมา ลองตรวจว่าข้อมูลผิดถูกนำไปแสดงเป็นค่าปกติหรือไม่

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เพิ่ม show("0") และ show("100") ต่อท้าย เพื่อยืนยันว่าปลายช่วงทั้งสองรับได้
2. ก่อน ParseFloat ให้ตรวจข้อความว่าง และคืน error ว่า `temperature is required` จากนั้นเปลี่ยน warm ใน main เป็นข้อความว่าง

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้า parseCelsius คืน 0 พร้อม error ค่า 0 นั้นนับเป็นอุณหภูมิที่อ่านสำเร็จหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ ต้องดู error ก่อน ผลลัพธ์ 0 ในทางที่ล้มเหลวเป็นเพียงค่าที่คืนให้ครบตามชนิดของฟังก์ชัน

</details>

**นำไปใช้ต่อ:** จัดการ input ที่เสียก่อนนำไปคำนวณหรือบันทึก

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/handle-errors) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep6)

</details>

[EP.5](../ep05-functions/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.7](../ep07-arrays-slices/README.md)
