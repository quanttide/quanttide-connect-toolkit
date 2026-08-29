# quanttide_connect

量潮沟通工程Dart SDK。

## 使用

```dart
import 'package:quanttide_connect/quanttide_connect.dart';

final consensus = Consensus.fromJson({
  'id': 'c1',
  'title': '采用 SQLite',
  'description': 'Provider 使用 SQLite 持久化共识。',
  'created_at': '2026-08-29T00:00:00Z',
});

final path = consensusPath(consensus.id);
```
