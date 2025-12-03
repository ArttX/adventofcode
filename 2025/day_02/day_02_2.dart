import 'dart:io';

void main() async {
  final file = File("day_02/input.txt");
  final input = await file.readAsString();

  List<(int, int)> ranges =
      input.split(",").map((str) {
        final start = int.parse(str.split("-")[0]);
        final end = int.parse(str.split("-")[1]);
        return (start, end);
      }).toList();

  int sum = 0;

  // Check ranges
  for (final range in ranges) {
    for (int i = range.$1; i <= range.$2; i++) {
      final numStr = i.toString();
      final len = numStr.length;
      final hlen = (len / 2).toInt();
      for (int y = hlen; y > 0; y--) {
        final parts = splitByLength(numStr, y);
        final isEqual = parts.every((x) => x == parts[0]);
        if (isEqual) {
          sum += i;
          break;
        }
      }
    }
  }

  print(sum);
}

List<String> splitByLength(String input, int length) {
  List<String> result = [];
  for (int i = 0; i < input.length; i += length) {
    int end = i + length;
    if (end > input.length) {
      end = input.length;
    }
    result.add(input.substring(i, end));
  }
  return result;
}
