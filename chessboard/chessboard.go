package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	counter := 0
	if _, ok := cb[rank]; !ok {
		return counter
	}
	for _, v := range cb[rank] {
		if v {
			counter++
		}
	}
	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	counter := 0
	if file < 1 || file > 8 {
		return counter
	}
	for _, v := range cb {
		if v[file-1] {
			counter++
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, rank := range cb {
		for _, v := range rank {
			if v {
				counter++
			}
		}
	}
	return counter
}
