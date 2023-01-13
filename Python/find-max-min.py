


def Find_Max_min():
    max, min = 0, 0
    while True:
        number = int(input("กรุณป้อนตัวเลข : "))
        min = number
        if number > 0:
            if number > max:
                max = number
            if number < min:
                min = number
        else:
            break
    print("max = ", max, "\nmin = ",min)
Find_Max_min()