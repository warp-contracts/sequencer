package types

import (
	"fmt"
)

func (next NextArweaveBlock) GetHeightString() string {
	return fmt.Sprintf("%d", next.BlockInfo.Height)
}
