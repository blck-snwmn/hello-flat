package main

import (
	"testing"

	psample "github.com/blck-snwmn/hello-flat/gen/buf/schema/proto/v1"
	"github.com/blck-snwmn/hello-flat/mygame/sample"
	"github.com/golang/protobuf/proto"
	flatbuffers "github.com/google/flatbuffers/go"
)

func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		const count = 3

		// Generate the content in advance, as nesting is not possible
		items := make([]flatbuffers.UOffsetT, count)
		{
			item := builder.CreateString("sword")
			sample.ItemStart(builder)
			sample.ItemAddName(builder, item)
			items = append(items, sample.ItemEnd(builder))
		}
		{
			item := builder.CreateString("shield")
			sample.ItemStart(builder)
			sample.ItemAddName(builder, item)
			items = append(items, sample.ItemEnd(builder))
		}
		{
			item := builder.CreateString("armor")
			sample.ItemStart(builder)
			sample.ItemAddName(builder, item)
			items = append(items, sample.ItemEnd(builder))
		}

		sample.UserStartInventoryVector(builder, count)
		builder.PrependUOffsetT(items[2])
		builder.PrependUOffsetT(items[1])
		builder.PrependUOffsetT(items[0])

		inv := builder.EndVector(count)

		name := builder.CreateString("John Doe")
		sample.UserStart(builder)
		sample.UserAddName(builder, name)
		sample.UserAddInventory(builder, inv)

		p := sample.CreatePosition(builder, 11, 12, 13)
		sample.UserAddPos(builder, p)
		sample.UserAddColor(builder, sample.Color(10))

		u := sample.UserEnd(builder)
		builder.Finish(u)
		buf := builder.FinishedBytes()

		user := sample.GetRootAsUser(buf, 0)

		var pos sample.Position
		user.Pos(&pos)

		for i := 0; i < user.InventoryLength(); i++ {
			var item sample.Item
			user.Inventory(&item, i)
		}
	}
}

func BenchmarkCreateWithFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		u := createUser(builder)
		builder.Finish(u)
		buf := builder.FinishedBytes()

		user := sample.GetRootAsUser(buf, 0)

		var pos sample.Position
		user.Pos(&pos)

		for i := 0; i < user.InventoryLength(); i++ {
			var item sample.Item
			user.Inventory(&item, i)
		}
	}
}

func BenchmarkCreateWithUserT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		u := sample.UserT{
			Name: "John Doe",
			Pos: &sample.PositionT{
				X: 11,
				Y: 12,
				Z: 13,
			},
			Color: sample.Color(10),
			Inventory: []*sample.ItemT{
				{Name: "sword"},
				{Name: "shield"},
				{Name: "armor"},
			},
		}

		builder.Finish(u.Pack(builder))
		buf := builder.FinishedBytes()

		uu := sample.GetRootAsUser(buf, 0)
		_ = uu.UnPack() // heavy
	}
}

func BenchmarkProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := psample.User{
			Name:  "John Doe",
			Pos:   &psample.Position{X: float32(11), Y: float32(12), Z: float32(13)},
			Color: psample.Color(10),
			Inventory: []*psample.Item{
				{Name: "sword"},
				{Name: "shield"},
				{Name: "armor"},
			},
		}
		buf, _ := proto.Marshal(&u)

		var uu psample.User
		proto.Unmarshal(buf, &uu)
	}
}
