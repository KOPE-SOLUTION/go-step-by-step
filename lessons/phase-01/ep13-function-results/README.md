# EP.13 — รับค่าที่ฟังก์ชันส่งกลับ

**เป้าหมาย:** แยกการคำนวณออกจากการแสดงผล

## 1. อ่านโค้ด

ชนิดหลังวงเล็บของฟังก์ชันบอกชนิดผลลัพธ์ `return` ส่งค่ากลับและจบการเรียกครั้งนั้น

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func main() {
	fahrenheit := celsiusToFahrenheit(25)
	fmt.Printf("%.1f F\n", fahrenheit)
}
```

## 2. ลองรัน

**ก่อนรัน:** ถ้าเรียกฟังก์ชันแต่ไม่ใช้ fmt จะเห็นคำตอบใน terminal เองไหม?

รันจากโฟลเดอร์ `lessons/phase-01/ep13-function-results`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
77.0 F
```

ฟังก์ชันคำนวณคำตอบ ส่วน `main` เลือกว่าจะนำคำตอบไปแสดงอย่างไร

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. แปลง 0 องศาเซลเซียสแล้วพิมพ์ทศนิยมหนึ่งตำแหน่ง
2. เขียน isWarning(temperature float64) bool ให้คืนค่าการเปรียบเทียบ temperature >= 30 แล้วพิมพ์ผลเมื่อส่ง 30

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

ทำไมแยกการคืนค่าจากการพิมพ์จึงช่วยการทดสอบ?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เหมือนส่งตัวเลขให้เครื่องคิดเลขแล้วรับคำตอบกลับมา ผู้เรียกเลือกเองว่าจะพิมพ์หรือเก็บคำตอบไว้

return คือส่งผลลัพธ์กลับและจบการเรียกครั้งนั้น; result type คือชนิดค่าที่เขียนหลังวงเล็บ parameter; ฟังก์ชันคืนค่าไม่จำเป็นต้องพิมพ์เอง

- float64 หลัง ) คือชนิดคำตอบ ไม่ใช่ parameter อีกตัว
- return ส่งค่ากลับมาที่ตำแหน่งเรียก จากนั้นเก็บใน fahrenheit
- แยกคำนวณกับพิมพ์เป็นแนวทางออกแบบที่ทำให้ตรวจผลคำนวณง่ายขึ้น ไม่ใช่ข้อบังคับ Go

**ข้อผิดพลาดที่พบบ่อย**

- ลืม return ในฟังก์ชันที่ต้องคืนค่า: ดู missing return
- คืนข้อความแทนตัวเลข: ผลลัพธ์ต้องตรงกับ float64 ที่ประกาศ

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/basics/6)

</details>

[ตอนก่อนหน้า](../ep12-function-parameters/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep14-errors/README.md)
