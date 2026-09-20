# go-step-by-step — เรียน Go ทีละก้าว

หลักสูตรภาษาไทยตั้งแต่ Basic ไปสู่ Back-end, Database และ Edge Gateway จำลอง

เรียน Go ผ่านการคาดเดา ลงมือรัน อ่านผล และอธิบายด้วยคำของตนเอง เริ่มจากพื้นฐานและค่อย ๆ ต่อเป็นโปรแกรมที่ใช้งานได้ ไม่จำกัดเวลาที่ใช้ทำความเข้าใจแต่ละตอน

ตอนนี้เตรียม **Phase 1 — Go Basic ครบ 28 EP** แล้ว เรียนทีละ EP ตามความพร้อมของตนเองได้ Phase 2–5 ยังเป็นแผน การมีไฟล์ครบไม่ใช่ข้อกำหนดให้รีบเรียน

## เส้นทางการเรียนรู้ตาม Phase

ใช้ Case Study **ข้อมูลการวัดจากอุปกรณ์จำลอง** ต่อเนื่องจาก Console ไปสู่บริการและ Gateway เริ่มได้โดยไม่ต้องมีพื้นฐาน Java, OOP หรือ Desktop App

| Phase | Playlist / Track | ผลลัพธ์เมื่อเรียนครบ | สถานะตอนนี้ |
|---|---|---|---|
| 1 | [Go Basic](docs/playlist-01-go-basic/README.md) | Console mini project และพื้นฐานภาษา | มีบทเรียน EP1–28 พร้อมตัวอย่าง แบบฝึกหัด และเฉลย |
| 2 | Practical Go / Application Core | อ่าน configuration/ไฟล์ จัดการงานและกฎการทำงาน | แผน |
| 3 | Go Back-end & REST API | บริการ HTTP/JSON บน localhost เริ่มจากข้อมูลในหน่วยความจำ | แผน |
| 4 | Database / Persistence | เก็บและค้นประวัติผ่าน Go/API ข้อมูลอยู่หลัง restart | แผน |
| 5 | IoT / Edge Gateway | อุปกรณ์จำลอง MQTT คิว offline และ ACK หลังบันทึกข้อมูล | แผน |

**เริ่มที่ Phase 1 — Basic ได้เลย** ไม่ต้องเรียนเฟสอื่นก่อน หนึ่งเฟสประกอบด้วยหลายบทและหลาย EP ดู [สารบัญ Playlist](docs/README.md), [ROADMAP](ROADMAP.md) และ [ขอบเขตเฟสอนาคต](docs/FUTURE_ROADMAP.md)

Back-end คือโปรแกรมฝั่งบริการ, API คือช่องทางให้โปรแกรมอื่นเรียกมัน, Database คือระบบเก็บและค้นข้อมูล สามเรื่องนี้มีเฟสแยกแล้ว แต่ยังไม่ติดตั้งหรือเขียนระบบเหล่านั้นในช่วง Basic

## เริ่มที่นี่

1. เครื่องนี้ติดตั้ง Go 1.27.1 แล้วเมื่อ 2026-09-19 ดู [ผลตรวจเครื่อง](notes/environment.md) ไม่ต้องติดตั้งซ้ำ
2. เปิด [EP.1 ฉบับทีละก้าว](lessons/01-hello-go/EP01.md) ทำทีละช่วง ยังไม่ต้องอ่านเนื้อหารวมทั้งบท
3. ทำแบบฝึกหัดสั้น 2 ข้อใน EP.1 ก่อนเปิดเฉลย เมื่อพร้อมให้เลือกตอนต่อไปจาก [สารบัญ Phase 1](docs/playlist-01-go-basic/README.md)

ตำแหน่งหลักสูตรบนเครื่องนี้:

```text
C:\Users\kopes\Documents\L\Go
```

เปิดโฟลเดอร์นี้ใน editor และใช้เป็นโฟลเดอร์หลักเมื่อทำตามคำสั่งในบทเรียน

## โครงสร้าง

```text
Go/
  README.md
  ROADMAP.md
  .gitignore
  docs/
    README.md
    PRACTICE.md
    FUTURE_ROADMAP.md
    playlist-01-go-basic/README.md
  lessons/
    01-hello-go/
      README.md
      EP01.md
      go.mod
      examples/hello/main.go
      exercises/README.md
      exercises/my-hello/main.go
      solutions/README.md
      solutions/01-greeting/main.go
      solutions/02-two-lines/main.go
      solutions/03-debug.md
      tests/README.md
      tests/verify.ps1
      tests/RESULTS.md
    phase-01/
      ep02-read-program/     # ถึง ep28-project-tests
        README.md
        go.mod
        examples/
        exercises/README.md
        solutions/
        tests/RESULTS.md
  scripts/
    verify-phase01.ps1
    phase01-checks.json
  practics/                  # งานฝึกส่วนตัว: Git ignore ทั้งหมด
    main.go                 # ผู้เรียนสร้างครั้งเดียวและใช้ต่อทุก EP
  projects/README.md
  notes/
    environment.md
    setup-windows.md
    glossary.md
    troubleshooting.md
```

## โค้ดตัวอย่างและพื้นที่ฝึก

**Module** คือกลุ่ม package ที่มีไฟล์ `go.mod` ระบุชื่อและข้อกำหนด ส่วน **package** คือกลุ่มไฟล์ Go ที่ทำงานร่วมกันในโฟลเดอร์เดียวกัน นี่เป็นภาพรวมพอให้เริ่มรัน รายละเอียดเรียนใน EP21–22

**หลักของเครื่องมือ Go:** ใช้ `go.mod` เป็นจุดเริ่มต้นของ module และระบุเวอร์ชันภาษาขั้นต่ำ ตัวอย่างแต่ละโปรแกรมมีโฟลเดอร์ของตนเอง จึงไม่ชนกันเพราะมี `main` หลายชุด

**แนวทางออกแบบหลักสูตร:** EP1 ใช้ module Hello Go เดิม ส่วน EP2–28 มี module แยกแต่ละ EP โค้ดที่รันไม่ผ่านระหว่างฝึกใน EP หนึ่งจะไม่ขัดขวางการทดสอบบทอื่น และบท MQTT ในอนาคตจะไม่เพิ่มไลบรารีให้บทแรก รันคำสั่งจากโฟลเดอร์บทที่กำลังเรียน ไม่รัน `go test ./...` จากราก `Go` ซึ่งไม่มี module

ตัวอย่างบทแรกใน `lessons/` ใช้ชื่อ `example.com/go-course/hello` ไม่มีการเผยแพร่หรือเชื่อมต่อโดเมนนี้ ใช้เฉพาะ `fmt` ที่ติดมากับ Go ไม่มีไลบรารีภายนอก ไม่ต้อง `go get` หรือ `go mod init` ซ้ำ เพราะตัวอย่างใน `lessons/` มี `go.mod` แล้ว ส่วนพื้นที่ฝึก `practics/main.go` ช่วง EP.1–20 รันด้วย `go run main.go` ได้โดยยังไม่ต้องมี `go.mod`

กำหนด `go 1.22.0` เป็นฐานขั้นต่ำที่ประกาศใน module ไม่ใช่เวอร์ชันที่ต้องดาวน์โหลด ตรวจ compile และรันจริงผ่านบน Go 1.27.1 windows/amd64 เมื่อ 2026-09-19 ยังไม่ได้ทดสอบด้วย toolchain 1.22 เองหรือทุกรุ่นระหว่างนั้น ไม่ต้องติดตั้งรุ่นเก่าให้ตรงตัวเลขในไฟล์

**พื้นที่ฝึก:** ใช้ `practics/main.go` ไฟล์เดียวและเขียนทับเมื่อขึ้น EP ใหม่ เริ่มมี module ฝึกหนึ่งชุดเมื่อถึง EP.21 และใช้ชุดเดิมต่อไป ยังไม่ใช้ `go.work`, framework หรือ Docker

ผลตรวจทุก EP และวิธีรันตรวจซ้ำ: [Phase 1 Results](notes/phase-01-results.md)

อ่านอ้างอิง: [Go modules](https://go.dev/doc/modules/managing-dependencies), [go directive](https://go.dev/doc/modules/gomod-ref#go)

## งานต้นฉบับและงานฝึก

`lessons/` เก็บตัวอย่างแยก EP สำหรับย้อนดู ส่วนงานฝึกใช้ `practics/main.go` ไฟล์เดิมตลอด ขึ้น EP ใหม่ก็เปลี่ยนโค้ดในไฟล์นี้ โจทย์อยู่ในหัวข้อ **3. ฝึกเอง** ของบท ดู [วิธีเริ่มฝึก](docs/PRACTICE.md)

ใช้ชื่อ `practics` ตามคำขอ ฝั่ง Java อ้างอิงใช้ชื่อ `practice` ตำแหน่งบทเดิมยังเหมือนเดิม จึงเปิดลิงก์ EP.1 และรันคำสั่งตัวอย่างเดิมต่อได้

เริ่มใช้ Git repository ภายในโฟลเดอร์ Go เมื่อ 2026-09-20 บน branch `main` เพื่อเก็บประวัติหลักสูตร ใช้ `origin` เป็น [KOPE-SOLUTION/go-step-by-step](https://github.com/KOPE-SOLUTION/go-step-by-step) `.gitignore` แยกงานฝึกออกจากไฟล์ที่เพิ่มตามปกติ การ ignore ไม่ลบไฟล์ที่เคยติดตามแล้วและไม่ป้องกันการฝืนเพิ่มด้วย `git add -f` จึงไม่ใช่ที่เก็บข้อมูลลับ

## วิธีเรียนและขอบเขตการทดลอง

- ทุกบทแยกเป้าหมาย ตัวอย่าง คำสั่ง ผลที่คาดหวัง แบบฝึกหัด เฉลย และหลักฐานการตรวจ
- ถ้าพบ error ให้เก็บคำสั่ง โฟลเดอร์ปัจจุบัน และข้อความ error ก่อนแก้ทีละจุด
- ใช้ชื่อไฟล์และตัวแปรภาษาอังกฤษ คำอธิบายภาษาไทย เรียนศัพท์ก่อนใช้
- ใช้ข้อมูลจำลองและ localhost ไม่ใช้ IP รหัสผ่าน หรือข้อมูลบริษัทจริง
- ไม่เชื่อมต่ออุปกรณ์หรือ MQTT production ไม่เปิดบริการออกอินเทอร์เน็ต ไม่ push Git ไม่ลบงานเดิม
- คิวและ ACK ในช่วงท้ายเป็นแบบฝึกหัด มีข้อจำกัด ไม่อ้างว่า production-ready

หลักสูตรต่อยอดงานเก็บข้อมูลและสื่อสารบนคอมพิวเตอร์ที่ทำหน้าที่ gateway บทแรกพิมพ์ข้อความเท่านั้น ยังไม่อ่านอุปกรณ์จริง

ดูแผนทั้งหมดใน [ROADMAP.md](ROADMAP.md) และคำศัพท์ใน [glossary.md](notes/glossary.md)
