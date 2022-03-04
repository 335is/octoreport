package main

import (
	"flag"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/335is/log"
	"github.com/335is/octomon/internal/config"
	"github.com/335is/octomon/internal/octopus"
)

const (
	// AppName defines the prefix for any configuration environment variables, as in OCTOMON_OCTOPUS_ADDRESS
	appName    = "octoreport"
	appVersion = "0.0.1"
)

var (
	appInstance string
)

func init() {
	appInstance = fmt.Sprintf("%s", uuid.NewV4())
}

func main() {
	// command line flags
	access := flag.Bool("access", false, "generate a report of all Octopus users and all their members")
	cfg := flag.Bool("config", false, "display configuration loaded")
	info := flag.Bool("info", false, "display information about this program")
	teams := flag.Bool("teams", false, "generate a report of all Octopus teams and their users")
	ver := flag.Bool("ver", false, "display the Octopus server version information")
	flag.Parse()

	if *info {
		DoInformation()
	}

	c := config.New()

	if *cfg {
		DoConfiguration(c)
	}

	octo := octopus.New(c.Octopus.Address, c.Octopus.APIKey, &http.Client{})

	if *ver {
		DoVersionReport(octo)
	}

	if *teams {
		DoTeamReport(octo)
	}

	if *access {
		DoAccessReport(octo)
	}
}

func DoInformation() {
	fmt.Printf("%s %s %s LOG_LEVEL=%s\n", appName, appVersion, appInstance, log.GetLevel().String())
}

func DoConfiguration(cfg *config.Config) {
	fmt.Println(cfg.Dump())
}

func DoVersionReport(octo *octopus.Client) {
	ver, err := octo.Version()
	if err != nil {
		log.Infof(err.Error())
	}
	fmt.Println(ver)
}

func DoTeamReport(octo *octopus.Client) {
	teams, err := octo.GetTeams()
	if err != nil {
		log.Infof(err.Error())
	}

	users, err := octo.GetUsers()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println()
	octopus.TeamMembersReport(teams, users)
}

func DoAccessReport(octo *octopus.Client) {
	teams, err := octo.GetTeams()
	if err != nil {
		log.Infof(err.Error())
	}

	scopedRoles, err := octo.GetScopedUserRoles()
	if err != nil {
		log.Infof(err.Error())
	}
	fmt.Println()
	scopedRoles.Print()

	environments, err := octo.GetEnvironments()
	if err != nil {
		log.Infof(err.Error())
	}
	fmt.Println()
	environments.Print()

	fmt.Println()
	octopus.TeamAccessReport(teams, scopedRoles, environments)
}

/*
	userRoles, err := octo.GetUserRoles()
	if err != nil {
		log.Infof(err.Error())
	}
	//	fmt.Println()
	//	userRoles.Print()
*/
