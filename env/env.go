package env

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func LoadEnv(filename string) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 忽略空行和注释行
		if strings.TrimSpace(line) == "" || line[0] == '#' {
			continue
		}

		// 按 "=" 分割每行
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		// 设置环境变量
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}

	// 检查读取过程中是否出错
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func StoInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func ParseTDuration(s string) time.Duration {
	time, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return time
}

func ParseSList(s string) []string {
	return strings.Split(s, ",")
}
