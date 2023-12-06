const fs = require("fs");
const path = require("path");

let seedRange = [];

const maps = {
  seedtosoil: [],
  soiltofertilizer: [],
  fertilizertowater: [],
  watertolight: [],
  lighttotemperature: [],
  temperaturetohumidity: [],
  humiditytolocation: [],
};

const main = async () => {
  const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf8");
  let currentKey = "";

  input.split("\n").forEach((line) => {
    if (line.trim() !== "") {
      if (line.startsWith("seeds:")) {
        const all = line
          .split(":")[1]
          .trim()
          .split(" ")
          .map((value) => Number(value.trim()));
        for (let index = 0; index < all.length; index += 2) {
          seedRange.push({
            start: all[index],
            end: all[index + 1],
          });
        }
      } else if (line.includes(":")) {
        const key = line.split(" ")[0].trim().split("-").join("");
        currentKey = key;
      } else {
        const [destination, source, range] = line
          .split(" ")
          .map((value) => Number(value.trim()));
        maps[currentKey].push({
          destination,
          source,
          range,
        });
      }
    }
  });
  const values = [];
  seedRange.some((range) => {
    console.log(range);
    return true;
  });

  console.log(Math.min(...values));
};

main();
