package types

// CustomError struct
type CustomError struct {
	ErrorCode    string
	ErrorMessage string
}

// SimpleMessage struct
type SimpleMessage struct {
	GroupId       string
	MessageId     string
	AccountId     string
	StatusMessage string
	StatusCode    string
	To            string
	From          string
	Type          string
	Country       string
}

// Count struct
type Count struct {
	Total             int
	SentTotal         int
	SentFailed        int
	SentSuccess       int
	SentPending       int
	SentReplacement   int
	Refund            int
	RegisteredFailed  int
	RegisteredSuccess int
}

// CountForCharge struct
type CountForCharge struct {
	SMS map[string]int
	LMS map[string]int
	MMS map[string]int
	ATA map[string]int
	CTA map[string]int
	CTI map[string]int
}

// Balance struct for Group struct
type BalanceForGroup struct {
	Requested   float32
	Replacement float32
	Refund      float32
	Sum         float32
}

// Point struct for Group struct
type Point struct {
	Requested   int
	Replacement int
	Refund      int
	Sum         int
}

// Profit struct for App struct
type Profit struct {
	SMS float32
	LMS float32
	MMS float32
	ATA float32
	CTA float32
	CTI float32
}

// App struct for Group struct
type App struct {
	Profit  Profit
	AppId   string
	Version string
}

// Price struct
type Price struct {
	SMS float32
	LMS float32
	MMS float32
	ATA float32
	CTA float32
	CTI float32
}

// Group struct
type Group struct {
	Count          Count
	CountForCharge CountForCharge
	Balance        BalanceForGroup
	Point          Point
	App            App
	SdkVersion     string
	OsPlatform     string
	Log            []map[string]interface{}
	Status         string
	DateSent       string
	DateCompleted  string
	IsRefunded     bool
	FlagUpdated    bool
	AccountId      string
	ApiVersion     string
	GroupId        string
	Price          map[string]Price
	DateCreated    string
	DateUpdated    string
	Id             string `json:"_id"`
}

// GroupList struct
type GroupList struct {
	StartKey  string
	Limit     int
	NextKey   string
	GroupList map[string]Group
}

// GroupMessageList struct
type GroupMessageList struct {
	StartKey    string
	Limit       int
	MessageList MessageList
}

// KakaoOptions struct
type KakaoOptions struct {
	SenderKey    string
	TemplateCode string
	ButtonName   string
	ButtonUrl    string
	PfId         string
	TemplateId   string
	ImageId      string
	DisableSms   bool
	Buttons      []map[string]string
}

// Message struct
type Message struct {
	KakaoOptions   KakaoOptions
	Type           string
	Country        string
	Subject        string
	ImageId        string
	DateProcessed  string
	DateReported   string
	DateReceived   string
	StatusCode     string
	NetworkCode    string
	Log            []map[string]interface{}
	Replacement    bool
	AutoTypeDetect bool
	RoutedQueue    string
	MessageId      string
	GroupId        string
	AccountId      string
	Text           string
	From           string
	To             string
	CustomFields   map[string]string
	DateCreated    string
	DateUpdated    string
	Reason         string
	NetworkName    string
}

// MessageList struct
type MessageList struct {
	// Data
	StartKey    string
	Limit       int
	MessageList map[string]Message
}

// AddGroupMessage struct
type AddGroupMessage struct {
	To            string
	From          string
	Type          string
	Country       string
	MessageId     string
	StatusCode    string
	StatusMessage string
	AccountId     string
}

// AddGroupMessageList struct
type AddGroupMessageList struct {
	ErrorCount int
	ResultList []AddGroupMessage
}

// LowBalanceAlert struct
type LowBalanceAlert struct {
	NotificationBalance string
	CurrentBalance      string
	Balances            []int
	Channels            []string
	Enabled             bool
}

// Balance struct
type Balance struct {
	LowBalanceAlert  LowBalanceAlert
	Point            int
	MinimumCash      int
	RechargeTo       int
	RechargeTryCount int
	AutoRecharge     int
	AccountId        string
	Balance          int
}

// File struct
type File struct {
	Kakao        map[string]string
	Type         string
	OriginalName string
	Link         string
	FileId       string
	Name         string
	Url          string
	AccountId    string
	References   []string
	DateCreated  string
	DateUpdated  string
}

// FileList struct
type FileList struct {
	FileList []File
}
