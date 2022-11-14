import "strings"

func simplifyPath(_path string) string {
    path := strings.Split(_path,"")
    stack :=[]string{}
    
    var i int
    for i < len(path){
        current_path := []string{}
        for i < len(path) && path[i] != "/" {
            current_path = append(current_path,path[i])
            i++
        }
        
        dir := strings.Join(current_path,"")
        

        if dir == "."{
            i++
            continue
        }
        
        if i <len(path) && path[i] =="/"  && len(dir) ==0{
            i++
            continue
        }
        
        if dir == ".." && len(stack) < 1 {
            i++
            continue
        }
        if dir ==".." {
            for len(stack) >0  && stack[len(stack)-1] != "/"{
                stack = stack[:len(stack)-1]
            }
            stack = stack[:len(stack)-1]
            i++
            continue
        }
        stack = append(stack, "/")    
        stack = append(stack, current_path...)
        i++
    }
    if len(stack) <1{
        return "/"
    }
    
    return strings.Join(stack,"")
}


