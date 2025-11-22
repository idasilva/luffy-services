package modules

import (
	"github.com/idasilva/luffy-services/api"
	"github.com/idasilva/luffy-services/app/modules/version"
)

func Application() *api.API {
	app := api.New()

	version.Init(app)

	return app
}
