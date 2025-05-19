export const getImports = (file: string): string[] => {
  const singleImports = [...file.matchAll(/import\s+"([^"]+)"/g)].map(
    (match) => match[1]
  );
  const blockImports = [...file.matchAll(/import\s+\(([\s\S]*?)\)/g)]
    .map((match) => match[1] || "")
    .flatMap((block) => [...block.matchAll(/"([^"]+)"/g)].map((m) => m[1]));
  return [...singleImports, ...blockImports].filter(Boolean) as string[];
};
