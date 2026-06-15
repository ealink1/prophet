# 完整页面路由表

## 主要功能页面

| 路径 | 页面 | 说明 |
|------|------|------|
| `/` | 首页 | 网站主页（推测为功能入口聚合） |
| `/qifu/` | 为家人祈福 | 点功德灯、供奉、灯墙 |
| `/almanac/` | 今日黄历 | 每日宜忌、时辰、七日历 |
| `/lottery/` | 求灵签 | 关帝灵签、师父开示 |
| `/bazi/` | 八字精批 | 生辰排盘、命理分析 |
| `/dream/` | 周公解梦 | 梦境描述解梦、分类查梦 |
| `/palmistry/` | 看手相 | 手相图解读 |
| `/naming/` | 宝宝起名 | 五行/音韵/笔画起名 |
| `/divination/` | 六爻占卜 | 六爻卦象分析 |
| `/meditation/` | 静心禅坐 | 禅修音乐播放 |

## 用户相关页面

| 路径 | 页面 | 说明 |
|------|------|------|
| `/profile/` | 我的 | 用户中心、历史记录 |
| `/more/` | 更多 | 功能入口聚合页 |

## 管理页面（隐藏）

| 路径 | 页面 | 说明 |
|------|------|------|
| `/admin/*` | 管理后台 | 隐藏导航和背景 |
| `/dy` | 抖音专用页 | 隐藏导航和背景 |

## 静态资源

| 路径 | 说明 |
|------|------|
| `/manifest.json` | PWA 配置 |
| `/favicon.svg` | SVG 图标 |
| `/favicon.ico` | ICO 图标（16x16） |
| `/share-cover.svg` | 分享封面图（1200x630） |
| `/fonts/ZhiMangXing-subset.woff2` | 志莽行体字体 |
| `/temple/temple-mountain.svg` | 背景山水图 |
| `/icons/icon-192.png` | PWA 图标 192px |
| `/icons/icon-512.png` | PWA 图标 512px |

## 路由守卫规则

```javascript
// 导航隐藏规则
if (pathname === "/dy" || pathname.startsWith("/admin")) {
  // 隐藏 TopNav
  // 隐藏 BottomBar
  // 隐藏 GlobalBackground
}

// 弹窗时隐藏规则
if (modalOpen) {
  // 隐藏 BottomBar
  // 隐藏 MiniPlayer
}

// 页面追踪排除
if (pathname.startsWith("/admin")) {
  // 不追踪
}
```

## 响应式断点

| 断点 | 宽度 | 行为 |
|------|------|------|
| 默认 | < 768px | 移动端布局，底部导航 |
| `md` | ≥ 768px | 桌面端布局，顶部导航 |
| `lg` | ≥ 1024px | 时辰网格 4 列 |

## 安全区域适配

```css
/* iPhone 底部安全区 */
pb-[max(env(safe-area-inset-bottom),0.5rem)]

/* iPhone 顶部安全区 */
safe-top

/* 底部导航避开安全区 */
bottom-[calc(env(safe-area-inset-bottom)+88px)]
```
