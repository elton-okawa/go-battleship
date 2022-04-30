package rest

import (
	"elton-okawa/battleship/internal/entity/board"
	"elton-okawa/battleship/internal/entity/gamestate"
)

type shootResponse struct {
	Hit   bool        `json:"hit"`
	Ships int         `json:"ships"`
	Board board.Board `json:"board"`
}

func (rp RestApiPresenter) StartResult(gs *gamestate.GameState, err error) {
	// if err != nil {
	// 	rp.MapError(err)
	// 	return
	// }

	// // TODO map to a real response
	// rp.responseBody(http.StatusCreated, gs.Board.String())
}

func (rp RestApiPresenter) ShootResult(gs *gamestate.GameState, hit bool, ships int, err error) {
	// if err != nil {
	// 	rp.MapError(err)
	// 	return
	// }

	// shootRes := shootResponse{
	// 	Hit:   hit,
	// 	Ships: ships,
	// 	Board: gs.Board,
	// }

	// fmt.Println(gs.Board.String())
	// resData, _ := json.Marshal(shootRes)
	// rp.responseBody(http.StatusOK, resData)
}
