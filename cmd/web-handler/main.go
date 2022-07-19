// Copyright (c) 2020. Sailpoint Technologies, Inc. All rights reserved.
package main

import (
	"github.com/sailpoint/${{ values.name }}/internal/${{ values.name }}/infra"
	log "github.com/sailpoint/atlas-go/atlas/log"
)

func main() {
	service, err := infra.NewService()
	if err != nil {
		log.Global().Sugar().Fatalf("init: %v", err)
	}
	defer service.Close()

	if err := service.Run(); err != nil {
		log.Global().Sugar().Fatalf("run: %v", err)
	}
}
