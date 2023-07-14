// 初始值全部为false的完整的license
import { reactive } from "vue"

export const defaultLicense = {
  apiManage: {
    display: false,
    childs: {
      apiUrl: false,
      domain: false,
      bodyModel: false,
      project: {
        display: false,
        childs: { d2a: false, file: false, integration: false, api: false, soap: false, nacos: false }
      },
      apiLog: false,
      apiSearch: false,
      dashboard: false
    }
  },
  integrationManage: { display: false, childs: { d2a: false, integration: false } },
  system: {
    display: false,
    childs: { licenseManage: false, orgManage: false, userManage: false, roleManage: false }
  },
  alarmManage: {
    display: false,
    childs: { alarmStatistics: false, alarmUser: false, alarmList: false, policy: false }
  },
  portalManage: {
    display: false,
    childs: { openPortal: false, preGrantPortal: false, partnerManage: false, transferStation: false }
  },
  menu: {},
  version: "4.0",
  envManage: { display: false, childs: { envList: false } },
  cmsManage: {
    display: false,
    childs: {
      navigation: false,
      logoSetting: false,
      material: false,
      apiDesc: false,
      carouselSetting: false,
      styleSetting: false
    }
  },
  apiSafe: {
    display: false,
    childs: {
      todayOverview: false,
      threatlist: false,
      strategic: false,
      riskmanagement: false,
      threatevent: false
    }
  },
  monitoringCenter: { display: false, childs: { overview: false } },
  actionLog: { display: false },
  roleManage: {
    display: false,
    childs: { read: false, create: false, edit: false, delete: false, copy: false }
  },
  userManage: {
    display: false,
    childs: { read: false, create: false, edit: false, delete: false, sourceManage: false }
  },
  orgManage: { display: false, childs: { read: false, edit: false } },
  licenseManage: { display: false, childs: {} }
}

// license所有的checkbox对应的keys
export const licenseChildKeys: Record<string, string[]> = {
  apiManage: ["apiUrl", "domain", "bodyModel", "project", "apiLog", "apiSearch", "dashboard"],
  integrationManage: ["d2a", "integration"],
  system: ["licenseManage", "orgManage", "userManage", "roleManage"],
  licenseManage: ["read"],
  orgManage: ["read", "create", "edit"],
  userManage: ["read", "create", "edit", "delete", "password", "sourceManage"],
  roleManage: ["read", "create", "edit", "delete", "copy"],
  project: ["d2a", "file", "integration", "api", "soap", "nacos", "zookeeper", "eureka"],
  alarmManage: ["alarmStatistics", "alarmUser", "alarmList", "policy"],
  portalManage: ["openPortal", "preGrantPortal", "partnerManage", "transferStation"],
  partner: ["read", "create", "edit", "delete", "password"],
  menu: [],
  envManage: ["envList"],
  cmsManage: ["navigation", "logoSetting", "material", "apiDesc", "carouselSetting", "styleSetting"],
  apiSafe: ["todayOverview", "threatlist", "strategic", "riskmanagement", "threatevent"],
  monitoringCenter: ["overview", "call", "relationShip", "customizeDashboard"],
  test: ["testSuite", "testPlan"],
  operationManagement: ["read", "monitor", "alert"]
}

export const LICENSEENTRIES = [
  {
    submenu: [
      {
        label: "操作记录",
        value: "actionLog",
        permissions: [{ label: "显示操作记录菜单", value: "display" }]
      }
    ]
  },
  {
    label: "告警管理",
    value: "alarmManage",
    menubar: [],
    submenu: [
      {
        label: "告警管理",
        value: "alarmManage",
        permissions: [
          { label: "显示告警管理菜单", value: "display" },
          { label: "告警列表", value: "alarmList" },
          { label: "告警统计", value: "alarmStatistics" },
          { label: "策略管理", value: "policy" },
          { label: "告警人员管理", value: "alarmUser" }
        ]
      }
    ]
  },
  {
    label: "API管理",
    value: "apiManage",
    menubar: [],
    submenu: [
      {
        label: "API管理",
        value: "apiManage",
        permissions: [
          { label: "显示项目管理菜单", value: "display" },
          { label: "API日志管理", value: "apiLog" },
          { label: "API搜索", value: "apiSearch" },
          { label: "API地址管理", value: "apiUrl" },
          { label: "Body模型管理", value: "bodyModel" },
          { label: "仪表盘", value: "dashboard" },
          { label: "域名管理", value: "domain" },
          { label: "项目管理", value: "project" }
        ]
      }
    ]
  },
  {
    label: "项目类型",
    value: "project",
    menubar: [],
    submenu: [
      {
        label: "项目类型",
        value: "project",
        permissions: [
          { label: "允许创建项目类型", value: "display" },
          { label: "D2A", value: "d2a" },
          { label: "代理已有API", value: "api" },
          { label: "文件导入", value: "file" },
          { label: "编排", value: "integration" },
          { label: "注册中心", value: "nacos" },
          { label: "SOAP", value: "soap" }
        ]
      }
    ]
  },
  {
    label: "API安全管理",
    value: "apiSafe",
    menubar: [],
    submenu: [
      {
        label: "API安全管理",
        value: "apiSafe",
        permissions: [
          { label: "显示安全管理菜单", value: "display" },
          { label: "风险处置", value: "riskmanagement" },
          { label: "策略管理", value: "strategic" },
          { label: "威胁事件统计", value: "threatevent" },
          { label: "威胁事件列表", value: "threatlist" },
          { label: "今日总览", value: "todayOverview" }
        ]
      }
    ]
  },
  {
    label: "开放平台",
    value: "cmsManage",
    menubar: [],
    submenu: [
      {
        label: "开放平台",
        value: "cmsManage",
        permissions: [
          { label: "显示开放平台菜单", value: "display" },
          { label: "API详情设置", value: "apiDesc" },
          { label: "轮播图设置", value: "carouselSetting" },
          { label: "Logo设置", value: "logoSetting" },
          { label: "素材内容管理", value: "material" },
          { label: "导航菜单", value: "navigation" },
          { label: "样式设置", value: "styleSetting" }
        ]
      }
    ]
  },
  {
    label: "环境管理",
    value: "envManage",
    menubar: [],
    submenu: [
      {
        label: "环境管理",
        value: "envManage",
        permissions: [
          { label: "显示环境管理菜单", value: "display" },
          { label: "环境列表", value: "envList" }
        ]
      }
    ]
  },
  {
    label: "集成管理",
    value: "integrationManage",
    menubar: [],
    submenu: [
      {
        label: "集成管理",
        value: "integrationManage",
        permissions: [
          { label: "显示集成管理菜单", value: "display" },
          { label: "D2A管理", value: "d2a" },
          { label: "编排管理", value: "integration" }
        ]
      }
    ]
  },
  {
    label: "监控中心",
    value: "monitoringCenter",
    menubar: [],
    submenu: [
      {
        label: "监控中心",
        value: "monitoringCenter",
        permissions: [
          { label: "显示监控中心菜单", value: "display" },
          { label: "租户总览", value: "overview" },
          { label: "调用统计", value: "call" },
          { label: "资产关系图", value: "relationShip" },
          { label: "自定义仪表盘", value: "customizeDashboard" }
        ]
      }
    ]
  },
  {
    label: "门户管理",
    value: "portal",
    menubar: [],
    submenu: [
      {
        label: "门户管理",
        value: "portalManage",
        permissions: [
          { label: "显示门户管理菜单", value: "display" },
          { label: "开放门户管理", value: "openPortal" },
          { label: "预授权门户管理", value: "preGrantPortal" },
          { label: "应用中转站", value: "transferStation" }
        ]
      }
    ]
  },
  {
    label: "测试中心",
    value: "portal",
    menubar: [],
    submenu: [
      {
        label: "测试中心",
        value: "test",
        permissions: [
          { label: "显示测试中心菜单", value: "display" },
          { label: "测试套件", value: "testSuite" },
          { label: "测试计划", value: "testPlan" }
        ]
      }
    ]
  },
  {
    label: "门户管理",
    value: "portal",
    menubar: [],
    submenu: [
      {
        label: "合作伙伴管理",
        value: "partner",
        permissions: [
          { label: "查看合作伙伴及菜单", value: "read" },
          { label: "创建合作伙伴", value: "create" },
          { label: "编辑合作伙伴", value: "edit" },
          { label: "删除伙伴", value: "delete" },
          { label: "修改密码", value: "password" }
        ]
      }
    ]
  },
  {
    label: "系统设置",
    value: "system",
    menubar: [
      // { label: '显示API安全管理菜单', value: 'display' },
    ],
    submenu: [
      {
        label: "账号管理",
        value: "userManage",
        permissions: [
          { label: "查看账号及菜单", value: "read" },
          { label: "创建账号", value: "create" },
          { label: "编辑账号", value: "edit" },
          { label: "删除账号", value: "delete" },
          { label: "修改密码", value: "password" },
          { label: "资源分配", value: "sourceManage" }
        ]
      },
      {
        label: "角色管理",
        value: "roleManage",
        permissions: [
          { label: "查看角色及菜单", value: "read" },
          { label: "创建角色", value: "create" },
          { label: "编辑角色", value: "edit" },
          { label: "删除角色", value: "delete" },
          { label: "复制角色", value: "copy" }
        ]
      },
      {
        label: "租户管理",
        value: "orgManage",
        permissions: [
          { label: "查看租户及菜单", value: "read" },
          { label: "创建租户", value: "create" },
          { label: "编辑租户", value: "edit" }
        ]
      },
      {
        label: "License管理",
        value: "licenseManage",
        permissions: [{ label: "显示License管理菜单", value: "read" }]
      }
    ]
  },
  {
    label: "运维管理",
    value: "devops",
    menubar: [
      // { label: '显示运维管理菜单', value: 'display' },
    ],
    submenu: [
      {
        label: "运维管理",
        value: "operationManagement",
        permissions: [
          { label: "显示运维管理菜单", value: "read" },
          { label: "运维监控", value: "monitor" },
          { label: "运维告警", value: "alert" }
        ]
      }
    ]
  }
]
