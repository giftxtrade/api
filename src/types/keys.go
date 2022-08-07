package types

import "time"

const NameKey string = "name"

type AuthKeyType string
const AuthKey AuthKeyType = "auth"
const AuthHeader string = "Authorization"

const DateTimeFormat string = time.RFC3339