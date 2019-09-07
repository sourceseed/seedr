package seedr

import (
	"errors"
	"fmt"
	"os"

	"github.com/sourceseed/seedr/internal/pkg/config"
	"github.com/sourceseed/seedr/internal/pkg/generate"
	"github.com/sourceseed/seedr/internal/pkg/seed"
	"github.com/sourceseed/seedr/internal/pkg/ui"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// cutCmd represents the cut command
var generateCmd = &cobra.Command{
	Use:   "generate <appname> <template>",
	Short: "Generate a skeleton.",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner()

		seedStr, _ := cmd.Flags().GetString("seed")
		targetStr, _ := cmd.Flags().GetString("target")

		seedConfig := configSeed(seedStr)

		params := map[string]string{}
		seedFile, err := config.ParseSeedfile(seedConfig.GetSeedfilePath())
		if err == nil {
			for _, info := range seedFile.Parameters {
				params[info.Variable] = askParam(info)
			}
		}

		generate.Generate(generate.GenerateConfig{
			Seed:       seedConfig,
			TargetDir:  targetStr,
			Parameters: params,
		})

		os.Exit(0)
	},
}

func init() {
	generateCmd.Flags().String("seed", "", "Seed source")
	generateCmd.Flags().String("target", ".", "Seed source")
	rootCmd.AddCommand(generateCmd)
}

func configSeed(str string) *seed.Seed {
	if v, err := seed.NewSeed(str); err == nil {
		return v
	}

	var v *seed.Seed
	var vs string

	survey.Ask([]*survey.Question{{
		Name:   "seed",
		Prompt: &survey.Input{Message: "What template do you want to use?"},
		Validate: func(val interface{}) error {
			str, _ := val.(string)
			if str == "" {
				return errors.New("Needs value")
			}

			sd, err := seed.NewSeed(str)
			if err != nil {
				return err
			}
			v = sd
			return nil
		},
	}}, &vs)

	return v
}

func askParam(info config.ParamOptions) string {
	var v string

	survey.Ask([]*survey.Question{{
		Name: info.Variable,
		Prompt: &survey.Input{
			Message: fmt.Sprintf("%s (%s)", info.Description, info.Variable),
		},
		Validate: func(ans interface{}) error {
			if !info.Optional && ans.(string) == "" {
				return errors.New("Value required")
			}
			return nil
		},
	}}, &v)

	return v
}
