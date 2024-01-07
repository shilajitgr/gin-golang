package service

import "gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	Findall() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

// Findall implements VideoService.
func (service *videoService) Findall() []entity.Video {
	return service.videos
}
