package handler

type Code string

const (
	codeBadBody Code = "BAD_BODY"
	codeContentEmpty Code = "EMPTY_CONTENT"
	codeInternalServerError Code = "INTERNAL_SERVER_ERROR"
	codeOK Code = "OK"
	codeNoParam Code = "NO_PARAM"
	codeInvalidConvertion Code = "INVALID_CONVERSION"
	codeInvalidPostId Code = "INVALID_POST_ID"
)