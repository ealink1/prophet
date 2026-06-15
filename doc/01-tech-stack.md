# 技术栈分析

## 前端框架

- **Next.js**（App Router 模式）
  - 路由结构：`/app/almanac/page.tsx`、`/app/layout.tsx`
  - 使用 Server Components + Client Components 混合
  - RSC Payload 传输（`self.__next_f` 堆栈）
  - 构建ID：`Bs5FTOAgZeLp0dXEXHsT5`

## UI 层

- **Tailwind CSS v3.4.19**（PostCSS 编译）
  - 大量使用自定义主题色（xuan、gold、vermillion、paper-dark 等）
  - 支持 dark 模式变体
  - 响应式断点：md（768px）、lg（1024px）
- **Framer Motion**（`framer-motion`）
  - `motion.div` 用于 `ScrollReveal` 动画
  - `whileInView` 视口触发
  - 缓动曲线：`[.25, .46, .45, .94]`
- **CVA（class-variance-authority）**
  - 用于组件变体管理（Card、Badge 等）
- **cn() 工具函数**
  - `clsx` + `tailwind-merge` 合并类名

## 图标

- **Lucide React**（SVG 图标库）
  - 全站使用 Lucide 图标体系
  - 图标通过 `lucide()` 工具函数创建

## 字体

- **ZhiMangXing**（志莽行体）- 中文书法字体
  - 用于 Logo、标题、Display 文字
  - WOFF2 格式，子集化加载（`/fonts/ZhiMangXing-subset.woff2`）
  - `font-display: block`（阻塞渲染等待字体加载）

## 数据层

- **Cloudflare 后端 API**
  - Base URL：`https://denied-accomplished-micro-drawn.trycloudflare.com/api/v1`
  - 所有请求使用 POST 方法
  - 统一响应格式：`{ code: 0, data: {...}, message: "..." }`
- **localStorage 持久化**
  - `lingji_token_v2`：JWT Token
  - `lingji_user_v2`：用户信息 JSON
  - `lingji_device_id`：设备标识
  - `lingji_invite_code`：邀请码

## 构建与部署

- **Cloudflare Pages**
  - `_next/static/chunks/` 静态资源
  - manifest.json 声明 PWA
  - Service Worker 支持
