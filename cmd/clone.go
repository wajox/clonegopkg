//nolint //reason: TODO. disabled just due to developing process. uncomment it when module will be done
package cmd

import (
	// "os"
	// "os/signal"
	// "syscall"
	// "time"

	// "github.com/wajox/clonegopkg/internal/app"
	// "github.com/rs/zerolog/log"
	"fmt"

	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wajox/clonegopkg/internal/helpers"
)

func NewCloneCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "c",
		Aliases: []string{"clone"},
		Short:   "clone pkg from remote git repository",
		Example: "clonegopkg clone git@github.com:ildarusmanov/gobase.git github.com/ildarusmanov/newproject",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				log.Error().Msg("please provide repository url, directory and pkg name")
				return
			}

			repo := args[0]
			newPkg := args[1]
			dir := fmt.Sprintf("%s/src/%s", envy.Get("GOPATH", "~/go"), newPkg)

			if err := helpers.CloneGitRepository(repo, dir); err != nil {
				log.Error().
					Err(err).
					Str("Repository", repo).
					Str("Destination", dir).
					Msg("can not clone the repository")

				return
			}

			oldPkg, err := helpers.GetPkgNameFromGomod(dir)
			if err != nil {
				log.Error().Err(err).Msg("can not parse pkg name from go.mod")

				return
			}

			if err := helpers.StrReplaceInDirectory(dir, oldPkg, newPkg); err != nil {
				log.Error().Err(err).Msg("can not update pkg name in all files")
			}
		},
	}
}
