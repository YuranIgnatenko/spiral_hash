package HSP

//convert data in spiral, spiral to x,y points
func GetHash(data []int) (int, int) {
	x := 0
	y := 0
	c := 0
	for counter, value := range data {
		value += counter + 1
		switch c {
		case 0:
			y += value
		case 1:
			x += value
		case 2:
			y -= value
		case 3:
			x -= value
		}
		c += 1
		if c == 4 {
			c = 0
		}
	}
	return x, y
}

//convert enum int
//Array(1,2,3) -> []int{1,2,3}
func Array(args ...int) []int {
	dt := make([]int, 0)
	for _, val := range args {
		dt = append(dt, val)
	}
	return dt
}
