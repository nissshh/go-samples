package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//make a flexi slice with provided params
	retPic := make([][]uint8, dx)
	//init
	for i := range retPic {
		retPic[i]=  make([]uint8,dy)
	}
	//set for each i,j an integer
	for n := range retPic{
		for m := range retPic[n]{
			retPic[n][m]= uint8(n*m/2);
		}
	}
	//return
	return retPic	
}

func main() {
	pic.Show(Pic)
}
