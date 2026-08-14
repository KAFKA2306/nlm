# 学習用語集

> `data/glossary/terms.yaml` から自動生成。詳しく見たい語は `nlm glossary show <term>`。直接編集しないでください。

| 用語 | 一言説明 | 一次情報 |
|---|---|---|
| GitHub Actions | GitHub上のイベントやスケジュールを契機に、ビルド、テスト、デプロイなどの自動処理を実行する仕組み。 | [Concepts for GitHub Actions](https://docs.github.com/en/actions/concepts) |
| Workflow | 1つ以上のjobを実行する設定可能な自動プロセス。 | [Workflows](https://docs.github.com/en/actions/concepts/workflows-and-actions/workflows) |
| Runner | GitHub Actionsのjobを実際に実行するマシンまたは実行環境。 | [GitHub Actions Runners](https://docs.github.com/en/actions/concepts/runners) |
| Continuous Integration | コードを共有リポジトリへ頻繁に統合し、その更新を継続的にビルド・テストして不具合を早期検出する実践。 | [Continuous integration](https://docs.github.com/en/actions/get-started/continuous-integration) |
| Pull Request | GitHub上でコード変更を提案し、議論・レビュー・mergeするための単位。 | [Pull requests documentation](https://docs.github.com/en/pull-requests) |
| Branch | Gitにおける独立した開発ライン。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Remote-tracking branch | 別のリポジトリ上のbranchの状態を追跡するための参照。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Merge | 別branchの内容を現在のbranchへ取り込む操作。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Squash and merge | PR内の複数commitを1つのcommitへまとめてbase branchへmergeする方式。 | [Pull request merges](https://docs.github.com/en/pull-requests/reference/pull-request-merges) |
| Draft Pull Request | まだmerge可能な完成状態ではないことを示すPull Request。 | [Merging a pull request](https://docs.github.com/en/pull-requests/how-tos/merge-and-close-pull-requests/merging-a-pull-request) |
| GitHub Pages | GitHubリポジトリからWebサイトを公開できるホスティング機能。 | [GitHub Pages documentation](https://docs.github.com/en/pages) |
| Workflow artifact | GitHub Actionsのworkflow run中に生成され、job完了後も保存・共有できるファイルまたはファイル群。 | [Workflow artifacts](https://docs.github.com/en/actions/concepts/workflows-and-actions/workflow-artifacts) |
| GitHub Actions budget | GitHub Actionsなどの従量課金製品について、支出額を監視または上限到達時に利用停止させるためのbudget設定。 | [Budgets and alerts](https://docs.github.com/en/billing/concepts/budgets-and-alerts) |
| Commit | Git履歴に保存されるプロジェクト状態のスナップショット。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| Object ID | Git objectを一意に識別するobject name。 | [Git Glossary](https://git-scm.com/docs/gitglossary.html) |
| SHA-256 | メッセージからダイジェストを生成するハッシュアルゴリズム。 | [SHA-256 - NIST CSRC Glossary](https://csrc.nist.gov/glossary/term/sha_256) |
| REST API | HTTP経由でリソースを操作するAPI設計の一形態。 | [About the REST API](https://docs.github.com/en/rest/about-the-rest-api) |
| Model Context Protocol | LLMアプリケーションを外部データソースやツールへ標準化された形で接続するためのオープンプロトコル。 | [Model Context Protocol Specification](https://modelcontextprotocol.io/specification/2025-11-25) |
| NotebookLM | GoogleのAIリサーチ支援サービス。 | [Learn about NotebookLM](https://support.google.com/notebooklm/answer/16164461) |
| Supabase | Postgresを中核にDatabase、Auth、Storage、Realtime、Edge Functionsなどを提供する開発プラットフォーム。 | [Supabase Docs](https://supabase.com/docs) |
| PostgreSQL | オープンソースのオブジェクト・リレーショナルデータベース管理システム。 | [PostgreSQL Documentation](https://www.postgresql.org/docs/current/) |
| Row Level Security | テーブルの行ごとに、誰がどの行を参照・変更できるかをpolicyで制御するPostgresのアクセス制御機能。 | [Row Level Security \| Supabase Docs](https://supabase.com/docs/guides/database/postgres/row-level-security) |
| OAuth 2.0 | 第三者アプリケーションにHTTPサービスへの限定的なアクセスを与えるための認可フレームワーク。 | [RFC 6749: The OAuth 2.0 Authorization Framework](https://www.rfc-editor.org/rfc/rfc6749.html) |
| Supabase Edge Functions | Supabaseが提供するserver-side function実行環境。 | [Securing Edge Functions](https://supabase.com/docs/guides/functions/auth) |
| Database schema | PostgreSQLではSQL objectのnamespace。 | [PostgreSQL Glossary](https://www.postgresql.org/docs/current/glossary.html) |
| Deployment | buildしたアプリケーションや静的成果物を実行・公開環境へ配置し、利用可能な状態にすること。 | [Deploying to Vercel](https://vercel.com/docs/deployments/overview) |
| Provenance | データや成果物が、どのentity・person・processによって生成、影響、提供されたかを記録する情報。 | [PROV Model Primer](https://www.w3.org/TR/prov-primer/) |
| Ontology | 特定domainの用語を形式化し、用語同士の関係を定義する共有語彙。 | [OWL 2 Web Ontology Language Document Overview](https://www.w3.org/TR/owl-overview/) |
| End-to-end testing | ユーザーが実際に行う操作に近い形で、アプリケーションを端から端まで通して挙動確認するテスト。 | [Playwright Library](https://playwright.dev/docs/library) |
