package service

import (
	api "github.com/wunderkraut/radi-api"		
	"github.com/wunderkraut/radi-api/operation"
)

// Create the service API handler
func New_ServiceAPI() *ServiceAPI {
	api := ServiceAPI{}

	return &api
}


/**
 * Custom API handler for the service
 */

type ServiceAPI struct {
	api.BaseAPI
}

func (api *ServiceAPI) Validate() bool {
	return true
}

/**
 * Settings struct for a service API which will
 * likely also get used in the handlers
 */

// Service API settings struct
type ServiceAPISettings struct {
	
}
