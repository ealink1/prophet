package dao

import "prophet/internal/model"

type MeditationDAO struct{}

func (d *MeditationDAO) FindActiveTracks() ([]model.MeditationTrack, error) {
	var tracks []model.MeditationTrack
	err := model.DB.Where("is_active = ?", true).Order("sort_order").Find(&tracks).Error
	return tracks, err
}
