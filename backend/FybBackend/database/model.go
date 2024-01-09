package database

import (
	"time"
)

type User struct {
	ID               int64 `gorm:"column:ID;primaryKey"`
	Account          string
	Password         string
	PhoneNumber      string `gorm:"column:phoneNumber"`
	NickName         string `gorm:"column:nickName"`
	Job              string
	AvatarUrl        string `gorm:"column:avatarUrl"`
	Sex              string
	Area             string
	Year             int64
	TargetEnterprise string `gorm:"column:targetEnterprise"`
	TargetJob        string `gorm:"column:targetJob"`
	Slogan           string
	Balance          int64
	College          string
	State            int64     `gorm:"column:state"`
	Email            int64     `gorm:"column:email"`
	RegisterTime     time.Time `gorm:"column:registerTime"`
}

type Admin struct {
	ID          int64 `gorm:"column:ID;primaryKey"`
	Account     string
	Password    string
	PhoneNumber string `gorm:"column:phoneNumber"`
	Token       string
}

type Enterprise struct {
	ID          int64 `gorm:"column:ID;primaryKey"`
	Account     string
	Password    string
	Name        string
	City        string
	Type        string
	Region      string
	Logo        string
	Scale       string
	Recruitment string
	Business    string
	Founder     string
	FoundDate   time.Time `gorm:"column:foundDate"`
	Desc        string
	PhoneNumber string `gorm:"column:phoneNumber"`
	Token       string
}

type News struct {
	ID          int64 `gorm:"column:ID;primaryKey"`
	Author      string
	Title       string
	Content     string
	PublishTime string `gorm:"column:publishTime"`
	UserID      int64  `gorm:"column:userID"`
	Type        string `gorm:"column:type"`
}

type Job struct {
	ID           int64      `gorm:"primaryKey"`
	Name         string     `gorm:"column:name"`
	Education    string     `gorm:"column:education"`
	Wage         string     `gorm:"column:wage"`
	City         string     `gorm:"column:city"`
	Region       string     `gorm:"column:region"`
	Type         string     `gorm:"column:type"`
	Tag          string     `gorm:"column:Tag"`
	Category     string     `gorm:"column:category"`
	EnterpriseID string     `gorm:"column:enterpriseID"`
	Benefit      string     `gorm:"column:benefit"`
	Requirement  string     `gorm:"column:requirement"`
	Message      string     `gorm:"column:message"`
	Introduction string     `gorm:"column:introduction"`
	Deadline     string     `gorm:"column:deadline"`
	Enterprise   Enterprise `gorm:"foreignKey:enterpriseID"`
}

type Part struct {
	ID       int64  `gorm:"column:ID;primaryKey"`
	PartName string `gorm:"column:partName"`
}

type Post struct {
	ID          int64     `gorm:"column:ID;primaryKey"`
	Summary     string    `gorm:"column:summary"`
	Content     string    `gorm:"column:content"`
	State       int64     `gorm:"column:state"`
	Author      User      `gorm:"foreignKey:authorID"`
	Part        Part      `gorm:"foreignKey:partID"`
	PartID      int64     `gorm:"column:partID"`
	AuthorID    int64     `gorm:"column:authorID"`
	Favorite    int64     `gorm:"column:favorite"`
	Like        int64     `gorm:"column:like"`
	PublishTime time.Time `gorm:"column:publishTime"`
	Title       string    `gorm:"column:title"`
	CoverUrl    string    `gorm:"column:coverUrl"`
	Reward      int64     `gorm:"column:reward"`
}

type Comment struct {
	ID          int64     `gorm:"column:ID;primaryKey"`
	UserID      string    `gorm:"column:userID"`
	Content     string    `gorm:"column:content"`
	CommentNum  int64     `gorm:"column:CommentNum"`
	TargetPost  int64     `gorm:"column:targetPost"`
	PublishTime time.Time `gorm:"column:publishTime"`
	State       int64     `gorm:"column:state"`
	Author      User      `gorm:"foreignKey:userID"`
}

type Feedback struct {
	ID      int64     `gorm:"column:ID;primaryKey"`
	UserID  string    `gorm:"column:userID"`
	Content string    `gorm:"column:content"`
	State   int64     `gorm:"column:state"`
	Time    time.Time `gorm:"column:time"`
	Author  User      `gorm:"foreignKey:userID"`
}

type FavoriteRecord struct {
	ID        int64  `gorm:"column:ID;primaryKey"`
	UserID    string `gorm:"column:userID"`
	ArticleID string `gorm:"column:articleID"`
	Article   Post   `gorm:"foreignKey:userID"`
	Author    User   `gorm:"foreignKey:articleID"`
}

type LikeRecord struct {
	ID     int64  `gorm:"column:ID;primaryKey"`
	UserId string `gorm:"column:userId"`
	PostId string `gorm:"column:postId"`
	Post   Post   `gorm:"foreignKey:userId"`
	Author User   `gorm:"foreignKey:postId"`
}

type AdoptRecord struct {
	ID        int64  `gorm:"column:ID;primaryKey"`
	PostId    string `gorm:"column:postId;primaryKey"`
	CommentId string `gorm:"column:commentId;primaryKey"`
}
