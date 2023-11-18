package cardsearch

import (
	"context"
	"flag"
	"fmt"

	//"os"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cardsearch",
		Short: "cardsearch",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return run(ctx)
		},
	}
	return cmd
}

func run(ctx context.Context) error {
	// var str = flag.String("オプション名", "初期値", "説明")
	f := flag.String("flag1", "hoge", "flag 1")
	flag.Parse()

	fmt.Printf("Hello %s", *f)

	return nil
}
