# PWA 与 SEO

## PWA 配置

### manifest.json

```json
{
  "name": "菩提苑",
  "short_name": "菩提苑",
  "description": "心诚则灵。点一盏灯，求一支签，看一卦命。一念慈悲，福报自来。",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#1a1410",
  "theme_color": "#1a1410",
  "orientation": "portrait",
  "icons": [
    { "src": "/icons/icon-192.png", "sizes": "192x192", "type": "image/png", "purpose": "any maskable" },
    { "src": "/icons/icon-512.png", "sizes": "512x512", "type": "image/png", "purpose": "any maskable" },
    { "src": "/favicon.svg", "sizes": "any", "type": "image/svg+xml" }
  ]
}
```

### PWA 特性

- `display: standalone`：全屏 App 体验
- `orientation: portrait`：锁定竖屏
- `maskable` 图标：支持 Android 自适应图标
- `theme_color: #1a1410`：深棕色主题

### Meta 标签（PWA 相关）

```html
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
<meta name="apple-mobile-web-app-title" content="菩提苑">
<meta name="x5-fullscreen" content="true">
<meta name="x5-page-mode" content="app">
```

- iOS Safari 全屏支持
- Android X5 内核（QQ/微信）全屏支持
- 状态栏半透明

## SEO 配置

### 基础 Meta

```html
<title>菩提苑 · 为家人祈福求灵签</title>
<meta name="description" content="心诚则灵。为家人点一盏祈福灯，求一支关帝灵签，看一卦命理八字。一念慈悲，福报自来。">
<meta name="keywords" content="菩提苑,祈福,求签,关帝灵签,八字精批,周公解梦,求灵签,看手相,命理">
<meta name="author" content="菩提苑">
```

### Open Graph

```html
<meta property="og:title" content="菩提苑 · 为家人祈福求灵签">
<meta property="og:description" content="心诚则灵。为家人点一盏祈福灯，求一支关帝灵签，看一卦命理八字。一念慈悲，福报自来。">
<meta property="og:site_name" content="菩提苑">
<meta property="og:image" content="https://putiyuan.pages.dev/share-cover.svg">
<meta property="og:image:width" content="1200">
<meta property="og:image:height" content="630">
<meta property="og:image:alt" content="菩提苑 · 为家人祈福">
<meta property="og:type" content="website">
```

### Twitter Card

```html
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:title" content="菩提苑 · 为家人祈福求灵签">
<meta name="twitter:description" content="心诚则灵。为家人点一盏祈福灯，求一支关帝灵签，看一卦命理八字。一念慈悲，福报自来。">
<meta name="twitter:image" content="https://putiyuan.pages.dev/share-cover.svg">
```

### 其他

```html
<meta name="format-detection" content="telephone=no, address=no, email=no">
<meta name="theme-color" content="#1A1410">
```

- 禁止自动识别电话/地址/邮箱
- 主题色与 PWA 一致

## 内嵌浏览器适配

### 微信/抖音检测

```javascript
// InstallGuide 组件
检测 userAgent 中的 micromessenger / aweme / bytedancewebview
显示安装引导浮窗（10秒后自动关闭）
提示用户"用浏览器打开"
```

### 抖音/TikTok 专用提示

```javascript
// InAppBrowserHint 组件
检测 aweme / bytedance / douyin / tiktok / musical_ly
显示顶部横幅提示：
- "请在浏览器中打开 · 抖音 内打不开支付与登录"
- "点右上角「⋮」选「在浏览器中打开」"
- 提供"复制链接"按钮
- 可关闭，关闭后 session 内不再显示
```

### iOS 添加到桌面引导

```javascript
// TopNav 中的安装按钮
检测 standalone 模式
提供 iOS 和 Android 分别的添加到桌面步骤
- iOS：分享 → 添加到主屏幕 → 添加
- Android：菜单 → 添加到主屏幕 → 确认
```

## 活动追踪

```javascript
// ActivityTracker 组件
POST /api/v1/activity/track
Body: {
  device_id: "...",
  event_type: "page_view",
  path: "/almanac",
  title: "黄历",
  referrer: "..."
}
```

- 800ms 延迟发送（不阻塞首屏）
- 使用 `keepalive: true` 确保页面关闭时也能发送
- 使用 AbortController 取消未完成的请求
