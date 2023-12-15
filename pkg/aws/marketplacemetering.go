package aws

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
)

func CreateMeteringRecords(productCode string, customerIdentifier string, cpuValue, memValue, storageValue int64, cpuTimestamp, memTimestamp, storageTimestamp time.Time) (*marketplacemetering.BatchMeterUsageInput) {
	timezone, _ := time.LoadLocation("UTC")
	meteringRecords := &marketplacemetering.BatchMeterUsageInput{
		ProductCode: aws.String(productCode),
		UsageRecords: []*marketplacemetering.UsageRecord{
			{
				CustomerIdentifier: aws.String(customerIdentifier),
				Dimension:          aws.String("cpu"),
				Quantity:           aws.Int64(cpuValue) ,
				Timestamp:          aws.Time(cpuTimestamp.In(timezone)),
			},
			{
				CustomerIdentifier: aws.String(customerIdentifier),
				Dimension:          aws.String("memory"),
				Quantity:           aws.Int64(memValue),
				Timestamp:          aws.Time(memTimestamp.In(timezone)),
			},
			{
				CustomerIdentifier: aws.String(customerIdentifier),
				Dimension:          aws.String("storage"),
				Quantity:           aws.Int64(storageValue),
				Timestamp:          aws.Time(storageTimestamp.In(timezone)),
			},
		},
	}

	return meteringRecords
}

func SendMeteringRecords(m *marketplacemetering.BatchMeterUsageInput) marketplacemetering.BatchMeterUsageOutput {
	// Initial credentials loaded from SDK's default credential chain. Such as
	// the environment, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role. These credentials will be used to to make the STS Assume Role API.
	mySession := session.Must(session.NewSession())
	svc := marketplacemetering.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))
	req, err := svc.BatchMeterUsage(m)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr)
		}
	}
	if err == nil { // resp is now filled
		fmt.Println("New meteringrecord sent")
    	fmt.Println(req)
	}
	return *req
}