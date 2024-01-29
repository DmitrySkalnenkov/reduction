package repo

import (
	"fmt"
	"os"
	"testing"
)

func TestURLStorageInit(t *testing.T) {

	tests := []struct {
		name        string
		filePath    string
		wantTypeStr string
	}{ //Test table
		{
			name:        "Positive test 1. Memory repository.",
			filePath:    "",
			wantTypeStr: "*memrepo.MemRepo",
		},
		{
			name:        "Positive test 1. Memory repository.",
			filePath:    "/tmp/temp_for_URLStorageInit.txt",
			wantTypeStr: "*filerepo.FileRepo",
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			us := URLStorageInit(tt.filePath)
			curTypeStr := fmt.Sprintf("%T", us)
			fmt.Printf("Current type is '%s'\n", curTypeStr)
			if curTypeStr != tt.wantTypeStr {
				t.Errorf("TEST_ERROR: Type of storage ('%s') doesn't match expected ('%s').\n", curTypeStr, tt.wantTypeStr)
			}
			if tt.filePath != "" {
				os.Remove(tt.filePath)
			}
		})
	}
}
