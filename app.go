package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx              context.Context
	selectedRepo     string
	localGitRepo     string
	claudeApiKey     string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowMaximise(ctx)
}

// SelectRepository opens a dialog to select a local repository
func (a *App) SelectRepository() (string, error) {
	selectedPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择本地代码仓库",
	})
	if err != nil {
		return "", err
	}
	
	a.selectedRepo = selectedPath
	return selectedPath, nil
}

// InitializeLocalGitRepo creates a local git repository for version control
func (a *App) InitializeLocalGitRepo() error {
	if a.selectedRepo == "" {
		return fmt.Errorf("请先选择代码仓库")
	}
	
	// Create a local git repository directory
	gitDir := filepath.Join(a.selectedRepo, ".gitclaude")
	if err := os.MkdirAll(gitDir, 0755); err != nil {
		return fmt.Errorf("创建本地Git目录失败: %v", err)
	}
	
	// Initialize git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = gitDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("初始化Git仓库失败: %v", err)
	}
	
	a.localGitRepo = gitDir
	return nil
}

// CommitChanges commits changes to the local git repository
func (a *App) CommitChanges(message string) error {
	if a.localGitRepo == "" {
		return fmt.Errorf("本地Git仓库未初始化")
	}
	
	// Copy current state to git repo
	err := a.copyRepoState()
	if err != nil {
		return fmt.Errorf("复制仓库状态失败: %v", err)
	}
	
	// Add all changes
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = a.localGitRepo
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("添加更改失败: %v", err)
	}
	
	// Commit changes
	commitMsg := fmt.Sprintf("Claude对话提交 - %s [%s]", message, time.Now().Format("2006-01-02 15:04:05"))
	cmd = exec.Command("git", "commit", "-m", commitMsg)
	cmd.Dir = a.localGitRepo
	if err := cmd.Run(); err != nil {
		// Check if there are no changes to commit
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return fmt.Errorf("没有更改需要提交")
		}
		return fmt.Errorf("提交失败: %v", err)
	}
	
	return nil
}

// copyRepoState copies the current repository state to git directory
func (a *App) copyRepoState() error {
	return filepath.Walk(a.selectedRepo, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip git directories and other irrelevant files
		if filepath.Base(path) == ".git" || filepath.Base(path) == ".gitclaude" {
			return filepath.SkipDir
		}
		
		relPath, err := filepath.Rel(a.selectedRepo, path)
		if err != nil {
			return err
		}
		
		destPath := filepath.Join(a.localGitRepo, relPath)
		
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}
		
		return a.copyFile(path, destPath)
	})
}

// copyFile copies a file from src to dst
func (a *App) copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	
	srcFile.Seek(0, 0)
	_, err = dstFile.ReadFrom(srcFile)
	return err
}

// GetCommitHistory returns the commit history
func (a *App) GetCommitHistory() ([]string, error) {
	if a.localGitRepo == "" {
		return nil, fmt.Errorf("本地Git仓库未初始化")
	}
	
	cmd := exec.Command("git", "log", "--oneline", "--max-count=20")
	cmd.Dir = a.localGitRepo
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("获取提交历史失败: %v", err)
	}
	
	// Parse output into slice
	lines := []string{}
	if len(output) > 0 {
		lines = []string{string(output)}
	}
	
	return lines, nil
}

// CallClaudeCode executes a claude code command in the selected repository
func (a *App) CallClaudeCode(prompt string) (string, error) {
	if a.selectedRepo == "" {
		return "", fmt.Errorf("请先选择代码仓库")
	}
	
	// Execute claude code command
	cmd := exec.Command("claude", "code", prompt)
	cmd.Dir = a.selectedRepo
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Claude Code执行失败: %v\n输出: %s", err, string(output))
	}
	
	// Auto-commit after claude code execution
	go func() {
		time.Sleep(2 * time.Second) // Wait for changes to be written
		commitMsg := fmt.Sprintf("Claude对话: %s", a.truncateString(prompt, 50))
		a.CommitChanges(commitMsg)
	}()
	
	return string(output), nil
}

// truncateString truncates a string to specified length
func (a *App) truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

// GetRepoFiles returns the file structure of the selected repository
func (a *App) GetRepoFiles() ([]string, error) {
	if a.selectedRepo == "" {
		return nil, fmt.Errorf("请先选择代码仓库")
	}
	
	var files []string
	err := filepath.Walk(a.selectedRepo, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip hidden directories and files
		if strings.HasPrefix(filepath.Base(path), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		
		if !info.IsDir() {
			relPath, err := filepath.Rel(a.selectedRepo, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}
		
		return nil
	})
	
	return files, err
}

// ReadFile reads a file from the selected repository
func (a *App) ReadFile(filename string) (string, error) {
	if a.selectedRepo == "" {
		return "", fmt.Errorf("请先选择代码仓库")
	}
	
	filePath := filepath.Join(a.selectedRepo, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %v", err)
	}
	
	return string(content), nil
}

// WriteFile writes content to a file in the selected repository
func (a *App) WriteFile(filename, content string) error {
	if a.selectedRepo == "" {
		return fmt.Errorf("请先选择代码仓库")
	}
	
	filePath := filepath.Join(a.selectedRepo, filename)
	
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}
	
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	
	return nil
}

// CheckGitRepoExists checks if .gitclaude directory exists
func (a *App) CheckGitRepoExists() (bool, error) {
	if a.selectedRepo == "" {
		return false, fmt.Errorf("请先选择代码仓库")
	}
	
	gitDir := filepath.Join(a.selectedRepo, ".gitclaude")
	if _, err := os.Stat(gitDir); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	
	// 检查是否是有效的git仓库
	gitConfigPath := filepath.Join(gitDir, ".git", "config")
	if _, err := os.Stat(gitConfigPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	
	a.localGitRepo = gitDir
	return true, nil
}

// GetSelectedRepo returns the currently selected repository path
func (a *App) GetSelectedRepo() string {
	return a.selectedRepo
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
