package registry

import "github.com/delta/DAuth-backend-v2/service"
import impl "github.com/delta/DAuth-backend-v2/service/impl"

func (r *registry) NewAuthService() service.AuthService{
	return impl.NewAuthServiceImpl(r.NewResourceOwnerRepository())
}
