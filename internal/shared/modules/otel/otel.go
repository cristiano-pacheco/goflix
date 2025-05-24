package otel

import (
	"log"

	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/pkg/otel/trace"
)

var (
	_global      Otel
	_initialized bool
)

type Otel struct {
	trace.Trace
}

func Init(cfg config.Config) {
	_global = new(cfg)
	_initialized = true
}

func get() Otel {
	if !_initialized {
		log.Fatalf("otel not initialized")
	}
	return _global
}

func Trace() trace.Trace {
	return get().Trace
}

func new(config config.Config) Otel {
	tc := trace.TracerConfig{
		AppName:      config.App.Name,
		AppVersion:   config.App.Version,
		TraceEnabled: config.Telemetry.Enabled,
		TracerVendor: config.Telemetry.TracerVendor,
		TraceURL:     config.Telemetry.TracerURL,
	}

	return Otel{
		Trace: trace.New(tc),
	}
}
