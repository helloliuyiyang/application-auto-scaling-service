package ascontroller

import (
	"testing"
)

func TestAsController_AddServers(t *testing.T) {
	InitAsController()
	//if _, err := GetAsController().AddServers("df107ffc-fecb-4733-afa5-8d6deec995b9", 2); err != nil {
	//	t.Error(err)
	//}

	if err := GetAsController().DelServer("df107ffc-fecb-4733-afa5-8d6deec995b9",
		"54f28d0f-a7ac-45cb-a5da-4d96d9f0946c"); err != nil {
		t.Error(err)
	}
}
