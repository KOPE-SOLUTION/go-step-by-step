# EP.11 — สลับแหล่งข้อมูลด้วย interface

**เป้าหมาย:** ใช้ฟังก์ชันเดียวอ่านตัวจำลองที่สำเร็จและล้มเหลว โดยไม่ผูกกับชนิดอุปกรณ์เดียว

**ก่อนเริ่ม:** EP.6, EP.9–10 — error, struct และ method

## ทำความเข้าใจ

**interface** ในบทนี้เป็นข้อตกลงว่าค่าที่ส่งเข้ามาต้องเรียก method อะไรได้ Reader ต้องมี `Read() (float64, error)` จึงใช้ได้ทั้งตัวจำลองที่คืนค่าแน่นอนและตัวจำลองที่คืน error

## ลงมือทำ

1. สร้าง FixedSensor พร้อม Read และลองเรียก Read โดยตรงก่อน
2. สร้าง Reader แล้วเปลี่ยน show ให้รับ Reader แทน FixedSensor
3. เพิ่ม FailedSensor และเรียก show ต่อเนื่องตามตัวอย่าง เดาก่อนว่ารายการที่สามจะยังแสดงได้หรือไม่

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

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 27.5 C
sensor-02: ERROR: simulated read failure
sensor-03: 30.0 C
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- ชนิดใดมี method ตรงตาม Reader ก็ใช้เป็น Reader ได้โดยไม่ต้องเขียนคำว่า implements
- show รู้เพียงว่าเรียก Read ได้ จึงไม่ต้องรู้รายละเอียดของแต่ละตัวจำลอง
- `struct{}` คือ struct ที่ไม่มี field ใช้เมื่อพฤติกรรมตัวอย่างไม่ต้องเก็บข้อมูล
- นี่เป็นการอ่านตามลำดับทีละตัว ยังไม่ใช่งานพร้อมกันหรือการต่ออุปกรณ์จริง
- interface เป็นทางเลือกในการออกแบบเมื่อจำเป็นต้องสลับพฤติกรรม ไม่จำเป็นต้องสร้างให้ทุก struct ตัวอย่างกำหนดไว้ฝั่ง show ซึ่งเป็นผู้ใช้งาน

**ลองตรวจเมื่อผิด:** ชื่อ method ชนิด parameter และชนิดผลลัพธ์ต้องตรงกันทั้งหมด ถ้า Read ใช้ pointer receiver ต้องส่ง pointer ของชนิดนั้นให้ Reader เช่น &FixedSensor{...} การย่อเวลาเรียก method ตรง ๆ ไม่ได้ทำให้ค่าธรรมดาตรงกับ interface เสมอไป

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. ให้ sensor-02 ใช้ FixedSensor ค่า 28.0 แทน FailedSensor โดยไม่แก้ show
2. เพิ่ม OffsetSensor มี field Base และ Offset ชนิด float64 ให้ Read คืนผลรวม แล้วใช้เป็น sensor-03 ด้วย Base=30, Offset=-1.5

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** เหตุใด show จึงใช้ FailedSensor ได้ทั้งที่ไม่ได้รับ parameter ชนิด FailedSensor โดยตรง?

<details>
<summary>แนวคำตอบ</summary>

FailedSensor มี Read ที่ตรงกับข้อตกลง Reader จึงส่งผ่าน interface ได้

</details>

**นำไปใช้ต่อ:** สลับแหล่งข้อมูลหรือที่เก็บข้อมูล และสร้างตัวจำลองสำหรับทดสอบในเฟสถัดไป

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Interface_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep11)

</details>

[EP.10](../ep10-methods-pointers/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.12](../ep12-packages-modules/README.md)
