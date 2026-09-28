# EP.11 — สลับแหล่งข้อมูลด้วย interface

**เป้าหมาย:** ใช้ฟังก์ชันเดียวอ่านตัวจำลองที่สำเร็จและล้มเหลว โดยไม่ผูกกับชนิดอุปกรณ์เดียว

**ก่อนเริ่ม:** EP.6, EP.9–10 — error, struct และ method

## ทำความเข้าใจ

**interface** ในบทนี้กำหนดว่าค่าที่นำมาใช้ต้องมี method อะไรบ้าง เราจะสร้าง interface ชื่อ `Reader` ที่กำหนด method `Read() (float64, error)` ทำให้ `show` รับได้ทั้งตัวจำลองที่คืนอุณหภูมิและตัวจำลองที่คืน error

## ลงมือทำ

1. สร้าง `FixedSensor` พร้อม method `Read` แล้วลองเรียก `Read` โดยตรงก่อน
2. สร้าง interface `Reader` แล้วเปลี่ยน `show` ให้รับ `Reader` แทน `FixedSensor`
3. เพิ่ม `FailedSensor` แล้วเรียก `show` กับตัวจำลองแต่ละตัวตามตัวอย่าง เดาก่อนว่ารายการที่สามจะยังแสดงได้หรือไม่

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

type Reader interface {
	Read() (float64, error)
}

type FixedSensor struct {
	Celsius float64
}

func (s FixedSensor) Read() (float64, error) {
	return s.Celsius, nil
}

type FailedSensor struct{}

func (s FailedSensor) Read() (float64, error) {
	return 0, fmt.Errorf("simulated read failure")
}

func show(name string, reader Reader) {
	value, err := reader.Read()
	if err != nil {
		fmt.Printf("%s: ERROR: %v\n", name, err)
		return
	}
	fmt.Printf("%s: %.1f C\n", name, value)
}

func main() {
	show("sensor-01", FixedSensor{Celsius: 27.5})
	show("sensor-02", FailedSensor{})
	show("sensor-03", FixedSensor{Celsius: 30})
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep11-interfaces` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าเมื่อ `sensor-02` คืน error โปรแกรมจะยังแสดงอุณหภูมิของ `sensor-03` หรือไม่

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 27.5 C
sensor-02: ERROR: simulated read failure
sensor-03: 30.0 C
```

</details>

<details>
<summary>อธิบายโค้ดและจุดที่ควรตรวจสอบ</summary>

- ชนิดที่มี method ตรงตาม `Reader` ใช้เป็น `Reader` ได้ทันที โดยไม่ต้องประกาศเพิ่มว่ารองรับ interface นี้
- `show` รู้เพียงว่าเรียก `Read` ได้ จึงไม่ต้องรู้รายละเอียดของแต่ละตัวจำลอง
- `struct{}` คือ struct ที่ไม่มี field ใช้เมื่อพฤติกรรมตัวอย่างไม่ต้องเก็บข้อมูล
- นี่เป็นการอ่านตามลำดับทีละตัว ยังไม่ใช่งานพร้อมกันหรือการต่ออุปกรณ์จริง
- interface เป็นทางเลือกในการออกแบบเมื่อจำเป็นต้องสลับพฤติกรรม ไม่จำเป็นต้องสร้างให้ทุก struct ตัวอย่างจึงกำหนด `Reader` ให้มีเฉพาะ method ที่ `show` ต้องเรียก

**ข้อผิดพลาดที่พบบ่อย:** ชื่อ method รวมถึงจำนวนและชนิดของ parameter กับผลลัพธ์ต้องตรงกับ interface ถ้าเปลี่ยน `Read` ของ `FixedSensor` ให้ใช้ pointer receiver ต้องส่ง `&FixedSensor{...}` ให้ `show` แทน `FixedSensor{...}`

</details>

## ฝึกเอง

ฝึกใน `practics/main.go` ไฟล์เดิม ก่อนทำแต่ละข้อให้ใส่โค้ดตัวอย่างเต็มของบทนี้ แล้วแก้ตามโจทย์ หากต้องการเก็บคำตอบข้อก่อนหน้า ให้คัดลอกเป็นไฟล์ `.txt` ก่อน

1. เปลี่ยนตัวจำลองของ `sensor-02` จาก `FailedSensor` เป็น `FixedSensor` ที่คืนค่า 28.0 โดยไม่แก้ฟังก์ชัน `show`
2. เพิ่มชนิด `OffsetSensor` ที่มี field `Base` และ `Offset` ชนิด `float64` ให้ method `Read` คืนผลรวมของสองค่าและ `nil` แล้วใช้กับ `sensor-03` โดยกำหนด `Base: 30` และ `Offset: -1.5`

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ทำไม `show` จึงรับ `FailedSensor` ได้ ทั้งที่ประกาศ parameter เป็นชนิด `Reader`?

<details>
<summary>แนวคำตอบ</summary>

`FailedSensor` มี method `Read` ตรงตามที่ `Reader` กำหนด จึงส่งให้ parameter ชนิด `Reader` ได้

</details>

**นำไปใช้ต่อ:** สลับแหล่งข้อมูลหรือที่เก็บข้อมูล และสร้างตัวจำลองสำหรับทดสอบในเฟสถัดไป

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Interface_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep11)

</details>

[EP.10](../ep10-methods-pointers/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.12](../ep12-packages-modules/README.md)
