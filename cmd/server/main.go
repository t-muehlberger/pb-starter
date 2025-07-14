package main

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	_ "muehlberger.dev/pb-starter/internal/migrations"
)

func main() {
	app := pocketbase.New()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	automigrateEnv := strings.ToLower(os.Getenv("PB_AUTOMIGRATE"))
	automigrate := isGoRun && automigrateEnv == "" || automigrateEnv == "true"

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         "internal/migrations",
		Automigrate: automigrate,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
