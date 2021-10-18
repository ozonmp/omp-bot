package announcement

type AnnouncementService interface {
	Describe(announcementID uint64) (*Announcement, error)
	List(cursor uint64, limit uint64) ([]Announcement, error)
	Create(announcement Announcement) (uint64, error)
	Update(announcementID uint64, announcement Announcement) error
	Remove(announcementID uint64) (bool, error)
}

type DummyAnnouncementService struct {
	Announcements []Announcement
}

func NewDummyAnnouncementService() *DummyAnnouncementService {
	return &DummyAnnouncementService{Announcements: allEntities}
}

func (d *DummyAnnouncementService) Describe(announcementID uint64) (*Announcement, error) {
	idx := announcementID - 1
	if idx < 0 || idx > uint64(len(d.Announcements) - 1) {
		return nil, nil
	}

	return &d.Announcements[idx], nil
}

func (d *DummyAnnouncementService) List(cursor uint64, limit uint64) ([]Announcement, error) {
	start := int(cursor)
	end := start + int(limit)
	length := len(d.Announcements)
	if end > length {
		end = length
	}
	return d.Announcements[start:end], nil
}

func (d *DummyAnnouncementService) Create(announcement Announcement) (uint64, error) {
	return 0, nil
}

func (d *DummyAnnouncementService) Update(announcementID uint64, announcement Announcement) error {
	return nil
}

func (d *DummyAnnouncementService) Remove(announcementID uint64) (bool, error) {
	return false, nil
}