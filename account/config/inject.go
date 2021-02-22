package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	_tokenRepository "github.com/jllanes-ss/avisos/account/token/repository"
	_userRepository "github.com/jllanes-ss/avisos/account/user/repository"
	"github.com/jllanes-ss/avisos/account/user/usecase"
	"github.com/spf13/viper"
)

func Inject(d *dataSources) (*fiber.App, error) {
	log.Println("Injecting data sources")

	userRepository := _userRepository.GetRepository(d.DB)
	tokenRepository := _tokenRepository.GetRepository(d.RedisClient)

	userUseCase := usecase.NewUserUseCase(&usecase.UUCConfig{
		UserRepository: userRepository,
	})

	// load rsa keys
	privKeyFile := viper.GetString("PRIV_KEY_FILE")
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	pubKeyFile := viper.GetString("PUB_KEY_FILE")
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not read public key pem file: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	// load refresh token secret from env variable
	refreshSecret := viper.GetString("REFRESH_SECRET")

	// load expiration lengths from env variables and parse as int
	idTokenExp := viper.GetString("ID_TOKEN_EXP")
	refreshTokenExp := viper.GetString("REFRESH_TOKEN_EXP")

	idExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse ID_TOKEN_EXP as int: %w", err)
	}

	refreshExp, err := strconv.ParseInt(refreshTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse REFRESH_TOKEN_EXP as int: %w", err)
	}

	tokenUseCase := useCase.NewTokenUseCase(&useCase.TUCConfig{
		TokenRepository:       tokenRepository,
		PrivKey:               privKey,
		PubKey:                pubKey,
		RefreshSecret:         refreshSecret,
		IDExpirationSecs:      idExp,
		RefreshExpirationSecs: refreshExp,
	})

	// initialize fiber APP
	app := fiber.New()

	// read in ACCOUNT_API_URL
	baseURL := viper.GetString("ACCOUNT_API_URL")

	// read in HANDLER_TIMEOUT
	handlerTimeout := viper.GetString("HANDLER_TIMEOUT")
	ht, err := strconv.ParseInt(handlerTimeout, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse HANDLER_TIMEOUT as int: %w", err)
	}

	maxBodyBytes := viper.GetString("MAX_BODY_BYTES")
	mbb, err := strconv.ParseInt(maxBodyBytes, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse MAX_BODY_BYTES as int: %w", err)
	}

	handler.NewHandler(&handler.Config{
		F:               app,
		UserUseCase:     userUseCae,
		TokenUseCase:    tokenUseCase,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(ht) * time.Second),
		MaxBodyBytes:    mbb,
	})

	return app, nil

}
