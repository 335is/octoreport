package octopus

import (
	"fmt"
)

// TeamMembersReport displays each team and all its members
func TeamMembersReport(teams *Teams, users *Users) {
	for _, t := range teams.Items {
		fmt.Println(t.Name)

		for _, tu := range t.MemberUserIds {
			for _, u := range users.Items {
				if tu == u.ID {
					fmt.Printf("    %s,%s,%s\n", u.DisplayName, u.Username, u.EmailAddress)
					break
				}
			}
		}
	}
}

// TeamAccessReport displays each team and all its roles (permissions)
func TeamAccessReport(teams *Teams, scopedRoles *ScopedUserRoles, environments *Environments) {
	for _, t := range teams.Items {
		fmt.Println(t.Name)

		for _, sr := range scopedRoles.Items {
			if sr.TeamID == t.ID {
				for _, eID := range sr.EnvironmentIDs {
					for _, env := range environments.Items {
						if env.ID == eID {
							fmt.Println("  ", env.Name)
							break
						}
					}
				}
			}
		}
	}
}
