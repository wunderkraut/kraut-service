package mongo

// Settings needed for connecting to a MongoDB
type MongoSettings struct {
	// The Dial Url, as seen in https://godoc.org/gopkg.in/mgo.v2#Dial
	DialUrl string
}
