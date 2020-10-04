package dojo

func checkCentralRoot(value, position int, arr []int) bool {
	return value == arr[position]
}

func binarySearch(value int, arr []int) bool {
	var resp bool
	length := len(arr) - 1
	halfPosition := length / 2

	if length == 0 {
		return false
	}

	for {

		if checkCentralRoot(value, halfPosition, arr) {
			resp = true
			break

		} else {

			if value < arr[halfPosition] {
				halfPosition = halfPosition / 2
			} else {

				halfPosition = ((length + halfPosition) / 2)

			}
		}
		if halfPosition == 0 || halfPosition == length {

			resp = checkCentralRoot(value, halfPosition, arr)
			break
		}

	}

	return resp
}
