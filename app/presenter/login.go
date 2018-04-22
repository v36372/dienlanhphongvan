package presenter

import "fmt"

type LoginPagePresenter struct {
	global
}

func NewLoginPagePresenter() LoginPagePresenter {
	return LoginPagePresenter{
		global: global{
			Categories: globalCategories,
			IsAdmin:    false,
			CurrentPageBreadCrumbs: []string{
				"Đăng nhập",
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, "Đăng nhập"),
		},
	}
}
