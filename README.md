# nlm

NotebookLM を CLI から操作する Go クライアントです。このリポジトリには CLI 本体に加えて、`KAFKA探究室` の Podcast 配信物も置かれています。

## Repository contract

README は入口であり、実装の代わりではありません。事実確認は次の順で行います。

1. CLI の実挙動: `cmd/nlm/main.go`
2. 認証: `cmd/nlm/auth.go`, `internal/auth/`
3. API / schema: `proto/`, `gen/`, `internal/notebooklm/`
4. build / test: `go.mod`, `Makefile`, `*_test.go`
5. 公開配信物: `docs/`

README と実装が食い違う場合は、実装と実行結果を優先します。未実行のものを `PASS` と扱いません。

## Current boundaries

- Go module: `github.com/tmc/nlm`
- Go directive: `1.24.0`
- 認証情報: `~/.nlm/env` に `0600` で保存
- NotebookLM との通信はこのリポジトリ内の RPC / batchexecute 実装に依存するため、上流の挙動変更で壊れる可能性があります
- `cmd/nlm/mcp.go` には stdio MCP server 実装がありますが、現行 CLI の command dispatcher には `mcp` が登録されていません。したがって **`nlm mcp` は現時点の supported command ではありません**

## Install

このリポジトリの内容をそのまま使う場合は source から install します。

```bash
git clone https://github.com/KAFKA2306/nlm.git
cd nlm
go version
make install
```

`go.mod` は Go `1.24.0` を要求します。

## Authentication

通常は browser login を使います。

```bash
nlm auth login
```

profile を指定する場合:

```bash
nlm auth login -profile Work
```

利用可能な profile を探索する場合:

```bash
nlm auth login -all -notebooks
```

既存 CDP session を使う場合:

```bash
nlm auth -cdp-url ws://localhost:9222
```

認証後は `NLM_COOKIES`, `NLM_AUTH_TOKEN`, `NLM_BROWSER_PROFILE` が `~/.nlm/env` に保存されます。値そのものを log / issue / artifact に残さないでください。

## CLI

完全な command surface はコードと `nlm help` を正準とします。代表的な command は以下です。

| Area | Commands |
|---|---|
| Notebook | `list`, `create`, `rm`, `analytics`, `list-featured` |
| Source | `sources`, `add`, `rm-source`, `rename-source`, `refresh-source`, `check-source`, `discover-sources` |
| Note | `notes`, `new-note`, `update-note`, `rm-note` |
| Audio | `audio-list`, `audio-create`, `audio-get`, `audio-download`, `audio-rm`, `audio-share` |
| Video | `video-list`, `video-create`, `video-download` |
| Artifact | `artifacts`, `create-artifact`, `get-artifact`, `rename-artifact`, `delete-artifact` |
| Generate | `generate-guide`, `generate-outline`, `generate-section`, `generate-chat`, `generate-magic`, `chat` |
| Transform | `summarize`, `rephrase`, `expand`, `critique`, `verify`, `explain`, `study-guide`, `faq`, `briefing-doc`, `mindmap`, `timeline`, `toc` |
| Share | `share`, `share-private`, `share-details` |
| Auth / misc | `auth`, `refresh`, `feedback`, `hb` |

最小例:

```bash
nlm list
nlm create "research"
nlm add <notebook-id> <input>
nlm sources <notebook-id>
nlm chat <notebook-id>
```

## Debug

Debug は global flag で有効化します。

```bash
nlm -debug list
```

実装上の主な debug flag は `-debug`, `-debug-dump-payload`, `-debug-parsing`, `-debug-field-mapping` です。旧文書にあった `NLM_DEBUG=true` は現行 CLI の正準設定ではありません。

`HTTPRR_RECORDING_DIR` は現行 `main.go` で検出されますが、recording client への接続は未実装です。record/replay が動くものとして運用しません。

## Build and verification

```bash
make build
make test
```

同等の最小確認:

```bash
go build ./cmd/nlm
go test ./...
```

protobuf / generated client を更新した場合:

```bash
make generate
```

完了判定は次の4値で扱います。

- `PASS`: 対象 command / test を実行し成功を確認した
- `FAIL`: 実行して失敗した
- `UNVERIFIED`: 実行環境または証跡がなく未確認
- `ASK_USER`: 外部判断が必要で機械的に決められない

README、コメント、Issue、LLM 出力だけを `PASS` の根拠にしません。

## Development

- 1変更1目的を基本にする
- user-facing command を変えたら `README.md` と test を同じ変更で更新する
- 同じ説明を複数の `.md` に複製しない
- generated code を手編集せず、`proto/` / template 側を変更して `make generate` する
- 認証情報・cookie・token を commit しない
- arbitrary な行数制限や、実装と矛盾する開発ルールを追加しない

人間向けドキュメントの正準入口はこの `README.md`、agent 用の実行契約は `AGENTS.md` です。

## Podcast: KAFKA探究室

公開ページと RSS:

- https://kafka2306.github.io/nlm/
- https://kafka2306.github.io/nlm/feed.xml

GitHub Pages の公開物は `docs/index.html`, `docs/feed.xml`, `docs/artwork.png` にあります。Podcast の内容と CLI の仕様は別の責務として扱います。

## License

See [`LICENSE`](LICENSE).
