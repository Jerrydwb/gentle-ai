package qwen

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/internal/model"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

var LookPathOverride = exec.LookPath

type statResult struct {
	isDir bool
	err   error
}

type Adapter struct {
	lookPath func(string) (string, error)
	statPath func(string) statResult
}

func NewAdapter() *Adapter {
	return &Adapter{
		lookPath: LookPathOverride,
		statPath: defaultStat,
	}
}

// --- Identity ---

func (a *Adapter) Agent() model.AgentID {
	return model.AgentQwenCode
}

func (a *Adapter) Tier() model.SupportTier {
	return model.TierFull
}

// --- Detection ---

func (a *Adapter) Detect(_ context.Context, homeDir string) (bool, string, string, bool, error) {
	configPath := filepath.Join(homeDir, ".qwen")

	binaryPath, err := a.lookPath("qwen")
	installed := err == nil

	stat := a.statPath(configPath)
	if stat.err != nil {
		if os.IsNotExist(stat.err) {
			return installed, binaryPath, configPath, false, nil
		}
		return false, "", "", false, stat.err
	}

	return installed, binaryPath, configPath, stat.isDir, nil
}

// --- Installation ---

func (a *Adapter) SupportsAutoInstall() bool {
	return true
}

func (a *Adapter) InstallCommand(profile system.PlatformProfile) ([][]string, error) {
	// Prefer the official install script on Linux/macOS.
	// Fall back to npm on Windows or when brew is not available.
	switch profile.OS {
	case "darwin":
		if profile.PackageManager == "brew" {
			return [][]string{{"brew", "install", "qwen-code"}}, nil
		}
		return [][]string{{"npm", "install", "-g", "@qwen-code/qwen-code@latest"}}, nil
	case "linux":
		if profile.PackageManager == "brew" {
			return [][]string{{"brew", "install", "qwen-code"}}, nil
		}
		if !profile.NpmWritable {
			return [][]string{{"sudo", "npm", "install", "-g", "@qwen-code/qwen-code@latest"}}, nil
		}
		return [][]string{{"npm", "install", "-g", "@qwen-code/qwen-code@latest"}}, nil
	default:
		return [][]string{{"npm", "install", "-g", "@qwen-code/qwen-code@latest"}}, nil
	}
}

// --- Config paths ---

func (a *Adapter) GlobalConfigDir(homeDir string) string {
	return filepath.Join(homeDir, ".qwen")
}

func (a *Adapter) SystemPromptDir(homeDir string) string {
	return filepath.Join(homeDir, ".qwen")
}

func (a *Adapter) SystemPromptFile(homeDir string) string {
	return filepath.Join(homeDir, ".qwen", "QWEN.md")
}

func (a *Adapter) SkillsDir(homeDir string) string {
	return filepath.Join(homeDir, ".qwen", "skills")
}

func (a *Adapter) SettingsPath(homeDir string) string {
	return filepath.Join(homeDir, ".qwen", "settings.json")
}

// --- Config strategies ---

func (a *Adapter) SystemPromptStrategy() model.SystemPromptStrategy {
	return model.StrategyFileReplace
}

func (a *Adapter) MCPStrategy() model.MCPStrategy {
	return model.StrategyMergeIntoSettings
}

// --- MCP ---

func (a *Adapter) MCPConfigPath(homeDir string, _ string) string {
	return filepath.Join(homeDir, ".qwen", "settings.json")
}

// --- Optional capabilities ---

func (a *Adapter) SupportsOutputStyles() bool {
	return false
}

func (a *Adapter) OutputStyleDir(_ string) string {
	return ""
}

func (a *Adapter) SupportsSlashCommands() bool {
	return true
}

func (a *Adapter) CommandsDir(homeDir string) string {
	return filepath.Join(homeDir, ".qwen", "commands")
}

func (a *Adapter) EmbeddedCommandsDir() string {
	return "qwen/commands"
}

func (a *Adapter) SupportsSkills() bool {
	return true
}

func (a *Adapter) SupportsSystemPrompt() bool {
	return true
}

func (a *Adapter) SupportsMCP() bool {
	return true
}

// --- Sub-agent support (Qwen native agents in ~/.qwen/agents/) ---
// Qwen Code sub-agents are a native stable feature — no experimental flag required.

func (a *Adapter) SupportsSubAgents() bool {
	return true
}

func (a *Adapter) SubAgentsDir(homeDir string) string {
	return filepath.Join(homeDir, ".qwen", "agents")
}

func (a *Adapter) EmbeddedSubAgentsDir() string {
	return "qwen/agents"
}

func defaultStat(path string) statResult {
	info, err := os.Stat(path)
	if err != nil {
		return statResult{err: err}
	}

	return statResult{isDir: info.IsDir()}
}
