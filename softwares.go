package main

type Software struct {
	Name string
	IPhone string
	IPad string
	Android string
	Windows64 string
	Windows32 string
	MacOsX string
}

var Softwares = map[string]*Software{
	"chrome" : &Software{
		Name: "Chrome",
		Windows64: "https://dl.google.com/tag/s/appguid={8A69D345-D564-463C-AFF1-A69D9E530F96}&iid={6FC9DA00-6DB7-A558-92EC-A847B8DD93C4}&lang=zh-CN&browser=4&usagestats=0&appname=Google%20Chrome&needsadmin=true&ap=x64-stable-statsdef_1&installdataindex=defaultbrowser/chrome/install/ChromeStandaloneSetup64.exe",
		Windows32: "https://dl.google.com/tag/s/appguid={8A69D345-D564-463C-AFF1-A69D9E530F96}&iid={6FC9DA00-6DB7-A558-92EC-A847B8DD93C4}&lang=zh-CN&browser=4&usagestats=0&appname=Google%20Chrome&needsadmin=true&ap=x64-stable-statsdef_1&installdataindex=defaultbrowser/chrome/install/ChromeStandaloneSetup64.exe",
		MacOsX: "https://dl.google.com/chrome/mac/stable/GGRO/googlechrome.dmg",
		IPhone: "https://itunes.apple.com/app/apple-store/id535886823",
		IPad: "https://itunes.apple.com/app/apple-store/id535886823",
		Android: "http://openbox.mobilem.360.cn/index/d/sid/21104",
	},
}