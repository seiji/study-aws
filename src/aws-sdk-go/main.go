//
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, _ := session.NewSession()
	svc := s3.New(sess)
	res, err := svc.ListBuckets(nil)
	if err != nil {
		panic(err)
	}

	for _, b := range res.Buckets {
		fmt.Printf("- %-50s\t%s \n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

