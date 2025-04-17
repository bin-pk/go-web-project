package service

import (
	"go-web-project/core"
	"go-web-project/repos"
)

type HelloService struct {
	input HelloInput
}

func (ptr *HelloService) AddParameter(input core.CuteData) *core.CuteError {
	helloInput, ok := input.(*HelloInput)
	if !ok {
		cuteErr := &core.CuteError{Message: "invalid input type!!!", Code: core.InternalError}
		return cuteErr
	} else {
		ptr.input = *helloInput
		return nil
	}
}
func (ptr *HelloService) AddRepository(db repos.CuteDB) *core.CuteError {
	return nil
}
func (ptr *HelloService) Execute(ctx core.CuteContext) (core.CuteData, *core.CuteError) {
	output := &HelloOutput{Data: "Hello World!"}
	return output, nil
}
func (ptr *HelloService) Destroy(ctx core.CuteContext) {
}

type OSMSearcherService struct {
}
