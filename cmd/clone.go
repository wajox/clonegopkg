//nolint //reason: TODO. disabled just due to developing process. uncomment it when module will be done
package cmd

import (
	// "os"
	// "os/signal"
	// "syscall"
	// "time"

	// "github.com/wajox/clonegopkg/internal/app"
	// "github.com/rs/zerolog/log"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wajox/clonegopkg/internal/helpers"
)

func NewCloneCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "c",
		Aliases: []string{"clone"},
		Short:   "clone pkg from remote git repository",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				log.Error().Msg("please provide repository url, directory and pkg name")
				return
			}

			repo := args[0]
			dir := args[1]
			pkg := args[2]

			if err := helpers.CloneGitRepository(repo, dir); err != nil {
				log.Error().
					Err(err).
					Str("Repository", repo).
					Str("Destination", dir).
					Msg("can not clone the repository")

				return
			}

			if err := helpers.StrReplaceInDirectory(dir, repo, pkg); err != nil {
				log.Error().Err(err).Msg("can not update pkg name in all files")
			}
		},
	}
}
