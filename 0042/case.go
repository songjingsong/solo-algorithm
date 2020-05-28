package case0042

func trap(height []int) int {
	leftmax, rightmax, ans := 0, 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftmax {
				leftmax = height[left]
			} else {
				ans += leftmax - height[left]
			}
			left++
		} else {
			if height[right] > rightmax {
				rightmax = height[right]
			} else {
				ans += rightmax - height[right]
			}
			right--

		}
	}
	return ans
}
