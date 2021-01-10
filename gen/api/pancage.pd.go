package api

import (
	"fmt"
	"math"

	"google.golang.org/protobuf/proto"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// メニュー表
type Pancake_Menu int32

const (
	Pancake_UNKNOWN           Pancake_Menu = 0
	Pancake_CLASSIC           Pancake_Menu = 1
	Pancake_BANANA_AND_WHIP   Pancake_Menu = 2
	Pancake_BACON_AND_CHEESE  Pancake_Menu = 3
	Pancake_MIX_BERRY         Pancake_Menu = 4
	Pancake_BAKED_MARSHMALLOW Pancake_Menu = 5
	Pancake_SPICY_CURRY       Pancake_Menu = 6
)
