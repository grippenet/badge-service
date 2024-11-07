package memory

import(
	"fmt"
	"log/slog"
)

type MemoryDBService struct {
	logger  *slog.Logger
	pioneer map[string]struct{}
}

func NewMemoryDBService() *MemoryDBService {
	return &MemoryDBService{
		logger: slog.Default()
		pioneer: make(map[string]struct{}, 0),
	}
}

func (svc *MemoryDBService) getPrefix(instanceID string, studyKey string, key string) string {
	return fmt.Sprintf("%s/%s|%s", instanceID, studyKey, key)
}

func (svc *MemoryDBService) FindPioneer(instanceID string, studyKey string, postalCode string) (bool, error) {
	shard := svc.getPrefix(instanceID, studyKey, postalCode)
	_, ok := svc.pioneer[shard]
	return ok, nil
}

func (svc *MemoryDBService) AddPioneer(instanceID string, studyKey string, postalCode string) error {
	shard := svc.getPrefix(instanceID, studyKey, postalCode)
	svc.pioneer[shard] = struct{}{}
	return nil
}
