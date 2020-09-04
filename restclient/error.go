package restclient

type ErrorResponse struct {
	Messages []Error  `json:"errorMessages"`
	Errors   []string `json:"errors"`
}

type Error struct {

}
