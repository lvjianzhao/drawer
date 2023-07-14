<template>
  <div class="app-container">
    <!-- 表格 -->
    <vxe-grid ref="xGridDom" v-bind="xGridOpt">
      <!-- 左侧按钮列表 -->
      <template #toolbar-btns>
        <vxe-button status="primary" icon="vxe-icon-add" @click="crudStore.onShowModal()">新增事件</vxe-button>
      </template>

      <template #row-name="{ row }">
        <div>
          <template v-if="row.description">
            <el-tooltip class="box-item" effect="dark" placement="right">
              <template #content>
                <div class="tooltip-content">
                  {{ row.description }}
                </div>
              </template>
              {{ row.name }}
            </el-tooltip>
          </template>
          <template v-else>
            {{ row.name }}
          </template>
        </div>
      </template>

      <!-- 操作 -->
      <template #row-operate="{ row }">
        <el-button type="primary" text icon="Edit" size="small" @click="crudStore.onShowModal(row)">编辑 </el-button>
        <el-button link icon="Delete" type="danger" @click="crudStore.onDelete(row)">删除</el-button>
      </template>
    </vxe-grid>
    <!-- 弹窗 -->
    <vxe-modal ref="xModalDom" v-bind="xModalOpt">
      <!-- 表单 -->
      <vxe-form ref="xFormDom" v-bind="xFormOpt"> </vxe-form>
    </vxe-modal>
  </div>
</template>

<script lang="ts" setup>
import { nextTick, reactive, ref } from "vue"
import { type ElMessageBoxOptions, ElMessageBox, ElMessage } from "element-plus"
import { shortcuts } from "@/config/dateDict"
import { formatDateTime, formatTimestamp } from "@/utils/index"

import {
  type VxeGridInstance,
  type VxeGridProps,
  type VxeModalInstance,
  type VxeModalProps,
  type VxeFormInstance,
  type VxeFormProps,
  type VxeGridPropTypes,
  type VxeFormDefines
} from "vxe-table"

import { selectOption } from "@/api/customer"
import { getAllCustomer } from "@/utils/customer"
import {
  addEventApi,
  deleteCustomerEventApi,
  EventTableData,
  getCustomerEventApi, RowMeta,
  updateEventApi
} from "@/api/customer/event"

defineOptions({
  name: "Event"
})

const customerOptions = ref<selectOption[]>([])
getAllCustomer().then((Options) => {
  if (Options) {
    customerOptions.value = Options || []
  }
})

//#region vxe-grid


const xGridDom = ref<VxeGridInstance>()
const xGridOpt: VxeGridProps = reactive({
  loading: true,
  autoResize: true,
  /** 分页配置项 */
  pagerConfig: {
    align: "right"
  },
  /** 表单配置项 */
  formConfig: {
    items: [
      {
        field: "customerName",
        itemRender: {
          name: "$input",
          props: { placeholder: "客户名称", clearable: true }
        }
      },
      {
        field: "eventKeyword",
        itemRender: {
          name: "$input",
          props: { placeholder: "事件关键字", clearable: true }
        }
      },
      {
        field: "datetimeRange",
        itemRender: {
          name: "ElDatePicker",
          props: {
            type: "datetimerange",
            valueFormat: "YYYY-MM-DD HH:mm:ss",
            startPlaceholder: "起始日期",
            endPlaceholder: "截止日期",
            unlinkPanels: true,
            shortcuts: shortcuts
          }
        }
      },
      {
        itemRender: {
          name: "$buttons",
          children: [
            {
              props: { type: "submit", content: "搜索", status: "primary" }
            }
          ]
        }
      }
    ]
  },
  /** 工具栏配置 */
  toolbarConfig: {
    refresh: true,
    custom: false,
    slots: { buttons: "toolbar-btns" }
  },
  /** 列配置 */
  columns: [
    {
      field: "id",
      title: "ID",
      width: "100px"
    },
    {
      field: "name",
      title: "客户名称",
      slots: { default: "row-name" }
    },
    {
      field: "eventTime",
      title: "事件发生时间",
      formatter: (params) => {
        return formatDateTime(params.row.eventTime)
      }
    },
    {
      field: "eventContent",
      title: "事件内容"
    },
    {
      field: "updatedAt",
      sortable: true,
      title: "更新时间",
      formatter: (params) => {
        return formatDateTime(params.row.updatedAt)
      }
    },
    {
      title: "操作",
      width: "250px",
      fixed: "right",
      showOverflow: false,
      slots: { default: "row-operate" }
    }
  ],
  /** 数据代理配置项（基于 Promise API） */
  proxyConfig: {
    /** 启用动态序号代理 */
    seq: true,
    /** 是否代理表单 */
    form: true,
    /** 是否自动加载，默认为 true */
    // autoLoad: false,
    props: {
      total: "total"
    },
    ajax: {
      query: ({ page, form }: VxeGridPropTypes.ProxyAjaxQueryParams) => {
        xGridOpt.loading = true
        crudStore.clearTable()
        return new Promise<any>((resolve: Function) => {
          let total = 0
          let result: RowMeta[] = []
          /** 加载数据 */
          const callback = (res: EventTableData) => {
            if (res && res.data) {
              const resData = res.data
              // 总数
              if (Number.isInteger(resData.total)) {
                total = resData.total
              }
              // 分页数据
              if (Array.isArray(resData.list)) {
                result = resData.list
              }
            }
            xGridOpt.loading = false
            resolve({ total, result })
          }

          let startTimestamp
          let endTimestamp

          if (
            typeof form.datetimeRange === "object" &&
            form.datetimeRange !== null &&
            Object.keys(form.datetimeRange).length === 2
          ) {
            let s = form.datetimeRange[0]
            let e = form.datetimeRange[1]
            startTimestamp = formatTimestamp(s)
            endTimestamp = formatTimestamp(e)
          }
          /** 接口需要的参数 */
          const params = {
            customerName: form.customerName || undefined,
            eventKeyword: form.eventKeyword || undefined,
            startTime: startTimestamp || undefined,
            endTime: endTimestamp || undefined,
            size: page.pageSize,
            currentPage: page.currentPage
          }
          /** 调用接口 */
          getCustomerEventApi(params).then(callback).catch(callback)
        })
      }
    }
  }
})
//#endregion

//#region vxe-modal
const xModalDom = ref<VxeModalInstance>()
const xModalOpt: VxeModalProps = reactive({
  title: "",
  width: "700px",
  height: "600px",
  showClose: true,
  escClosable: true,
  maskClosable: false,
  beforeHideMethod: () => {
    xFormDom.value?.clearValidate()
    return Promise.resolve()
  }
})
//#endregion

//#region vxe-form
const xFormDom = ref<VxeFormInstance>()
const xFormOpt = reactive<VxeFormProps>({
  span: 20,
  titleWidth: "30%",
  loading: false,
  /** 是否显示标题冒号 */
  titleColon: false,
  titleAlign: "right",
  /** 表单数据 */
  data: {
    id: "",
    customerID: "",
    eventContent: "",
    eventTime: ""
  },
  /** 项列表 */
  items: [
    {
      field: "customerID",
      title: "客户名称",
      itemRender: {
        name: "$select",
        props: {
          placeholder: "请选择",
          options: customerOptions,
          filterable: true
        }
      }
    },
    {
      field: "eventTime",
      title: "事件发生时间",
      itemRender: {
        name: "ElDatePicker",
        props: {
          type: "datetime",
          valueFormat: "YYYY-MM-DD HH:mm:ss",
          placeholder: "请选择",
          unlinkPanels: true
        }
      }
    },
    {
      field: "eventContent",
      title: "事件详情",
      itemRender: {
        name: "$textarea",
        props: {
          placeholder: "请输入",
          autosize: { minRows: 13, maxRows: 13 }
        }
      }
    },
    {
      align: "right",
      itemRender: {
        name: "$buttons",
        children: [
          { props: { content: "取消" }, events: { click: () => xModalDom.value?.close() } },
          {
            props: { type: "submit", content: "确定", status: "primary" },
            events: { click: () => crudStore.onSubmitForm() }
          }
        ]
      }
    }
  ],
  /** 校验规则 */
  rules: {
    customerID: [
      {
        required: true,
        validator: ({ itemValue }) => {
          if (itemValue === "") {
            return new Error("请选择客户")
          }
        }
      }
    ]
  }
})
//#endregion

//#region CRUD
const crudStore = reactive({
  /** 表单类型：修改：true 新增：false */
  isUpdate: true,
  /** 加载表格数据 */
  commitQuery: () => xGridDom.value?.commitProxy("query"),
  /** 清空表格数据 */
  clearTable: () => xGridDom.value?.reloadData([]),
  /** 点击显示弹窗 */
  onShowModal: (row?: RowMeta) => {
    if (row) {
      crudStore.isUpdate = true
      xModalOpt.title = "编辑事件"
      // 赋值
      xFormOpt.data.id = row.id
      xFormOpt.data.customerID = row.customerID
      xFormOpt.data.eventContent = row.eventContent
      xFormOpt.data.eventTime = formatDateTime(row.eventTime)
    } else {
      xFormOpt.data.id = 0
      crudStore.isUpdate = false
      xModalOpt.title = "新增事件"
    }
    xModalDom.value?.open()
    nextTick(() => {
      !crudStore.isUpdate && xFormDom.value?.reset()
      xFormDom.value?.clearValidate()
    })
  },
  /** 确定并保存 */
  onSubmitForm: () => {
    if (xFormOpt.loading) return
    xFormDom.value?.validate(async (errMap?: VxeFormDefines.ValidateErrorMapParams) => {
      if (errMap) return
      xFormOpt.loading = true
      const callback = (err?: any) => {
        xFormOpt.loading = false
        if (err) return
        xModalDom.value?.close()
        !crudStore.isUpdate && crudStore.afterInsert()
        crudStore.commitQuery()
      }
      try {
        let apiFunction
        if (crudStore.isUpdate) {
          apiFunction = updateEventApi
        } else {
          apiFunction = addEventApi
        }

        const res = await apiFunction({
          id: xFormOpt.data.id,
          customerID: xFormOpt.data.customerID,
          eventTime: xFormOpt.data.eventTime,
          eventContent: xFormOpt.data.eventContent
        })

        ElMessage({ type: "success", message: res.msg })
      } catch (error) {
        console.log(error)
      } finally {
        callback()
      }
    })
  },

  /** 新增后是否跳入最后一页 */
  afterInsert: () => {
    const pager: VxeGridPropTypes.ProxyAjaxQueryPageParams | undefined = xGridDom.value?.getProxyInfo()?.pager
    if (pager) {
      const currTotal: number = pager.currentPage * pager.pageSize
      if (currTotal === pager.total) {
        ++pager.currentPage
      }
    }
  },
  /** 删除 */
  onDelete: (row: RowMeta) => {
    const tip = `确定 <strong style='color:red;'>删除</strong> 客户 <strong style='color:#409eff;'>${row.name}</strong> ？`
    const config: ElMessageBoxOptions = {
      type: "warning",
      showClose: true,
      closeOnClickModal: true,
      closeOnPressEscape: true,
      cancelButtonText: "取消",
      confirmButtonText: "确定",
      dangerouslyUseHTMLString: true
    }
    ElMessageBox.confirm(tip, "提示", config)
      .then(() => {
        deleteCustomerEventApi(row.id)
          .then(() => {
            ElMessage.success("删除成功")
            crudStore.afterDelete()
            crudStore.commitQuery()
          })
          .catch(() => 1)
      })
      .catch(() => 1)
  },
  /** 删除后是否返回上一页 */
  afterDelete: () => {
    const tableData: RowMeta[] = xGridDom.value!.getData()
    const pager: VxeGridPropTypes.ProxyAjaxQueryPageParams | undefined = xGridDom.value?.getProxyInfo()?.pager
    if (pager && pager.currentPage > 1 && tableData.length === 1) {
      --pager.currentPage
    }
  }
})
//#endregion
</script>

<style lang="scss" scoped>
.expand-wrapper {
  padding: 10px;
}

.tooltip-content {
  white-space: pre-wrap; /* 或者使用 pre-line */
  word-break: break-all; /* 根据需要使用 word-break 属性 */
}
</style>
