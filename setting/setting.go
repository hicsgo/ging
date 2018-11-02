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
		DatabaseConfig DatabaseOption //数据库
		IsPro          bool           //是否生产环境
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 数据库选项
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	DatabaseOption struct {
		Connections []*DatabaseConnectionOption
		IsLog       bool
	}

	DatabaseConnectionOption struct {
		Key      string
		Username string
		Password string
		Host     string
		Database string
		Dialect  string
	}
)
