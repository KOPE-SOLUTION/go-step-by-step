# EP.18 — รวมข้อมูลหนึ่งรายการด้วย struct

**เป้าหมาย:** สร้างชนิดข้อมูลที่มีชื่ออุปกรณ์และอุณหภูมิ

## 1. อ่านโค้ด

`struct` รวมข้อมูลหลายช่องไว้เป็นชนิดเดียว แต่ละช่องเรียกว่า field เช่น ชื่ออุปกรณ์กับอุณหภูมิ

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 27.5}
	fmt.Println(reading.DeviceID, reading.Celsius)
}
```

## 2. ลองรัน

**ก่อนรัน:** ถ้าระบุเฉพาะ DeviceID ช่อง Celsius จะมีค่าอะไร?

รันจากโฟลเดอร์ `lessons/phase-01/ep18-structs`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
sensor-01 27.5
```

ใช้จุดเพื่ออ่าน field เช่น `reading.Celsius`

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. สร้าง Reading ที่ DeviceID เป็น sensor-02 โดยไม่ใส่ Celsius แล้วพิมพ์ทั้งสองช่อง
2. คัดลอก Reading ค่า 25 แล้วแก้ Celsius ของสำเนาเป็น 30 พิมพ์ Celsius ของทั้งสอง

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

struct ต่างจาก map[string]float64 อย่างไรในตัวอย่างการวัด?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

struct เหมือนแบบฟอร์มหนึ่งใบ มีช่องชื่อกับช่องค่าการวัด เมื่อส่งหนึ่งใบจะรู้ว่าข้อมูลส่วนต่าง ๆ เป็นของรายการเดียวกัน

type ใช้ประกาศชนิด; struct คือกลุ่ม field; field คือช่องข้อมูลที่มีชื่อและชนิด; Reading{...} เป็นการสร้างค่าชนิด Reading

- reading.DeviceID อ่าน field ของค่าที่เก็บใน reading
- ใช้ชื่อ field ใน literal ช่วยให้เห็นความหมาย และไม่ต้องจำลำดับ เป็นทางเลือกด้านความอ่านง่าย
- struct เป็นค่า การกำหนดให้อีกตัวจะคัดลอก field; ถ้า field ภายในเป็น slice/map ต้องคิดเรื่องข้อมูลร่วมกันเพิ่ม แต่บทนี้มีแค่ string/float64

**ข้อผิดพลาดที่พบบ่อย**

- ใช้ reading.Celcius สะกดผิด: Go แยกชื่อ field ตามตัวอักษร
- ใส่ "27.5" ให้ Celsius: ข้อความไม่ใช่ float64

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/moretypes/2)

</details>

[ตอนก่อนหน้า](../ep17-maps/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep19-methods/README.md)
