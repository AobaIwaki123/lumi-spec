# 02: lumiシリーズの一覧性とOrganization移行判断

## 1. 背景と現象の観察

`lumi`シリーズのリポジトリが増えると、個人アカウント配下でも一覧性を確保したくなる。一方で、Organizationへの移行には、リポジトリURL、Go module、GHCR、Argo CD、OpenAPI参照先などの更新が発生する。

## 2. 構造分析 / 判断

Organizationは会社用に限定されず、プロジェクト群の名前空間として利用できる。ただし、今回の主目的は権限管理や共同開発ではなく、自分自身がLumi関連リポジトリを一覧しやすくすることである。

この目的だけなら、現時点でOrganizationへ移行する必要はない。個人配下のまま、次の構成で同じ目的を達成できる。

```text
AobaIwaki123/AobaIwaki123
├── README.md
└── docs/
    ├── lumi-architecture.svg
    ├── lumi-architecture.drawio
    └── lumi-services.yaml
```

- 個人プロフィールREADMEにLumiシリーズの一覧と入口を置く。
- GitHub Projectで複数リポジトリを横断して進捗を確認する。
- `lumi-services.yaml`をサービス一覧の基準データにし、アーキテクチャ図を生成する。
- 公開サービス、共通仕様、運用ツールを表形式で整理する。

## 3. アクションプロトコル / 移行条件

当面は個人アカウント配下で運用し、次の条件のいずれかを満たしたらOrganizationを再検討する。

- Lumi関連リポジトリが5〜10個程度になる。
- 共同開発者、Bot、外部コントリビューターを招待する。
- Organization単位の権限、Secrets、Actionsポリシーが必要になる。
- Lumiを独立した公開ブランドとして管理する。

移行する場合は、依存の中心である`lumi-spec`、利用者である各サービス、`lumi-exporter`の順に、リポジトリ転送後にcanonical URLを更新する。旧GitHub URLのリダイレクトに依存し続けず、Go module、GHCRイメージ、Argo CD、OpenAPI `$ref`、READMEを新しいOrganization名へ揃える。

---

## 参考文献 (Authoritative References)

- [GitHub: Transferring a repository](https://docs.github.com/en/repositories/creating-and-managing-repositories/transferring-a-repository)
- [GitHub: Customizing your organization's profile](https://docs.github.com/en/organizations/collaborating-with-groups-in-organizations/customizing-your-organizations-profile)
- [GitHub: Managing your profile README](https://docs.github.com/en/account-and-profile/how-tos/profile-customization/managing-your-profile-readme)
