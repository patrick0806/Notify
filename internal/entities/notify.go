package entities

type NewRecomendation struct {
	Recomendation string `json:"recomendation"`
}

type Notify struct {
	Identity         string           `json:"identity"`
	NewRecomendation NewRecomendation `json:"newRecomendation"`
}
