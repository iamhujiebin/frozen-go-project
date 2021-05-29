package 模板Template

import "testing"

func TestNewGoogleWorker(t *testing.T) {
	NewWorker(&GoogleWorker{}).Create()
}
