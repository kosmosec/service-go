package main

import (
	"fmt"
	"github.com/ardanlabs/tomhol/06-configuration/internal/platform/conf"
	"github.com/ardanlabs/tomhol/06-configuration/internal/platform/database"
	"github.com/ardanlabs/tomhol/06-configuration/internal/schema"
	"log"
	"os"
)

func main() {

	// configuration
	var cfg struct {
		DB struct {
			User string `conf:"default:postgres"`
			Password string `conf:"default:postgres,noprint"`
			Host string `conf:"default:localhost"`
			Name string `conf:"default:postgres"`
			DisableTLS bool `conf:"default:false"`
		}
		Args conf.Args
	}

	if err := conf.Parse(os.Args[1:], "SALES", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("SALES", &cfg)
			if err != nil {
				log.Fatalf("Main : generating usage : %v", err)
			}
			fmt.Println(usage)
			return
		}
		log.Fatalf("error: parsing conf: %s", err)
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
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	switch cfg.Args.Num(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		fmt.Println("Migrations complete")
		return
	case "seed":
		if err := schema.Seed(db); err != nil {
			log.Println("error seeding database", err)
			os.Exit(1)
		}
		fmt.Println("Seed data complete")
		return
	}
}
