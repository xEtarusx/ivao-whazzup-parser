package connections

type Connections struct {
	ATCs         int64 `json:"atc"`
	FollowMeCars int64 `json:"followMe"`
	Observers    int64 `json:"observers"`
	Pilots       int64 `json:"pilot"`
	Supervisors  int64 `json:"supervisors"`
	WorldTours   int64 `json:"worldTours"`
	Total        int64 `json:"total"`
}
