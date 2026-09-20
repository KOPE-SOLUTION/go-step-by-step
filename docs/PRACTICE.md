# ใช้ practics สำหรับฝึกโดยไม่ส่งงานฝึกเข้า Git

ใช้ชื่อ **practics** ตามที่ขอ โฟลเดอร์นี้มีหน้าที่เหมือน `practice` ในหลักสูตร Java อย่าผสมสองชื่อในคำสั่งของหลักสูตร Go

| ตำแหน่ง | หน้าที่ | ตั้งใจเก็บใน Git |
|---|---|---|
| `docs/` | สารบัญ Playlist แผน และวิธีฝึก | ใช่ |
| `lessons/` | บทเรียน ตัวอย่าง starter และเฉลย | ใช่ |
| `projects/` | โปรเจกต์อ้างอิงเมื่อเริ่มทำจริง | ใช่ |
| `practics/` | โค้ดที่ลองเอง ผลทดลอง และบันทึกส่วนตัว | ไม่ใช่ — ignore ทั้งโฟลเดอร์ |

`lessons/01-hello-go/exercises/my-hello` เป็น starter ที่แชร์ให้คนอื่นได้ ส่วนไฟล์ที่คุณแก้ฝึกตั้งแต่นี้ให้ใช้ `practics/phase-01/ep01-hello-go/main.go` ซึ่งสร้างแยกไว้แล้ว งานในตำแหน่งเดิมไม่ได้ถูกลบหรือย้าย

## ใช้งานบนเครื่องนี้

เปิด PowerShell แล้วเลือกโฟลเดอร์ฝึก:

```powershell
Set-Location -LiteralPath 'C:\Users\kopes\Documents\Codex\2026-09-17\go-youtube-go-iot-edge-gateway\Go\practics\phase-01\ep01-hello-go'
go run .
```

`.` หมายถึงโฟลเดอร์ปัจจุบัน จึงรันโปรแกรมในพื้นที่ฝึก ผลเริ่มต้นคือ `Hello, Go!` ถ้าคุณแก้ข้อความแล้ว ผลก็เปลี่ยนตามที่เขียน เก็บการคาดเดาและผลจริงก่อนเปิดเฉลย

พื้นที่ฝึกมี go.mod แยกจากตัวอย่าง แม้ลองทำผิดจน compile ไม่ผ่าน ก็ไม่ทำให้การตรวจต้นฉบับใน lessons ล้มตาม และไม่ต้องมี go.work หรือ module กลาง

## เมื่อต้องสร้างพื้นที่ฝึกในเครื่องใหม่

Git ไม่ส่ง `practics` ไปกับการ clone จึงเก็บวิธีสร้างไว้ในเอกสารนี้ หลังติดตั้ง Go แล้วเปิด PowerShell ที่ราก `Go` ใช้คำสั่งต่อไปนี้ครั้งเดียว **คำสั่งหยุดถ้าโฟลเดอร์มีอยู่ เพื่อไม่ทับงานเดิม**:

```powershell
$practiceTarget = Join-Path (Get-Location).Path 'practics/phase-01/ep01-hello-go'
if (Test-Path -LiteralPath $practiceTarget) {
    throw 'Practice folder already exists. Open your existing work; do not overwrite it.'
}
New-Item -ItemType Directory -Path $practiceTarget -ErrorAction Stop | Out-Null
Copy-Item -LiteralPath './lessons/01-hello-go/examples/hello/main.go' -Destination (Join-Path $practiceTarget 'main.go') -ErrorAction Stop
Push-Location -LiteralPath $practiceTarget
try {
    go mod init example.com/go-course/practics/hello
    if ($LASTEXITCODE -ne 0) { throw 'Could not create go.mod. Keep the error for diagnosis.' }
    go run .
} finally {
    Pop-Location
}
```

เครื่องนี้เตรียมโฟลเดอร์ให้แล้วจึง **ไม่ต้องรันชุดสร้างซ้ำ** คำสั่ง New-Item/Copy-Item ภายใน Documents เคยถูก Windows ปฏิเสธในการเตรียมเครื่อง หากเกิดซ้ำให้หยุดและส่ง error ไม่ต้องปิดระบบป้องกัน การสร้างผ่าน editor เป็นอีกทางเลือกที่ต้องตรวจว่าไม่ทับไฟล์เดิม

## Git ignore ทำอะไรและไม่ทำอะไร

`Go/.gitignore` มีบรรทัด `/practics/` จึงไม่เพิ่มไฟล์ใหม่ในพื้นที่ฝึกด้วย `git add` ตามปกติ ไฟล์เหล่านี้จึงไม่อยู่ใน commit ที่จะ push ส่วนคำสั่ง push ส่ง commit ไม่ได้อ่านทุกไฟล์จากโฟลเดอร์โดยตรง

เริ่มใช้ Git repository จริงในโฟลเดอร์ Go เมื่อ 2026-09-20 บน branch `main` ยังไม่มี remote และไม่ได้ push งานออกไปภายนอก ผลตรวจครั้งก่อนใน repository ชั่วคราวเก็บไว้เป็นประวัติใน notes/structure-check.md

ตรวจได้จากราก Go:

```powershell
git check-ignore -v --no-index practics/phase-01/ep01-hello-go/main.go
git ls-files -- practics/
```

คำสั่งแรกควรบอกกฎ `/practics/` คำสั่งที่สองควรไม่มีไฟล์ที่ถูกติดตาม หากเคยติดตามมาก่อน .gitignore จะไม่เอาไฟล์นั้นออกให้อัตโนมัติ ให้ตรวจสถานะก่อนจัดการ ไม่ลบไฟล์ฝึกเอง และไม่ใช้ `git add -f` กับพื้นที่นี้

การ ignore ไม่ใช่การล็อกห้ามอัปโหลด: ยังบังคับเพิ่มไฟล์ได้ หรือส่งผ่าน ZIP ได้ จึงไม่ควรเก็บรหัสผ่าน/ข้อมูลจริงในงานฝึก หากจัดแพ็กเผยแพร่ในอนาคตให้เลือกจากไฟล์ที่ตั้งใจเผยแพร่และตรวจว่าไม่มี practics ปะปน

อ้างอิง: [Git — gitignore](https://git-scm.com/docs/gitignore)

## สร้างพื้นที่ฝึก EP2–28 ทีละตอน

ใช้ PowerShell 7 ที่ติดตั้งในเครื่องนี้ จากรากโฟลเดอร์ Go รันตัวอย่างสำหรับ EP2:

```powershell
./scripts/new-practice.ps1 -Episode 2
Set-Location -LiteralPath './practics/phase-01/ep02-read-program'
go run ./examples
```

เปลี่ยนหมายเลข Episode เป็นตอนที่กำลังเรียน สคริปต์บอกตำแหน่งโฟลเดอร์ที่สร้าง คัดลอก go.mod, examples และ package ย่อยที่จำเป็น เช่น sensor แต่ไม่คัดลอกเฉลย ถ้า target มีอยู่แล้วจะหยุดก่อนเขียน จึงให้เปิดงานเดิมแทนการรันสร้างซ้ำ

ไฟล์ examples/main.go เป็นจุดเริ่มแก้ของทุก EP ใหม่ เมื่อมีไฟล์อื่นประกอบให้เก็บไว้ครบ รันทั้ง package ด้วย go run ./examples ไม่ใช่รัน main.go เพียงไฟล์เดียว ส่วน EP1 ที่เตรียมไว้เดิมยังใช้ go run . ตามคำสั่งก่อนหน้า

หากเพิ่งสร้าง EP1 ด้วยสคริปต์ใหม่นี้ จะได้ examples/hello และใช้ go run ./examples/hello ตามที่สคริปต์บอก โฟลเดอร์ EP1 ที่มีอยู่จะไม่ถูกแปลงหรือเขียนทับ

ถ้าระบบไม่ยอมรันสคริปต์หรือเขียนโฟลเดอร์ ให้เก็บ error ไม่ต้องเปลี่ยนการตั้งค่าความปลอดภัย ใช้ editor/ตัวจัดการไฟล์สร้างโฟลเดอร์ชื่อใหม่ใน practics แล้วคัดลอก go.mod กับ examples (และ sensor สำหรับ EP21–22) ด้วยตนเองโดยตรวจว่าไม่ทับงานเดิม

รายละเอียดการตรวจสคริปต์และข้อจำกัดของเครื่อง: [ผลตรวจ Phase 1](../notes/phase-01-results.md)
