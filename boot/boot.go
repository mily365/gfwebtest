package boot

import (
	_ "xpass/packed"

	_ "xpass/app/dao/driver"

	_ "xpass/app/dao/base"

	_ "xpass/app/dao"

	_ "xpass/app/service/base"

	_ "xpass/app/service"
)

func init() {
	//读取实体表类型的方案，构造配置对像需要的映射

}
