# ผลตรวจสภาพแวดล้อม

ตำแหน่งหลักสูตรปัจจุบัน: `C:\Users\kopes\Documents\L\Go`

## อัปเดต 2026-09-19 — ติดตั้งและรันได้แล้ว

- ผู้เรียนอนุญาตให้ตรวจและติดตั้ง Go ในข้อความวันที่ 2026-09-19
- ดาวน์โหลด Go 1.27.1 windows-amd64 จาก go.dev ตรวจ SHA256 ตรงกับ metadata ทางการ
- MSI ตรวจลายเซ็น Google LLC ผ่าน แต่ติดตั้งไม่สำเร็จ: Windows Installer error 1925 (สิทธิ์ไม่พอสำหรับทุกบัญชี), exit code 1603
- จึงติดตั้งจาก ZIP ทางการเฉพาะบัญชีผู้ใช้ที่ `C:\Users\kopes\AppData\Local\Programs\Go-1.27.1\go` ตรวจ SHA256 ZIP: `a3911b5e0e1b1053f25ed0675f4c1c6aad1e2bfcf253df2b9be4caabd2edd95d`
- เพิ่มเฉพาะ `C:\Users\kopes\AppData\Local\Programs\Go-1.27.1\go\bin` ใน User PATH โดยเก็บค่าเดิมไว้ใน TEMP และไม่แทนที่รายการอื่น ไม่มีการแก้ execution policy หรือปิดระบบป้องกัน
- สำเนา PATH เดิม: `%TEMP%\go-course-setup-20260919\user-path-before-d3a5934749f74f53bf4b4d4aaf4c26d3.txt`
- ผลจริง: `go version go1.27.1 windows/amd64` ตรวจรูปแบบ/compile/vet/run/build ผ่าน ผลครั้งนั้นเป็นประวัติการติดตั้ง ดูผลหลักสูตรปัจจุบันใน [phase-01-results.md](phase-01-results.md)
- การเรียกครั้งทดสอบโหลด PATH ล่าสุดเข้าหน้าต่างทดสอบแล้ว ผู้เรียนยังต้องลอง `go version` ใน PowerShell ใหม่ของตนเอง จึงจะยืนยันหน้าต่างฝั่งผู้เรียนได้
- ใช้ TEMP เก็บไฟล์ติดตั้งและ build เพราะการสร้างไฟล์ด้วยโปรแกรมภายใต้ Documents ล้มเหลว ยังไม่สรุปว่ามาจากซอฟต์แวร์ป้องกันตัวใด

## ประวัติการตรวจครั้งแรก (ก่อนติดตั้ง)

วันที่ 2026-09-17 (Asia/Bangkok)

| รายการ | สิ่งที่ตรวจพบจริง |
|---|---|
| ระบบ | Windows NT 10.0.26200.0, สถาปัตยกรรม X64 |
| PowerShell | 7.6.5 |
| Git | `git version 2.50.0.windows.1` |
| Git executable | `C:\Program Files\Git\cmd\git.exe` |
| Go ใน PATH | `Get-Command go` ไม่พบ; ลอง `go version` แล้วคำสั่งไม่เป็นที่รู้จัก |
| Go ตำแหน่งทั่วไป | ไม่พบ executable ทั้ง 5 ตำแหน่งที่ระบุด้านล่าง |
| Registry รายการติดตั้ง | ไม่พบชื่อที่ตรงกับ Go Programming Language / Golang / Go version ใน uninstall keys ที่ค้น |
| โฟลเดอร์ปลายทางที่ตรวจเมื่อ 2026-09-17 | ยังว่างในวันตรวจ ต่อมาใช้เป็นที่เก็บหลักสูตร |

ตำแหน่งที่ตรวจ: `C:\Program Files\Go\bin\go.exe`, `C:\Program Files (x86)\Go\bin\go.exe`, `C:\Go\bin\go.exe`, `%LOCALAPPDATA%\Programs\Go\bin\go.exe`, `%USERPROFILE%\scoop\apps\go\current\bin\go.exe`

ข้อจำกัด: ไม่ได้ค้นทั้งดิสก์ อาจมี Go แบบ portable อยู่ที่อื่น ผลนี้หมายถึงยังไม่มี Go ที่เรียกใช้งานได้ในการตรวจครั้งนี้ ไม่ใช่ข้อยืนยันว่าไม่มีสำเนาใด ๆ บนเครื่อง

Registry แสดง ProductName `Windows 10 Pro`, DisplayVersion `25H2`, CurrentBuildNumber `26200` ชื่อรุ่นอาจเป็นค่าที่คงอยู่จากระบบเดิม จึงรายงาน Windows พร้อม build ตามที่อ่านได้ การขอข้อมูลยืนยันชื่อ edition ผ่าน `Get-CimInstance Win32_OperatingSystem` ถูกปฏิเสธการเข้าถึง ไม่มีการยกระดับสิทธิ์เพื่ออ่าน ข้อมูลที่มีเพียงพอเลือกตัวติดตั้ง Windows amd64

ไม่ได้ติดตั้งโปรแกรม แก้ PATH ตั้งค่า Git เริ่มบริการ เปิดพอร์ต หรือติดต่ออุปกรณ์ใด ๆ
