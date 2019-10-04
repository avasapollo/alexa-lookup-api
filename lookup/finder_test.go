// +build integration

package lookup

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinder_Do(t *testing.T) {
	testKey := os.Getenv("GOOGLE_API_KEY")
	client, err := baseClient(&Config{GoogleApiKey: testKey})
	require.NoError(t, err)

	type fields struct {
		config *Config
		opts   *Options
	}
	type args struct {
		ctx context.Context
		req *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    func(*NearbyResult)
		wantErr bool
	}{
		{
			name: "search",
			fields: fields{
				config: &Config{GoogleApiKey: testKey},
				opts: &Options{
					Logger: baseLogger(),
					Client: client,
				},
			},
			args: args{
				ctx: context.Background(),
				req: &Request{
					Location: &Location{
						Lat: 52.127820,
						Lng: -0.476020,
					},
					Radius:  50000,
					Keyword: "cinema",
					OpenNow: true,
				},
			},
			want: func(result *NearbyResult) {
				assert.True(t, len(result.List) > 0)
				for _, p := range result.List {
					t.Log(p)
				}

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Finder{
				config: tt.fields.config,
				opts:   tt.fields.opts,
			}
			got, err := s.Do(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				tt.want(got)
			}
		})
	}
}
