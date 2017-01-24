package mongo

import (
	"gopkg.in/mgo.v2"

	"github.com/wunderkraut/radi-api/operation/config"
)

const(
	MONGO_CONFIG_TABLE
)

// A Config handler that uses MongoDB as a backend
type MongoConfigHandler struct {
	MongoHandler_Base
}

// [Handler.]Init tells the LocalHandler_Orchestrate to prepare it's operations
func (handler *MongoConfigHandler) Init() operation.Result {
	result := operation.BaseResult{}
	result.Set(true, nil)

	ops := operation.Operations{}

	// build a ConfigConnector for use with the Config operations.
	connector := New_MongoConfigConnector(handler.settings)

	// Build this base operation to be shared across all of our config operations
	baseConnectorOperation := config.New_BaseConfigConnectorOperation(config.ConfigConnector(&connector))

	// Now we can add config operations that use that Base class
	ops.Add(operation.Operation(&config.ConfigSimpleConnectorReadersOperation{BaseConfigConnectorOperation: *baseConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigSimpleConnectorWritersOperation{BaseConfigConnectorOperation: *baseConnectorOperation}))
	ops.Add(operation.Operation(&config.ConfigSimpleConnectorListOperation{BaseConfigConnectorOperation: *baseConnectorOperation}))

	handler.operations = &ops

	return operation.Result(&result)
}

func (handler *MongoConfigHandler) ConfigWrapper() config.ConfigWrapper {
	return config.ConfigWrapper(config.New_SimpleConfigWrapper(handler.Operations()))
}


/**
 * Here we provide a ConfigConnector for MongoDB so that we don't have to 
 * write all new operations for the MongoDB backend, we just write the 
 * connector, and reused the existing ConfigConnect operations.
 */

// Constructor for MongoConfigConnector
func New_MongoConfigConnector(settings MongoSettings) *MongoConfigConnector {
	return &MongoConfigConnector{
		settings: settings
	}
} 

// Provide a Connector to MongoDB to leverage the existing connector operations
type MongoConfigConnector struct {
	settings MongoSettings
	session mgo.Session
}

func (connector *MongoConfigConnector) Dial() error {
	var err error
	connector.session, err = mgo.Dial(connector.settings.DialUrl)
	return err
}

func (connector *MongoConfigConnector) Readers(key string) config.ScopedReaders {
	connector.safe()

	return config.ScopedReaders{}
}

func (connector *MongoConfigConnector) Writers(key string) config.ScopedWriters {
	connector.safe()

	return config.ScopedWriters{}	
}

func (connector *MongoConfigConnector) List() []string {
	connector.safe()
}

// Ensure that we have a session safely
func (connector *MongoConfigConnector) safe() bool {
	if connector.session = nil {
		err := connector.Dial()
		if err != nil {
			log.WithError(err).Error("Could not connect to MongoDB")
			return false
		}
	}
	return true
}

