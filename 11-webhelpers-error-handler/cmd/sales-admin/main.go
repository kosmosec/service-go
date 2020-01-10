package main

import (
	"fmt"
	"github.com/ardanlabs/tomhol/11-webhelpers-error-handler/internal/platform/conf"
	"github.com/ardanlabs/tomhol/11-webhelpers-error-handler/internal/platform/database"
	"github.com/ardanlabs/tomhol/11-webhelpers-error-handler/internal/schema"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: shutting down: %s", err)
		os.Exit(1)
	}
}

func run() error {
	// configuration
	var cfg struct {
		DB struct {
			User       string `conf:"default:postgres"`
			Password   string `conf:"default:postgres,noprint"`
			Host       string `conf:"default:localhost"`
			Name       string `conf:"default:postgres"`
			DisableTLS bool   `conf:"default:false"`
		}
		Args conf.Args
	}

	if err := conf.Parse(os.Args[1:], "SALES", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("SALES", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "error: parsing config")
	}

	// init dep
	db, err := database.Open(database.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	})
	if err != nil {
		return errors.Wrap(err, "error: connecting to db")
	}
	defer db.Close()

	switch cfg.Args.Num(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			return errors.Wrap(err, "error applying migrations")
		}
		fmt.Println("Migrations complete")
		return nil
	case "seed":
		if err := schema.Seed(db); err != nil {
			return errors.Wrap(err, "error seeding db")
		}
		fmt.Println("Seed data complete")
		return nil

	}
	return nil
}
