package types

type DBService interface {
	FindPioneer(instanceID string, studyKey string, postalCode string) (bool, error)
	AddPioneer(instanceID string, studyKey string, postalCode string) error
}