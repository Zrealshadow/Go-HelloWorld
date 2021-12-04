package errcode

var (
	ErrorGetTagListFail     = NewError(20010001, "获取标签列表失败")
	ErrorGetAccessTokenFail = NewError(20010002, "获取权限失败")
)
