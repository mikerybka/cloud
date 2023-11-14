package cloud

type Platform struct {
	Settings *Settings
	Auth     *Auth
	Teams    map[string]*Team
}

type Auth struct {
	Login *AuthLogin
}

type AuthLogin struct {
	Code   *LoginCode
	Github *LoginGithub
}

type LoginCode struct{}
type LoginGithub struct{}

type Settings struct{}

type Team struct {
	Apps        map[string]*App
	Deployments map[string]*Deployment
	Packages    map[string]*Package
}

type App struct {
	Constants map[string]*Constant
	Functions map[string]*Function
	Secrets   map[string]*Secret
	Types     map[string]*Type
	Variables map[string]*Variable
}

type Constant struct{}
type Function struct{}
type Secret struct{}
type Type struct{}
type Variable struct{}

type Deployment struct{}

type Package struct {
	Constants map[string]*Constant
	Functions map[string]*Function
	Secrets   map[string]*Secret
	Types     map[string]*Type
	Variables map[string]*Variable
}
