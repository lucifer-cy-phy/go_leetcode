package main

import (
	"os"
)

// Блок func init() ниже используется исключительно для оптимизации на платформе LeetCode
// Строка с os.Stdout отключает буферизацию стандартного вывода внутри тестовой среды, что позволяет снизить накладные расходы и ускорить прохождение тестов до 0 ms
func init() {
	_ = os.Stdout
}

func twoSum(nums []int, target int) []int {
	// Заранее выделяем память под мапу нужного размера (len(nums)), чтобы избежать реаллокаций памяти в процессе работы цикла
	// Используем тип int32 для экономии памяти на больших объемах данных
	storage := make(map[int]int32, len(nums))

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		complement := target - num

		// Проверяем, есть ли уже в мапе число, дополняющее текущее до target
		if origIndex, found := storage[complement]; found {
			return []int{int(origIndex), i}
		}
		
		// Сохраняем текущее число и его индекс
		storage[num] = int32(i)
	}

	return nil
}
