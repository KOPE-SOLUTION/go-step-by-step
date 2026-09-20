# EP.8 — จัดข้อความด้วย Printf

**เป้าหมาย:** ประกอบรายงานจากค่าหลายชนิดด้วย format ที่ตรงกัน

## 1. อ่านโค้ด

`Printf` จัดรูปแบบข้อความด้วย `%s` สำหรับ string, `%t` สำหรับ bool และ `%.1f` สำหรับทศนิยมหนึ่งตำแหน่ง ส่วน `\n` ขึ้นบรรทัดใหม่

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func main() {
	name := "sensor-01"
	temperature := 27.56
	online := true
	fmt.Printf("%s: %.1f C, online=%t\n", name, temperature, online)
}
```

## 2. ลองรัน

**ก่อนรัน:** 27.56 จะแสดงเป็น 27.5 หรือ 27.6 เมื่อใช้ %.1f?

รันจากโฟลเดอร์ `lessons/phase-01/ep08-format-output`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
sensor-01: 27.6 C, online=true
```

รูปแบบกับชนิดข้อมูลต้องตรงกัน เช่น ใช้ `%.1f` กับค่าอุณหภูมิแบบ float64

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. แสดง sensor-02: 28.25 C ด้วยทศนิยมสองตำแหน่ง
2. ประกาศ name ว่างและ count เท่ากับ 3 แสดง name="" count=3

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

Printf %.1f ทำให้ค่าตัวแปรเหลือทศนิยมหนึ่งตำแหน่งด้วยหรือไม่?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เหมือนประโยคที่เว้นช่องให้เติมชื่อกับอุณหภูมิ Printf ใช้ข้อความรูปแบบกำหนดว่าค่าแต่ละตัวจะไปตรงไหน เหมาะกับรายงานที่ต้องอ่านซ้ำหลายรายการ

**format string** คือข้อความที่มีช่องแทนค่า; `%s` string, `%d` int, `%.1f` ทศนิยมหนึ่งตำแหน่ง, `%t` bool, `%q` ข้อความแบบมี quote; `\n` ขึ้นบรรทัดใหม่

- ค่าหลัง format string จับคู่กับช่องจากซ้ายไปขวา
- %.1f เปลี่ยนรูปที่แสดง ไม่ได้เปลี่ยนค่าตัวแปร temperature
- Printf ไม่เติม newline เอง จึงใส่ \n ท้ายข้อความ

**ข้อผิดพลาดที่พบบ่อย**

- ลืม \n: prompt หรือข้อความถัดไปมาต่อบรรทัดเดียวกัน
- ใช้ %d กับ float64: ชนิดไม่ตรง มักเห็นข้อความ %! ในผลและ go vet ตรวจพบได้

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://pkg.go.dev/fmt)

</details>

[ตอนก่อนหน้า](../ep07-zero-values/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep09-arithmetic/README.md)
