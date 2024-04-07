package configtx

const (
	// ChannelV2_0 is the capabilities string for standard new non-backwards compatible fabric v2.0 channel capabilities.
	ChannelV2_0 = "V2_0"
	// ConsensusTypeSmartBFT identifies the SmartBFT-based consensus implementation.
	ConsensusTypeSmartBFT = "smartbft"
	// DefaultConsortiumValue value for consortium
	DefaultConsortiumValue = "DefaultConsortium"
	// SHA2 is an identifier for SHA2 hash family
	SHA2 = "SHA2"
	// SHA256
	SHA256 = "SHA256"
	// Role values for principals
	RoleAdmin   = "admin"
	RoleMember  = "member"
	RoleClient  = "client"
	RolePeer    = "peer"
	RoleOrderer = "orderer"
	// Standart rules
	AnyReaders          = "ANY Readers"
	AnyWriters          = "ANY Writers"
	MajorityAdmins      = "MAJORITY Admins"
	MajorityEndorsement = "MAJORITY Endorsement"
)
