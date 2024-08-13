package board

func DrawBoard() [8][8]string {

	var points [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			points[i][j] = "x"
		}
	}

	return points
}
