package presenter

import "fmt"

type ContactPagePresenter struct {
	global
}

func NewContactPagePresenter(isAdmin bool) ContactPagePresenter {
	return ContactPagePresenter{
		global: global{
			Categories:       globalCategories,
			IsAdmin:          isAdmin,
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, "Liên hệ"),
			CurrentPageBreadCrumbs: []string{
				"Liên hệ",
			},
		},
	}
}
