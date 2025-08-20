package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	rankSlice, ok := cb[rank]
	if !ok {
		return 0
	}
	nbOccupedSquare := 0
	for _, isSquareOccuped := range rankSlice {
		if isSquareOccuped {
			nbOccupedSquare++
		}
	}
	return nbOccupedSquare
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	nbOccupedSquare := 0
	for _, rank := range cb {
		for column, isSquareOccuped := range rank {
			if column == file-1 && isSquareOccuped {
				nbOccupedSquare++
			}
		}

	}
	return nbOccupedSquare
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	size := 0
	for _, rank := range cb {
		size += len(rank)
	}
	return size
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	nbOccupied := 0
	for rankLetter := range cb {
		nbOccupied += CountInRank(cb, rankLetter)
	}
	return nbOccupied
}
