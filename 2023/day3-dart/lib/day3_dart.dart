import 'dart:io';

import 'package:day3_dart/visual.dart';

void solve(List<String> args) async {
  var input = await File('input.txt').readAsString();
  List<List<String>> grid = [];
  Map<String, Set<int>> gears = new Map();
  for (var element in input.split("\n")) {
    grid.add(element.runes.map(String.fromCharCode).toList());
  }

  List<int> numbers = [];
  for (var y = 0; y < grid.length; y++) {
    var row = grid[y];
    var numAccum = "";
    bool hasHit = false;
    List<(int, int)> tempGears = [];
    for (var x = 0; x < row.length; x++) {
      var char = row[x];
      if (int.tryParse(char) != null) {
        numAccum += char;
        var response = checkSymbols(grid, x, y);
        if (response.$1) {
          hasHit = true;
          tempGears.addAll(response.$2);
        }
      } else {
        if (numAccum != "") {
          if (hasHit) {
            var number = int.parse(numAccum);
            numbers.add(number);
            grid.printBoundingBox(x, y, numAccum.length);
            print("");
            for (var coord in tempGears) {
              var gearLoc = "${coord.$1},${coord.$2}}";
              if (gears[gearLoc] == null) {
                gears[gearLoc] = new Set();
              }
              gears[gearLoc]?.add(number);
            }
            tempGears = [];
          }
        }
        hasHit = false;
        numAccum = "";
      }
    }
  }
  print(
      "part 1: ${numbers.reduce((previousValue, element) => previousValue + element)}");
  var validGears = gears.values.where((element) => element.length == 2);
  var total = 0;
  for (var list in validGears) {
    total += list.reduce((previousValue, element) => previousValue * element);
  }

  print("part 2: ${total}");
}

(bool, List<(int, int)>) checkSymbols(List<List<String>> grid, int x, int y) {
  List<(int, int)> coords = [];

  void Function(int, int) isGear = (int x, int y) {
    coords.add((x, y));
  };

  return (
    (scan(grid, x, y + 1, isGear) ||
        scan(grid, x - 1, y, isGear) ||
        scan(grid, x + 1, y, isGear) ||
        scan(grid, x, y - 1, isGear) ||
        scan(grid, x + 1, y + 1, isGear) ||
        scan(grid, x - 1, y + 1, isGear) ||
        scan(grid, x + 1, y - 1, isGear) ||
        scan(grid, x - 1, y - 1, isGear)),
    coords
  );
}

bool scan(
    List<List<String>> grid, int x, int y, void Function(int, int) isGear) {
  var hasHit = false;
  if (isValidIndex(grid, x, y)) {
    var char = grid[y][x].trim();
    if (char.isEmpty) false;
    if (char == "*") {
      isGear(x, y);
      return true;
    }
    return !"0123456789.".contains(char);
  }
  return hasHit;
}

bool isValidIndex(List<List<String>> grid, int x, int y) {
  return y >= 0 && y < grid.length && x >= 0 && x < grid[y].length;
}
