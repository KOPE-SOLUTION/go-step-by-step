# เฉลย EP.22

เปิดหลังลอง [แบบฝึกหัด](../exercises/README.md) คำสั่งด้านล่างรันจากโฟลเดอร์ต้นฉบับ `lessons/phase-01/ep22-modules`

## ข้อ 1

module คือชื่อใน go.mod; sensor กับ examples เป็นคนละ package ภายใน module

ดู [main.go](01/main.go)

```powershell
go run ./solutions/01
```

ผล:

```text
false
```

## ข้อ 2

การเปลี่ยน input ไม่เปลี่ยนชื่อชุดโค้ด

ดู [main.go](02/main.go)

```powershell
go run ./solutions/02
```

ผล:

```text
true
```

<details>
<summary>คำตอบคำถามทบทวน</summary>

หนึ่ง module มีหลาย package ได้ไหม และ fmt อยู่ใน module ของบทนี้หรือไม่?

ได้; sensor กับ examples เป็นตัวอย่าง ส่วน fmt เป็น standard library ที่มากับ Go ไม่ใช่ package ที่เราเขียนใน module นี้

</details>

[กลับบทเรียน](../README.md)
