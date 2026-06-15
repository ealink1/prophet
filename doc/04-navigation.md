# 导航系统

## 导航层级

```
一级导航（桌面端 9 个链接）
├── 为家人祈福
├── 今日黄历
├── 求灵签
├── 八字精批
├── 周公解梦
├── 看手相
├── 宝宝起名
├── 六爻占卜
└── 静心禅坐

二级导航（移动端底部 6 个 Tab）
├── 首页（🏠）
├── 祈福（❤️）
├── 黄历（📅）
├── 灵签（📜）
├── 我的（👤）
└── 更多（⊞）
```

## 导航数据配置

```javascript
// 桌面端导航
cx = [
  { href: "/qifu",    label: "为家人祈福" },
  { href: "/almanac", label: "今日黄历" },
  { href: "/lottery", label: "求灵签" },
  { href: "/bazi",    label: "八字精批" },
  { href: "/dream",   label: "周公解梦" },
  { href: "/palmistry", label: "看手相" },
  { href: "/naming",  label: "宝宝起名" },
  { href: "/divination", label: "六爻占卜" },
  { href: "/meditation", label: "静心禅坐" },
]

// 移动端底部导航
P$ = [
  { href: "/",      label: "首页", icon: "home" },
  { href: "/qifu",  label: "祈福", icon: "flame" },
  { href: "/almanac", label: "黄历", icon: "calendar" },
  { href: "/lottery", label: "灵签", icon: "scroll" },
  { href: "/profile", label: "我的", icon: "user" },
  { href: "/more",  label: "更多", icon: "more" },
]
```

## 顶部导航行为

1. **初始状态**：透明背景
2. **滚动触发**：`window.scrollY > 20` 时切换为毛玻璃背景
3. **分割线**：滚动时 `gold-divider` opacity 从 0 过渡到 100
4. **路由变化**：`usePathname()` 监听当前路径，高亮对应链接
5. **隐藏规则**：
   - `/dy` 路径隐藏
   - `/admin/*` 路径隐藏
   - 弹窗打开时隐藏

## 底部导航行为

1. **响应式**：仅移动端显示（`md:hidden`）
2. **当前页高亮**：`text-gold` vs `text-ink-muted`
3. **安全区域**：适配 iPhone 底部安全区
4. **弹窗联动**：
   - 监听 `lingji:modal-open` / `lingji:modal-close` 自定义事件
   - 弹窗打开时隐藏，关闭后显示
5. **路由切换后重置**：`useEffect` 监听 `pathname` 变化重置状态

## 图标映射

```javascript
iconMap = {
  home:     Lucide House 图标
  flame:    Lucide Heart 图标（祈福用爱心）
  calendar: Lucide CalendarDays 图标
  scroll:   Lucide ScrollText 图标
  compass:  Lucide Compass 图标
  user:     Lucide User 图标
  more:     Lucide LayoutGrid 图标
}
```

## 二级页面导航

- 八字精批、求灵签页面有**师父选择器**（角色扮演式 AI 导师）
- 每个师父有头像 emoji、名称、角色描述、风格描述
- 三个角色：
  - 🧘 慧明长老 - 古寺住持，庄重持重
  - 🙏 明心师父 - 尼众法师，慈悲温柔
  - ☯️ 玄真道长 - 山中道人，直爽通透

## 寻找记录/用户入口

- 桌面端：顶部导航右侧 "找回记录" 按钮（未登录态）/ 用户头像+吉祥号（已登录态）
- 使用 `useAuth` hook 获取用户状态
- `openRestore()` 打开找回记录弹窗
