import { type App } from "vue"
import Clipboard from "v-clipboard"

export function loadVueClipboard(app: App) {
  app.use(Clipboard)
}
