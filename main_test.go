package main

import (
	"testing"

	"github.com/blck-snwmn/hello-flat/mygame/sample"
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
		for _, item := range items {
			builder.PrependUOffsetT(item)
		}
		inv := builder.EndVector(count)

		name := builder.CreateString("John Doe")
		sample.UserStart(builder)
		sample.UserAddName(builder, name)
		sample.UserAddInventory(builder, inv)

		pos := sample.CreatePosition(builder, 11, 12, 13)
		sample.UserAddPos(builder, pos)
		sample.UserAddColor(builder, sample.Color(10))

		u := sample.UserEnd(builder)
		builder.Finish(u)
		buf := builder.FinishedBytes()

		_ = sample.GetRootAsUser(buf, 0)
	}
}

func BenchmarkCreateWithFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		u := createUser(builder)
		builder.Finish(u)
		buf := builder.FinishedBytes()

		_ = sample.GetRootAsUser(buf, 0)
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
			Color: 0,
			Inventory: []*sample.ItemT{
				{Name: "sword"},
				{Name: "shield"},
				{Name: "armor"},
			},
		}

		builder.Finish(u.Pack(builder))
		buf := builder.FinishedBytes()

		_ = sample.GetRootAsUser(buf, 0)
		// ut := uu.UnPack()
	}
}
