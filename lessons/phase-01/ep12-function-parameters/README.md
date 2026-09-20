# EP.12 — แยกงานเป็นฟังก์ชันและส่งค่าเข้าไป

**เป้าหมาย:** เรียกงานเดิมด้วยชื่ออุปกรณ์ที่ต่างกัน

## 1. อ่านโค้ด

ฟังก์ชันคือชุดคำสั่งที่มีชื่อ ส่วน parameter คือตัวแปรรับค่า เช่น `name string` รับข้อความจากจุดที่เรียก

ไฟล์ [examples/main.go](examples/main.go)

```go
package main

import "fmt"

func showDevice(name string) {
	fmt.Println("Device:", name)
}

func main() {
	showDevice("sensor-01")
	showDevice("sensor-02")
}
```

## 2. ลองรัน

**ก่อนรัน:** ประกาศ showDevice ก่อน main แล้วมันจะพิมพ์เองก่อนถูกเรียกหรือไม่?

รันจากโฟลเดอร์ `lessons/phase-01/ep12-function-parameters`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
Device: sensor-01
Device: sensor-02
```

การประกาศฟังก์ชันยังไม่ทำให้มันทำงาน ต้องเรียกด้วยชื่อและส่งค่าที่ต้องการ

</details>

<a id="practice"></a>

## 3. ฝึกเอง

ใช้ **`practics/main.go` ไฟล์เดิม** เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดจาก [ตัวอย่าง EP นี้](examples/main.go) แล้วทำโจทย์ด้านล่าง

ถ้ายังไม่มีไฟล์ ให้สร้างโฟลเดอร์ `practics` ใน `Go` แล้วสร้าง `main.go` ข้างใน ไม่ต้องมี `examples` หรือ `go.mod` ดู [วิธีสร้างครั้งแรก](../../../docs/PRACTICE.md)

**แก้ `main.go` ในโฟลเดอร์ฝึก** ทีละข้อ:

1. เปลี่ยนให้เรียกด้วย gateway-01 เพียงครั้งเดียว
2. เขียน showCount(count int) ให้พิมพ์ Count: แล้วจำนวน เรียกด้วย 3

บันทึกไฟล์ (Ctrl+S) แล้วรันจาก **`practics`**:

```shell
go run main.go
```

<details>
<summary>คำถามทบทวนหลังทำโจทย์</summary>

name ใน showDevice ใช้ตรง ๆ ใน main ได้หรือไม่?

[ดูเฉลยหลังลองทำ](solutions/README.md)

</details>

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

ฟังก์ชันเหมือนขั้นตอนติดป้ายชื่อ เรากำหนดวิธีติดป้ายครั้งเดียว แล้วส่งชื่อที่จะใช้ในแต่ละครั้ง

function คือชุดคำสั่งที่ตั้งชื่อไว้; parameter คือตัวแปรรับค่าที่ประกาศในวงเล็บ; argument คือค่าที่ส่งตอนเรียก; ฟังก์ชันบทนี้ยังไม่คืนค่า

- name string ระบุว่าฟังก์ชันต้องได้รับข้อความหนึ่งค่า
- ตอนเรียกแต่ละครั้ง name รับค่าของครั้งนั้น จึงแสดงชื่อไม่เหมือนกัน
- โปรแกรมเริ่มจาก main; การประกาศฟังก์ชันเป็นการบอกวิธีทำงาน ยังไม่ใช่การสั่งให้ทำ

**ข้อผิดพลาดที่พบบ่อย**

- เรียก showDevice() โดยไม่ส่งค่า: อ่านข้อความจำนวน argument ไม่พอ
- เรียก showDevice(1): int ไม่ใช่ string ให้ส่งข้อความตามชนิดที่ประกาศ

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/tour/basics/4)

</details>

[ตอนก่อนหน้า](../ep11-loops/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep13-function-results/README.md)
