<template>
  <div class="app-container">
    <!-- 表格 -->
    <vxe-grid ref="xGridDom" v-bind="xGridOpt">
      <!-- 左侧按钮列表 -->
      <template #toolbar-btns>
        <vxe-button status="primary" icon="vxe-icon-add" @click="toRoute('CreateLicense')">创建license</vxe-button>
      </template>

      <!--      展开行的内容-->
      <template #expand-content="{ row }">
        <div>
          <el-card>
            <el-descriptions>
              <el-descriptions-item v-for="item in licenseDescribe" :label="item.label">
                <el-tag>{{ row[item.key] }}</el-tag>
              </el-descriptions-item>
            </el-descriptions>

            <el-descriptions title="集成组件黑名单">
              <el-descriptions-item>
                <span v-if="row.studioCompBlackList !== ''" class="tooltip-content">
                  {{ row.studioCompBlackList }}
                </span>
                <el-tag v-else>N/A</el-tag>
              </el-descriptions-item>
            </el-descriptions>
          </el-card>
          <!--          权限设置 -->
          <Permission :readonly="true" :propCheckList="row.checkList" />
        </div>
      </template>

      <!--      状态-->
      <template #row-status="{ row }">
        <div>
          <el-tag v-if="row.active === 1">生效中</el-tag>
          <el-tag v-else type="danger">已作废</el-tag>
        </div>
      </template>

      <!-- License -->
      <template #row-license="{ row }">
        <div v-if="row.license">
          <span>*****************</span>
          <el-tooltip class="box-item" effect="dark" content="点我复制 license" placement="right">
            <el-button
              icon="DocumentCopy"
              size="small"
              text
              v-clipboard="row.license"
              @click="ElMessage.success('已复制到剪贴板')"
            >
            </el-button>
          </el-tooltip>
        </div>
        <div v-else>
          <el-text class="mx-1" type="primary">license 为空</el-text>
        </div>
      </template>

      <!-- 操作 -->
      <template #row-operate="{ row }">
        <el-button link icon="Delete" type="danger" @click="crudStore.onDelete(row)">删除</el-button>
      </template>
    </vxe-grid>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue"
import { type ElMessageBoxOptions, ElMessageBox, ElMessage } from "element-plus"
import DaysColumnSolts from "./tsx/DaysColumnSolts"
import { formatDate, formatDateTime } from "@/utils/index"
import Permission from "../components/permissionSetting/index.vue"

import { type VxeGridInstance, type VxeGridProps, type VxeGridPropTypes } from "vxe-table"

import { useRouter } from "vue-router"
import {deleteLicenseApi, getLicenseApi, GetLicenseResponseData, RowMeta} from "@/api/license/license"

defineOptions({
  name: "License"
})

// 展开的license详情
const licenseDescribe = [
  { label: "管理平台API数", key: "managerApis" },
  { label: "管理平台用户数", key: "managerUsers" },
  { label: "管理平台项目数", key: "managerProjects" },
  { label: "管理平台集成平台限制", key: "managerIntegrations" },
  { label: "管理平台支持组织数", key: "managerOrganizations" },
  { label: "管理平台测试限制", key: "managerTests" },
  { label: "管理平台发布限制", key: "managerProducts" },
  { label: "集成平台用户数", key: "studioUsers" }
]

// 把权限设置的数据恢复为复选框可用的数据
const restoreDefaultCheckList = (processedData: Record<string, any>): Record<string, string[]> => {
  const restoredList: Record<string, string[]> = {}

  Object.entries(processedData).forEach(([key, value]) => {
    if (key === "actionLog") {
      // 单独处理 actionLog 属性
      restoredList.actionLog = value.display ? ["display"] : []
      return
    }

    if (key === "apiManage") {
      // 单独处理 apiManage 属性
      const projectChilds = value.childs?.project?.childs
      restoredList.apiManage = ["display"]

      if (projectChilds) {
        restoredList.apiManage = restoredList.apiManage.concat(
          Object.keys(projectChilds).filter((k) => projectChilds[k])
        )
      }

      return
    }

    if (value.display) {
      restoredList[key] = ["display"]
    } else {
      restoredList[key] = []
    }

    if (value.childs) {
      Object.entries(value.childs).forEach(([childKey, childValue]) => {
        if (childValue) {
          restoredList[key].push(childKey)
        }
      })
    }
  })

  return restoredList
}

const router = useRouter()
const toRoute = (componentName: string) => {
  router.push({ name: componentName })
}

//#region vxe-grid
const xGridDom = ref<VxeGridInstance>()
const xGridOpt: VxeGridProps = reactive({
  /** 展开行 */
  expandConfig: {
    iconColumnIndex: 0, // 展开图标所在的列索引
    expandAll: false, // 是否默认展开所有行
    expandRowKeys: [], // 默认展开的行的 key
    accordion: false // 是否每次只能展开一行
  },
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
        title: "状态",
        titleColon: true,
        field: "active",
        itemRender: {
          name: "$select",
          props: { placeholder: "请选择", clearable: true },
          options: [
            { label: "生效中", value: 1 },
            { label: "已作废", value: 2 }
          ]
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
      field: "customerName",
      title: "客户名称"
    },
    {
      field: "env",
      title: "ENV"
    },
    {
      field: "expiredDate",
      title: "过期日期",
      formatter: (params) => {
        return formatDate(params.row.expiredDate)
      }
    },
    {
      field: "expiredDate",
      title: "剩余天数",
      sortable: true,
      slots: DaysColumnSolts
    },
    {
      field: "active",
      title: "状态",
      slots: { default: "row-status" }
    },
    {
      field: "license",
      title: "License",
      slots: { default: "row-license" }
    },
    {
      field: "createdAt",
      sortable: true,
      title: "创建时间",
      formatter: (params) => {
        return formatDateTime(params.row.createdAt)
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
          const callback = (res: GetLicenseResponseData) => {
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
            active: form.active || undefined,
            size: page.pageSize,
            currentPage: page.currentPage
          }
          /** 调用接口 */
          getLicenseApi(params).then(callback).catch(callback)
        })
      }
    }
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
  /** 删除 */
  onDelete: (row: RowMeta) => {
    const tip = `确定 <strong style='color:red;'>删除 </strong>客户 <strong style='color:#409eff;'>${row.customerName}</strong> <strong style='color:#409eff;'>${row.env}的license</strong>信息？`
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
        deleteLicenseApi(row.id)
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
