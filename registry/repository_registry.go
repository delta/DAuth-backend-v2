package registry

import "github.com/delta/DAuth-backend-v2/repository"
import impl "github.com/delta/DAuth-backend-v2/repository/impl"


func (r *registry) NewResourceOwnerRepository() repository.ResourceOwnerRepository{
	return impl.NewResourceOwnerRepositoryImpl(r.db)
}
