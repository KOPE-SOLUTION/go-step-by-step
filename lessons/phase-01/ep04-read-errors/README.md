# EP.4 — อ่านข้อผิดพลาดจากการสะกดชื่อ

**เป้าหมาย:** แยกการหาเครื่องมือไม่พบจาก compile error และแก้ทีละจุด

## 1. อ่านโค้ด

Compiler คือเครื่องมือที่ตรวจและแปลงโค้ดก่อนรัน หากสะกดชื่อผิด โปรแกรมจะยังไม่เริ่มทำงาน

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func main() {
	fmt.Println("Check complete")
}
```

## 2. ลองรัน

**ก่อนรัน:** ถ้าเขียน fmt.println โปรแกรมจะพิมพ์ก่อนแล้วค่อยแจ้ง error หรือยังไม่เริ่มทำงาน?

จากโฟลเดอร์หลักสูตรที่มี `lessons` เปิด PowerShell แล้วใช้:

```powershell
Set-Location -LiteralPath './lessons/phase-01/ep04-read-errors'
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
Check complete
```

Go แยกตัวอักษรใหญ่กับเล็ก: `fmt.Println` ใช้ได้ แต่ `fmt.println` ใช้ไม่ได้

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: go ไม่เป็นที่รู้จัก กับ undefined: fmt.println ต่างกันที่ขั้นไหน?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

compiler เหมือนคนตรวจคำสั่งก่อนเริ่มงาน ถ้าพบชื่อที่ใช้ไม่ได้ โปรแกรมยังไม่ถูกสั่งทำงาน ให้เริ่มอ่านชื่อไฟล์และเลขบรรทัดของ error ก่อนเดาว่าเครื่องมือเสีย

**compile** คือการแปลงและตรวจ source ก่อนรัน; **error** คือข้อความแจ้งปัญหา; เลขหลังชื่อไฟล์มักบอกบรรทัดและคอลัมน์

- ตัวอย่างที่ให้เป็นโค้ดถูกต้อง เก็บไว้เทียบกับงานฝึก
- ในพื้นที่ฝึกเปลี่ยนเป็น `fmt.println` เพียงจุดเดียว บันทึกแล้วใช้คำสั่งรันเดิม
- มองหา `undefined: fmt.println` และเทียบกับ `Println` ก่อนแก้กลับ; ข้อความต่อท้ายอาจต่างตามรุ่น

**ข้อผิดพลาดที่พบบ่อย**

- เปลี่ยนหลายจุดก่อนรัน: ทำให้แยกสาเหตุยาก ให้กลับมาเทียบทีละบรรทัด
- แก้ชื่อแต่ยังไม่บันทึก: compiler จะอ่านไฟล์ที่บันทึกอยู่

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/getting-started)

</details>

[ตอนก่อนหน้า](../ep03-output-lines/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep05-variables/README.md)
