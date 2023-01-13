

def Mutiple_table():
    number = int(input("กรุณาป้อนเลขแม่สูตรคูณที่ต้องการจะดู : "))
    for i in range(1, 13):
        print("%i x %i =" %(number, i), number*i)

def All_table():
    for i in range(1, 13):
        for j in range(1, 13):
            print("%i x %i =" %(i, j), i*j)

All_table()