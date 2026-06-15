# 动画效果详解

## CSS 动画

### glow-rise（光点上升）

```css
@keyframes glow-rise {
  /* 从底部向上漂浮，同时透明度变化 */
}

.animate-glow-rise {
  animation: glow-rise 5s infinite;
}
```

- 用于 GlobalBackground 中的 6 个金色光点
- 每个光点有不同的 `animation-delay`（0s ~ 3.1s）
- 大小：`size-1.5`（6px）
- 颜色：`bg-gold/40`
- 无限循环

### spin-slow（慢速旋转）

```css
@keyframes spin-slow {
  /* 360° 旋转 */
}

.animation: spin-slow 18s linear infinite
```

- 用于 MiniPlayer 中的曲目图标
- 播放时旋转，暂停时停止
- 18 秒一圈

### ping（脉冲环）

```css
.animate-ping {
  /* 同时放大 + 淡出 */
}
```

- 用于音乐播放按钮的脉冲效果
- 播放中时显示金色脉冲圆环

### pulse（脉冲）

```css
.animate-pulse {
  /* 透明度脉冲 */
}
```

- 用于加载状态、播放中指示

### bounce（弹跳）

```css
.animate-bounce {
  /* 上下弹跳 */
}
```

- 用于安装引导的 👇 箭头

### fadeInDown

```css
@keyframes fadeInDown {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}
```

- 用于安装引导浮窗的入场动画

## Framer Motion 动画

### ScrollReveal

```javascript
{
  initial: { opacity: 0, y: 24 },
  whileInView: { opacity: 1, y: 0 },
  viewport: { once: true, margin: "-40px" },
  transition: {
    duration: 0.5,
    delay: 0,           // 可配置
    ease: [0.25, 0.46, 0.45, 0.94]
  }
}
```

- 页面滚动时触发
- 从下方/上方/右侧淡入
- 使用 `viewport.once: true` 确保只触发一次
- 提前 40px 触发（`margin: "-40px"`）
- 多个 ScrollReveal 使用递增 delay（0, 0.05, 0.1, 0.15, 0.2）

### 弹窗动画

```javascript
// 分享弹窗
<div className="fixed inset-0 z-[120] flex items-end justify-center
                bg-xuan/95 backdrop-blur-sm sm:items-center">
```

- 从底部滑入（移动端）/ 居中显示（桌面端）
- 背景毛玻璃效果

## 过渡动画

### 导航栏滚动过渡

```css
transition-all duration-base
```

- 从透明到毛玻璃的背景过渡
- 分割线 opacity 过渡

### 页面切换

```css
transition-colors duration-fast
```

- 导航链接 hover 颜色过渡
- 按钮状态过渡

### 卡片悬停

```css
transition-all duration-base
hover:border-gold/30 hover:shadow-card
```

- 边框颜色过渡
- 阴影过渡

## 动画时长规范

| 类名 | 时长 | 用途 |
|------|------|------|
| `duration-fast` | ~150ms | 颜色切换、悬停 |
| `duration-base` | ~300ms | 背景过渡、导航 |
| `duration-slow` | ~500ms | 分割线、大范围过渡 |

## 性能优化

- 光点动画使用 `will-change: transform`（通过 animation 属性隐含）
- ScrollReveal 使用 `viewport.once: true` 避免重复触发
- 背景图使用 `opacity-[0.20]` 降低渲染负担
- `pointer-events-none` 用于不交互的装饰层
