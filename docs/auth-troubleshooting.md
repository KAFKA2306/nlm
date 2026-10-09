# 認証トラブルシュート 🔐

認証情報の値（Cookie・トークン）は記載しない。管理ルールは [AGENTS.md](../AGENTS.md) を参照。

## 症状と原因

| 症状 | 原因 | 対処 |
|------|------|------|
| `401 Unauthorized` / `API error 16 (Authentication)` | 保存済みの Cookie が失効 | `make auth` で再ログイン |
| `auth login` が `auth data not found after timeout` | 対象プロファイルに notebook.google.com の Cookie がない | `nlm auth login -p "<プロファイル名>"` で Cookie を持つプロファイルを指定 |
| `NLM_COOKIES` が空のまま保存される | 抽出に失敗したまま env が書き出された | 保存内容を確認し、空なら再ログイン |
| `token not found or invalid` が繰り返される | コピーしたプロファイルでセッションが復号できていない | 下記「CDP 接続」を使う |
| `ERROR: could not unmarshal event: unknown PrivateNetworkRequestPolicy value` | 原因未確認。今回の調査では認証結果への影響は見られなかった | 認証結果だけで判断する |

## 既定プロファイルの罠

- 既定の `Default` プロファイルには notebook.google.com の Cookie が無い場合がある。
- `nlm auth login` はプロファイル一覧に `[✓ Has notebook.google.com cookies]` が出るものを選ぶ。
- `NLM_BROWSER_PROFILE` が `Default` に固定されていると、毎回失敗する。

## CDP 接続（推奨の代替手段）

プロファイルのコピーで失敗する場合は、ログイン済みの Chrome に接続する。

1. ログイン済みの Chrome を remote debugging 付きで起動する（既存の Chrome は終了してから）。
2. `./nlm auth -c ws://localhost:9222` を実行する。
3. 接続先が無い場合は `curl -s http://localhost:9222/json/version` で確認する。

## 資格情報の扱い（エージェント向け）

- `~/.nlm/env` はエージェントが作成・読み取り・コミットしない（AGENTS.md 準拠）。今回の調査で `nlm auth login` が平文ファイルを書き出したため、利用後は `make auth-clear` で破棄する。
- Chrome の Cookie DB を直接読む操作は、自動権限チェックで拒否された。ログインは人間が行い、資格情報マネージャへ登録する。

## 公式 API の代替

- Cookie を使わない方式は、NotebookLM Enterprise API（alpha）の audio overview 管理のみ確認できた。一覧 API は未確認。
- 出典: <https://docs.cloud.google.com/gemini/enterprise/notebooklm-enterprise/docs/api-audio-overview>

## 調査の記録（2026-10-09）

- `audio-list` は未検証。認証が通らないため、RSS 差分検出（feed の guid と audio-list の突き合わせ）は未実施。
- 次回の手順: 1) `make auth` で再ログイン 2) `scripts/secrets.sh exec ./nlm list` で疎通確認 3) `./nlm audio-list <notebook-id>` 4) feed.xml の guid と比較。
