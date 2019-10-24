package igo

type clientType uint8

const (
	REST clientType = 0
	RPC  clientType = 1
)

type WeIdClient struct {
	server   []string
	callType clientType
}

func NewWeIdClient(server []string) *WeIdClient {
	we := &WeIdClient{
		server:   server,
		callType: REST,
	}
	return we
}

func NewWeIdRestClient(server []string) *WeIdClient {
	we := &WeIdClient{
		server:   server,
		callType: REST,
	}
	return we
}

func NewWeIdRpcClient(server []string) *WeIdClient {
	we := &WeIdClient{
		server:   server,
		callType: RPC,
	}
	return we
}
