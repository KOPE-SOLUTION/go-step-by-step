# go-step-by-step — เรียน Go ทีละขั้น

เรียน Go ภาษาไทยสำหรับผู้เริ่มต้น พร้อมตัวอย่างที่รันได้และโจทย์ในหน้าบทเรียน ใช้ข้อมูลอุณหภูมิจำลองค่อย ๆ พัฒนาเป็นโปรแกรมที่ใช้งานได้

## เริ่มเรียน: Playlist 01 — Go Basic

**[เปิด Playlist 01 — สารบัญบทเรียนพื้นฐาน Go ทั้ง 14 EP](docs/playlist-01-go-basic/README.md)**

Playlist 01 คือบทเรียน **Phase 1** เริ่มจากโปรแกรมแรกไปจนถึงรายงานข้อมูลพร้อมการทดสอบ แต่ละ EP มีตัวอย่างและโจทย์ในหน้าเดียวกัน ส่วน Phase 2–5 ยังเป็นแผน

- **เริ่มตอนแรก:** [EP.1 — เขียน รัน และแก้ข้อผิดพลาดโปรแกรมแรก](lessons/phase-01/ep01-hello-go/README.md)
- **เตรียมพื้นที่ฝึก:** [ใช้ practics/main.go ไฟล์เดิม](docs/PRACTICE.md)

ก่อนเริ่ม ลอง `go version` ใน terminal หากยังรันไม่ได้ ดู [วิธีตรวจปัญหา](notes/troubleshooting.md)

## เส้นทางการเรียนรู้

| Phase | เนื้อหา | สิ่งที่จะได้ทำ |
|---|---|---|
| 1 | [Playlist 01 — Go Basic](docs/playlist-01-go-basic/README.md) | เขียนและทดสอบโปรแกรมรายงานการวัดใน terminal |
| 2 | โปรแกรมใช้งาน | อ่าน JSON และไฟล์ เรียก HTTP และจัดการงานพร้อมกัน |
| 3 | Back-end และ API | รับส่งข้อมูลผ่านบริการบน localhost |
| 4 | Database | เก็บข้อมูลและค้นคืนหลัง restart |
| 5 | IoT และ Edge Gateway | อ่านอุปกรณ์จำลอง ส่ง MQTT และส่งข้อมูลค้างย้อนหลัง |

รายละเอียดอยู่ใน [ROADMAP](ROADMAP.md) และ [แผน Phase 2–5](docs/FUTURE_ROADMAP.md)

ตั้งแต่ Phase 3 จะฝึกแยก HTTP ออกจากกฎของโปรแกรม และท้าย Phase 4 จะปรับ API ที่เชื่อมฐานข้อมูลแล้วด้วยแนวคิด Clean Architecture ดู [เส้นทางจัดโครงสร้างโค้ด](docs/ARCHITECTURE.md)

## พื้นที่ฝึกและโค้ดอ้างอิง

ใช้ `practics/main.go` ไฟล์เดิมสำหรับฝึก เปลี่ยนเนื้อหาเมื่อขึ้นบทใหม่ โดยมีโค้ดอ้างอิงเก็บแยกไว้ใน `lessons/phase-01/`

EP.1–11 ใช้ `go run main.go` เมื่อถึง EP.12 จึงสร้าง module และเพิ่ม package `sensor` ส่วน EP.13–14 เพิ่มไฟล์ทดสอบตามบท

โฟลเดอร์ `practics/` ไม่เก็บใน Git ผู้ที่ดาวน์โหลดหรือ clone หลักสูตรจึงต้องสร้างพื้นที่ฝึกเองตาม [วิธีเริ่มฝึก](docs/PRACTICE.md)

<details>
<summary>โครงสร้างและเวอร์ชัน</summary>

```text
Go/
  README.md
  ROADMAP.md
  docs/               # สารบัญและวิธีฝึก
  lessons/phase-01/
    go.mod            # module ร่วมของโค้ดอ้างอิง
    ep01-hello-go/
      README.md
      main.go
      solutions/
    ...
  practics/           # สร้างเอง ไม่เก็บใน Git
    main.go
  projects/           # สำหรับโปรเจกต์เฟสถัดไป
  notes/              # คำศัพท์และผลตรวจ
  scripts/            # เครื่องมือตรวจตัวอย่าง
```

โค้ดอ้างอิง Phase 1 ใช้ module เดียว เพราะใช้เฉพาะ package ที่มากับ Go (standard library) จึงยังไม่ต้องจัดการเวอร์ชัน package ภายนอกแยกตามบท ส่วน `practics` มี module ของตัวเองเมื่อถึง EP.12

ประกาศ Go 1.22.0 เป็นขั้นต่ำ ตรวจจริงด้วย Go 1.27.1 บน Windows ยังไม่ได้ตรวจด้วย compiler Go 1.22 หรือระบบอื่น ดู [ผลตรวจ](notes/phase-01-results.md)

</details>

ใช้ข้อมูลจำลองและ localhost เป็นหลัก ยังไม่ใช้ Docker ในพื้นฐาน บทคิวและ ACK เป็นแบบฝึกหัดเพื่อเข้าใจพฤติกรรมและข้อจำกัด ไม่ใช่ระบบ production-ready

หากติดขัด ดู [คำศัพท์](notes/glossary.md) และ [วิธีตรวจปัญหา](notes/troubleshooting.md)
