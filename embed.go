// GENERATED FILE

package vi

import (
	"embed"
	"github.com/amitybell/piper-asset"
)

var (
	//go:embed dist.tzst dist.json MODEL_CARD.txt
	fs embed.FS

	Asset = asset.Asset{Name: "vi", FS: fs}
)
