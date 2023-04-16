package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tm",
		Short: "Todo meister is a fast way to find todo comments in your code and convert them into issues.",
		Long: `Todo meister is simple tool for looking up all TODO / REFACTOR / .. comments in your code, that allows you to easily convert them into issues.

Meister is german and means something like master and is most likely be used as a public degree in the field of the vocational education.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewGenerateCmd())

	return cmd
}
