import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    coverage: {
      reporter: ['text', 'json-summary', 'json', 'html'],
      reportOnFailure: true,
    },
    globals: true,
    environment: 'jsdom',
    setupFiles: './vitest.setup.ts',
  },
});
