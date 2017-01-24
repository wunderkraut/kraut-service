package mongo 

/**
 * Create a handler that provides operations related to connecting
 * to a MongoDB
 */

// Constructor for a localHandlerBase
func New_ServiceHandler_Base(settings *MongoSettings) *MongoHandler_Base {
	return &MongoHandler_Base{
		settings:   settings,
		operations: &operation.Operations{},
	}
}

// MongoDB based Config Handler
type MongoHandler_Base struct {
	settings   *MongoSettings
	operations *operation.Operations
}

// Return the stored operatons
func (base *MongoHandler_Base) Operations() *operation.Operations {
	return base.operations
}
