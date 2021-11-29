package service

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/matheustex/go-reservation-api/util"
)

func GetItemByKey(tableName string, key AWSObject, out interface{}) (bool, error) {
	input := dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	output, err := DynamoDB().GetItem(&input)
	if err != nil {
		return false, err
	}

	if len(output.Item) == 0 {
		return false, nil
	}

	err = dynamodbattribute.UnmarshalMap(output.Item, out)
	if err != nil {
		return false, err
	}

	return true, nil
}

func QueryAll(input *dynamodb.QueryInput) ([]AWSObject, error) {
	var items = make([]AWSObject, 0)

	for {
		log.Print(input)

		req, res := DynamoDB().QueryRequest(input)

		err := req.Send()
		if err != nil {
			return items, err
		}

		log.Print(res)

		if len(res.Items) > 0 {
			items = append(items, res.Items...)
		}

		if res.LastEvaluatedKey == nil || len(res.LastEvaluatedKey) == 0 {
			break
		}

		input.ExclusiveStartKey = res.LastEvaluatedKey
	}

	return items, nil
}

func QueryItems(queryInput *dynamodb.QueryInput, offset, cap int) ([]AWSObject, error) {
	items := make([]AWSObject, 0, cap)
	resultIndex := 0

	err := DynamoDB().QueryPages(queryInput, func(page *dynamodb.QueryOutput, lastPage bool) bool {
		pageCount := len(page.Items)

		if resultIndex+pageCount > offset {
			start := util.MaxInt(0, offset-resultIndex)
			for i := start; i < pageCount; i++ {
				items = append(items, page.Items[i])
			}
		}

		resultIndex += pageCount
		return true
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}

func BatchGetItems(batchGetInput *dynamodb.BatchGetItemInput, cap int) ([]map[string][]AWSObject, error) {
	responses := make([]map[string][]AWSObject, 0, cap)

	err := DynamoDB().BatchGetItemPages(batchGetInput, func(page *dynamodb.BatchGetItemOutput, lastPage bool) bool {
		responses = append(responses, page.Responses)
		return true
	})

	if err != nil {
		return nil, err
	}

	return responses, nil
}
