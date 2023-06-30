# go-utils
Random GO codes for convenience

In package reverse, 
reverse(), prints out the reverse array of a integer or string array.
if the array is []int{1, 2, 3, 4, 5, 6} or []string{"apple", orange", "banana", "cherry"}

reverse(number) or reverse(fruits) will print out the reverse array only.

ureverse(), returns the output as a struct that allows you to have a variable to contain the reverse array value
array := []int{1, 2, 3, 4, 5, 6}
a:= ureverse(array)
ureverse() function will work with string array.
