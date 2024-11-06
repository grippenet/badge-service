package db

import(
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/grippenet/badge-service/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dbService *BadgeDBService) FindPioneer(instanceID string, studyKey string, postalCode string) (bool, error) {
	ctx, cancel := dbService.GetContext()
	defer cancel()

	filter := bson.M{
		"key": postalCode,
	}

	elem := types.PioneerPostalCode{}
	err := dbService.CollectionPioneer(instanceID, studyKey).FindOne(ctx, filter).Decode(&elem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (dbService *BadgeDBService) AddPioneer(instanceID string, studyKey string, postalCode string) error {
	ctx, cancel := dbService.GetContext()
	defer cancel()
	dbService.ensurePioneerIndex(instanceID, studyKey)
	elem := types.PioneerPostalCode{Key: postalCode}
	_, err := dbService.CollectionPioneer(instanceID, studyKey).InsertOne(ctx, elem)
	return  err
}

func (dbService *BadgeDBService) ensurePioneerIndex(instanceID string, studyKey string) error {
	name := fmt.Sprintf("%s:%s:pionneer")
	_, ok := dbService.indexesDone[name]
	if(ok) {
		return nil
	}
	err := dbService.CreatePioneerIndex(instanceID, studyKey)
	dbService.indexesDone[name] = struct{}{}
	return err
}

func (dbService *BadgeDBService) CreatePioneerIndex(instanceID string, studyKey string) error {
	ctx, cancel := dbService.GetContext()
	defer cancel()
	_, err := dbService.CollectionPioneer(instanceID, studyKey).Indexes().CreateMany(
		ctx, []mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "key", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
		},
	)
	return err
}

