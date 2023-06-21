package revint

func reverse(arr []int) (rev []int) {
	a := len(arr)
	for i := 1; i <= a; i++ {
		rev = append(rev, arr[a-i])
	}
	return rev
}
