package comment

func (service CommentService) DeleteComment(CommentID uint) (err error) {
	err = service.CommentRepository.DeleteComment(service.DB, CommentID)
	if err != nil {
		return err
	}
	return nil
}
