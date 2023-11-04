package registry

import (
	"github.com/delta/DAuth-backend-v2/service"
	impl "github.com/delta/DAuth-backend-v2/service/impl"
)

func (r *registry) NewResourceService() service.ResourceService {
	return impl.NewResourceServiceImpl(r.NewResourceOwnerRepository())
}

func (r *registry) NewEmailService() service.EmailService {
	return impl.NewEmailServiceImpl(r.NewEmailRepository())
}
