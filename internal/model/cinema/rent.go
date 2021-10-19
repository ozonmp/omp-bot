package cinema

import (
	"fmt"
	"time"
)

type Rent struct {
	// UglyHack. т.к. элементы могут удалятся, но при этом следование идентификаторов
	// у нас остаётся и в интерфейсе сервиса List мы отдаём slice, нам нужно
	// каким-то образом отобразить идентифиакторы записей
	RecordIndex  uint64
	FilmID       int64
	SerialID     int64
	PriceInKopec int64

	Deleted   bool
	DeletedAt time.Time
}

func NewFilmRent(FilmID int64, PriceInKopek int64) *Rent {
	return &Rent{
		FilmID:       FilmID,
		SerialID:     -1,
		PriceInKopec: PriceInKopek,
		Deleted:      false,
	}
}

func NewSerialRent(SerialID int64, PriceInKopec int64) *Rent {
	return &Rent{
		FilmID:       -1,
		SerialID:     SerialID,
		PriceInKopec: PriceInKopec,
		Deleted:      false,
	}
}

func (rent *Rent) String() string {
	return fmt.Sprintf(
		"film_id: %d, serial_id: %d, price(в копейках): %d",
		rent.FilmID,
		rent.SerialID,
		rent.PriceInKopec,
	)
}
