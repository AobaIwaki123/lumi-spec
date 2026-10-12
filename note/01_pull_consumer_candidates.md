# pull連携の将来候補

## 方針

`lumi-spec`は特定サービスの実装や登録順序を持たず、共通データとHTTP feed envelopeだけを定義する。pullの登録・表示・再試行は`lumi-exporter`の責務とする。

現在の対象は`lumitlink`、`lumitree`、`lumidive`である。`lumibot`にも将来的に`entity.v1`や`event-feed.v1`を提供させられる可能性があるが、現時点では候補メモに留める。サービス固有のレスポンスが先に存在する場合は、exporter側のadapterで変換し、将来サービス側が共通feedを返す段階でadapterを薄くする。

参考:

- [lumi-exporter](https://github.com/AobaIwaki123/lumi-exporter)
- [lumitree](https://github.com/AobaIwaki123/lumitree)
- [lumidive](https://github.com/AobaIwaki123/lumidive)
- [JSON Schema 2020-12](https://json-schema.org/draft/2020-12/json-schema-core)
