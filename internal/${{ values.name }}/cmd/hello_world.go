// Copyright (c) 2020, SailPoint Technologies, Inc. All rights reserved.
package cmd

import (
	"context"

	"github.com/sailpoint/${{ values.name }}/internal/${{ values.name }}/model"
)

type HelloWorld struct {
	genericModelID model.GenericModel
}

// NewHelloWorld Create new helloworld object
func NewHelloWorld() (*HelloWorld, error) {
	cmd := &HelloWorld{}
	return cmd, nil
}

//Handle return "hello world" string
func (cmd *HelloWorld) Handle(ctx context.Context) (string, error) {
	return "Hello World!", nil
}
