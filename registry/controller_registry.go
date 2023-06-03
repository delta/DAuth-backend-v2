package registry

import (
	"github.com/delta/DAuth-backend-v2/controller"
	impl "github.com/delta/DAuth-backend-v2/controller/impl"
)

func (r *registry) NewAuthController() controller.AuthController {
	return impl.NewAuthControllerImpl(r.NewAuthService())
}
