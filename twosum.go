func twoSum(nums []int, target int) []int {
    matchingMap := map[int]int{}
    for index, value := range(nums){
        complement := target - value
        if index2, ok := matchingMap[complement]; ok{
            return []int{index2, index}
        }
        matchingMap[value] = index
    }
    return []int{}
}
