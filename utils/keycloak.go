package utils

import "github.com/buyanbadrakh/keycloak-group/model"

func HasRole(user model.AccessModel, role string) bool {
	access, ok := user.ResourceAccess[user.Azp]
	if !ok {
		return false
	}

	for _, r := range access.Roles {
		if r == role {
			return true
		}
	}
	return false
}
