package impl

import (
	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/delta/DAuth-backend-v2/service"
)

type resourceControllerImpl struct {
	resourceService service.ResourceService
}

func NewResourceControllerImpl(resourceService service.ResourceService) controller.ResourceController {
	return &resourceControllerImpl{resourceService}
}
