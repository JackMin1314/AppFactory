package fileopr

const (
	FILE_CSV  = "csv"
	FILE_XLSX = "xlsx"
	FILE_TXT  = "txt"
	FILE_LOG  = "log"
)

/* 关于gf gen dao生成命令说明   gf gen dao -c .\configs\config.toml --path ./internal/
1. dao/internal 以及model/internal 下面的文件由工具生成，多次生成会被覆盖，因此不要手动修改。采用internal包名的目的是仅作为dao或model的内部包引用，不对外开放。
2. dao 目录下的文件 可以做一些数据库的定制化操作，通过工具多次生成不会覆盖，但是更多建议用户在自己的service中实现。
3. model目录下的文件，可以做自定义的一些数据结构定义，通过工具多次生成不会覆盖。

*/
