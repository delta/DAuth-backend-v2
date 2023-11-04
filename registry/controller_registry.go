package registry

import (
	"github.com/delta/DAuth-backend-v2/controller"
	impl "github.com/delta/DAuth-backend-v2/controller/impl"
)

func (r *registry) NewAuthController() controller.AuthController {
	return impl.NewAuthControllerImpl(r.NewResourceService(), r.NewEmailService())
}

func (r *registry) NewResourceController() controller.ResourceController {
	return impl.NewResourceControllerImpl(r.NewResourceService())
}
