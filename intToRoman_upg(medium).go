func intToRoman(num int) string {
    answer:=""
    thousant_number:=num/1000
    num=num-thousant_number*1000
    hundred_number:=num/100
    num=num-hundred_number*100
    ten_number:=num/10
    num=num-ten_number*10
    for i:=0; i<thousant_number; i++ {
        answer+="M"
    }
    answer+=Romanhelp (hundred_number, "C", "D", "M")
    answer+=Romanhelp (ten_number, "X", "L", "C")
    answer+=Romanhelp (num, "I", "V", "X")
    return answer
}

func Romanhelp (x int, y,z,n string) string {
    answer:=""
    if x==9{
        answer+=y+n
        x=0
    }else if x==4{
        answer+=y+z
        x=0
    }else if x>=5{
        answer+=z
        x-=5
    }
    for j:=0; j<x; j++ {
        answer+=y
    }
    return answer

}
