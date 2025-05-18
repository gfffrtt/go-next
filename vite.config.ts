import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  plugins: [react()],
  build: {
    outDir: "build/router",
    lib: {
      entry: { counter: "src/app/counter/_components/counter.tsx" },
      name: "go-next",
    },
  },
  define: {
    "process.env": { NODE_ENV: "development" },
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "pkg/client"),
    },
  },
});
