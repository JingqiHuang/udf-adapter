package util_3gpp

import (
	"fmt"

	"github.com/omec-project/openapi/models"
)

func SNssaiToString(snssai *models.Snssai) (str string) {
	if snssai.Sd == "" {
		return fmt.Sprintf("%d-%s", snssai.Sst, snssai.Sd)
	}
	return fmt.Sprintf("%d", snssai.Sst)
}
