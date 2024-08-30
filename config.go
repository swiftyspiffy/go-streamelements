package streamelements

type Config struct {
	WebsocketEndpoint string
	JWT               string
}

func DefaultConfig(jwt string) *Config {
	return &Config{
		WebsocketEndpoint: "wss://realtime.streamelements.com/socket.io/?cluster=main&EIO=3&transport=websocket",
		JWT:               jwt,
	}
}
