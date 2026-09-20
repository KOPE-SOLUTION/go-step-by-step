# EP.21 — แยกงานเป็น package

**เป้าหมาย:** เรียกฟังก์ชันในโฟลเดอร์ sensor จาก main

## 1. อ่านโค้ด

Package รวมไฟล์ Go ในโฟลเดอร์เดียวกัน ใช้ `import` เพื่อเรียกจากอีก package และตั้งชื่อขึ้นต้นตัวใหญ่เพื่อให้ภายนอกเรียกได้

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import (
	"fmt"

	"example.com/go-course/basic/ep21/sensor"
)

func main() {
	fmt.Println(sensor.IsWarning(30))
}
```

ไฟล์ [sensor/reading.go](sensor/reading.go)

```go
package sensor

func IsWarning(celsius float64) bool {
	return celsius >= 30
}
```

## 2. ลองรัน

**ก่อนรัน:** ทำไมใช้ sensor.IsWarning แทนเรียก IsWarning ตรง ๆ?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep21-packages'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
true
```

ชื่อ import มาจาก module ใน go.mod ต่อด้วย `/sensor` รายละเอียด module อยู่ใน EP22

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: ถ้าเปลี่ยนชื่อฟังก์ชันเป็น isWarning ทำไมอีก package เรียกไม่ได้?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เมื่อโต๊ะเริ่มมีหลายชิ้นงาน เราแยกกลุ่มให้มีหน้าที่ชัดเจน package รวมไฟล์ Go ในโฟลเดอร์หนึ่งเพื่อให้โค้ดอื่น import มาใช้

package คือกลุ่มไฟล์ที่ทำงานร่วมกันในโฟลเดอร์; import path คือชื่อที่ระบุแหล่ง package; ชื่อที่ขึ้นต้นตัวใหญ่ เช่น IsWarning ถูก export ให้ package อื่นใช้; module คือชุด package ที่มี go.mod เป็นขอบเขต จะอ่านรายละเอียดใน EP22

- sensor/reading.go ประกาศ package sensor ส่วน examples/main.go เป็น package main คนละโฟลเดอร์
- import ใช้ชื่อ module ตาม go.mod ต่อด้วย /sensor ไม่ใช่ path ของไดรฟ์ C:
- โค้ด sensor อยู่บนเครื่องใน module เดียวกัน จึงไม่ต้องดาวน์โหลด; แยกเมื่อมีหน้าที่ชัดเจน ไม่จำเป็นต้องแยกทุกฟังก์ชัน

**ข้อผิดพลาดที่พบบ่อย**

- เปลี่ยน IsWarning เป็น isWarning แล้วเรียกจาก main: ชื่อเล็กไม่ถูก export
- เอา package sensor และ package main ไว้โฟลเดอร์เดียวกัน: แยกตามต้นไม้ไฟล์ในบท

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/basics/3)

</details>

[ตอนก่อนหน้า](../ep20-pointers/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep22-modules/README.md)
