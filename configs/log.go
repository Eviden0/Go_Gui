package configs

const (
	AddTodoItemErr = "添加待办事项[%s]失败--异常信息[%s]"
	GetTodoListErr = "获取待办事项列表失败--异常信息[%s]"
	DelTodoItemErr = "删除待办事项[%s]失败--异常信息[%s]"
)
const (
	AddCmdItemErr     = "添加快捷指令[%s]失败--异常信息[%s]"
	UpdateCmdItemErr  = "更新指令状态[%s]失败--异常信息[%s]"
	CmdItemHandlerErr = "[%s]指令执行失败--异常信息[%s]"
	DelCmdItemErr     = "删除快捷指令[%s]失败--异常信息[%s]"
)
const (
	GetConfigListErr    = "获取软件配置列表失败--异常信息[%s]"
	AddConfigItemErr    = "新增[%s]软件配置项失败--异常信息[%s]"
	GetConfigItemErr    = "查询[%s]软件配置项失败--异常信息[%s]"
	UpdateConfigItemErr = "更新[%s]软件配置项失败--异常信息[%s]"
	NoImgBedConfigErr   = "系统无图床配置"
)
const (
	InitSysClipboardErr      = "初始化剪贴板功能失败--异常信息[%s]"
	GetSystemClipboardImgErr = "获取系统剪贴板图片失败"
	SaveImgToLocalErr        = "图片本地存储失败"
)

const (
	CreateAliOSSClientErr = "创建阿里云OSS客户端失败--异常信息[%s]"
	GetAliOSSBucketErr    = "获取AliOSS存储空间失败--异常信息[%s]"
)
const (
	NoTransConfigErr   = "系统无翻译配置"
	BdOcrUrlParseErr   = "百度OCR url解析错误--异常信息[%s]"
	BdOcrTokenGetErr   = "百度OCR Token获取失败--异常信息[%s]"
	BdOcrRequestErr    = "百度OCR Request请求失败--异常信息[%s]"
	BdOcrResponseErr   = "百度OCR Response读取失败--异常信息[%s]"
	BdTransRequestErr  = "百度翻译请求失败--异常信息[%s]"
	TxTransRequestErr  = "腾讯翻译请求失败--异常信息[%s]"
	AliTransCreateErr  = "阿里翻译初始化失败--异常信息[%s]"
	AliTransRequestErr = "阿里翻译请求失败--异常信息[%s]"
)
