package singleSubscription

import (
	"encoding/json"

	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

type PaginationList struct {
	GetList getter `json:"-"`
	Cursor  uint64 `json:"cursor"`
	Limit   uint64 `json:"limit"`
}

type button struct {
	Text string
	Data string
}

type getter func(cursor, limit uint64) ([]subscription.SingleSubscription, error)

func NewPaginationList(g getter, cursor, limit uint64) *PaginationList {
	if cursor == 0 {
		cursor = 1
	}
	return &PaginationList{
		GetList: g,
		Cursor:  cursor,
		Limit:   limit,
	}
}

func (pl *PaginationList) Page() string {
	list, err := pl.GetList(pl.Cursor, pl.Limit)
	if len(list) == 0 || err != nil {
		return ErrEmptyList
	}
	txt := "Список элементов:\n"
	txt += list[0].String() + "\n"
	for _, v := range list[1:] {
		txt += v.String() + "\n"
	}
	return txt
}

func (pl *PaginationList) Prev() string {
	var prevCursor uint64 = 1
	needRender := true
	if pl.Cursor == 1 {
		needRender = false
	} else if pl.Limit > pl.Cursor && pl.Cursor != 1 {
		prevCursor = 1
	} else {
		prevCursor = pl.Cursor - pl.Limit
		if prevCursor == 0 {
			prevCursor = 1
		}
	}
	if !needRender {
		return ""
	}
	return pl.serialize(prevCursor, pl.Limit)
}

func (pl *PaginationList) Next() string {
	nextCursor := pl.Cursor + pl.Limit
	l, _ := pl.GetList(nextCursor, pl.Limit)
	if len(l) == 0 {
		return ""
	}
	return pl.serialize(nextCursor, pl.Limit)
}

func (pl *PaginationList) Buttons() []*button {
	buttons := make([]*button, 0, 2)

	if b := pl.Prev(); b != "" {
		buttons = append(buttons, &button{
			Text: "⬅️ Prev",
			Data: b,
		})
	}

	if b := pl.Next(); b != "" {
		buttons = append(buttons, &button{
			Text: "Next ➡️",
			Data: b,
		})
	}

	return buttons
}

func (pl *PaginationList) serialize(cursor, limit uint64) string {
	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: uint64(cursor),
		Limit:  uint64(limit),
	})

	callbackPath := path.CallbackPath{
		Domain:       DomainName,
		Subdomain:    SubdomainName,
		CallbackName: CallbackNameList,
		CallbackData: string(serializedData),
	}
	return callbackPath.String()
}