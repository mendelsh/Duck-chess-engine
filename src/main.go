package main

import (
	"fmt"

	"github.com/mendelsh/Duck-chess-engine/board"
)

func main() {
	fmt.Println("White Pawns:")
	board.PrintBitBoard(board.StartingPosition.WhitePawns)
	fmt.Println("White Knights:")
	board.PrintBitBoard(board.StartingPosition.WhiteKnights)
	fmt.Println("White Bishops:")
	board.PrintBitBoard(board.StartingPosition.WhiteBishops)
	fmt.Println("White Rooks:")
	board.PrintBitBoard(board.StartingPosition.WhiteRooks)
	fmt.Println("White Queens:")
	board.PrintBitBoard(board.StartingPosition.WhiteQueens)
	fmt.Println("White King:")
	board.PrintBitBoard(board.StartingPosition.WhiteKing)

	fmt.Println("Black Pawns:")
	board.PrintBitBoard(board.StartingPosition.BlackPawns)
	fmt.Println("Black Knights:")
	board.PrintBitBoard(board.StartingPosition.BlackKnights)
	fmt.Println("Black Bishops:")
	board.PrintBitBoard(board.StartingPosition.BlackBishops)
	fmt.Println("Black Rooks:")
	board.PrintBitBoard(board.StartingPosition.BlackRooks)
	fmt.Println("Black Queens:")
	board.PrintBitBoard(board.StartingPosition.BlackQueens)
	fmt.Println("Black King:")
	board.PrintBitBoard(board.StartingPosition.BlackKing)
}