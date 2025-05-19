import fs from "fs/promises";
import { getNonBaseHtmlComponents } from "./get-non-base-html-components";
import { getImports } from "./get-imports";
import { glob } from "glob";

export const compilePath = async (
  path: string,
  memory: Map<string, string[]> = new Map<string, string[]>()
): Promise<{ root: string; clientComponents: string[] }> => {
  if (memory.has(path))
    return { root: path, clientComponents: memory.get(path) ?? [] };
  const file = await fs.readFile(path, "utf-8");
  const imports = getImports(file);
  const nonBaseHtmlComponents = getNonBaseHtmlComponents(file);
  const nestedComponents: [string, string][] = imports
    .map(
      (path) =>
        [
          path,
          nonBaseHtmlComponents.find((component) =>
            component.startsWith(path.split("/").at(-1) ?? "")
          ),
        ] as const
    )
    .filter(([, comp]) => Boolean(comp)) as [string, string][];

  const nestedComponentsPaths = await Promise.all(
    nestedComponents.map(async ([path, component]) => {
      const relativePath = path
        .split("/")
        .slice(path.split("/").indexOf("src"))
        .join("/");
      const files = await glob(`${relativePath}/**/*.go`, {
        ignore: [`${relativePath}/_*/**`],
      });
      const components = await Promise.all(
        files.map(async (file) => {
          const content = await fs.readFile(file, "utf-8");
          const func = `func ${component.slice(
            component.indexOf(".") + 1,
            -1
          )}`;
          if (!content.includes(func)) return null;
          return file;
        })
      );
      return components.filter(Boolean) as string[];
    })
  );
  const nestedFiles = nestedComponentsPaths.flat();

  const nestedFilesContent = await Promise.all(
    nestedFiles.map(async (file) => compilePath(file, memory))
  );

  const nestedFilesClientComponents = nestedFilesContent
    .flat()
    .map(({ clientComponents }) => clientComponents)
    .flat();

  const matches = [
    ...file.matchAll(/html\.Client\(\s*"[^"]+"\s*,\s*"([^"]+)"/g),
  ];

  const clientComponents = matches
    .map((match) => match[1])
    .filter(Boolean) as string[];

  return {
    root: path,
    clientComponents: Array.from(
      new Set([...clientComponents, ...nestedFilesClientComponents])
    ),
  };
};
