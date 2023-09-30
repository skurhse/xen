package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/pkg/apps"
	"github.com/transprogrammer/xenia/pkg/aspects"
	"github.com/transprogrammer/xenia/pkg/stacks"
)

func main() {
	app := cdktf.NewApp(nil)

	config, err := apps.LoadConfig()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	core := stacks.NewCore(app)

	jumpBeat := core.JumpBeat()

	mongoEnvs := config.MongoEnvironments()
	mongoBeats := core.MongoBeats()

	mongoDevEnv := mongoEnvs.Development()
	mongoProdEnv := mongoEnvs.Production()

	mongoDevBeat := mongoBeats.Development()
	mongoProdBeat := mongoBeats.Production()

	drums := [4]stacks.StackDrum{
		core,
		stacks.NewJump(app, jumpBeat),
		stacks.NewMongo(app, mongoDevBeat, mongoDevEnv),
		stacks.NewMongo(app, mongoProdBeat, mongoProdEnv),
	}

	for _, drum := range drums {
		aspects.AddTags(drum)
	}

	app.Synth()
}
