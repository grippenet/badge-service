package config

import(
	"os"
	"fmt"
	"strings"
	"github.com/influenzanet/go-utils/pkg/configs"
	"github.com/grippenet/badge-service/pkg/types"
)

func LoadConfig() (types.AppConfig, error) {
	dbConfig := configs.GetMongoDBConfig("BADGE_")
	initIndexes, err := parseIndexes(os.Getenv("BADGE_DB_INITIAL_INDEXES"))
	return types.AppConfig{
		DBConfig: types.DBConfig{
			DBConfig: dbConfig,
			InitialIndexes:initIndexes,
		},
	}, err
}

func parseIndexes(spec string) ([]types.DBIndexRef, error) {
	indexes := make([]types.DBIndexRef, 0, 0)
	if(spec == "") {
		return indexes, nil
	}
	specs := strings.Split(spec, ",")
	for i, spec := range specs {
		index := strings.Split(spec, ":")
		if(len(index) != 2) {
			return indexes, fmt.Errorf("Initial index %d, must have instance:studyKey form", i)
		}
		indexes = append(indexes, types.DBIndexRef{InstanceID: index[0], StudyKey: index[1]})
	}
	return indexes, nil
}