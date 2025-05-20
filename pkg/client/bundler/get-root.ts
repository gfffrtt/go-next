import path from "path";
import fs from "fs";

export const getRoot = () => {
  let currentDir = process.cwd();
  while (currentDir !== path.parse(currentDir).root) {
    if (fs.existsSync(path.join(currentDir, "package.json"))) {
      return currentDir;
    }
    currentDir = path.dirname(currentDir);
  }
  throw new Error("Could not find project root (package.json)");
};
