package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/stk"
)

func main() {
	app := cdktf.NewApp(nil)

	cfg, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	tokens := stk.NewTokens(cfg)

	core := stk.NewCore(app, cfg, tokens)

	drums := [5]stk.Drum{
		core,
		stk.NewJump(app, cfg, core.JumpBeat(), tokens.Jump),
		stk.NewMongo(app, cfg, core.MongoBeats().Dev(), tokens.Mongo.Dev),
		stk.NewMongo(app, cfg, core.MongoBeats().Prod(), tokens.Mongo.Prod),
		stk.NewCluster(app, cfg, core.ClusterBeat(), tokens.Cluster),
	}

	for _, drum := range drums {
		// aspects.AddTags(drum, cfg)
		fmt.Printf(*drum.StackName())
	}

	app.Synth()
}
