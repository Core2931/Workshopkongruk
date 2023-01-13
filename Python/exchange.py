# แลกเงินและจำนวนแบงค์
# ธนบัตร 1000, 500, 100, 50, 20, 10, 5, 2, 1

def Exchange():
    money = int(input("ป้อนจำนวนเงินที่ต้องการแลก : "))
    total = money
    if money >=1000:
        print("จำนวนแบงค์ 1000 : %i ใบ" %(money//1000))
        money%=1000

    if money >=500:
        print("จำนวนแบงค์ 500 : %i ใบ" %(money//500))
        money%=500

    if money >=100:
        print("จำนวนแบงค์ 100 : %i ใบ" %(money//100))
        money%=100

    if money >=50:
        print("จำนวนแบงค์ 50 : %i ใบ" %(money//50))
        money%=50
        
    if money >=20:
        print("จำนวนแบงค์ 20 : %i ใบ" %(money//20))
        money%=20
        
    if money >=10:
        print("จำนวนเหรียญ 10 บาท : %i ใบ" %(money//10))
        money%=10

    if money >=5:
        print("จำนวนเหรียญ 5 บาท : %i ใบ" %(money//5))
        money%=5

    if money >=2:
        print("จำนวนเหรียญ 2 บาท : %i ใบ" %(money//2))
        money%=2

    if money >=1:
        print("จำนวนเหรียญ 1 บาท : %i ใบ" %(money//1))
        money%=1
    print("จำนวนเงินทั้งหมดที่ต้องการแลก : %i บาท" %total)

Exchange()