package cli

import (
	"github.com/spf13/cobra"
)

func newSrcCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "src [@HANDLE]",
		RunE: execSrc,
		Example: `  # get active data source
   $ sq src

   # set @my1 as active data source
   $ sq src @my1`,
		Args:              cobra.MaximumNArgs(1),
		ValidArgsFunction: completeHandle(1),
		Short:             "Get or set active data source",
		Long: `Get or set active data source. If no argument provided, get the active data
source. Otherwise, set @HANDLE as the active data source.`,
	}

	return cmd
}

func execSrc(cmd *cobra.Command, args []string) error {
	rc := RunContextFrom(cmd.Context())
	cfg := rc.Config

	if len(args) == 0 {
		// Get the active data source
		src := cfg.Sources.Active()
		if src == nil {
			return nil
		}

		return rc.writers.srcw.Source(src)
	}

	src, err := cfg.Sources.SetActive(args[0])
	if err != nil {
		return err
	}

	err = rc.ConfigStore.Save(cfg)
	if err != nil {
		return err
	}

	return rc.writers.srcw.Source(src)
}
