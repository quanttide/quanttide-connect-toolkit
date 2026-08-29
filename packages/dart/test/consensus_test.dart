import 'package:quanttide_connect/quanttide_connect.dart';
import 'package:test/test.dart';

void main() {
  test('parses a consensus from the specification shape', () {
    final consensus = Consensus.fromJson({
      'id': 'c1',
      'title': '采用 SQLite',
      'description': 'Provider 使用 SQLite 持久化共识。',
      'created_at': '2026-08-29T00:00:00Z',
      'updated_at': '2026-08-29T00:05:00Z',
    });

    expect(consensus.title, '采用 SQLite');
    expect(consensus.description, 'Provider 使用 SQLite 持久化共识。');
    expect(consensus.updatedAt, DateTime.parse('2026-08-29T00:05:00Z'));
  });

  test('constructs consensus detail route', () {
    expect(consensusPath('c1'), '/consensuses/c1');
  });
}
