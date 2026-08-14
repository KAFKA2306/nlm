# 学習用語集

このページは、直近の会話で実際に使われた技術用語を学習用に整理した人間向けビューです。

- 正準データ: [`data/glossary/terms.yaml`](../data/glossary/terms.yaml)
- 抽出監査: [`data/glossary/recent-term-inventory.yaml`](../data/glossary/recent-term-inventory.yaml) — 294語（verified 29 / needs_review 265）
- 確認日: 2026-08-14
- `verified` は一次情報または標準仕様で定義を確認した項目だけに付けます。
- 定義・数値・仕様を確認できない語は、この一覧では `verified` として追加しません。
- 抽出監査は、このタスクで取得できた直近ChatGPT会話コンテキストとrecent-work retrievalを対象にしています。ChatGPT全履歴の完全exportとは扱いません。
- このMarkdownを直接の正準データにはしません。Issue #4 で正準YAMLからの自動生成を実装します。

## Git / GitHub / CI

| 用語 | 短い説明 | 一次情報 |
|---|---|---|
| GitHub Actions | GitHub上のイベントやスケジュールを契機に自動処理を実行する仕組み。 | [GitHub Docs](https://docs.github.com/en/actions/concepts) |
| Workflow | GitHub Actionsで1つ以上のjobを実行する自動プロセス。 | [GitHub Docs](https://docs.github.com/en/actions/concepts/workflows-and-actions/workflows) |
| Runner | GitHub Actionsのjobを実際に実行するマシンまたは実行環境。 | [GitHub Docs](https://docs.github.com/en/actions/concepts/runners) |
| Continuous Integration (CI) | 変更を継続的に統合し、build・testして不具合を早く検出する実践。 | [GitHub Docs](https://docs.github.com/en/actions/get-started/continuous-integration) |
| Pull Request (PR) | GitHub上で変更を提案し、議論・review・mergeする単位。 | [GitHub Docs](https://docs.github.com/en/pull-requests) |
| Branch | Gitにおける独立した開発ライン。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Remote-tracking branch | 別repositoryのbranch状態を追跡する参照。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Merge | 別branchの内容を取り込む操作。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Squash and merge | PR内の複数commitを1commitへまとめてmergeする方式。 | [GitHub Docs](https://docs.github.com/en/pull-requests/reference/pull-request-merges) |
| Draft Pull Request | まだmerge可能な完成状態ではないことを示すPR。 | [GitHub Docs](https://docs.github.com/en/pull-requests/how-tos/merge-and-close-pull-requests/merging-a-pull-request) |
| GitHub Pages | GitHub repositoryからWebサイトを公開するホスティング機能。 | [GitHub Docs](https://docs.github.com/en/pages) |
| Workflow artifact | workflow runで生成され、job完了後も保存・共有できるファイル群。 | [GitHub Docs](https://docs.github.com/en/actions/concepts/workflows-and-actions/workflow-artifacts) |
| GitHub Actions budget | Actions等の従量課金製品の支出を監視・制限するbudget設定。 | [GitHub Docs](https://docs.github.com/en/billing/concepts/budgets-and-alerts) |
| Commit | Git履歴に保存されるプロジェクト状態のスナップショット。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Object ID / commit SHA | Git objectを識別するobject name。特定commitを固定するときに使う。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |

## AI / API / 学習

| 用語 | 短い説明 | 一次情報 |
|---|---|---|
| REST API | HTTP経由でresourceを操作するAPI設計の一形態。 | [GitHub Docs](https://docs.github.com/en/rest/about-the-rest-api) |
| Model Context Protocol (MCP) | LLMアプリを外部data sourceやtoolへ接続するためのオープンプロトコル。 | [MCP Specification](https://modelcontextprotocol.io/specification/2025-11-25) |
| NotebookLM | sourceに基づく質問回答、要約、学習支援を行うGoogleのAIサービス。 | [Google Help](https://support.google.com/notebooklm/answer/16164461) |

## Database / Supabase / 認証

| 用語 | 短い説明 | 一次情報 |
|---|---|---|
| Supabase | Postgresを中核にDatabase、Auth、Storage、Realtime、Edge Functions等を提供する開発platform。 | [Supabase Docs](https://supabase.com/docs) |
| PostgreSQL / Postgres | オープンソースのobject-relational database management system。 | [PostgreSQL Docs](https://www.postgresql.org/docs/current/) |
| Row Level Security (RLS) | tableの行ごとに参照・変更権限をpolicyで制御するPostgresの機能。 | [Supabase Docs](https://supabase.com/docs/guides/database/postgres/row-level-security) |
| OAuth 2.0 | 第三者applicationへ限定的なresource accessを与える認可framework。 | [RFC 6749](https://www.rfc-editor.org/rfc/rfc6749.html) |
| Supabase Edge Functions | Supabaseが提供するserver-side function実行環境。 | [Supabase Docs](https://supabase.com/docs/guides/functions/auth) |
| Database schema | PostgreSQLではSQL objectのnamespace。広義にはdata structureの定義。 | [PostgreSQL Docs](https://www.postgresql.org/docs/current/ddl-schemas.html) |

## 検証 / Web / Knowledge representation

| 用語 | 短い説明 | 一次情報 |
|---|---|---|
| SHA-256 | messageからdigestを生成するhash algorithm。成果物の同一性確認に使える。 | [NIST CSRC](https://csrc.nist.gov/glossary/term/sha_256) |
| Deployment | buildしたapplicationや成果物を公開・実行環境へ配置すること。 | [Vercel Docs](https://vercel.com/docs/deployments/overview) |
| Provenance | dataや成果物が何によって生成・影響・提供されたかを記録する来歴情報。 | [W3C PROV Primer](https://www.w3.org/TR/prov-primer/) |
| Ontology | domainの用語と、その用語間の関係を形式化した共有語彙。 | [W3C OWL 2 Overview](https://www.w3.org/TR/owl-overview/) |
| End-to-end testing (E2E) | user操作に近い形でapplicationを端から端まで通して確認するtest。 | [Playwright Docs](https://playwright.dev/docs/library) |

## 未検証語の扱い

このタスクで取得できた直近会話から抽出したうち、まだ一次情報・標準・原論文で定義を確認していない265語は、[`recent-term-inventory.yaml`](../data/glossary/recent-term-inventory.yaml) に `needs_review` として保持します。

対象には `fail-closed`、`version pin`、`data lake`、`data mart`、`future leakage`、`walk-forward`、`out-of-sample`、`PITR`、`Neo4j`、`Graphiti`、`RAG` に加え、VR/3Dの `Shader` / `Single Pass Instanced` / `SRP Batcher` / `UdonSharp`、GitHub automationの `workflow_call` / `OIDC` / `GITHUB_TOKEN`、agent系の `ReAct` / `context compaction`、音声・home automationの `Whisper` / `Piper` / `Matter` / `Thread` なども含みます。確認が完了した語だけを `data/glossary/terms.yaml` へ昇格します。
