### เลขขั้นบันได
def Floorloop():
    numbers = int(input("กรุณาป้อนตัวเลข: "))
    for i in range(1,numbers+1,):
        for y in range(1, i+1):
            print(y, end="")
        print("")

### ตารางหมากฮอส
def Checker():
    numbers = int(input("กรุณากรอกขนาดบล็อค : "))
    for i in range(1,numbers+1):
        for y in range(1, numbers+1):
            if (i+y)%2==0:
                print("X", end="")
            else:
                print("O", end="")
            ### ใช้ได้เหมือนกันแต่ยาว
            # if y%2 and i%2 != 0:
            #     print("X", end="")
            # elif y%2!=0 and i%2==0:
            #     print("O", end="")
            # elif y%2==0 and i%2!=0:
            #     print("O", end="")
            # else:
            #     print("X", end="")
        print("")

## สร้างขอบตาราง
def Border_box():
    numbers = int(input("กรุณากรอกขนาดบล๊อค : "))
    for row in range(1, numbers+1):
        print("X", end="")
        for column in range(1, numbers):
            if row == 1 or row == numbers or column == numbers-1:
                print("X", end="")
            else:
                print(" ",end="")
        print("")

# Checker()
Border_box()