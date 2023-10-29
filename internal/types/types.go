// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	ID        int64  `json:"id"`
	Phone     string `json:"phone"`
	NickName  string `json:"nickName"`
	Sex       int64  `json:"sex"`
	AvatarUrl string `json:"avatarUrl"`
	Info      string `json:"info"`
}

type RegisterReq struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName,option"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type UpdateAvatarReq struct {
	Image string `json:"image" binding:"Required"`
}

type UpdateAvatarResp struct {
}

type FeedReq struct {
	LatestTime int64  `json:"latestTime,optional" form:"latestTime,optional"` // 最新视频时间(毫秒时间戳)
	Limit      int    `json:"limit,optional" form:"limit,optional"`           // 请求数量
	Token      string `json:"token,optional" form:"token,optional"`           // 登录token
	AuthorID   uint   `json:"authorId,optional" form:"authorID,optional"`     // 作者ID
	Tag        string `json:"tag,optional" form:"tag,optional"`               // 标签
	CategoryID uint   `json:"categoryId,optional" form:"categoryId,optional"` // 分类
}

type FeedResp struct {
	VideoList   []*Video     `json:"videoList"`
	Items       []*VideoItem `json:"items"`
	NextTime    int64        `json:"nextTime"`    // 下次请求时间(毫秒时间戳)
	CursorScore string       `json:"cursorScore"` // 下次请求时间(毫秒时间戳)
	Length      int          `json:"length"`      // 视频列表长度
}

type CategoryReq struct {
}

type CategoryResp struct {
	CategoryList []*Category `json:"categoryList"`
}

type VideoUserInfo struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	UserID   string `json:"userId"`
}

type VideoCoverInfo struct {
	ImageScene string `json:"imageScene"` //视频封面类型
	URL        string `json:"url"`        //视频封面地址
}

type VideoCover struct {
	TraceID  string           `json:"traceId"`
	InfoList []VideoCoverInfo `json:"infoList"`
	FileID   string           `json:"fileId"`
	Height   int              `json:"height"`
	Width    int              `json:"width"`
	URL      string           `json:"url"`
}

type VideoInteractInfo struct {
	Liked      bool   `json:"liked"`      // 当前用户是否已点赞
	LikedCount string `json:"likedCount"` // 点赞数
}

type VideoItem struct {
	Ignore    bool      `json:"ignore"`    // 是否忽略
	ID        int64     `json:"id"`        // 视频ID
	ModelType string    `json:"modelType"` // 发布类型:[video,article]
	VideoCard VideoCard `json:"videoCard"` // 视频信息
	TraceID   string    `json:"traceId"`   // 视频traceID
}

type VideoCard struct {
	Cover        VideoCover        `json:"cover"`        // 视频封面
	Type         string            `json:"type"`         // 视频分类
	DisplayTitle string            `json:"displayTitle"` // 视频标题
	User         VideoUserInfo     `json:"user"`         // 视频作者信息
	InteractInfo VideoInteractInfo `json:"interactInfo"` // 视频互动信息
}

type Video struct {
	ID            int64         `json:"id"`         // 视频ID
	Title         string        `json:"title"`      // 视频标题
	PlayUrl       string        `json:"playUrl"`    // 视频播放地址
	ThumbUrl      string        `json:"thumbUrl"`   // 视频封面地址
	FavoriteCount int64         `json:"favNum"`     // 点赞数
	CommentCount  int64         `json:"commentNum"` // 评论数
	ShareNum      int64         `json:"shareNum"`   // 分享数
	CreatedAt     string        `json:"createdAt"`  // 视频创建时间(毫秒时间戳)
	IsFav         bool          `json:"isFav"`      // 当前用户是否已点赞
	IsFollow      bool          `json:"isFollow"`   // 当前用户是否已关注该用户
	Tags          []string      `json:"tags"`       // 视频标签
	Author        VideoUserInfo `json:"author"`     // 作者信息
}

type Category struct {
	ID   uint   `json:"id"`   // 分类ID
	Name string `json:"name"` // 分类名称
}

type CreateUpTokenReq struct {
}

type CreateUpTokenResp struct {
	UpToken string `json:"upToken"` // 上传凭证
	Expires int64  `json:"expires"` // 上传凭证过期时间(秒)
}

type CountFollowReq struct {
	ContextUserId uint `json:"userId" path:"user_id" desc:"用户id" validate:"required"`
}

type CountFollowResp struct {
	FollowingCount int64 `json:"followingCount" desc:"关注数量"`
	FollowersCount int64 `json:"followerCount" desc:"粉丝数量"`
}

type FollowUser struct {
	UserId   string `json:"userId" desc:"用户id"`
	UserName string `json:"userName" desc:"用户名"`
	Avatar   string `json:"avatar" desc:"头像"`
}

type FollowReq struct {
	ContextUserId uint `json:"userId" path:"user_id" desc:"关注用户id" validate:"required"`
}

type FollowResp struct {
}

type UnFollowReq struct {
	ContextUserId uint `json:"userId" path:"user_id" desc:"关注用户id" validate:"required"`
}

type UnFollowResp struct {
}

type CheckMyFollowingReq struct {
	ContextUserId uint `json:"userId" path:"user_id" desc:"关注用户id" validate:"required"`
}

type CheckMyFollowingResp struct {
	Status bool `json:"status" desc:"是否关注"`
}

type ListMyFollowersReq struct {
	Page  int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}

type ListMyFollowersResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type ListMyFollowingReq struct {
	Page  int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}

type ListMyFollowingResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type ListFollowersReq struct {
	ContextUserId uint `json:"userId" path:"user_id" desc:"用户id" validate:"required"`
	Page          int  `json:"page" desc:"页码" validate:"required"`
	Limit         int  `json:"limit" desc:"每页数量" validate:"required"`
}

type ListFollowersResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type ListFollowingReq struct {
	ContextUserId uint `json:"userId" path:"user_id" desc:"用户id" validate:"required"`
	Page          int  `json:"page" desc:"页码" validate:"required"`
	Limit         int  `json:"limit" desc:"每页数量" validate:"required"`
}

type ListFollowingResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type CheckFollowingReq struct {
	DoerUserId    uint `json:"doerUserId" path:"doer_user_id" desc:"用户id" validate:"required"`
	ContextUserId uint `json:"contextUserId" path:"doer_user_id" desc:"用户id" validate:"required"`
}

type CheckFollowingResp struct {
	Status bool `json:"status" desc:"是否关注"`
}

type CommentUserInfo struct {
	UserID   string `json:"userId"`
	Nickname string `json:"nickname"`
	Image    string `json:"image"`
}

type SubComments struct {
	ID                  string               `json:"id"`
	VideoID             string               `json:"videoId"`
	Content             string               `json:"content"`
	CreateTime          int64                `json:"createTime"`
	Status              int                  `json:"status"`
	AtUsers             []*CommentUserInfo   `json:"atUsers"`
	UserInfo            *CommentUserInfo     `json:"userInfo"`
	CommentInteractInfo *CommentInteractInfo `json:"commentInteractInfo"`
	ShowTags            []string             `json:"showTags"`
}

type TargetCommnet struct {
	ID       string           `json:"id"`
	UserInfo *CommentUserInfo `json:"userInfo"`
}

type CommentInteractInfo struct {
	Liked      bool   `json:"liked"`      // 当前用户是否已点赞
	LikedCount string `json:"likedCount"` // 点赞数
}

type Commnent struct {
	ID                  string               `json:"id"`
	VideoID             string               `json:"videoId"`
	Status              int                  `json:"status"`
	Content             string               `json:"content"`
	CreateTime          int64                `json:"createTime"`
	UserInfo            *CommentUserInfo     `json:"userInfo"`
	CommentInteractInfo *CommentInteractInfo `json:"commentInteractInfo"`
	SubComments         []*SubComments       `json:"subComments"`
	SubCommentCount     string               `json:"subCommentCount"`
	SubCommentHasMore   bool                 `json:"subCommentHasMore"`
	SubCommentCursor    string               `json:"subCommentCursor"`
	AtUsers             []*CommentUserInfo   `json:"atUsers"`
	ShowTags            []string             `json:"showTags"`
	IPAddress           string               `json:"ipAddress"`
}

type CreateVideoCommentReq struct {
	Content string `json:"content"`
	VideoId uint   `json:"videoId" path:"video_id"`
}

type CreateVideoCommentResp struct {
	Commnent *Commnent `json:"comment"`
}

type GetVideoCommentReq struct {
	VideoId   uint `json:"videoId" path:"video_id"`
	CommentId uint `json:"commentId" path:"comment_id"`
}

type GetVideoCommentResp struct {
	Commnent *Commnent `json:"comment"`
}

type DeleteVideoCommentReq struct {
	VideoId   uint `json:"videoId" path:"video_id"`
	CommentId uint `json:"commentId" path:"comment_id"`
}

type DeleteVideoCommentResp struct {
}

type EditVideoCommentReq struct {
	VideoId   uint   `json:"videoId" path:"video_id"`
	CommentId uint   `json:"commentId" path:"comment_id"`
	Content   string `json:"content"`
}

type EditVideoCommentResp struct {
}

type ListVideoCommentsReq struct {
	VideoId uint `json:"videoId" path:"video_id"`
}

type ListVideoCommentsResp struct {
	CreatedAt string      `json:"createdAt"`
	UserID    string      `json:"userId"`
	Comments  []*Commnent `json:"comments"`
}

type CreateReplyCommentReq struct {
	Content  string `json:"content"`
	VideoId  uint   `json:"videoId" path:"video_id"`
	ParentId uint   `json:"parentId"`
	TargetId uint   `json:"targetId"`
}

type CreateReplyCommentResp struct {
	Commnent *Commnent `json:"comment"`
}
