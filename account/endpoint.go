package account

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		
	}
}