package model

type Photo struct {
	ID            int64
	RestaurantID  string
	Name          string
	Image         string
	ImageWebp     string
	Thumbnail     string
	ThumbnailWebp string
}
