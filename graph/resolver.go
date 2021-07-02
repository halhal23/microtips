package graph

import (
	articleClient "microtips/article/client"
	userClient "microtips/user/client"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleClient *articleClient.Client
	UserClient    *userClient.Client
}
