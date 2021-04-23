package model

// 定义爬虫获取信息的结构体
type Profile struct {
	// 价格
	Price   		string
	// 厂商
	Factory			string
	// 级别
	Level   		string
	// 发动机
	Motor			string
	// 动力类型
	PowerType		string
	// 上市年份
	Year			string
	// 最高车速
	Speed			string
	// 0-100加速时间
	Acceleration	string
	// 油耗
	Oil				string
}
