# lumiシリーズ共通仕様 ディスカッション & 進捗管理

このドキュメントでは、lumiシリーズ共通仕様 に関する目的、ロードマップ、議論の経緯、進捗を記録・管理します。

---

## 目的・ゴール

- `lumitree`、`lumidive`、`lumitlink`など複数システム間で交換するデータ契約を一元化する。
- イベント候補の出所と原文を失わず、名寄せ結果を別レコードとして追跡可能にする。
- 複数のlumiサービスをpullするためのfeed envelopeを定義する。
- OpenAPIに限定せず、イベント単体で検証しやすいJSON Schemaを主契約にする。

---

## ロードマップ / タスク

| Step | 項目 | 内容 | 状況 |
| :--- | :--- | :--- | :--- |
| **Step 1** | **契約設計** | 複数システムからのイベント、エンティティ、名寄せ判定schemaを定義 | 完了（v0.1.1） |
| **Step 2** | **利用者実装** | lumi-exporterとlumitlinkから契約を利用 | 完了（実データ確認済み） |
| **Step 3** | **入力接続** | lumitree/lumidiveの出力を候補schemaへ変換 | 進行中（exporter adapterは完了） |

---

## 決定事項 (ADR一覧)

- [ADR 0001: lumiシリーズ共通イベント契約](adr/0001-common-event-contract.md)
- [ADR 0002: lumiシリーズの命名規約](adr/0002-lumi-naming-convention.md)

## 現在地 (2026-10-12)

- Entity、Event、Feed、Matchの共通型とJSON Schemaを公開している。
- `lumi-exporter`はv0.1.1を利用し、3つの公開lumiサービスをpullして正規化できる。
- `lumitlink`は`/api/v1/entities/amamiyasaku`で`lumi.entity.v1`を返す。
- Organizationは一覧性だけを理由に急いで作らず、個人プロフィールREADMEとGitHub Projectで代替する。
- 共同開発、権限管理、独立ブランド化が必要になった段階でOrganization移行を再検討する。

---

## 直近のネクストアクション

- [ ] schemaのJSON Schema検証をCIに追加する
- [x] lumi-exporterからGo型をimportして利用する
- [x] lumi-exporter側にlumitree/lumidiveの各フィールド変換を追加する
- [ ] lumitree/lumidive自身からfeed envelopeを返す契約を追加する
