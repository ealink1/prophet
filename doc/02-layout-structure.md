# 页面布局结构

## 全局层级（z-index）

```
z-0   背景层（GlobalBackground）
z-10  内容区（PageTransition/main）
z-40  底部导航栏、浮动按钮、迷你播放器
z-50  顶部导航栏
z-110 抖音/TikTok 内嵌浏览器提示
z-120 分享弹窗
z-199 弹窗遮罩
z-200 安装引导浮窗
```

## 全局组件树

```
<html lang="zh-CN">
  <body className="antialiased">
    ├── GlobalBackground（z-0，固定背景）
    ├── AuthProvider（认证上下文）
    │   └── PlayerProvider（音频播放上下文）
    │       ├── ActivityTracker（页面访问追踪，无渲染）
    │       ├── InAppBrowserHint（抖音/TikTok 内嵌提示）
    │       ├── InstallGuide（微信/抖音内安装引导）
    │       ├── TopNav（顶部导航栏，z-50）
    │       ├── PageTransition（main 容器，z-10）
    │       │   └── {children}（页面内容）
    │       ├── MiniPlayer（迷你音乐播放器）
    │       ├── ShareFAB（浮动分享按钮）
    │       └── BottomBar（底部导航栏，移动端 z-40）
```

## GlobalBackground 背景层

由 5 层叠加构成：

1. **渐变背景**：`bg-gradient-to-b from-xuan via-xuan-card to-xuan`
   - 从深色到卡片色再到深色的垂直渐变
2. **山水纹理**：`temple/temple-mountain.svg`
   - 背景图，opacity 20%，覆盖模式
3. **径向暗角**：
   - 中心透明、四周暗角的椭圆径向渐变
   - `rgba(10,6,4,0.55)` → `transparent` → `rgba(10,6,4,0.6)`
4. **顶部金色光晕**：`from-gold/15 to-transparent`
   - 高 32px，从顶部向下的金色渐变
5. **浮动光点粒子**：
   - 6 个金色圆点，`bg-gold/40`
   - 使用 `animate-glow-rise` 动画上升
   - 不同位置和延迟（0s ~ 3.1s），循环播放

## TopNav 顶部导航

- 固定在顶部，高度 `h-14`（56px）
- **透明模式**（默认）→ **毛玻璃模式**（滚动 > 20px 后触发）
  - 背景：`bg-xuan/95 backdrop-blur-md`
  - 阴影：`shadow-[0_1px_0_rgba(201,169,110,0.12)]`
  - 底部金色分割线：`gold-divider`，滚动时 opacity 0→100

### Desktop 布局（md+）

```
[Logo + "菩提苑"] ──── [导航链接 × 9] ──── [音乐按钮] [安装按钮] [找回记录/用户头像]
```

### Mobile 布局

```
[Logo + "菩提苑"] ──────────────────────────── [音乐按钮] [安装按钮]
```

## BottomBar 底部导航

- 仅在移动端显示（`md:hidden`）
- 固定在底部，`border-t border-gold/20 bg-xuan-card/97 backdrop-blur-md`
- 6 列网格：首页 / 祈福 / 黄历 / 灵签 / 我的 / 更多
- 适配 iPhone 安全区域：`pb-[max(env(safe-area-inset-bottom),0.5rem)]`
- 当前页高亮：`text-gold`，其他：`text-ink-muted`
- 弹窗打开时自动隐藏

## PageTransition 内容区

- `min-h-[calc(100vh-3.5rem)]`（减去顶部导航高度）
- `pt-14`（顶部留白）`pb-24 md:pb-8`（底部留白，移动端考虑底部导航）

## ShareFAB 浮动按钮

- 固定右下角
- 移动端：`right-3 bottom-[calc(env(safe-area-inset-bottom)+88px)]`（避开底部导航）
- 桌面端：`right-4 bottom-4`
- 圆形渐变背景：`from-gold/35 via-gold/20 to-vermillion/20`
- 点击打开分享返佣弹窗

## MiniPlayer 迷你播放器

- 固定左下角，仅在非 meditation 页面显示
- 显示当前播放曲目、播放/暂停按钮、关闭按钮
- 进度条：`from-gold to-vermillion` 渐变
- 弹窗打开时自动隐藏
