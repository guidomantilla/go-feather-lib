package config

const (
	Application = "golang-engineer-udolyj"
	Version     = "v0.1.0"
)

type Config struct {
	Host                 *string `env:"HOST,default=localhost"`
	HttpPort             *string `env:"HTTP_PORT,default=8080"`
	GrpcPort             *string `env:"GRPC_PORT,default=50051"`
	TokenSignatureKey    *string `env:"TOKEN_SIGNATURE_KEY,default=SecretYouShouldHide"`
	TokenVerificationKey *string `env:"TOKEN_VERIFICATION_KEY,default=SecretYouShouldHide"`
	TokenTimeout         *string `env:"TOKEN_TIMEOUT,default=24h"`
	DatasourceDriver     *string `env:"DATASOURCE_DRIVER,required"`
	DatasourceUsername   *string `env:"DATASOURCE_USERNAME,required"`
	DatasourcePassword   *string `env:"DATASOURCE_PASSWORD,required"`
	DatasourceServer     *string `env:"DATASOURCE_SERVER,required"`
	DatasourceService    *string `env:"DATASOURCE_SERVICE,required"`
	DatasourceUrl        *string `env:"DATASOURCE_URL,required"`
}
