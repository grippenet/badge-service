package services

import(
	"github.com/grippenet/badge-service/pkg/types"
	"github.com/grippenet/badge-service/pkg/services/pioneer"
)


func InitServices(dbService types.DBService) types.BadgeServices {

	pioneerSvc := pioneer.NewPioneerService(dbService)

	return types.BadgeServices{
		Pioneer: pioneerSvc,
	}
}