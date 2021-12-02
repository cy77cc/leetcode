package main

func reverseString(s []byte)  {
	//left, right := 0, len(s)-1
	//for left <= right {
	//	s[left], s[right] = s[right], s[left]
	//	left++
	//	right--
	//}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func main() {

}
