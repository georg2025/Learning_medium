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
    if hundred_number==9{
        answer+="CM"
        hundred_number=0
    }else if hundred_number==4{
        answer+="CD"
        hundred_number=0
    }else if hundred_number>=5{
        answer+="D"
        hundred_number-=5
    }
    for j:=0; j<hundred_number; j++ {
        answer+="C"
    }
    if ten_number==9{
        answer+="XC"
        ten_number=0
    }else if ten_number==4{
        answer+="XL"
        ten_number=0
    }else if ten_number>=5{
        answer+="L"
        ten_number-=5
    }
    for k:=0; k<ten_number; k++ {
        answer+="X"
    }
    if num==9{
        answer+="IX"
        num=0
    }else if num==4{
        answer+="IV"
        num=0
    }else if num>=5{
        answer+="V"
        num-=5
    }
    for k:=0; k<num; k++ {
        answer+="I"
    }
    return answer
}
