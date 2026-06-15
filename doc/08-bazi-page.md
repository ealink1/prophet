# 八字精批功能详解

## 路由

`/bazi/` → 八字排盘页面

## 页面结构

```
┌─────────────────────────────────────────┐
│  # 八字精批                              │
│  输入生辰，洞悉天命，先看命盘，再看流年。    │
├─────────────────────────────────────────┤
│  请选一位师父为您开示                      │ ← 师父选择（同灵签）
├─────────────────────────────────────────┤
│  出生年  [1990年 点击选择]                │ ← 年份选择器
│  出生月  [5月 点击选择]                   │ ← 月份选择器
│  出生日  [15日 点击选择]                  │ ← 日期选择器
│  出生时辰 [未时 (13:00-15:00)]            │ ← 时辰下拉
│  性别    [男] [女]                       │ ← 性别选择
│                                         │
│  [请师父排盘]                            │ ← CTA 按钮
└─────────────────────────────────────────┘
```

## 输入字段

| 字段 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| 出生年 | 选择器 | 1990 | 点击弹出年份列表 |
| 出生月 | 选择器 | 5 | 点击弹出月份列表 |
| 出生日 | 选择器 | 15 | 点击弹出日期列表 |
| 出生时辰 | 下拉选择 | 未时 | 12 时辰可选 |
| 性别 | 按钮组 | 男 | 男/女 |

## 时辰选项

```javascript
shichen = [
  { value: "zi",  label: "子时 (23:00-01:00)", hour: 23 },
  { value: "chou", label: "丑时 (01:00-03:00)", hour: 1 },
  { value: "yin",  label: "寅时 (03:00-05:00)", hour: 3 },
  { value: "mao",  label: "卯时 (05:00-07:00)", hour: 5 },
  { value: "chen", label: "辰时 (07:00-09:00)", hour: 7 },
  { value: "si",   label: "巳时 (09:00-11:00)", hour: 9 },
  { value: "wu",   label: "午时 (11:00-13:00)", hour: 11 },
  { value: "wei",  label: "未时 (13:00-15:00)", hour: 13 },
  { value: "shen", label: "申时 (15:00-17:00)", hour: 15 },
  { value: "you",  label: "酉时 (17:00-19:00)", hour: 17 },
  { value: "xu",   label: "戌时 (19:00-21:00)", hour: 19 },
  { value: "hai",  label: "亥时 (21:00-23:00)", hour: 21 },
]
```

## 默认值配置

```javascript
defaultValues = {
  birth_year: 1990,
  birth_month: 5,
  birth_day: 15,
  birth_hour: 14,
  birth_minute: 30,
  gender: "male",
  shichen: "wei"
}
```

## API 接口（推测）

```
POST /api/v1/bazi/analyze
Body: {
  year: 1990,
  month: 5,
  day: 15,
  shichen: "wei",
  gender: "male",
  master: "huiming" | "mingxin" | "xuanzhen"
}
Response: {
  code: 0,
  data: {
    bazi: { ... },           // 八字信息
    wuxing: { ... },         // 五行分析
    shishen: { ... },        // 十神
    dayun: [ ... ],          // 大运
    liunian: [ ... ],        // 流年
    master_reading: "...",   // 师父开示
    personality: "...",      // 性格分析
    career: "...",           // 事业分析
    wealth: "...",           // 财运分析
    relationship: "...",     // 感情分析
    health: "...",           // 健康分析
  }
}
```

## 命理维度

```javascript
analysisDimensions = [
  { value: "personality", label: "性格" },
  { value: "career",      label: "事业" },
  { value: "wealth",      label: "财运" },
  { value: "relationship", label: "感情" },
  { value: "health",      label: "健康" },
]
```

## 付费产品关联

- **流年运势详批**：¥9.9（首单特惠）
  - 12 月逐月运势
  - 贵人/桃花/财禄提示
  - 师父开示
- **八字精批深度版**：¥19.9（明星产品）
  - 完整十神/大运/流年
  - 古籍引用
  - PDF 报告
- **两人合婚**：¥29.9（情感推荐）
  - 双方八字配对
  - 五行互补分析
  - 古籍参考

## UI 细节

- 年/月/日选择器：点击弹出选择面板
- 时辰下拉：标准 select 或自定义下拉
- 性别按钮组：两列，选中高亮
- "请师父排盘" 按钮：金色 CTA
- 排盘结果可能使用 Card + Tab 展示不同维度
