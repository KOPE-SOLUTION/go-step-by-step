# ผลตรวจจริง — บทที่ 1

## ผลล่าสุด 2026-09-19 — ผ่านด้วย Go 1.27.1

ระบบ Windows x64, `go version go1.27.1 windows/amd64` ติดตั้งจาก ZIP ทางการเฉพาะบัญชีผู้ใช้ ตรวจ SHA256 ก่อนติดตั้ง รันคำสั่งจากรากบท `Go/lessons/01-hello-go`

| คำสั่ง/กรณี | ผลจริง |
|---|---|
| `gofmt -l .` | ไม่มีไฟล์ที่ต้องจัดรูปแบบ, exit 0 ไม่แก้ source |
| `go test ./...` | compile ผ่านทั้ง 4 package, exit 0; ทุก package รายงาน `[no test files]` ไม่มี unit test function |
| `go vet ./...` | ผ่าน ไม่มีข้อความเตือน, exit 0 |
| `go run ./examples/hello` | `Hello, Go!` 1 บรรทัด, exit 0 |
| `go run ./exercises/my-hello` | starter เดิม: `Hello, Go!` 1 บรรทัด, exit 0 |
| `go run ./solutions/01-greeting` | `Hello, learner!` 1 บรรทัด, exit 0 |
| `go run ./solutions/02-two-lines` | `Hello, learner!` แล้ว `Gateway simulator is starting.` รวม 2 บรรทัด, exit 0 |
| `go build -o <TEMP>/hello.exe ./examples/hello` และเรียก exe | build ผ่าน โปรแกรมพิมพ์ `Hello, Go!`, exit 0 |
| `tests/verify.ps1` หลังปรับที่เก็บ build | ผ่านครบทุกขั้น จบ exit 0 |
| ตั้งใจเปลี่ยนเป็น `fmt.println` ในสำเนาชั่วคราว | compiler ปฏิเสธจริง: `main.go:6:6: undefined: fmt.println (but have Println)`, exit 1 ตามที่คาด |

ครั้งแรกของการตรวจผ่าน gofmt/compile/vet/run แล้ว แต่สคริปต์สร้างโฟลเดอร์ build ภายใต้ Documents ไม่ได้ จึงปรับให้เก็บ executable ใน TEMP และรันตรวจใหม่ผ่านครบ ตำแหน่ง executable ของรอบที่ผ่าน: `%TEMP%\go-course-verify-f48453e44a814bb2954056c9e2216508\hello.exe` ไม่มีการลบหรือเขียนทับไฟล์ฝึกเพื่อทดสอบ

ตั้ง `GOTOOLCHAIN=local` และ `GOPROXY=off` เฉพาะ process ทดสอบเพื่อใช้ Go ที่ติดตั้งจริงโดยไม่ดาวน์โหลด dependency ไม่บันทึกเป็นค่าระบบ การเทียบ stdout สนใจข้อความ/ลำดับ/จำนวนบรรทัด ไม่ใช่การเทียบ byte ของ CRLF/LF

**ยังไม่ได้ตรวจ:** Go รุ่นอื่น/OS อื่น, หน้าต่าง PowerShell ที่ผู้เรียนเปิดเอง, error ทุกกรณีในตาราง troubleshooting, MQTT/อุปกรณ์/เครือข่ายหรือบทถัดไป ผลนี้ไม่เท่ากับผู้เรียนผ่านบทแล้ว

## ประวัติวันที่ 2026-09-17 (ก่อนติดตั้ง ไม่ใช่สถานะล่าสุด)

วันที่เตรียม: 2026-09-17 (Asia/Bangkok)

## สิ่งที่ตรวจพบแล้ว

- Windows x64, PowerShell 7.6.5, Git 2.50.0.windows.1
- เรียก `go version` ไม่ได้: `The term 'go' is not recognized as a name of a cmdlet, function, script file, or executable program.`
- ไม่พบ Go ใน PATH และตำแหน่งติดตั้งทั่วไปที่ตรวจ รายละเอียดใน `notes/environment.md`
- ตรวจเทียบโครงสร้างโค้ดด้วยการอ่านกับเอกสาร Go ทางการแล้ว ใช้ package main, func main และ fmt.Println ไม่มี dependency ภายนอก

## การตรวจไฟล์ที่รันจริงได้โดยไม่ใช้ Go

- **ผ่าน:** PowerShell parser ตรวจ syntax ของ `tests/verify.ps1` ไม่พบ parse error
- **ผ่าน:** เรียก `tests/verify.ps1` จริง สคริปต์หยุดด้วยข้อความ `Go is not available. Follow notes/setup-windows.md, reopen PowerShell, then retry. No Go checks were run.` ตามเงื่อนไขเมื่อไม่มี Go ไม่ได้ไปถึงคำสั่งทดสอบ Go
- **ผ่าน:** ตรวจปลายทางลิงก์ Markdown ภายในชุดบทเรียน ทุกลิงก์ชี้ไปยังไฟล์ที่มีอยู่
- **ผ่าน:** ตรวจรายการไฟล์ 23 ไฟล์ มี source Go 4 ไฟล์ และ go.mod 1 ไฟล์ ทุก source มีโครงสร้าง package main / import fmt / func main การตรวจข้อความนี้ไม่ใช่การ compile
- **ผ่าน:** ตรวจซ้ำว่าโฟลเดอร์แนบ `C:\Users\kopes\Documents\L\Go` ยังว่าง ไม่ถูกแก้ไข

การตรวจข้างต้นยืนยันโครงสร้างเอกสารและการหยุดของสคริปต์เท่านั้น ไม่ยืนยันพฤติกรรมโปรแกรม Go หรือเส้นทางสำเร็จของสคริปต์

## สิ่งที่ยังไม่ได้ทดสอบ

| รายการ | สถานะ |
|---|---|
| `gofmt -l .` | ยังไม่ได้รัน — ไม่มี Go toolchain |
| `go run ./examples/hello` | ยังไม่ได้รัน — ไม่มี Go toolchain |
| รันไฟล์ฝึกและเฉลย | ยังไม่ได้รัน — ไม่มี Go toolchain |
| `go test ./...` | ยังไม่ได้รัน — ไม่มี Go toolchain; ไม่มี unit test ในบทนี้ |
| `go vet ./...` | ยังไม่ได้รัน — ไม่มี Go toolchain |
| build และรัน executable | ยังไม่ได้รัน — ไม่มี Go toolchain |
| error สะกด `Println` ผิด | ยังไม่ได้ทดสอบกับ compiler จริง |
| ความเข้ากันได้กับ Go ที่ติดตั้ง | ยังตรวจไม่ได้ ไม่มีเวอร์ชัน Go ให้เทียบ |

ผลลัพธ์ที่แสดงใน README และ EP01.md เป็นผลที่คาดหวังตามโค้ด ไม่ใช่การอ้างผลรันจริง

## บันทึกหลังติดตั้ง (ยังว่าง)

- วันที่/เวลา: __________
- `go version`: __________
- โฟลเดอร์ที่ใช้รัน: __________
- คำสั่ง: __________
- stdout / stderr / exit code จริง: __________
- ผลเทียบกับที่คาด: __________
- สิ่งที่ยังไม่ได้ตรวจ: __________
