package config

// Shard is a struct that represents a shard
type Shard struct {
	Name    string
	Idx     int
	Address string
}

// Config is a struct that represents the configuration of the sharding
type Config struct {
	Shards []Shard
}
