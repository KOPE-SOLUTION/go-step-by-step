# EP.7 — จัดการรายการข้อมูลด้วย array และ slice

**เป้าหมาย:** เก็บหลายค่าการวัด เพิ่มรายการ และสรุปค่าเฉลี่ยกับจำนวนคำเตือน

**ก่อนเริ่ม:** EP.4 — ลูป และ EP.5 — ฟังก์ชัน

## ทำความเข้าใจ

**array** เก็บข้อมูลชนิดเดียวกันจำนวนคงที่ เช่น `[3]float64` ส่วน **slice** อ้างถึงข้อมูลช่วงหนึ่งใน array และใช้จัดการรายการที่เพิ่มจำนวนได้ เขียนชนิดเป็น `[]float64`

**index** คือตำแหน่ง เริ่มจาก 0 ส่วน `len` บอกจำนวนสมาชิกปัจจุบัน

## ลงมือทำ

1. สร้าง `initial` แบบ array แล้วอ่าน `initial[0]` และ `len(initial)`
2. คัดลอกค่าจาก `initial` มาเก็บใน slice ชื่อ `readings` ด้วย `range` และ `append` แล้วเพิ่มค่า 32 ตามตัวอย่าง
3. ใช้ `range` สรุปค่าเฉลี่ยและนับคำเตือน ลองทำ slice ว่างแล้วตรวจว่าส่วนค่าเฉลี่ยทำงานอย่างไร

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

ก่อนเปิดผลลัพธ์ ลองคำนวณว่าเมื่อเพิ่มค่า 32 จะมีค่าการวัดทั้งหมดกี่ค่า ค่าเฉลี่ยเท่าไร และมีกี่ค่าที่ถึงเกณฑ์เตือน

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
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `[3]float64` มีจำนวน 3 เป็นส่วนหนึ่งของชนิด ส่วน `[]float64{}` สร้าง slice ว่าง
- `range` ให้ index กับค่าของแต่ละรายการ ถ้าไม่ต้องใช้ index เขียน `_` เพื่อทิ้งได้
- ต้องรับผลจาก `append` กลับมา เพราะ slice ที่ได้อาจอ้างถึงพื้นที่เก็บใหม่
- `len` บอกจำนวนสมาชิกปัจจุบัน ส่วน `cap` บอกความจุของ array ที่รองรับ slice โดยนับจากจุดเริ่มต้นของ slice จึงอาจมากกว่า `len` ได้
- การสร้าง slice ด้วย `initial[:]` เป็นอีกวิธี แต่จะใช้ข้อมูลเบื้องหลังร่วมกับ array การเปลี่ยนสมาชิกอาจกระทบกัน ตัวอย่างนี้จึงใช้ `append` คัดลอกแต่ละค่าลง slice ใหม่

**ข้อผิดพลาดที่พบบ่อย:** อ่าน `readings[len(readings)]` จะเกินขอบเขต ตำแหน่งสุดท้ายคือ `len(readings)-1` และต้องตรวจว่า slice ไม่ว่างก่อนอ่าน ส่วนตัวแปร `value` ใน `range` เป็นสำเนาค่า การกำหนดค่าใหม่ให้ `value` จึงไม่ได้แก้สมาชิกใน `readings`

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เพิ่มค่า `26.5` ด้วย `append` อีกหนึ่งรายการ แล้วตรวจค่าเฉลี่ยและจำนวนคำเตือน
2. หลังเพิ่มค่า 32 แล้ว ให้กำหนด `readings = []float64{}` ก่อนเริ่มคำนวณ โปรแกรมต้องพิมพ์ `no readings`

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ทำไมต้องเก็บค่าที่ `append` คืนมา เช่น `readings = append(readings, value)`?

<details>
<summary>แนวคำตอบ</summary>

`append` คืน slice ที่มีจำนวนสมาชิกใหม่ และอาจอ้างถึง array ที่สร้างขึ้นใหม่เมื่อความจุเดิมไม่พอ จึงต้องเก็บ slice ที่คืนมาเพื่อใช้สมาชิกที่เพิ่มเข้าไป

</details>

**นำไปใช้ต่อ:** รับข้อมูลจำนวนเปลี่ยนแปลงได้ และประมวลผลทั้งชุด

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Slice_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep7)

</details>

[EP.6](../ep06-errors/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.8](../ep08-maps/README.md)
