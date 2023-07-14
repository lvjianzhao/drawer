/** 布局配置 */
interface LayoutSettings {
  /** 是否显示 Settings Panel */
  showSettings: boolean
  /** 是否显示标签栏 */
  showTagsView: boolean
  /** 是否显示侧边栏 Logo */
  showSidebarLogo: boolean
  /** 是否固定 Header */
  fixedHeader: boolean
  /** 是否显示全屏按钮 */
  /** 是否显示切换主题按钮 */
  showThemeSwitch: boolean
  showScreenfull: boolean
}

const layoutSettings: LayoutSettings = {
  showSettings: false,
  showTagsView: false,
  fixedHeader: true,
  showSidebarLogo: true,
  showThemeSwitch: false,
  showScreenfull: false
}

export default layoutSettings
