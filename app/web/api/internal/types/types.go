// Code generated by goctl. DO NOT EDIT.
package types

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

type RankingReq struct {
	ListUserOption
}

type RankingResp struct {
	Users []*UserBasicInfo `json:"users"`
}

type UpdateInfoReq struct {
	NickName      string `json:"nickName"`      // 用户名
	Slogan        string `json:"slogan"`        // 个性签名
	Gender        int    `json:"gender"`        // 性别
	Age           int    `json:"age"`           // 年龄
	AvatarUrl     string `json:"avatarUrl"`     // 头像
	BackgroundUrl string `json:"backgroundUrl"` // 用户主页背景图
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
	Total      int64         `json:"total"`
}

type UserBasicInfo struct {
	ID            string `json:"userId" copier:"IDString"` // 用户ID
	NickName      string `json:"nickName"`                 // 用户名
	AvatarUrl     string `json:"avatarUrl"`                // 头像
	Slogan        string `json:"slogan"`                   // 个性签名
	Gender        int64  `json:"gender"`                   // 性别
	Age           int    `json:"age"`
	BackgroundUrl string `json:"backgroundUrl,optional"` //用户主页背景图
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

type FeedReq struct {
	Cursor     int64  `form:"cursor,default=0"`    // 最新视频时间(毫秒时间戳)
	Limit      int    `form:"limit,default=10"`    // 请求数量
	AuthorID   string `form:"authorID,default=0"`  // 作者ID(是否根据用户ID过滤)
	Tag        string `form:"tag,optional"`        // 标签(是否根据标签过滤)
	CategoryID string `form:"categoryId,optional"` // 分类(是否根据分类过滤)
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

type VideoSearchReq struct {
	Keyword string `json:"keyword"`         // 搜索关键字
	Offset  int64  `json:"offset,optional"` // 偏移量
	Limit   int64  `json:"limit,optional"`  // 请求数量
}

type VideoSearchResp struct {
	Videos             []*FeedVideoItem `json:"list"`               // 视频列表
	Query              string           `json:"query"`              // 搜索关键字
	ProcessingTimeMs   int64            `json:"processingTimeMs"`   // 搜索耗时(毫秒)
	Offset             int64            `json:"offset"`             // 偏移量
	Limit              int64            `json:"limit"`              // 请求数量
	EstimatedTotalHits int64            `json:"estimatedTotalHits"` // 搜索结果总数
}

type CreateVideoReq struct {
	Title         string   `json:"title"`                // 视频标题
	PlayUrl       string   `json:"playUrl"`              // 视频播放地址
	ThumbUrl      string   `json:"thumbUrl"`             // 视频封面地址
	Description   string   `json:"description,optional"` // 视频描述
	CategoryID    string   `json:"category"`             // 视频分类
	Tags          []string `json:"tags"`                 // 视频标签
	PublishTime   int64    `json:"publishTime,optional"` // 视频发布时间(毫秒时间戳)
	VideoKey      string   `json:"videoKey,optional"`    // 视频上传key
	Visibility    int      `json:"visibility"`           // 视频可见性(1:公开,2:私密)
	VideoDuration float32  `json:"videoDuration"`        // 视频时长(秒)
	VideoHeight   uint     `json:"videoHeight"`          // 视频高度(像素)
	VideoWidth    uint     `json:"videoWidth"`           // 视频宽度(像素)
}

type CreateVideoResp struct {
}

type DeleteVideoReq struct {
	VideoID int64 `path:"videoId"`
}

type DeleteVideoResp struct {
}

type GetVideoInfoReq struct {
	VideoId int64 `path:"videoId"`
}

type GetVideoInfoResp struct {
	Video *VideoBasicInfo `json:"video"`
}

type Category struct {
	ID   string `json:"id"`   // 分类ID
	Name string `json:"name"` // 分类名称
}

type Tag struct {
	Id   string `json:"id" copier:"IDString"` // 标签id
	Name string `json:"name"`                 // 标签名称
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
	ID            string         `json:"id" copier:"IDString"` // 视频ID
	Title         string         `json:"title"`                // 视频标题
	Description   string         `json:"description"`          // 视频描述
	PlayUrl       string         `json:"playUrl"`              // 视频播放地址
	ThumbUrl      string         `json:"thumbUrl"`             // 视频封面地址
	CreatedAt     int64          `json:"createdAt"`            // 视频创建时间(毫秒时间戳)
	Category      *Category      `json:"category"`             // 视频分类
	Tags          []*Tag         `json:"tags"`                 // 视频标签
	VideoUserInfo *VideoUserInfo `json:"author"`               // 视频作者信息
	VideoHeight   float32        `json:"videoHeight"`          // 视频高度
	VideoWidth    float32        `json:"videoWidth"`           // 视频宽度
	FavoriteCount int64          `json:"favoriteCount"`        // 点赞数
	CommentCount  int64          `json:"commentNum"`           // 评论数
	VideoDuration float32        `json:"videoDuration"`        // 视频时长         //视频宽度
}

type VideoManageInfo struct {
	PublishTime   string `json:"publishTime"`   // 视频发布时间
	PublishStatus int    `json:"publishStatus"` // 视频发布状态
	Visibility    int    `json:"visibility"`    // 视频可见性
}

type VideoStatisticalInfo struct {
	ShareCount int64 `json:"shareNum"` // 分享数
}

type VideoInteractInfo struct {
	IsFavorite bool `json:"isFavorite"` // 当前用户是否已点赞
}

type CreateVideoCommentReq struct {
	VideoId   string   `json:"videoId"`
	Content   string   `json:"content"`
	AtUsersId []string `json:"atUsersId,optional"`
}

type CreateVideoCommentResp struct {
	Comment *ParentComment `json:"comment"`
}

type CreateChildCommentReq struct {
	VideoId         string   `json:"videoId"`
	Content         string   `json:"content"`
	AtUsersId       []string `json:"atUsersId,optional"`
	ParentCommentId string   `json:"parentCommentId"`
	TargetCommentId string   `json:"targetCommentId,optional"`
}

type CreateChildCommentResp struct {
	Comment *ChildComment `json:"comment"`
}

type CreateReplyCommentReq struct {
	VideoId         string   `json:"videoId"`
	Content         string   `json:"content"`
	AtUsersId       []string `json:"atUsersId,optional"`
	ParentCommentId string   `json:"parentCommentId,optional"`
	TargetCommentId string   `json:"targetCommentId,optional"`
}

type CreateReplyCommentResp struct {
	Comment *ChildComment `json:"comment"`
}

type GetVideoCommentReq struct {
	CommentId string `json:"commentId" path:"comment_id"`
}

type GetVideoCommentResp struct {
	Commnent *ParentComment `json:"comment"`
}

type DeleteVideoCommentReq struct {
	CommentId string `path:"comment_id"`
	Type      string `form:"type"`
}

type DeleteVideoCommentResp struct {
}

type EditVideoCommentReq struct {
	VideoId   string `json:"videoId" path:"video_id"`
	CommentId string `json:"commentId" path:"comment_id"`
	Content   string `json:"content"`
}

type EditVideoCommentResp struct {
}

type ListVideoCommentsReq struct {
	VideoId string `path:"video_id,optional"`
}

type ListVideoCommentsResp struct {
	Comments []*ParentComment `json:"comments"`
}

type ListCommentOption struct {
	PageSize int  `form:"pageSize,default=10"`   // 分页大小,默认为 10
	Page     int  `form:"page,default=1"`        // 当前页码,默认为 1
	ListAll  bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type TargetComment struct {
	ID       string         `json:"id" copier:"IDString"` // 回复的目标评论ID
	UserInfo *UserBasicInfo `json:"userInfo"`             // 回复的目标评论用户信息
}

type CommentBasicInfo struct {
	ID         string           `json:"id" copier:"IDString"`                   // 评论ID
	Content    string           `json:"content"`                                // 评论内容
	AtUsers    []*UserBasicInfo `json:"atUsers"`                                // @用户列表(暂未实现)
	UserInfo   *UserBasicInfo   `json:"userInfo"`                               // 发布评论的用户信息
	ShowTags   []string         `json:"showTags"`                               // 标签列表(暂未实现)
	LikedCount int64            `json:"likedCount"`                             // 点赞数
	Liked      bool             `json:"liked"`                                  // 当前用户是否已点赞
	CreatedAt  int64            `json:"createTime" copier:"CreatedAtUnixMilli"` // 创建时间(毫秒时间戳)
	Status     int              `json:"status"`
}

type ParentComment struct {
	CommentBasicInfo
	VideoID       string          `json:"videoId"`       // 视频ID
	ChildComments []*ChildComment `json:"childComments"` // 二级评论列表
	ChildCount    int64           `json:"childCount"`    // 二级评论数
	ChildHasMore  bool            `json:"childHasMore"`  // 是否还有更多二级评论
}

type ChildComment struct {
	CommentBasicInfo
	TargetComment *TargetComment `json:"targetComment"` // 回复的目标评论
}

type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}

type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
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

type NoticeItem struct {
	NoticeId          string `json:"noticeId"`          // 通知ID
	NoticeType        string `json:"noticeType"`        // 通知类型
	NoticeTime        int64  `json:"noticeTime"`        // 通知时间
	FromUserID        string `json:"fromUserId"`        // 用户ID
	FromUserNickName  string `json:"fromUserNickName"`  // 用户名
	FromUserAvatarUrl string `json:"fromUserAvatarUrl"` // 头像
}

type FollowNoticeReq struct {
	Cursor     int64  `form:"cursor,default=0"` // 最新通知时间(毫秒时间戳)
	Limit      int    `form:"limit,default=10"` // 请求数量
	NoticeType string `form:"noticeType"`       // 通知类型
}

type FollowNoticeResp struct {
	Next  string        `json:"next"`  // 请求游标
	List  []*NoticeItem `json:"list"`  // 通知列表
	IsEnd bool          `json:"isEnd"` // 是否已到最后一页
}
