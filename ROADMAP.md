# แผนหลักสูตร Go — 5 Phase

Phase 1 มีเนื้อหาพร้อมฝึก 14 EP ส่วน Phase 2–5 เป็นแผน ยังไม่ได้สร้างบริการหรือฐานข้อมูลในเฟสเหล่านั้น

| Phase | เนื้อหา | ผลลัพธ์ |
|---|---|---|
| 1 — Basic | ภาษา Go และการทดสอบพื้นฐาน | รายงานข้อมูลการวัดใน terminal |
| 2 — Practical Go | JSON, ไฟล์, HTTP client, งานพร้อมกัน | โปรแกรมอ่านข้อมูลและจัดการความผิดพลาด |
| 3 — Back-end & API | HTTP server, JSON, validation, CRUD และ test | API สำหรับข้อมูลจำลองบน localhost |
| 4 — Database | SQL, database/sql, migration, transaction | ข้อมูลอยู่หลัง restart และป้องกัน event ซ้ำ |
| 5 — IoT / Edge Gateway | อุปกรณ์จำลอง, MQTT, offline queue และ ACK | ทดลองส่งย้อนหลังและรับมือความขัดข้อง |

## Phase 1 — พื้นฐานที่นำมาประกอบเป็นโปรแกรมได้

| EP | เนื้อหา | นำไปต่อยอด |
|---|---|---|
| 1 | [เขียนและรันโปรแกรม Go แรก](lessons/phase-01/ep01-hello-go/README.md) | อ่านโค้ดและตรวจปัญหาจากเครื่องมือ |
| 2 | [เก็บข้อมูลด้วยตัวแปรและชนิดข้อมูล](lessons/phase-01/ep02-variables-types/README.md) | รูปข้อมูลและการแสดงค่าการวัด |
| 3 | [คำนวณและตรวจเงื่อนไขของข้อมูล](lessons/phase-01/ep03-calculations-conditions/README.md) | กฎตรวจ input และสถานะ |
| 4 | [ทำซ้ำและสรุปค่าการวัดด้วย for](lessons/phase-01/ep04-loops/README.md) | ประมวลผลหลายรายการ |
| 5 | [แยกงานด้วยฟังก์ชัน รับค่าและคืนค่า](lessons/phase-01/ep05-functions/README.md) | กฎที่ใช้ซ้ำได้ |
| 6 | [รับมือข้อมูลผิดด้วย error](lessons/phase-01/ep06-errors/README.md) | ข้อมูลผิดและงานที่ล้มเหลว |
| 7 | [จัดการรายการข้อมูลด้วย array และ slice](lessons/phase-01/ep07-arrays-slices/README.md) | รายการอุปกรณ์และค่าการวัด |
| 8 | [ค้นหาและอัปเดตข้อมูลด้วย map](lessons/phase-01/ep08-maps/README.md) | ค้นค่าล่าสุดตามชื่อ |
| 9 | [สร้างชนิดข้อมูลการวัดด้วย struct](lessons/phase-01/ep09-structs/README.md) | ข้อมูลหนึ่งรายการที่มีหลายช่อง |
| 10 | [เพิ่ม method และแก้ต้นฉบับด้วย pointer](lessons/phase-01/ep10-methods-pointers/README.md) | การอ่านและแก้สถานะข้อมูล |
| 11 | [สลับแหล่งข้อมูลด้วย interface](lessons/phase-01/ep11-interfaces/README.md) | สลับตัวจำลองกับแหล่งข้อมูลอื่น |
| 12 | [จัดโค้ดเป็น package และ module](lessons/phase-01/ep12-packages-modules/README.md) | โค้ดที่โตขึ้นโดยแยกหน้าที่ |
| 13 | [ทดสอบกฎและค่าขอบเขตด้วย unit test](lessons/phase-01/ep13-unit-tests/README.md) | ตรวจค่าขอบเขตและการเปลี่ยนกฎ |
| 14 | [โปรเจกต์สรุป: รายงานข้อมูลการวัด](lessons/phase-01/ep14-measurement-report/README.md) | ส่วนประมวลผลสำหรับ JSON และ API |

ลำดับใหม่นี้รวมการรัน อ่านโค้ด แสดงข้อความ และอ่านข้อผิดพลาดไว้ในบทแรก รวมตัวแปรกับชนิดข้อมูลและค่าศูนย์ไว้ในบทเดียว ส่วนฟังก์ชันกับ error แยกกันเพื่อฝึกรับ/คืนค่าให้เข้าใจก่อนจัดการทางที่ล้มเหลว

ย้าย interface มาอยู่หลัง method และ pointer เพื่อให้พื้นฐานครบก่อนต่อยอดบริการ สอน package/module เต็มบทเมื่อมีโค้ดให้แยกจริง แล้วจึงเขียน test และโปรเจกต์สรุป รายละเอียดการเทียบแหล่งอ้างอิงอยู่ใน [REFERENCES](docs/REFERENCES.md)

## Phase 2 — โปรแกรมใช้งาน

| งานหลัก | สิ่งที่เรียนร่วมกัน | นำไปต่อยอด |
|---|---|---|
| โหลด configuration | JSON, ค่าที่ขาด, validation | เปลี่ยนรายการอุปกรณ์โดยไม่แก้โค้ด |
| บันทึกและอ่านไฟล์ | file I/O, defer, error, log | ประวัติและคิวบนดิสก์ |
| เรียก HTTP | request/response, timeout, error | อ่านข้อมูลจากบริการ |
| อ่านหลายแหล่งข้อมูล | interface จาก Phase 1, goroutine, channel, context | อ่านพร้อมกันและยกเลิกงาน |
| หยุดโปรแกรมอย่างมีลำดับ | graceful shutdown และทดสอบงานพร้อมกัน | จบงานและปิดทรัพยากร |

เรียน defer เมื่อมีไฟล์หรือทรัพยากรที่ต้องปิดจริง ทบทวนค่าพิเศษ NaN/Inf และการเก็บสาเหตุของ error ก่อนรับข้อมูลจากภายนอก ดูรายละเอียดใน [แผน Phase 2–5](docs/FUTURE_ROADMAP.md)

## Phase 3 — Back-end และ API

สร้างบริการด้วย net/http → รับส่ง JSON → นำกฎจากโปรแกรมเดิมมาใช้ → ตรวจ request → CRUD ในหน่วยความจำ → ทดสอบ handler และ shutdown เริ่มบน localhost และไม่บังคับ framework

## Phase 4 — Database

เรียนตารางและ SQL → เชื่อม Go ผ่าน database/sql กับ driver → เพิ่ม/ค้นข้อมูล → ต่อ API → ทดสอบ restart → migration และ transaction → event ID และ unique constraint → ค้นประวัติและสำรองข้อมูล

เริ่มวางแผนด้วย SQLite ส่วน PostgreSQL เป็นทางเลือกภายหลัง ต้องตรวจ driver และ Go ที่ใช้งานจริงก่อนเริ่มบทฐานข้อมูล

## Phase 5 — Gateway จำลอง

| งานหลัก | การทดลองที่ต้องทำ | นำไปต่อยอด |
|---|---|---|
| สร้างข้อมูลการวัด | event ID และเวลาอ่านเดิม | ระบุเหตุการณ์ตลอดการส่งและ retry |
| อ่านหลายอุปกรณ์ | ตัวหนึ่งเสียแล้วตัวอื่นยังทำงาน | แยกผลกระทบของความล้มเหลว |
| ส่ง MQTT ในเครื่อง | broker, topic, publish/subscribe, QoS | ส่งข้อมูลถึงผู้รับจำลอง |
| บันทึกก่อนส่ง | เขียนไม่ได้ต้องไม่รายงานว่ารับสำเร็จ | คิวบนดิสก์ |
| ตัดการเชื่อมต่อและส่งย้อนหลัง | retry แบบหน่วงเวลา ใช้ ID และเวลาเดิม | store-and-forward |
| รับ ACK หลังบันทึก | แยก MQTT PUBACK ออกจาก ACK หลังฐานข้อมูล commit | เงื่อนไขนำข้อมูลออกจากคิว |
| ทดสอบความขัดข้อง | restart, ACK หาย, ข้อมูลซ้ำ, ไฟล์ท้ายขาด, คิวเต็ม | เห็นข้อจำกัดและนโยบายเมื่อรับเพิ่มไม่ได้ |

API และ Database อยู่ก่อน Gateway เพื่อให้เข้าใจฝั่งผู้รับและความหมายของการบันทึกสำเร็จ โดยไม่ต้องเรียน SQL, HTTP และ MQTT พร้อมกัน

ใช้ข้อมูลจำลองและ localhost ไม่ใช้ข้อมูลบริษัทหรือ production ไม่ต้องใช้ Docker ใน Phase 1 ตัวอย่างคิวและ ACK เป็นแบบฝึกหัด ไม่อ้าง exactly-once ทั้งระบบหรือ production-ready
