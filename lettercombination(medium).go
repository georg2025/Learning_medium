//this function take sting with numbers 2-9 and find out, what combinations of words can it make on old mobile phones (where number 2 had letters a,b and c; number 3 - letters d, e and f). Max - 4 numbers
func letterCombinations(digits string) []string {
    x:=[]string{}
    z:=strings.Split(digits, "")
    if len(z)==0 {
        return x
    }else if len(z)==1 {
        return numberletter(z[0])
    }else if len(z)==2 {
        for _,k:= range numberletter(z[0]){
                for _,l:= range numberletter(z[1]){
                    n:=k+l
                    x=append(x, n)
                }
            }
    }else if len(z)==3 {
        for _,j:= range numberletter(z[0]){
            for _,k:= range numberletter(z[1]){
                for _,l:= range numberletter(z[2]){
                    n:=j+k+l
                    x=append(x, n)
                }
            }
        }
    }else if len(z)==4 {
    for _,i:= range numberletter(z[0]){
        for _,j:= range numberletter(z[1]){
            for _,k:= range numberletter(z[2]){
                for _,l:= range numberletter(z[3]){
                    n:=i+j+k+l
                    x=append(x, n)
                }
            }
        }

    }
    }
return x
}




func numberletter (x string) []string {
    z:=[]string{}
	switch x {
    case "2":
	z=[]string{"a", "b", "c"}
    case "3":
	z=[]string{"d", "e", "f"}
    case "4":
	z=[]string{"g", "h", "i"}
    case "5":
	z=[]string{"j", "k", "l"}
    case "6":
	z=[]string{"m", "n", "o"}
    case "7":
	z=[]string{"p", "q", "r", "s"}
    case "8":
	z=[]string{"t", "u", "v"}
    case "9":
	z=[]string{"w", "x", "y", "z"}
}
return z
}
