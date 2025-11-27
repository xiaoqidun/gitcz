package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CzType struct {
	Type    string
	Message string
}

type CzCommit struct {
	Type           *CzType
	Scope          *string
	Subject        *string
	Body           *string
	BreakingChange *string
	Closes         *string
}

var StdinInput = bufio.NewReader(os.Stdin)

var (
	InputTypePrompt           = "选择或输入一个提交类型(必填): "
	InputScopePrompt          = "说明本次提交的影响范围(必填): "
	InputSubjectPrompt        = "对本次提交进行简短描述(必填): "
	InputBodyPrompt           = "对本次提交进行完整描述(选填): "
	InputBreakingChangePrompt = "如果当前代码版本与上一版本不兼容, 对变动、变动的理由及迁移的方法进行描述(选填): "
	InputClosesPrompt         = "如果本次提交针对某个issue, 列出关闭的issues(选填): "
)

var CzTypeList = []CzType{
	{
		Type:    "feat",
		Message: "新的功能",
	},
	{
		Type:    "fix",
		Message: "修补错误",
	},
	{
		Type:    "docs",
		Message: "文档修改",
	},
	{
		Type:    "style",
		Message: "格式变化",
	},
	{
		Type:    "refactor",
		Message: "重构代码",
	},
	{
		Type:    "perf",
		Message: "性能提高",
	},
	{
		Type:    "test",
		Message: "测试用例",
	},
	{
		Type:    "chore",
		Message: "构建变动",
	},
}

func main() {
	amend := flag.Bool(
		"amend",
		false,
		"覆盖上次提交信息",
	)
	sign := flag.Bool("S", false, "对commit进行签名")
	author := flag.Bool(
		"author",
		false,
		"关于本软件开发者",
	)
	flag.Parse()
	if *author {
		Author()
		return
	}
	czCommit := &CzCommit{}
	czCommit.Type = InputType()
	czCommit.Scope = InputScope()
	czCommit.Subject = InputSubject()
	czCommit.Body = InputBody()
	czCommit.BreakingChange = InputBreakingChange()
	czCommit.Closes = InputCloses()
	commit := GenerateCommit(czCommit)
	if err := GitCommit(commit, *amend, *sign); err != nil {
		fmt.Println(err)
	}
}

func Author() {
	fmt.Println("welcome to our website https://aite.xyz/")
	fmt.Println("----------------------------------------")
	fmt.Println("腾讯扣扣：88966001")
	fmt.Println("电子邮箱：xiaoqidun@gmail.com")
	fmt.Println("----------------------------------------")
	fmt.Println("Copyright (c) 2020 xiaoqidun@gmail.com")
}

func NewLine() {
	fmt.Println()
}

func GitCommit(commit string, amend bool, sign bool) (err error) {
	tempFile, err := os.CreateTemp("", "git_commit_")
	if err != nil {
		return
	}
	defer func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}()
	if _, err = tempFile.WriteString(commit); err != nil {
		return
	}
	args := []string{"commit"}
	if amend {
		args = append(args, "--amend")
	}
	if sign {
		args = append(args, "-S")
	}
	args = append(args, "-F", tempFile.Name())
	cmd := exec.Command("git", args...)
	result, err := cmd.CombinedOutput()
	if err != nil && !strings.ContainsAny(err.Error(), "exit status") {
		return
	} else {
		fmt.Println(string(bytes.TrimSpace(result)))
	}
	return nil
}

func InputType() *CzType {
	typeNum := len(CzTypeList)
	for i := 0; i < typeNum; i++ {
		fmt.Printf("[%d] %s:\t%s\n", i+1, CzTypeList[i].Type, CzTypeList[i].Message)
	}
	fmt.Print(InputTypePrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	selectId, err := strconv.Atoi(text)
	if err == nil && (selectId > 0 && selectId <= typeNum) {
		NewLine()
		return &CzTypeList[selectId-1]
	}
	for i := 0; i < typeNum; i++ {
		if text == CzTypeList[i].Type {
			NewLine()
			return &CzTypeList[i]
		}
	}
	NewLine()
	return InputType()
}

func InputScope() *string {
	fmt.Print(InputScopePrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	NewLine()
	return InputScope()
}

func InputSubject() *string {
	fmt.Print(InputSubjectPrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	NewLine()
	return InputSubject()
}

func InputBody() *string {
	fmt.Print(InputBodyPrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	NewLine()
	return nil
}

func InputBreakingChange() *string {
	fmt.Print(InputBreakingChangePrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	NewLine()
	return nil
}

func InputCloses() *string {
	fmt.Print(InputClosesPrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	NewLine()
	return nil
}

func GenerateCommit(czCommit *CzCommit) string {
	commit := fmt.Sprintf(
		"%s(%s): %s\n\n",
		czCommit.Type.Type,
		*czCommit.Scope,
		*czCommit.Subject,
	)
	if czCommit.Body != nil {
		commit += *czCommit.Body
		commit += "\n\n"
	}
	if czCommit.BreakingChange != nil {
		commit += "BREAKING CHANGE: " + *czCommit.BreakingChange
		commit += "\n\n"
	}
	if czCommit.Closes != nil {
		commit += "Closes fix " + *czCommit.Closes
	}
	return commit
}
