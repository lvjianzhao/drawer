import { type App } from "vue"
// https://vxetable.cn/#/table/start/install
import VXETable from "vxe-table"
// https://github.com/x-extends/vxe-table-plugin-element
import VXETablePluginElement from "vxe-table-plugin-element"
import "vxe-table-plugin-element/dist/style.css"

VXETable.use(VXETablePluginElement)

/** 全局默认参数 */
VXETable.setup({
  /** 全局尺寸 */
  size: "medium",
  /** 全局 zIndex 起始值，如果项目的的 z-index 样式值过大时就需要跟随设置更大，避免被遮挡
      不要超过1999，因为from表单中使用了DatePicker组件，它的默认zindex值是2000
   **/
  zIndex: 1500,
  /** 版本号，对于某些带数据缓存的功能有用到，上升版本号可以用于重置数据 */
  version: 0,
  /** 全局 loading 提示内容，如果为 null 则不显示文本 */
  loadingText: "数据加载...",
  table: {
    showHeader: true,
    showOverflow: "tooltip",
    showHeaderOverflow: "tooltip",
    autoResize: true,
    /*斑马线条纹*/
    stripe: false,
    border: "inner",
    // round: false,
    emptyText: "暂无数据",
    rowConfig: {
      isHover: true,
      isCurrent: true
    },
    sortConfig: {
      defaultSort: { field: "updatedAt", order: "desc" },
      orders: ["desc", "asc", null]
    },
    columnConfig: {
      resizable: false // 是否启用列宽拖动功能
    },
    align: "center",
    headerAlign: "center",
    /** 行数据的唯一主键字段名 */
    rowId: "_VXE_ID"
  },
  pager: {
    // size: "medium",
    /** 配套的样式 */
    perfect: false,
    pageSize: 20,
    pagerCount: 7,
    pageSizes: [20, 50, 100, 200],
    layouts: ["Total", "PrevJump", "PrevPage", "Number", "NextPage", "NextJump", "Sizes", "FullJump"]
  },
  modal: {
    minWidth: 500,
    minHeight: 400,
    lockView: true,
    mask: true,
    // duration: 3000,
    // marginSize: 20,
    dblclickZoom: false,
    showTitleOverflow: true,
    transfer: true,
    draggable: false
  }
})

export function loadVxeTable(app: App) {
  /** Vxe Table 组件完整引入 */
  app.use(VXETable)
}
