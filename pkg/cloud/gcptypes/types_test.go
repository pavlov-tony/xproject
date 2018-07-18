package gcptypes

import (
	"testing"
	"time"
)

func Test_Objects_after(t *testing.T) {
	// empty objects list
	var objs0 Objects
	objs0 = objs0.after(time.Now())
	if len(objs0) != 0 {
		t.Errorf("Objs list len Exp: %d, Got: %d", 0, len(objs0))
	}

	// 1 element in list
	objs1 := Objects{Object{
		Created: time.Date(2018, 7, 1, 0, 0, 0, 0, time.UTC)}}

	res1 := objs1.after(
		time.Date(2018, 8, 1, 0, 0, 0, 0, time.UTC))
	if len(res1) != 0 {
		t.Errorf("Objs list len Exp: %d, Got: %d", 0, len(res1))
	}

	res1 = objs1.after(
		time.Date(2018, 6, 1, 0, 0, 0, 0, time.UTC))
	if len(res1) != 1 {
		t.Errorf("Objs list len Exp: %d, Got: %d", 1, len(res1))
	}
}
