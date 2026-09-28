# EP.6 — รับมือข้อมูลผิดด้วย error

**เป้าหมาย:** แปลงข้อความเป็นตัวเลข ตรวจ error และให้โปรแกรมรับข้อมูลรายการถัดไปได้

**ก่อนเริ่ม:** EP.5 — ฟังก์ชันรับและคืนค่า

## ทำความเข้าใจ

**error** เป็นค่าที่บอกว่างานไม่สำเร็จ ฟังก์ชันใน Go มักคืนทั้งผลลัพธ์และ error ส่วน **nil** ใช้ในตัวอย่างนี้เพื่อบอกว่าไม่มี error ผู้เรียกต้องตรวจเองว่าจะทำอะไรต่อ

`strconv` เป็น package ที่ช่วยแปลงข้อความและตัวเลข เราจะลองข้อความที่แปลงเป็นอุณหภูมิได้ ข้อความที่ไม่ใช่ตัวเลข และตัวเลขนอกช่วง 0–100

## ลงมือทำ

1. เรียก `strconv.ParseFloat("27.5", 64)` เพื่อรับค่าทศนิยมและ error แล้วลองเปลี่ยนข้อความเป็น `warm`
2. รวมการแปลงและตรวจช่วงไว้ใน `parseCelsius` ตามตัวอย่าง ทุกคำสั่ง `return` ในฟังก์ชันนี้ต้องคืนสองค่า คือ `float64` และ `error`
3. เรียก `show` หลายครั้งตามตัวอย่าง แล้วสังเกตว่าเมื่อพบข้อมูลผิด โปรแกรมยังตรวจข้อมูลถัดไปหรือไม่

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

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าเมื่อพบ `"warm"` และ `"101"` โปรแกรมจะยังแสดงบรรทัด `accepted` ของข้อมูลถัดไปหรือไม่

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
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `value, err := ...` รับผลลัพธ์สองค่า ต้องตรวจ `err != nil` ก่อนนำ `value` ไปใช้เป็นอุณหภูมิที่ผ่านการตรวจ
- `fmt.Errorf` สร้าง error พร้อมข้อความอธิบาย ในตัวอย่างนี้ `parseCelsius` คืน `0` คู่กับ error เมื่อแปลงหรือตรวจข้อมูลไม่สำเร็จ ผู้เรียกจึงไม่ควรนำ `0` ไปใช้เป็นอุณหภูมิเมื่อ `err != nil`
- `return` ใน `show` จบการเรียก `show` ครั้งนั้น แล้วกลับไปทำคำสั่งถัดไปใน `main` จึงยังเรียก `show` เพื่อตรวจข้อมูลถัดไปได้
- เราแยก error ระหว่างการแปลงกับกฎช่วงข้อมูล เพื่อให้รู้ว่าข้อมูลที่รับมาผิดตรงไหน การคืน error ให้ผู้เรียกตัดสินใจเป็นแนวทางที่ใช้ทั่วไป
- ตัวอย่างนี้ตรวจข้อความธรรมดาและช่วง 0–100 เท่านั้น ยังไม่ได้ทดสอบค่าพิเศษของเลขทศนิยม เช่น `NaN` และ `Inf` ส่วนการเก็บ error เดิมไว้พร้อมข้อความเพิ่มเติมจะเรียนเมื่อรับข้อมูลจากภายนอกในบทต่อไป

**ข้อผิดพลาดที่พบบ่อย:** อย่าเขียน `value, _ := ...` เพื่อทิ้ง error แล้วใช้ value ต่อ เครื่องหมาย `_` คือทิ้งค่าที่รับมา ลองตรวจว่าข้อมูลผิดถูกนำไปแสดงเป็นค่าปกติหรือไม่

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่ม `show("0")` และ `show("100")` ท้ายฟังก์ชัน `main` เพื่อตรวจว่าค่า 0 และ 100 ยังอยู่ในช่วงที่รับได้
2. ใน `parseCelsius` ให้ตรวจข้อความว่างก่อนเรียก `strconv.ParseFloat` และคืน error ว่า `temperature is required` จากนั้นเปลี่ยน `show("warm")` ใน `main` เป็น `show("")` เพื่อทดลอง

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้า `parseCelsius` คืน `0` พร้อม error ค่า 0 นั้นนับเป็นอุณหภูมิที่อ่านสำเร็จหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ เมื่อมี error แสดงว่าข้อมูลไม่ผ่านการตรวจ ค่า `0` ที่คืนคู่กันจึงไม่ใช่อุณหภูมิที่อ่านสำเร็จ

</details>

**นำไปใช้ต่อ:** ตรวจและจัดการข้อมูลที่ไม่ถูกต้องก่อนนำไปคำนวณหรือบันทึก

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/doc/tutorial/handle-errors) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep6)

</details>

[EP.5](../ep05-functions/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.7](../ep07-arrays-slices/README.md)
