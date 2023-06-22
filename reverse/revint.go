// Copyright 2023 Max NG. All Rights Reserved!
package reverse

func revint(arr []int) (rev []int) {
	a := len(arr)
	for i := 1; i <= a; i++ {
		rev = append(rev, arr[a-i])
	}
	return rev
}
