
// Input: s = "(*)"
// stack = 1 0 -> true, 1 2 1 -> false, 1 0 -1
// index = 0 1 2 , 2 ,2



// * = ( " " ) tree possible cases

// return stack == 0 
// " " = ignore and continue
// "(" add to stack and continue
// ")" pop from stack and continue

//  ( 
//  1-0
// "" - ( - )
// 1-1 2-1 0-1
// ) ) )
// 0-2 1-2 -1-2

// min 1 0 0
// max 1 2 1
// i  0 1 2

// in max * is always (
// for min * is always ) 
// reset min to 0 
// if max < 0 return false
// if min > 0 return false

func checkValidString(s string) bool {
    var r, min,max int
    
    for r < len(s){
        if s[r] == '('{
            min++
            max++
        }else if s[r] == '*'{
            min--
            max++
        }else{
            min--
            max--
        }
        
        if max < 0{
            return false
        }
        if min < 0{
            min = 0
        }
        r++
    }
    return min == 0
}

