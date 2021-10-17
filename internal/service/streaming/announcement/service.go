package announcement

type AnnouncementService interface {
	Describe(AnnouncementID uint64) (*Announcement, error)
	List(cursor uint64, limit uint64) ([]Announcement, error)
	Create(Announcement) (uint64, error)
	Update(AnnouncementID uint64, announcement Announcement) error
	Remove(AnnouncementID uint64) (bool, error)
}

type DummyAnnouncementService struct {}

func NewDummyAnnouncementService() AnnouncementService {
	return DummyAnnouncementService{}
}

func (d DummyAnnouncementService) Describe(AnnouncementID uint64) (*Announcement, error) {
	return nil, nil
}

func (d DummyAnnouncementService) List(cursor uint64, limit uint64) ([]Announcement, error) {
	return nil, nil
}

func (d DummyAnnouncementService) Create(Announcement) (uint64, error) {
	return 0, nil
}

func (d DummyAnnouncementService) Update(AnnouncementID uint64, announcement Announcement) error {
	return nil
}

func (d DummyAnnouncementService) Remove(AnnouncementID uint64) (bool, error) {
	return false, nil
}