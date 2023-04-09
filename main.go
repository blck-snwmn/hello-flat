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
}

func createUser(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	name := builder.CreateString("John Doe")
	sample.UserStart(builder)
	sample.UserAddName(builder, name)

	pos := sample.CreatePosition(builder, 11, 12, 13)
	sample.UserAddPos(builder, pos)
	sample.UserAddColor(builder, sample.Color(10))

	return sample.UserEnd(builder)
}
