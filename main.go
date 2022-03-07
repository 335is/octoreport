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
	// AppName defines the prefix for any configuration environment variables, as in OCTOREPORT_OCTOPUS_ADDRESS
	appName    = "octoreport"
	appVersion = "0.0.2"
)

var (
	appInstance string
)

func init() {
	appInstance = fmt.Sprintf("%s", uuid.NewV4())
}

func main() {
	// command line flags
	help := flag.Bool("help", false, "display usage of this program")
	info := flag.Bool("info", false, "display information about this program")
	cfg := flag.Bool("config", false, "display program configuration that was loaded")
	server := flag.Bool("server", false, "display the Octopus server information")
	report := flag.Bool("report", false, "display a report of all teams and their users/roles/environments/projects/project groups/tenants")
	teams := flag.Bool("teams", false, "display all teams")
	users := flag.Bool("users", false, "display all users")
	roles := flag.Bool("roles", false, "display all user roles")
	environments := flag.Bool("environments", false, "display all environments")
	projectGroups := flag.Bool("projectgroups", false, "display all project groups")
	projects := flag.Bool("projects", false, "display all projects")
	tenants := flag.Bool("tenants", false, "display all tenants")
	channels := flag.Bool("channels", false, "display all channels")
	feeds := flag.Bool("feeds", false, "display all feeds")
	lifecycles := flag.Bool("lifecycles", false, "display all lifecycles")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
	}

	if *info {
		fmt.Printf("%s %s %s LOG_LEVEL=%s\n", appName, appVersion, appInstance, log.GetLevel().String())
	}

	c := config.New()

	if *cfg {
		fmt.Println(c.Dump())
	}

	client := octopus.New(c.Octopus.Address, c.Octopus.APIKey, &http.Client{})

	if *server {
		octopus.PrintServerReport(client)
	}

	if *report {
		octopus.PrintTeamReport(client)
	}

	if *teams {
		octopus.PrintTeams(client)
	}

	if *users {
		octopus.PrintUsers(client)
	}

	if *roles {
		octopus.PrintUserRoles(client)
	}

	if *environments {
		octopus.PrintEnvironments(client)
	}

	if *projects {
		octopus.PrintProjects(client)
	}

	if *projectGroups {
		octopus.PrintProjectGroups(client)
	}

	if *tenants {
		octopus.PrintTenants(client)
	}

	if *channels {
		octopus.PrintChannels(client)
	}

	if *feeds {
		octopus.PrintFeeds(client)
	}

	if *lifecycles {
		octopus.PrintLifecycles(client)
	}
}
