package db

import (
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func TestNewPrice(t *testing.T) {
	ctx := getTestData()
	c, _ := GetSession(ctx)

	type args struct {
		c       *sql.DB
		fruitID int
		amount  int
		where   int
		by      int
		start   time.Time
		end     time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    Price
		wantErr bool
	}{
		{
			"Default",
			args{
				c,
				1,
				1400,
				99,
				100,
				time.Now(),
				time.Now(),
			},
			Price{
				Fruit:        1,
				Amount:       1400,
				CreatedWhere: 99,
				CreatedBy:    100,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPrice(tt.args.c, tt.args.fruitID, tt.args.amount, tt.args.where, tt.args.by, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
