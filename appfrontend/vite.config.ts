import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'
import tailwindcss from '@tailwindcss/vite'
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';
function pathResolve(dir: string): string {
  return resolve(__dirname, dir)
}

const alias: Record<string, string> = {
    '/@': pathResolve('./src/'),
    '@views': pathResolve('./src/views'),
    'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
}

export default defineConfig({
  plugins: [
    vue(),
    tailwindcss(),
    Components({
      resolvers: [
        AntDesignVueResolver({
          importStyle: false, // css in js
        }),
      ],
    }),
  ],
  resolve: {
    alias,
  }
})