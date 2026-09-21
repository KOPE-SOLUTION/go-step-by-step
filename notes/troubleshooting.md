# ตรวจปัญหาที่พบบ่อย

| อาการ | ตรวจและแก้อย่างไร |
|---|---|
| ไม่รู้จักคำสั่ง go | รัน `go version` ใน terminal ใหม่ หากยังไม่พบ ดู [การติดตั้ง Windows](setup-windows.md) และ [ตัวติดตั้งทางการ](https://go.dev/doc/install) |
| หา main.go ไม่พบ | ตรวจว่า terminal อยู่ที่ practics และบันทึกไฟล์ชื่อ main.go จริง |
| ผลยังเหมือนเดิม | บันทึกไฟล์ แล้วตรวจว่ารันจากโฟลเดอร์ที่กำลังแก้ |
| undefined: fmt.PrintIn | ชื่อที่ถูกคือ Println ใช้ตัว l เล็ก อ่านชื่อและเลขบรรทัดใน error |
| declared and not used | ใช้ตัวแปรที่ประกาศไว้ หรือนำประกาศที่ไม่จำเป็นออก |
| no new variables on left side of := | ถ้าต้องการเปลี่ยนตัวแปรเดิมให้ใช้ = |
| index out of range | index เริ่ม 0 และต้องน้อยกว่า len ตรวจ slice ว่างก่อนอ่านสมาชิก |
| assignment to entry in nil map | สร้าง map ด้วย literal หรือ make ก่อนเขียนสมาชิก |
| หา package ของงานฝึกไม่พบ | EP.12: ตรวจ go.mod และ import ต้องใช้ชื่อ module เดียวกัน |
| go.mod file not found | EP.1–11 ใช้ go run main.go; EP.12 เป็นต้นไปสร้าง module ครั้งเดียวตามบท |
| main redeclared | มีไฟล์ .go เก่าที่ประกาศ main ซ้ำ เก็บสำเนาเก่าเป็น .txt หรือโฟลเดอร์อื่น |
| [no test files] | EP.13–14: ใช้ชื่อ main_test.go และรัน go test -v . จากโฟลเดอร์เดียวกับไฟล์นั้น |
| test ไม่ผ่านหลังขึ้น EP.14 | เปลี่ยน main_test.go เป็นของโปรเจกต์พร้อม main.go อย่าใช้ test กฎทดลองจาก EP.13 ค้างไว้ |

ถ้ารันโค้ดอ้างอิง ให้เปิด terminal ที่โฟลเดอร์ EP ใต้ lessons/phase-01 แล้วใช้ `go run .` go.mod ร่วมอยู่ที่ phase-01 หากดาวน์โหลดเพียง main.go ให้ใช้วิธีฝึกใน practics แทน

[วิธีใช้พื้นที่ฝึก](../docs/PRACTICE.md) · [สารบัญ](../docs/playlist-01-go-basic/README.md)
