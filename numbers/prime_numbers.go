package numbers

import "math"

// нахождение простых множителей числа
// 1. нужно рассмотреть лишь делимость на 2 и на нечетные числа
// 2. проверять множители только до квадратного корня числа
// 3. при делении числа на множитель обновлять максимальное количество
// потенциальных множителей, которые необходимо проверить
// время работы O(sqrt(N))
func FindFactors(num int) []int {
	factors := make([]int, 0)
	for num%2 == 0 {
		factors = append(factors, 2)
		num = num / 2
	}

	maxFactor := int(math.Sqrt(float64(num)))
	for i := 3; i <= maxFactor; i += 2 {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
			maxFactor = int(math.Sqrt(float64(num)))
		}
	}

	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}
