package service 

/**
 * Some base handler structs
 */

// Constructor for a localHandlerBase
func New_ServiceHandler_Base(settings *ServiceAPISettings) *ServiceHandler_Base {
	return &ServiceHandler_Base{
		settings:   settings,
		operations: &operation.Operations{},
	}
}

// A handler for base local handlers
type ServiceHandler_Base struct {
	settings   *ServiceAPISettings
	operations *operation.Operations
}

// Return the stored operatons
func (base *ServiceHandler_Base) Operations() *operation.Operations {
	return base.operations
}
