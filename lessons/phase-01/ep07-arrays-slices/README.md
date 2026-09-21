# EP.7 — จัดการรายการข้อมูลด้วย array และ slice

**เป้าหมาย:** เก็บหลายค่าการวัด เพิ่มรายการ และสรุปค่าเฉลี่ยกับจำนวนคำเตือน

**ก่อนเริ่ม:** EP.4 — ลูป และ EP.5 — ฟังก์ชัน

## ทำความเข้าใจ

**array** เก็บข้อมูลชนิดเดียวกันจำนวนคงที่ เช่น `[3]float64` ส่วน **slice** ใช้อ้างถึงช่วงของ array เบื้องหลังและเหมาะกับรายการที่เพิ่มจำนวนได้ เขียนชนิดเป็น `[]float64`

**index** คือตำแหน่ง เริ่มจาก 0 ส่วน `len` บอกจำนวนรายการที่ใช้ได้

## ลงมือทำ

1. สร้าง initial แบบ array แล้วอ่าน `initial[0]` และ `len(initial)`
2. ย้ายค่าเข้าสู่ readings แบบ slice ด้วย range และ append แล้วเพิ่มค่า 32 ตามตัวอย่าง
3. ใช้ range สรุปค่าเฉลี่ยและนับคำเตือน ลองทำ slice ว่างแล้วตรวจว่าส่วนค่าเฉลี่ยทำงานอย่างไร

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

func main() {
	initial := [3]float64{27.5, 30.0, 28.5}
	readings := []float64{}
	for _, value := range initial {
		readings = append(readings, value)
	}
	readings = append(readings, 32.0)

	total := 0.0
	warnings := 0
	for index, value := range readings {
		total += value
		if value >= 30 {
			warnings++
		}
		fmt.Printf("%d: %.1f C\n", index, value)
	}
	if len(readings) > 0 {
		fmt.Printf("average: %.1f C, warnings: %d\n",
			total/float64(len(readings)), warnings)
	} else {
		fmt.Println("no readings")
	}
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep07-arrays-slices` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
0: 27.5 C
1: 30.0 C
2: 28.5 C
3: 32.0 C
average: 29.5 C, warnings: 2
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- `[3]float64` มีจำนวน 3 เป็นส่วนหนึ่งของชนิด ส่วน `[]float64{}` สร้าง slice ว่าง
- `range` ให้ index กับค่าของแต่ละรายการ ถ้าไม่ต้องใช้ index เขียน `_` เพื่อทิ้งได้
- ต้องรับผลจาก `append` กลับมา เพราะ slice ที่ได้อาจอ้างถึงพื้นที่เก็บใหม่
- `len` เป็นจำนวนที่เข้าถึงได้ ส่วน `cap` เป็นความจุของ array เบื้องหลังจากจุดเริ่ม slice ไม่ใช่จำนวนข้อมูลที่พร้อมอ่าน
- การสร้าง slice ด้วย `initial[:]` เป็นอีกวิธี แต่จะใช้ข้อมูลเบื้องหลังร่วมกับ array การเปลี่ยนสมาชิกอาจกระทบกัน ตัวอย่างนี้ append ค่าลง slice ใหม่ให้เห็นขั้นตอนชัด

**ลองตรวจเมื่อผิด:** อ่าน `readings[len(readings)]` จะเกินขอบเขต ตำแหน่งสุดท้ายคือ len-1 และต้องไม่ว่างก่อนอ่าน ส่วนตัวแปร value ใน range เป็นสำเนาค่า การกำหนด value ใหม่ไม่ได้แก้ readings[index]

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. เพิ่มค่า `26.5` ด้วย append อีกหนึ่งรายการ แล้วตรวจค่าเฉลี่ยและจำนวนคำเตือน
2. หลังสร้าง readings ครบแล้ว ให้เปลี่ยนเป็น slice ว่าง ก่อนเริ่มคำนวณ โปรแกรมต้องพิมพ์ no readings

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ทำไมต้องเขียน readings = append(readings, value) แทนการเรียก append แล้วทิ้งผล?

<details>
<summary>แนวคำตอบ</summary>

append คืน slice ที่อัปเดตจำนวนและอาจย้ายพื้นที่เบื้องหลัง เราจึงต้องเก็บ slice ที่คืนมาไว้ใช้งานต่อ

</details>

**นำไปใช้ต่อ:** รับข้อมูลจำนวนเปลี่ยนแปลงได้ และประมวลผลทั้งชุด

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Slice_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep7)

</details>

[EP.6](../ep06-errors/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.8](../ep08-maps/README.md)
