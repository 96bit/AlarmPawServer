package modal

type Params struct {
	DeviceKey   string `json:"deviceKey"`
	DeviceToken string `json:"deviceToken"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Category    string `json:"category"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
	Url         string `json:"url"`
	IsArchive   string `json:"isArchive"`
	Group       string `json:"group"`
	Sound       string `json:"sound"`
	AutoCopy    string `json:"autoCopy"`
	Copy        string `json:"copy"`
	Badge       string `json:"badge"`
	Level       string `json:"level"`
	CipherText  string `json:"cipherText"`
}

// ios notification sound(system sound please refer to http://iphonedevwiki.net/index.php/AudioServices)

type Message struct {
	DeviceToken string                 `form:"-" json:"-" xml:"-" query:"-"`
	DeviceKey   string                 `form:"device_key,omitempty" json:"device_key,omitempty" xml:"device_key,omitempty" query:"device_key,omitempty"`
	Category    string                 `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty" query:"category,omitempty"`
	Title       string                 `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty" query:"title,omitempty"`
	Body        string                 `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty" query:"body,omitempty"`
	Sound       string                 `form:"sound,omitempty" json:"sound,omitempty" xml:"sound,omitempty" query:"sound,omitempty"`
	ExtParams   map[string]interface{} `form:"ext_params,omitempty" json:"ext_params,omitempty" xml:"ext_params,omitempty" query:"ext_params,omitempty"`
}

type DeviceInfo struct {
	DeviceKey   string `form:"device_key,omitempty" json:"device_key,omitempty" xml:"device_key,omitempty" query:"device_key,omitempty"`
	DeviceToken string `form:"device_token,omitempty" json:"device_token,omitempty" xml:"device_token,omitempty" query:"device_token,omitempty"`
}
