# 产品配置与定价

## 付费产品列表

```javascript
products = [
  {
    product_id: "single_liunian",
    name: "流年运势详批",
    price: 9.9,
    badge: "首单特惠",
    benefits: [
      "12 月逐月运势",
      "贵人 / 桃花 / 财禄提示",
      "师父开示"
    ]
  },
  {
    product_id: "single_bazi_deep",
    name: "八字精批深度版",
    price: 19.9,
    badge: "明星产品",
    benefits: [
      "完整十神 / 大运 / 流年",
      "古籍引用",
      "PDF 报告"
    ]
  },
  {
    product_id: "unlock_palmistry",
    name: "手相图解读",
    price: 19.9,
    badge: "新功能",
    benefits: [
      "拍照上传",
      "手纹细看",
      "手纹命理详解"
    ]
  },
  {
    product_id: "single_hehun",
    name: "两人合婚",
    price: 29.9,
    badge: "情感推荐",
    benefits: [
      "双方八字配对",
      "五行互补分析",
      "古籍参考"
    ]
  },
  {
    product_id: "single_naming_pro",
    name: "宝宝起名 VIP",
    price: 49.9,
    badge: "热销",
    benefits: [
      "30 个候选名",
      "音韵 / 笔画 / 五行评分",
      "典故出处"
    ]
  },
  {
    product_id: "single_company",
    name: "公司起名",
    price: 99.9,
    badge: "企业版",
    benefits: [
      "行业五行匹配",
      "品牌寓意",
      "5 个候选方案"
    ]
  }
]
```

## 定价策略

| 产品 | 价格 | 定位 |
|------|------|------|
| 流年运势 | ¥9.9 | 引流产品（首单特惠） |
| 八字深度版 | ¥19.9 | 核心产品（明星产品） |
| 手相解读 | ¥19.9 | 新功能推广 |
| 合婚 | ¥29.9 | 情感场景 |
| 宝宝起名 VIP | ¥49.9 | 高价值服务 |
| 公司起名 | ¥99.9 | 企业级服务 |

## 价格区间

- 最低：¥6.6（祈福灯）
- 最高：¥99.9（公司起名）
- 主力区间：¥9.9 ~ ¥29.9

## 分享返佣系统

### 机制

1. 用户分享专属链接（含 `?invite=<lucky_code>` 参数）
2. 朋友通过链接访问并付款
3. 奖励自动记入分享者的"福报金"

### 分享链接格式

```
https://putiyuan.pages.dev/<当前页面>?invite=<用户lucky_code>
```

### 分享渠道

- 微信：直接发链接 / 长按二维码识别
- 抖音私信：复制链接后让对方在浏览器打开
- 系统分享：调用 Web Share API
- 复制链接 / 复制标题和链接

### 二维码

- 使用 `qrcode` 库生成 Data URL
- 尺寸：280px
- 颜色：深色 `#1A1410` / 浅色 `#F5F0E8`
- 纠错级别：M

## 邀请码系统

```javascript
// 应用邀请码
POST /referral/apply
Body: { invite_code: "XXXX", device_id: "..." }

// 获取邀请信息
POST /referral/me

// 提现
POST /referral/withdraw
Body: { amount, note }
```

## 产品展示卡片（更多页面）

```javascript
moreProducts = [
  { href: "/bazi",       label: "八字精批", icon: "compass" },
  { href: "/dream",      label: "周公解梦", icon: "scroll" },
  { href: "/palmistry",  label: "看手相",   icon: "hand" },
  { href: "/naming",     label: "宝宝起名", icon: "baby" },
  { href: "/divination", label: "六爻占卜", icon: "yin-yang" },
  { href: "/meditation", label: "静心禅坐", icon: "lotus" },
]
```

## 五种起名风格

```javascript
namingStyles = ["诗意", "刚毅", "儒雅", "清逸", "典雅", "温润"]
```
