// Code generated by goctl. DO NOT EDIT.
package types

type UserBasicInfo struct {
	ID        string `json:"userId"`    // 用户ID
	NickName  string `json:"nickName"`  // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan    string `json:"slogan"`    // 个性签名
	Gender    int64  `json:"gender"`    // 性别
}

type UserStatisticalInfo struct {
	FollowingCount        int `json:"followingCount"`        // 关注数
	FollowerCount         int `json:"followerCount"`         // 粉丝数
	LikeCount             int `json:"likeCount"`             // 获赞数量
	PublishedVideoCount   int `json:"publishVideoCount"`     // 发布作品数量
	LikeVideoCount        int `json:"likeVideoCount"`        // 点赞作品数量
	CollectionsVideoCount int `json:"collectionsVideoCount"` // 收藏作品数量
}

type UserInteractionInfo struct {
	IsFollow bool `json:"isFollow"` // 是否关注
}

type ListUserOption struct {
	PageSize int  `form:"pageSize,default=10"`   // 分页大小,默认为 10
	Page     int  `form:"page,default=1"`        // 当前页码,默认为 1
	ListAll  bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
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

type CurrentUserInfoReq struct {
}

type CurrentUserInfoResp struct {
	UserBasicInfo
	UserStatisticalInfo
	UserInteractionInfo
}

type UserDetailInfoReq struct {
	ContextUserId int64 `path:"userId"`
}

type UserDetailInfoResp struct {
	UserBasicInfo
	UserStatisticalInfo
	UserInteractionInfo
}

type UpdateAvatarReq struct {
	Image string `json:"image" binding:"Required"`
}

type UpdateAvatarResp struct {
}

type RankingReq struct {
	ListUserOption
}

type RankingResp struct {
	Users []*UserBasicInfo `json:"users"`
}

type UpdateInfoReq struct {
}

type UpdateInfoResp struct {
}

type UpdatePasswordReq struct {
}

type UpdatePasswordResp struct {
}

type FollowReq struct {
	ContextUserId int64 `path:"user_id"`
}

type FollowResp struct {
}

type UnFollowReq struct {
	ContextUserId int64 `json:"userId" path:"user_id" desc:"关注用户id" validate:"required"`
}

type UnFollowResp struct {
}

type FollowUser struct {
	UserBasicInfo
	UserInteractionInfo
}

type ListMyFollowersReq struct {
	ListUserOption
}

type ListMyFollowersResp struct {
	FollowUser []*FollowUser `json:"users" desc:"用户id"`
}

type ListFollowingReq struct {
	ContextUserId int64 `path:"user_id"`
	ListUserOption
}

type ListFollowingResp struct {
	FollowUser []*FollowUser `json:"users" desc:"用户id"`
}

type ListMyFollowingReq struct {
	ListUserOption
}

type ListMyFollowingResp struct {
	FollowUser []*FollowUser `json:"users" desc:"用户id"`
}

type ListFollowersReq struct {
	ContextUserId int64 `path:"user_id"`
	ListUserOption
}

type ListFollowersResp struct {
	FollowUser []*FollowUser `json:"users" desc:"用户id"`
}

type Category struct {
	ID   string `json:"id"`   // 分类ID
	Name string `json:"name"` // 分类名称
}

type Tag struct {
	Id   string `json:"id"`   // 标签id
	Name string `json:"name"` // 标签名称
}

type VideoUserInfo struct {
	ID        string `json:"userId"`    // 用户ID
	NickName  string `json:"nickName"`  // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan    string `json:"slogan"`    // 个性签名
	Gender    int64  `json:"gender"`    // 性别
	IsFollow  bool   `json:"isFollow"`  // 是否关注
}

type VideoBasicInfo struct {
	ID            string         `json:"id"`          // 视频ID
	Title         string         `json:"title"`       // 视频标题
	Description   string         `json:"description"` // 视频描述
	PlayUrl       string         `json:"playUrl"`     // 视频播放地址
	ThumbUrl      string         `json:"thumbUrl"`    // 视频封面地址
	CreatedAt     int64          `json:"createdAt"`   // 视频创建时间(毫秒时间戳)
	Category      *Category      `json:"category"`    // 视频分类
	Tags          []*Tag         `json:"tags"`        // 视频标签
	VideoUserInfo *VideoUserInfo `json:"author"`      // 视频作者信息
}

type VideoManageInfo struct {
	PublishTime   string `json:"publishTime"`   // 视频发布时间
	PublishStatus int    `json:"publishStatus"` // 视频发布状态
	Visibility    int    `json:"visibility"`    // 视频可见性
}

type VideoStatisticalInfo struct {
	FavoriteCount int64 `json:"favoriteCount"` // 点赞数
	CommentCount  int64 `json:"commentNum"`    // 评论数
	ShareCount    int64 `json:"shareNum"`      // 分享数
}

type VideoInteractInfo struct {
	IsFavorite bool `json:"isFavorite"` // 当前用户是否已点赞
}

type FeedReq struct {
	Cursor     int64  `form:"cursor,default=0"`    // 最新视频时间(毫秒时间戳)
	Limit      int    `form:"limit,default=10"`    // 请求数量
	AuthorID   int64  `form:"authorID,default=0"`  // 作者ID(是否根据用户ID过滤)
	Tag        string `form:"tag,optional"`        // 标签(是否根据标签过滤)
	CategoryID int64  `form:"categoryId,optional"` // 分类(是否根据分类过滤)
}

type FeedVideoItem struct {
	VideoBasicInfo
	VideoStatisticalInfo
	VideoInteractInfo
}

type FeedResp struct {
	Next  string           `json:"next"`  // 请求游标
	List  []*FeedVideoItem `json:"list"`  // 视频列表
	IsEnd bool             `json:"isEnd"` // 是否已到最后一页
}

type CategoryResp struct {
	CategoryList []*Category `json:"categoryList"`
}

type SearchReq struct {
	Keyword string `json:"keyword"` // 搜索关键字
	Offset  int64  `json:"offset"`  // 偏移量
	Limit   int64  `json:"limit"`   // 请求数量
}

type SearchResp struct {
	Hits               interface{} `json:"hits"`               // 搜索结果
	Query              string      `json:"query"`              // 搜索关键字
	ProcessingTimeMs   int64       `json:"processingTimeMs"`   // 搜索耗时(毫秒)
	Offset             int64       `json:"offset"`             // 偏移量
	Limit              int64       `json:"limit"`              // 请求数量
	EstimatedTotalHits int64       `json:"estimatedTotalHits"` // 搜索结果总数
}

type CreateVideoReq struct {
	Title       string   `json:"title"`                // 视频标题
	PlayUrl     string   `json:"playUrl"`              // 视频播放地址
	ThumbUrl    string   `json:"thumbUrl"`             // 视频封面地址
	Description string   `json:"description"`          // 视频描述
	Category    int64    `json:"category"`             // 视频分类
	Tags        []string `json:"tags"`                 // 视频标签
	PublishTime int64    `json:"publishTime,optional"` // 视频发布时间(毫秒时间戳)
	VideoKey    string   `json:"videoKey,optional"`    // 视频上传key
	Visibility  int      `json:"visibility"`           // 视频可见性(1:公开,2:私密)
}

type CreateVideoResp struct {
}

type DeleteVideoReq struct {
}

type DeleteVideoResp struct {
}

type FollowingReq struct {
	PageSize int  `form:"pageSize"` // 分页大小,默认为 10
	Page     int  `form:"page"`     // 当前页码,默认为 1
	ListAll  bool `form:"listAll"`  // 是否列出所有,默认为 false
}

type FollowingResp struct {
	VideoList   []*VideoBasicInfo `json:"videoList"`
	NextTime    int64             `json:"nextTime"`    // 下次请求时间(毫秒时间戳)
	CursorScore string            `json:"cursorScore"` // 下次请求时间(毫秒时间戳)
	Length      int               `json:"length"`      // 视频列表长度
}

type GetVideoInfoReq struct {
	VideoId int64 `path:"videoId"`
}

type GetVideoInfoResp struct {
	Video *VideoBasicInfo `json:"video"`
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
	VideoId int64  `json:"videoId" path:"video_id"`
}

type CreateVideoCommentResp struct {
	Commnent *Commnent `json:"comment"`
}

type GetVideoCommentReq struct {
	VideoId   int64 `json:"videoId" path:"video_id"`
	CommentId int64 `json:"commentId" path:"comment_id"`
}

type GetVideoCommentResp struct {
	Commnent *Commnent `json:"comment"`
}

type DeleteVideoCommentReq struct {
	VideoId   int64 `json:"videoId" path:"video_id"`
	CommentId int64 `json:"commentId" path:"comment_id"`
}

type DeleteVideoCommentResp struct {
}

type EditVideoCommentReq struct {
	VideoId   int64  `json:"videoId" path:"video_id"`
	CommentId int64  `json:"commentId" path:"comment_id"`
	Content   string `json:"content"`
}

type EditVideoCommentResp struct {
}

type ListVideoCommentsReq struct {
	VideoId int64 `json:"videoId" path:"video_id"`
}

type ListVideoCommentsResp struct {
	CreatedAt string      `json:"createdAt"`
	UserID    string      `json:"userId"`
	Comments  []*Commnent `json:"comments"`
}

type CreateReplyCommentReq struct {
	Content  string `json:"content"`
	VideoId  int64  `json:"videoId" path:"video_id"`
	ParentId int64  `json:"parentId"`
	TargetId int64  `json:"targetId"`
}

type CreateReplyCommentResp struct {
	Commnent *Commnent `json:"comment"`
}

type FavoriteVideoReq struct {
	VideoId    int64 `json:"videoId"`
	IsFavorite bool  `json:"isFavorite"` //true: favorite, false: unfavorite
}

type FavoriteVideoResp struct {
	IsFavorite bool `json:"favorite"`
}

type FavoriteCommentReq struct {
	VideoId    int64 `json:"videoId"`
	IsFavorite bool  `json:"isFavorite"` //true: favorite, false: unfavorite
}

type FavoriteCommentResp struct {
	IsFavorite bool `json:"favorite"`
}

type CheckVideoFavoriteReq struct {
	VideoId int64 `json:"videoId"`
}

type CheckVideoFavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}

type CheckCommentFavoriteReq struct {
	VideoId int64 `json:"videoId"`
}

type CheckCommentFavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}

type ListVideoFavoriteReq struct {
}

type ListVideoFavoriteResp struct {
}

type ListCommentFavoriteReq struct {
}

type ListCommentFavoriteResp struct {
}

type OssEndpointResponse struct {
	EndPoint string `json:"endpoint"`
}

type CreateUpTokenReq struct {
	UploadType string `form:"uploadType"` // 上传类型(video:视频,cover:封面,avatar:头像)
}

type CreateUpTokenResp struct {
	UpToken string `json:"upToken"` // 上传凭证
	Expires int64  `json:"expires"` // 上传凭证过期时间(秒)
}
