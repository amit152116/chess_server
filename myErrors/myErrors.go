package myErrors

import "errors"

var SessionExpired = errors.New("token: session expired")

var InvalidCredentials = errors.New("token: invalid credentials")

var UserNotFound = errors.New("token: user not found")

var UserAlreadyExists = errors.New("token: user already exists")

var InvalidSession = errors.New(" token: invalid session id")

var SessionMissing = errors.New("token: session-id header is missing")

var Unauthorized = errors.New("token: unauthorized")

var Forbidden = errors.New("token: forbidden")

var InvalidEmail = errors.New("token: invalid email")

var InvalidFen = errors.New("invalid FEN")

var ErrBodyLenTooLarge = errors.New("protocol.Header: Body Length is Larger than 2 Byte")
