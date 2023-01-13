
#เปรียบเทียบตัวเลข มากกว่า น้อยกว่า
def Compare_numbers():
    number_a = float(input("ระบุตัวเลขที่ A :"))
    number_b = float(input("ระบุตัวเลขที่ B :"))

    if number_a > number_b:
        print("A มีค่ามากกว่า B")
    elif number_b > number_a:
        print("B มีค่ามากกว่า A")
    elif number_b == number_a:
        print("ทั้งA และ B มีค่าเท่ากัน")
    else :
        print("พบข้อผิดพลาดไม่สามารถเปรียบเทียบค่าได้")

#หาเลขคู่ / เลขคี่
def Find_odd_even():
    number_1 = int(input("ระบุตัวเลขที่ 1 :"))
    
    if number_1 != 0 and number_1%2 == 0:
        print(number_1,"เป็นเลขคู่")
    elif number_1 != 0 and number_1%2 != 0:
        print(number_1, "เป็นเลขคี่")
    else:
        print("พบข้อผิดพลาด")

Find_odd_even()
