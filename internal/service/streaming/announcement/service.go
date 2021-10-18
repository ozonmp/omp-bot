package announcement

import (
	"errors"
)

type AnnouncementService interface {
	Describe(announcementID uint64) (*Announcement, error)
	List(cursor uint64, limit uint64) ([]Announcement, error)
	Create(announcement Announcement) (uint64, error)
	Update(announcementID uint64, announcement Announcement) error
	Remove(announcementID uint64) (bool, error)
}

type DummyAnnouncementService struct {
	Announcements []Announcement
	LastID uint64
}

func NewDummyAnnouncementService() *DummyAnnouncementService {
	return &DummyAnnouncementService{Announcements: allEntities, LastID: uint64(len(allEntities))}
}

func (d *DummyAnnouncementService) Describe(announcementID uint64) (*Announcement, error) {
	var item *Announcement
	for _, val := range d.Announcements {
		if val.ID == announcementID {
			item = &val
			break
		}
	}
	if item == nil {
		return nil, nil
	}

	return item, nil
}

func (d *DummyAnnouncementService) List(cursor uint64, limit uint64) ([]Announcement, error) {
	start := cursor * limit
	end := start + limit
	length := uint64(len(d.Announcements))
	if end > length {
		end = length
	}
	return d.Announcements[start:end], nil
}

func (d *DummyAnnouncementService) Create(announcement Announcement) (uint64, error) {
	d.LastID++

	announcement.ID = d.LastID
	d.Announcements = append(d.Announcements, announcement)

	return d.LastID, nil
}

func (d *DummyAnnouncementService) Update(announcementID uint64, announcement Announcement) error {
	var item *Announcement
	idx := -1
	for i, val := range d.Announcements {
		if val.ID == announcementID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("item not found")
	}

	item = &d.Announcements[idx]
	item.Author = announcement.Author
	item.TimePlanned = announcement.TimePlanned
	item.Title = announcement.Title
	item.Description = announcement.Description
	item.ThumbnailUrl = announcement.ThumbnailUrl

	return nil
}

func (d *DummyAnnouncementService) Remove(announcementID uint64) (bool, error) {
	idx := -1
	for i, val := range d.Announcements {
		if val.ID == announcementID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false, nil
	}
	d.Announcements = append(d.Announcements[:idx], d.Announcements[idx + 1:]...)
	return true, nil
}