package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func ExecuteRootCmd() {
	c := cobra.Command{}

	c.AddCommand(NewCloneCmd())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
