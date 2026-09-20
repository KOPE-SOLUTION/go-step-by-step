# ฝึกในไฟล์เดียว

ใช้ `practics/main.go` เป็นสมุดฝึกเล่มเดียว ขึ้น EP ใหม่ก็เปลี่ยนโค้ดในไฟล์เดิม ส่วนตัวอย่างแต่ละตอนยังเปิดย้อนดูได้ใน `lessons/`

## เตรียมครั้งเดียว

1. ใน VS Code คลิกขวาที่โฟลเดอร์ `Go` → **New Folder** ตั้งชื่อ `practics` ถ้ามีแล้วใช้โฟลเดอร์เดิม
2. คลิกขวา `practics` → **New File** ตั้งชื่อ `main.go`
3. พิมพ์โค้ดตั้งต้นตามบทที่กำลังเรียนในไฟล์นี้ แล้วบันทึกด้วย **Ctrl+S**

```text
Go/
  practics/
    main.go
```

คลิกขวา `practics` → **Open in Integrated Terminal** แล้วรันจากโฟลเดอร์นี้:

```shell
go run main.go
```

ใช้ได้ทั้ง terminal ใน VS Code, PowerShell และ cmd โดย terminal ต้องอยู่ที่ `practics`

## เมื่อขึ้น EP ใหม่

เปิด `main.go` เดิม เปลี่ยนเนื้อหาทั้งไฟล์เป็นโค้ดตั้งต้นของ EP ใหม่ แล้วทำโจทย์ในหัวข้อ **3. ฝึกเอง** ไม่ต้องสร้างโฟลเดอร์แยก EP หรือโฟลเดอร์ `examples`

หากอยากเก็บงานที่ปรับเอง ให้คัดลอกไว้ก่อนเขียนทับ เช่น `main-ep02.txt` ใช้นามสกุล `.txt` เพื่อไม่ให้ Go นำมารันรวม

`practics/` ถูก Git ignore ผู้ที่ clone repository จึงสร้างพื้นที่นี้เอง งานฝึกจะไม่ติดไปกับการเพิ่มไฟล์เข้า Git ตามปกติ

<details>
<summary>เมื่อถึงบทที่ต้องมีไฟล์เพิ่ม</summary>

- **EP.1–20:** ใช้ `main.go` และ `go run main.go` ยังไม่ต้องมี `go.mod`
- **EP.21:** เพิ่ม `sensor/reading.go` เพื่อเรียน package และใช้ `go mod init example.com/go-practice` ครั้งเดียวเพื่อสร้าง `go.mod` รายละเอียดอยู่ในบท
- **EP.22 เป็นต้นไป:** ใช้ `go.mod` เดิม รันโปรแกรมด้วย `go run .` จาก `practics`
- **EP.23–24 และ EP.28:** เพิ่มหรือเขียนทับ `main_test.go` แล้วรันทดสอบด้วย `go test .`
- **EP.25:** เริ่มโปรเจกต์ใหม่ ให้เปลี่ยนเนื้อหา `main_test.go` เหลือเพียง `package main` เพราะ test ของ EP.24 ใช้กับโปรแกรมเก่า เราจะเขียน test ของโปรเจกต์ใน EP.28

ทั้งหมดใช้โฟลเดอร์ `practics` เดิม เพิ่มไฟล์เฉพาะเมื่อบทเรียนนั้นต้องใช้ ดู [คำสั่ง go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program), [การสร้าง module](https://go.dev/doc/tutorial/create-module) และ [ไฟล์ทดสอบ](https://go.dev/doc/tutorial/add-a-test)

</details>

[กลับไปฝึก EP.2](../lessons/phase-01/ep02-read-program/README.md#practice)
