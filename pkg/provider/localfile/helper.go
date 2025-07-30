package localfile

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// readAndFilter is a function that reads a file and filters its lines based on the given keywords.
func readAndFilter(filePath string, keywords []string) ([]string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 将 keywords 转换为正则表达式模式
	var keywordPatterns []string
	for _, keyword := range keywords {
		// 对每个关键字进行转义，避免特殊字符干扰正则表达式
		keywordPatterns = append(keywordPatterns, regexp.QuoteMeta(keyword))
	}

	// 构造正则表达式：匹配任意一个关键字
	pattern := strings.Join(keywordPatterns, "|")
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	// 读取文件并过滤行
	var filteredLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 如果当前行匹配任意一个关键字，则保留该行
		if re.MatchString(line) {
			filteredLines = append(filteredLines, line)
		}
	}

	// 检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return filteredLines, nil
}
