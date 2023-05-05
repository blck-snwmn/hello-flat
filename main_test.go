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
		builder := flatbuffers.NewBuilder(2024)
		const count = 3

		// Generate the content in advance, as nesting is not possible
		items := make([]flatbuffers.UOffsetT, 0, count)
		{
			posionStr := builder.CreateString("healing")
			typ := builder.CreateString("heal")
			sample.PosionStart(builder)
			sample.PosionAddName(builder, posionStr)
			sample.PosionAddEffectType(builder, typ)
			posion := sample.PosionEnd(builder)

			item := builder.CreateString("posion")
			sample.ItemStart(builder)
			sample.ItemAddName(builder, item)
			sample.ItemAddContentType(builder, sample.ItemContentPosion)
			sample.ItemAddContent(builder, posion)
			items = append(items, sample.ItemEnd(builder))
		}
		{
			posionStr := builder.CreateString("damage")
			typ := builder.CreateString("heal")
			sample.PosionStart(builder)
			sample.PosionAddName(builder, posionStr)
			sample.PosionAddEffectType(builder, typ)
			posion := sample.PosionEnd(builder)

			item := builder.CreateString("shield")
			sample.ItemStart(builder)
			sample.ItemAddName(builder, item)
			sample.ItemAddContentType(builder, sample.ItemContentPosion)
			sample.ItemAddContent(builder, posion)
			items = append(items, sample.ItemEnd(builder))
		}
		{
			name := builder.CreateString("apple juice")
			taste := builder.CreateString("sweet")
			sample.DrinkStart(builder)
			sample.DrinkAddName(builder, name)
			sample.DrinkAddTaste(builder, taste)
			drink := sample.DrinkEnd(builder)

			item := builder.CreateString("armor")
			sample.ItemStart(builder)
			sample.ItemAddName(builder, item)
			sample.ItemAddContentType(builder, sample.ItemContentDrink)
			sample.ItemAddContent(builder, drink)
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

			// unionTable := new(flatbuffers.Table)
			// item.Content(unionTable)

			// contentType := item.ContentType()
			// switch contentType {
			// case sample.ItemContentPosion:

			// 	posion := new(sample.Posion)
			// 	posion.Init(unionTable.Bytes, unionTable.Pos)
			// }

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

func BenchmarkCreateHybrid(b *testing.B) {
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
				{Name: "healing posion", Content: &sample.ItemContentT{
					Type: sample.ItemContentPosion,
					Value: &sample.PosionT{
						Name:       "healing",
						EffectType: "heal",
					},
				}},
				{Name: "damage posion", Content: &sample.ItemContentT{
					Type: sample.ItemContentPosion,
					Value: &sample.PosionT{
						Name:       "dagame",
						EffectType: "bob",
					},
				}},
				{Name: "juice", Content: &sample.ItemContentT{
					Type: sample.ItemContentDrink,
					Value: &sample.DrinkT{
						Name:  "juice",
						Taste: "sweet",
					},
				}},
			},
		}

		builder.Finish(u.Pack(builder))
		buf := builder.FinishedBytes()

		user := sample.GetRootAsUser(buf, 0)

		var pos sample.Position
		user.Pos(&pos)

		for i := 0; i < user.InventoryLength(); i++ {
			var item sample.Item
			user.Inventory(&item, i)

			unionTable := new(flatbuffers.Table)
			item.Content(unionTable)

			contentType := item.ContentType()
			switch contentType {
			case sample.ItemContentPosion:
				posion := new(sample.Posion)
				posion.Init(unionTable.Bytes, unionTable.Pos)
			case sample.ItemContentDrink:
				drink := new(sample.Drink)
				drink.Init(unionTable.Bytes, unionTable.Pos)
			}
		}
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
