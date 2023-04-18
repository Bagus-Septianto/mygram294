package photo

func (service PhotoService) DeletePhoto(PhotoID uint) (err error) {
	err = service.PhotoRepository.DeletePhoto(service.DB, PhotoID)
	if err != nil {
		return err
	}
	return nil
}
