package main

import "fmt"

func canAssignPrimaryContact(org_team OrganizingTeam, name string) bool {

	for _, team_m := range org_team.Members {
		if team_m == name {
			return true
		} else {
			return false
		}
	}
	return false
}

type NSPEvent struct {
	Name string

	Date string

	Primary_contact string

	organizing_team OrganizingTeam
}

type OrganizingTeam struct {
	Members []string

	Primary_contact string
}

func main() {

	org_team := OrganizingTeam{

		//Primary_contact: "Teja",

		Members: []string{"Teja", "Ravi", "Ram", "Anand"},
	}
	name := "Cahitu"
	if canAssignPrimaryContact(org_team, name) {
		org_team.Primary_contact = name
	} else {
		org_team.Primary_contact = " not allocated"
	}

	nspEvent := NSPEvent{

		Name: "Holi celebrations",

		Date: "2024-03-01",

		Primary_contact: "Sai",

		organizing_team: org_team,
	}

	// Display information about the organized event

	fmt.Printf("Event: %s\n\nDate: %s\n\nContact: %s\n\nOrganizing Team:\n", nspEvent.Name, nspEvent.Date, nspEvent.Primary_contact)
	fmt.Printf("  Primary contact: %s\n", nspEvent.organizing_team.Primary_contact)
	//for _, team := range nspEvent.organizing_team {

	//fmt.Printf("Primary contact %s\n", team.Primary_contact)

	//fmt.Printf("  Leader: %s (%s)\n", team.Leader.Name, team.Leader.Role)

	fmt.Print("  Members:")

	for _, member := range nspEvent.organizing_team.Members {

		fmt.Printf(" %s ", member)

	}

	fmt.Println()

}
