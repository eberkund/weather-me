import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
  assetsInclude: ["**/*.{jpg,svg}"],
  plugins: [react(), tsconfigPaths()],
});
