package creditcard

type card struct {
	number string
}

func (card card) New(number string) (card, error) {
	return card{number}, nil
}

func (c *card) Number() string {
	return c.number
}
