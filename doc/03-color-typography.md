# 配色与字体系统

## 主色调

网站采用**中国传统玄学/佛教风格**的配色方案，以深色为底、金色为主色调。

### 核心色板（从 CSS 变量推断）

| 色名 | 用途 | 视觉效果 |
|------|------|----------|
| `xuan` | 主背景色 | 深棕/墨色（约 `#1a1410`） |
| `xuan-card` | 卡片背景 | 比 xuan 稍浅的棕色 |
| `xuan-surface` | 表面/容器 | 浅棕色 |
| `gold` | 主强调色 | 金色（约 `#c9a05c`） |
| `gold-dark` | 深金色 | 用于文字阴影 |
| `vermillion` | 次强调色 | 朱红色（中国传统红） |
| `vermillion-light` | 亮朱红 | 用于"忌"、凶神等 |
| `paper` | 正文文字色 | 浅色文字 |
| `paper-dark` | 次要文字 | 用于副标题、说明文字 |
| `paper-warm` | 暖白 | 用于 scroll 卷轴背景 |
| `ink-muted` | 弱化文字 | 用于底部导航未选中项 |

### 功能色

| 场景 | 配色 |
|------|------|
| 宜（好事） | 翡翠绿系：`emerald-500/30`、`emerald-900/10`、`emerald-300` |
| 忌（避讳） | 朱红系：`vermillion/30`、`vermillion/10`、`vermillion-light` |
| 五行-金 | 黄色系：`yellow-300`、`yellow-800` |
| 五行-木 | 绿色系：`green-300`、`green-800` |
| 五行-水 | 蓝色系：`blue-300`、`blue-800` |
| 五行-火 | 红色系：`red-300`、`red-800` |
| 五行-土 | 琥珀系：`amber-300`、`amber-800` |
| 吉 | 朱红系 |
| 凶 | 玄色系 |
| 平 | 灰色系：`ink-muted` |

## 渐变效果

```
Logo 文字渐变：linear-gradient(180deg, #f5e6b8 0%, #c9a05c 50%, #8b6914 100%)
  - 从亮金 → 中金 → 暗金，模拟金属质感

分享按钮渐变：from-gold/35 via-gold/20 to-vermillion/20
  - 金到红的对角渐变

进度条渐变：from-gold to-vermillion
  - 金到红的水平渐变
```

## 字体

### ZhiMangXing（志莽行体）

- **用途**：Logo、页面标题、Display 级文字
- **CSS 名称**：`font-display`
- **加载方式**：`@font-face` 声明，WOFF2 格式
- **字体子集化**：`ZhiMangXing-subset.woff2`
- **font-display: block**：阻塞渲染，确保字体加载后显示

### 正文字体

- **CSS 名称**：`font-body`
- 回退字体栈：`system-ui, "Segoe UI", Roboto, Helvetica, Arial, sans-serif`

### 等宽字体

- 用于 URL/代码展示：`font-mono`

## 文字大小规范

| 元素 | 类名 | 大小 |
|------|------|------|
| Logo | `text-[1.4rem] md:text-[1.65rem]` | 22.4px / 26.4px |
| 页面大标题 | `text-4xl` | 36px |
| 模块标题 | `text-2xl` / `text-xl` | 24px / 20px |
| 正文 | `text-base` | 16px |
| 副标题/说明 | `text-sm` | 14px |
| 小字/标签 | `text-xs` / `text-[11px]` / `text-[10px]` | 12px / 11px / 10px |

## 特殊视觉效果

### Logo 文字

```css
font-family: 'ZhiMangXing', cursive;
background: linear-gradient(180deg, #f5e6b8 0%, #c9a05c 50%, #8b6914 100%);
-webkit-background-clip: text;
-webkit-text-fill-color: transparent;
background-clip: text;
letter-spacing: 0.12em;
filter: drop-shadow(0 1px 2px rgba(0,0,0,0.3));
```

### Logo SVG

- 菩提叶形状，SVG 内联
- `fill="currentColor" fill-opacity="0.12"` 半透明填充
- 多条对称曲线模拟叶脉
- `drop-shadow-[0_0_8px_rgba(201,160,94,0.4)]` 金色辉光

### 金色分割线

- `.gold-divider`：用于顶部导航底部和卷轴上下装饰

### 圆角卡片

- 默认卡片：`rounded-lg`（8px）
- 卷轴变体：`rounded-t-xl rounded-b-xl`（12px）
- 浮动按钮：`rounded-full`
- 标签：`rounded-full`
