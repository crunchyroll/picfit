package config

import "time"

// DefaultFormat is the default image format
const DefaultFormat = "png"

// DefaultQuality is the default quality for processed images
const DefaultQuality = 95

// DefaultPort is the default port of the application server
const DefaultPort = 3001

// DefaultShardWidth is the default shard width
const DefaultShardWidth = 0

// DefaultShardDepth is the default shard depth
const DefaultShardDepth = 0

const DefaultCacheControlDuration = int(5 * time.Minute)
