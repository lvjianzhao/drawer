<template>
  <div>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>权限设置</span>
        </div>
      </template>
      <div class="box-row">
        <template v-for="firstClass in LICENSEENTRIES">
          <template v-for="secClass in firstClass.submenu" :key="secClass.value">
            <el-form-item :label="secClass.label">
              <el-checkbox-group v-if="!readonly" v-model="defaultCheckList[secClass.value]">
                <template v-for="per in secClass.permissions" :key="per.value">
                  <el-checkbox :label="per.value">
                    {{ per.label }}
                  </el-checkbox>
                </template>
              </el-checkbox-group>
              <el-checkbox-group v-else v-model="checkedlist[secClass.value]">
                <template v-for="per in secClass.permissions" :key="per.value">
                  <el-checkbox :label="per.value" :disabled="readonly">
                    {{ per.label }}
                  </el-checkbox>
                </template>
              </el-checkbox-group>
            </el-form-item>
          </template>
        </template>
      </div>
    </el-card>

    <el-affix position="bottom" :offset="20" v-if="!readonly">
      <el-card shadow="never">
        <el-footer style="float: right">
          <el-button type="primary" @click="onSubmit">确定</el-button>
        </el-footer>
      </el-card>
    </el-affix>
  </div>
</template>

<script lang="ts" setup>
import { LICENSEENTRIES, licenseChildKeys } from "./dict"
import { reactive } from "vue"

defineOptions({
  name: "Permission"
})

const emits = defineEmits(["submitPermission"])
const props = defineProps({
  readonly: {
    type: Boolean,
    default: false
  },
  propCheckList: {
    type: String,
    default: "{}"
  }
})

/** 默认选中的复选框 */
const defaultCheckList: Record<string, string[]> = reactive({
  apiManage: ["display", "domain", "bodyModel", "project", "apiLog", "apiSearch", "dashboard", "apiUrl"],
  integrationManage: ["display", "d2a"],
  licenseManage: ["read"],
  orgManage: ["read", "create", "edit"],
  userManage: ["read", "create", "edit", "delete", "password", "sourceManage"],
  roleManage: ["read", "create", "edit", "delete", "copy"],
  project: ["display", "d2a", "file", "api", "soap"],
  alarmManage: ["display", "alarmStatistics", "alarmUser", "alarmList", "policy"],
  portalManage: ["display", "preGrantPortal", "partnerManage"],
  partner: ["read", "create", "edit", "delete", "password"],
  envManage: ["display", "envList"],
  cmsManage: [],
  apiSafe: [],
  monitoringCenter: ["display", "overview", "call"],
  actionLog: ["display"],
  test: [],
  operationManagement: []
})

const checkedlist = JSON.parse(props.propCheckList)
const translatePermissionToObj = (): Object => {
  interface Result {
    [key: string]: {
      display: boolean
      childs?: {
        [key: string]: boolean
      }
    }
  }

  const result: Result = {
    actionLog: { display: defaultCheckList.actionLog.includes("display") || false }
  }

  Object.keys(defaultCheckList)
    .filter((k) => !["version", "actionLog", "project"].includes(k)) // actionLog/project需要放到apiManage下  所以排除
    .forEach((k) => {
      result[k] = {
        display: defaultCheckList[k].includes("display") || defaultCheckList[k].includes("read"),
        childs: {}
      }

      if (result[k].childs) {
        licenseChildKeys[k].forEach((sk) => {
          //@ts-ignore
          result[k].childs[sk] = defaultCheckList[k].includes(sk)
        })
      }

      if (result[k].childs) {
        //@ts-ignore
        delete result[k].childs.display
      }
    })

  const project = { display: result.apiManage.childs?.project, childs: {} as Record<string, boolean> }

  licenseChildKeys.project.forEach((p) => {
    project.childs[p] = defaultCheckList.project.includes(p)
  })

  //@ts-ignore
  result.apiManage.childs.project = project

  result.system = {
    display:
      result.userManage.display ||
      result.orgManage.display ||
      result.roleManage.display ||
      result.licenseManage.display,
    childs: {
      userManage: result.userManage.display,
      orgManage: result.orgManage.display,
      roleManage: result.roleManage.display,
      licenseManage: result.licenseManage.display
    }
  }

  if (result.portalManage.childs) {
    result.portalManage.childs.partnerManage = defaultCheckList.partner.includes("display")
  }

  return result
}

const onSubmit = () => {
  const obj = translatePermissionToObj()
  emits("submitPermission", obj, defaultCheckList)
}

/**  默认禁用的复选框
 const licenseDefaultDisabledKeys = {
  apiManage: ['display', 'domain', 'bodyModel', 'project', 'apiLog', 'apiSearch', 'dashboard'],
  integrationManage: ['display', 'd2a'],
  licenseManage: [],
  orgManage: [],
  userManage: [],
  roleManage: [],
  project: ['display', 'd2a', 'file', 'api', 'soap'],
  alarmManage: ['display', 'alarmStatistics', 'alarmUser', 'alarmList', 'policy'],
  portalManage: ['display', 'preGrantPortal', 'partnerManage'],
  partner: [],
  envManage: ['display', 'envList'],
  cmsManage: [],
  apiSafe: [],
  monitoringCenter: ['display', 'overview', 'call'],
  actionLog: ['display'],
  test: [],
  operationManagement: []
}
 */
</script>

<style lang="scss" scoped>
.box-row {
  margin-left: 5%;
  margin-right: 5%;
}
</style>
