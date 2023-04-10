package main

import (
	"fmt"

	"github.com/blck-snwmn/hello-flat/mygame/sample"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	fmt.Println("=====\nno pack")
	normal()
	fmt.Println("=====\npack")
	usePack()
}

func usePack() {
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

	fmt.Printf("length=%d, msg=`%X`\n", len(buf), buf)

	uu := sample.GetRootAsUser(buf, 0)
	user := uu.UnPack()

	fmt.Printf("name=`%s`\n", user.Name)
	fmt.Printf("color=%v\n", user.Color)
	fmt.Printf("position{x, y, z} = {%v, %v, %v}\n", user.Pos.X, user.Pos.Y, user.Pos.Z)

	for i, item := range user.Inventory {
		fmt.Printf("item[%d]:name=%s\n", i, item.Name)
	}
}

func normal() {
	builder := flatbuffers.NewBuilder(1024)
	u := createUser(builder)
	builder.Finish(u)
	buf := builder.FinishedBytes()

	fmt.Printf("length=%d, msg=`%X`\n", len(buf), buf)

	user := sample.GetRootAsUser(buf, 0)

	var pos sample.Position
	user.Pos(&pos)

	fmt.Printf("name=`%s`\n", user.Name())
	fmt.Printf("color=%v\n", user.Color())
	fmt.Printf("position{x, y, z} = {%v, %v, %v}\n", pos.X(), pos.Y(), pos.Z())

	for i := 0; i < user.InventoryLength(); i++ {
		var item sample.Item
		user.Inventory(&item, i)
		fmt.Printf("item[%d]:name=%s\n", i, item.Name())
	}
}

func createUser(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	inv := createInv(builder)

	name := builder.CreateString("John Doe")
	sample.UserStart(builder)
	sample.UserAddName(builder, name)
	sample.UserAddInventory(builder, inv)

	pos := sample.CreatePosition(builder, 11, 12, 13)
	sample.UserAddPos(builder, pos)
	sample.UserAddColor(builder, sample.Color(10))

	return sample.UserEnd(builder)
}

func createItem(builder *flatbuffers.Builder, name string) flatbuffers.UOffsetT {
	item := builder.CreateString(name)
	sample.ItemStart(builder)
	sample.ItemAddName(builder, item)
	return sample.ItemEnd(builder)
}

func createInv(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	const count = 3

	// Generate the content in advance, as nesting is not possible
	items := make([]flatbuffers.UOffsetT, 0, count)
	items = append(items, createItem(builder, "sword"))
	items = append(items, createItem(builder, "shield"))
	items = append(items, createItem(builder, "armor"))

	sample.UserStartInventoryVector(builder, count)
	builder.PrependUOffsetT(items[2])
	builder.PrependUOffsetT(items[1])
	builder.PrependUOffsetT(items[0])
	return builder.EndVector(count)
}
