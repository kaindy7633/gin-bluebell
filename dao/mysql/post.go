package mysql

import "gin-bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO  post(post_id, title, content, author_id, community_id) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostDetailById
func GetPostDetailById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time from post WHERE post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}
