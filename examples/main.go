package main

import (
	"context"
	"fmt"

	grpc_metadata "github.com/go-godin/grpc-metadata"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()
	md := metadata.Pairs("requestId", uuid.New().String(), "accountId", "123", "userId", "some_user")
	ctx = metadata.NewIncomingContext(ctx, md)

	requestID := grpc_metadata.GetRequestID(ctx)
	fmt.Println(requestID)

	if grpc_metadata.HasAccountID(ctx) {
		fmt.Println(grpc_metadata.GetAccountID(ctx))
	}

	if grpc_metadata.HasUserID(ctx) {
		fmt.Println(grpc_metadata.GetUserID(ctx))
	}
}
