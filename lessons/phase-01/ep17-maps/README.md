# EP.17 — ค้นหาค่าด้วยชื่อใน map

**เป้าหมาย:** แยกค่าศูนย์จริงออกจากชื่อที่ไม่มีใน map

## 1. อ่านโค้ด

Map จับคู่ key กับ value เช่น ชื่ออุปกรณ์กับอุณหภูมิ `value, ok := m[key]` ให้ทั้งค่าและผลว่าพบ key หรือไม่

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func main() {
	temperatures := map[string]float64{"sensor-01": 0}
	value, ok := temperatures["sensor-01"]
	fmt.Println(value, ok)
	value, ok = temperatures["missing"]
	fmt.Println(value, ok)
}
```

## 2. ลองรัน

**ก่อนรัน:** ค่าที่ค้นได้เป็น 0 ทั้งคู่ หมายความว่าพบอุปกรณ์ทั้งคู่หรือไม่?

รันจากโฟลเดอร์ `lessons/phase-01/ep17-maps`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
0 true
0 false
```

ถ้าไม่พบจะได้ค่าศูนย์ จึงต้องดู `ok` เพื่อแยกจากค่าศูนย์ที่มีอยู่จริง

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. สร้าง map ว่าง เพิ่ม sensor-02 ค่า 28.5 แล้วค้นและพิมพ์พร้อม ok
2. สร้าง map ของ sensor-01 ค่า 25 ลบด้วย delete แล้วค้นพร้อม ok

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

ทำไมการทดสอบไม่ควรอิงลำดับที่ range map ให้มา?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

map เหมือนสมุดที่ค้นค่าจากป้ายชื่อ เช่น ชื่ออุปกรณ์ แทนการจำว่ามันอยู่ช่องที่เท่าไร

map จับคู่ key กับ value; key ในตัวอย่างเป็น string; value เป็น float64; value, ok := m[key] ให้ทั้งค่ากับ bool ว่าพบหรือไม่; delete ลบคู่ข้อมูลตาม key

- map literal ที่มี { } สร้าง map พร้อมใช้งาน เขียนเพิ่มได้ด้วย m[key] = value
- ถ้าไม่พบ key จะได้ค่าศูนย์ของชนิด value จึงต้องดู ok เมื่อต้องแยกกรณีนี้
- Go ไม่รับประกันลำดับการวน map; ตัวอย่างนี้ค้นด้วย key ตรง ๆ เพื่อให้ผลแน่นอน

**ข้อผิดพลาดที่พบบ่อย**

- ดูเฉพาะ value == 0 แล้วบอกว่าไม่พบ: ค่าจริงอาจเป็นศูนย์
- ประกาศ var m map[string]int แล้วเขียนทันที: map ที่เป็น nil ต้องสร้างก่อน เช่น map[string]int{}

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/moretypes/19)

</details>

[ตอนก่อนหน้า](../ep16-slices/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep18-structs/README.md)
