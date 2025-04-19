package hello

type Service struct{}
type Response struct {
	Message string `json:"message"`
}

func CreateService() *Service {
	return &Service{}
}

func (s Service) Get() Response {
	return Response{
		Message: "Hello, World",
	}
}
