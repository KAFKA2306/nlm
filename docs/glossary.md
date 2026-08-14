# 学習用語集

> このファイルは `data/glossary/terms.yaml` から `nlm glossary generate` で生成します。直接編集しないでください。

- Schema version: `1`
- Terms: **29**
- Verified at: `2026-08-14`
- Source scope: Recent ChatGPT conversations around KAFKA2306 development, research, deployment, and data workflows

## GitHub Actions

`github-actions` · `github` · `verified`

GitHub上のイベントやスケジュールを契機に、ビルド、テスト、デプロイなどの自動処理を実行する仕組み。

- **Aliases:** `Actions`
- **なぜ重要か:** 直近の作業では、CI、Pages配信、定期実行、Issue/PR運用の自動化基盤として繰り返し登場した。
- **例:** pushやpull requestを契機にテストworkflowを実行する。
- **関連語:** `workflow`, `runner`, `continuous-integration`, `github-pages`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Concepts for GitHub Actions](https://docs.github.com/en/actions/concepts) — `official_docs`

## Workflow

`workflow` · `github` · `verified`

1つ以上のjobを実行する設定可能な自動プロセス。GitHub ActionsではYAMLファイルとしてリポジトリの .github/workflows に定義する。

- **Aliases:** `GitHub Actions workflow`
- **なぜ重要か:** Actionsの失敗原因や定期処理を理解するときの最上位単位になる。
- **例:** CI workflow、GitHub Pages deploy workflow。
- **関連語:** `github-actions`, `runner`, `workflow-artifact`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Workflows](https://docs.github.com/en/actions/concepts/workflows-and-actions/workflows) — `official_docs`

## Runner

`runner` · `github` · `verified`

GitHub Actionsのjobを実際に実行するマシンまたは実行環境。GitHub-hostedとself-hostedがある。

- **Aliases:** `GitHub Actions runner`
- **なぜ重要か:** 直近ではBillingやspending制約によりrunner起動前に停止するケースの切り分けで重要だった。
- **例:** runs-on: ubuntu-latest でGitHub-hosted runner上にjobを割り当てる。
- **関連語:** `github-actions`, `workflow`, `github-actions-budget`
- **確認日:** `2026-08-14`
- **Sources:**
  - [GitHub Actions Runners](https://docs.github.com/en/actions/concepts/runners) — `official_docs`

## Continuous Integration

`continuous-integration` · `software-engineering` · `verified`

コードを共有リポジトリへ頻繁に統合し、その更新を継続的にビルド・テストして不具合を早期検出する実践。

- **Aliases:** `CI`, `継続的インテグレーション`
- **なぜ重要か:** PRをmergeする前に変更が既存挙動を壊していないか確認する基本ゲートになる。
- **例:** PR head SHAに対してgo testやlintを自動実行する。
- **関連語:** `github-actions`, `workflow`, `pull-request`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Continuous integration](https://docs.github.com/en/actions/get-started/continuous-integration) — `official_docs`

## Pull Request

`pull-request` · `github` · `verified`

GitHub上でコード変更を提案し、議論・レビュー・mergeするための単位。

- **Aliases:** `PR`
- **なぜ重要か:** 直近の自律開発では、実装をmainへ入れる前の正準な作業線として繰り返し使われた。
- **例:** feature branchからmainへの変更をPRとして提案する。
- **関連語:** `branch`, `merge`, `squash-merge`, `draft-pull-request`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Pull requests documentation](https://docs.github.com/en/pull-requests) — `official_docs`

## Branch

`branch` · `git` · `verified`

Gitにおける独立した開発ライン。既存の安定版へ直接影響させずに変更を進めるために使う。

- **Aliases:** `Git branch`
- **なぜ重要か:** 作業branch、main、不要branchのcleanupという形で直近のGitHub運用に頻出した。
- **例:** feat/... branchで作業し、PR merge後に削除する。
- **関連語:** `pull-request`, `remote-tracking-branch`, `merge`, `commit`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Git Glossary](https://git-scm.com/docs/gitglossary.html) — `official_docs`

## Remote-tracking branch

`remote-tracking-branch` · `git` · `verified`

別のリポジトリ上のbranchの状態を追跡するための参照。通常 refs/remotes/<remote>/<branch> の形を取る。

- **Aliases:** `remote branch`, `リモート追跡ブランチ`
- **なぜ重要か:** merge後にremote側へ不要branchが残っていないか監査する文脈で使われた。
- **例:** origin/main はremote repositoryのmainを追跡する参照。
- **関連語:** `branch`, `commit`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Git Glossary](https://git-scm.com/docs/gitglossary.html) — `official_docs`

## Merge

`merge` · `git` · `verified`

別branchの内容を現在のbranchへ取り込む操作。GitHubではPRを通じて実行されることが多い。

- **Aliases:** `マージ`
- **なぜ重要か:** 実装完了後にmainへ成果を確定させる操作として、CI成功と並んで完了条件になっている。
- **例:** PRの変更をmainへmergeする。
- **関連語:** `pull-request`, `branch`, `squash-merge`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Git Glossary](https://git-scm.com/docs/gitglossary.html) — `official_docs`

## Squash and merge

`squash-merge` · `github` · `verified`

PR内の複数commitを1つのcommitへまとめてbase branchへmergeする方式。

- **Aliases:** `squash merge`
- **なぜ重要か:** 直近のPR完了操作で明示的に選ばれており、履歴を1変更1commitへ整理するために使われた。
- **例:** 小さなfixup commitをまとめてmainへ1commitとして入れる。
- **関連語:** `merge`, `pull-request`, `commit`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Pull request merges](https://docs.github.com/en/pull-requests/reference/pull-request-merges) — `official_docs`

## Draft Pull Request

`draft-pull-request` · `github` · `verified`

まだmerge可能な完成状態ではないことを示すPull Request。GitHubではdraftのPRはmergeできない。

- **Aliases:** `Draft PR`
- **なぜ重要か:** 実装途中の正準作業線を残す場合と、不要・重複Draft PRをcleanupする場合の区別に使われた。
- **例:** 未完了の実装をDraft PRとして共有し、完成後にready for reviewへ移す。
- **関連語:** `pull-request`, `merge`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Merging a pull request](https://docs.github.com/en/pull-requests/how-tos/merge-and-close-pull-requests/merging-a-pull-request) — `official_docs`

## GitHub Pages

`github-pages` · `github` · `verified`

GitHubリポジトリからWebサイトを公開できるホスティング機能。branchまたはGitHub Actions workflowを公開元にできる。

- **Aliases:** `Pages`
- **なぜ重要か:** 直近ではCI成功後の公開先、deploy確認先、repository settingのblockerとして登場した。
- **例:** Actionsで静的ファイルをartifact化し、deploy-pagesで公開する。
- **関連語:** `github-actions`, `workflow-artifact`, `deployment`
- **確認日:** `2026-08-14`
- **Sources:**
  - [GitHub Pages documentation](https://docs.github.com/en/pages) — `official_docs`
  - [Using custom workflows with GitHub Pages](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages) — `official_docs`

## Workflow artifact

`workflow-artifact` · `github` · `verified`

GitHub Actionsのworkflow run中に生成され、job完了後も保存・共有できるファイルまたはファイル群。

- **Aliases:** `Actions artifact`, `artifact`
- **なぜ重要か:** テスト結果、build成果、証拠ファイル、Pages deploy入力などを残す用途で繰り返し登場した。
- **例:** build出力をupload-artifactで保存して別jobから利用する。
- **関連語:** `workflow`, `github-actions`, `provenance`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Workflow artifacts](https://docs.github.com/en/actions/concepts/workflows-and-actions/workflow-artifacts) — `official_docs`

## GitHub Actions budget

`github-actions-budget` · `github` · `verified`

GitHub Actionsなどの従量課金製品について、支出額を監視または上限到達時に利用停止させるためのbudget設定。

- **Aliases:** `spending limit`, `Actions billing budget`
- **なぜ重要か:** private repositoryのActionsがrunner起動前にBilling / spending条件で停止した原因切り分けに直接関係した。
- **例:** Actionsのbudget上限到達により追加のrunner利用を止める。
- **関連語:** `github-actions`, `runner`, `workflow`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Budgets and alerts](https://docs.github.com/en/billing/concepts/budgets-and-alerts) — `official_docs`

## Commit

`commit` · `git` · `verified`

Git履歴に保存されるプロジェクト状態のスナップショット。親commit、author、committer、日時、treeなどの情報を持つ。

- **Aliases:** `Git commit`
- **なぜ重要か:** PR headの固定、mainの状態確認、成果の証拠URL提示でcommit単位の識別が使われた。
- **例:** CIが特定のcommitに対して成功したことを確認する。
- **関連語:** `branch`, `object-id`, `squash-merge`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Git Glossary](https://git-scm.com/docs/gitglossary.html) — `official_docs`

## Object ID

`object-id` · `git` · `verified`

Git objectを一意に識別するobject name。会話ではcommit SHAとして、PR headやmainの状態固定に使われた。

- **Aliases:** `OID`, `object name`, `commit SHA`
- **なぜ重要か:** branch名のような可変参照ではなく、特定のcommitを固定してCIやmerge対象を確認できる。
- **例:** expected head SHAを指定してPR headが動いていないことを確認してからmergeする。
- **関連語:** `commit`, `sha-256`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Git Glossary](https://git-scm.com/docs/gitglossary.html) — `official_docs`

## SHA-256

`sha-256` · `security` · `verified`

メッセージからダイジェストを生成するハッシュアルゴリズム。生成後に内容が変わっていないか確認する用途に使える。

- **Aliases:** `Secure Hash Algorithm 256`
- **なぜ重要か:** 画像・PDF・artifactの再取得後一致確認やprovenance証跡で頻出した。
- **例:** アップロード前後のファイルSHA-256を比較して同一性を確認する。
- **関連語:** `provenance`, `object-id`
- **確認日:** `2026-08-14`
- **Sources:**
  - [SHA-256 - NIST CSRC Glossary](https://csrc.nist.gov/glossary/term/sha_256) — `standard_body`

## REST API

`rest-api` · `software-engineering` · `verified`

HTTP経由でリソースを操作するAPI設計の一形態。GitHubではREST APIを用いてIssue、PR、repository、Actionsなどを自動操作できる。

- **Aliases:** `REST`
- **なぜ重要か:** GitHubの自動操作や外部サービス連携の実装面で頻繁に登場した。
- **例:** GitHub REST APIでIssueを作成する。
- **関連語:** `github-actions`, `mcp`
- **確認日:** `2026-08-14`
- **Sources:**
  - [About the REST API](https://docs.github.com/en/rest/about-the-rest-api) — `official_docs`

## Model Context Protocol

`mcp` · `ai` · `verified`

LLMアプリケーションを外部データソースやツールへ標準化された形で接続するためのオープンプロトコル。

- **Aliases:** `MCP`
- **なぜ重要か:** GitHub、研究、データベース、NotebookLMなどの外部機能をAIから扱う文脈で繰り返し登場した。
- **例:** LLMクライアントからMCP serverが公開するtoolを呼び出す。
- **関連語:** `rest-api`, `notebooklm`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Model Context Protocol Specification](https://modelcontextprotocol.io/specification/2025-11-25) — `protocol_spec`

## NotebookLM

`notebooklm` · `ai` · `verified`

GoogleのAIリサーチ支援サービス。ソースを取り込み、そのソースに基づく質問回答、要約、学習ガイド等の生成を行える。

- **Aliases:** -
- **なぜ重要か:** nlm repository自体の対象であり、用語集を学習ガイドやsourceへ再利用する先として選定した。
- **例:** 用語集MarkdownをNotebookLM sourceとして追加し、学習ガイドを生成する。
- **関連語:** `mcp`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Learn about NotebookLM](https://support.google.com/notebooklm/answer/16164461) — `official_docs`
  - [Add or discover new sources for your notebook](https://support.google.com/notebooklm/answer/16215270) — `official_docs`

## Supabase

`supabase` · `database` · `verified`

Postgresを中核にDatabase、Auth、Storage、Realtime、Edge Functionsなどを提供する開発プラットフォーム。

- **Aliases:** -
- **なぜ重要か:** 画像資産、認証、RLS、Webアプリのbackend設定に関する直近会話で頻出した。
- **例:** Supabase Storageへ画像を保存し、RLSでアクセス制御する。
- **関連語:** `postgresql`, `row-level-security`, `edge-functions`, `oauth-2`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Supabase Docs](https://supabase.com/docs) — `official_docs`

## PostgreSQL

`postgresql` · `database` · `verified`

オープンソースのオブジェクト・リレーショナルデータベース管理システム。SupabaseのDatabase基盤としても使われる。

- **Aliases:** `Postgres`
- **なぜ重要か:** Supabase、RLS、schema、データ管理の会話で基礎技術として登場した。
- **例:** tableとschemaをPostgres上で定義する。
- **関連語:** `supabase`, `row-level-security`, `database-schema`
- **確認日:** `2026-08-14`
- **Sources:**
  - [PostgreSQL Documentation](https://www.postgresql.org/docs/current/) — `official_docs`

## Row Level Security

`row-level-security` · `database-security` · `verified`

テーブルの行ごとに、誰がどの行を参照・変更できるかをpolicyで制御するPostgresのアクセス制御機能。

- **Aliases:** `RLS`
- **なぜ重要か:** Supabase上のユーザー別データ、collection、Like等の認可条件として直近のIssue設計に登場した。
- **例:** auth.uid()とuser_idが一致する行だけSELECT可能にするpolicyを設定する。
- **関連語:** `supabase`, `postgresql`, `oauth-2`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Row Level Security | Supabase Docs](https://supabase.com/docs/guides/database/postgres/row-level-security) — `official_docs`

## OAuth 2.0

`oauth-2` · `security` · `verified`

第三者アプリケーションにHTTPサービスへの限定的なアクセスを与えるための認可フレームワーク。

- **Aliases:** `OAuth`
- **なぜ重要か:** Supabase Google OAuthなど、外部identity providerとの認証・認可設定のblockerとして登場した。
- **例:** ユーザー承認を経てGoogleアカウント由来の認可情報をアプリへ渡す。
- **関連語:** `supabase`, `row-level-security`
- **確認日:** `2026-08-14`
- **Sources:**
  - [RFC 6749: The OAuth 2.0 Authorization Framework](https://www.rfc-editor.org/rfc/rfc6749.html) — `standard`

## Supabase Edge Functions

`edge-functions` · `serverless` · `verified`

Supabaseが提供するserver-side function実行環境。認証情報を検証し、Databaseや外部APIと連携する処理を置ける。

- **Aliases:** `Edge Functions`
- **なぜ重要か:** Supabaseのserver-side処理、API、認証境界を理解するための基礎語として登場した。
- **例:** 認証済みユーザーのJWTを検証してPostgresへアクセスする。
- **関連語:** `supabase`, `postgresql`, `row-level-security`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Securing Edge Functions](https://supabase.com/docs/guides/functions/auth) — `official_docs`
  - [Integrating with Supabase Database (Postgres)](https://supabase.com/docs/guides/functions/connect-to-postgres) — `official_docs`

## Database schema

`database-schema` · `database` · `verified`

PostgreSQLではSQL objectのnamespace。より広い意味では、table定義やconstraintなどデータ構造の記述全体を指す。

- **Aliases:** `schema`, `スキーマ`
- **なぜ重要か:** 二重schemaを作らない、正準schemaを固定する、という設計判断で頻出した。
- **例:** collection用tableを既存schemaへ統合し、別schemaを重複作成しない。
- **関連語:** `postgresql`, `ontology`
- **確認日:** `2026-08-14`
- **Sources:**
  - [PostgreSQL Glossary](https://www.postgresql.org/docs/current/glossary.html) — `official_docs`
  - [Schemas](https://www.postgresql.org/docs/current/ddl-schemas.html) — `official_docs`

## Deployment

`deployment` · `web` · `verified`

buildしたアプリケーションや静的成果物を実行・公開環境へ配置し、利用可能な状態にすること。Vercelでは成功したbuildからdeploymentが生成される。

- **Aliases:** `deploy`, `デプロイ`
- **なぜ重要か:** Vercel、GitHub Pages、本番検証、production反映の会話で頻出した。
- **例:** mainへのpushを契機にproduction deploymentを作成する。
- **関連語:** `github-pages`, `continuous-integration`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Deploying to Vercel](https://vercel.com/docs/deployments/overview) — `official_docs`

## Provenance

`provenance` · `data-governance` · `verified`

データや成果物が、どのentity・person・processによって生成、影響、提供されたかを記録する情報。

- **Aliases:** `来歴`, `data provenance`
- **なぜ重要か:** 画像・render metadata・paper dataset・artifactの出所と検証可能性を示す証拠チェーンとして頻出した。
- **例:** assetのcommit、blob SHA、SHA-256、生成元を記録する。
- **関連語:** `sha-256`, `workflow-artifact`, `ontology`
- **確認日:** `2026-08-14`
- **Sources:**
  - [PROV Model Primer](https://www.w3.org/TR/prov-primer/) — `w3c_recommendation`

## Ontology

`ontology` · `knowledge-representation` · `verified`

特定domainの用語を形式化し、用語同士の関係を定義する共有語彙。

- **Aliases:** `オントロジー`
- **なぜ重要か:** repositoryのsemantic zone分類やfactory domain modelを一貫した意味体系で扱う文脈に登場した。
- **例:** repositoryをagent-zone-*の意味体系に従って分類する。
- **関連語:** `database-schema`, `provenance`
- **確認日:** `2026-08-14`
- **Sources:**
  - [OWL 2 Web Ontology Language Document Overview](https://www.w3.org/TR/owl-overview/) — `w3c_recommendation`

## End-to-end testing

`end-to-end-testing` · `software-testing` · `verified`

ユーザーが実際に行う操作に近い形で、アプリケーションを端から端まで通して挙動確認するテスト。

- **Aliases:** `E2E`, `E2E test`
- **なぜ重要か:** CRUD、認可、mobile UI、本番導線までAcceptance Criteriaに含める文脈で繰り返し登場した。
- **例:** ブラウザからログインし、list作成・並べ替え・削除まで通して検証する。
- **関連語:** `continuous-integration`, `deployment`
- **確認日:** `2026-08-14`
- **Sources:**
  - [Playwright Library](https://playwright.dev/docs/library) — `official_docs`
  - [Playwright Best Practices](https://playwright.dev/docs/best-practices) — `official_docs`
