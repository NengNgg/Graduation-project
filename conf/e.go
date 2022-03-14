/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 19:34
 */
package conf

var Msg = map[int]string{

	//default
	404: "未找到页面~",
	500: "未知错误~",
	528: "rss",
	633: "atom",

	400001000: "表单参数解析失败,请检查代码",
	400001001: "控制器参数断言失败,请检查后再试",
	400001002: "数据类型转换失败,请检查后再试",
	400001003: "表单参数获取失败,请检查后再试",
	400001004: "数据获取失败,请检查后再试",
	400001005: "权限不足或未登录,请检查后再试",

	//post
	401000000: "文章标题不能为空,请检查后再试",
	401000001: "文章分类不能为空,请检查后再试",
	401000002: "文章标签不能为空,请检查后再试",
	401000003: "文章摘要不能为空,请检查后再试",
	401000004: "表单参数获取失败,请检查后再试",
	401000005: "图片上传失败,请检查后再试",

	//cate
	402000000: "查询分类数据失败,请检查后再试",
	402000001: "获取分类数据列表失败,请检查后再试",
	402000002: "分类名不能为空,请检查后再试",
	402000003: "分类别名不能为空,请检查后再试",
	402000004: "分类父级不能为空,请检查后再试",
	402000005: "分类描述不能为空,请检查后再试",
	402000006: "分类名长度超出范围,请检查后再试",
	402000007: "分类别名长度超出范围,请检查后再试",
	402000008: "分类描述长度超出范围,请检查后再试",
	402000009: "分类修改失败,请检查后再试",
	402000010: "分类创建失败,请检查后再试",
	402000011: "分类还有子节点,请检查后再试",

	//tag
	403000000: "标签名不能为空,请检查后再试",
	403000001: "标签名超过最大长度,请检查后再试",
	403000002: "标签别名不能为空,请检查后再试",
	403000003: "标签别名超过最大长度,请检查后再试",
	403000004: "标签描述不能为空,请检查后再试",
	403000005: "标签描述超过最大长度,请检查后再试",
	403000006: "创建标签失败,请检查后再试",
	403000007: "标签修改失败,请检查后再试",
	403000008: "标签查询失败,请检查后再试",

	//system
	405000000: "系统信息更新失败,请检查后再试",
	405000001: "请输入网站title",
	405000002: "网站Title长度超出最大长度,请检查后再试",
	405000003: "网站关键词不能为空,请检查后再试",
	405000004: "网站关键词长度超出最大长度,请检查后再试",
	405000005: "网站描述不能为空,请检查后再试",
	405000006: "网站描述长度超出最大范围,请检查后再试",
	405000007: "网站备案号不能为空,请检查后再试",
	405000008: "网站备案号长度超出,请检查后再试",
	405000009: "请选择网站前台主题,请检查后再试",

	// link
	406000000: "友链别名不能为空,请检查后再试",
	406000001: "友链别名长度超出,请检查后再试",
	406000002: "友链链接不能为空,请检查后再试",
	406000003: "友链链接长度超出,请检查后再试",
	406000004: "友链排序号不正确,请检查后再试",
	406000005: "友链创建失败,请检查后再试",
	406000006: "友链信息查询失败,请检查后再试",
	406000007: "友链更新失败,请检查后再试",

	// auth
	407000000: "登录账号不能为空,请检查后再试",
	407000001: "登录账号类型不正确,请检查后再试",
	407000002: "密码不能为空,请检查后再试",
	407000003: "密码长度超出,请检查后再试",
	407000004: "验证码不能为空,请检查后再试",
	407000005: "验证码长度不正确,请检查后再试",
	407000006: "验证码key不存在,请检查后再试",
	407000007: "验证码key长度不正确,请检查后再试",
	407000008: "验证码不正确,请检查后再试",
	407000009: "用户数据获取失败,请检查后再试",
	407000010: "用户名账号或密码不正确,请检查后再试",
	407000011: "登录失败,请检查后再试",
	407000012: "用户名不能为空,请检查后再试",
	407000013: "用户名长度超出,请检查后再试",
	407000014: "操作失败,请检查后再试",
	407000015: "用户数量已超限",

	// index
	408000000: "常规信息查询失败,请检查后再试",
	408000001: "文章列表获取失败,请检查后再试",
	408000002: "文章详情获取失败,请检查后再试",
	408000003: "文章标签获取失败,请检查后再试",
	408000004: "文章标签获取失败,请检查后再试",
	408000005: "文章分类获取失败,请检查后再试",
	408000006: "归档列表获取失败,请检查后再试",
}
