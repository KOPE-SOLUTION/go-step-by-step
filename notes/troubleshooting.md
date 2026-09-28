# ตรวจปัญหาที่พบบ่อย

| อาการ | ตรวจและแก้อย่างไร |
|---|---|
| ไม่รู้จักคำสั่ง go | รัน `go version` ใน terminal ใหม่ หากยังไม่พบ ดู [การติดตั้ง Windows](setup-windows.md) และ [ตัวติดตั้งทางการ](https://go.dev/doc/install) |
| หา main.go ไม่พบ | ตรวจว่า terminal อยู่ที่ `practics` และบันทึกไฟล์ชื่อ `main.go` แล้ว |
| ผลยังเหมือนเดิม | บันทึกไฟล์ แล้วตรวจว่ารันจากโฟลเดอร์ที่กำลังแก้ |
| undefined: fmt.println หรือ fmt.PrintIn | ชื่อที่ถูกคือ `Println` ขึ้นต้นด้วย P ใหญ่ และใช้ l เล็กก่อน n ตรวจชื่อและเลขบรรทัดใน error |
| declared and not used | มีตัวแปรภายในฟังก์ชันที่ประกาศแล้วไม่ใช้ ให้นำตัวแปรไปใช้ หรือลบคำประกาศถ้าไม่จำเป็น |
| no new variables on left side of := | ถ้าต้องการเปลี่ยนค่าตัวแปรเดิมให้ใช้ `=` |
| index out of range | index เริ่มจาก 0 และต้องน้อยกว่าจำนวนสมาชิก เช่น `len(readings)` ตรวจว่า slice ไม่ว่างก่อนอ่านสมาชิกแรกหรือสมาชิกสุดท้าย |
| assignment to entry in nil map | สร้าง map ด้วย `map[string]float64{}` หรือ `make(map[string]float64)` ก่อนเพิ่มหรือแก้สมาชิก |
| หา package ของงานฝึกไม่พบ | EP.12: import ต้องเริ่มด้วยชื่อ module ใน `go.mod` แล้วตามด้วยโฟลเดอร์ package เช่น `/sensor` |
| go.mod file not found | EP.1–11 ใช้ `go run main.go` ส่วน EP.12 เป็นต้นไปสร้าง module ครั้งเดียวตามบท |
| main redeclared | มีไฟล์ `.go` เก่าที่ประกาศ `main` ซ้ำ ให้เก็บสำเนาเก่าเป็น `.txt` หรือย้ายไปโฟลเดอร์อื่น |
| [no test files] | EP.13–14: ตั้งชื่อไฟล์ `main_test.go` แล้วรัน `go test -v .` จากโฟลเดอร์ที่เก็บ `main_test.go` |
| test ไม่ผ่านหลังขึ้น EP.14 | เปลี่ยนทั้ง `main.go` และ `main_test.go` เป็นตัวอย่างของ EP.14 เพราะ test ที่แก้กฎไว้ใน EP.13 อาจคาดหวังผลต่างกัน |

ถ้ารันโค้ดอ้างอิง ให้เปิด terminal ที่โฟลเดอร์ EP ใต้ `lessons/phase-01` แล้วใช้ `go run .` ตัวอย่างแต่ละ EP ใช้ `lessons/phase-01/go.mod` ร่วมกัน หากดาวน์โหลดเพียง `main.go` ให้ทำตาม [วิธีเริ่มฝึก](../docs/PRACTICE.md)

[วิธีใช้พื้นที่ฝึก](../docs/PRACTICE.md) · [สารบัญ](../docs/playlist-01-go-basic/README.md)
