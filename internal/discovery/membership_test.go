package discovery

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
	"testing"
	"time"
)

func TestMembership(t *testing.T) {
	m, handler := setupMember(t, nil)
	m, _ = setupMember(t, m)
	m, _ = setupMember(t, m)
	require.Eventually(t, func() bool {
		return 2 == len(handler.joins) && 3 == len(m[0].Members()) && 0 == len(handler.leaves)
	}, 10*time.Second, 250*time.Millisecond)
}

type handler struct {
	joins  chan map[string]string
	leaves chan string
}

func (h *handler) Join(id, addr string) error {
	if h.joins != nil {
		h.joins <- map[string]string{
			"id":   id,
			"addr": addr,
		}
	}
	return nil
}

func (h *handler) Leave(id string) error {
	if h.leaves != nil {
		h.leaves <- id
	}
	return nil
}

func setupMember(t *testing.T, members []*Membership) ([]*Membership, *handler) {
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", dynaport.Get(1)[0])
	fmt.Println(addr)
	c := Config{
		NodeName: fmt.Sprintf("%d", len(members)), // id
		BindAddr: addr,
		Tags:     map[string]string{
			"rpc_addr": addr,
		},
	}

	h := &handler{}
	if len(members) == 0 {
		h.joins = make(chan map[string]string, 3)
		h.leaves = make(chan string, 3)
	} else {
		c.StartJoinAddrs = []string{
			members[0].BindAddr,
		}
	}
	m, err := New(h, c)
	require.NoError(t, err)
	members = append(members, m)
	return members, h
}
