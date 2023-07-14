import { type App } from "vue"
import { loadElementPlus } from "./element-plus"
import { loadElementPlusIcon } from "./element-plus-icon"
import { loadVxeTable } from "./vxe-table"
import { loadMdEditor } from "./v-md-editor"
import { loadVueClipboard } from "./v-clipboard"

export function loadPlugins(app: App) {
  loadElementPlus(app)
  loadElementPlusIcon(app)
  loadVxeTable(app)
  loadMdEditor(app)
  loadVueClipboard(app)
}
