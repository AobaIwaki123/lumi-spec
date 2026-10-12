# ADR 0001: lumiシリーズ共通イベント契約

## ステータス

提案

## 背景

`lumitree`はTimeTree由来のイベントを返し、`lumidive`は別の取得・配信系統を持つ。`lumi-exporter`はこれら複数ソースの候補を名寄せするため、システムごとにイベントJSONを持つと、日付・出所・原文の扱いがずれ、判定の再現やレビューが難しくなる。

`lumitlink`のOpenAPIはプロフィール取得APIの契約であり、イベント名寄せの入力型をそこへ混ぜると、Lit.link固有APIとlumiシリーズ共通契約の責務が混ざる。

## 決定

共通契約を`lumi-spec`に分離し、JSON Schemaを正とする。

- `event-candidate.v1`: ソースごとの正規化イベント。`source_ref`、原文タイトル、`local_date`を必須にする。
- `match-decision.v1`: `matched` / `ambiguous` / `unmatched` と候補別スコア、アラート理由を表す。
- `lumitree`、`lumidive`などの入力アダプターは、各システム固有のIDやURLを`source_ref`に保持する。
- 破壊的変更は新しいschema versionを発行する。
- Go利用者向けに`event`と`match`の型を提供するが、他言語利用者はJSON Schemaを直接利用できる。
- `lumitlink`はプロフィールAPIのOpenAPIを維持し、将来イベント情報を提供するときに`event-candidate.v1`へ変換する。

## 捨てたものとリスク

OpenAPIだけを共通形式にしなかった。`lumi-exporter`の初期PoCはCLIであり、HTTPサーバーの有無に依存せずイベント単体を検証する必要があるためである。将来HTTP入出力を追加する場合は、このJSON SchemaをOpenAPIのcomponentsへ取り込む。

`local_date`はタイムゾーン変換後の値を入力アダプターが責任を持って作る。変換を名寄せ器内で暗黙に行わないことで、日付境界の誤判定を追跡できる反面、入力アダプターの品質に依存する。

参考:

- [JSON Schema 2020-12](https://json-schema.org/draft/2020-12/json-schema-core)
- [Unicode Standard Annex #15](https://unicode.org/reports/tr15/)
- [Fellegi and Sunter, A Theory for Record Linkage](https://doi.org/10.1080/01621459.1969.10501049)
