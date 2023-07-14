import {type App} from "vue"

import VMdEditor from "@kangc/v-md-editor/lib/codemirror-editor"
import "@kangc/v-md-editor/lib/style/codemirror-editor.css"
import githubTheme from "@kangc/v-md-editor/lib/theme/github.js"
import "@kangc/v-md-editor/lib/theme/style/github.css"
// @ts-ignore
import createLineNumbertPlugin from "@kangc/v-md-editor/lib/plugins/line-number/index"
// highlightjs
import hljs from "highlight.js"

// codemirror 编辑器的相关资源
import Codemirror from "codemirror"
// mode
import "codemirror/mode/markdown/markdown"
import "codemirror/mode/javascript/javascript"
import "codemirror/mode/css/css"
import "codemirror/mode/htmlmixed/htmlmixed"
import "codemirror/mode/vue/vue"
// edit
import "codemirror/addon/edit/closebrackets"
import "codemirror/addon/edit/closetag"
import "codemirror/addon/edit/matchbrackets"
// placeholder
import "codemirror/addon/display/placeholder"
// active-line
import "codemirror/addon/selection/active-line"
// scrollbar
import "codemirror/addon/scroll/simplescrollbars"
import "codemirror/addon/scroll/simplescrollbars.css"
// style
import "codemirror/lib/codemirror.css"

VMdEditor.Codemirror = Codemirror
// 支持代码高亮
VMdEditor.use(githubTheme, {
  Hljs: hljs
})
// 代码块支持行号
VMdEditor.use(createLineNumbertPlugin())

export function loadMdEditor(app: App) {
  /** markdown-editor 组件完整引入 */
  app.use(VMdEditor)
}
