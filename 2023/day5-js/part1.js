const fs = require("fs");
const path = require("path");

let seeds = [];

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
        seeds = line
          .split(":")[1]
          .trim()
          .split(" ")
          .map((value) => Number(value.trim()));
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

  const values = seeds.map((seed) => {
    const stuff = Object.values(maps);
    let result = seed;

    for (const item of stuff) {
      for (const { destination, source, range } of item) {
        if (source <= result && source + range > result) {
          result = destination + result - source;
          break;
        }
      }
    }

    return result;
  });

  console.log(Math.min(...values));
};

main();
