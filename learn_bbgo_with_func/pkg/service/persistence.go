package service

type PersistenceService interface {
	NewStore(id string, subIDs ...string) interface{}
}
