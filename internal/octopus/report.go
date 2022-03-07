package octopus

import (
	"fmt"

	"github.com/335is/log"
)

// PrintServerReport displays the Octopus server information
func PrintServerReport(octo *Client) {
	ver, err := octo.Version()
	if err != nil {
		log.Infof(err.Error())
	}
	fmt.Println(ver)
}

// TeamReport displays each team, its members, roles, projects, environments, tenants, tenant tags, project groups
//	All of the parameters are pointers and can be nil, so check before dereferencing!
func PrintTeamReport(octo *Client) {
	teams, err := octo.GetAllTeams()
	if err != nil {
		log.Infof(err.Error())
	}

	users, err := octo.GetAllUsers()
	if err != nil {
		log.Infof(err.Error())
	}

	userRoles, err := octo.GetAllUserRoles()
	if err != nil {
		log.Infof(err.Error())
	}

	environments, err := octo.GetAllEnvironments()
	if err != nil {
		log.Infof(err.Error())
	}

	projectGroups, err := octo.GetAllProjectGroups()
	if err != nil {
		log.Infof(err.Error())
	}

	projects, err := octo.GetAllProjects()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Teams Report -----")
	defer fmt.Println("----- End -----")

	// bail out if no teams or users
	if teams == nil {
		fmt.Println("No teams found")
		return
	}
	if users == nil {
		fmt.Println("No users found")
		return
	}

	// warn but continue to report if no roles, environments, project groups
	if userRoles == nil {
		fmt.Println("No user roles found")
	}
	if environments == nil {
		fmt.Println("No environments found")
	}
	if projectGroups == nil {
		fmt.Println("No project groups found")
	}
	if projects == nil {
		fmt.Println("No projects found")
	}

	// display each team's properties
	for _, t := range *teams {
		// users
		fmt.Println(t.Name)
		fmt.Println("  Users:")
		for _, tu := range t.MemberUserIds {
			for _, u := range *users {
				if tu == u.ID {
					fmt.Printf("    %s,%s,%s\n", u.DisplayName, u.Username, u.EmailAddress)
					break
				}
			}
		}

		// roles
		if userRoles != nil {
			fmt.Println("  Roles:")
			for _, ur := range t.UserRoleIds {
				for _, r := range *userRoles {
					if ur == r.ID {
						fmt.Printf("    %s\n", r.Name)
						break
					}
				}
			}
		}

		// environments
		if environments != nil {
			fmt.Println("  Environments:")
			for _, eid := range t.EnvironmentIds {
				for _, e := range *environments {
					if eid == e.ID {
						fmt.Printf("    %s\n", e.Name)
						break
					}
				}
			}
		}

		// project groups
		if projectGroups != nil {
			fmt.Println("  Project Groups:")
			for _, gid := range t.ProjectGroupIds {
				for _, g := range *projectGroups {
					if gid == g.ID {
						fmt.Printf("    %s\n", g.Name)
						break
					}
				}
			}
		}

		// projects
		if projects != nil {
			fmt.Println("  Projects:")
			for _, pid := range t.ProjectIds {
				for _, p := range *projects {
					if pid == p.ID {
						fmt.Printf("    %s\n", p.Name)
						break
					}
				}
			}
		}
	}
}

// PrintTeams displays just the team names
func PrintTeams(octo *Client) {
	teams, err := octo.GetAllTeams()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Teams List -----")
	defer fmt.Println("----- End -----")

	if teams == nil {
		fmt.Println("No teams found")
		return
	}

	for _, t := range *teams {
		fmt.Printf("%s (%s)\n", t.Name, t.ID)
	}
}

// PrintUsers displays the teams and their users
func PrintUsers(octo *Client) {
	users, err := octo.GetAllUsers()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Users List -----")
	defer fmt.Println("----- End -----")

	if users == nil {
		fmt.Println("No users found")
		return
	}

	for _, t := range *users {
		fmt.Printf("%s (%s)\n", t.Username, t.ID)
	}
}

// PrintUserRoles displays the user roles
func PrintUserRoles(octo *Client) {
	userRoles, err := octo.GetAllUserRoles()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Roles List -----")
	defer fmt.Println("----- End -----")

	if userRoles == nil {
		fmt.Println("No user roles found")
		return
	}

	for _, r := range *userRoles {
		fmt.Printf("%s (%s)\n", r.Name, r.ID)
	}
}

// PrintEnvironments displays the deployment environments
func PrintEnvironments(octo *Client) {
	environments, err := octo.GetAllEnvironments()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Environments List -----")
	defer fmt.Println("----- End -----")

	if environments == nil {
		fmt.Println("No environments found")
		return
	}

	for _, e := range *environments {
		fmt.Printf("%s (%s)\n", e.Name, e.ID)
	}
}

// PrintProjectGroups displays the project groups
func PrintProjectGroups(octo *Client) {
	projectGroups, err := octo.GetAllProjectGroups()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Project Groups List -----")
	defer fmt.Println("----- End -----")

	if projectGroups == nil {
		fmt.Println("No project groups found")
		return
	}

	for _, g := range *projectGroups {
		fmt.Printf("%s (%s)\n", g.Name, g.ID)
	}
}

// PrintProjects displays the projects
func PrintProjects(octo *Client) {
	projects, err := octo.GetAllProjects()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Projects List -----")
	defer fmt.Println("----- End -----")

	if projects == nil {
		fmt.Println("No projects found")
		return
	}

	for _, p := range *projects {
		fmt.Printf("%s (%s)\n", p.Name, p.ID)
	}
}

// PrintChannels displays the channels
func PrintChannels(octo *Client) {
	channels, err := octo.GetAllChannels()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Channels List -----")
	defer fmt.Println("----- End -----")

	if channels == nil {
		fmt.Println("No Channels found")
		return
	}

	for _, p := range *channels {
		fmt.Printf("%s (%s)\n", p.Name, p.ID)
	}
}

// PrintFeeds displays the feeds
func PrintFeeds(octo *Client) {
	feeds, err := octo.GetAllFeeds()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Feeds List -----")
	defer fmt.Println("----- End -----")

	if feeds == nil {
		fmt.Println("No Feeds found")
		return
	}

	for _, p := range *feeds {
		fmt.Printf("%s (%s)\n", p.Name, p.ID)
	}
}

// PrintLifecycles displays the projects
func PrintLifecycles(octo *Client) {
	lifecycles, err := octo.GetAllLifecycles()
	if err != nil {
		log.Infof(err.Error())
	}

	fmt.Println("----- Octopus Lifecycles List -----")
	defer fmt.Println("----- End -----")

	if lifecycles == nil {
		fmt.Println("No Lifecycles found")
		return
	}

	for _, p := range *lifecycles {
		fmt.Printf("%s (%s)\n", p.Name, p.ID)
	}
}
