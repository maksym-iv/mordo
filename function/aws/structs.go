package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"

	ms "github.com/mitchellh/mapstructure"
)

type QS struct {
	// resize Resize `mapstructure:",squash"`
	Image     string
	Resize    Resize `mapstructure:",squash"`
	Watermark `mapstructure:",squash"`
	SCrop     SCrop `mapstructure:",squash"`
	Crop      Crop  `mapstructure:",squash"`
	DPR       float64
	Sharpen   bool
}

type Resize struct {
	Width  int
	Heigth int
}
type SCrop struct {
	Width  int `mapstructure:"sc_x"`
	Heigth int `mapstructure:"sc_y"`
}
type Crop struct {
	Left   int `mapstructure:"c_top"`
	Top    int `mapstructure:"c_left"`
	Width  int `mapstructure:"c_x"`
	Heigth int `mapstructure:"c_y"`
}
type Watermark struct {
	WX string  `mapstructure:"w_x"`
	WY string  `mapstructure:"w_y"`
	WS float64 `mapstructure:"w_s"`
}

func newQs(image string, qs map[string]string) (*QS, error) {
	result := &QS{}
	result.Image = image

	// qsi := atoiface(qs)

	config := &ms.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := ms.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(qs); err != nil {
		return nil, err
	}

	return result, nil
}

func (qss *QS) hashPath() string {
	// m := structs.Map(qss)
	// elements := []byte(fmt.Sprintf("%v", qss))
	elements := append([]byte(fmt.Sprintf("%v", *qss)), []byte(fmt.Sprintf("%v", *config))...)

	h := sha1.New()
	h.Write(elements)

	p := hex.EncodeToString(h.Sum(nil))
	p = fmt.Sprintf("prc_%v_%v", p, qss.Image)
	return p

}

// errorJSON Custom error type with JSON marshaling
type errorJSON interface {
	ErrorJSON() string
}

type gwError struct {
	S string `json:"err"`
}

func newErr(err string) *gwError {
	return &gwError{err}
}

func (e *gwError) Error() string {
	return e.S
}

func (e *gwError) ErrorJSON() string {
	body, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(body[:])
}
