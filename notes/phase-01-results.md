# ผลตรวจ Phase 1 — Go Basic

วันที่ 2026-09-19 • Windows x64 • Go 1.27.1 windows/amd64 • PowerShell 7.6.5

## ขอบเขตที่ตรวจจริง

- ครบ 28 module: EP1 เดิม และ EP2–28 ที่เพิ่ม
- ตัวอย่าง 28 โปรแกรม และเฉลย 56 โปรแกรม รวม **84 ผลลัพธ์** ตรงตามที่กำหนดใน phase01-checks.json
- gofmt -l ไม่มีไฟล์ที่ต้องจัดรูปแบบ
- go test -count=1 ./... ผ่านในทุก module; **มี unit test จริงใน EP23, EP24 และ EP28** ส่วน EP อื่นได้ [no test files] และใช้คำสั่งนี้ตรวจ compile เท่านั้น
- EP4: สำเนาที่สะกด fmt.Println เป็น fmt.println ล้มตามคาดและมีข้อความ undefined: fmt.println
- EP23 และ EP24: ทดลองเปลี่ยน >= 30 เป็น > 30 ในสำเนา test จับข้อผิดพลาดที่ค่า 30 ได้จริง (exit 1)
- EP28: ทดลองเปลี่ยน continue เป็น return lines ในสำเนา TestBuildReport ตรวจพบว่ารายงานหยุดก่อนรายการถัดไป (exit 1)
- new-practice.ps1: สร้างสำเนา EP2, EP21, EP23, EP28 ใน TEMP แล้ว compile/test ผ่าน ไม่คัดลอก solutions; เรียกซ้ำ EP2 ถูกปฏิเสธก่อนเขียน และ hash ของไฟล์เดิมไม่เปลี่ยน

ผลจากคำสั่งตรวจรวม:

```text
PASS: 28 episodes; 84 program outputs matched; gofmt clean.
```

ไฟล์ [phase01-checks.json](../scripts/phase01-checks.json) ระบุ module คำสั่งรัน และข้อความคาดหวัง ส่วน [สารบัญ](../docs/playlist-01-go-basic/README.md) เชื่อมไปผลตรวจราย EP

ตรวจเป้าหมายลิงก์ภายใน Markdown ว่ามีไฟล์จริง และตรวจโค้ด Go ที่แสดงใน README ทั้ง 38 ชุดว่าตรงกับไฟล์ที่รันผ่าน หลังปรับเฉลย EP23 ให้เก็บ test เดิมแล้วเพิ่มกรณีด้วยชื่อไม่ซ้ำ ได้ตรวจ EP23 ซ้ำผ่านทั้ง compile, test, gofmt และผลทั้งสามโปรแกรม

## วิธีตรวจซ้ำ

จากราก Go ใน PowerShell 7:

```powershell
./scripts/verify-phase01.ps1
```

ตรวจเฉพาะ EP เช่น EP24:

```powershell
./scripts/verify-phase01.ps1 -Episode 24
```

บนเครื่องนี้หากหน้าต่างเดิมยังไม่เห็น PATH สามารถระบุเครื่องมือที่ติดตั้งแล้ว:

```powershell
./scripts/verify-phase01.ps1 -GoCommand 'C:/Users/kopes/AppData/Local/Programs/Go-1.27.1/go/bin/go.exe'
```

สคริปต์ตั้ง GOTOOLCHAIN=local และ GOPROXY=off ชั่วคราว แล้วคืนค่าเดิมเมื่อจบ ไม่ติดตั้งเครื่องมือหรือดาวน์โหลด dependency ถ้าไม่รู้จัก Go ให้ตรวจตาม [troubleshooting](troubleshooting.md)

## ข้อจำกัดของผลตรวจ

ครั้งแรกที่ตรวจใน sandbox EP2 เขียน build cache ของบัญชีไม่ได้ (Access is denied) การรันด้วยสิทธิ์เครื่องมือที่ได้รับอนุญาตจึงตรวจผ่านครบ ไม่ปิดระบบป้องกันและไม่เปลี่ยนการตั้งค่าระบบเพิ่มในงานนี้

สคริปต์สร้างพื้นที่ฝึกทดสอบปลายทาง TEMP; ยังไม่ยืนยันว่าหน้าต่าง PowerShell ของผู้เรียนเขียนใน Documents ได้ทุกกรณี เพราะเครื่องเคยปฏิเสธการสร้างไฟล์จากบาง process งานฝึก EP1 เดิมไม่ได้ถูกเขียนทับ และไม่ได้สร้างงานฝึกใหม่ครบทุกตอนล่วงหน้า

go.mod ประกาศ Go 1.22.0 เป็นขั้นต่ำ แต่ทดสอบจริงเฉพาะ Go 1.27.1 บน Windows เครื่องนี้ ยังไม่ทดสอบ toolchain 1.22, Linux/macOS, อุปกรณ์จริง, MQTT, API หรือฐานข้อมูล

โปรเจกต์ท้ายเฟสใช้ตัวเลขปกติและชื่อที่กำหนดในโค้ด ไม่ครอบคลุม NaN ชื่อที่มีแต่ช่องว่าง หรือข้อมูลจากภายนอก ยังไม่มีระบบเก็บถาวร งานพร้อมกัน restart หรือ ACK

การตรวจโปรแกรมกับการเรียนผ่านเป็นคนละเรื่อง เช็กลิสต์ของผู้เรียนใน [PROGRESS](../PROGRESS.md) ยังรอการทดลองและคำอธิบายของผู้เรียนเอง
