package dbgamestate

import (
	"elton-okawa/battleship/internal/entity/board"
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/infra/database"
	"elton-okawa/battleship/internal/test"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// Mudar o time do Entity para int64 mesmo -> objeto time falha na comparação e não aarrumaram
// arrumar os erros nos outros testes que quebrei refatorando
func TestSuite_GameState(t *testing.T) {
	suite.Run(t, new(TestGameStateSuite))
}

type TestGameStateSuite struct {
	suite.Suite
}

func (s *TestGameStateSuite) SetupTest() {
	test.CleanupDatabase()
}

func (s TestGameStateSuite) TestSaveAndRead() {
	gsId := "id"
	pOneId := "accOneId"
	pTwoId := "accTwoId"

	bOne := board.Initialize(8, 3)
	bTwo := board.Initialize(8, 3)
	history := shootMany(
		[]string{pOneId, pTwoId},
		[]*board.Board{bOne, bTwo},
		true, false, false, true, true, false, false,
	)

	gs := gamestate.New(
		gsId,
		pOneId,
		pTwoId,
		bOne,
		bTwo,
		history,
		pTwoId,
		true,
	)

	repoOption := database.RepositoryOption{Path: test.DbFilePath()}
	db := New(repoOption.File("game-state"))

	saveErr := db.Save(gs)
	s.Nilf(saveErr, "unexpected error %v", saveErr)

	savedGs, readErr := db.Get(gsId)
	s.Nilf(readErr, "unexpected error %v", readErr)

	s.Equal(gs, savedGs)
}

func shootMany(playersId []string, boards []*board.Board, hits ...bool) []gamestate.Turn {
	hist := make([]gamestate.Turn, len(hits))

	// last place searched to avoid same spot
	startHitCoord := make([][2]int, 2)
	startMissCoord := make([][2]int, 2)

	baseTime := time.Now().Unix()

	for i, hit := range hits {
		startCoord := startMissCoord
		mark := board.MISS
		if hit {
			startCoord = startHitCoord
			boards[(i+1)%2].ShipCount -= 1
			mark = board.HIT
		}

		row, col := coordinateTo(boards[(i+1)%2], hit, startCoord[i%2])
		boards[(i+1)%2].State[row][col] = mark
		startCoord[i%2][0] = row
		startCoord[i%2][1] = col

		hist[i] = gamestate.NewTurn(
			strconv.Itoa(i),
			playersId[i%2],
			row,
			col,
			hit,
			baseTime,
		)

		baseTime += 1
	}

	return hist
}

func coordinateTo(b *board.Board, hit bool, startCoord [2]int) (row, col int) {
	mark := board.EMPTY
	if hit {
		mark = board.SHIP
	}

	for i := startCoord[0]; i < b.Size; i++ {
		for j := startCoord[1]; j < b.Size; j++ {
			if b.Placement[i][j] == mark && b.State[i][j] == board.EMPTY {
				return i, j
			}
		}
	}

	return -1, -1
}
