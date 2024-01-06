package database

import (
	"time"
)

type User struct {
	ID            int64 `gorm:"column:ID;primaryKey"`
	Account       string
	Password      string
	PhoneNumber   string `gorm:"column:phoneNumber"`
	NickName      string `gorm:"column:nickName"`
	Job           string
	AvatarUrl     string `gorm:"column:avatarUrl"`
	Sex           string
	Area          string
	Year          int64
	TargetCollege string `gorm:"column:targetCollege"`
	TargetJob     string `gorm:"column:targetJob"`
	Slogan        string
	Balance       int64
	College       string
	State         int64     `gorm:"column:state"`
	RegisterTime  time.Time `gorm:"column:registerTime"`
}

type Admin struct {
	ID          int64 `gorm:"column:ID;primaryKey"`
	Account     string
	Password    string
	PhoneNumber string `gorm:"column:phoneNumber"`
	Token       string
}

type Company struct {
	ID       int64 `gorm:"column:ID;primaryKey"`
	Name     string
	Code     string
	Guide    string
	Type     string
	Belong   string
	Logo     string
	Level    string
	LineType string `gorm:"column:lineType"`
	Profile  string
	Region   string
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
	ID                    int64  `gorm:"primaryKey;comment:岗位id"`
	Name                  string `gorm:"column:name;comment:岗位名称"`
	Code                  string `gorm:"column:code;comment:岗位代码"`
	Profile               string `gorm:"column:profile;comment:岗位简介"`
	JobDirection          string `gorm:"column:jobDirection;comment:岗位培养方向"`
	JobProspect           string `gorm:"column:jobProspect;comment:岗位就业前景"`
	SubjectCategory       string `gorm:"column:subjectCategory;comment:岗位学科门类"`
	FirstLevelDiscipline  string `gorm:"column:firstLevelDiscipline;comment:岗位一级学科"`
	SecondLevelDiscipline string `gorm:"column:secondLevelDiscipline;comment:岗位二级学课"`
	ScoreUrl              string `gorm:"column:scoreUrl;comment:考研分数线url"`
	ForeignType           string `gorm:"column:foreignType"`
	MathType              string `gorm:"column:mathType"`
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
	Author      User      `gorm:"foreignKey:AuthorID;"`
	Part        Part      `gorm:"foreignKey:PartID"`
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
	Author      User      `gorm:"foreignKey:UserID"`
}

type Feedback struct {
	ID      int64     `gorm:"column:ID;primaryKey"`
	UserID  string    `gorm:"column:userID"`
	Content string    `gorm:"column:content"`
	State   int64     `gorm:"column:state"`
	Time    time.Time `gorm:"column:time"`
	Author  User      `gorm:"foreignKey:UserID"`
}

type FavoriteRecord struct {
	ID        int64  `gorm:"column:ID;primaryKey"`
	UserID    string `gorm:"column:userID"`
	ArticleID string `gorm:"column:articleID"`
	Article   Post   `gorm:"foreignKey:articleID"`
	Author    User   `gorm:"foreignKey:UserID"`
}

type LikeRecord struct {
	ID     int64  `gorm:"column:ID;primaryKey"`
	UserId string `gorm:"column:userId"`
	PostId string `gorm:"column:postId"`
	Post   Post   `gorm:"foreignKey:postId"`
	Author User   `gorm:"foreignKey:UserId"`
}

type AdoptRecord struct {
	ID        int64  `gorm:"column:ID;primaryKey"`
	PostId    string `gorm:"column:postId;primaryKey"`
	CommentId string `gorm:"column:commentId;primaryKey"`
}
