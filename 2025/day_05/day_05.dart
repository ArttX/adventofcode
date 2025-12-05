import "dart:io";

void main() async {
  final file = File("day_05/input.txt");
  final input = await file.readAsLines();

  final sepIdx = input.indexOf("");
  final ranges = input.sublist(0, sepIdx);
  final products = input.sublist(sepIdx + 1).map(int.parse).toList();

  var freshCount = 0;
  for (final product in products) {
    for (final range in ranges) {
      final parts = range.split("-");
      if (int.parse(parts[0]) <= product && product <= int.parse(parts[1])) {
        freshCount++;
        break;
      }
    }
  }

  print("Fresh count: ${freshCount}");
}
