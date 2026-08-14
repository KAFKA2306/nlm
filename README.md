# nlm

**ブラウザでしかできない作業を、毎回手で繰り返す必要はない。**

NotebookLMでnotebookを作る、sourceを追加する、内容を確認する、audioやartifactを生成する――同じ操作を何度も行うなら、手順をコマンドとして再実行できる方が、何を実行したかも検証しやすくなります。実装ファイルが存在していても、command dispatcherへ登録されていない機能はCLIから利用できません。

nlmは、NotebookLMの操作をCLIから再現可能に扱うGoクライアントです。このrepositoryにはCLI本体に加えて、`KAFKA探究室` のPodcast配信物と、会話・開発・研究で実際に使った技術用語を学習用に整理する用語集を置いています。NotebookLMとの通信にはRPC / batchexecute実装を使い、実装済みでもdispatcher未登録の`nlm mcp`は現時点のsupported commandとして扱いません。

## Current boundaries

- Go module: `github.com/tmc/nlm`
- Go directive: `1.24.0`
- 認証情報: `~/.nlm/env` に `0600` で保存
- NotebookLM との通信はこのリポジトリ内の RPC / batchexecute 実装に依存するため、上流の挙動変更で壊れる可能性があります
- `cmd/nlm/mcp.go` には stdio MCP server 実装がありますが、現行 CLI の command dispatcher には `mcp` が登録されていません。したがって **`nlm mcp` は現時点の supported command ではありません**
- `nlm glossary ...` はNotebookLM認証を使わないローカルcommand familyです。標準ではrepository内の `data/glossary/terms.yaml` を探索し、別パスを使う場合は `NLM_GLOSSARY_PATH` を指定します

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

NotebookLM側のcommand surfaceはコードと `nlm help` を正準とします。学習用語集は認証前に処理するローカルcommand familyとして別に提供します。代表的なcommandは以下です。

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
| Glossary | `glossary list`, `glossary search`, `glossary show`, `glossary check`, `glossary generate` |
| Share | `share`, `share-private`, `share-details` |
| Auth / misc | `auth`, `refresh`, `feedback`, `hb` |

最小例:

```bash
nlm list
nlm create "research"
nlm add <notebook-id> <input>
nlm sources <notebook-id>
nlm chat <notebook-id>
nlm glossary search RLS
```

## 学習用語集

会話・開発・研究で**実際に使った技術用語**を拾い、学習用の正準データとして継続的に育てます。LLMがもっともらしい定義を生成して終わりにせず、一次情報・標準仕様・原論文で確認できた語だけを `verified` に昇格させます。

### Current snapshot — 2026-08-14

- 抽出監査: **294語**
- `verified`: **29語**
- `needs_review`: **265語**
- schema: [`data/glossary/schema.json`](data/glossary/schema.json)
- 正準データ: [`data/glossary/terms.yaml`](data/glossary/terms.yaml)
- 抽出監査inventory: [`data/glossary/recent-term-inventory.yaml`](data/glossary/recent-term-inventory.yaml)
- 人間向けビュー: [`docs/glossary.md`](docs/glossary.md)
- 追跡Issue: [#4 学習用の正準用語集を追加し、継続更新できる仕組みにする](https://github.com/KAFKA2306/nlm/issues/4)

現在の294語は、この作業で取得できた直近ChatGPT会話コンテキストとrecent-work retrievalから抽出した監査集合です。**ChatGPT全履歴の完全exportとは扱いません。**

### Promotion rule

用語は次の順で正準用語集へ昇格します。

1. 会話・Issue・PR・研究作業で実際に使われた語をinventoryへ追加する
2. aliasesとdomainを正規化する
3. 一次情報、標準仕様、原論文のいずれかで意味を確認する
4. 日本語の短い定義、なぜ重要か、具体例、関連語、source URL、確認日を付ける
5. 検証できたentryだけを `data/glossary/terms.yaml` の `verified` として保持する

正準entryの主な項目は以下です。

```yaml
id: row-level-security
term: Row Level Security
aliases: [RLS]
domain: database-security
definition_ja: ...
why_it_matters: ...
example: ...
related_terms: [...]
sources:
  - title: ...
    url: ...
    source_type: official_docs
verified_at: "2026-08-14"
status: verified
```

`needs_review` の語には、定義を推測して埋めません。現在は `RAG`, `Graphiti`, `PITR`, `fail-closed`, `walk-forward`, `Shader`, `Single Pass Instanced`, `SRP Batcher`, `UdonSharp`, `OIDC`, `ReAct`, `Whisper`, `Matter`, `Thread` などが検証待ちに含まれます。

### Glossary CLI

用語集commandはNotebookLMのcookie/tokenを必要としません。

```bash
nlm glossary list
nlm glossary search <query>
nlm glossary show <term|id|alias>
nlm glossary check
nlm glossary generate [output-path]
```

- `list`: 正準用語をterm/domain/status/idで一覧表示
- `search`: term、id、alias、domain、説明、例から部分一致検索
- `show`: term、id、aliasの完全一致で学習用詳細を表示
- `check`: strict YAML schema、必須項目、ID/term重複、status、一次情報URL、日付形式、`related_terms`参照切れを検査し、異常時は非0終了
- `generate`: `terms.yaml` から決定論的にMarkdownを生成。output-pathを省略した場合は `docs/glossary.md` を更新

標準パス以外の正準YAMLを検査する場合:

```bash
NLM_GLOSSARY_PATH=/path/to/terms.yaml nlm glossary check
```

**`docs/glossary.md` は生成物です。直接編集せず、必ず `nlm glossary generate` で再生成してください。** 正準データは `data/glossary/terms.yaml` です。

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
go run ./cmd/nlm glossary check
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
- 用語集の`verified` entryは、一次情報・標準仕様・原論文の根拠URLなしで追加しない
- `recent-term-inventory.yaml` は抽出監査、`terms.yaml` は検証済み正準データとして責務を分ける
- `docs/glossary.md` は手編集せず、`nlm glossary generate` で再生成する

## Podcast: KAFKA探究室

公開ページと RSS:

- https://kafka2306.github.io/nlm/
- https://kafka2306.github.io/nlm/feed.xml

GitHub Pages の公開物は `docs/index.html`, `docs/feed.xml`, `docs/artwork.png` にあります。Podcast の内容と CLI の仕様は別の責務として扱います。

## License

See [`LICENSE`](LICENSE).
