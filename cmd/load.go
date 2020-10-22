package cmd

import (
	"github.com/apex/log"
	"github.com/dhdanie/itinerary-db/db"
	"github.com/dhdanie/itinerary-db/load"
	"github.com/spf13/cobra"
	"os"
)

var (
	loadType string
	loadFile string
)

var typeLoaders map[string]load.Loader

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load DB data from file",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof("loadType              = %s", loadType)
		log.Infof("loadFile              = %s", loadFile)

		typeLoaders = make(map[string]load.Loader)
		typeLoaders["venue"] = load.NewVenueLoader(db.GetClient())
		typeLoaders["itinerary"] = load.NewItineraryLoader(db.GetClient())
	},
	Run: func(cmd *cobra.Command, args []string) {
		loader, ok := typeLoaders[loadType]
		if !ok {
			log.Errorf("No loader exists for type '%s' - Exiting", loadType)
			os.Exit(1)
		}
		err := loader.LoadFile(loadFile)
		if err != nil {
			log.WithError(err).Error("Error loading file")
			os.Exit(1)
		}
	},
}

func init() {
	loadCmd.Flags().StringVar(&loadType, "type", "venue", "Load data type (venue | itinerary)")
	_ = loadCmd.MarkFlagRequired("type")

	loadCmd.Flags().StringVar(&loadFile, "file", "", "File containing data to load")
	_ = loadCmd.MarkFlagRequired("file")
}
