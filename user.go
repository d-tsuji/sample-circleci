package sample_circleci

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	UserID   string `dynamodbav:"user_id"`
	UserName string `dynamodbav:"user_name"`
}

func FetchUserByID(ctx context.Context, userID string) (*User, error) {
	q := &dynamodb.GetItemInput{
		TableName: aws.String(userTable),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
		},
		ConsistentRead: aws.Bool(true),
	}
	out, err := db.GetItemWithContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("dynamodb fetch user: %w", err)
	}

	var user *User
	if err := dynamodbattribute.UnmarshalMap(out.Item, &user); err != nil {
		return nil, fmt.Errorf("dynamodb unmarshal: %w", err)
	}

	return user, nil
}
