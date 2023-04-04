package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type s3Config struct {
	AccessKey string `fig:"access_key,required"`
	SecretKey string `fig:"secret_key,required"`
	Endpoint  string `fig:"endpoint,required"`
	Region    string `fig:"region"`

	Bucket string `fig:"bucket,required"`
}

type S3 interface {
	AwsConfig() *aws.Config
	Bucket() string
}

func NewS3(getter kv.Getter) S3 {
	return &s3{
		getter: getter,
	}
}

type s3 struct {
	getter        kv.Getter
	onceAwsConfig comfig.Once
	onceBucket    comfig.Once
}

func (s *s3) AwsConfig() *aws.Config {
	return s.onceAwsConfig.Do(func() interface{} {
		raw := kv.MustGetStringMap(s.getter, "s3")
		config := new(s3Config)
		err := figure.Out(config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out s3 config"))
		}
		return &aws.Config{
			Region:           aws.String(config.Region),
			Credentials:      credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
			Endpoint:         aws.String(config.Endpoint),
			S3ForcePathStyle: aws.Bool(true),
		}
	}).(*aws.Config)
}

func (s *s3) Bucket() string {
	return s.onceBucket.Do(func() interface{} {
		raw := kv.MustGetStringMap(s.getter, "s3")
		config := new(s3Config)
		err := figure.Out(config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out s3 config"))
		}
		return config.Bucket
	}).(string)
}
