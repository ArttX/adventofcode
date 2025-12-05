import "dart:io";
import "dart:math";

void main() async {
  final file = File("day_05/input.txt");
  final input = await file.readAsLines();

  final sepIdx = input.indexOf("");
  final ranges = input.sublist(0, sepIdx);

  final mergedRanges = mergeOverlap(
    ranges.map((range) {
      final parts = range.split("-");
      return (int.parse(parts[0]), int.parse(parts[1]));
    }).toList(),
  );

  var freshCount = 0;
  for (final range in mergedRanges) {
    final count = range.$2 - range.$1 + 1;
    freshCount += count;
  }

  print("Fresh count: ${freshCount}");
}

List<(int, int)> mergeOverlap(List<(int, int)> ranges) {
  var n = ranges.length;

  ranges.sort((a, b) => a.$1 - b.$1);
  List<(int, int)> res = [];

  // Checking for all possible overlaps
  for (var i = 0; i < n; i++) {
    var start = ranges[i].$1;
    var end = ranges[i].$2;

    // Skipping already merged intervals
    if (res.length > 0 && res[res.length - 1].$2 >= end) {
      continue;
    }

    // Find the end of the merged range
    for (var j = i + 1; j < n; j++) {
      if (ranges[j].$1 <= end) {
        end = max(end, ranges[j].$2);
      }
    }
    res.add((start, end));
  }
  return res;
}
