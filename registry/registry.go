package registry

import (
	"github.com/delta/DAuth-backend-v2/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{Auth: r.NewAuthController(),}
}
