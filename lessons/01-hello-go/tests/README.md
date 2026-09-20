# ตรวจบทที่ 1

**ต้องมี Go ก่อน** คำสั่งทั้งหมดเริ่มจาก `Go/lessons/01-hello-go` และต้องแก้ error ที่ตั้งใจทำในแบบฝึกหัดข้อ 3 กลับแล้ว

การตรวจมี 3 ชั้น: รูปแบบและการ compile, การรันเทียบข้อความ, และ build เป็นไฟล์โปรแกรม ไม่มี unit test function ในบทนี้ เพราะโปรแกรมยังไม่มีตรรกะที่ควรทดสอบแยกส่วน

## วิธีอัตโนมัติสำหรับ Windows PowerShell

```powershell
& ./tests/verify.ps1
```

สคริปต์ตรวจ Go/gofmt ก่อน จากนั้นตรวจรูปแบบโดยไม่แก้ source, สั่ง go test เพื่อ compile package, go vet เพื่อตรวจปัญหาที่เครื่องมือพบได้, รันตัวอย่างและเฉลย 3 โปรแกรมเทียบข้อความและจำนวนบรรทัด แล้ว build/run ตัวอย่างในโฟลเดอร์ `%TEMP%/go-course-verify-<รหัสไม่ซ้ำ>` ไม่เขียนทับไฟล์ executable เก่า ไม่ลบไฟล์ ไม่ติดตั้งเครื่องมือ ไม่เปลี่ยนการตั้งค่าระบบ ใช้ TEMP เพราะเครื่องนี้ไม่ยอมให้คำสั่งสร้างไฟล์ build ภายใต้ Documents

การเทียบข้อความไม่สนใจความต่าง CRLF/LF ของระบบ แต่สนใจตัวใหญ่เล็กและจำนวนบรรทัด ไม่ตรวจ byte ทุกตัว สคริปต์ไม่กำหนดผลข้อความของไฟล์ฝึกเพราะผู้เรียนต้องแก้เอง อย่างไรก็ตาม `go test ./...` จะตรวจว่าไฟล์ฝึก compile ได้ด้วย

ถ้า execution policy บล็อกสคริปต์ ใช้วิธีทีละคำสั่งด้านล่าง ไม่ต้องปรับ policy

## ตรวจทีละคำสั่ง

```powershell
go version
gofmt -l .
go test ./...
go vet ./...
go run ./examples/hello
go run ./solutions/01-greeting
go run ./solutions/02-two-lines
```

หลังแต่ละคำสั่ง Go ใช้ `$LASTEXITCODE` เพื่อตรวจสถานะสำเร็จ (`0`) และอ่าน error ก่อนทำต่อ

ผลที่คาดหวัง:

- `gofmt -l .`: ไม่มีรายชื่อไฟล์ ถ้ามีชื่อให้เปิดดู แล้วใช้ `gofmt -w <ไฟล์ที่ตั้งใจแก้>` เมื่อต้องการจัดรูปแบบจริง `-w` จะเขียนไฟล์
- `go test ./...`: exit code 0 และรายงาน `[no test files]` สำหรับ package ต่าง ๆ **ไม่ได้หมายความว่ามี unit test ผ่านแล้ว**
- `go vet ./...`: ไม่มีข้อความเตือนและ exit code 0
- โปรแกรมตัวอย่าง: `Hello, Go!`
- เฉลยข้อ 1: `Hello, learner!`
- เฉลยข้อ 2: `Hello, learner!` ตามด้วย `Gateway simulator is starting.` คนละบรรทัด

ขั้นเพิ่มเติม: build และเรียกโปรแกรมที่ได้ โดยสร้างโฟลเดอร์ใหม่ทุกครั้ง:

```powershell
$lessonBuildDir = Join-Path ([IO.Path]::GetTempPath()) ('go-course-build-' + [guid]::NewGuid().ToString('N'))
New-Item -ItemType Directory -Path $lessonBuildDir | Out-Null
$lessonExe = Join-Path $lessonBuildDir 'hello.exe'
go build -o $lessonExe ./examples/hello
if ($LASTEXITCODE -eq 0) { & $lessonExe }
```

คาดว่าได้ `Hello, Go!` หนึ่งบรรทัด การรันครั้งแรกอาจช้ากว่าครั้งต่อไปเพราะ Go ต้อง compile และสร้าง cache ตามการทำงานปกติของเครื่องมือ

บันทึกเวอร์ชัน คำสั่ง ผลจริง และข้อจำกัดใน [RESULTS.md](RESULTS.md) อย่าใช้รายการผลที่คาดหวังด้านบนแทนหลักฐาน
