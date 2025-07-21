# ChatGPT Mock with Server-Sent Events

このプロジェクトは、GoのEchoフレームワークとNext.jsを使用して、Server-Sent Events（SSE）を活用したChatGPTのモックサイトです。

## 機能

- リアルタイムのチャットインターフェース
- Server-Sent Eventsによるストリーミングレスポンス
- ChatGPTライクなUI/UX
- 日本語対応
- タイピングアニメーション効果

## 技術スタック

### バックエンド
- **Go** - プログラミング言語
- **Echo** - Webフレームワーク
- **Server-Sent Events** - リアルタイム通信

### フロントエンド
- **Next.js 15** - Reactフレームワーク
- **TypeScript** - 型安全な開発
- **Tailwind CSS** - スタイリング
- **Lucide React** - アイコン

## セットアップ

### 前提条件
- Go 1.21以上
- Node.js 18以上
- npm

### 1. リポジトリのクローン
```bash
git clone <repository-url>
cd sse
```

### 2. バックエンド（Goサーバー）の起動

```bash
cd server
go mod tidy
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

### 3. フロントエンド（Next.js）の起動

新しいターミナルで：

```bash
cd frontend
npm install
npm run dev
```

フロントエンドは `http://localhost:3000` で起動します。

## API エンドポイント

### POST /api/chat
通常のチャットレスポンスを返します。

### POST /api/chat/stream
Server-Sent Eventsを使用してストリーミングレスポンスを返します。

## 使用方法

1. ブラウザで `http://localhost:3000` にアクセス
2. チャットボックスにメッセージを入力
3. 送信ボタンをクリックまたはEnterキーを押す
4. AIのレスポンスがリアルタイムで表示されます

## プロジェクト構造

```
sse/
├── server/          # Go Echo サーバー
│   ├── main.go
│   └── go.mod
├── frontend/        # Next.js フロントエンド
│   ├── src/
│   │   └── app/
│   │       └── page.tsx
│   ├── package.json
│   └── ...
└── README.md
```

## 開発

### バックエンドの開発
```bash
cd server
go run main.go
```

### フロントエンドの開発
```bash
cd frontend
npm run dev
```

## ライセンス

MIT License 