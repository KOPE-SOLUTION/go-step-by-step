# Hello Go — โค้ดโปรแกรมแรก

เริ่มลงมือจาก [EP.1 — รัน Go ครั้งแรก](EP01.md) หน้านี้ใช้อ่านโค้ดประกอบเมื่อสงสัย

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

| ส่วนของโค้ด | หน้าที่ |
|---|---|
| `package main` | ระบุ package ของโปรแกรมที่สั่งรันได้ |
| `import "fmt"` | เรียกใช้เครื่องมือจัดรูปแบบและแสดงข้อความ |
| `func main()` | ฟังก์ชันเริ่มต้นของโปรแกรม |
| `fmt.Println(...)` | พิมพ์ข้อความแล้วขึ้นบรรทัดใหม่ |

Package คือกลุ่มไฟล์ Go ในโฟลเดอร์เดียวกัน ฟังก์ชันคือชุดคำสั่งที่มีชื่อ และ string คือค่าข้อความ

<details>
<summary>อ่านเพิ่มเติม</summary>

- ชื่อไฟล์ main.go เป็นชื่อที่นิยมใช้ จุดเริ่มทำงานอยู่ที่ฟังก์ชัน main ใน package main
- เครื่องหมายคำพูดกำหนดขอบเขตข้อความ จึงไม่ปรากฏในผลลัพธ์
- ตัวอย่างนี้ยืนยันว่ารันโปรแกรมได้ การอ่านอุปกรณ์จำลองจะอยู่ในเฟสถัดไป

[แบบฝึกหัดเพิ่มเติม](exercises/README.md) · [เฉลย](solutions/README.md) · [ผลตรวจจริง](tests/RESULTS.md) · [เอกสาร Go](https://go.dev/doc/tutorial/getting-started)

</details>

[กลับ EP.1](EP01.md) · [สารบัญ](../../docs/playlist-01-go-basic/README.md)
