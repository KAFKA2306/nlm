# nlm

**ブラウザでしかできない作業を、毎回手で繰り返す必要はない。**

NotebookLMでnotebookを作る、sourceを追加する、内容を確認する、audioやartifactを生成する――同じ操作を何度も行うなら、手順をコマンドとして再実行できる方が、何を実行したかも検証しやすくなります。実装ファイルが存在していても、command dispatcherへ登録されていない機能はCLIから利用できません。

nlmは、NotebookLMの操作をCLIから再現可能に扱うGoクライアントです。このrepositoryにはCLI本体に加えて、`KAFKA探究室` のPodcast配信物も置かれています。NotebookLMとの通信にはRPC / batchexecute実装を使い、実装済みでもdispatcher未登録の`nlm mcp`は現時点のsupported commandとして扱いません。

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

`PASS`には対象commandまたはtestの実行証跡が必要です。

## Development

- 1変更1目的を基本にする
- generated code を手編集せず、`proto/` / template 側を変更して `make generate` する
- 認証情報・cookie・token を commit しない
- arbitrary な行数制限や、実装と矛盾する開発ルールを追加しない

## Podcast: KAFKA探究室

公開ページと RSS:

- https://kafka2306.github.io/nlm/
- https://kafka2306.github.io/nlm/feed.xml

GitHub Pages の公開物は `docs/index.html`, `docs/feed.xml`, `docs/artwork.png` にあります。Podcast の内容と CLI の仕様は別の責務として扱います。

## License

See [`LICENSE`](LICENSE).
