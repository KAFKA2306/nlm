# AGENTS.md

このファイルは automation / coding agent 向けの実行契約です。人間向けの説明は `README.md` に集約します。

## Priority of evidence

1. 実行結果・test・artifact
2. 実装 (`cmd/`, `internal/`, `proto/`, `gen/`)
3. build 定義 (`go.mod`, `Makefile`, workflow)
4. README / Issue / comment
5. 推測

下位の記述だけで上位の事実を上書きしません。証拠がなければ `UNVERIFIED` とします。

## Repository invariants

- module path は `github.com/tmc/nlm`
- Go directive は `1.24.0`
- CLI surface は `cmd/nlm/main.go` の `isValidCommand` / `runCmd` を正準とする
- 認証情報は secret として扱い、値を stdout / log / Issue / artifact に残さない
- `~/.nlm/env` は `0600` で保存される前提を壊さない
- generated code は `gen/` を直接直さず、`proto/` / template を直して `make generate` する
- `cmd/nlm/mcp.go` の存在だけで MCP を supported と判定しない。command routing と test が揃うまで `nlm mcp` は未提供
- `docs/` の HTML/XML/画像は公開 artifact。README の代替にしない

## Documentation contract

- 人間向けの正準入口は `README.md` 1本
- agent 固有ルールだけを `AGENTS.md` に置く
- 小さな例、debug 手順、test 手順のためだけに新しい `.md` を増やさない
- command、path、環境変数、version は実装と照合してから書く
- obsolete な説明は redirect 文書として残さず削除する
- README 自体を audit evidence と扱わない

## Change workflow

1. 対象実装と現在の repository state を読む
2. 外部仕様に依存する変更は公式一次情報を確認する
3. 最小差分で変更する
4. 可能なら以下を実行する

```bash
go test ./...
go build ./cmd/nlm
```

5. generated code を変えた場合は `make generate` 後の差分を確認する
6. 結果を `PASS` / `FAIL` / `UNVERIFIED` / `ASK_USER` で報告する

## Do not

- 実行していない test を `PASS` と書かない
- 存在しない command / path / env var を README に追加しない
- stale な snippet を別 `.md` に複製しない
- arbitrary な 80行 / 200行制限を repository contract にしない
- blanket な "crash-driven" ルールで必要な error handling を禁止しない
- secret を debug 目的で露出しない
- docs-only の変更を runtime correctness の証明として扱わない
