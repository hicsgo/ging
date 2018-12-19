package setting

/* ================================================================================
 * 设置数据域结构
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
     * 全局设置数据模型
     * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	Setting struct {
		DatabaseConfig *DatabaseConfig //数据库
		IsPro          bool            //是否生产环境
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 数据库选项
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	DatabaseConfig struct {
		DatabaseOptions []*DatabaseOption
	}

	DatabaseOption struct {
		ProjectName  string
		ReadDBConns  []*DatabaseConnectionOption
		WirteDBConns []*DatabaseConnectionOption
	}

	DatabaseConnectionOption struct {
		Key      string //用来查找当前配置项
		Username string //登录名
		Password string //密码
		Host     string //host
		Database string //数据库名称
		Dialect  string //数据库类型
		IsLog    bool   //是否记录日志
		Weight   int    //权重
	}

	//达观数据项
	DataGrand struct {
		UrlPrefix string //地址前缀eg:http://datareportapi.datagrand.com/data/
		AppId     string //appid
		AppName   string //your_app_name
	}
)
