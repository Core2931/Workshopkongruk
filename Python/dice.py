import random

def Random():
    number = random.randrange(1,99)
    while True:
        numbers = int(input("กรุณาป้อนคำตอบที่คุณคิดว่าใช่ : "))
        if numbers < number:
            print("มากกว่านั้น T_T")
        elif numbers > number:
            print("น้อยกว่านั้น T_T")
        elif numbers == number:
            print("ถูกต้องแล้วค้าบบ เลขที่ถูกต้องคือ %i" %number)
            break
        else:
            print("ข้อผิดพลาด")
     
    print("จบการทำงานโปรแกรมทายเลข")

def Rice():
    number = random.randrange(1,6)
    print("เลขที่คุณทอยได้คือ : %i" %(number))
