package odd

func Process(numberList []int) []int {

	numberMap := make(map[int]int)
	var oddList []int

	// loop เก็บข้อมูลใส่ map โดยให้ key = ข้อมูลตัวเลข และ value = จำนวนตัวเลข
	for _, number := range numberList {
		numberMap[number] = numberMap[number] + 1
	}

	// วนข้อมูล map เพื่อคำนวณหาว่าข้อมูลของเลขตัวไหน ที่มีจำนวนเป็นเลขคี่
	for key, val := range numberMap {
		if val%2 != 0 {
			oddList = append(oddList, key)
		}
	}

	// กรณีที่ไม่มีเลขตัวใดมีจำนวนเป็นเลขคี่ ให้ทำการ ส่งค่า [0] ออกไป
	if len(oddList) == 0 {
		return []int{0}
	}

	return oddList
}
