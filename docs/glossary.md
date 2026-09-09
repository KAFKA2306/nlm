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
| MCP server | MCP clientへcontextやcapabilityを提供し、resources・prompts・toolsなどのserver featureを公開するserver。 | [Architecture overview](https://modelcontextprotocol.io/docs/learn/architecture) |
| NotebookLM source | NotebookLMへimportまたはuploadした文書のcopyまたは自動同期版で、モデルが質問回答やrequest処理の根拠として使う情報源。 | [Add or discover new sources for your notebook](https://support.google.com/notebooklm/answer/16215270) |
| NotebookLM note | NotebookLMで情報を記録・整理し、sourceから得た洞察や解釈、自分の考え、保存したchat responseなどを保持するnotebook内のメモ。 | [Create & add notes in NotebookLM](https://support.google.com/notebooklm/answer/16262519) |
| Large Language Model | 非常に大規模なデータセットにdeep learningを適用し、自然なテキストを予測・構成する計算モデル。 | [Large Language Models MeSH Descriptor Data 2026](https://meshb.nlm.nih.gov/record/ui?ui=D000098342) |
| Retrieval-Augmented Generation | 事前学習済み生成モデルのparametric memoryと、検索でアクセスするnon-parametric memoryを組み合わせて文章生成を行う手法。 | [Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks](https://papers.nips.cc/paper/2020/hash/6b493230205f780e1bc26945df7481e5-Abstract.html) |
| Remote MCP server | local machineではなくinternet上にhostされ、MCP clientへtools・prompts・resourcesなどを提供するMCP server。 | [Connect to remote MCP Servers](https://modelcontextprotocol.io/docs/develop/connect-remote-servers) |
| tools/list | MCP clientがserverへ利用可能なtool一覧を要求するためのJSON-RPC method。 | [Schema Reference - tools/list](https://modelcontextprotocol.io/specification/2025-11-25/schema) |
| DINKs | 結婚するが子どもは持たず、仕事を一生続けるライフコースを指す語。 | [平成25年版 男女共同参画白書 第1部 第1節](https://www.gender.go.jp/about_danjo/whitepaper/h25/zentai/html/honpen/b1_s00_03.html) |
| ひとり親世帯 | 父または母の一方と、20歳未満の未婚の子どもから構成される世帯を扱う公的調査上の概念。 | [令和3年度 全国ひとり親世帯等調査の結果](https://www.mhlw.go.jp/stf/seisakunitsuite/bunya/0000188147_00013.html) |
| 共働き世帯 | 夫妻の双方が就業または有業である世帯を扱う統計上の概念。 | [統計局FAQ 16A-Q12 共働き世帯に関する統計](https://www.stat.go.jp/library/faq/faq16/faq16a12.html) |
| Private repository | アクセスできるユーザーが所有者、明示的に共有されたユーザー、または組織内で許可されたメンバーに制限されるGitHub repository。 | [About repositories](https://docs.github.com/en/repositories/creating-and-managing-repositories/about-repositories) |
| Self-hosted runner | 利用者側が管理する環境でGitHub Actionsのjobを実行するrunner。 | [Self-hosted runners reference](https://docs.github.com/en/actions/reference/runners/self-hosted-runners) |
| GitHub Issue | GitHub上でアイデア、フィードバック、タスク、bugなどを計画・議論・追跡する作業項目。 | [GitHub Issues documentation](https://docs.github.com/en/issues) |
| Repository topic | repositoryの目的、分野、言語などを分類し、関連repositoryの発見や検索に使うGitHubのmetadata。 | [Classifying your repository with topics](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/classifying-your-repository-with-topics) |
| Workflow run | GitHub Actions workflowがtriggerされて作成される1回の実行で、状態、結果、job、stepのlogを持つ実行単位。 | [Using workflow run logs](https://docs.github.com/en/actions/how-tos/monitor-workflows/use-workflow-run-logs) |
| GitHub Actions job | GitHub Actions workflow run内で1つ以上のstepをrunner上で実行する実行単位。 | [Using workflow run logs](https://docs.github.com/en/actions/how-tos/monitor-workflows/use-workflow-run-logs) |
| Reusable workflow | 別のworkflowから呼び出して共通処理を再利用できるGitHub Actions workflow。 | [Reusing workflow configurations](https://docs.github.com/en/actions/reference/workflows-and-actions/reusing-workflow-configurations) |
| workflow_call | GitHub Actions workflowを他のworkflowから呼び出せるreusable workflowとして定義するためにonで指定するtrigger。 | [Reusing workflow configurations](https://docs.github.com/en/actions/reference/workflows-and-actions/reusing-workflow-configurations) |
| Caller workflow | reusable workflowをjobのusesで呼び出す側のGitHub Actions workflow。 | [Reusing workflow configurations](https://docs.github.com/en/actions/reference/workflows-and-actions/reusing-workflow-configurations) |
| Called workflow | caller workflowから呼び出され、定義済みの処理を提供するreusable workflow。 | [Reusing workflow configurations](https://docs.github.com/en/actions/reference/workflows-and-actions/reusing-workflow-configurations) |
| GITHUB_TOKEN | workflow job開始時にGitHubが自動生成し、そのrepositoryに対する認証に使えるGitHub App installation access token。 | [GITHUB_TOKEN](https://docs.github.com/en/actions/concepts/security/github_token) |
| GitHub permissions | GitHub ActionsでGITHUB_TOKENなどに許可する操作範囲をworkflowまたはjob単位で制御する設定。 | [Reusing workflow configurations](https://docs.github.com/en/actions/reference/workflows-and-actions/reusing-workflow-configurations) |
| OpenID Connect | GitHub Actions workflowが外部cloud providerなどへ短命tokenで認証するために利用できるidentity federation方式。 | [OpenID Connect reference](https://docs.github.com/en/actions/reference/security/oidc) |
| id-token: write | GitHub ActionsのjobまたはworkflowにOIDC tokenの要求を許可するpermissions設定で、外部resourceへの書き込み権限そのものではない。 | [OpenID Connect reference](https://docs.github.com/en/actions/reference/security/oidc) |
| JSON | 構造化データを可搬な形で表現するための、軽量・テキストベース・言語非依存のデータ交換形式。 | [RFC 8259: The JavaScript Object Notation (JSON) Data Interchange Format](https://www.rfc-editor.org/info/rfc8259) |
| Parquet | 効率的な保存と取得を目的に設計された、オープンソースの列指向データファイル形式。 | [Apache Parquet](https://parquet.apache.org/) |
| CSV | レコードを行、fieldをcommaで区切って表す広く使われるテキスト形式。 | [RFC 4180: Common Format and MIME Type for Comma-Separated Values (CSV) Files](https://www.rfc-editor.org/info/rfc4180) |
| Unique constraint | 1列または複数列の値の組み合わせがtable内の全rowで一意になることを保証するdatabase constraint。 | [PostgreSQL 18: Constraints — Unique Constraints](https://www.postgresql.org/docs/current/ddl-constraints.html) |
| Indexing | table全体を順番に走査せず、条件に合うrowをより効率よく探せるようにindexを維持・利用する仕組み。 | [PostgreSQL 18: Indexes — Introduction](https://www.postgresql.org/docs/current/indexes-intro.html) |
| Point-in-Time Recovery | base backupとarchived Write-Ahead Logを使い、WAL replayを指定時点で止めてdatabaseをその時点の整合した状態へ戻す復旧方式。 | [PostgreSQL 18: Continuous Archiving and Point-in-Time Recovery (PITR)](https://www.postgresql.org/docs/current/continuous-archiving.html) |
| SQLite | 別server processを必要とせずapplication process内で動作する、self-contained・serverless・zero-configuration・transactionalなSQL database engine。 | [About SQLite](https://sqlite.org/about.html) |
| RLS policy | Row-Level Securityを有効にしたtableで、userやcommandごとにどのrowを参照・追加・更新・削除できるかをBoolean expressionで制御するpolicy。 | [PostgreSQL 18: Row Security Policies](https://www.postgresql.org/docs/current/ddl-rowsecurity.html) |
