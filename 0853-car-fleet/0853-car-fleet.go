import "sort"

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}
func carFleet(target int, position []int, speed []int) int {
    cars := [][]int{}
    for i := range position{
        cars = append(cars,[]int{position[i],speed[i]})
    }
    sort.Slice(cars, func(a,b int)bool{
        return cars[a][0] >= cars[b][0] 
    })
        
    
    var i int
    stack := []float64{}   
    for i < len(position) {    
        stack = append(stack,((float64(target - cars[i][0])) / float64(cars[i][1])))
        if len(stack) > 1 && stack[len(stack)-1] <= stack[len(stack)-2] {
            stack = stack[:len(stack)-1]
        }
        
        i++            
    }    
    return len(stack)
}


// position = [10,8,0,5,3], speed = [2,4,1,1,3]

// move firts the values on top
// [10,8,5,3,0] , [2,4,1,3,1] // sorted by position, update speed 
// stack = [12,6,1] // 


//[0,2,4], speed = [4,2,1]
// [4,2,0] [1,2,4]

// s = [5,4]