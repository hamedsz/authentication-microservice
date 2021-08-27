package auth

import "auth_micro/helpers/env"

var DefaultAuthService = env.GetEnv("AUTH_DEFAULT_SERVICE")