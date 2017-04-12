package GoJsoner

import (
	"fmt"
	"testing"
)

func TestDiscardingComments(t *testing.T) {
	result, err := Discarding(`
		{//test comment1
			"name": "测试",
			/**
			test comment2
			1
			2
			3
			end
			*/
			"age":26 //test comment3
			/*****/
		}
	`)
	if err != nil {
		t.Error("discarding failed")
	}
	fmt.Println(result)
}
