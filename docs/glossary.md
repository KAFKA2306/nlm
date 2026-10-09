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
| Document Object Model | HTMLやXMLなどのdocumentをnode treeとして表し、documentの構造へアクセス・操作するためのWeb platformのmodelとAPI。 | [DOM Standard](https://dom.spec.whatwg.org/) |
| CSS align-content | alignment containerの中でcontentをblock/cross axis方向に配置・分配するCSS property。 | [CSS Box Alignment Module Level 3](https://www.w3.org/TR/css-align-3/) |
| CSS align-items | containerの子boxに対するdefault align-selfを設定し、block/cross axis方向のself-alignmentを決めるCSS property。 | [CSS Box Alignment Module Level 3](https://www.w3.org/TR/css-align-3/) |
| Vercel | source codeやframework projectからWeb applicationやserviceをbuild・deployし、global deliveryやserver-side computeなどを提供するdeployment platform。 | [Vercel](https://vercel.com/) |
| FastAPI | Pythonの標準type hintを中心にAPIを定義し、OpenAPI schemaやinteractive API documentationを生成できるPython Web API framework。 | [FastAPI Features](https://fastapi.tiangolo.com/features/) |
| Home Assistant | local controlとprivacyを重視する、freeかつopen-sourceのhome automation platform。 | [Home Assistant](https://www.home-assistant.io/) |
| Matter | Connectivity Standards Allianceが策定する、異なるbrandやecosystemのIoT device間のinteroperabilityを目的としたIP-based connectivity standard。 | [Build With Matter](https://csa-iot.org/all-solutions/matter/) |
| Thread | IoT向けに設計された、IPv6-based・low-power・self-healing meshのwireless networking protocol。 | [Thread Group — What is Thread?](https://threadgroup.org/) |
| Home APIs | Android/iOS applicationからGoogle Homeのdevice、structure、Matter infrastructure、commissioning、automation機能へaccessするためのGoogle Home developer API群。 | [Home APIs](https://developers.home.google.com/apis) |
| Voice Match | Google Homeのshared device上でvoice assistantが話者の声を認識し、そのGoogle Accountに基づくpersonalized responseやservice selectionへ使う機能。 | [Set up and manage Voice Match for your home or devices](https://support.google.com/googlehome/answer/7320960) |
| Deep Q-Network | Q学習などの強化学習と深層ニューラルネットワークを組み合わせ、高次元の感覚入力から行動価値を近似して方策を学習する手法。 | [Human-level control through deep reinforcement learning](https://www.nature.com/articles/nature14236) |
| Long Short-Term Memory | 長い時間間隔をまたぐ依存関係を学習しやすくするため、一定の誤差信号の流れとゲート機構を導入した再帰型ニューラルネットワークの方式。 | [Long Short-Term Memory](https://direct.mit.edu/neco/article/9/8/1735/6109/Long-Short-Term-Memory) |
| Walk-forward evaluation | 時系列の順序を保ち、過去側のデータを学習に使って、その後の未来側区間を評価しながら評価時点を前へ進める方法。 | [TimeSeriesSplit — scikit-learn 1.9.0 documentation](https://scikit-learn.org/stable/modules/generated/sklearn.model_selection.TimeSeriesSplit.html) |
| Data split | モデルの学習用データと評価用データなど、異なる役割を持つ部分集合へデータを分割すること。 | [train_test_split — scikit-learn 1.9.0 documentation](https://scikit-learn.org/stable/modules/generated/sklearn.model_selection.train_test_split.html) |
| Exploratory Data Analysis | データの構造、外れ値、仮定、重要な変数などへの洞察を得るため、主に可視化を含む多様な技法でデータを探索する分析アプローチ。 | [NIST/SEMATECH e-Handbook of Statistical Methods; Chapter 1: Exploratory Data Analysis](https://www.nist.gov/publications/nistsematech-e-handbook-statistical-methods-chapter-1-exploratory-data-analysis) |
| Feature engineering | モデル学習に有用な特徴を選び、元のデータを学習しやすい特徴表現へ変換するプロセス。 | [Machine Learning Glossary — Google for Developers](https://developers.google.com/machine-learning/glossary) |
| Google OAuth | Google AccountのOAuth 2.0認可endpointとtoken endpointを使い、userの同意に基づいてapplicationへGoogle APIへの限定的なaccess権を付与する認可フロー。 | [Using OAuth 2.0 for Web Server Applications](https://developers.google.com/identity/protocols/oauth2/web-server) |
| Fail-closed | 認可判定や保護機構が失敗または不確定になったとき、accessを許可せず拒否側へ倒す安全設計。 | [OWASP Authorization Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html) |
| Authorization | 主体がresourceやoperationへaccessする権限を持つかを決定し、許可または拒否すること。 | [NIST CSRC Glossary — Authorization](https://csrc.nist.gov/glossary/term/authorization) |
| SSL/TLS | network通信を暗号学的に保護するprotocol群を指す慣用的な表現。 | [RFC 9846 — The Transport Layer Security (TLS) Protocol Version 1.3](https://www.rfc-editor.org/rfc/rfc9846.html) |
| Supabase Google OAuth | Supabase AuthでGoogleを外部identity providerとして設定し、GoogleのOAuth flowを介してuserをsign inさせるintegration。 | [Sign in with Google — Supabase Docs](https://supabase.com/docs/guides/auth/social-login/auth-google) |
| HF_TOKEN | Hugging Face Hubへの認証に使うUser Access Tokenを環境変数から指定する設定。 | [Environment variables — Hugging Face Hub Python Library](https://huggingface.co/docs/huggingface_hub/package_reference/environment_variables) |
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
