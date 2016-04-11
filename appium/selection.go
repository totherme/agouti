package appium

import "github.com/totherme/agouti/internal/element"

type elementRepository interface {
	Get() ([]element.Element, error)
	GetAtLeastOne() ([]element.Element, error)
	GetExactlyOne() (element.Element, error)
}
