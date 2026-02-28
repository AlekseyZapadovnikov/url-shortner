package conf

type ShortenerConfig struct {
	shortUrlLength int
}

func (sh *ShortenerConfig) GetAlphabet() []rune {
	return make([]rune, 0) // TODO add impl
}

func (sh *ShortenerConfig) GetAlphabetLen() int {
	return sh.shortUrlLength
}
