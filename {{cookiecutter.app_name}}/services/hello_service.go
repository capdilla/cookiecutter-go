// THIS IS A EXAMPLE FILE

package services

import (
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/repository"
	"github.com/gin-gonic/gin"
)

type HelloService struct {
	helloRepository repository.HelloRepository
}

func NewHelloService(helloRepository repository.HelloRepository) *HelloService {
	return &HelloService{helloRepository: helloRepository}
}

func (s *HelloService) SayHello() string {
	return "OK"
}

func (s *HelloService) GetAll() (*gin.H, error) {

	data, err := s.helloRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return &gin.H{
		"data": data,
	}, nil
}
