package capMonsterTool

//This file reference all CapMonster available Captcha Task Types
//documentation page: https://zennolab.atlassian.net/wiki/spaces/APIS/pages/557229/Captcha+Task+Types

type ImageToTextTask struct {
	//Defines the type of the task.
	Type string `json:"type"`
	//File body encoded in base64. Make sure to send it without line breaks.
	//Default value: ImageToTextTask
	Body string `json:"body"`
	//Name of recognizing module, for example, “yandex“. Alternative way to pass module name and list of all available modules look here
	//Default value: yandex
	CapMonsterModule string `json:"capMonsterModule"`
}

type NoCaptchaTaskProxyless struct {
	//value: NoCaptchaTaskProxyless
	Type string `json:"type"`
	//Address of a webpage with Google ReCaptcha
	WebsiteURL string `json:"websiteURL"`
	//Recaptcha website key.
	//<div class="g-recaptcha" data-sitekey="THAT_ONE"></div>
	WebsiteKey string `json:"websiteKey"`
	//not required
	//Some custom implementations may contain additional "data-s" parameter in ReCaptcha2 div, which is in fact a one-time token and must be grabbed every time you want to solve a ReCaptcha2.
	//<div class="g-recaptcha" data-sitekey="some sitekey" data-s="THIS_ONE"></div>
	RecaptchaDataSValue string `json:"recaptchaDataSValue"`
	//not required
	//Browser's User-Agent which is used in emulation. It is required that you use a signature of a modern browser, otherwise Google will ask you to "update your browser".
	UserAgent string `json:"userAgent"`
	//not required
	//Additional cookies which we must use during interaction with target page or Google.
	//Format: cookiename1=cookievalue1; cookiename2=cookievalue2
	Cookies string `json:"cookies"`
}

type NoCaptchaTask struct {
	//default value: NoCaptchaTask
	Type string `json:"type"`
	//Address of a webpage with Google ReCaptcha
	WebsiteURL string `json:"websiteURL"`
	//Recaptcha website key.
	//<div class="g-recaptcha" data-sitekey="THAT_ONE"></div>
	WebsiteKey string `json:"websiteKey"`
	//not required
	//Some custom implementations may contain additional "data-s" parameter in ReCaptcha2 div, which is in fact a one-time token and must be grabbed every time you want to solve a ReCaptcha2.
	//<div class="g-recaptcha" data-sitekey="some sitekey" data-s="THIS_ONE"></div>
	RecaptchaDataSValue string `json:"recaptchaDataSValue"`
	//Type of the proxy
	//
	//http - usual http/https proxy
	//https - try this only if "http" doesn't work (required by some custom proxy servers)
	//socks4 - socks4 proxy
	//socks5 - socks5 proxy
	ProxyType string `json:"proxyType"`
	//Proxy IP address IPv4/IPv6. Not allowed to use:
	//
	//host names instead of IPs
	//
	//transparent proxies (where client IP is visible)
	//
	//proxies from local networks (192.., 10.., 127...)
	ProxyAddress string `json:"proxyAddress"`
	//Proxy port
	ProxyPort string `json:"proxyPort"`
	//not required
	//Login for proxy which requires authorizaiton (basic)
	ProxyLogin string `json:"proxyLogin"`
	//not required
	//Proxy password
	ProxyPassword string `json:"proxyPassword"`
	//not required
	//Browser's User-Agent which is used in emulation. It is required that you use a signature of a modern browser, otherwise Google will ask you to "update your browser".
	UserAgent string `json:"userAgent"`
	//not required
	//Additional cookies which we must use during interaction with target page or Google.
	//
	//Format: cookiename1=cookievalue1; cookiename2=cookievalue2
	Cookies string `json:"cookies"`
}

type RecaptchaV3TaskProxyless struct {
	//default value: RecaptchaV3TaskProxyless
	Type string `json:"type"`
	//Address of a webpage with Google ReCaptcha
	WebsiteURL string `json:"websiteURL"`
	//Recaptcha website key.
	//https://www.google.com/recaptcha/api.js?render="THAT_ONE"
	WebsiteKey string `json:"websiteKey"`
	//Value from 0.1 to 0.9.
	MinScore string `json:"minScore"`
	//not required
	//Widget action value. Website owner defines what user is doing on the page through this parameter. Default value: verify
	//
	//Example:
	//grecaptcha.execute('site_key', {action:'login_test'}).
	PageAction string `json:"pageAction"`
}

type FunCaptchaTask struct {
	//default value: FunCaptchaTask
	Type string `json:"type"`
	//Address of a webpage with FunCaptcha
	WebsiteURL string `json:"websiteURL"`
	//not required
	//A special subdomain of funcaptcha.com, from which the JS captcha widget should be loaded.
	//Most FunCaptcha installations work from shared domains, so this option is only needed in certain rare cases.
	FuncaptchaApiJSSubdomain string `json:"funcaptchaApiJSSubdomain"`
	//FunCaptcha website key.
	//<div id="funcaptcha" data-pkey="THAT_ONE"></div>
	WebsitePublicKey string `json:"websitePublicKey"`
	//not required
	//Additional parameter that may be required by FunCaptcha implementation.
	//Use this property to send "blob" value as a stringified array. See example how it may look like.
	//{"\blob\":\"HERE_COMES_THE_blob_VALUE\"}
	Data string `json:"data"`
	//Type of the proxy
	//
	//http - usual http/https proxy
	//https - try this only if "http" doesn't work (required by some custom proxy servers)
	//socks4 - socks4 proxy
	//socks5 - socks5 proxy
	ProxyType string `json:"proxyType"`
	//Proxy IP address IPv4/IPv6. Not allowed to use:
	//
	//host names instead of IPs
	//
	//transparent proxies (where client IP is visible)
	//
	//proxies from local networks (192.., 10.., 127...)
	ProxyAddress string `json:"proxyAddress"`
	//Proxy port
	ProxyPort string `json:"proxyPort"`
	//not required
	//Login for proxy which requires authorization (basic)
	ProxyLogin string `json:"proxyLogin"`
	//not required
	//Proxy password
	ProxyPassword string `json:"proxyPassword"`
	//Browser's User-Agent which is used in emulation. It is required that you use a signature of a modern browser, otherwise Google will ask you to "update your browser".
	UserAgent string `json:"userAgent"`
	//not required
	//Additional cookies which we must use during interaction with target page or Google.
	//
	//Format: cookiename1=cookievalue1; cookiename2=cookievalue2
	Cookies string `json:"cookies"`
}

type FunCaptchaTaskProxyless struct {
	//default value: FunCaptchaTaskProxyless
	Type string `json:"type"`
	//Address of a webpage with FunCaptcha
	WebsiteURL string `json:"websiteURL"`
	//not required
	//A special subdomain of funcaptcha.com, from which the JS captcha widget should be loaded.
	//Most FunCaptcha installations work from shared domains, so this option is only needed in certain rare cases.
	FuncaptchaApiJSSubdomain string `json:"funcaptchaApiJSSubdomain"`
	//FunCaptcha website key.
	//<div id="funcaptcha" data-pkey="THAT_ONE"></div>
	WebsitePublicKey string `json:"websitePublicKey"`
	//not required
	//Additional parameter that may be required by Funcaptcha implementation.
	//Use this property to send "blob" value as a stringified array. See example how it may look like.
	//{"\blob\":\"HERE_COMES_THE_blob_VALUE\"}
	Data string `json:"data"`
}

type HCaptchaTask struct {
	//default value: HCaptchaTask
	Type string `json:"type"`
	//Address of a webpage with hCaptcha
	WebsiteURL string `json:"websiteURL"`
	//hCaptcha website key.
	WebsiteKey string `json:"websiteKey"`
	//Type of the proxy
	//
	//http - usual http/https proxy
	//https - try this only if "http" doesn't work (required by some custom proxy servers)
	//socks4 - socks4 proxy
	//socks5 - socks5 proxy
	ProxyType string `json:"proxyType"`
	//Proxy IP address IPv4/IPv6. Not allowed to use:
	//
	//host names instead of IPs
	//
	//transparent proxies (where client IP is visible)
	//
	//proxies from local networks (192.., 10.., 127...)
	ProxyAddress string `json:"proxyAddress"`
	//Proxy port
	ProxyPort string `json:"proxyPort"`
	//not required
	//Login for proxy which requires authorization (basic)
	ProxyLogin string `json:"proxyLogin"`
	//not required
	//Proxy password
	ProxyPassword string `json:"proxyPassword"`
	//not required
	//Browser's User-Agent which is used in emulation. It is required that you use a signature of a modern browser.
	UserAgent string `json:"userAgent"`
	//not required
	//Additional cookies which we must use during interaction with target page.
	//
	//Format: cookiename1=cookievalue1; cookiename2=cookievalue2
	Cookies string `json:"cookies"`
}

type HCaptchaTaskProxyless struct {
	//default value: HCaptchaTaskProxyless
	Type string
	//Address of a webpage with hCaptcha
	WebsiteURL string
	//hCaptcha website key.
	WebsiteKey string
	//not required
	//Browser's User-Agent which is used in emulation. It is required that you use a signature of a modern browser.
	UserAgent string
	//not required
	//Additional cookies which we must use during interaction with target page.
	//
	//Format: cookiename1=cookievalue1; cookiename2=cookievalue2
	Cookies string
}