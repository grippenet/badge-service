package db

import (
	"context"
	"time"
	//"github.com/influenzanet/go-utils/pkg/configs"
	"github.com/grippenet/badge-service/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BadgeDBService struct {
	DBClient     *mongo.Client
	timeout      int
	DBNamePrefix string
	indexesDone map[string]struct{}
}

func NewBadgeDBService(conf types.DBConfig) (*BadgeDBService, error) {
	var err error
	dbClient, err := mongo.NewClient(
		options.Client().ApplyURI(conf.URI),
		options.Client().SetMaxConnIdleTime(time.Duration(conf.IdleConnTimeout)*time.Second),
		options.Client().SetMaxPoolSize(conf.MaxPoolSize),
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Timeout)*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	ctx, conCancel := context.WithTimeout(context.Background(), time.Duration(conf.Timeout)*time.Second)
	err = dbClient.Ping(ctx, nil)
	defer conCancel()
	if err != nil {
		return nil, err
	}

    indexesDone := make(map[string]struct{}, 1)

	svc := &BadgeDBService{
		DBClient:     dbClient,
		timeout:      conf.Timeout,
		DBNamePrefix: conf.DBNamePrefix,
		indexesDone: indexesDone,
	}

	for _, index := range conf.InitialIndexes {
		err = svc.ensurePioneerIndex(index.InstanceID, index.StudyKey)		
		if(err != nil) {
			return nil, err
		}
	}
 
	return svc, nil
}


func (dbService *BadgeDBService) getDbName(instanceID string) string {
	return dbService.DBNamePrefix + instanceID + "_badgeDB"
}

func (dbService *BadgeDBService) CollectionPioneer(instanceID string, studyKey string) *mongo.Collection {
	dbName := dbService.getDbName(instanceID)
	return dbService.DBClient.Database(dbName).Collection(studyKey + "_pioneer")
}

// DB utils
func (dbService *BadgeDBService) GetContext() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(dbService.timeout)*time.Second)
}
