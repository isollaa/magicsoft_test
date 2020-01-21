package main

import "log"

type ServerConfig struct {
	Driver  string
	URI     string
	Enabled bool
}

type Server struct {
	config ServerConfig
}

type Config interface {
	DBConfigCreation() *ServerConfig
}

func NewServer(c Config) *Server {
	return &Server{
		config: *c.DBConfigCreation(),
	}
}

type sqlConfig ServerConfig

func (s *sqlConfig) DBConfigCreation() *ServerConfig {
	return &ServerConfig{
		Driver:  s.Driver,
		URI:     s.URI,
		Enabled: s.Enabled,
	}
}

type mongoConfig ServerConfig

func (s *mongoConfig) DBConfigCreation() *ServerConfig {
	return &ServerConfig{
		Driver:  s.Driver,
		URI:     s.URI,
		Enabled: s.Enabled,
	}
}

func main() {
	server := NewServer(&sqlConfig{
		Driver:  "sql",
		URI:     "sss",
		Enabled: true})
	// Do something with server
	log.Printf("%v", server)

	server = NewServer(&mongoConfig{
		Driver:  "mongo",
		URI:     "mmm",
		Enabled: false})
	// Do something with server
	log.Printf("%v", server)
}
