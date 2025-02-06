package main

import (
	"context"
	"log"
	"os"
	"verify-cli-plugin/pkg/commands/upload"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "verify",
	Description: "verifies connectivity with the platform",
	Target:      types.TargetGlobal,
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.ManageCmdGroup, // set group
}

func main() {
	ctx := context.Background()
	var logger = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
	if descriptor.Version == "" {
		descriptor.Version = "dev"
	}
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		logger.Fatal(err, "")
	}
	p.Cmd.SilenceErrors = true
	p.Cmd.SilenceUsage = true
	p.AddCommands(
		upload.NewUploadCommand(ctx),
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
