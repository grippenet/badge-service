package types

import(
	"github.com/influenzanet/go-utils/pkg/configs"
)

type DBIndexRef struct {
	InstanceID string
	StudyKey   string
}

type DBConfig struct {
	configs.DBConfig
	InitialIndexes []DBIndexRef
}

type AppConfig struct {
	DBConfig DBConfig
	Http	HttpConfig
}

type HttpConfig struct {
	Port int
}

