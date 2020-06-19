package sample_circleci

import (
	"context"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func setupDB(t *testing.T) {
	cmds := []string{
		`aws dynamodb --endpoint-url http://localhost:8000 create-table --cli-input-json file://./testdata/local_user.json`,
		`aws dynamodb put-item --endpoint-url http://localhost:8000 --table-name local_user --item file://./testdata/input_user.json`,
	}
	for _, cmd := range cmds {
		args := strings.Split(cmd, " ")
		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			t.Logf("setup DynamoDB %v %s", err, cmd)
		}
	}
}

func teardownDB(t *testing.T) {
	cmds := []string{
		`aws dynamodb --endpoint-url http://localhost:8000 delete-table --table local_user`,
	}
	for _, cmd := range cmds {
		args := strings.Split(cmd, " ")
		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			t.Logf("teardown DynamoDB %v %s", err, cmd)
		}
	}
}

func TestFetchUserByID(t *testing.T) {
	setupDB(t)
	t.Cleanup(func() { teardownDB(t) })

	tests := []struct {
		name    string
		userID  string
		want    *User
		wantErr bool
	}{
		{
			name:   "normal",
			userID: "001",
			want: &User{
				UserID:   "001",
				UserName: "gopher_1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchUserByID(context.TODO(), tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
