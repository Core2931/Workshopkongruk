# เกี่ยวกับ String พิมเล็ก พิมใหญ่ ลบช่องว่าง พิมใหญ่ตัวแรกสุด

name = "hello world or or ro or"
name_up = "HELLO UPPERCASE"
name_strip = "  SPACE left & right  "

### แปลงเป็นพิมใหญ่
print(name.upper())

### แปลงเป็นพิมเล็ก
print(name_up.lower())

### แปลงตัวแรกสุดเป็นตัวพิมใหญ่
print(name.capitalize())

### ลบช่องว่าง
print(len(name_strip))
print(name_strip.strip())
#ลบช่องว่างซ้าย ขวา
print(name_strip.lstrip())
print(name_strip.rstrip())

### เปลี่ยนคำแก้คำแทนคำในประโยค replace
print(name.replace("hello","Holaaaaa",1))

### เช็คว่ามีคำนั้นในประโยคหรือป่าว Bool และ เช็ค Count ของคำที่เจอ 
fname = "นาย พัชรพล"
print("wor" in name)
print("or Count: ",name.count("or"))
print("Check คำขึ้นต้น ขึ้นต้นด้วย นาย หรีอไม่ :", fname.startswith("นาย"))
print("Check คำลงท้าย ลงท้ายด้วย มาร์ค หรีอไม่ :", fname.endswith("มาร์ค"))

### แทนค่าในตัวแปร
nickname = "mark"
age= 23
salary = 25400.33223
text = "ชื่อเล่น: {0} \n อายุ: {1}\n เงินเดือน: {3:.2f}\n อื่นๆ : {0}"
print(text.format(nickname, age, "Programmer", salary))


