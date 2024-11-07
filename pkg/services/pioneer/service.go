package pioneer

import(
	"github.com/grippenet/badge-service/pkg/types"
	"fmt"
	"log/slog"
)

type PioneerService struct {
	dbService types.DBService
}

func NewPioneerService(dbService types.DBService) *PioneerService {
	return &PioneerService{
		dbService: dbService,
	}
}

func (svc *PioneerService) Check(instanceID string, studyKey string, postalCode string) (bool, error) {
	found, err := svc.dbService.FindPioneer(instanceID, studyKey, postalCode)
	if(err != nil) {
		slog.Error(fmt.Sprintf("Error loading pioneer : %s", err))
		return false, err
	}
	if(found) {
		return false, nil
	}
	err = svc.dbService.AddPioneer(instanceID, studyKey, postalCode)
	return true, err
}

