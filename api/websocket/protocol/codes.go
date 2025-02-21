package protocol

type StatusCode int

const (
	StatusOK               = 200 // RFC 9110, 15.3.1
	StatusCreated          = 201 // RFC 9110, 15.3.2
	StatusAccepted         = 202 // RFC 9110, 15.3.3
	StatusNoContent        = 204 // RFC 9110, 15.3.5
	StatusResetContent     = 205 // RFC 9110, 15.3.6
	StatusPartialContent   = 206 // RFC 9110, 15.3.7
	StatusBadRequest       = 400 // RFC 9110, 15.5.1
	StatusUnauthorized     = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired  = 402 // RFC 9110, 15.5.3
	StatusForbidden        = 403 // RFC 9110, 15.5.4
	StatusNotFound         = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable    = 406 // RFC 9110, 15.5.7
	StatusRequestTimeout   = 408 // RFC 9110, 15.5.9
	StatusConflict         = 409 // RFC 9110, 15.5.10
	StatusGone             = 410 // RFC 9110, 15.5.11

	StatusInternalServerError = 500 // RFC 9110, 15.6.1
	StatusNotImplemented      = 501 // RFC 9110, 15.6.2
	StatusBadGateway          = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable  = 503 // RFC 9110, 15.6.4
)

type PacketType byte

type RequestPacketType PacketType

const (
	Move RequestPacketType = iota + 1
	GameStateRequest
	LegalMovesRequests
	Promotion
	Resignation
	PingRequest
	UndoMoveRequest
	DrawOfferRequest
	AbortRequest
	Chat
)

type ResponsePacketType PacketType

const (
	Acknowledgement ResponsePacketType = iota + 100
	GameStateResponse
	LegalMovesResponse
	GameOver
	PlayerStatusChanged
	PingResponse
	UndoMoveResponse
	DrawOfferResponse
)
