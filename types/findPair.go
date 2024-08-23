package types

type FindPairreq struct {
	Number []int `json:"number"`
	Target int   `json:"target"`
}

type FindPairRes struct {
	Solutions [][]int `json:"solution"`
}
