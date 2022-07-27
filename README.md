# GoLang Lumberjack Logger ✍️
golang lumberjack logger example

---
[![shields](https://img.shields.io/badge/made%20with-go-blue?logo=go&style=for-the-badge&logoColor=white)](https://golang.org) ![shields](https://img.shields.io/badge/License-GPL-green.svg?logo=read-the-docs&style=for-the-badge&logoColor=white)
[![socialbadge](https://img.shields.io/twitter/follow/yakutozcan.svg?style=social)](https://twitter.com/yakutozcan)

---

```  
log.SetOutput(&lumberjack.Logger{
		Filename:   "log/logfile.log",
		MaxSize:    1, // megabytes
		MaxBackups: 10,
		MaxAge:     90,   //days
		Compress:   true, // disabled by default
	})
	log.SetFlags(log.Lshortfile | log.LstdFlags)
```

### Screenshot 

![goLumberjackLogger](https://github.com/yakutozcan/goLumberjackLogger/blob/main/screen.png?raw=true)
