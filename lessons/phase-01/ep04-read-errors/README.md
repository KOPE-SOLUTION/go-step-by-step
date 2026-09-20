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

รันจากโฟลเดอร์ `lessons/phase-01/ep04-read-errors`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
Check complete
```

Go แยกตัวอักษรใหญ่กับเล็ก: `fmt.Println` ใช้ได้ แต่ `fmt.println` ใช้ไม่ได้

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. เปลี่ยน `fmt.Println` เป็น `fmt.println` ลองรันและอ่านข้อความผิดพลาด แล้วแก้กลับให้แสดง `Check complete`
2. ลบ `import "fmt"` ชั่วคราว ลองรันและอ่าน error จากนั้นใส่กลับ แล้วเปลี่ยนข้อความเป็น `Fixed`

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

go ไม่เป็นที่รู้จัก กับ undefined: fmt.println ต่างกันที่ขั้นไหน?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

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
