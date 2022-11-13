package DTO

type PatchRequest struct {
	ID    string `json:"id"`
	FIELD string `json:"field"`
	VALUE string `json:"value"`
}
