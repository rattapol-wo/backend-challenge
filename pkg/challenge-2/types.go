package challenge2

type requestKeyBoardEncoded struct {
	Encoded []string `json:"encoded" validate:"required"`
}

type responseKeyBoardDecoded struct {
	Decoded string
	Sum      int
}