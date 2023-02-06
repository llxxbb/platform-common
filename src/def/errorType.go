package def

type ErrorType string

const (
	ET_BIZ ErrorType = "BIZ" // 业务问题，请求参数验证失败
	ET_ENV ErrorType = "ENV" // 环境问题，如数据库无法连接
	ET_SYS ErrorType = "SYS" // 系统本身出现的问题，一般需要逻辑修正的。
	ET_COM ErrorType = "COM" // 具体错误要看结果中的子错误
)

var constStr = [...]string{"BIZ", "ENV", "SYS", "COM"}

func (et ErrorType) IsValid() bool {
	switch et {
	case ET_BIZ, ET_ENV, ET_SYS, ET_COM:
		return true
	}
	return false
}
