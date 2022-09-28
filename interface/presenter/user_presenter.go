package presenter

type UserPresenter struct{}
type userResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
