package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nabeken/psadm"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type ExportCommand struct {
	KeyPrefix string `long:"key-prefix" description:"Specify a key prefix to be exported"`
}

func (cmd *ExportCommand) Execute(args []string) error {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	client := psadm.NewClient(cfg)

	params, err := client.GetParametersByPath(ctx, cmd.KeyPrefix)
	if err != nil {
		return err
	}

	out, err := yaml.Marshal(params)
	if err != nil {
		return errors.Wrap(err, "failed to marshal into YAML")
	}

	fmt.Print(string(out))

	return nil
}

func init() {
	parser.AddCommand(
		"export",
		"Export parameters",
		"The export command exports parameters from Parameter Store.",
		&ExportCommand{},
	)
}
