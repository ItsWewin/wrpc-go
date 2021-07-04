package conf

type conf struct {
	addr string
	network string
}

var confDefault = conf{
	addr: "127.0.0.1:8082",
	network: "tcp",
}

func Addr() string {
	return confDefault.addr
}

func Network() string {
	return confDefault.network
}
