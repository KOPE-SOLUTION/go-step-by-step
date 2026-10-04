# EP.2 — เก็บข้อมูลด้วยตัวแปรและชนิดข้อมูล

**เป้าหมาย:** เก็บชื่ออุปกรณ์ ค่าการวัด และสถานะ แล้วจัดรูปแบบรายงานได้

**ก่อนเริ่ม:** EP.1 — รันโปรแกรมและอ่าน main ได้

## ทำความเข้าใจ

**ตัวแปร** คือชื่อที่ใช้อ้างถึงค่าซึ่งเปลี่ยนได้ ส่วน **ชนิดข้อมูล** ระบุประเภทของค่าที่ตัวแปรเก็บได้ เช่น `string` สำหรับข้อความ `int` สำหรับจำนวนเต็ม `float64` สำหรับเลขทศนิยม และ `bool` สำหรับ `true` หรือ `false`

เราจะเริ่มจากชื่ออุปกรณ์และอุณหภูมิ แล้วเพิ่มข้อมูลทีละส่วนก่อนจัดรูปแบบรายงาน

## ลงมือทำทีละขั้น

ใช้ `practics/main.go` ไฟล์เดิม เริ่มด้วยโค้ดขั้นที่ 1 แล้วแก้ต่อทีละขั้น หากต้องการเก็บงานเดิม ให้คัดลอกเป็น `.txt` ก่อนเปลี่ยนโค้ด บันทึกและรันจาก terminal ที่ `practics` ทุกครั้ง:

```shell
go run main.go
```

### 1. เก็บชื่ออุปกรณ์และอุณหภูมิ

เริ่ม `main.go` ของบทนี้ด้วยโค้ดต่อไปนี้:

```go
package main

import "fmt"

func main() {
	var deviceID string = "sensor-01"
	celsius := 27.5

	fmt.Println(deviceID, celsius)
}
```

`var deviceID string = ...` ประกาศตัวแปรพร้อมระบุชนิดเป็นข้อความ ส่วน `celsius := 27.5` ประกาศแบบย่อภายในฟังก์ชัน ให้ Go เลือกชนิดจากค่า ซึ่งกรณีนี้เป็น `float64`

`Println` รับหลายค่าได้ โดยพิมพ์เว้นช่องว่างระหว่างค่า

**ลองคิดก่อนรัน:** โปรแกรมจะพิมพ์ชื่อ `deviceID` หรือค่าที่ตัวแปรนี้เก็บอยู่?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01 27.5
```

</details>

### 2. เพิ่มหน่วยวัดและสถานะการเชื่อมต่อ

**ค่าคงที่** ประกาศด้วย `const` ใช้กับค่าที่ไม่เปลี่ยน ส่วน `true` และ `false` เป็นค่าจริงและเท็จของชนิด `bool`

แทนเฉพาะ `main` ทั้งฟังก์ชัน เก็บ `package` และ `import` ไว้เหมือนเดิม:

```go
func main() {
	const unit = "C"
	var deviceID string = "sensor-01"
	celsius := 27.5
	connected := true

	fmt.Println(deviceID, celsius, unit)
	fmt.Println("connected:", connected)
}
```

`unit` เก็บหน่วยวัด C ส่วน `connected` เก็บสถานะจำลองเป็น `true` ยังไม่ได้ตรวจการเชื่อมต่ออุปกรณ์จริง

**ลองคิดก่อนรัน:** ข้อความหน่วยวัดกับค่า `true` จะแสดงที่บรรทัดใด?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01 27.5 C
connected: true
```

</details>

### 3. ดูค่าของตัวแปรที่ยังไม่ได้กำหนดเอง

เพิ่มคำประกาศเหล่านี้ต่อจาก `connected := true`:

```go
	var retries int
	var label string
	var failed bool
```

แล้วเพิ่มสองคำสั่งนี้ต่อจาก `fmt.Println("connected:", connected)`:

```go
	fmt.Println("retries:", retries)
	fmt.Println("label:", label, "failed:", failed)
```

ตัวแปรที่ประกาศโดยไม่กำหนดค่าเองจะได้รับ **zero value** หรือค่าเริ่มต้นตามชนิดของมัน เราจะลองพิมพ์ดูก่อน

**ลองคิดก่อนรัน:** จำนวนเต็ม ข้อความ และ bool ที่ยังไม่ได้กำหนดค่า จะแสดงเป็นอะไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01 27.5 C
connected: true
retries: 0
label:  failed: false
```

`retries` เป็น `0`, `label` เป็นข้อความว่าง `""` และ `failed` เป็น `false` ส่วนตัวเลขชนิดทศนิยมมีค่าเริ่มต้นเป็น `0` เช่นกัน

ช่องระหว่าง `label:` กับ `failed:` ดูว่าง เพราะ `label` ไม่มีตัวอักษร ขั้นที่ 6 จะทำให้มองเห็นข้อความว่างชัดขึ้น

</details>

### 4. เปลี่ยนค่าตัวแปรเดิมด้วย =

เพิ่มบรรทัดนี้ต่อจาก `var failed bool` ก่อนคำสั่งแสดงผล:

```go
	celsius = 28.26
```

`:=` ใช้ประกาศตัวแปร ส่วน `=` ในบรรทัดนี้เปลี่ยนค่าของ `celsius` ที่ประกาศไว้แล้ว ค่าที่ใช้ตอนพิมพ์จึงเป็นค่าล่าสุด

**ลองคิดก่อนรัน:** โปรแกรมจะพิมพ์อุณหภูมิเดิม 27.5 หรือค่าใหม่ 28.26?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01 28.26 C
connected: true
retries: 0
label:  failed: false
```

</details>

**ลองใส่ค่าผิดชนิด:** เปลี่ยน `celsius = 28.26` เป็นบรรทัดนี้แล้วรัน:

```go
	celsius = "warm"
```

<details>
<summary>ตรวจสาเหตุและวิธีแก้</summary>

คอมไพล์ไม่ผ่าน โดยข้อความ error จะระบุว่าใช้ string เป็น float64 ไม่ได้ เพราะ `celsius` ถูกประกาศให้เก็บตัวเลขทศนิยมแล้ว การใช้ `=` เปลี่ยนค่าได้แต่ไม่เปลี่ยนชนิด แก้กลับเป็น `celsius = 28.26` แล้วรันยืนยันก่อนทำต่อ

</details>

### 5. จัดรูปแบบอุณหภูมิด้วย Printf

`Printf` มาจาก print formatted หมายถึงพิมพ์ตามรูปแบบที่กำหนด โดยรับข้อความรูปแบบก่อน แล้วเติมค่าตามลำดับ `%s` ใช้กับข้อความ `%.1f` แสดงเลขทศนิยมหนึ่งตำแหน่ง และ `\n` ใช้ขึ้นบรรทัดใหม่

แทนเฉพาะ `fmt.Println(deviceID, celsius, unit)` ด้วย:

```go
	fmt.Printf("%s: %.1f %s\n", deviceID, celsius, unit)
```

ค่าที่เติมตามลำดับคือ `deviceID`, `celsius` และ `unit` รูปแบบ `%.1f` ปัดค่าที่แสดงเป็นหนึ่งตำแหน่ง แต่ค่าที่เก็บในตัวแปรยังเป็น 28.26

**ลองคิดก่อนรัน:** รายงานจะเหลือทศนิยมกี่ตำแหน่ง? ถ้าไม่มี `\n` ข้อความถัดไปจะเริ่มที่ใด?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 28.3 C
connected: true
retries: 0
label:  failed: false
```

</details>

### 6. จัดรูปแบบข้อมูลที่เหลือให้ครบรายงาน

เพิ่มรูปแบบอีกสามตัว:

| รูปแบบ | ใช้แสดง |
|---|---|
| `%t` | bool เป็น `true` หรือ `false` |
| `%d` | จำนวนเต็ม |
| `%q` | ข้อความพร้อมเครื่องหมายคำพูด จึงเห็นข้อความว่างเป็น `""` |

แทนคำสั่ง `Println` สามบรรทัดที่พิมพ์ connected, retries และ label ด้วยสองบรรทัดนี้ เก็บ `Printf` แสดงอุณหภูมิไว้:

```go
	fmt.Printf("connected=%t retries=%d\n", connected, retries)
	fmt.Printf("label=%q failed=%t\n", label, failed)
```

**ลองคิดก่อนรัน:** ตอนนี้เราจะมองเห็นค่า `label` ที่เป็นข้อความว่างได้อย่างไร?

<details>
<summary>รันแล้วค่อยเปิดตรวจผล</summary>

```text
sensor-01: 28.3 C
connected=true retries=0
label="" failed=false
```

</details>

ลองเปลี่ยนเฉพาะ `%.1f` เป็น `%.2f` แล้วคาดเดาก่อนรันว่าอุณหภูมิจะแสดงอย่างไร จากนั้นเปลี่ยนกลับเป็น `%.1f` ก่อนทำโจทย์ท้ายบท

<details>
<summary>ผลเมื่อลองใช้ %.2f</summary>

```text
sensor-01: 28.26 C
connected=true retries=0
label="" failed=false
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
	const unit = "C"
	var deviceID string = "sensor-01"
	celsius := 27.5
	connected := true
	var retries int
	var label string
	var failed bool
	celsius = 28.26

	fmt.Printf("%s: %.1f %s\n", deviceID, celsius, unit)
	fmt.Printf("connected=%t retries=%d\n", connected, retries)
	fmt.Printf("label=%q failed=%t\n", label, failed)
}
```

</details>

<details>
<summary>วิธีรันตัวอย่างต้นฉบับและจุดที่ควรตรวจ</summary>

เปิด terminal ที่ `lessons/phase-01/ep02-variables-types` แล้วรัน `go run .` ผลต้องเหมือนขั้นที่ 6 ส่วนไฟล์ที่เขียนเองยังรัน `go run main.go` จาก `practics`

- ถ้าแจ้งว่าประกาศตัวแปรแล้วไม่ใช้ ให้อ่านชื่อใน error แล้วนำตัวแปรนั้นไปใช้ หรือลบคำประกาศถ้าไม่จำเป็น
- ถ้าพิมพ์ได้ข้อความอย่าง `%!d(...)` ให้ตรวจรูปแบบกับชนิดค่า เช่น `%d` ใช้กับจำนวนเต็ม ส่วน `%f` ใช้กับเลขทศนิยม

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เปลี่ยนค่า `deviceID` เป็น `"sensor-02"` แก้บรรทัด `celsius = 28.26` ให้เป็น `celsius = 26.75` และแสดงอุณหภูมิเป็นทศนิยมสองตำแหน่ง
2. ก่อนพิมพ์รายงาน กำหนด `retries = 2`, `label = "backup"` และ `failed = true` โดยไม่ประกาศตัวแปรซ้ำ

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** การพิมพ์ `28.3` ด้วย `%.1f` ทำให้ค่าที่เก็บในตัวแปร `celsius` เปลี่ยนเป็น `28.3` ด้วยหรือไม่?

<details>
<summary>แนวคำตอบ</summary>

ไม่ ค่าตัวแปรยังเป็น `28.26` ส่วน `%.1f` เปลี่ยนเฉพาะรูปแบบข้อความที่แสดง

</details>

**นำไปใช้ต่อ:** เป็นพื้นฐานข้อมูลการวัดและสถานะที่จะนำมาคำนวณและตัดสินใจ

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Variables) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep2)

</details>

[EP.1](../ep01-hello-go/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.3](../ep03-calculations-conditions/README.md)
