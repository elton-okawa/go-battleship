package gamerequest

type GameRequest struct {
	Id           string
	OwnerId      string
	ChallengerId string
	Pending      bool
}

func New(id, owner, chal string, pending bool) GameRequest {
	return GameRequest{
		Id:           id,
		OwnerId:      owner,
		ChallengerId: chal,
		Pending:      pending,
	}
}
