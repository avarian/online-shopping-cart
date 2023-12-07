package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getPostgresDSN(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "postgres config",
			request:  "postgres",
			expected: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", "", "", "", "", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getPostgresDSN(tt.request)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_getMysqlDSN(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "mysql config",
			request:  "mysql",
			expected: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=Local&parseTime=true", "", "", "", "", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMysqlDSN(tt.request)
			assert.Equal(t, tt.expected, result)
		})
	}
}
