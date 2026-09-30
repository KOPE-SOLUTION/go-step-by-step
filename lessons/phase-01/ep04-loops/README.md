# EP.4 — ทำซ้ำและสรุปค่าการวัดด้วย for

**เป้าหมาย:** ประมวลผลหลายรอบ ข้ามรอบที่อ่านไม่ได้ และหาค่าเฉลี่ยเฉพาะรอบที่อ่านค่าได้

**ก่อนเริ่ม:** EP.3 — การคำนวณและเงื่อนไข

## ทำความเข้าใจ

**ลูป** คือการทำคำสั่งเดิมซ้ำ `for` ของ Go กำหนดค่าเริ่มต้น เงื่อนไขทำต่อ และการเปลี่ยนค่าแต่ละรอบได้ เราจะสร้างอุณหภูมิจำลองจากเลขรอบ จึงยังไม่ต้องมีอุปกรณ์หรือสุ่มตัวเลข

## ลงมือทำทีละขั้น

ใช้ `practics/main.go` ไฟล์เดิม เริ่มด้วยโค้ดขั้นที่ 1 แล้วแก้ต่อทีละขั้น บันทึกและรัน `go run main.go` จาก terminal ที่ `practics` ทุกครั้ง ก่อนดูผล ให้ลองคาดเดาสิ่งที่จะพิมพ์

### 1. เขียนลูปพิมพ์เลขรอบ

เริ่ม `main.go` ของบทนี้ด้วยโค้ดต่อไปนี้:

```go
package main

import "fmt"

func main() {
	for round := 1; round <= 5; round++ {
		fmt.Println(round)
	}
}
```

`round := 1` กำหนดค่าเริ่มต้นครั้งเดียว ก่อนทำแต่ละรอบจะตรวจ `round <= 5` และหลังทำคำสั่งในลูปจะเพิ่มค่าด้วย `round++` ซึ่งหมายถึงเพิ่มทีละ 1

**ลองคิดก่อนรัน:** โปรแกรมจะพิมพ์เลข 5 ด้วยหรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
1
2
3
4
5
```

</details>

### 2. สร้างอุณหภูมิจากเลขรอบ

แทนเฉพาะฟังก์ชัน `main` ทั้งฟังก์ชัน เก็บ `package` และ `import` ไว้เหมือนเดิม:

```go
func main() {
	for round := 1; round <= 5; round++ {
		celsius := 25.0 + float64(round)
		fmt.Printf("round %d: %.1f C\n", round, celsius)
	}
}
```

`float64(round)` แปลงเลขรอบชนิด `int` เป็น `float64` เพื่อใช้คำนวณกับค่าทศนิยม อุณหภูมินี้เป็นข้อมูลจำลองจากสูตร ยังไม่ได้อ่านอุปกรณ์

**ลองคิดก่อนรัน:** รอบแรกกับรอบสุดท้ายจะได้อุณหภูมิเท่าไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: 28.0 C
round 4: 29.0 C
round 5: 30.0 C
```

</details>

### 3. เพิ่มผลรวม total และจำนวน count

แทน `main` ด้วยเวอร์ชันนี้ เพื่อเพิ่มตัวสะสมและตรวจผลหลังจบลูป:

```go
func main() {
	total := 0.0
	count := 0
	for round := 1; round <= 5; round++ {
		celsius := 25.0 + float64(round)
		total += celsius
		count++
		fmt.Printf("round %d: %.1f C\n", round, celsius)
	}
	fmt.Printf("total: %.1f, count: %d\n", total, count)
}
```

ประกาศ `total` และ `count` ก่อนลูปเพื่อเก็บค่าต่อเนื่อง `total += celsius` ย่อจาก `total = total + celsius` ส่วน `count++` นับค่าที่นำมารวมแต่ละครั้ง

**ลองคิดก่อนรัน:** หลังจบห้ารอบ ผลรวมและจำนวนจะเป็นเท่าไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: 28.0 C
round 4: 29.0 C
round 5: 30.0 C
total: 140.0, count: 5
```

</details>

### 4. คำนวณค่าเฉลี่ยหลังจบลูป

แทนบรรทัด `fmt.Printf("total: ...")` หลังลูปด้วยบล็อกนี้:

```go
if count > 0 {
	fmt.Printf("average: %.2f C (%d readings)\n", total/float64(count), count)
}
```

ค่าเฉลี่ยคือผลรวมหารจำนวนค่า ใช้ `float64(count)` ให้ชนิดตรงกับ `total` และตรวจ `count > 0` ก่อนหาร

**ลองคิดก่อนรัน:** หากไม่มีค่าการวัดเลย เงื่อนไขนี้จะยอมให้หารหรือไม่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: 28.0 C
round 4: 29.0 C
round 5: 30.0 C
average: 28.00 C (5 readings)
```

</details>

### 5. เพิ่ม continue เพื่อข้ามรอบที่อ่านไม่ได้

ในลูป แทนบรรทัด `celsius := ...` ด้วยโค้ดนี้ เพื่อแทรกเงื่อนไขก่อนคำนวณและสะสมค่า:

```go
if round == 3 {
	fmt.Println("round 3: skipped")
	continue
}
celsius := 25.0 + float64(round)
```

`continue` ข้ามคำสั่งที่เหลือในรอบนั้น แต่ยังทำ `round++` ก่อนตรวจรอบถัดไป รอบที่ข้ามจึงไม่เพิ่มทั้ง `total` และ `count`

**ลองคิดก่อนรัน:** เมื่อข้ามรอบ 3 ต้องหารผลรวมด้วย 5 หรือ 4? อธิบายก่อนรัน

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: skipped
round 4: 29.0 C
round 5: 30.0 C
average: 28.00 C (4 readings)
```

</details>

### ตัวอย่างเมื่อทำครบ

<details>
<summary>เปิดเทียบโค้ดฉบับเต็มหลังทำครบทุกขั้น</summary>

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

func main() {
	total := 0.0
	count := 0
	for round := 1; round <= 5; round++ {
		if round == 3 {
			fmt.Println("round 3: skipped")
			continue
		}
		celsius := 25.0 + float64(round)
		total += celsius
		count++
		fmt.Printf("round %d: %.1f C\n", round, celsius)
	}
	if count > 0 {
		fmt.Printf("average: %.2f C (%d readings)\n", total/float64(count), count)
	}
}
```

</details>

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep04-loops` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคำนวณว่าเมื่อข้ามรอบ 3 จะเหลือค่าการวัดกี่ค่า และได้ค่าเฉลี่ยเท่าไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
round 1: 26.0 C
round 2: 27.0 C
round 3: skipped
round 4: 29.0 C
round 5: 30.0 C
average: 28.00 C (4 readings)
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- `round++` เพิ่มทีละหนึ่ง และ `total += celsius` ย่อจาก `total = total + celsius`
- `continue` ข้ามคำสั่งที่เหลือของรอบนั้นแล้วไปทำรอบถัดไป ส่วน `break` ออกจากลูป
- เพิ่ม `count` หลังเงื่อนไขข้ามรอบ จึงนับเฉพาะค่าที่นำมารวมใน `total` แล้วใช้จำนวนนี้หารหาค่าเฉลี่ย
- `total` และ `count` อยู่นอกลูปเพื่อเก็บค่าต่อเนื่อง ส่วน `celsius` ที่ประกาศภายในลูปใช้ได้เฉพาะภายในลูป
- Go ยังเขียน `for เงื่อนไข { ... }` ได้เมื่อไม่ต้องการส่วนเริ่มต้นและเพิ่มค่า

**ข้อผิดพลาดที่พบบ่อย:** ใช้ `< 5` จะได้ถึงรอบ 4 เท่านั้น ส่วนการประกาศ `total` ไว้ภายในลูปจะทำให้ผลรวมเริ่มใหม่ทุกครั้ง ลองพิมพ์ `total` ในแต่ละรอบเพื่อตรวจสอบ และตรวจว่า `count > 0` ก่อนหารเพื่อป้องกันการหารด้วยศูนย์

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เปลี่ยนเงื่อนไขลูปจาก `round <= 5` เป็น `round <= 3` โดยยังข้ามรอบ 3 ค่าเฉลี่ยต้องคำนวณจากสองรอบแรก
2. ให้หยุดลูปเมื่อถึงรอบ 3 โดยเปลี่ยน `continue` เป็น `break` และเปลี่ยนข้อความเป็น `round 3: stopped`

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ทำไมจึงต้องเพิ่มค่า `count` หลังเงื่อนไขข้ามรอบ?

<details>
<summary>แนวคำตอบ</summary>

เพราะเราต้องการนับค่าที่นำมารวมจริง ถ้านับรอบที่ข้ามด้วย ตัวหารจะมากเกินไปและค่าเฉลี่ยผิด

</details>

**นำไปใช้ต่อ:** เป็นพื้นฐานการอ่านหลายรายการและทำงานต่อเมื่อบางรายการใช้ไม่ได้

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#For_statements) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep4)

</details>

[EP.3](../ep03-calculations-conditions/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.5](../ep05-functions/README.md)
