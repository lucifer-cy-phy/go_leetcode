func init() {
    _ = os.Stdout
}

func twoSum(nums []int, target int) []int {
    storage := make(map[int]int32, len(nums))

  for i := 0; i < len(nums); i++ {
    num := nums[i]
    complement := target - num

    if origIndex, found := storage[complement]; found {
             return []int{int(origIndex), i}
    }
    storage[num] = int32(i)
  }  
  
  return nil
}
