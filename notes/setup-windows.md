# เตรียม Go บน Windows x64

**อัปเดต 2026-09-19: เครื่องนี้ติดตั้ง Go 1.27.1 แล้ว ไม่ต้องติดตั้งซ้ำ** ใช้ ZIP ทางการติดตั้งเฉพาะบัญชีที่ `%LOCALAPPDATA%\Programs\Go-1.27.1\go` และเพิ่มโฟลเดอร์ bin ใน User PATH เนื่องจาก MSI ต้องการสิทธิ์ติดตั้งสำหรับทุกบัญชี ดู [บันทึกการติดตั้งจริง](environment.md)

ให้เปิด PowerShell ใหม่จาก Start แล้วลอง `go version` หากเปิดผ่าน editor ที่เปิดค้างอยู่ก่อนติดตั้ง อาจต้องเปิด editor ใหม่ด้วย ขั้นตอน MSI ด้านล่างเก็บไว้สำหรับเครื่องใหม่ที่ยังไม่ได้ติดตั้ง ไม่ใช่สิ่งที่ต้องทำซ้ำใน EP.1

## ก่อนติดตั้ง

ตัวติดตั้ง MSI จะเพิ่มไฟล์โปรแกรมและอาจปรับ PATH ซึ่งเป็นรายการตำแหน่งที่ terminal ใช้ค้นหาคำสั่ง หากเคยติดตั้ง Go เอง ให้ตรวจตำแหน่งเดิมก่อนติดตั้งซ้ำ

1. เปิด [หน้าดาวน์โหลด Go ทางการ](https://go.dev/dl/)
2. เลือก stable สำหรับ **Windows / amd64 / .msi** เพราะเครื่องนี้เป็น X64 (`amd64` ใช้ได้กับ CPU x64 ทั้ง Intel และ AMD) ไม่เลือกรุ่น beta/rc หรือ ARM64
3. หน้าที่ตรวจวันที่ 2026-09-17 แสดง `go1.27.1.windows-amd64.msi` เป็น featured download หากทำในวันอื่นให้ยึดรุ่น stable ที่ทางการแสดงขณะนั้น
4. เปิด MSI แล้วทำตามหน้าจอ อ่านตำแหน่งติดตั้งและรายการเปลี่ยนแปลงก่อนยืนยัน ใช้ค่าปริยายได้ ไม่ต้องติดตั้ง Docker หรือ broker ในบทนี้
5. ปิดแล้วเปิด PowerShell ใหม่ ถ้าใช้ terminal ภายใน editor ให้เปิด editor ใหม่เมื่อยังไม่เห็น PATH ล่าสุด
6. ตรวจด้วยคำสั่งด้านล่าง ไม่ต้องเปลี่ยน execution policy

```powershell
go version
Get-Command go
git --version
```

ผล `go version` ควรมีรูปแบบ `go version go<version> windows/amd64` โดยเลขเวอร์ชันขึ้นกับที่ติดตั้งจริง

หากยังหา `go` ไม่พบ อ่าน [troubleshooting](troubleshooting.md) ก่อนแก้ PATH ด้วยตนเอง ไม่ต้องกำหนด GOPATH หรือ GOROOT เพื่อทำบทนี้

## ไปยังบทที่ 1

```powershell
Set-Location -LiteralPath 'C:\Users\kopes\Documents\L\Go\lessons\01-hello-go'
Get-Location
go version
go run ./examples/hello
```

ผลที่คาดหวัง (ยังไม่ใช่ผลรันจริงบนเครื่องนี้):

```text
Hello, Go!
```

อ่าน [บทเรียน](../lessons/01-hello-go/README.md) ก่อนทดลองแก้โค้ด

อ้างอิง: [การติดตั้ง Go ทางการ — แท็บ Windows](https://go.dev/doc/install), [รายการดาวน์โหลด](https://go.dev/dl/)
