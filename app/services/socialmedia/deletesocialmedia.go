package socialmedia

func (service SocialMediaService) DeleteSocialMedia(SocialMediaID uint) (err error) {
	err = service.SocialMediaRepository.DeleteSocialMedia(service.DB, SocialMediaID)
	if err != nil {
		return err
	}
	return nil
}
