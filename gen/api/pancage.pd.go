package api

import (
	"fmt"
	"math"

	"github.com/golang/protobuf/ptypes/timestamp"
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

var Pancake_Menu_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLASSIC",
	2: "BANANA_AND_WHIP",
	3: "BACON_AND_CHEESE",
	4: "MIX_BERRY",
	5: "BAKED_MARSHMALLOW",
	6: "SPICY_CURRY",
}

var Pancake_Menu_value = map[string]int32{
	"UNKNOWN":           0,
	"CLASSIC":           1,
	"BANANA_AND_WHIP":   2,
	"BACON_AND_CHEESE":  3,
	"MIX_BERRY":         4,
	"BAKED_MARSHMALLOW": 5,
	"SPICY_CURRY":       6,
}

type Pancake struct {
	// シェフの名前
	ChefName string `protobuf:"bytes,1,opt,name=chef_name,json=chefName,proto3" json:"chef_name,omitempty"`
	// メニュー
	Menu Pancake_Menu `protobuf:"varint,2,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
	// 焼き具合を表すスコアです(0-0.9)
	TechnicalScore float32 `protobuf:"fixed32,3,opt,name=technical_score,json=technicalScore,proto3" json:"technical_score,omitempty"`
	// 焼いた日時
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pancake) Reset()         { *m = Pancake{} }
func (m *Pancake) String() string { return proto.CompactTextString(m) }
func (*Pancake) ProtoMessage()    {}
func (*Pancake) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{0}

}
