# EP.8 — ค้นหาและอัปเดตข้อมูลด้วย map

**เป้าหมาย:** เก็บค่าล่าสุดตามชื่ออุปกรณ์ และแยกค่าศูนย์จริงออกจากชื่อที่ไม่พบได้

**ก่อนเริ่ม:** EP.5 และ EP.7 — ฟังก์ชันและรายการข้อมูล

## ทำความเข้าใจ

**map** เก็บข้อมูลเป็นคู่ **key** กับ **value** เช่น key เป็นชื่อ sensor-01 และ value เป็นอุณหภูมิล่าสุด ไม่ต้องจำว่าอุปกรณ์อยู่ตำแหน่งไหนเหมือน slice

## ลงมือทำ

1. สร้าง map สองอุปกรณ์ แล้วอ่านค่าด้วย `latest["sensor-01"]`
2. เพิ่มชื่อใหม่และกำหนดค่าของชื่อเดิมตามตัวอย่าง สังเกตว่าชื่อเดิมถูกอัปเดต ไม่เพิ่มรายการซ้ำ
3. ค้นชื่อที่ไม่มีด้วยผลลัพธ์สองค่า value, ok แล้วทดลอง delete และ len

### ตัวอย่างเมื่อทำครบ

ไฟล์ [main.go](main.go):

```go
package main

import "fmt"

func show(latest map[string]float64, deviceID string) {
	value, ok := latest[deviceID]
	if !ok {
		fmt.Println(deviceID + ": not found")
		return
	}
	fmt.Printf("%s: %.1f C\n", deviceID, value)
}

func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	latest["sensor-01"] = 29
	latest["sensor-03"] = 31
	show(latest, "sensor-01")
	show(latest, "sensor-02")
	show(latest, "sensor-04")
	delete(latest, "sensor-03")
	fmt.Println("devices:", len(latest))
}
```

### รันและตรวจผล

งานฝึก: เปิด terminal ที่ `practics` แล้วรัน:

```shell
go run main.go
```

ถ้ารันตัวอย่างที่ให้มาโดยตรง ให้เปิด terminal ที่ `lessons/phase-01/ep08-maps` แล้วใช้ `go run .` ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd

ก่อนเปิดผลลัพธ์ ลองคาดเดาว่าข้อมูลแต่ละรายการจะถูกจัดการอย่างไร

<details>
<summary>ผลลัพธ์ที่คาดหวัง</summary>

```text
sensor-01: 29.0 C
sensor-02: 0.0 C
sensor-04: not found
devices: 2
```

</details>

<details>
<summary>อธิบายโค้ดและวิธีตรวจเมื่อผลไม่ตรง</summary>

- `map[string]float64` ระบุชนิด key และ value ต้องเตรียม map ก่อนเขียน เช่น literal ในตัวอย่าง หรือ `make(map[string]float64)`
- อ่าน key ที่ไม่มีจะได้ zero value ของชนิดนั้น จึงใช้ bool `ok` เพื่อแยกว่าไม่พบหรือมีค่าศูนย์จริง
- `delete` ลบคู่ข้อมูลของ key นั้น ถ้าไม่มี key ก็ไม่เกิด error
- เมื่อ range บน map ไม่รับประกันลำดับ จึงใช้การค้นตามชื่อในตัวอย่างเพื่อให้ผลสาธิตแน่นอน
- map นี้เก็บค่าล่าสุดเท่านั้น ถ้าต้องการประวัติหลายครั้งต่ออุปกรณ์ ต้องเปลี่ยนรูปข้อมูลในบทต่อยอด

**ลองตรวจเมื่อผิด:** `var latest map[string]float64` ยังเป็น nil map อ่านได้แต่เขียนสมาชิกจะ panic ให้สร้าง map ก่อน ส่วนการเช็ก `value == 0` แทน ok จะทำให้ข้อมูลศูนย์จริงถูกมองว่าหาย

</details>

## ฝึกเอง

ใช้ `practics/main.go` เดิม เริ่มแต่ละข้อจากตัวอย่างของบทนี้ ไม่ต้องสร้างโฟลเดอร์แยกโจทย์ หากต้องการเก็บงานเดิมให้คัดลอกเป็นไฟล์ .txt ก่อนเปลี่ยนโค้ด

1. ก่อนค้น sensor-04 ให้เพิ่มชื่อนี้ด้วยค่า 25 แล้วดูว่าจำนวนอุปกรณ์หลัง delete เป็นเท่าไร
2. หลัง delete sensor-03 ให้ลองค้นชื่อนี้อีกครั้ง ต้องรายงาน not found

ทำก่อนแล้วค่อยดู [เฉลยพร้อมเหตุผล](solutions/README.md)

**ลองอธิบาย:** ถ้า sensor-02 มีค่า 0 จะตรวจว่ามีอุปกรณ์นี้จริงอย่างไร?

<details>
<summary>แนวคำตอบ</summary>

รับ value, ok จาก map แล้วดู ok ซึ่งเป็น true สำหรับ key ที่มีอยู่ แม้ value จะเป็น 0

</details>

**นำไปใช้ต่อ:** ค้นสถานะหรือค่าล่าสุดของอุปกรณ์ตามชื่อ

<details>
<summary>อ้างอิงและผลตรวจ</summary>

อิง [เอกสาร Go ทางการ](https://go.dev/ref/spec#Map_types) การเลือกตัวอย่างและการแบ่ง EP เป็นการจัดหลักสูตรนี้ ดู [หลักการจัดลำดับ](../../../docs/REFERENCES.md)

ผลรันจริง ขอบเขต test และสิ่งที่ยังไม่ได้ทดสอบ: [ผลตรวจ Phase 1](../../../notes/phase-01-results.md#ep8)

</details>

[EP.7](../ep07-arrays-slices/README.md) · [สารบัญ Phase 1](../../../docs/playlist-01-go-basic/README.md) · [EP.9](../ep09-structs/README.md)
