package config

type Config struct {
	Port string
	AccessKey string
	AccessSecret string
	JWTSecret string
	BuckerName string
	Region string
	AdminKey string
}

func NewConfig() *Config {
	return &Config{
		Port: ":9090",
		AccessKey: "AKIA46M3QOWCYOCU46OZ",
		AccessSecret: "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",
		JWTSecret: "mdcyu76r3jevwneqwe@WEQREFQ@373t58dggt",
		BuckerName: "karma.files",
		Region: "us-east-1",
		AdminKey: "ealfnl32rmwef342ER#@$TRSdf3rqar",
	}
}