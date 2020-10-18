package cmd

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/dhdanie/itinerary-db/db"
	"github.com/spf13/cobra"
	"os"
)

var (
	dbConn                string
	dbConnTimeout         int
	dbName                string
	itinerariesCollection string
	venuesCollection      string
)

var rootCmd = &cobra.Command{
	Use:   "itinerary-db",
	Short: "Itinerary DB Loader",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.DebugLevel)
		log.SetHandler(cli.Default)

		log.Infof("dbConn                = %s", dbConn)
		log.Infof("dbConnTimeout         = %d", dbConnTimeout)
		log.Infof("dbName                = %s", dbName)
		log.Infof("itinerariesCollection = %s", itinerariesCollection)
		log.Infof("venuesCollection      = %s", venuesCollection)

		dbCli := db.NewClient(dbConn, dbConnTimeout, dbName, itinerariesCollection, venuesCollection)
		err := dbCli.Connect()
		if err != nil {
			log.WithError(err).Error("Error connecting to database - Exiting")
			os.Exit(1)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		dbCli := db.GetClient()
		if dbCli != nil {
			dbCli.Disconnect()
			log.Debug("Disconnected from database")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbConn, "dbconn", "", "Database connection string")
	_ = rootCmd.MarkFlagRequired("dbconn")

	rootCmd.PersistentFlags().IntVar(&dbConnTimeout, "dbconn-timeout", 10, "DB connection timeout in seconds")

	rootCmd.PersistentFlags().StringVar(&dbName, "dbname", "itinerary", "DB name")
	rootCmd.PersistentFlags().StringVar(&itinerariesCollection, "itineraries-collection", "itineraries", "Itinerary collection name")
	rootCmd.PersistentFlags().StringVar(&venuesCollection, "venues-collection", "venues", "Venue collection name")

	rootCmd.AddCommand(loadCmd)
}
