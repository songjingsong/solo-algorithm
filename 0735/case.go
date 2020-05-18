package case0735

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		win := true
		for len(stack) > 0 && asteroids[i] < 0 && stack[len(stack)-1] > 0 {
			temp := stack[len(stack)-1] + asteroids[i]
			if temp < 0 {
				stack = stack[0 : len(stack)-1]
				continue
			}
			if temp == 0 {
				stack = stack[0 : len(stack)-1]
			}
			win = false
			break
		}
		if win {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}
