import { glob } from "glob";
import fs from "fs/promises";

const APP = "src/app";

const randomId = () => Math.random().toString(36).substring(2, 15);

const getImports = (file: string): string[] => {
  const singleImports = [...file.matchAll(/import\s+"([^"]+)"/g)].map(
    (match) => match[1]
  );
  const blockImports = [...file.matchAll(/import\s+\(([\s\S]*?)\)/g)]
    .map((match) => match[1] || "")
    .flatMap((block) => [...block.matchAll(/"([^"]+)"/g)].map((m) => m[1]));
  return [...singleImports, ...blockImports].filter(Boolean) as string[];
};

const getNonBaseHtmlComponents = (file: string) => {
  const matches = [
    ...file.matchAll(
      /(?<=return\s+(?:[\s\S]*?))(?!html\.)[a-zA-Z_][a-zA-Z0-9_]*\.[a-zA-Z_][a-zA-Z0-9_]*\(\)/g
    ),
  ];
  return matches.map((match) => match[0]);
};

const compilePath = async (
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
          path, // "github.com/gfffrtt/go-next/src/app/counter/_components"
          nonBaseHtmlComponents.find((component) =>
            component.startsWith(path.split("/").at(-1) ?? "")
          ), // "_components.Counter()"
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

const formatAsInputs = (
  scripts: { root: string; clientComponents: string[] }[]
) => {};

export const bundle = async () => {
  const paths = glob.sync(`${APP}/**/+(page.go|layout.go)`, {
    ignore: [`${APP}/_*/**`],
  });
  const scripts = await Promise.all(paths.map((path) => compilePath(path)));
  const inputs = formatAsInputs(scripts);
  return inputs;
};

const components = await bundle();
console.log(components);
