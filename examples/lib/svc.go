package lib

import (
	"net/http"

	"github.com/iceyang/boom"
)

type Service struct{}

func (s *Service) FindOne() *User {
	return &User{"Justin"}
}

func (s *Service) PanicBoomError() {
	panic(boom.Boom(http.StatusInternalServerError, "数据库出错"))
}

func NewService() *Service {
	return &Service{}
}
