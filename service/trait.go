package service

import (
	"fmt"
	"go-web-project/core"
	"go-web-project/repos"
)

type CuteService interface {
	AddParameter(input core.CuteData) *core.CuteError
	AddRepository(db repos.CuteDB) *core.CuteError
	Execute(ctx core.CuteContext) (core.CuteData, *core.CuteError)
	Destroy(ctx core.CuteContext)
}

type CuteServiceMap struct {
	Constructors map[string]CuteService
}

func CreateServiceMap() CuteServiceMap {
	return CuteServiceMap{Constructors: make(map[string]CuteService)}
}

func (ptr *CuteServiceMap) AddService(key string, service CuteService) {
	ptr.Constructors[key] = service
}

func (ptr *CuteServiceMap) GetService(key string, input core.CuteData) (CuteService, *core.CuteError) {
	service, exists := ptr.Constructors[key]
	if exists {
		service.AddParameter(input)
		return service, nil
	} else {
		return nil, &core.CuteError{
			Message: fmt.Sprintf("service key '%s' is not found", key),
			Code:    core.InternalError,
		}
	}
}
