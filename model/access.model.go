package model

import "github.com/golang-jwt/jwt/v4"

type RealmAccessAndAccount struct {
	Roles []string `json:"roles"`
}
type ResourceAccess struct {
	Account RealmAccessAndAccount `json:"account"`
	RealmAccessAndAccount
}

// type GetRoles GetRolesFunc(access interface{}) RealmAccessAndAccount

type AccessModel struct {
	Exp                   int64                     `json:"exp"`
	Iat                   int64                     `json:"iat"`
	AuthTime              int64                     `json:"auth_time"`
	Jti                   string                    `json:"jti"`
	Iss                   string                    `json:"iss"`
	Aud                   interface{}               `json:"aud"`
	Sub                   string                    `json:"sub"`
	Typ                   string                    `json:"typ"`
	Azp                   string                    `json:"azp"`
	SessionState          string                    `json:"session_state"`
	AllowedOrigins        []string                  `json:"allowed-origins"`
	RealmAccessAndAccount RealmAccessAndAccount     `json:"realm_access"`
	ResourceAccess        map[string]ResourceAccess `json:"resource_access"`
	Scope                 string                    `json:"scope"`
	EmailVerified         bool                      `json:"email_verified"`
	Phone                 int                       `json:"phone"`
	Name                  string                    `json:"name"`
	Company               string                    `json:"company"`
	PreferredUsername     string                    `json:"preferred_username"`
	GivenName             string                    `json:"given_name"`
	FamilyName            string                    `json:"family_name"`
	Email                 string                    `json:"email"`
	Group                 []string                  `json:"group"`
	Realm                 string                    `json:"realm"`
	jwt.RegisteredClaims
}

// jwt.Claims интерфэйсийг бүрэн хэрэгжүүлэх
func (a AccessModel) Valid() error {
	return a.RegisteredClaims.Valid()
}
