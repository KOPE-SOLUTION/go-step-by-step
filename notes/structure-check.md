# ผลตรวจการแบ่ง Phase และพื้นที่ฝึก

วันที่ 2026-09-19

## สิ่งที่ทำ

อ้างอิง README, docs/FUTURE_ROADMAP.md, สารบัญ Playlist และ .gitignore ของหลักสูตร Java ที่ระบุในแช็ต “จัดทำ README Java Case Study” ใช้เป็นแบบรูปแบบเอกสาร ไม่คัดลอกเทคโนโลยี Java หรือเปลี่ยนงานในหลักสูตรนั้น

จัด Go เป็น 5 Phase เพิ่ม docs/playlist-01-go-basic, docs/FUTURE_ROADMAP.md และ docs/PRACTICE.md เตรียม practics แยก module สำหรับ Hello Go ไม่มีการย้ายหรือลบไฟล์บทเรียนเดิม ไม่มีการแก้ main.go เดิม

## ผลตรวจจริง

| รายการ | ผล |
|---|---|
| ลิงก์ Markdown ใน Go/ และ outputs/ รวม heading fragment | ผ่าน ณ เวลาตรวจ |
| `git check-ignore --no-index -v` กับ practics/README.md, main.go, go.mod | ทั้งหมดตรงกฎ `.gitignore:2:/practics/` |
| ตรวจเส้นทางฝึกที่อาจเพิ่มในอนาคต `practics/future/session/main.go` | ตรงกฎเดียวกัน |
| ตรวจ README, ROADMAP, .gitignore, สารบัญ Basic, คู่มือฝึก, ตัวอย่างและ go.mod ต้นฉบับ | ไม่ถูก ignore |
| `git ls-files --others --exclude-standard` | ไม่พบ practics ในไฟล์ที่ Git มองเห็นสำหรับเพิ่มตามปกติ |
| `gofmt -l .` ใน module ฝึก | ไม่มีรายชื่อไฟล์ที่ต้องจัดรูปแบบ, exit 0 |
| `go run .` ใน practics/phase-01/ep01-hello-go | `Hello, Go!` หนึ่งบรรทัด, exit 0 บน Go 1.27.1 windows/amd64 |

ใช้ Git metadata ใน repository ชั่วคราวใต้ TEMP โดยชี้ work-tree มาที่ Go และปิดผลจาก global excludes ในการตรวจ ไม่ init repository จริง ไม่มีการ add/commit/remote/push งานผู้เรียน

ครั้งแรกที่รัน Go ภายใต้ sandbox เขียน build cache ของบัญชีไม่ได้ จึงรันด้วยสิทธิ์เครื่องมือตามที่ได้รับอนุมัติและผ่าน ตั้ง GOPROXY=off/GOTOOLCHAIN=local เฉพาะ process ไม่ดาวน์โหลด dependency

## ข้อจำกัด

ยังไม่มี Git repository/remote ในโฟลเดอร์หลักสูตรจริง จึงยังไม่มี commit หรือการ push ให้ตรวจ การ ignore มีผลต่อไฟล์ที่ยังไม่ถูกติดตาม ไม่ห้ามการบังคับ add -f หรือการรวมไฟล์ลง ZIP ต้องตรวจรายการก่อนเผยแพร่ในอนาคต

คำสั่งสร้าง practics สำหรับเครื่องที่ clone ใหม่ใน docs/PRACTICE.md ยังไม่ได้รันทับบนเครื่องนี้ เพราะมีโฟลเดอร์ฝึกอยู่แล้ว ไม่ทดสอบ API/Database/Gateway เพราะมีเพียง roadmap การเปลี่ยนเอกสารไม่ถือว่าผู้เรียนผ่าน EP.1 หรือ Phase 1
