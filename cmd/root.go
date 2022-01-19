package cmd

import (
	"flag"
	"fmt"
	"os"

	"twiscode-test/config"
	"twiscode-test/internal/app/appcontext"
	"twiscode-test/internal/app/commons"
	"twiscode-test/internal/app/migration"
	"twiscode-test/internal/app/server"
	"twiscode-test/internal/app/service"
	"twiscode-test/internal/app/service/funcService"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "twiscode-test",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.Config()
	app := appcontext.NewAppContext(cfg)
	var err error

	logLevel := zerolog.InfoLevel
	logLevelP, err := zerolog.ParseLevel(os.Getenv("APP_LOG_LEVEL"))
	if err == nil {
		logLevel = logLevelP
	}
	zerolog.SetGlobalLevel(logLevel)

	validator := validator.New()

	var mysqlDB *gorm.DB
	if app.GetMysqlOption().IsEnable {
		mysqlDB, err = app.GetDBInstance(appcontext.DBDialectMysql)
		if err != nil {
			log.Info().Msgf("Failed to start, error connect to DB MySQL | %v", err)
			return
		}
	}

	opt := commons.Options{
		AppCtx:    app,
		Db:        mysqlDB,
		UUID:      commons.NewUuid(),
		Validator: validator,
	}

	//run migration
	runMigration(&opt)

	// repo := wiringRepository(repository.Option{
	// 	Options: opt,
	// })

	service := wiringService(service.Option{
		Options: opt,
		// Repositories: repo,
	})

	server := server.NewServer(opt, service)

	// run app
	server.StartApp()
}

func runMigration(opt *commons.Options) {
	pathMigration := os.Getenv("APP_MIGRATION_PATH")
	migrationDir := flag.String("migration-dir", pathMigration, "migration directory")
	log.Info().Msg("path migration : " + pathMigration)

	migrationConf, errMigrationConf := migration.NewMigrationConfig(*migrationDir,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE_NAME"),
		"mysql")
	if errMigrationConf != nil {
		log.Error().Msg(errMigrationConf.Error())
		return
	}

	errMigration := migration.MigrateUp(migrationConf)
	if errMigration != nil {
		if errMigration.Error() != "no change" {
			log.Error().Msg(errMigration.Error())
			return
		}
		log.Info().Msg("Migration success : no change table . . .")
	}
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	svc := service.Services{
		FuncService: funcService.NewFuncService(),
	}
	return &svc
}
