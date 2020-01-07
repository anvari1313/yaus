package db

//func MongoConnect(c config.MongoDB) (*mongo.Database, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI))
//	if err != nil {
//		return nil, err
//	}
//
//	return client.Database(c.DatabaseName), nil
//}
