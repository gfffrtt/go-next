import { APP } from "./constants";
import type { Tree } from "./convert-to-tree";
import fs from "fs/promises";
import { constants } from "fs";
import { build } from "vite";
import { getRoot } from "./get-root";

const CHUNKS_DIR = "build/chunks";
const ROUTER_DIR = "build/router";

type Chunk = {
  route: string;
  clientComponents: string[];
};

export const buildChunks = async (trees: Tree[]) => {
  const chunks: Chunk[] = [];

  const traverse = (node: Tree, clientComponents: string[] = []) => {
    const { root, clientComponents: nodeClientComponents, nodes } = node;

    const currentClientComponents = [
      ...clientComponents,
      ...nodeClientComponents,
    ];

    if (root.endsWith("page.go")) {
      const chunk: Chunk = {
        route: root.replace("src/app", "").replace("/page.go", "") || "/",
        clientComponents: currentClientComponents,
      };
      chunks.push(chunk);
    }

    nodes.forEach((childNode) => traverse(childNode, currentClientComponents));
  };

  trees.forEach((tree) => traverse(tree));

  const entrypoints = new Map<string, string>();

  await Promise.all(
    chunks.map(async (chunk) => {
      const { route, clientComponents } = chunk;
      const scripts = clientComponents.reduce((acc, component) => {
        acc += `<script type="module" src="${getRoot()}/${APP}${component}"></script>\n`;
        return acc;
      }, "");
      if (!scripts) return;
      const formattedRoute = route === "/" ? "/index" : route;
      const chunkDir = `${CHUNKS_DIR}${formattedRoute}`;
      try {
        await fs.access(chunkDir, constants.F_OK);
        await fs.rm(chunkDir, { recursive: true });
      } catch {}
      await fs.mkdir(chunkDir, { recursive: true });
      const file = `${chunkDir}/index.html`;
      await fs.writeFile(file, scripts);
      entrypoints.set(`${formattedRoute.slice(1)}/index`, file);
    })
  );

  try {
    await fs.access(ROUTER_DIR, constants.F_OK);
    await fs.rm(ROUTER_DIR, { recursive: true });
  } catch {}

  for (const [output, input] of entrypoints) {
    await build({
      build: {
        outDir: ROUTER_DIR,
        lib: {
          entry: { [output]: input },
          name: "go-next",
        },
        emptyOutDir: false,
      },
      resolve: {
        alias: {
          "pkg/client/index": getRoot() + "/pkg/client/index",
        },
      },
      define: {
        "process.env": Object.fromEntries(
          Object.entries(process.env).filter(
            ([key]) => key.startsWith("GO_NEXT_PUBLIC_") || key === "NODE_ENV"
          )
        ),
      },
    });
  }

  await fs.rm(`${ROUTER_DIR}/build`, { recursive: true });

  return chunks;
};
