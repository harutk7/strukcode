<p align="center">
  <a href="https://strukcode.ai">
    <picture>
      <source srcset="packages/console/app/src/asset/logo-ornate-dark.svg" media="(prefers-color-scheme: dark)">
      <source srcset="packages/console/app/src/asset/logo-ornate-light.svg" media="(prefers-color-scheme: light)">
      <img src="packages/console/app/src/asset/logo-ornate-light.svg" alt="StrukCode logo">
    </picture>
  </a>
</p>
<p align="center">開源的 AI Coding Agent。</p>
<p align="center">
  <a href="https://strukcode.ai/discord"><img alt="Discord" src="https://img.shields.io/discord/1391832426048651334?style=flat-square&label=discord" /></a>
  <a href="https://www.npmjs.com/package/strukcode-ai"><img alt="npm" src="https://img.shields.io/npm/v/strukcode-ai?style=flat-square" /></a>
  <a href="https://github.com/harutk7/strukcode/actions/workflows/publish.yml"><img alt="Build status" src="https://img.shields.io/github/actions/workflow/status/harutk7/strukcode/publish.yml?style=flat-square&branch=dev" /></a>
</p>

[![StrukCode Terminal UI](packages/web/src/assets/lander/screenshot.png)](https://strukcode.ai)

---

### 安裝

```bash
# 直接安裝 (YOLO)
curl -fsSL https://strukcode.ai/install | bash

# 套件管理員
npm i -g strukcode-ai@latest        # 也可使用 bun/pnpm/yarn
scoop bucket add extras; scoop install extras/strukcode  # Windows
choco install strukcode             # Windows
brew install strukcode              # macOS 與 Linux
paru -S strukcode-bin               # Arch Linux
mise use -g github:harutk7/strukcode    # 任何作業系統
nix run nixpkgs#strukcode           # 或使用 github:harutk7/strukcode 以取得最新開發分支
```

> [!TIP]
> 安裝前請先移除 0.1.x 以前的舊版本。

### 桌面應用程式 (BETA)

StrukCode 也提供桌面版應用程式。您可以直接從 [發佈頁面 (releases page)](https://github.com/harutk7/strukcode/releases) 或 [strukcode.ai/download](https://strukcode.ai/download) 下載。

| 平台                  | 下載連結                              |
| --------------------- | ------------------------------------- |
| macOS (Apple Silicon) | `strukcode-desktop-darwin-aarch64.dmg` |
| macOS (Intel)         | `strukcode-desktop-darwin-x64.dmg`     |
| Windows               | `strukcode-desktop-windows-x64.exe`    |
| Linux                 | `.deb`, `.rpm`, 或 AppImage           |

```bash
# macOS (Homebrew Cask)
brew install --cask strukcode-desktop
```

#### 安裝目錄

安裝腳本會依據以下優先順序決定安裝路徑：

1. `$STRUKCODE_INSTALL_DIR` - 自定義安裝目錄
2. `$XDG_BIN_DIR` - 符合 XDG 基礎目錄規範的路徑
3. `$HOME/bin` - 標準使用者執行檔目錄 (若存在或可建立)
4. `$HOME/.strukcode/bin` - 預設備用路徑

```bash
# 範例
STRUKCODE_INSTALL_DIR=/usr/local/bin curl -fsSL https://strukcode.ai/install | bash
XDG_BIN_DIR=$HOME/.local/bin curl -fsSL https://strukcode.ai/install | bash
```

### Agents

StrukCode 內建了兩種 Agent，您可以使用 `Tab` 鍵快速切換。

- **build** - 預設模式，具備完整權限的 Agent，適用於開發工作。
- **plan** - 唯讀模式，適用於程式碼分析與探索。
  - 預設禁止修改檔案。
  - 執行 bash 指令前會詢問權限。
  - 非常適合用來探索陌生的程式碼庫或規劃變更。

此外，StrukCode 還包含一個 **general** 子 Agent，用於處理複雜搜尋與多步驟任務。此 Agent 供系統內部使用，亦可透過在訊息中輸入 `@general` 來呼叫。

了解更多關於 [Agents](https://strukcode.ai/docs/agents) 的資訊。

### 線上文件

關於如何設定 StrukCode 的詳細資訊，請參閱我們的 [**官方文件**](https://strukcode.ai/docs)。

### 參與貢獻

如果您有興趣參與 StrukCode 的開發，請在提交 Pull Request 前先閱讀我們的 [貢獻指南 (Contributing Docs)](./CONTRIBUTING.md)。

### 基於 StrukCode 進行開發

如果您正在開發與 StrukCode 相關的專案，並在名稱中使用了 "strukcode"（例如 "strukcode-dashboard" 或 "strukcode-mobile"），請在您的 README 中加入聲明，說明該專案並非由 StrukCode 團隊開發，且與我們沒有任何隸屬關係。

### 常見問題 (FAQ)

#### 這跟 Claude Code 有什麼不同？

在功能面上與 Claude Code 非常相似。以下是關鍵差異：

- 100% 開源。
- 不綁定特定的服務提供商。雖然我們推薦使用透過 [StrukCode Zen](https://strukcode.ai/zen) 提供的模型，但 StrukCode 也可搭配 Claude, OpenAI, Google 甚至本地模型使用。隨著模型不斷演進，彼此間的差距會縮小且價格會下降，因此具備「不限廠商 (provider-agnostic)」的特性至關重要。
- 內建 LSP (語言伺服器協定) 支援。
- 專注於終端機介面 (TUI)。StrukCode 由 Neovim 愛好者與 [terminal.shop](https://terminal.shop) 的創作者打造；我們將不斷挑戰終端機介面的極限。
- 客戶端/伺服器架構 (Client/Server Architecture)。這讓 StrukCode 能夠在您的電腦上運行的同時，由行動裝置進行遠端操控。這意味著 TUI 前端只是眾多可能的客戶端之一。

#### 另一個同名的 Repo 是什麼？

另一個名稱相近的儲存庫與本專案無關。您可以點此[閱讀背後的故事](https://x.com/thdxr/status/1933561254481666466)。

---

**加入我們的社群** [Discord](https://discord.gg/strukcode) | [X.com](https://x.com/strukcode)
