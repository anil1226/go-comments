//go:build integration

package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/anil1226/go-banking/internal/comment"
)

func TestDatabase_GetComment(t *testing.T) {
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		d       *Database
		args    args
		want    comment.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetComment(tt.args.ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_PostComment(t *testing.T) {
	type args struct {
		ctx context.Context
		cmt comment.Comment
	}
	tests := []struct {
		name    string
		d       *Database
		args    args
		want    comment.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:"positive"
			d: db.NewDatabase
			args: args{
				ctx: context.Background(),
				cmt: comment.Comment{}
			}
			want: comment.Comment{}
			wantErr: false
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.PostComment(tt.args.ctx, tt.args.cmt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.PostComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.PostComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
