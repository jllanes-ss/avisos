package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	//"cloud.google.com/go/storage"
	//"github.com/go-redis/redis/v8"
	//"github.com/jmoiron/sqlx"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

type dataSources struct {
	DB          *sql.DB
	RedisClient *redis.Client
	//StorageClient *storage.Client
}

// InitDS establishes connections to fields in dataSources
func GetDS() (*dataSources, error) {

	engine := viper.GetString(`USER_REPOSITORY_ENGINE`)

	var (
		dbHost     = ""
		dbPort     = ""
		dbUser     = ""
		dbPass     = ""
		dbName     = ""
		dbSSL      = ""
		connection = ""
		dsn        = ""
	)
	if engine == "postgres" {
		dbHost = viper.GetString(`PG_HOST`)
		dbPort = viper.GetString(`PG_PORT`)
		dbUser = viper.GetString(`PG_USER`)
		dbPass = viper.GetString(`PG_PASSWORD`)
		dbName = viper.GetString(`PG_DB`)
		dbSSL = viper.GetString(`PG_SSL`)

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSSL)

	} else {
		dbHost = viper.GetString(`MYSQL_HOST`)
		dbPort = viper.GetString(`MYSQL_PORT`)
		dbUser = viper.GetString(`MYSQL_USER`)
		dbPass = viper.GetString(`MYSQL_PASSWORD`)
		dbName = viper.GetString(`MYSQL_DATABASE_NAME`)
		connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		val := url.Values{}
		val.Add("parseTime", "1")
		val.Add("loc", "Asia/Jakarta")
		dsn = fmt.Sprintf("%s?%s", connection, val.Encode())

	}
	db, err := sql.Open(engine, dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	//redis datasource

	// Initialize redis connection
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetString("REDIS_PORT")

	log.Printf("Connecting to Redis\n")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	// verify redis connection

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}

	return &dataSources{
		DB:          db,
		RedisClient: rdb,
	}, nil
}

// close to be used in graceful server shutdown
func (d *dataSources) Close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	if err := d.RedisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}
