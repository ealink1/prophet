# 黄历页面详解

## 路由

`/almanac/` → `app/almanac/page.tsx`

## 数据获取

### 今日黄历

```javascript
POST /api/v1/almanac/today
Body: { date: "YYYY-MM-DD" }
Response: {
  code: 0,
  data: {
    solar: { year, month, day, weekday_full },
    lunar: { year_in_chinese, month_in_chinese, day_in_chinese, year_zodiac },
    ganzhi: { year, month, day, nayin },
    overall_level: { level: "上上"|"上吉"|"中吉"|"中平"|"下下", summary: "..." },
    jieqi: { today: "..." | null },
    yi: ["宜事1", "宜事2", ...],
    ji: ["忌事1", "忌事2", ...],
    shen: { lucky: [...], unlucky: [...] },
    chong: "冲煞信息",
    tai_position: "胎神方位",
    xiu: "28宿名称",
    xiu_luck: "28宿吉凶",
    zhixing: "12建除",
    shichen: [
      { name: "子", lucky: "吉", ganzhi: "天干地支", chong: "冲..." },
      ...
    ]
  }
}
```

### 未来七日

```javascript
POST /api/v1/almanac/week
Body: {}
Response: {
  code: 0,
  data: {
    items: [
      { date: "YYYY-MM-DD", weekday: "一", lunar_day: "初一", level: "中吉" },
      ...
    ]
  }
}
```

## 页面状态管理

```javascript
const [data, setData] = useState(null);          // 当日黄历数据
const [weekData, setWeekData] = useState([]);     // 七日数据
const [error, setError] = useState(null);         // 错误信息
const [currentDate, setCurrentDate] = useState(g(new Date)); // 当前查看日期
```

- `currentDate` 变化时重新 fetch 今日黄历
- 七日数据仅在组件挂载时 fetch 一次
- `g()` 函数：`Date → "YYYY-MM-DD"` 格式化

## 页面布局

```
┌─────────────────────────────────────────┐
│  [<]   今日黄历   [>]   ← 日期切换器     │
│       2026年06月15日                      │
│       星期一 · 农历丙午年五月廿一          │
│       干支：丙午年 甲午月 ...              │
├─────────────────────────────────────────┤
│  ☀ 今日 · 上上  万事皆宜，大吉大利       │ ← 综合评级
│  今日节气：夏至                           │ ← 节气（条件显示）
├─────────────────────────────────────────┤
│  ┌──────────────┐ ┌──────────────┐      │
│  │  宜 今日适合   │ │  忌 今日避开   │      │ ← 宜忌双栏
│  │  [标签] [标签] │ │  [标签] [标签] │      │
│  └──────────────┘ └──────────────┘      │
├─────────────────────────────────────────┤
│  神煞 · 冲煞                              │
│  ┌────────────┐ ┌────────────┐          │
│  │ 吉神宜趋    │ │ 凶神宜避    │          │ ← 神煞信息
│  │ 文昌 天德... │ │ 天火 ...    │          │
│  └────────────┘ └────────────┘          │
│  ┌────────────┐ ┌────────────┐          │
│  │ 冲煞        │ │ 胎神方位    │          │
│  └────────────┘ └────────────┘          │
│  ┌────────────┐ ┌────────────┐          │
│  │ 28 宿       │ │ 12 建除     │          │
│  └────────────┘ └────────────┘          │
├─────────────────────────────────────────┤
│  时辰吉凶                                │
│  ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐   │
│  │子时 吉│ │丑时 凶│ │寅时 平│ │卯时 吉│   │ ← 时辰网格
│  │甲子   │ │乙丑   │ │丙寅   │ │丁卯   │   │
│  │冲：...│ │       │ │冲：...│ │       │   │
│  └──────┘ └──────┘ └──────┘ └──────┘   │
├─────────────────────────────────────────┤
│  未来七日                                │
│  [一] [二] [三] [四] [五] [六] [日]     │ ← 七日横向
│  初一 初二 初三 初四 初五 初六 初七        │
│  上吉 中吉 中平 下下 上吉 上吉 中吉      │
└─────────────────────────────────────────┘
```

## 组件结构

```
<ScrollReveal>
  <Card variant="default">
    ← 日期切换 + 基本信息
    ← 综合评级 Badge
    ← 节气（条件渲染）
  </Card>
</ScrollReveal>

<ScrollReveal delay={0.05}>
  <div className="grid md:grid-cols-2">
    <Card variant="default">← 宜</Card>
    <Card variant="default">← 忌</Card>
  </div>
</ScrollReveal>

<ScrollReveal delay={0.1}>
  <Card variant="default">← 神煞 · 冲煞</Card>
</ScrollReveal>

<ScrollReveal delay={0.15}>
  <Card variant="default">← 时辰吉凶</Card>
</ScrollReveal>

<ScrollReveal delay={0.2}>
  <Card variant="default">← 未来七日</Card>
</ScrollReveal>
```

## 交互细节

### 日期切换

- 左右箭头按钮：`k(-1)` / `k(1)` 加减天数
- 日期格式化为 `"YYYY-MM-DD"` 发送给 API
- 当前日期（今天）显示 "今日黄历"，其他日期仅显示 "黄历"
- `_` 变量：`useMemo` 判断是否为今日

### 灵签评级颜色映射

```javascript
levelColors = {
  "上上": "text-vermillion-light",  // 亮朱红
  "上吉": "text-vermillion-light",
  "中吉": "text-gold",              // 金色
  "中平": "text-paper",             // 白色
  "下下": "text-paper-dark/70",     // 暗色
}

levelBorders = {
  "上上": "border-vermillion/40 bg-vermillion/10",
  "上吉": "border-vermillion/30 bg-vermillion/5",
  "中吉": "border-gold/40 bg-gold/10",
  "中平": "border-gold/20 bg-xuan-surface/50",
  "下下": "border-paper-dark/20 bg-paper-dark/5",
}
```

### 时辰吉凶网格

- 移动端：2 列
- md：3 列
- lg：4 列
- 每个时辰显示：名称、吉凶标签（绿色"吉" / 红色"凶"）、干支、冲信息

### 七日横向滚动

- `grid-cols-7` 七列等宽
- 当前选中日期：`border-gold bg-gold/15` 高亮
- 可点击切换查看历史/未来日期
