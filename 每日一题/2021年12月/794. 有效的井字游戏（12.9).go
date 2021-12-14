package main
func validTicTacToe(board []string) bool {
	var count = []int{0, 0}  // X, O

	for _, row := range board{
		for _, char := range row{
			if char == 'O'{
				count[1]++
			}else if char == 'X'{
				count[0]++
			}
		}
	}

	if (count[1] == 1 && count[0] == 0) || (count[1] - count[0] >0) || (count[0] - count[1] >1){
		return false
	}

	for i:=0;i<3;i++{
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][2] != ' '{
			if board[i][0] == 'O' && count[0] > count[1]{
				return false
			}
			if board[i][0] == 'X' && count[1] >= count[0]{
				return false
			}
		}

		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[2][i] != ' '{
			if board[0][i] == 'O' && count[0] > count[1]{
				return false
			}
			if board[0][i] == 'X' && count[1] >= count[0]{
				return false
			}
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] != ' '{
		if board[1][1] == 'O' && count[0] > count[1]{
			return false
		}
		if board[1][1] == 'X' && count[1] >= count[0]{
			return false
		}

	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] != ' '{
		if board[1][1] == 'O' && count[0] > count[1]{
			return false
		}
		if board[1][1] == 'X' && count[1] >= count[0]{
			return false
		}
	}

	return true

}

