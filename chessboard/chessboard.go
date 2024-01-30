package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	currentFile := cb[file]
	numberOfOccupied := 0
	for _, isOccupied := range currentFile {
		if isOccupied {
			numberOfOccupied++
		}
	}
	return numberOfOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	numberOfOccupied := 0
	for _, file := range cb {
		if rank <= 8 && rank >= 1 {
			if file[rank-1] {
				numberOfOccupied++
			}
		}
	}
	return numberOfOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sum := 0
	for _, file := range cb {
		for _, _ = range file {
			sum += 1
		}
	}
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0
	for file, _ := range cb {
		sum += CountInFile(cb, file)
	}
	return sum
}
