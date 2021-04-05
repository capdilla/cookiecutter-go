package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/kelseyhightower/envconfig"
)

// Config secrets for app
// envconvfig is only used on dev and json for prod
type Config struct {
	Database struct {
		User     string `envconfig:"DATABASE_USERNAME" json:"DATABASE_USERNAME"`
		Password string `envconfig:"DATABASE_PASSWORD" json:"DATABASE_PASSWORD"`
		Name     string `envconfig:"DATABASE" json:"DATABASE"`
		Host     string `envconfig:"DATABASE_HOST" json:"DATABASE_HOST"`
		Port     string `envconfig:"DB_PORT" json:"DB_PORT"`
	}
}

// GetConfig sets config depending on the enviroment
func GetConfig() (*Config, error) {
	var cfg Config
	var err error

	mode := os.Getenv("GIN_MODE")
	_, getFromSSM := os.LookupEnv("USE_SECRETS")
	if getFromSSM == false && mode != "release" {
		// Load configuration from env variables on development
		err = envconfig.Process("", &cfg)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		// Load configuration from SSM on staging/production
		cfg, err = getSSMConfig()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getSSMConfig() (Config, error) {
	secretName := os.Getenv("SECRET_NAME")
	region := os.Getenv("AWS_DEFAULT_REGION")

	// Create a Secrets Manager client
	svc := secretsmanager.New(
		session.New(),
		aws.NewConfig().WithRegion(region),
	)
	input := &secretsmanager.GetSecretValueInput{SecretId: &secretName}

	// SSM saves JSON configuration as a string of values
	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err.Error())
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	// Decode SSM string into JSON
	var config Config
	err = json.Unmarshal([]byte(secretString), &config)
	if err != nil {
		panic(err.Error())
	}

	return config, nil
}

// GetSecretByName returns a decrypted string value from AWS Systems Manager Parameter Store.
func GetSecretByName(name string) (string, error) {
	sess := session.New()
	svc := ssm.New(sess)
	output, err := svc.GetParameter(
		&ssm.GetParameterInput{
			Name:           aws.String(name),
			WithDecryption: aws.Bool(true),
		},
	)
	if err != nil {
		return "", err
	}
	return aws.StringValue(output.Parameter.Value), nil
}
