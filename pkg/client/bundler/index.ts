#!/usr/bin/env node
import { glob } from "glob";
import { APP } from "./constants";
import { convertToTree } from "./convert-to-tree";
import { compilePath } from "./compile-path";
import { buildChunks } from "./build-chunks";

export const bundle = async () => {
  const paths = glob.sync(`${APP}/**/+(page.go|layout.go)`, {
    ignore: [`${APP}/_*/**`],
  });
  const scripts = await Promise.all(paths.map((path) => compilePath(path)));
  const tree = convertToTree(scripts);
  await buildChunks(tree);
};

bundle();
