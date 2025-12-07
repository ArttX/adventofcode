import "dart:io";

void main() async {
  final file = File("day_07/input.txt");
  final input = await file.readAsLines();

  final grid = input.map((row) => row.split("")).toList();

  var beams = Set()..add(input[0].indexOf("S"));
  var splitters = 0;

  // Move beams
  for (var currRow = 1; currRow < grid.length; currRow++) {
    final newList = Set<int>();
    for (final beamIdx in beams) {
      if (grid[currRow][beamIdx] == "^") {
        splitters++;
        newList.add(beamIdx - 1);
        grid[currRow][beamIdx - 1] = "|";
        newList.add(beamIdx + 1);
        grid[currRow][beamIdx + 1] = "|";
      } else {
        newList.add(beamIdx);
        grid[currRow][beamIdx] = "|";
      }
    }
    beams = newList;
  }
  print("Splitters: $splitters");
}
