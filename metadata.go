package grpc_metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type MetadataContextKey string

const (
	RequestID MetadataContextKey = "requestId"
	AccountID MetadataContextKey = "accountId"
	UserID    MetadataContextKey = "userId"
)

// GetRequestID tries to extract the requestId key from the given context.
func GetRequestID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		requestID := md.Get(string(RequestID))
		if len(requestID) > 0 {
			return requestID[0]
		}
	}
	return ""
}

// GetAccountID tries to extract the accountId key from the given context.
// If no AccountID exists, an empty string is returned.
func GetAccountID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		accountID := md.Get(string(AccountID))
		if len(accountID) > 0 {
			return accountID[0]
		}
	}
	return ""
}

// GetUserID tries to extract the userId key from the given context.
// If no UserID exists, an empty string is returned
func GetUserID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		userID := md.Get(string(UserID))
		if len(userID) > 0 {
			return userID[0]
		}
	}
	return ""
}

// Has checks whether the passed key exists in the context metadata
func Has(ctx context.Context, key MetadataContextKey) bool {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		vals := md.Get(string(key))
		if len(vals) > 0 {
			return true
		}
	}
	return false
}

func HasAccountID(ctx context.Context) bool {
	return Has(ctx, AccountID)
}

func HasUserID(ctx context.Context) bool {
	return Has(ctx, UserID)
}
