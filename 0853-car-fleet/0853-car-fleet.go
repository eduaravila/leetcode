func carFleet(target int, position []int, speed []int) int {
    cars := [][]int{}
    
    for i := range position{
        cars = append(cars,[]int{position[i],speed[i]})
    }
    sort.Slice(cars, func(a,b int)bool{
        return cars[a][0] > cars[b][0]
    })
    
    
    res := []float64{}
    for i := range position{
        res = append(res,(float64(target - cars[i][0]))/ float64(cars[i][1])) // remaining step to get to target from current position
        for len(res) > 1 && res[len(res)-2] >= res[len(res)-1]{
            res = res[:len(res)-1]
        }
    }
    return len(res)
}