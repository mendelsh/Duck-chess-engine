package board

import (
	"fmt"
)

const (
	A1 = iota; A2; A3; A4; A5; A6; A7; A8
	B1; B2; B3; B4; B5; B6; B7; B8
	C1; C2; C3; C4; C5; C6; C7; C8
	D1; D2; D3; D4; D5; D6; D7; D8
	E1; E2; E3; E4; E5; E6; E7; E8
	F1; F2; F3; F4; F5; F6; F7; F8
	G1; G2; G3; G4; G5; G6; G7; G8
	H1; H2; H3; H4; H5; H6; H7; H8
) // Enum for the board squares

type Position struct {
	WhitePawns, WhiteKnights, WhiteBishops, WhiteRooks, WhiteQueens, WhiteKing uint64
	BlackPawns, BlackKnights, BlackBishops, BlackRooks, BlackQueens, BlackKing uint64
}


var StartingPosition Position = Position{
	WhitePawns:   0b_00000000_00000000_00000000_00000000_00000000_00000000_11111111_00000000,
	WhiteKnights: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_01000010,
	WhiteBishops: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00100100,
	WhiteRooks:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_10000001,
	WhiteQueens:  0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00001000,
	WhiteKing:    0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00010000,

	BlackPawns:   0b_00000000_11111111_00000000_00000000_00000000_00000000_00000000_00000000,
	BlackKnights: 0b_01000010_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
	BlackBishops: 0b_00100100_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
	BlackRooks:   0b_10000001_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
	BlackQueens:  0b_00001000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
	BlackKing:    0b_00010000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
}  

// Debug function to print the bitboard
func PrintBitBoard(board uint64) {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			// Get the bit for the current square
			if board&(1<<uint(8*rank+file)) != 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println() 
	}
	fmt.Println()
}

func (b *Position) Copy() Position {
	return *b
}

func NewBoard() *Position {
	sp := StartingPosition.Copy()
	return &sp
}



