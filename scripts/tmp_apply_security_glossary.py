from pathlib import Path

TERMS = Path("data/glossary/terms.yaml")
INVENTORY = Path("data/glossary/recent-term-inventory.yaml")

entries = r'''

  - id: "google-oauth"
    term: "Google OAuth"
    aliases: ["Google OAuth 2.0"]
    domain: "security"
    definition_ja: "Google AccountのOAuth 2.0認可endpointとtoken endpointを使い、userの同意に基づいてapplicationへGoogle APIへの限定的なaccess権を付与する認可フロー。"
    why_it_matters: "Google API連携で、sign-inそのものと、user dataへ許可されたscopeでaccessするauthorizationを区別するために重要。"
    example: "Web applicationがauthorization codeを受け取り、access tokenへ交換して許可されたGoogle APIを呼び出す。"
    related_terms: ["authorization"]
    sources:
      - title: "Using OAuth 2.0 for Web Server Applications"
        url: "https://developers.google.com/identity/protocols/oauth2/web-server"
        source_type: "official_docs"
    verified_at: "2026-09-10"
    status: "verified"

  - id: "fail-closed"
    term: "Fail-closed"
    aliases: ["fail closed", "deny on failure"]
    domain: "security"
    definition_ja: "認可判定や保護機構が失敗または不確定になったとき、accessを許可せず拒否側へ倒す安全設計。"
    why_it_matters: "authorization checkの例外や障害をaccess許可として扱うfail-openを避け、異常時にも権限境界を維持するために重要。"
    example: "authorization serviceが応答できない場合、resourceへのrequestを許可せず拒否する。"
    related_terms: ["authorization"]
    sources:
      - title: "OWASP Authorization Cheat Sheet"
        url: "https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html"
        source_type: "official_docs"
    verified_at: "2026-09-10"
    status: "verified"

  - id: "authorization"
    term: "Authorization"
    aliases: ["認可"]
    domain: "security"
    definition_ja: "主体がresourceやoperationへaccessする権限を持つかを決定し、許可または拒否すること。"
    why_it_matters: "identityを確認するauthenticationと、何を実行してよいかを決めるauthorizationを混同しないための基本概念。"
    example: "認証済みuserが特定database rowを更新できるroleを持つか確認する。"
    related_terms: []
    sources:
      - title: "NIST CSRC Glossary — Authorization"
        url: "https://csrc.nist.gov/glossary/term/authorization"
        source_type: "official_docs"
    verified_at: "2026-09-10"
    status: "verified"

  - id: "ssl-tls"
    term: "SSL/TLS"
    aliases: ["TLS", "Transport Layer Security", "Secure Sockets Layer"]
    domain: "security"
    definition_ja: "network通信を暗号学的に保護するprotocol群を指す慣用的な表現。現行標準はTLSであり、SSLv3はIETFにより使用禁止とされている。"
    why_it_matters: "HTTPSなどのtransport protectionを説明するとき、現行TLSと廃止されたSSLを同じ現役protocolとして扱わないために重要。"
    example: "clientとserverがTLS handshakeで暗号parameterを合意し、その後のapplication dataを保護する。"
    related_terms: []
    sources:
      - title: "RFC 9846 — The Transport Layer Security (TLS) Protocol Version 1.3"
        url: "https://www.rfc-editor.org/rfc/rfc9846.html"
        source_type: "official_standard_body"
      - title: "RFC 7568 — Deprecating Secure Sockets Layer Version 3.0"
        url: "https://www.rfc-editor.org/rfc/rfc7568.html"
        source_type: "official_standard_body"
    verified_at: "2026-09-10"
    status: "verified"

  - id: "supabase-google-oauth"
    term: "Supabase Google OAuth"
    aliases: ["Supabase Sign in with Google"]
    domain: "security"
    definition_ja: "Supabase AuthでGoogleを外部identity providerとして設定し、GoogleのOAuth flowを介してuserをsign inさせるintegration。"
    why_it_matters: "Supabase applicationでGoogle Accountをidentity providerとして使う設定と、一般的なGoogle API authorizationを区別するために重要。"
    example: "Supabase projectでGoogle providerを設定し、applicationからGoogle sign-inを開始する。"
    related_terms: ["google-oauth"]
    sources:
      - title: "Sign in with Google — Supabase Docs"
        url: "https://supabase.com/docs/guides/auth/social-login/auth-google"
        source_type: "official_docs"
    verified_at: "2026-09-10"
    status: "verified"

  - id: "hf-token"
    term: "HF_TOKEN"
    aliases: ["Hugging Face token environment variable"]
    domain: "security"
    definition_ja: "Hugging Face Hubへの認証に使うUser Access Tokenを環境変数から指定する設定。設定時はmachineに保存されたtokenより優先される。"
    why_it_matters: "CIや一時実行環境でHub認証情報をcodeへ埋め込まず注入し、どのcredentialが実際に使われるか理解するために重要。"
    example: "CIのsecret storeからHF_TOKENを環境変数へ渡し、private Hub resourceへ認証する。"
    related_terms: []
    sources:
      - title: "Environment variables — Hugging Face Hub Python Library"
        url: "https://huggingface.co/docs/huggingface_hub/package_reference/environment_variables"
        source_type: "official_docs"
    verified_at: "2026-09-10"
    status: "verified"
'''

terms = TERMS.read_text(encoding="utf-8")
if '  - id: "google-oauth"' not in terms:
    TERMS.write_text(terms.rstrip() + entries + "\n", encoding="utf-8")

inventory = INVENTORY.read_text(encoding="utf-8")
old_counts = "  verified: 74\n  needs_review: 201\n"
if old_counts not in inventory:
    raise SystemExit("expected baseline counts not found")
inventory = inventory.replace(old_counts, "  verified: 80\n  needs_review: 195\n", 1)

marker = "excluded_from_canonical_review:\n"
verified = ["Google OAuth", "Fail-closed", "Authorization", "SSL/TLS", "Supabase Google OAuth", "HF_TOKEN"]
inventory = inventory.replace(marker, "".join(f"- {term}\n" for term in verified) + marker, 1)

old_security = "  security:\n  - Google OAuth\n  - Fail-closed\n  - Authorization\n  - SSL/TLS\n  - Network restriction\n  - Supabase Google OAuth\n  - HF_TOKEN\n"
if old_security not in inventory:
    raise SystemExit("expected security review block not found")
inventory = inventory.replace(old_security, "  security:\n  - Network restriction\n", 1)
INVENTORY.write_text(inventory, encoding="utf-8")
