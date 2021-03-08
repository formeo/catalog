package server

import (
	"context"

	"go.uber.org/zap"
)

const metadataKey = "METADATA_KEY"

type Metadata struct {
	Log *zap.Logger
}

func GetMetadataFromContext(c context.Context) *Metadata {
	md, ok := c.Value(metadataKey).(*Metadata)
	if !ok {
		return nil
	}

	return md
}

func GetLogger(ctx context.Context) *zap.Logger {
	md := GetMetadataFromContext(ctx)
	if md != nil {
		return md.Log
	}
	return zap.NewNop()
}

func contextWithMetadata(c context.Context, md *Metadata) context.Context {
	return context.WithValue(c, metadataKey, md)
}
