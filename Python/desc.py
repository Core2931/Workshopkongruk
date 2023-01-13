# เรียงลำดับตัวเลขที่รับเข้ามา
def Seqnumber():
    seqlist = []
    while True:
        number = int(input("กรุณาระบุตัวเลข : "))
        if number >= 0:
            seqlist.append(number)
            seqlist.sort()
            print("เพิ่มเลขที่คุณระบุ เข้า List เรียบร้อย :",seqlist)
        else:
            print("กรุณาระบุตัวเลข และค่าที่มากกว่า 0 ")
            break
    print("ค่าที่น้อยที่สุด Min : ", min(seqlist))
    print("ค่าที่มากที่สุด Max : ", max(seqlist))
    print("ผลรวมทั้งหมดของใน List : ", sum(seqlist))

# รับค่าเลขคู่และเลขคี่
def Odd_even():
    odd_list = []
    even_list = []
    while True:
        number = int(input("กรุณาระบุตัวเลข : "))
        if number >= 0:
            if number %2== 0:
                even_list.append(number)
                print("เลขคู่ทั้งหมดใน Even list : ", even_list)
            else:
                odd_list.append(number)
                print("เลขคี่ใน Odd list :" ,odd_list)
        else:
            print("คุณได้ระบุค่าที่น้อยกว่า 0")
            break

# เรียงลำดับชื่อ

def Seq_str():
    student = ["สมพร","ก้อง","พัชรพล" ]
    print("ก่อนเรียง => ", student)
    student.sort()
    print("หลังเรียง => ", student)
    student.reverse()
    print("มากไปน้อย =>", student)
    
Seq_str()