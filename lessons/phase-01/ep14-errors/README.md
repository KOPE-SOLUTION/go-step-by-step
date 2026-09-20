# EP.14 — ตรวจสอบและจัดการ error

**เป้าหมาย:** แยกผลสำเร็จออกจากข้อมูลที่ใช้ไม่ได้

## 1. อ่านโค้ด

`error` ใช้บอกปัญหา และ `nil` ในที่นี้หมายถึงไม่มี error ฟังก์ชันคืนสองค่าได้ จึงรับด้วย `value, err := ...` `!=` คือไม่เท่ากับ และ `||` คืออย่างใดอย่างหนึ่งเป็นจริง

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import (
	"errors"
	"fmt"
)

func validateTemperature(value float64) (float64, error) {
	if value < 0 || value > 100 {
		return 0, errors.New("temperature outside simulated range")
	}
	return value, nil
}

func main() {
	value, err := validateTemperature(27.5)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Temperature:", value)
}
```

## 2. ลองรัน

**ก่อนรัน:** ถ้าส่ง -1 จะพิมพ์ Temperature: 0 ต่อจาก error หรือไม่?

รันจากโฟลเดอร์ `lessons/phase-01/ep14-errors`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
Temperature: 27.5
```

ตรวจ `err != nil` ก่อนใช้ผลลัพธ์ `||` คืออย่างใดอย่างหนึ่งเป็นจริง และ `return` ใน main จบโปรแกรมส่วนนี้

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. ส่ง -1 และตรวจว่ามีเพียงบรรทัด ERROR
2. ส่ง 100 แล้วตรวจว่าขอบเขตบนผ่าน

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

ทำไมรับค่ามาแล้วต้องตรวจ err ก่อน?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เครื่องตรวจข้อมูลให้ทั้งค่าและเหตุผลเมื่อไม่ผ่าน เราจะใช้ค่าต่อเมื่อผ่านการตรวจแล้ว ตัวอย่างรับเฉพาะตัวเลขจำลอง 0 ถึง 100 เท่านั้น ไม่ใช่ขอบเขตของเซนเซอร์จริงทุกชนิด

error คือชนิดค่าที่ใช้บอกปัญหา; nil ในที่นี้หมายถึงไม่มี error; errors.New สร้าง error จากข้อความ; != คือไม่เท่ากับ; || คืออย่างใดอย่างหนึ่งเป็นจริง; ฟังก์ชัน Go คืนหลายค่าได้

- (float64, error) ระบุผลสองค่า ใช้ value, err := รับทั้งคู่
- ตรวจ err ก่อนใช้ value; ค่า 0 เมื่อผิดพลาดเป็นเพียงค่าประกอบ ไม่ใช่ค่าที่อ่านสำเร็จ
- return ใน main จบฟังก์ชันทันที จึงไม่ไหลต่อไปพิมพ์ผลสำเร็จ

**ข้อผิดพลาดที่พบบ่อย**

- ละเลย err แล้วใช้ค่า 0: จะทำให้ข้อมูลเสียดูเหมือนค่าที่ใช้ได้
- ใช้ err == nil เป็นทางแสดง ERROR: ตรวจทิศทางเงื่อนไขใหม่

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/handle-errors)

</details>

[ตอนก่อนหน้า](../ep13-function-results/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep15-arrays/README.md)
