package server

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	fiber "github.com/gofiber/fiber/v2"
	redis "github.com/0187773933/RedisManagerUtils/manager"
	buttons_route_handler "github.com/0187773933/FireStickC2Server/v1/routes/buttons"
	utils "github.com/0187773933/FireStickC2Server/v1/utils"
	// misc_route_handler "firestickc2server/v1/routes/misc"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Redis utils.RedisConfig `json:"redis"`
	Apps struct {
		Instagram struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Following []string `json:"following"`
		} `json:"instagram"`
		Skype struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Contacts []string `json:"contacts"`
		} `json:"skype"`
		Spotify struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Playlists []string `json:"playlists"`
		} `json:"spotify"`
		Twitch struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Following []string `json:"following"`
		} `json:"twitch"`
		Youtube struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Following []string `json:"following"`
		} `json:"youtube"`
		Zoom struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Contacts []string `json:"contacts"`
		} `json:"zoom"`
	}`json:"apps"`
}
type Server struct {
	FiberApp *fiber.App `json:"fiber_app"`
	RedisClient redis.Manager `json:"redis_client"`
	Config ServerConfig `json:"config"`
}

func NewServer( config_file_path string ) ( server Server ) {
	var config ServerConfig
	config_file , _ := ioutil.ReadFile( config_file_path )
	json.Unmarshal( []byte( config_file ) , &config )
	server = Server{}
	server.Config = config
	server.FiberApp = fiber.New()
	server.SetupRedis()
	server.SetupRoutes()
	return
}

func ( s *Server ) SetupRedis() {
	s.RedisClient = redis.Manager{}
	// test = test
	s.RedisClient.Connect(
		fmt.Sprintf( "%s:%s" , s.Config.Redis.Host , s.Config.Redis.Port ) ,
		s.Config.Redis.DB ,
		s.Config.Redis.Password ,
	)
}

func ( s *Server ) SetupRoutes() {
	buttons_route_handler.RegisterRoutes( s.FiberApp , &s.Config.Redis )
}

func ( s *Server ) Listen() {
	s.FiberApp.Listen( fmt.Sprintf( ":%s" , s.Config.Port ) )
}