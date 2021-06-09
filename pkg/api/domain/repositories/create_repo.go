package repositories

// this is request sent by client of our microservice
type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

// this is what we will give back to client using our microservice
type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}