import { type VxeColumnPropTypes } from "vxe-table/types/column"

const solts: VxeColumnPropTypes.Slots = {
  default: ({ row, column }) => {
    const cellValue = row[column.field]
    let type = "danger"
    let value = "已过期"
    if (cellValue) {
      type = "success"
      value = "维保中"
    }
    return [<span class={`el-tag el-tag--${type} el-tag--plain`}>{value}</span>]
  }
}

export default solts
