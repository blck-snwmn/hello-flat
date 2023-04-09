package main

import (
	"fmt"

	"github.com/blck-snwmn/hello-flat/mygame/sample"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
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
	fmt.Printf("{x, y, z} = {%v, %v, %v}\n", pos.X(), pos.Y(), pos.Z())

	for i := 0; i < user.InventoryLength(); i++ {
		var item sample.Item
		user.Inventory(&item, i)
		fmt.Printf("[%d]:name=%s\n", i, item.Name())
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
	items := make([]flatbuffers.UOffsetT, count)
	items = append(items, createItem(builder, "sword"))
	items = append(items, createItem(builder, "shield"))
	items = append(items, createItem(builder, "armor"))

	sample.UserStartInventoryVector(builder, count)
	for _, item := range items {
		builder.PrependUOffsetT(item)
	}
	return builder.EndVector(count)
}
