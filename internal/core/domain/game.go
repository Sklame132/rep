package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/Sklame132/rep/internal/core/errors"
)

type Game struct {
	ID        string
	FenStart  string
	FenEnd    string
	PlayerW   string
	PlayerB   string
	Type      string
	Mode      string
	Result    string
	History   any
	CreatedAt time.Time
}

func NewGame(
	id string,
	fenStart string,
	fenEnd string,
	playerW string,
	playerB string,
	gameType string,
	mode string,
	result string,
	history any,
	createdAt time.Time,
) Game {
	return Game{
		ID:        id,
		FenStart:  fenStart,
		FenEnd:    fenEnd,
		PlayerW:   playerW,
		PlayerB:   playerB,
		Type:      gameType,
		Mode:      mode,
		Result:    result,
		History:   history,
		CreatedAt: createdAt,
	}
}

func NewGameUnitialized(
	fen_start string,
	fen_end string,
	playerW string,
	playerB string,
	gameType string,
	mode string,
	result string,
	history any,
) Game {
	return NewGame(
		UninitializedString,
		fen_start,
		fen_end,
		playerW,
		playerB,
		gameType,
		mode,
		result,
		history,
		UninitializedTime,
	)
}

func (g *Game) Validate() error {
	re := regexp.MustCompile(`^((?:[pnbrqkPNBRQK1-8]{1,8}\/){7}[pnbrqkPNBRQK1-8]{1,8})\s(?:w|b)\s(?:-|K?Q?k?q?)\s(?:-|[a-h][36])\s\d+\s\d+$`)
	if !re.MatchString(g.FenStart) {
		return fmt.Errorf("invalid `FenStart` format: %w", core_errors.ErrInvalidArgument)
	}
	if !re.MatchString(g.FenEnd) {
		return fmt.Errorf(`invalid "FenEnd" format: %w`, core_errors.ErrInvalidArgument)
	}

	return nil
}
