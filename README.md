# lumi-spec

`lumitree`、`lumidive`、`lumitlink`、`lumi-exporter`などのlumiシリーズで共有するデータ契約です。

## 契約

| 契約 | 用途 |
|---|---|
| `schema/event-candidate.v1.json` | 各システムから正規化したイベント候補 |
| `schema/match-decision.v1.json` | 名寄せの判定、候補スコア、レビューアラート |

JSON Schemaを正とし、Go型は`event`と`match`パッケージで提供します。

## 方針

- 元ソースのID・URL・原文を捨てない
- 日付は入力アダプターがタイムゾーン変換した`local_date`を渡す
- 同率1位やマージン不足は`ambiguous`として自動統合しない
- 破壊的な変更はschema versionを上げる

## 検証

```bash
go test ./...
```
