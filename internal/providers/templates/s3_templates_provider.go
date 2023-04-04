package templates

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"io"
)

func NewS3TemplatesProvider(cfg *aws.Config, bucket string) TemplatesProvider {
	sess, err := session.NewSession(cfg)
	if err != nil {
		panic(errors.Wrap(err, "failed to create new aws session"))
	}
	instance := s3.New(sess)
	return &s3TemplatesProvider{
		instance: instance,
		bucket:   bucket,
	}
}

type s3TemplatesProvider struct {
	instance *s3.S3
	bucket   string
}

func (s *s3TemplatesProvider) GetTemplate(topic, channel, locale string) (raw []byte, errClose error) {
	path := fmt.Sprintf("/templates/%s-%s-%s", channel, topic, locale)

	file, err := s.instance.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get template file object")
	}
	defer func() {
		if err := file.Body.Close(); err != nil {
			errClose = err
		}
	}()
	raw, err = io.ReadAll(file.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the file body")
	}

	return raw, nil
}
