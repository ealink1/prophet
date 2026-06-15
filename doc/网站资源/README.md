# 菩提苑网站资源清单

来源：https://putiyuan.pages.dev/

## 资源总览

| 类型 | 文件数 | 总大小 |
|------|--------|--------|
| 🎵 音频 (MP3) | 10 | ~39 MB |
| 🖼️ 图片 (PNG/SVG) | 5 | ~1.5 MB |
| 📝 字体 (WOFF2) | 1 | 1.7 KB |
| 📜 JavaScript | 19 | ~450 KB |
| 🎨 CSS | 1 | 68 KB |
| 📋 配置 | 2 | ~9 KB |
| **合计** | **38** | **~41 MB** |

## 音频文件 (meditation/)

禅修音乐，共 10 首，均为项目原创。

| 文件 | 标题 | 时长 | 分类 | 说明 |
|------|------|------|------|------|
| `bodhi_theme.mp3` | 菩提苑主题曲 | 2:57 | 主题 | 金光普照感，开场冥想 |
| `bodhi_garden.mp3` | 菩提苑 | 2:51 | 禅意 | 苑中清雅，万缘澄定 |
| `bodhi_light.mp3` | 菩提苑·轻音乐 | 3:15 | 轻禅 | 轻柔版主题，长时间陪伴 |
| `bodhi_crossing.mp3` | 菩提苑·渡尘缘 | 3:39 | 禅悟 | 渡过尘缘，返照本心 |
| `palace_dawn.mp3` | 宝殿晨曦 | 2:48 | 晨修 | 晨曦宝殿，适合清晨 |
| `zen_sit.md3` | 禅坐 | 2:36 | 正念 | 结跏趺坐，身心安住 |
| `zen_mind.mp3` | 禅意 | 3:12 | 禅意 | 万像皆禅，处处是道场 |
| `crystal_moon.mp3` | 琉璃月 | 3:31 | 禅韵 | 月光琉璃，照见五蕴皆空 |
| `great_compassion.mp3` | 大悲咒 | 4:06 | 梵音 | 观世音菩萨大悲咒 |
| `heart_sutra.mp3` | 心经 | 3:55 | 梵音 | 般若波罗蜜多心经 |

## 图片资源

| 文件 | 说明 | 尺寸 |
|------|------|------|
| `temple-mountain.svg` | 背景山水图（寺庙+山+月亮） | SVG 800×600 |
| `share-cover.svg` | 分享封面图（OG/Twitter Card） | SVG 1200×630 |
| `favicon.svg` | 网站图标（菩提叶） | SVG 64×64 |
| `icon-192.png` | PWA 图标 192px | PNG 192×192 |
| `icon-512.png` | PWA 图标 512px | PNG 512×512 |
| `palm-guide.png` | 手相图解指南（看手相页面） | PNG ~290KB |
| `books/classics-strip-matted.png` | 古籍装饰图（首页展示） | PNG ~1MB |

## 字体

| 文件 | 说明 |
|------|------|
| `ZhiMangXing-subset.woff2` | 志莽行体（中文书法字体），用于 Logo 和标题 |

## JavaScript 文件

### 核心框架

| 文件 | 说明 |
|------|------|
| `webpack.js` | Webpack runtime |
| `main-app.js` | Next.js App 入口 |
| `chunk-117.js` | React 相关 |
| `chunk-398.js` | 通用依赖 |
| `chunk-529.js` | UI 组件 |
| `chunk-92.js` | 动画库 (Framer Motion) |
| `chunk-420.js` | 工具函数 |
| `chunk-815.js` | 其他依赖 |

### 布局与页面

| 文件 | 说明 |
|------|------|
| `layout.js` | 全局布局（导航/背景/播放器/底部栏） |
| `homepage.js` | 首页组件 |
| `almanac-page.js` | 今日黄历页面 |
| `qifu-page.js` | 为家人祈福页面 |
| `lottery-page.js` | 求灵签页面 |
| `bazi-page.js` | 八字精批页面 |
| `dream-page.js` | 周公解梦页面 |
| `meditation-page.js` | 静心禅坐页面 |
| `palmistry-page.js` | 看手相页面 |
| `naming-page.js` | 宝宝起名页面 |
| `divination-page.js` | 六爻占卜页面 |
| `profile-page.js` | 我的页面 |
| `more-page.js` | 更多页面 |

## 配置文件

| 文件 | 说明 |
|------|------|
| `manifest.json` | PWA 配置 |
| `meditation-catalog.json` | 禅修曲目 API 数据 |
| `main.css` | Tailwind CSS 编译输出 |

## API 数据

| 端点 | 说明 |
|------|------|
| `POST /api/v1/meditation/catalog` | 禅修曲目列表（含音频URL） |
| `POST /api/v1/auth/anonymous/init` | 匿名用户初始化 |
| `POST /api/v1/almanac/today` | 今日黄历数据 |
| `POST /api/v1/almanac/week` | 未来七日数据 |

## 未获取的资源

以下资源通过 API 动态获取，无法静态爬取：

- 灵签数据（`/lottery` API）
- 八字排盘结果（`/bazi` API）
- 解梦结果（`/dream` API）
- 祈福灯墙数据（`/qifu` API）
- 用户上传的支付凭证图片
- 用户手相照片
