// Copyright of Max NG. All Rights Reserved!
package reverse

import "fmt"

type output struct {
	num []int
	str []string
}

func reverse(arr interface{}) {
  switch arr := arr.(type) {
	case []int:
		a := len(arr)
		rev := []int{}
		for i := 1; i <= a; i++ {
      rev = append(rev, arr[a-i])
		}
		fmt.Println(rev)
	case []string:
		a := len(arr)
    rev := []string{}
    for i := 1; i <= a; i++ {
      rev = append(rev, arr[a-i])
    }
    fmt.Println(rev)
	}
}


func ureverse(arr interface{}) *output {
  switch arr := arr.(type) {
  case []int:
    a := len(arr)
    rev := []int{}
    for i := 1; i <= a; i++ {
      rev = append(rev, arr[a-i])
    }
    //fmt.Println(rev)
    return &output{ num: rev}
		
  case []string:
    a := len(arr)
    rev := []string{}
    for i := 1; i <= a; i++ {
      rev = append(rev, arr[a-i])
    }
    //fmt.Println(rev)
    return &output{ str: rev}
    }
return &output{}
}
