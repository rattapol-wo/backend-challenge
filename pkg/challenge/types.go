package challenge

type requestNumberSummary struct {
	ListNumber [][]int `json:"listNumber" validate:"required"`
}