package healthcheck

import "github.com/abmpio/abmp/app"

func init() {
	app.RegisterStartupAction(healthcheckStartup)
}
