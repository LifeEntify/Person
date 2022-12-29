package person_db

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	person_config "github.com/lifeentify/person/config"
)

const (
	COLLECTION = "person"
)

type MongoDB struct {
	uri      string
	database string
}

func NewMongoDB(config *person_config.Config) *MongoDB {
	return &MongoDB{
		uri: config.MongoUrl,
		database: config.DatabaseName,
	}
}
func MongoConnection(db *MongoDB) (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
func (db *MongoDB) FindPersonByPhone(ctx context.Context, phone string) ([]byte, error) {
	client, collection := MongoConnection(db)
	defer MongoDisconnect(client)
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "profile.phone_number", Value: phone}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the phone %s\n", phone)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) FindPersonByID(ctx context.Context, accountId string) ([]byte, error) {
	client, collection := MongoConnection(db)
	defer MongoDisconnect(client)
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "account_id", Value: accountId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the accountId %s\n", accountId)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) FindPersonProfile(ctx context.Context, _id string) ([]byte, error) {
	client, coll := MongoConnection(db)
	defer MongoDisconnect(client)
	var result bson.M
	opts := options.FindOne().SetProjection(bson.D{{Key: "profile", Value: 1}, {Key: "account_id", Value: 1}})
	err := coll.FindOne(ctx, bson.D{{Key: "_id", Value: _id}}, opts).Decode(result)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) FindPersonCrendential(ctx context.Context, _id string) ([]byte, error) {
	client, coll := MongoConnection(db)
	defer MongoDisconnect(client)
	var result bson.M
	opts := options.FindOne().SetProjection(bson.D{{Key: "credential", Value: 1}, {Key: "account_id", Value: 1}})
	err := coll.FindOne(ctx, bson.D{{Key: "_id", Value: _id}}, opts).Decode(result)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) Save(ctx context.Context, person any) (*mongo.InsertOneResult, error) {
	client, coll := MongoConnection(db)
	defer MongoDisconnect(client)
	value, err := coll.InsertOne(ctx, person)
	if err != nil {
		return nil, err
	}
	return value, nil
}
func (db *MongoDB) Update(ctx context.Context, _id string, person any) (*mongo.UpdateResult, error) {
	client, coll := MongoConnection(db)
	defer MongoDisconnect(client)
	bPerson, err := bson.Marshal(person)
	if err != nil {
		return nil, err
	}
	var update bson.M
	err = bson.Unmarshal(bPerson, &update)
	if err != nil {
		return nil, err
	}
	value, err := coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: _id}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		panic(err)
	}
	return value, nil
}
