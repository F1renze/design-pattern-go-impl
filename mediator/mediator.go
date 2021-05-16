package main

import (
	"errors"
	"fmt"
)

type IMediator interface {
	Register(name string, service IService)
	Call(name string, req string) (reply string, err error)
}

type IService interface {
	HandleRequest(req string) (reply string, err error)
}

var ErrServiceNotFound = errors.New("service not found")

func NewMediator() IMediator {
	return &Mediator{
		serviceMap: map[string]IService{},
	}
}

type Mediator struct {
	serviceMap map[string]IService
}

func (s *Mediator) Register(name string, service IService) {
	s.serviceMap[name] = service
}

func (s *Mediator) Call(name string, req string) (reply string, err error) {
	srv, ok := s.serviceMap[name]
	if !ok {
		err = ErrServiceNotFound
		return
	}

	return srv.HandleRequest(req)
}

func NewServiceA(name string) IService {
	return &ServiceA{name: name}
}

type ServiceA struct {
	name string
}

func (s *ServiceA) HandleRequest(req string) (reply string, err error) {
	return fmt.Sprintf("[%s] %s", s.name, req), nil
}
