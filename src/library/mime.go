package library

/*
	<summary>
		Multipurpose Internet Mail Extensions
		Override default system types
	</summary>
*/
var MIMEs = map[string]string{
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":       ".xlsx",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
	"application/zip": ".zip",
	"video/mp4":       ".mp4",
	"image/webp":      ".webp",
}
