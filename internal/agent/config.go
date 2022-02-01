package agent

import (
	"crypto/tls"
	"fmt"
	"net"
)

type Config struct {
	ServeTLSConfig *tls.Config
	PeerTLSConfig  *tls.Config

	DataDir        string
	BindAddr       string
	RPCPort        int
	NodeName       string
	StartJoinAddrs []string

	ACLModeFile   string
	ACLPolicyFile string
}

func (c Config) RPCAddr() (string, error) {
	host, _, err := net.SplitHostPort(c.BindAddr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%d", host, c.RPCPort), nil
}

