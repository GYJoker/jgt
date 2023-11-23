# YJ_GO_TOOL Go语言常用方法集

安装使用
```shell
go get github.com/GYJoker/jgt 
```


> 自己平时用到的方法合集，有什么问题欢迎大家指正修改

> 缓慢补充中...


## 内容介绍
### money-元和分的相互转换
- IntCentsToFloatYuan 整数分转浮点元
- FloatYuanToIntCents 浮点元转整数分

### times-时间相关方法集
- GetTimePointer 获取时间指针
- FormatTime 格式化时间  2006-01-02 15:04:05
- FormatSystemStrTime 格式化时间 将RFC3339格式的转换成2006-01-02 15:04:05
- ParseTime 解析时间  2006-01-02 15:04:05
- OffsetMinuteTime 偏移时间 -- 分钟
- OffsetHourTime 偏移时间 -- 小时
- OffsetDayTime 偏移时间 -- 天
- OffsetWeekTime 偏移时间 -- 周
- OffsetMonthTime 偏移时间 -- 月
- OffsetYearTime 偏移时间 -- 年
- BeginOfDay 一天开始时间
- EndOfDay 一天结束时间
- BeginOfWeek 一周开始时间
- EndOfWeek 一周结束时间
- BeginOfMonth 一个月开始时间
- EndOfMonth 一个月结束时间
- BeginOfYear 一年开始时间
- EndOfYear 一年结束时间

### common_func 常用工具函数
- StrIsEmpty 判断是否是空字符串
- StrIsNotEmpty 判断是否不是空字符串
- StrToUint64 将字符串转换为uint64
- GetMapKeys 获取map的key
- StrArrayToString 将字符串数组转换为字符串
- IntArrayToString 将整型数组转换为字符串
- FormatStrMoney 将数字转换为金额格式
- ArrContains 判断数组是否包含某个元素
- GetNickNameByPhone 获取昵称 将中间四位替换成*
- ValueToJsonStr 将任意类型的值转换为json字符串
- MaxInt64 获取最大值
- CalRateToStr 计算百分比 输出字符串
- StrVal 获取变量的字符串值 浮点型 3.0将会转换成字符串3, "3"  非数值或字符类型的变量将会被转换成JSON格式字符串
- ArrayInGroupsOf 将数组分割为多个数组
- ThreeWayOperator 三目运算
