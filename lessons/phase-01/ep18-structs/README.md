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

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep18-structs'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
sensor-01 27.5
```

ใช้จุดเพื่ออ่าน field เช่น `reading.Celsius`

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: struct ต่างจาก map[string]float64 อย่างไรในตัวอย่างการวัด?

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
