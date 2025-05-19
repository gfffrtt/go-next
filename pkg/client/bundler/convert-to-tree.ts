import { APP } from "./constants";

export type Tree = {
  root: string;
  clientComponents: string[];
  nodes: Tree[];
};

export const convertToTree = (
  scripts: { root: string; clientComponents: string[] }[]
) => {
  const layouts = scripts.filter((script) => script.root.endsWith("layout.go"));
  const pages = scripts.filter((script) => script.root.endsWith("page.go"));

  const layoutMap = new Map<string, Tree>();

  const rootLayout = layouts.find(
    (layout) => layout.root === `${APP}/layout.go`
  );
  if (!rootLayout) return [];

  const tree: Tree = {
    root: rootLayout.root,
    clientComponents: rootLayout.clientComponents,
    nodes: [],
  };
  layoutMap.set("/", tree);

  layouts.forEach((layout) => {
    if (layout.root === `${APP}/layout.go`) return;

    const routePath = layout.root.replace(APP, "").replace("/layout.go", "");
    const layoutNode = {
      root: layout.root,
      clientComponents: layout.clientComponents,
      nodes: [],
    };
    layoutMap.set(routePath, layoutNode);
  });

  pages.forEach((page) => {
    const routePath = page.root.replace(APP, "").replace("/page.go", "");
    const routeParts = routePath.split("/").filter(Boolean);

    let currentPath = routeParts.join("/");
    let parentLayout = null;

    while (currentPath !== "") {
      if (layoutMap.has("/" + currentPath)) {
        parentLayout = layoutMap.get("/" + currentPath);
        break;
      }
      routeParts.pop();
      currentPath = routeParts.join("/");
    }

    parentLayout = parentLayout || layoutMap.get("/");

    if (parentLayout) {
      parentLayout.nodes.push({
        root: page.root,
        clientComponents: page.clientComponents,
        nodes: [],
      });
    }
  });

  layoutMap.forEach((layout, path) => {
    if (path === "/") return;

    const routeParts = path.split("/").filter(Boolean);
    routeParts.pop();
    let currentPath = routeParts.join("/");
    let parentLayout = null;

    while (currentPath !== "") {
      if (layoutMap.has("/" + currentPath)) {
        parentLayout = layoutMap.get("/" + currentPath);
        break;
      }
      routeParts.pop();
      currentPath = routeParts.join("/");
    }

    parentLayout = parentLayout || layoutMap.get("/");

    if (parentLayout && layout !== parentLayout) {
      parentLayout.nodes.push(layout);
    }
  });

  return [tree];
};
