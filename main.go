package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/whipshout/gator/internal/config"
	"github.com/whipshout/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{db: dbQueries, cfg: cfg}

	cmd := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmd.register("login", handlerLogin)
	cmd.register("register", handlerRegister)
	cmd.register("reset", handlerReset)
	cmd.register("users", handlerGetUsers)
	cmd.register("agg", handlerAgg)
	cmd.register("addfeed", middlewareLoggedIn(handlerCreateFeed))
	cmd.register("feeds", handlerGetFeeds)
	cmd.register("follow", middlewareLoggedIn(handlerFollow))
	cmd.register("following", middlewareLoggedIn(handlerFollowing))
	cmd.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmd.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmd.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
