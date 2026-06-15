# 全局组件详解

## Card 组件

三种变体：

### 1. default（默认卡片）

```css
rounded-lg border border-gold/20 bg-xuan-card/95 p-card-pad
shadow-paper backdrop-blur-sm
hover:border-gold/30 hover:shadow-card
```

- 圆角 8px
- 金色边框 20% 透明度
- 卡片背景 95% 透明度
- 悬停时边框加深 + 阴影增强

### 2. scroll（卷轴）

```css
rounded-t-xl rounded-b-xl border-x-4 border-gold/40
bg-paper-warm px-8 py-6 text-ink shadow-card
```

- 上下圆角 12px
- 左右 4px 金色边框
- 暖白背景（模拟纸张）
- 上下有 `gold-divider` 装饰线

### 3. stamp（印章）

```css
rounded-md border-2 border-vermillion/30 bg-vermillion/5
p-4 shadow-stamp
```

- 2px 朱红边框
- 朱红背景 5% 透明度
- 印章风格阴影

## Badge 组件

多变体标签组件：

### variant（变体）

| 变体 | 样式 |
|------|------|
| `element` | 金色边框 + 金色背景 + 金色文字 |
| `shensha` | 朱红边框 + 朱红背景 + 朱红文字 |
| `gongde` | 金色渐变边框 + 渐变背景 + 深金色文字 |

### tone（五行/吉凶色调）

| 色调 | 边框 | 背景 | 文字 |
|------|------|------|------|
| 金 | `yellow-300` | `yellow-100` | `yellow-800` |
| 木 | `green-300` | `green-100` | `green-800` |
| 水 | `blue-300` | `blue-100` | `blue-800` |
| 火 | `red-300` | `red-100` | `red-800` |
| 土 | `amber-300` | `amber-100` | `amber-800` |
| 吉 | `vermillion/30` | `vermillion/10` | `vermillion` |
| 凶 | `xuan/30` | `xuan/10` | `paper-dark` |
| 平 | `ink-muted/30` | `white/5` | `ink-muted` |

## ScrollReveal 组件

滚动触发动画组件：

```javascript
ScrollReveal({ children, delay=0, direction="bottom", className })
```

### 参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `children` | - | 内容 |
| `delay` | 0 | 延迟秒数 |
| `direction` | "bottom" | 动画方向：bottom/top/right |
| `className` | - | 自定义类名 |

### 动画配置

```javascript
initial: { opacity: 0, y: 24 }  // bottom
         { opacity: 0, y: -24 } // top
         { opacity: 0, x: 24 }  // right

whileInView: { opacity: 1, x: 0, y: 0 }

viewport: { once: true, margin: "-40px" }

transition: {
  duration: 0.5,
  delay: delay,
  ease: [0.25, 0.46, 0.45, 0.94]  // 自定义缓动
}
```

- `once: true`：仅触发一次
- `margin: "-40px"`：提前 40px 触发

## Button 组件

使用 CVA 管理变体：

```javascript
Button({ variant, size, children, className, ...props })
```

### 已知变体

- `ritual`：仪式感按钮（祈福/分享场景）
- `secondary`：次要按钮

## cn() 工具函数

```javascript
cn(...classes) // clsx + tailwind-merge
```

合并类名，处理冲突（如同时存在 `text-gold` 和 `text-red` 时后者覆盖前者）。

## formatCurrency() 工具函数

```javascript
formatCurrency(9.9)  // "¥9.9"
formatCurrency(100)  // "¥100"
```

使用 `Intl.NumberFormat("zh-CN", { style: "currency", currency: "CNY" })`。

## parseSections() 工具函数

将长文本按段落和标题分割为结构化数组：

```javascript
parseSections(text)
// 按 \n\n 分割，再按数字编号或 ** 粗体标题细分
```

用于解析 AI 生成的长文本内容。
