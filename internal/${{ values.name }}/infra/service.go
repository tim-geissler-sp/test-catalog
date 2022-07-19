// Copyright (c) 2020, SailPoint Technologies, Inc. All rights reserved.
package infra

import (
	"context"
	"github.com/sailpoint/${{ values.name }}/internal/${{ values.name }}/cmd"
	"github.com/sailpoint/atlas-go/atlas"
	"github.com/sailpoint/atlas-go/atlas/application"
)

// Service is the main application structure.
type Service struct {
	*application.Application
	app cmd.App
}

// NewService constructs a new service instance.
func NewService() (*Service, error) { //ConnectService -> Service
	application, err := application.New("${{ values.name }}") //Change to val from wizard
	if err != nil {
		return nil, err
	}

	s := &Service{}
	s.Application = application

	return s, nil
}

// Run will start the server execution, waiting for the system to exit.
func (s *Service) Run() error {
	ctx, done := context.WithCancel(context.Background())

	ar, ctx := atlas.NewRoutineWithContext(ctx)
	ar.Go(ctx, func() error { return s.StartBeaconHeartbeat(ctx) })
	ar.Go(ctx, func() error { return s.StartMetricsServer(ctx) })
	ar.Go(ctx, func() error { return s.StartWebServer(ctx, s.buildRoutes()) })
	ar.Go(ctx, func() error { return s.WaitForInterrupt(ctx, done) })

	if err := ar.Wait(); err != nil && err != context.Canceled {
		return err
	}

	return nil
}
