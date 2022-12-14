package config

import (
	"errors"
	"os"
)

type Mongo struct {
	HostURI          string
	BlaastDB         string
	BlaastCollection string
}

type Server struct {
	Port string
}

type Config struct {
	Server *Server
	Mongo  *Mongo
}

var config = Config{}

func GetConfig() (*Config, error) {
	server, err := getServerConfig()
	if err != nil {
		return nil, err
	}
	config.Server = server
	mongoConfig, err := getMongoConfig()
	if err != nil {
		return nil, err
	}
	config.Mongo = mongoConfig
	return &config, nil
}

func getServerConfig() (*Server, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Server{Port: port}, nil
}

func getMongoConfig() (*Mongo, error) {
	host := os.Getenv("MONGODB_HOST_URI")
	if host == "" {
		return nil, errors.New("missing mongodb host uri in env")
	}
	blaastDb := os.Getenv("MONGODB_BLAAST_DB")
	if blaastDb == "" {
		return nil, errors.New("missing blaastDb name in env")
	}
	blaastCollection := os.Getenv("MONGODB_BLAAST_COLLECTION")
	if blaastCollection == "" {
		return nil, errors.New("missing blaast collection name in env")
	}

	return &Mongo{
		HostURI:          host,
		BlaastDB:         blaastDb,
		BlaastCollection: blaastCollection,
	}, nil
}
