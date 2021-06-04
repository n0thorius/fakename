package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "fakename",
	Short: "create fake identity",
	Long: `
fakename allows you to create fake identity. 
It also allows you to post that fake identity to pastebin, ghostbin`,
}

func Execute() {
	RootCmd.Execute()
}
