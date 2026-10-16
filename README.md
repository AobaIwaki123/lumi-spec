# lumi-spec

`lumitree`、`lumidive`、`lumitlink`、`lumi-exporter`などのlumiシリーズで共有するデータ契約です。

## 契約

| 契約 | 用途 |
|---|---|
| `schema/event-candidate.v1.json` | 各システムから正規化したイベント候補 |
| `schema/event-feed.v1.json` | イベント候補をHTTPでpullするためのレスポンス包み |
| `schema/entity.v1.json` | アイドル、グループ、会場、主催者の名前・アイコン・ロゴ・SNS |
| `schema/entity-feed.v1.json` | エンティティをHTTPでpullするためのレスポンス包み |
| `schema/match-decision.v1.json` | 名寄せの判定、候補スコア、レビューアラート |

JSON Schemaを正とし、Go型は`event`、`entity`、`match`パッケージで提供します。

命名は、直接使うサービスを`lumitree`のように直結、共通契約・開発者向けツールを`lumi-spec`のようにハイフン付きとする方針です。詳細は[ADR 0002](adr/0002-lumi-naming-convention.md)を参照してください。

## 方針

- 元ソースのID・URL・原文を捨てない。共通の出所IDは`source_ref.record_id`で表す
- `source_ref.url`は元レコードの出典URL、イベントに紐づくチケット等のURLは`links`で表す
- `links.kind`は`ticket`、`official`、`related`、`social`を初期値とし、未知の種別も文字列として保持できる
- `links.origin`には`button`、`description`、`structured_url`など、抽出元を記録できる
- 日付は入力アダプターがタイムゾーン変換した`local_date`を渡す
- アイコン・ロゴ・SNSは`entity.v1`のアセットとソーシャルリンクで表す
- 共通コアに入らないソース固有値は`extensions`に保持する。ただし連携ロジックは共通フィールドだけに依存する
- 同率1位やマージン不足は`ambiguous`として自動統合しない
- 破壊的な変更はschema versionを上げる

## 検証

```bash
go test ./...
```
