package stuutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Sex  int    `json:"sex"`
}

func TestStructToUrlEncode(t *testing.T) {
	data := Data{
		Name: "Xiao",
		Age:  "18",
		Sex:  1,
	}
	assert.Equal(t, StructToUrlEncode(data), "age=18&name=Xiao&sex=1")
}

func TestStructToSliceAndMap(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name          string
		args          args
		wantSliceKey  []string
		wantSliceVale []string
		wantMap       map[string]string
	}{
		{
			name: "1",
			args: args{
				Data{
					Name: "Xiao",
					Age:  "18",
					Sex:  1,
				},
			},
			wantSliceKey:  []string{"name", "age", "sex"},
			wantSliceVale: []string{"Xiao", "18", "1"},
			wantMap:       map[string]string{"age": "18", "name": "Xiao", "sex": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSliceKey, gotSliceVale, gotMap := StructToSliceAndMap(tt.args.data)
			assert.Equalf(t, tt.wantSliceKey, gotSliceKey, "StructToSliceAndMap(%v)", tt.args.data)
			assert.Equalf(t, tt.wantSliceVale, gotSliceVale, "StructToSliceAndMap(%v)", tt.args.data)
			assert.Equalf(t, tt.wantMap, gotMap, "StructToSliceAndMap(%v)", tt.args.data)
		})
	}
}
