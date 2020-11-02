package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_a "github.com/youk-h/practice/a/mock"
)

func TestPan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_a.NewMockA(ctrl)
	mock.EXPECT().Hoge(4)

	b := Pang(mock)

	if b == 10 {
		t.Fatalf("b is %d", b)
	}
}
