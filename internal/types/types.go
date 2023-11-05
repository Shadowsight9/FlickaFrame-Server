// Code generated by goctl. DO NOT EDIT.
package types

type UserBasicInfo struct {
	ID        string `json:"userId"`    // 用户ID
	NickName  string `json:"nickName"`  // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan    string `json:"slogan"`    // 个性签名
	Gender    int64  `json:"gender"`    // 性别
	Age       int    `json:"age"`
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
	ContextUserId int64 `path:"userId,optional"`
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
	NickName string `json:"nickName,optional"` // 用户名
	Slogan   string `json:"slogan,optional"`   // 个性签名
	Gender   int    `json:"gender,optional"`   // 性别
	Age      int    `json:"age,optional"`
	Password string `json:"password,optional"`
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
	ContextUserId int64 `path:"user_id"`
}

type UnFollowResp struct {
}

type FollowUser struct {
	UserBasicInfo
	UserInteractionInfo
}

type ListFollowReq struct {
	ContextUserId int64 `path:"user_id,optional"`
	ListUserOption
}

type ListFollowUserResp struct {
	FollowUser []*FollowUser `json:"users"`
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

type ListCommentOption struct {
	PageSize int  `form:"pageSize,default=10"`   // 分页大小,默认为 10
	Page     int  `form:"page,default=1"`        // 当前页码,默认为 1
	ListAll  bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type CommonTag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TargetComment struct {
	ID       string         `json:"id"`       // 回复的目标评论ID
	UserInfo *UserBasicInfo `json:"userInfo"` // 回复的目标评论用户信息
}

type CommentBasicInfo struct {
	ID         string           `json:"id"`         // 评论ID
	Content    string           `json:"content"`    // 评论内容
	AtUsers    []*UserBasicInfo `json:"atUsers"`    // @用户列表(暂未实现)
	UserInfo   *UserBasicInfo   `json:"userInfo"`   // 发布评论的用户信息
	ShowTags   []*CommonTag     `json:"showTags"`   // 标签列表(暂未实现)
	LikedCount int64            `json:"likedCount"` // 点赞数
	Liked      bool             `json:"liked"`      // 当前用户是否已点赞
	CreatedAt  int64            `json:"createTime"` // 创建时间(毫秒时间戳)
	Status     int              `json:"status"`
}

type ParentComment struct {
	CommentBasicInfo
	VideoID       string          `json:"videoId"`       // 视频ID
	ChildComments []*ChildComment `json:"childComments"` // 二级评论列表
	ChildCount    string          `json:"childCount"`    // 二级评论数
	ChildHasMore  bool            `json:"childHasMore"`  // 是否还有更多二级评论
}

type ChildComment struct {
	CommentBasicInfo
	TargetComment *TargetComment `json:"targetComment"` // 回复的目标评论
}

type CreateVideoCommentReq struct {
	VideoId   int64   `json:"videoId"`
	Content   string  `json:"content"`
	AtUsersId []int64 `json:"atUsersId,optional"`
}

type CreateVideoCommentResp struct {
	Comment *ParentComment `json:"comment"`
}

type CreateChildCommentReq struct {
	VideoId         int64   `json:"videoId"`
	Content         string  `json:"content"`
	AtUsersId       []int64 `json:"atUsersId,optional"`
	ParentCommentId int64   `json:"parentCommentId"`
	TargetCommentId int64   `json:"targetCommentId,optional"`
}

type CreateChildCommentResp struct {
	Comment *ChildComment `json:"comment"`
}

type CreateReplyCommentReq struct {
	VideoId         int64   `json:"videoId"`
	Content         string  `json:"content"`
	AtUsersId       []int64 `json:"atUsersId,optional"`
	ParentCommentId int64   `json:"parentCommentId,optional"`
	TargetCommentId int64   `json:"targetCommentId,optional"`
}

type CreateReplyCommentResp struct {
	Comment *ChildComment `json:"comment"`
}

type GetVideoCommentReq struct {
	CommentId int64 `json:"commentId" path:"comment_id"`
}

type GetVideoCommentResp struct {
	Commnent *ParentComment `json:"comment"`
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
	VideoId int64 `path:"video_id,optional"`
}

type ListVideoCommentsResp struct {
	Comments []*ParentComment `json:"comments"`
}

type FavoriteReq struct {
	TargetId int64 `json:"targetId"`
}

type FavoriteResp struct {
}

type UnFavoriteReq struct {
	TargetId int64 `json:"targetId"`
}

type UnFavoriteResp struct {
}

type CheckFavoriteReq struct {
	TargetId int64 `path:"targetId"`
}

type CheckFavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}

type ListFavoriteReq struct {
}

type ListFavoriteResp struct {
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
