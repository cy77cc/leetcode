package main

import "fmt"

/*
	1. 给点一个坐标，获取坐标的像素值
	2. 从初始坐标(0,0)开始搜索所有走的通，且像素的值和给定坐标的值的像素值相等，则修改这个位置的像素值
	3. 广度搜索
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var dfs func(image [][]int, sr int, sc int, newColor int, oldColor int)
	dfs = func(image [][]int, sr int, sc int, newColor int, oldColor int) {
		if !(sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] == newColor || image[sr][sc] != oldColor) {
			temp := image[sr][sc]
			image[sr][sc] = newColor
			dfs(image, sr-1, sc, newColor, temp)
			dfs(image, sr, sc+1, newColor, temp)
			dfs(image, sr+1, sc, newColor, temp)
			dfs(image, sr, sc-1, newColor, temp)
		}
	}
	dfs(image, sr, sc, newColor, image[sr][sc])
	return image
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	arr := floodFill(image, 1, 1, 2)
	for _, v := range arr {
		for _, v := range v {
			fmt.Print(v, "")
		}
		fmt.Println()
	}
}
