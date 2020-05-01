package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path"
	"tatry/app"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "tatry",
		Short: "Auto update your background",
		Long: `Tatry is application that updates your background with images from given URLs, 
by default images are taken from Livecams from Polish mountains called Tatry`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	cams := []string{
		"http://pogoda.topr.pl/download/current/kwgs.jpeg",
		"http://pogoda.topr.pl/download/current/hala.jpeg",
		"http://pogoda.topr.pl/download/current/kscw.jpeg",
		"http://pogoda.topr.pl/download/current/dcho.jpeg",
		"http://pogoda.topr.pl/download/current/mors.jpeg",
		"http://pogoda.topr.pl/download/current/momn.jpeg",
		"http://pogoda.topr.pl/download/current/psps.jpeg",
		"http://pogoda.topr.pl/download/current/psdb.jpeg",
		"http://pogoda.topr.pl/download/current/kwgr.jpeg",
		"http://pogoda.topr.pl/download/current/hgkw.jpeg"}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tatry.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Konrad Kuznicki", "Konrad Kuznicki")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "GNU GPLv3", "Open")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Konrad Kuznicki konrad@kuznicki.me")
	viper.SetDefault("license", "GNU GPLv3")

	config := &app.Config{}

	rootCmd.Flags().IntP("download-parallelism", "p", 2, "how many download routines should be started")
	rootCmd.Flags().DurationP("download-interval", "i", time.Minute*6, "how often should app refresh background images")
	rootCmd.Flags().DurationP("background-rotation", "r", time.Second*30, "how often should app rotate background image")
	rootCmd.Flags().StringSliceP("URLs", "u", []string{}, fmt.Sprintf("list of URLs from which images should be rotated on background (default [%s, ...%d])", cams[0], len(cams)))
	rootCmd.Flags().StringP("cache-dir", "c", "", "directory to store downloaded background images (default is $HOME/.cache/tatry)")
	viper.BindPFlags(rootCmd.Flags())

	rootCmd.Run = func(cmd *cobra.Command, args []string) {

		if err := viper.Unmarshal(config); err != nil {
			log.Fatal(err)
		}

		if config.DownloadConcurrency <= 0 {
			log.Fatal("download-parallelism has to be at least 1")
		}

		if len(config.Cams) == 0 {
			config.Cams = cams
		}

		if config.CacheLocation == "" {
			homePath, err := homedir.Dir()
			if err != nil {
				log.Fatal(errors.Wrap(err, "could not find home dir"))
			}
			config.CacheLocation = path.Join(homePath, ".cache/tatry")
		}

		if err := os.MkdirAll(config.CacheLocation, os.ModePerm); err != nil {
			log.Fatal(errors.Wrap(err, "could not create cache dir"))
		}

		b := app.NewApp(config)
		b.Run()
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	defaultFileExists := true
	if cfgFile != "" {
		log.Println("test", cfgFile)
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		if _, err := os.Stat(path.Join(home, ".tatry")); os.IsNotExist(err) {
			defaultFileExists = false
		}
		// Search config in home directory with name ".tatry" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tatry")
	}

	viper.SetEnvPrefix("TATRY")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil && (cfgFile != "" || defaultFileExists) {
		log.Fatal(err)
	} else {
		if defaultFileExists {
			log.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			log.Println("Using config file:", "default config file does not exit")
		}
	}
}
