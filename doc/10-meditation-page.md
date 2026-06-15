# 静心禅坐功能详解

## 路由

`/meditation/` → 禅修音乐播放页面

## 页面结构

```
┌─────────────────────────────────────────┐
│  加载禅修曲目...                          │ ← 加载状态
├─────────────────────────────────────────┤
│  禅修曲目列表                             │
│  ┌─────────────────────────────────┐    │
│  │ 🎵 曲目1 - 时长                  │    │
│  │ 🎵 曲目2 - 时长                  │    │
│  │ 🎵 曲目3 - 时长                  │    │
│  └─────────────────────────────────┘    │
└─────────────────────────────────────────┘
```

## 曲目数据结构

```javascript
track = {
  id: "track_id",
  title: "曲目名称",
  subtitle: "副标题/描述",
  duration: 300,        // 秒
  icon: "🧘",          // emoji 图标
  url: "https://...",   // 音频文件 URL
}
```

## API 接口

```javascript
POST /api/v1/meditation/catalog
Body: {}
Response: {
  code: 0,
  data: {
    tracks: [
      { id, title, subtitle, duration, icon, url },
      ...
    ]
  }
}
```

## 音频播放系统

### PlayerProvider 上下文

全局音频播放状态管理，通过 Context 提供：

```javascript
{
  track: null | Track,        // 当前播放曲目
  playing: false,             // 是否播放中
  muted: false,               // 是否静音
  volume: 0.7,                // 音量 0-1
  elapsed: 0,                 // 已播放秒数
  duration: 0,                // 总时长秒数
  meritSeconds: 0,            // 功德计时（累计播放时长）
  audioRef: React.RefObject,  // audio 元素引用
  
  play: (track) => {},        // 播放指定曲目
  toggle: () => {},           // 播放/暂停切换
  stop: () => {},             // 停止播放
  setMuted: (muted) => {},    // 设置静音
  setVolume: (vol) => {},     // 设置音量
  seekTo: (percent) => {},    // 跳转到指定位置
  resetMerit: () => {},       // 重置功德计时
  flushMerit: () => number,   // 获取并重置功德秒数
  getLiveMerit: () => number, // 实时获取功德秒数
}
```

### 功德计时机制

- 播放时累计 `meritSeconds`
- 暂停时停止计时
- `flushMerit()` 获取累计值并重置
- `getLiveMerit()` 实时获取（含当前播放段）
- 最大计时上限：86400 秒（24小时）

### 音频事件处理

```javascript
onPlay:    → 设置 playing=true, 记录开始时间
onPause:   → 设置 playing=false, 累计播放时长
onEnded:   → 设置 playing=false, 累计播放时长, 自动播放下一首
onTimeUpdate: → 更新 elapsed
onLoadedMetadata: → 更新 duration
```

### 自动播放下一首

- 当前曲目播放结束后，自动随机选择其他曲目
- 排除当前曲目避免重复
- 从曲目列表中随机选取

## 迷你播放器（MiniPlayer）

非 meditation 页面时显示的浮动播放控件：

```
┌─────────────────────────────────────────┐
│  [旋转图标] 曲目名称          [播放/暂停] [关闭] │
│  正在播放 · 1:23 / 5:00                   │
│  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░  进度条            │
└─────────────────────────────────────────┘
```

- 固定左下角，不遮挡底部导航
- 图标旋转动画（播放时 18s 一圈）
- 进度条：`from-gold to-vermillion` 渐变
- 点击曲目名称跳转到 meditation 页面
- 弹窗打开时自动隐藏

## 顶栏音乐按钮

- 在 TopNav 右侧
- 未播放时：音乐图标，border `gold/25`
- 播放中：金色高亮 + `animate-ping` 脉冲圆环
- 点击切换播放/暂停
- 首次点击自动从曲目列表随机选一首播放

## UI 细节

- 曲目列表可能使用 Card 组件
- 每首曲目显示：emoji 图标、标题、副标题、时长
- 当前播放曲目高亮
- 背景可能有冥想相关视觉元素
