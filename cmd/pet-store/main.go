package main

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/DeoEsor/MAI-course-backend-go/internal"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
)

func main() {
	ctx := context.Background() // TODO graceful
	cfg, err := config.Load()
	if err != nil {
		// structured logging
		log.Fatalf("error while parsing config: %v", err)
	}

	app, err := internal.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err = app.Run(ctx); err != nil {
		log.Fatal(err)
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
