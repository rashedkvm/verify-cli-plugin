package upload

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

type Options struct {
	Path string
}

func NewUploadCommand(ctx context.Context) *cobra.Command {
	opts := &Options{}
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload files to S3",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid number of arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			//implementation
			return nil
		},
	}
	cmd.SetFlagErrorFunc(func(command *cobra.Command, err error) error {
		_ = command.Help()
		return err
	})

	cmd.Flags().StringVar(&opts.Path, "path", "", "path to upload files")
	return cmd
}
