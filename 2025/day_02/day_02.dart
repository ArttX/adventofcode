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
      if (len % 2 != 0) continue;
      if (numStr.substring(0, (len / 2).toInt()) ==
          numStr.substring((len / 2).toInt())) {
        sum += i;
      }
    }
  }

  print(sum);
}
