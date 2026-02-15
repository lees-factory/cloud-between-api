# 🌥 Cloud Between Us PRD v2

# PRD: Cloud Between Us

**감성 궁합 서비스 브랜드 & 제품 설계 문서**

---

## 0. 브랜드 정체성

### 브랜드 정의

**Cloud Between Us**는 연애 중인 두 사람 사이의 감정적 공기를 구름이라는 시각적 은유로 표현하는 감성 궁합 서비스다.

### 우리는 무엇이 아닌가

- ❌ MBTI 테스트
- ❌ 심리 분석 도구
- ❌ 과학 기반 성격 평가
- ❌ 점술/타로

✅ 우리는 **감정 해석기**다.

### 포지셔닝 문장

> **"We don't measure love. We visualize it."**
>

이 문장이 모든 방향성을 결정한다.

### 왜 "Cloud"인가

구름은:

- 형태가 고정되지 않는다
- 두 개가 만나면 섞이거나 충돌한다
- 부드럽지만 강력하다
- **감정과 닮았다**

"연애 감정"을 설명하기에 가장 적합한 메타포다.

---

## 1. 프로젝트 개요

### 1.1 목표

**감정을 숫자가 아닌 풍경으로 보여주는 서비스**

- 단순 궁합 퍼센트가 아닌 "하늘의 날씨"로 관계를 시각화
- IP 기반 세계관으로 브랜드 차별화
- 글로벌 확장 가능한 영어 브랜드

### 1.2 타겟 사용자

### 페르소나 1: 연애 초기 커플

- 나이: 20-30대
- 니즈: "우리 잘 맞을까?" 궁금증
- 기대: 재미있고 감성적인 결과

### 페르소나 2: 연애 콘텐츠 소비자

- 플랫폼: Instagram, TikTok
- 니즈: 공유 가능한 감성 콘텐츠
- 기대: 비주얼이 예쁜 결과 카드

### 1.3 성공 지표

- 테스트 완료율: 70% 이상
- 결과 공유율: 40% 이상
- 프리미엄 전환율: 8-12%

---

## 2. 브랜드 구조

### 2.1 브랜드 계층

**Brand:** Cloud Between Us

**Core Concept:** Cloud Chemistry

**IP Layer:** 6 Cloud Types

**Future Expansion:**

- Couple Cloud Avatar
- Anniversary Report
- AI Love Forecast

### 2.2 왜 글로벌 영어 브랜드인가

- 연애 콘텐츠는 글로벌 소비 강함
- TikTok, Instagram 바이럴은 영어가 유리
- 감성 단어(Cloud, Chemistry)가 직관적
- 리브랜딩 비용 방지

---

## 3. 비주얼 방향

### 3.1 디자인 철학

**"Soft but intentional."**

- 너무 귀엽지 않다
- 너무 심리학적이지 않다
- 너무 키치하지 않다
- **감정적이지만 세련된**

### 3.2 컬러 시스템

```css
/* Primary: 감정의 공기 */
--sky-blue: #A7D8F5;

/* Accent: 연애의 온기 */
--warm-peach: #FFC6A8;

/* Background: 부드러운 공간감 */
--off-white: #FAFAF8;

/* Text */
--text-dark: #111827;
--text-gray: #6B7280;
```

**왜 이 조합인가?**

- 차가운 파랑 + 따뜻한 피치 = 두 사람의 대비
- 궁합 서비스 특성상 대비와 균형이 중요

### 3.3 UI 스타일 가이드

```tsx
// 브랜드 UI 규칙
const BRAND_STYLE = {
  borderRadius: {
    card: '16px',
    button: '24px',
    large: '32px',
  },
  spacing: {
    section: '4rem',   // 충분한 여백
    card: '2rem',
  },
  shadow: {
    soft: '0 4px 24px rgba(0, 0, 0, 0.06)',  // 약하게
  },
  animation: {
    duration: '0.3s',  // 과하지 않게
    easing: 'ease-out',
  },
};
```

**금지 사항:**

- ❌ 과한 애니메이션 (우리는 게임이 아니다)
- ❌ 날카로운 모서리
- ❌ 고채도 네온 컬러
- ❌ 복잡한 그래디언트

---

## 4. 카피 전략

### 4.1 톤앤매너

- **Soft**: 공격적이지 않음
- **Intimate**: 개인적이고 친밀함
- **Slightly poetic**: 약간 시적인 표현
- **Short sentences**: 짧은 문장
- **No technical jargon**: 전문 용어 금지

### 4.2 핵심 카피

### 랜딩 페이지

**Hero Headline:**

> What kind of cloud are you in love?
>

**Sub:**

> Take the 2-minute love test and discover the chemistry between you.
>

**CTA:**

> Start the Test ☁️
>

**왜 이렇게 썼나?**

- "Personality"라는 단어 회피 → 심리 테스트 느낌 제거
- "Love" 직접 사용 → 명확한 타겟팅
- "2-minute" → 진입 장벽 낮춤
- "Chemistry" → 궁합 암시

### 궁합 결과 카피

**상단:**

> The cloud between you is glowing.
>

**퍼센트 아래:**

> Some clouds collide.
>
>
> Yours blend.
>

**왜 이렇게 쓰는가?**

- 직접적 확언 피함 (과학적 주장 아님)
- 시적인 언어 사용 (감성 유지)
- 공유 시 멋있어 보임

### 프리미엄 유도 카피

**Blur 영역 위:**

> Want to see how your love evolves?
>

**버튼:**

> Unlock Full Love Report – $6.99
>

**하단:**

> One-time payment. No subscription.
>

**전략:**

- 구독 공포 제거 ("한 번만" 강조)
- "Love Report" (심리 리포트 ❌)
- 가격 명시로 투명성 확보

---

## 5. IP 세계관: The Sky Lore

### 세계관 설정

```
하늘에는 여섯 가지 구름이 존재한다.
그들은 날씨를 만들지 않는다.
그들은 사랑의 공기를 만든다.

어떤 구름은 빛을 머금고,
어떤 구름은 비를 품고,
어떤 구름은 바람을 따라 흩어진다.

두 사람이 만날 때,
그들의 구름은 하늘에서 먼저 만난다.

그것이 바로
The Cloud Between Us.
```

### 6가지 Cloud Types

### ☀️ 햇살 (Sunlit) — The Warm Leader

**대표 단어:** Warmth · Direction · Loyalty · Radiance

**세계관 서사:**

```
햇살은 해를 가장 오래 품고 있는 구름이다.
이 구름은 빛을 통과시키지 않는다.
빛을 머금고 주변을 밝힌다.
사랑에 빠지면 길을 잃지 않게 하려 한다.
관계를 앞으로 움직이게 만든다.
```

**연애 특징:**

- 고백을 먼저 하는 구름
- 미래를 그리는 구름
- "우리"라는 말을 자주 쓰는 구름

**그림자:**
빛이 강해질수록 상대의 그림자를 보지 못할 수 있다.

---

### 🌫 안개 (Mist) — The Sensitive Soul

**대표 단어:** Sensitivity · Intuition · Depth · Fragility

**세계관 서사:**

```
안개는 해 뜨기 전 공기를 떠다닌다.
보이지 않지만 가장 많은 감정을 품고 있다.
사랑은 말보다 분위기다.
눈빛과 공기의 온도다.
```

**연애 특징:**

- 작은 변화도 알아차린다
- 말 대신 표정을 읽는다
- 깊이 연결되길 원한다

**그림자:**
감정을 너무 많이 흡수해 스스로 흐려질 수 있다.

---

### ⛈ 천둥 (Storm) — The Passion Spark

**대표 단어:** Intensity · Desire · Energy · Impulse

**세계관 서사:**

```
천둥은 하늘에서 가장 소리가 크다.
사랑을 시작할 때 번개처럼 빠르다.
감정은 숨기지 않는다. 폭발한다.
```

**연애 특징:**

- 빠르게 빠진다
- 빠르게 싸운다
- 빠르게 화해한다

**그림자:**
폭풍은 오래 지속되지 않는다. 지치게 만들 수 있다.

---

### 🌤 여명 (Dawn) — The Peace Keeper

**대표 단어:** Calm · Stability · Patience · Balance

**세계관 서사:**

```
여명은 바람을 거스르지 않는다.
하지만 사라지지도 않는다.
천천히 밝아오며 하늘을 오래 지킨다.
```

**연애 특징:**

- 싸움을 키우지 않는다
- 감정을 크게 드러내지 않는다
- 오래 가는 관계를 만든다

**그림자:**
너무 잔잔하면 열정을 잃을 수 있다.

---

### 🌪 바람 (Wild) — The Free Spirit

**대표 단어:** Freedom · Playfulness · Unpredictability · Spark

**세계관 서사:**

```
바람은 형태가 없다.
부는 방향으로 자유롭게 춤춘다.
사랑은 재미여야 한다.
숨 막히면 사라진다.
```

**연애 특징:**

- 갑자기 여행 제안
- 갑자기 고백
- 갑자기 거리두기

**그림자:**
구름이 너무 빠르면 붙잡히지 않는다.

---

### 🌥 그늘 (Shade) — The Quiet Anchor

**대표 단어:** Depth · Reliability · Composure · Devotion

**세계관 서사:**

```
그늘은 빛을 드러내지 않는다.
하지만 가장 오래 남는다.
말은 적지만 감정은 깊다.
누군가에게는 쉼터가 된다.
```

**연애 특징:**

- 행동으로 증명
- 쉽게 떠나지 않음
- 오래 버팀

**그림자:**
조용함이 차가움으로 오해받을 수 있다.

---

## 6. 15가지 커플 조합 서사

### 기상 현상 3가지

1. **Glow** – 따뜻하게 섞이는 사랑
2. **Rain** – 감정이 과해지는 사랑
3. **Thunder** – 강하게 충돌하는 사랑

---

### 1. ☀️ 햇살 × 🌫 안개

**"Morning Light Through Fog"**

**현상:** Glow

```
햇살의 따뜻한 빛이 안개의 감정을 천천히 녹인다.
안개는 이해받는다고 느끼고,
햇살은 보호하고 싶어진다.
```

**⚠️ 위험:** 빛이 너무 강하면 안개는 사라진다.

---

### 2. ☀️ 햇살 × ⛈ 천둥

**"Lightning at Noon"**

**현상:** Thunder

```
둘 다 강하다.
햇살은 방향을 잡고, 천둥은 속도를 올린다.
🔥 케미는 강렬하다.
💥 충돌도 강렬하다.
```

---

### 3. ☀️ 햇살 × 🌤 여명

**"Golden Afternoon"**

**현상:** Stable Glow

```
햇살이 앞으로 이끌고 여명이 균형을 맞춘다.
가장 "결혼형" 하늘.
```

**⚠️ 문제:** 열정이 식을 수 있음.

---

### 4. ☀️ 햇살 × 🌪 바람

**"Bright Winds"**

**현상:** Unpredictable Glow

```
햇살은 구조를 원하고 바람은 자유를 원한다.
사랑은 재미있다. 하지만 줄다리기다.
```

---

### 5. ☀️ 햇살 × 🌥 그늘

**"Royal Sky"**

**현상:** Deep Glow

```
둘 다 리더 기질. 겉은 단단하고 속은 깊다.
존중이 유지되면 아주 오래 간다.
```

---

### 6. 🌫 안개 × ⛈ 천둥

**"Heavy Rain"**

**현상:** Rain

```
천둥은 감정을 쏟고 안개는 그걸 전부 흡수한다.
강렬하지만 피로도가 높다.
```

---

### 7. 🌫 안개 × 🌤 여명

**"Silent Dawn"**

**현상:** Soft Glow

```
말이 많지 않아도 감정은 통한다.
조용하지만 깊다.
```

**⚠️:** 표현 부족 위험.

---

### 8. 🌫 안개 × 🌪 바람

**"Fog in the Wind"**

**현상:** Disappearing Sky

```
바람은 자유롭고 안개는 붙잡고 싶다.
균형이 어렵다.
```

---

### 9. 🌫 안개 × 🌥 그늘

**"Winter Morning"**

**현상:** Calm Glow

```
둘 다 조용하다. 감정은 깊지만 표현은 적다.
서로 이해하면 아주 단단하다.
```

---

### 10. ⛈ 천둥 × 🌤 여명

**"Storm Over Still Water"**

**현상:** Tension Sky

```
천둥은 흔들고 여명은 버틴다.
버티는 쪽이 지치면 관계가 무너진다.
```

---

### 11. ⛈ 천둥 × 🌪 바람

**"Hurricane Love"**

**현상:** Explosive Thunder

```
열정 × 자유. 속도 × 즉흥.
강렬하고 짧을 가능성 높음.
하지만 잊히지 않는다.
```

---

### 12. ⛈ 천둥 × 🌥 그늘

**"Thunder Against Stone"**

**현상:** Resistance

```
천둥은 밀어붙이고 그늘은 흔들리지 않는다.
존중하면 균형, 아니면 충돌.
```

---

### 13. 🌤 여명 × 🌪 바람

**"Soft Wind"**

**현상:** Uneven Sky

```
바람은 움직이고 여명은 머문다.
여명이 답답해질 수 있음.
```

---

### 14. 🌤 여명 × 🌥 그늘

**"Endless Horizon"**

**현상:** Long-Term Glow

```
안정 × 안정. 가장 오래 가는 구조.
드라마는 적다. 평화는 많다.
```

---

### 15. 🌪 바람 × 🌥 그늘

**"Tethered Wind"**

**현상:** Controlled Freedom

```
바람은 날고 그늘은 줄을 쥔다.
잘 맞으면 성숙한 균형.
잘못 맞으면 속박.
```

---

## 7. 기능 요구사항

### 7.1 필수 기능 (MVP)

### 7.1.1 랜딩 페이지

- Hero 섹션 (브랜드 카피)
- CTA 버튼 (테스트 시작)
- 6가지 Cloud Types 미리보기
- 푸터 (소셜 링크)

**수용 기준:**

- 로딩 3초 이내
- 모바일 최적화
- 브랜드 컬러 정확히 적용

---

### 7.1.2 테스트 플로우

- 12-15개 질문 (각 Cloud Type 판별)
- 프로그레스 바
- 부드러운 페이지 전환
- 뒤로 가기 가능

**수용 기준:**

- 질문 진행 상태 저장 (새로고침 대응)
- 2분 이내 완료 가능
- 접근성 (키보드 네비게이션)

---

### 7.1.3 결과 페이지 (무료 버전)

- 내 Cloud Type 표시
- 4가지 대표 단어
- 세계관 서사 (2-3문장)
- **블러 처리된 궁합 결과**
- 프리미엄 CTA

**수용 기준:**

- 결과 카드 이미지 다운로드 가능
- 소셜 공유 버튼 (Instagram, Twitter)
- Open Graph 메타태그 적용

---

### 7.1.4 결과 페이지 (프리미엄 버전)

- 커플 조합 서사 전체
- 15가지 조합 중 해당 조합 상세
- 기상 현상 비주얼
- 장점/위험 포인트
- PDF 다운로드

**수용 기준:**

- 결제 후 즉시 언락
- 영구 액세스 (재방문 가능)

---

### 7.2 선택 기능 (Phase 2)

- Couple Avatar 생성 (AI 이미지)
- Anniversary Report (연애 기간 기반)
- AI Love Forecast (미래 예측 콘텐츠)

---

## 8. 기술 요구사항

### 8.1 프론트엔드

```tsx
// 기술 스택
const TECH_STACK = {
  framework: 'SvelteKit',
  language: 'TypeScript',
  styling: 'Tailwind CSS + Custom CSS Variables',
  stateManagement: 'Svelte Stores',
  animation: 'CSS Transitions + Svelte transitions',
  routing: 'SvelteKit File-based Routing',
};
```

### 8.2 타입 정의

```tsx
// src/lib/types/cloud.ts

export type CloudType =
  | 'sunlit'    // 햇살
  | 'mist'      // 안개
  | 'storm'     // 천둥
  | 'dawn'      // 여명
  | 'wild'      // 바람
  | 'shade';    // 그늘

export type WeatherPhenomenon =
  | 'glow'
  | 'rain'
  | 'thunder';

export interface CloudProfile {
  type: CloudType;
  emoji: string;
  name: string;
  subtitle: string;
  keywords: [string, string, string, string];
  lore: string;
  traits: {
    strengths: string[];
    shadows: string[];
  };
}

export interface CoupleChemistry {
  user: CloudType;
  partner: CloudType;
  skyName: string;           // "Morning Light Through Fog"
  phenomenon: WeatherPhenomenon;
  narrative: string;
  warning: string | null;
}
```

### 8.3 디렉토리 구조

```
src/
├── lib/
│   ├── components/
│   │   ├── landing/
│   │   │   ├── HeroSection.svelte
│   │   │   ├── CloudTypesPreview.svelte
│   │   │   └── Footer.svelte
│   │   ├── test/
│   │   │   ├── QuestionCard.svelte
│   │   │   ├── ProgressBar.svelte
│   │   │   └── NavigationButtons.svelte
│   │   ├── result/
│   │   │   ├── CloudReveal.svelte
│   │   │   ├── SkyNarrative.svelte
│   │   │   ├── ChemistryCard.svelte (블러)
│   │   │   └── PremiumCTA.svelte
│   │   └── shared/
│   │       ├── CloudIcon.svelte
│   │       └── ShareButtons.svelte
│   ├── data/
│   │   ├── cloudProfiles.ts      # 6가지 Cloud 데이터
│   │   ├── questions.ts           # 테스트 질문
│   │   └── chemistryMatrix.ts     # 15가지 조합 서사
│   ├── stores/
│   │   └── testProgress.ts        # 진행 상태 관리
│   ├── utils/
│   │   ├── calculateCloud.ts      # Cloud Type 판별 로직
│   │   └── shareUtils.ts          # 소셜 공유 기능
│   └── types/
│       └── cloud.ts
├── routes/
│   ├── +page.svelte               # 랜딩
│   ├── test/
│   │   └── +page.svelte           # 테스트
│   ├── result/
│   │   ├── +page.ts               # 결과 로드
│   │   └── +page.svelte           # 결과 페이지
│   └── api/
│       └── payment/
│           └── +server.ts         # Stripe 결제
└── app.css                        # 브랜드 CSS Variables
```

### 8.4 성능 목표

- Lighthouse Score: 95+ (Performance)
- First Contentful Paint: < 1.5초
- Time to Interactive: < 3초
- 이미지 최적화: WebP 포맷
- Code Splitting: 자동 (SvelteKit)

### 8.5 접근성

- WCAG 2.1 AA 준수
- 키보드 네비게이션
- 스크린 리더 호환 (ARIA 속성)
- Color Contrast: 4.5:1 이상

---

## 9. 수익화 구조

### 9.1 프리미엄 모델

**가격:** $6.99 (일회성)

**포함 내용:**

- 전체 궁합 서사 언락
- PDF 다운로드
- 영구 액세스

**전략:**

- 구독 없음 (구독 공포 제거)
- 명확한 가격 (숨김 없음)
- 즉시 가치 제공

### 9.2 결제 플로우

```
1. 결과 페이지 도달
2. 블러 처리된 궁합 확인
3. "Unlock Full Love Report" CTA
4. Stripe Checkout
5. 결제 완료 → 즉시 리다이렉트
6. 전체 내용 표시
```

---

## 10. 개발 일정

### Phase 1 (2주) - MVP

- 랜딩 페이지
- 테스트 플로우
- 결과 페이지 (무료 버전)
- 브랜드 컬러/카피 적용

### Phase 2 (1주) - 프리미엄

- 결제 연동 (Stripe)
- 프리미엄 결과 페이지
- PDF 생성

### Phase 3 (1주) - 폴리싱

- 애니메이션 추가
- 소셜 공유 최적화
- SEO/OG 태그

---

## 11. 핵심 메시지

**이 서비스는 분석 도구가 아니다.**

**감정 경험 제품이다.**

우리는:

- 데이터를 파는 게 아니라
- **감정 해석**을 판다.

모든 코드, 모든 단어, 모든 픽셀이

이 철학을 반영해야 한다.

---

## 12. 브랜드 메시지 (복창용)

개발하기 전 항상 읽어라:

> Every love has a sky.
>
>
> What does yours look like?
>

---

**END OF DOCUMENT**