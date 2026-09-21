# แหล่งอ้างอิงและการจัดลำดับ

ใช้เอกสาร Go ทางการตรวจหลักภาษาและคำสั่ง ส่วนการแบ่งเป็น 14 EP และตัวอย่างอุณหภูมิเป็นการออกแบบของหลักสูตรนี้ เอกสารทางการไม่ได้กำหนดว่าต้องมีจำนวนตอนเท่านี้

| ส่วนที่ตรวจเทียบ | เอกสารอ้างอิง | ใช้ในบท |
|---|---|---|
| โปรแกรมแรก main/import/fmt และ go run | [Getting started](https://go.dev/doc/tutorial/getting-started) | EP.1 |
| ตัวแปร ค่าคงที่ ตัวดำเนินการ เงื่อนไข และลูป | [Language specification](https://go.dev/ref/spec) | EP.2–4 |
| ฟังก์ชันและการคืนค่า | [Function declarations](https://go.dev/ref/spec#Function_declarations) | EP.5 |
| คืน error และตรวจ nil | [Return and handle an error](https://go.dev/doc/tutorial/handle-errors) | EP.6 |
| array, slice, map และ struct | [Types](https://go.dev/ref/spec#Types) | EP.7–9 |
| pointer, method และ interface | [Method declarations](https://go.dev/ref/spec#Method_declarations), [Interfaces](https://go.dev/ref/spec#Interface_types) | EP.10–11 |
| package/module และ import path | [Create a Go module](https://go.dev/doc/tutorial/create-module) | EP.12 |
| testing และ go test | [Add a test](https://go.dev/doc/tutorial/add-a-test) | EP.13–14 |

[หน้ารวม Tutorials](https://go.dev/doc/tutorial/) ใช้ดูเส้นทางต่อยอด เช่น ฐานข้อมูลและ API ส่วน [Effective Go](https://go.dev/doc/effective_go) ใช้ประกอบแนวทางเขียนโค้ด เอกสารระบุว่าไม่ได้ปรับปรุงตามระบบ Go ใหม่ทั้งหมด จึงไม่ใช้เป็นแหล่งเดียว

## เหตุผลการจัดบท

- รวมแนวคิดที่ต้องใช้ร่วมกันเพื่อทำงานหนึ่งอย่าง เช่น ตัวแปร ชนิดข้อมูล และรูปแบบรายงาน
- ใช้การทดลองหลายขั้นภายในบท แทนการแยกหนึ่งคำสั่งเป็นหนึ่ง EP
- ฟังก์ชันมาก่อน error เพื่อเข้าใจการรับและคืนค่าก่อนเพิ่มผลลัพธ์ทางที่ล้มเหลว
- struct มาก่อน pointer/method และ interface เพื่อให้เห็นข้อมูลที่กำลังอ่านหรือเปลี่ยน
- package/module เต็มบทอยู่หลังเขียนโปรแกรมเล็กได้แล้ว ต่างจาก tutorial ทางการที่เริ่ม module ตั้งแต่แรก งานฝึกช่วงต้นใช้ `go run main.go` จึงไม่ต้องตั้งค่า module เองทันที
- ตัวอย่างอ้างอิงทั้งหมดอยู่ใน module เดียวเพื่อให้รันและตรวจทั้งเฟสได้ ส่วน practics ใช้ module แยกเมื่อถึง EP.12
- concurrency, defer กับการปิดไฟล์, generics และรายละเอียด string/Unicode จะเพิ่มเมื่อมีโจทย์ที่ต้องใช้ ไม่ใช่เงื่อนไขเริ่มโปรเจกต์พื้นฐานนี้

**หลักภาษา** เช่นชนิดและการคัดลอกค่าต้องทำงานตาม Go ส่วน **แนวทางออกแบบ** เช่นแยกคำนวณจากการพิมพ์ และ **ทางเลือก** เช่นใช้ interface หรือ switch เลือกตามปัญหาที่กำลังแก้

ตรวจเอกสารวันที่ 2026-09-21 โค้ดประกาศ Go 1.22.0 และตรวจด้วย Go 1.27.1 บน Windows ดูขอบเขตที่ยืนยันได้ใน [ผลตรวจ](../notes/phase-01-results.md)
