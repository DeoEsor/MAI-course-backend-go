package main

import (
	"context"
	log "github.com/rs/zerolog"
	"os"
	"time"

	"github.com/DeoEsor/MAI-course-backend-go/internal"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
)

func main() {
	ctx := context.Background() // TODO graceful
	cfg, err := config.Load()
	logger := log.New(log.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Caller().
		Ctx(ctx).
		Logger()

	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("error while parsing config")
	}

	app, err := internal.New(ctx, cfg)
	if err != nil {
		logger.Fatal().
			Err(err)
	}
	if err = app.Run(ctx); err != nil {
		logger.Fatal().
			Err(err)
	}
}

// mvc (model view controller)
// 			model (model данные) view(слой представления) controller (логика)
//    Startup, Small business

// mvc (model view component)
// Front,  Unity GameDev

// mvvm (model view view-model)
//      Desktop, Microsoft, Web API
///-------------
// layered architecture
//     Go, ASP NET, JAVA, Python microservices
// 3-4 layered
// Api, Domain (Buisness Layer), Infrastructure

// clean architecture
//     Go, ASP NET, JAVA, Python (Controller, Consumer, Domain, Use Case,  Services, Repository, Producers, External) microservices

// onion architecture
//     Go, ASP NET, JAVA, Python (Presentation, Domain, Application, Infrastructure, External) microservices
