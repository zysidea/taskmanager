package common

func StartUp() {
	//初始化配置文件
	initConfig()
	//初始化公钥私钥
	initKeys()
	//获取MongoDB的session
	createDbSession()
	//添加索引
	addIndexes()
}
