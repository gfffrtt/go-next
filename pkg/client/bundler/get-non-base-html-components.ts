export const getNonBaseHtmlComponents = (file: string) => {
  const matches = [
    ...file.matchAll(
      /(?<=return\s+(?:[\s\S]*?))(?!html\.)[a-zA-Z_][a-zA-Z0-9_]*\.[a-zA-Z_][a-zA-Z0-9_]*\(\)/g
    ),
  ];
  return matches.map((match) => match[0]);
};
