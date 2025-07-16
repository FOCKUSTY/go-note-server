package status

import (
	"main/route"
)

func Router(data *route.Route) {
	controller := CreateContoller(CreateService(data.Database), data.Path)

	data.Router.Get(controller.Get())
	data.Router.Get(controller.GetDatabase())
}
