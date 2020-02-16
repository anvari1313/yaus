package shorturl

import (
	"github.com/anvari1313/yaus/config"
	"github.com/speps/go-hashids"
)

type HashID struct {
	h *hashids.HashID
}

func Create(c config.ShortenedURL) (*HashID, error) {
	hd := hashids.NewData()
	hd.Salt = c.Salt
	hd.MinLength = c.MinLength
	hd.Alphabet = c.Alphabet
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return &HashID{
		h: h,
	}, nil
}

func (h *HashID) Generate(id int64) (string, error) {
	e, err := h.h.EncodeInt64([]int64{id})
	if err != nil {
		return "", err
	}

	return e, nil
}
