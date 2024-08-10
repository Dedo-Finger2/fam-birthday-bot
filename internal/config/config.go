package config

import (
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"
)

var PermissionIDS []int64 = []int64{6105709676, 7496811705}

type ConfigJsonPaylod struct {
	Dates []types.Birthday `json:"dates"`
}
