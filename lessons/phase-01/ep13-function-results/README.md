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

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep13-function-results'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
77.0 F
```

ฟังก์ชันคำนวณคำตอบ ส่วน `main` เลือกว่าจะนำคำตอบไปแสดงอย่างไร

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: ทำไมแยกการคืนค่าจากการพิมพ์จึงช่วยการทดสอบ?

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
