package demo

import (
	demoModels "go-api-template/domain/models/demo"
	demoRepo "go-api-template/repositories/demo"
)

func DemoSvc(demo demoModels.DemoRequest) (*demoModels.DemoResPonse, error) {

	_, err := demoRepo.GetDemoRepo("Hello")
	if err != nil {
		return nil, err
	}

	res := demoModels.DemoResPonse{
		Message: "Hello",
	}

	return &res, nil
}
