<template>
  <div class="app-container">
    <!-- 表格 -->
    <vxe-grid ref="xGridDom" v-bind="xGridOpt">
      <!-- 左侧按钮列表 -->
      <template #toolbar-btns>
        <vxe-button status="primary" icon="vxe-icon-add" @click="crudStore.onShowModal()">新增客户环境</vxe-button>
      </template>

      <!--      展开行的内容-->
      <template #expand-content="{ row }">
        <markdown :text="row.deploy_report" @save-doc="saveDoc" :row="row" />
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
      <vxe-form ref="xFormDom" v-bind="xFormOpt">
        <!--        <template #upload-deployReport>-->
        <!--          <el-upload-->
        <!--            v-model:file-list="fileList"-->
        <!--            class="upload-demo"-->
        <!--            action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"-->
        <!--            multiple-->
        <!--            :on-preview="handlePreview"-->
        <!--            :on-remove="handleRemove"-->
        <!--            :before-remove="beforeRemove"-->
        <!--            :limit="3"-->
        <!--            :on-exceed="handleExceed"-->
        <!--          >-->
        <!--            <el-button type="primary">Click to upload</el-button>-->
        <!--          </el-upload>-->
        <!--        </template>-->
      </vxe-form>
    </vxe-modal>
  </div>
</template>

<script lang="ts" setup>
import { nextTick, reactive, ref } from "vue"
import { type ElMessageBoxOptions, ElMessageBox, ElMessage } from "element-plus"
import { formatDateTime } from "@/utils/index"

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

import { getAllCustomerApi, selectOption } from "@/api/customer"
import {
  addCustomerEnvApi,
  deleteCustomerEnvApi,
  getCustomerEnvApi, GetEnvResponseData,
  RowMeta,
  updateCustomerEnvApi
} from "@/api/customer/env"
import Markdown from "@/views/customer/components/markdown.vue"

defineOptions({
  name: "CustomerEnv"
})

//#region vxe-grid


const allCustomerName = ref<selectOption[]>([])
const getAllCustomer = async () => {
  try {
    const res = await getAllCustomerApi()
    allCustomerName.value = res.data.list.map((item) => {
      return {
        value: item.id,
        label: String(item.name)
      }
    })
  } catch (error) {
    console.log(error)
  }
}
getAllCustomer()

const callback = (err?: any) => {
  xFormOpt.loading = false
  if (err) return
  xModalDom.value?.close()
  !crudStore.isUpdate && crudStore.afterInsert()
  crudStore.commitQuery()
}
const saveDoc = async (doc: string, row: RowMeta) => {
  xFormOpt.data = Object.assign(row)
  xFormOpt.data.deploy_report = doc
  try {
    const res = await updateCustomerEnvApi({
      id: xFormOpt.data.id,
      customerId: xFormOpt.data.customerId,
      env: xFormOpt.data.env,
      vpn: xFormOpt.data.vpn,
      deploy_report: xFormOpt.data.deploy_report,
      description: xFormOpt.data.description
    })

    ElMessage({ type: "success", message: res.msg })
  } catch (error) {
    console.log(error)
  } finally {
    callback()
  }
}

const xGridDom = ref<VxeGridInstance>()
const xGridOpt: VxeGridProps = reactive({
  loading: true,
  autoResize: true,
  /** 展开行 */
  expandConfig: {
    iconColumnIndex: 0, // 展开图标所在的列索引
    expandAll: false, // 是否默认展开所有行
    expandRowKeys: [], // 默认展开的行的 key
    accordion: false // 是否每次只能展开一行
  },
  /** 分页配置项 */
  pagerConfig: {
    align: "right"
  },
  /** 表单配置项 */
  formConfig: {
    items: [
      {
        title: "客户",
        titleColon: true,
        field: "name",
        titlePrefix: {
          content: "支持模糊搜索"
        },
        itemRender: {
          name: "$input",
          props: { placeholder: "客户名称", clearable: true }
        }
      },
      {
        title: "主机IP",
        titleColon: true,
        field: "ip",
        titlePrefix: {
          content: "检索目标为部署报告中的IP，支持模糊搜索"
        },
        itemRender: {
          name: "$input",
          props: { placeholder: "ip", clearable: true }
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
      type: "expand", // 设置列类型为展开列
      width: 30, // 设置展开列的宽度
      slots: { content: "expand-content" }
    },
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
      field: "env",
      title: "ENV"
    },
    {
      field: "vpn",
      title: "vpn"
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
          const callback = (res: GetEnvResponseData) => {
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

          /** 接口需要的参数 */
          const params = {
            name: form.name || undefined,
            ip: form.ip || undefined,
            size: page.pageSize,
            currentPage: page.currentPage
          }
          /** 调用接口 */
          getCustomerEnvApi(params).then(callback).catch(callback)
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
  className: "textarea-drag-down-only",
  /** 表单数据 */
  data: {
    id: "",
    customerId: "请选择",
    env: "",
    vpn: "",
    deploy_report: "",
    license: "",
    description: ""
  },
  /** 项列表 */
  items: [
    {
      field: "customerId",
      title: "客户名称",
      itemRender: {
        name: "$select",
        props: {
          placeholder: "请选择",
          options: allCustomerName,
          filterable: true
        }
      }
    },
    {
      field: "env",
      title: "环境名称",
      itemRender: {
        name: "$select",
        props: {
          placeholder: "请选择",
          options: [
            { label: "生产环境", value: "生产环境" },
            { label: "测试环境", value: "测试环境" },
            { label: "开发环境", value: "开发环境" },
            { label: "UAT环境", value: "UAT环境" },
            { label: "SIT环境", value: "SIT环境" },
            { label: "POC环境", value: "POC环境" }
          ]
        }
      }
    },
    // 表单中使用插槽
    // {
    //   field: "deploy_report",
    //   title: "部署报告",
    //   slots: {
    //     default: "upload-deployReport"
    //   }
    // },
    {
      field: "vpn",
      title: "VPN信息",
      itemRender: {
        name: "$textarea",
        props: {
          placeholder: "请输入",
          autosize: { minRows: 5, maxRows: 10 }
        }
      }
    },
    {
      field: "deploy_report",
      title: "部署报告",
      itemRender: {
        name: "$textarea",
        props: {
          placeholder: "请输入",
          autosize: { minRows: 5, maxRows: 10 }
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
    customerId: [
      {
        required: true,
        validator: ({ itemValue }) => {
          if (!itemValue) {
            return new Error("请选择")
          }
        }
      }
    ],
    env: [
      {
        required: true,
        validator: ({ itemValue }) => {
          if (!itemValue) {
            return new Error("请选择")
          }
        }
      }
    ],
    // deploy_report: [
    //   {
    //     required: true,
    //     validator: ({itemValue}) => {
    //       if (!itemValue) {
    //         return new Error("至少要填写客户的管理平台访问地址以及用户名/密码")
    //       }
    //     }
    //   }
    // ],
    vpn: [
      {
        required: true,
        validator: ({ itemValue }) => {
          if (!itemValue) {
            return new Error("必须填写客户环境访问方式")
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
    console.log(row)
    if (row) {
      crudStore.isUpdate = true
      xModalOpt.title = "修改客户名称"
      // 赋值
      xFormOpt.data.id = row.id
      xFormOpt.data.customerId = row.customerId
      xFormOpt.data.env = row.env
      xFormOpt.data.vpn = row.vpn
      xFormOpt.data.deploy_report = row.deploy_report
      xFormOpt.data.license = row.license
      xFormOpt.data.description = row.description
    } else {
      xFormOpt.data.id = 0
      crudStore.isUpdate = false
      xModalOpt.title = "新增客户信息"
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

      try {
        let apiFunction
        if (crudStore.isUpdate) {
          apiFunction = updateCustomerEnvApi
        } else {
          apiFunction = addCustomerEnvApi
        }

        const res = await apiFunction({
          id: xFormOpt.data.id,
          customerId: xFormOpt.data.customerId,
          env: xFormOpt.data.env,
          vpn: xFormOpt.data.vpn,
          deploy_report: xFormOpt.data.deploy_report,
          description: xFormOpt.data.description
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
    const tip = `确定 <strong style='color:red;'>删除 </strong>客户 <strong style='color:#409eff;'>${row.name}</strong> 的<strong style='color:#409eff;'>${row.env}</strong>信息？`
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
        deleteCustomerEnvApi(row.id)
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

.tooltip-base-box {
  width: 600px;
}
</style>
