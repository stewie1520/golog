package discovery

import (
	"github.com/hashicorp/serf/serf"
	"go.uber.org/zap"
)

// isLocal Check if a member is a local member
func (m *Membership) isLocal(member serf.Member) bool {
	return m.serf.LocalMember().Name == member.Name
}

// Members Get all members of current membership
func (m *Membership) Members() []serf.Member {
	return m.serf.Members()
}

// logError log error with member's name and rpc_addr
func (m *Membership) logError(err error, msg string, member serf.Member) {
	m.logger.Error(
		msg,
		zap.Error(err),
		zap.String("name", member.Name),
		zap.String("rpc_addr", member.Tags["rpc_addr"]))
}

// Leave member leave current membership
func (m *Membership) Leave() error {
	return m.serf.Leave()
}
