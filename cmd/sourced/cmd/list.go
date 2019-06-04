package cmd

import (
	"fmt"

	"github.com/src-d/sourced-ce/cmd/sourced/compose/workdir"
)

type listCmd struct {
	Command `name:"list" short-description:"List working directories" long-description:"List previously initialized working directories."`
}

func (c *listCmd) Execute(args []string) error {
	dirs, err := workdir.ListRepoDirs()
	if err != nil {
		return err
	}

	active, err := workdir.ActiveRepoDir()
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if dir == active {
			fmt.Printf("* %s\n", dir)
		} else {
			fmt.Printf("  %s\n", dir)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(&listCmd{})
}
