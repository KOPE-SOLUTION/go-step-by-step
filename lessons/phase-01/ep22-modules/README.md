# EP.22 — อ่าน go.mod และเข้าใจขอบเขต module

**เป้าหมาย:** แยก module, package และโฟลเดอร์ที่ใช้รันคำสั่งได้

## 1. อ่านโค้ด

Module คือชุด package ที่มี `go.mod` ระบุชื่อและ Go version ขั้นต่ำ ตัวอย่างนี้ใช้ package ในเครื่องและไม่ต้องดาวน์โหลดเพิ่ม

ไฟล์ [go.mod](go.mod)

```text
module example.com/go-course/basic/ep22

go 1.22.0
```

ดูการใช้ชื่อ module ใน [main.go](examples/main.go) และ package [sensor](sensor/reading.go)

## 2. ลองรัน

**ก่อนรัน:** import ที่เริ่ม example.com ในบทนี้ต้องต่ออินเทอร์เน็ตหรือไม่?

รันจากโฟลเดอร์ `lessons/phase-01/ep22-modules`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
true
```

ชื่อ module ต่อด้วย `/sensor` ต้องตรงกับ import ใน main.go

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: หนึ่ง module มีหลาย package ได้ไหม และ fmt อยู่ใน module ของบทนี้หรือไม่?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

module เหมือนแฟ้มใหญ่ที่บอกชื่อชุดโค้ด ภายในมี package ได้หลายอัน บทนี้ดูแฟ้มที่เตรียมไว้ก่อน ยังไม่เพิ่มไลบรารีภายนอก

go.mod คือไฟล์ประกาศชื่อ module และ Go version ขั้นต่ำ; module path เป็นคำนำหน้าของ import ภายในชุดนี้; standard library คือ package ที่มากับ Go เช่น fmt; dependency คือโค้ดอีกชุดที่โปรแกรมต้องใช้

- go.mod กำหนด module example.com/go-course/basic/ep22 จึง import package ย่อยด้วยชื่อนี้ต่อ /sensor
- go 1.22.0 คือเวอร์ชันขั้นต่ำที่ module ประกาศ ไม่ใช่คำสั่งติดตั้ง; ตัวอย่างตรวจด้วย Go ที่มีในเครื่อง
- เราเลือกหนึ่ง module ต่อ EP เพื่อรันและลองแต่ละบทแยกกัน เป็นทางเลือกสำหรับหลักสูตร ไม่ใช่กฎว่าทุกแอปต้องแยกแบบนี้

**ข้อผิดพลาดที่พบบ่อย**

- รัน go run ./examples จาก Go/ ซึ่งไม่มี go.mod: เปิด terminal ที่โฟลเดอร์ EP ก่อน
- แก้ module path แล้วไม่แก้ import ภายใน: ชื่อสองฝั่งต้องตรงกัน; ไม่ต้องใช้ go get เพื่อแก้การสะกด path ในบทนี้

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/modules/managing-source)

</details>

[ตอนก่อนหน้า](../ep21-packages/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep23-first-test/README.md)
