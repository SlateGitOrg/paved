// Command paved is the internal-developer-platform CLI: software catalog,
// graph queries, scaffolder, scorecards and REST API.
package main

import (
	"os"

	"github.com/archit2501/paved/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
