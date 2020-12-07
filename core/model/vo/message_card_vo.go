package vo

import "github.com/galaxy-book/feishu-sdk-golang/core/consts"

// 消息卡片数据类型定义

// 模块
type CardModuleContent struct {
	Tag string `json:"tag"`

	Text   *CardObjText   `json:"text"`
	Fields []CardObjField `json:"fields"`
	Extra  interface{}    `json:"extra,omitempty"`
}

type CardModuleHr struct {
	Tag string `json:"tag"`
}

type CardModuleImage struct {
	Tag string `json:"tag"`

	ImgKey string       `json:"img_key"`
	Alt    CardObjText  `json:"alt"`
	Title  *CardObjText `json:"title,omitempty"`
	Mode   string       `json:"mode"`
}

type CardModuleAction struct {
	Tag     string        `json:"tag"`
	Actions []interface{} `json:"actions"`
	Layout  string        `json:"layout,omitempty"`
}

type CardModuleNote struct {
	Tag string `json:"tag"`
	Elements []interface{} `json:"elements"` // text对象或者image元素
}

// 元素

type CardElementImage struct {
	Tag    string       `json:"tag"`
	ImgKey string       `json:"img_key"`
	Alt    *CardObjText `json:"alt,omitempty"`
}

type CardElementButton struct {
	Tag string `json:"tag"`

	Text     CardObjText            `json:"text"`
	Url      string                 `json:"url,omitempty"`
	MultiUrl *CardObjUrl            `json:"multi_url"`
	Type     string                 `json:"type,omitempty"`
	Value    map[string]interface{} `json:"value,omitempty"`
	Confirm  *CardObjConfirm        `json:"confirm,omitempty"`
}

type CardElementSelectMenu struct {
	Tag string `json:"tag"`

	Placeholder   *CardObjText           `json:"placeholder,omitempty"`
	InitialOption string                 `json:"initial_option,omitempty"`
	Options       []CardObjOption        `json:"options"`
	Value         map[string]interface{} `json:"value,omitempty"`
	Confirm       *CardObjConfirm        `json:"confirm,omitempty"`
}

type CardElementOverflow struct {
	Tag string `json:"tag"`

	Options []CardObjOption        `json:"options"`
	Value   map[string]interface{} `json:"value,omitempty"`
	Confirm *CardObjConfirm        `json:"confirm,omitempty"`
}

type CardElementDatePicker struct {
	Tag string `json:"tag"`

	InitialDate     string                 `json:"initial_date,omitempty"`
	InitialTime     string                 `json:"initial_time,omitempty"`
	InitialDatetime string                 `json:"initial_datetime,omitempty"`
	Placeholder     *CardObjText           `json:"placeholder,omitempty"`
	Value           map[string]interface{} `json:"value,omitempty"`
	Confirm         *CardObjConfirm        `json:"confirm,omitempty"`
}

// 对象

type CardObjText struct {
	Tag     string                    `json:"tag"`
	Content string                    `json:"content"`
	Lines   int                       `json:"lines,omitempty"`
	//Href    map[string]CardObjUrl `json:"href,omitempty"`
}

type CardObjField struct {
	IsShort bool        `json:"is_short"`
	Text    CardObjText `json:"text,omitempty"`
}

type CardObjUrl struct {
	Url        string `json:"url"`
	AndroidUrl string `json:"android_url"`
	IosUrl     string `json:"ios_url"`
	PcUrl      string `json:"pc_url"`
}

type CardObjOption struct {
	Text     *CardObjText `json:"text,omitempty"`
	Value    string       `json:"value"`
	Url      string       `json:"url,omitempty"`
	MultiUrl *CardObjUrl  `json:"multi_url,omitempty"`
}

type CardObjConfirm struct {
	Title *CardHeaderTitle `json:"title,omitempty"`
	Text  *CardObjText     `json:"text,omitempty"`
}

//机器人消息Card字段数据格式定义
type Card struct {
	Config       *CardConfig   `json:"config,omitempty"`
	CardLink     *CardObjUrl   `json:"card_link,omitempty"`
	Header       *CardHeader   `json:"header,omitempty"`
	Elements     []interface{} `json:"elements"`
	I18nElements *I18nElement  `json:"i18n_elements,omitempty"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type CardHeader struct {
	Title *CardHeaderTitle `json:"title,omitempty"`
}

type CardHeaderTitle struct {
	Tag     string    `json:"tag"`
	Content string    `json:"content"`
	Lines   int       `json:"lines,omitempty"`
	I18n    *CardI18n `json:"i18n,omitempty"`
}

type CardI18n struct {
	ZhCn string `json:"zh_cn"`
	EnUs string `json:"en_us"`
	JaJp string `json:"ja_jp"`
}

//type CardElementTextAlt struct {
//	Tag     string                `json:"tag"`
//	Content string                `json:"content"`
//	Lines   int                   `json:"lines,omitempty"`
//	Href    map[string]CardObjUrl `json:"href,omitempty"`
//}

type I18nElement struct {
	ZhCn []interface{} `json:"zh_cn"`
	EnUs []interface{} `json:"en_us"`
	JaJp []interface{} `json:"ja_jp"`
}

//type CardElementAction struct {
//	Tag  string       `json:"tag"`
//	Text *CardObjText `json:"text,omitempty"`
//	Type string       `json:"type"`
//}


func NewButton(btnType string,
	           content string,
	           url string,
	           multiUrl *CardObjUrl,
	           value map[string]interface{},
		       confirm *CardObjConfirm) *CardElementButton {
	return &CardElementButton{
		Tag: consts.TagBtn,
		Text: CardObjText{
			Tag: consts.TagLarkMd,
			Content: content,
		},
		Url: url,
		MultiUrl: multiUrl,
		Type: btnType,
		Value: value,
		Confirm: confirm,
	}
}

func NewConfirm(title string, text string) *CardObjConfirm{
	return &CardObjConfirm{
		Title: &CardHeaderTitle{
			Tag: consts.TagLarkMd,
			Content: title,
		},
		Text: &CardObjText{
			Tag: consts.TagLarkMd,
			Content: text,
		},
	}
}
