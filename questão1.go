package main

import "fmt"

func check(W int, H int, N int, pairsMatrix [][4]int) string { //RETURNS Y/N FOR FAIR OR UNFAIR
	var x, y int //ACCOUNTANTS
	for i := 0; i < N; i++ {
		for j := 0; j < 4; j++ {
			if pairsMatrix[i][j] == W || pairsMatrix[i][j] == 0 { //EVALUATES WHETHER THE POSITION IS AT ONE OF THE END
				x++
			}
			if pairsMatrix[i][j] == H {
				y++
			}
			if x >= W/2 && y >= H/2 { //EVALUATE ONLY THE X AXIS
				return "N" //RETURN NO
			}
		}
	}
	return "Y" //RETURN YES
}

func main() {
	var W, H, N int  //DEFINING WIDTH, HEIGHT AND PAIRS
	fmt.Scan(&W, &H) //CAPTURE WIDTH AND HEIGHT
	fmt.Scan(&N)     //CAPTURE NUMBER PAIRS

	pairsMatrix := make([][4]int, N) //DEFINING PAIRS MATRIX

	for i := 0; i < N; i++ { //DEFINING MATRIX COMPONENTS
		fmt.Scan(&pairsMatrix[i][0], &pairsMatrix[i][1], &pairsMatrix[i][2], &pairsMatrix[i][3])
	}

	fmt.Println(check(W, H, N, pairsMatrix)) //EVALUATE Y/N
}
