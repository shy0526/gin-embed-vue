// https://vitejs.dev/config/
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import yaml from 'js-yaml'
import fs from 'fs'
import path from 'path'

interface Config {
  static_url_prefix?: string;
}

export default defineConfig(() => {
  const configPath = path.resolve(__dirname, '../config.yaml')
  if (! fs.existsSync(configPath)) {
    throw new Error('config.yaml not found')
  }

  const config = yaml.load(fs.readFileSync(configPath, 'utf8')) as Config
  if (! config.static_url_prefix) {
      throw new Error('config.yaml must have static_url_prefix')
  }
  const staticBaseUrl = config.static_url_prefix
  console.log(`staticBaseUrl: ${staticBaseUrl}`)

  return {
    plugins: [vue()],
    base: staticBaseUrl,
    build: {
      outDir: '../static',
      emptyOutDir: true,
    },
  }
})
