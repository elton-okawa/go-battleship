package rest

import (
	"elton-okawa/battleship/internal/entity/board"
	"elton-okawa/battleship/internal/entity/gamestate"
	"net/http"
)

type shootResponse struct {
	Hit   bool        `json:"hit"`
	Ships int         `json:"ships"`
	Board board.Board `json:"board"`
}

func (rp *RestApiPresenter) StartResult() {
	rp.response(http.StatusCreated)
}

func (rp *RestApiPresenter) ShootResult(gs *gamestate.GameState, hit bool, ships int, err error) {
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
