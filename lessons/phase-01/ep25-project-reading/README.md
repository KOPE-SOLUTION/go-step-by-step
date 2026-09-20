# EP.25 — โปรเจกต์สรุป ตอนที่ 1 — ข้อมูลการวัดและสถานะ

**เป้าหมาย:** ประกอบ struct ฟังก์ชัน และการแสดงผลเป็นรายงานหนึ่งรายการ

## 1. อ่านโค้ด

นำ `struct` และฟังก์ชันมารวมกันเป็นรายงานหนึ่งรายการ: ชื่ออุปกรณ์ ค่าอุณหภูมิ และสถานะ

ไฟล์ [examples/model.go](examples/model.go)

```go
package main

type Reading struct {
	DeviceID string
	Celsius  float64
}

func status(reading Reading) string {
	if reading.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}
```

ไฟล์ที่ใช้ร่วมกัน: [examples/main.go](examples/main.go)

## 2. ลองรัน

**ก่อนรัน:** ตั้ง Celsius เป็น 30 แล้วส่วนใดของบรรทัดจะเปลี่ยน?

รันจากโฟลเดอร์ `lessons/phase-01/ep25-project-reading`:

```shell
go run ./examples
```

<details>
<summary>ดูผลที่คาดหวังหลังลองรัน</summary>

```text
sensor-01: 27.5 C [OK]
```

main.go สร้าง Reading แล้วเรียก status เพื่อแสดงผล เกณฑ์เตือนจำลองคือ 30

</details>

## 3. ฝึกเอง

ทำ [แบบฝึกหัด 2 ข้อ](exercises/README.md) ใน [practics](../../../docs/PRACTICE.md) แล้วลองตอบ: สถานะควรคำนวณจากค่าจริงหรือข้อความที่จัดรูปแบบแล้ว?

<details>
<summary>อ่านเพิ่มเติมเมื่อสงสัย</summary>

เริ่มจากใบรายงานหนึ่งใบก่อน รายงานต้องมีชื่ออุปกรณ์ ค่า และสถานะ ทุกค่าในบทนี้พิมพ์ไว้ในโค้ด ไม่ได้อ่านอุปกรณ์

ข้อกำหนดคือสิ่งที่โปรแกรมต้องทำให้ได้; snapshot ในหลักสูตรคือโค้ดแต่ละขั้นที่เก็บแยก EP เพื่อกลับไปดูได้; ชนิดและฟังก์ชันใน model.go อยู่ package main เดียวกับ main.go

- Reading เก็บข้อมูล status เลือกสถานะ main จัดรูปแบบและพิมพ์
- ข้อกำหนด: แสดงทศนิยมหนึ่งตำแหน่ง และเตือนตั้งแต่ 30; การปัดตัวเลขตอนแสดงไม่เปลี่ยนค่าใน status
- แยกสองไฟล์เพื่ออ่านหน้าที่ง่ายขึ้น เป็นทางเลือก ยังไม่มีความจำเป็นต้องแยกหลาย package

**ข้อผิดพลาดที่พบบ่อย**

- รันเฉพาะ main.go แล้วไม่พบ Reading: ใช้ go run ./examples
- ใช้ข้อความที่ปัดเศษตัดสินสถานะ: ให้ตรวจค่าตัวเลขเดิมก่อนจัดรูปแบบ

[เฉลย](solutions/README.md) · [ผลตรวจและข้อจำกัด](tests/RESULTS.md) · [เอกสาร Go](https://pkg.go.dev/fmt)

</details>

[ตอนก่อนหน้า](../ep24-boundary-tests/README.md) · [สารบัญ](../../../docs/playlist-01-go-basic/README.md) · [ตอนถัดไป](../ep26-project-validation/README.md)
