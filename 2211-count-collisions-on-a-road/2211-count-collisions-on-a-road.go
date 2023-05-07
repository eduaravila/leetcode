func countCollisions(directions string) int {
    stack := []rune{}
    var collisions, prevR int 
    for _, c := range directions {
        n := len(stack)-1
        if len(stack) >0 && ((stack[n] == 'S' && c == 'L') ||  (stack[n] == 'R' && c == 'S')){
            collisions++ 
            if stack[n]=='R'{
                prevR--
            }
            stack[n] = 'S'
        }else if len(stack) > 0 && stack[n] == 'R' && c == 'L' {
            collisions+=2
            stack[n] = 'S'
            prevR--
        }else if c == 'R'  || c == 'S' {
            if c == 'R'{
                prevR++
            }
            stack = append(stack, c)
        }
        if len(stack) > 0 && stack[len(stack)-1] == 'S'{
            collisions+= prevR
            prevR=0
        } 
    }
    return collisions
}

/*
 left and right
 right and left
 +2
 
 
 stop and left 
 stop and right 
 +1
*/