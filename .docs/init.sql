-- =============================================
-- Cloud Between Us - Database Schema
-- cloud → persona 네이밍 적용
-- =============================================

-- 사용자 기본 정보
CREATE TABLE cloud_between.user_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    social_id TEXT,
    social_provider VARCHAR(20),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT,
    profile_image_url TEXT,
    is_paid BOOLEAN DEFAULT FALSE,
    last_login_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(social_id, social_provider)
);

COMMENT ON TABLE cloud_between.user_profiles IS '사용자 기본 정보 및 결제 상태 관리 테이블';
COMMENT ON COLUMN cloud_between.user_profiles.social_id IS '소셜 로그인 제공자의 고유 ID';
COMMENT ON COLUMN cloud_between.user_profiles.social_provider IS '소셜 로그인 제공자 (GOOGLE, APPLE, KAKAO)';
COMMENT ON COLUMN cloud_between.user_profiles.email IS '사용자 이메일 (로그인 식별자)';
COMMENT ON COLUMN cloud_between.user_profiles.password_hash IS '암호화된 비밀번호';
COMMENT ON COLUMN cloud_between.user_profiles.is_paid IS '유료 컨텐츠(상세 결과) 접근 권한 여부';

-- =============================================
-- 심리 테스트 메타데이터
-- =============================================

-- 테스트 스텝 (12개 카테고리)
CREATE TABLE cloud_between.test_steps (
    id INT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    emoji VARCHAR(10),
    order_index INT NOT NULL,
    locale VARCHAR(5) DEFAULT 'ko',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.test_steps IS '심리 테스트 스텝(카테고리) 메타데이터';
COMMENT ON COLUMN cloud_between.test_steps.title IS '스텝 제목 (사랑의 시작, 감정 표현 등)';

-- 테스트 질문 (48개)
CREATE TABLE cloud_between.test_questions (
    id SERIAL PRIMARY KEY,
    step_id INT NOT NULL REFERENCES cloud_between.test_steps(id),
    question_text TEXT NOT NULL,
    options JSONB NOT NULL,
    locale VARCHAR(5) DEFAULT 'ko',
    order_index INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.test_questions IS '심리 테스트 질문 및 선택지 관리 테이블';
COMMENT ON COLUMN cloud_between.test_questions.step_id IS '소속 스텝 ID (FK → test_steps)';
COMMENT ON COLUMN cloud_between.test_questions.options IS '선택지 배열: [{"text": "...", "cloudType": "..."}]';
COMMENT ON COLUMN cloud_between.test_questions.locale IS '다국어 코드 (ko, en 등)';
COMMENT ON COLUMN cloud_between.test_questions.order_index IS '질문 노출 순서';

-- =============================================
-- 페르소나 (성격 유형) 프로필
-- =============================================

CREATE TABLE cloud_between.persona_profiles (
    type_key VARCHAR(20),
    locale VARCHAR(5) DEFAULT 'ko',
    emoji VARCHAR(10),
    name VARCHAR(50) NOT NULL,
    subtitle VARCHAR(100),
    keywords JSONB,
    lore TEXT,
    strengths JSONB,
    shadows JSONB,
    PRIMARY KEY (type_key, locale)
);

COMMENT ON TABLE cloud_between.persona_profiles IS '각 페르소나 유형별 상세 설명 및 특징 정보';
COMMENT ON COLUMN cloud_between.persona_profiles.type_key IS '페르소나 유형 식별자 (sunlit, mist, storm, dawn, shade, wild)';
COMMENT ON COLUMN cloud_between.persona_profiles.keywords IS '페르소나 특징 키워드 배열';
COMMENT ON COLUMN cloud_between.persona_profiles.lore IS '페르소나에 담긴 이야기/설명';

-- =============================================
-- 궁합 매트릭스
-- =============================================

CREATE TABLE cloud_between.chemistry_matrix (
    id SERIAL PRIMARY KEY,
    persona_type_1 VARCHAR(20) NOT NULL,
    persona_type_2 VARCHAR(20) NOT NULL,
    sky_name VARCHAR(100),
    phenomenon VARCHAR(50),
    narrative TEXT,
    warning TEXT,
    UNIQUE(persona_type_1, persona_type_2)
);

COMMENT ON TABLE cloud_between.chemistry_matrix IS '두 페르소나 유형 간의 궁합 정보 매트릭스';
COMMENT ON COLUMN cloud_between.chemistry_matrix.phenomenon IS '함께 있을 때 나타나는 현상 (glow, rain 등)';

-- =============================================
-- 사용자 테스트 결과
-- =============================================

CREATE TABLE cloud_between.user_test_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    result_persona_type VARCHAR(20) NOT NULL,
    answers JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.user_test_results IS '사용자 심리 테스트 수행 결과 기록';
COMMENT ON COLUMN cloud_between.user_test_results.result_persona_type IS '최종 판정된 페르소나 유형';

-- =============================================
-- ALTER: updated_at 컬럼 추가
-- =============================================

ALTER TABLE cloud_between.user_profiles
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

ALTER TABLE cloud_between.test_steps
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

ALTER TABLE cloud_between.test_questions
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

ALTER TABLE cloud_between.persona_profiles
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

ALTER TABLE cloud_between.chemistry_matrix
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

ALTER TABLE cloud_between.user_test_results
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

-- =============================================
-- 테스트 데이터 초기화 및 재삽입
-- =============================================

DELETE FROM cloud_between.test_questions;
DELETE FROM cloud_between.test_steps;

-- 스텝 데이터 (12개)
INSERT INTO cloud_between.test_steps (id, title, emoji, order_index, locale) VALUES
(1, '첫 빛이 드는 순간', '🌅', 1, 'ko'),
(2, '마음의 언어', '💬', 2, 'ko'),
(3, '폭풍우가 치는 날', '⚡', 3, 'ko'),
(4, '흐르는 일상', '☕', 4, 'ko'),
(5, '나와 우리 사이의 여백', '🕊️', 5, 'ko'),
(6, '먼 미래의 스카이라인', '🔮', 6, 'ko'),
(7, '소유와 질투의 그림자', '👀', 7, 'ko'),
(8, '살결과 숨결', '💫', 8, 'ko'),
(9, '흔들리는 하늘 아래서', '🌊', 9, 'ko'),
(10, '사랑의 빛깔', '💝', 10, 'ko'),
(11, '혼자와 둘 사이', '🎭', 11, 'ko'),
(12, '영원히 머무를 하늘', '🌡️', 12, 'ko');

-- Step 1: 첫 빛이 드는 순간
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(1, '마음이 가는 사람을 발견했을 때, 당신의 첫 움직임은?', '[{"text":"먼저 다가가 따스한 대화를 건넨다","cloudType":"sunlit"},{"text":"상대의 결을 살피며 조용히 기다린다","cloudType":"mist"},{"text":"강렬한 이끌림을 숨기지 않고 표현한다","cloudType":"storm"},{"text":"자연스럽게 곁에 스며들어 친구가 된다","cloudType":"dawn"}]'::jsonb, 'ko', 1),
(1, '누군가를 처음 마주할 때, 당신이 풍기는 공기는?', '[{"text":"모두를 환하게 만드는 밝은 에너지","cloudType":"sunlit"},{"text":"말수는 적어도 깊이 공감하는 눈빛","cloudType":"mist"},{"text":"한 번 보면 잊히지 않는 뚜렷한 존재감","cloudType":"storm"},{"text":"누구든 편히 쉴 수 있는 부드러운 여백","cloudType":"dawn"}]'::jsonb, 'ko', 2),
(1, '연애에서 당신이 가장 지키고 싶은 가치는?', '[{"text":"서로를 향한 신뢰와 선명한 미래","cloudType":"sunlit"},{"text":"마음 깊은 곳까지 닿는 이해","cloudType":"mist"},{"text":"멈추지 않는 뜨거운 불꽃","cloudType":"storm"},{"text":"바람처럼 자유롭되 변치 않는 안식","cloudType":"shade"}]'::jsonb, 'ko', 3),
(1, '사랑에 물드는 당신의 속도는?', '[{"text":"신중하게 관찰하다 확신이 들면 단단하게","cloudType":"sunlit"},{"text":"이슬이 맺히듯 천천히, 어느샌가 깊게","cloudType":"mist"},{"text":"벼락이 치듯 순식간에, 거부할 수 없게","cloudType":"storm"},{"text":"동이 트듯 나도 모르는 사이에 서서히","cloudType":"dawn"}]'::jsonb, 'ko', 4);

-- Step 2: 마음의 언어
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(2, '''사랑해''라는 고백이 당신의 입술을 떠나는 시점은?', '[{"text":"상대가 내 확신을 느낄 수 있도록 먼저","cloudType":"sunlit"},{"text":"그의 마음을 확인하고 조심스럽게 화답하며","cloudType":"mist"},{"text":"참을 수 없는 감정이 차오르는 바로 그 순간","cloudType":"storm"},{"text":"말보다는 따뜻한 눈빛과 행동으로 대신하며","cloudType":"shade"}]'::jsonb, 'ko', 5),
(2, '마음속에 먹구름이 끼었을 때, 당신은 어떻게 하나요?', '[{"text":"투명하게 이야기하고 함께 해법을 찾는다","cloudType":"sunlit"},{"text":"혼자만의 방에 들어가 감정을 가라앉힌다","cloudType":"mist"},{"text":"서운함을 숨기지 않고 격렬히 표현한다","cloudType":"storm"},{"text":"시간이 흘러 자연히 흩어지길 기다린다","cloudType":"dawn"}]'::jsonb, 'ko', 6),
(2, '당신이 사랑을 건네는 가장 선명한 방법은?', '[{"text":"미래를 약속하는 말과 든든한 리드","cloudType":"sunlit"},{"text":"세심한 배려가 담긴 작은 쪽지와 선물","cloudType":"mist"},{"text":"숨 막힐 듯 강렬한 스킨십과 열정","cloudType":"storm"},{"text":"그저 묵묵히 곁을 지켜주는 존재감","cloudType":"shade"}]'::jsonb, 'ko', 7),
(2, '소중한 사람이 비를 맞고 있다면 당신은?', '[{"text":"우산을 씌워주며 갈 길을 안내한다","cloudType":"sunlit"},{"text":"함께 비를 맞으며 그의 눈물에 공감한다","cloudType":"mist"},{"text":"비를 멈추게 할 방법을 찾아 온몸으로 부딪친다","cloudType":"storm"},{"text":"말없이 겉옷을 벗어준 뒤 등을 토닥인다","cloudType":"shade"}]'::jsonb, 'ko', 8);

-- Step 3: 폭풍우가 치는 날
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(3, '의견 차이로 관계에 바람이 거세질 때 당신은?', '[{"text":"흐트러진 구름을 정리하듯 대화를 시도한다","cloudType":"sunlit"},{"text":"바람이 잦아들 때까지 생각에 잠긴다","cloudType":"mist"},{"text":"감정의 번개를 숨기지 않고 터뜨린다","cloudType":"storm"},{"text":"갈등 자체가 생기지 않도록 미리 피한다","cloudType":"dawn"}]'::jsonb, 'ko', 9),
(3, '상대가 큰 잘못을 했을 때, 당신의 첫 반응은?', '[{"text":"잘못된 부분을 명확히 짚고 사과를 요구한다","cloudType":"sunlit"},{"text":"깊게 상처받고 마음의 문을 잠근다","cloudType":"mist"},{"text":"화가 풀릴 때까지 감정을 쏟아낸다","cloudType":"storm"},{"text":"이유가 있었을 거라며 애써 이해해본다","cloudType":"dawn"}]'::jsonb, 'ko', 10),
(3, '화해의 손길을 먼저 내미는 쪽은?', '[{"text":"우리의 시간을 위해 주로 내가 먼저","cloudType":"sunlit"},{"text":"상대의 태도를 지켜보다 조심스럽게","cloudType":"mist"},{"text":"싸움만큼이나 뜨겁고 빠르게 화해를 시도","cloudType":"storm"},{"text":"시간이 지나 어색함이 사라지길 기다리며","cloudType":"dawn"}]'::jsonb, 'ko', 11),
(3, '다툼 끝에 혼자 남겨졌을 때 드는 생각은?', '[{"text":"우리가 왜 싸웠는지, 다음엔 어떡할지 분석한다","cloudType":"sunlit"},{"text":"그가 했던 말들을 곱씹으며 눈물 흘린다","cloudType":"mist"},{"text":"답답함에 어디론가 훌쩍 떠나고 싶다","cloudType":"wild"},{"text":"이 또한 지나갈 일이라며 마음을 비운다","cloudType":"shade"}]'::jsonb, 'ko', 12);

-- Step 4: 흐르는 일상
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(4, '기다리던 주말, 연인과 함께하고 싶은 모습은?', '[{"text":"미리 짜놓은 완벽한 계획대로 움직이기","cloudType":"sunlit"},{"text":"둘만의 아늑한 공간에서 조용히 머물기","cloudType":"mist"},{"text":"발길 닿는 대로 떠나는 즉흥적인 모험","cloudType":"wild"},{"text":"익숙하고 편안한 곳에서 쉬어가는 하루","cloudType":"dawn"}]'::jsonb, 'ko', 13),
(4, '상대의 아주 작은 기분 변화를 알아차리나요?', '[{"text":"표정만 봐도 무엇이 필요한지 바로 안다","cloudType":"sunlit"},{"text":"공기의 온도만으로도 그의 슬픔을 느낀다","cloudType":"mist"},{"text":"말해주지 않으면 가끔 놓치기도 한다","cloudType":"storm"},{"text":"그저 곁에서 편안하게 만들어주려 노력한다","cloudType":"shade"}]'::jsonb, 'ko', 14),
(4, '반복되는 데이트 루틴에 대해 어떻게 생각하나요?', '[{"text":"우리의 역사가 쌓이는 것 같아 든든하다","cloudType":"sunlit"},{"text":"익숙해서 좋지만 가끔은 낯선 설렘을 꿈꾼다","cloudType":"mist"},{"text":"새로운 자극이 없으면 금방 지루해진다","cloudType":"wild"},{"text":"변함없는 평온함이 가장 큰 행복이다","cloudType":"shade"}]'::jsonb, 'ko', 15),
(4, '두 사람만의 기념일을 맞이하는 당신의 자세는?', '[{"text":"기억에 남을 만한 특별한 이벤트를 준비한다","cloudType":"sunlit"},{"text":"진심을 꾹꾹 눌러 담은 편지를 전한다","cloudType":"mist"},{"text":"가장 화려하고 즐거운 곳으로 떠난다","cloudType":"storm"},{"text":"맛있는 음식을 먹으며 소박하게 축하한다","cloudType":"shade"}]'::jsonb, 'ko', 16);

-- Step 5: 나와 우리 사이의 여백
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(5, '연애 중에도 당신에게 꼭 필요한 ''혼자만의 시간''은?', '[{"text":"성장을 위해 계획적으로 확보해야 하는 시간","cloudType":"sunlit"},{"text":"감정의 과부하를 식히기 위한 필수적인 고립","cloudType":"mist"},{"text":"내 마음대로 자유를 만끽하고 싶은 순간","cloudType":"wild"},{"text":"그저 아무 생각 없이 멍하게 머무는 휴식","cloudType":"shade"}]'::jsonb, 'ko', 17),
(5, '연인이 갑자기 혼자 있고 싶다고 한다면?', '[{"text":"혹시 내가 뭘 잘못했는지 이유를 묻는다","cloudType":"sunlit"},{"text":"나를 멀리하는 것 같아 왠지 서글퍼진다","cloudType":"mist"},{"text":"나도 내 시간을 즐기면 된다고 생각한다","cloudType":"wild"},{"text":"그의 마음을 이해하며 기꺼이 내버려 둔다","cloudType":"shade"}]'::jsonb, 'ko', 18),
(5, '연락의 빈도가 사랑의 척도라고 생각하나요?', '[{"text":"어느 정도의 규칙적인 연락은 신뢰의 기본이다","cloudType":"sunlit"},{"text":"마음이 연결되어 있다면 횟수는 중요치 않다","cloudType":"mist"},{"text":"생각날 때마다 뜨겁게 쏟아붓는 편이다","cloudType":"storm"},{"text":"각자의 생활을 존중하며 여유 있게 하고 싶다","cloudType":"wild"}]'::jsonb, 'ko', 19),
(5, '상대에게 나의 모든 일상을 공유하고 싶나요?', '[{"text":"중요한 계획과 일과는 모두 알고 싶고 알리고 싶다","cloudType":"sunlit"},{"text":"나의 깊은 내면은 가끔 나만의 비밀로 남기고 싶다","cloudType":"mist"},{"text":"그때그때 기분과 재미있는 일들 위주로 공유한다","cloudType":"wild"},{"text":"상대가 묻는다면 편하게 대답해준다","cloudType":"shade"}]'::jsonb, 'ko', 20);

-- Step 6: 먼 미래의 스카이라인
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(6, '연애 초기, 당신은 어디까지 내다보나요?', '[{"text":"함께 집을 꾸미고 미래를 설계하는 상상까지","cloudType":"sunlit"},{"text":"이 사람이 나의 영혼을 알아줄지 깊게 고민","cloudType":"mist"},{"text":"지금 이 순간의 강렬한 스파크에 집중","cloudType":"storm"},{"text":"일단 흐르는 대로 지켜보며 천천히","cloudType":"dawn"}]'::jsonb, 'ko', 21),
(6, '결혼이나 동거 같은 구체적인 약속에 대한 생각은?', '[{"text":"관계가 안정되려면 명확한 목표가 필요하다","cloudType":"sunlit"},{"text":"언젠가 찾아올 운명적인 과정이라 믿는다","cloudType":"mist"},{"text":"구속받는 것 같아 가끔은 망설여진다","cloudType":"wild"},{"text":"때가 되면 자연스럽게 이루어질 일이다","cloudType":"shade"}]'::jsonb, 'ko', 22),
(6, '우리가 함께 늙어가는 모습을 상상하면?', '[{"text":"든든하고 단단한 뿌리를 내린 나무 같다","cloudType":"sunlit"},{"text":"서로의 주름까지 사랑하는 애틋한 풍경","cloudType":"mist"},{"text":"여전히 손잡고 새로운 세상을 누비는 모험가","cloudType":"wild"},{"text":"오랜 친구처럼 편안하게 마주 앉은 모습","cloudType":"shade"}]'::jsonb, 'ko', 23),
(6, '만약 장거리 연애를 해야 하는 상황이 온다면?', '[{"text":"계획적인 방문과 연락으로 충분히 극복한다","cloudType":"sunlit"},{"text":"그리움에 사무쳐 매일 밤 잠 못 이룰 것 같다","cloudType":"mist"},{"text":"몸이 멀어지면 마음도 멀어질까 두렵다","cloudType":"storm"},{"text":"신뢰만 있다면 각자의 자리에서 잘 지낼 수 있다","cloudType":"shade"}]'::jsonb, 'ko', 24);

-- Step 7: 소유와 질투의 그림자
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(7, '연인이 낯선 이성과 다정하게 웃고 있다면?', '[{"text":"기분이 상하지만 나중에 조용히 대화로 푼다","cloudType":"sunlit"},{"text":"말 못 할 서운함이 가슴 속에 차곡차곡 쌓인다","cloudType":"mist"},{"text":"그 즉시 불편한 기색을 내비치며 개입한다","cloudType":"storm"},{"text":"그럴 수도 있다며 쿨하게 넘기려 노력한다","cloudType":"wild"}]'::jsonb, 'ko', 25),
(7, '상대의 과거 연애가 궁금해질 때 당신은?', '[{"text":"관계를 위해 알아야 할 부분만 질문한다","cloudType":"sunlit"},{"text":"상상만으로도 아파서 아예 묻지 않는다","cloudType":"mist"},{"text":"하나부터 열까지 모두 다 알고 싶다","cloudType":"storm"},{"text":"지나간 일은 궁금해하지 않는 것이 예의다","cloudType":"dawn"}]'::jsonb, 'ko', 26),
(7, '연인이 나를 친구들에게 소개하지 않는다면?', '[{"text":"진지한 관계가 아닌지 의심하고 이유를 묻는다","cloudType":"sunlit"},{"text":"내가 부끄러운 존재인가 싶어 자존감이 떨어진다","cloudType":"mist"},{"text":"자존심이 상해서 크게 한판 싸운다","cloudType":"storm"},{"text":"그만의 사정이 있겠거니 하고 기다려준다","cloudType":"shade"}]'::jsonb, 'ko', 27),
(7, '당신이 생각하는 사랑의 ''소유''란?', '[{"text":"서로의 일상을 책임지는 기분 좋은 구속","cloudType":"sunlit"},{"text":"마음의 깊이가 같아야만 성립되는 것","cloudType":"mist"},{"text":"너는 내 것, 나는 네 것이라는 강렬한 선언","cloudType":"storm"},{"text":"서로를 소유하기보다 나란히 걷는 것","cloudType":"wild"}]'::jsonb, 'ko', 28);

-- Step 8: 살결과 숨결
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(8, '당신에게 스킨십은 어떤 의미인가요?', '[{"text":"우리의 친밀감을 확인하는 가장 확실한 도구","cloudType":"sunlit"},{"text":"백 마디 말보다 깊은 정서적 교감의 순간","cloudType":"mist"},{"text":"사랑을 타오르게 하는 본능적인 에너지","cloudType":"storm"},{"text":"포근하고 안정적인 애정의 확인","cloudType":"shade"}]'::jsonb, 'ko', 29),
(8, '연인과 밤새 대화를 나누는 것에 대해?', '[{"text":"가치관을 맞추어가는 꼭 필요한 시간이다","cloudType":"sunlit"},{"text":"영혼이 맞닿는 듯한 가장 행복한 순간이다","cloudType":"mist"},{"text":"가끔은 좋지만 잠도 중요하다고 생각한다","cloudType":"shade"},{"text":"말보다는 같이 무언가를 하는 게 더 좋다","cloudType":"wild"}]'::jsonb, 'ko', 30),
(8, '사람들이 많은 곳에서의 애정표현은?', '[{"text":"적당한 선에서는 자랑스럽고 좋다","cloudType":"sunlit"},{"text":"부끄러워서 단둘이 있을 때만 하고 싶다","cloudType":"mist"},{"text":"주변 시선 따위 신경 쓰지 않고 표현한다","cloudType":"storm"},{"text":"상대가 원한다면 맞춰주는 편이다","cloudType":"dawn"}]'::jsonb, 'ko', 31),
(8, '자기 전, 마지막 연락은 어떤 모습인가요?', '[{"text":"내일의 안녕을 빌어주는 다정한 약속","cloudType":"sunlit"},{"text":"꿈속에서도 만나길 바라는 애틋한 인사","cloudType":"mist"},{"text":"오늘의 뜨거웠던 마음을 전하는 고백","cloudType":"storm"},{"text":"오늘 하루도 수고했다는 편안한 위로","cloudType":"shade"}]'::jsonb, 'ko', 32);

-- Step 9: 흔들리는 하늘 아래서
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(9, '관계가 예전 같지 않다고 느껴질 때 당신은?', '[{"text":"문제를 분석하고 해결하기 위해 대화를 제안한다","cloudType":"sunlit"},{"text":"혼자 가슴 앓이 하며 관계의 끝을 두려워한다","cloudType":"mist"},{"text":"변한 태도에 대해 화를 내며 긴장감을 조성한다","cloudType":"storm"},{"text":"다시 좋아질 때까지 묵묵히 자리를 지킨다","cloudType":"shade"}]'::jsonb, 'ko', 33),
(9, '상대방이 이별을 암시하는 말을 한다면?', '[{"text":"우리가 노력할 수 있는 부분을 마지막으로 설득한다","cloudType":"sunlit"},{"text":"무너져 내리지만 그의 뜻을 존중하려 애쓴다","cloudType":"mist"},{"text":"잡거나 혹은 냉정하게 돌아서거나, 극단적으로 반응","cloudType":"storm"},{"text":"올 것이 왔다는 듯 담담하게 받아들이려 한다","cloudType":"shade"}]'::jsonb, 'ko', 34),
(9, '연애 중 가장 용서할 수 없는 것은?', '[{"text":"거짓말과 신뢰를 저버리는 행위","cloudType":"sunlit"},{"text":"나를 더 이상 사랑하지 않는다는 차가운 눈빛","cloudType":"mist"},{"text":"나를 통제하려 하거나 자유를 억압하는 것","cloudType":"wild"},{"text":"무례한 태도와 나를 무시하는 말들","cloudType":"storm"}]'::jsonb, 'ko', 35),
(9, '헤어진 후, 당신의 이별 극복 방식은?', '[{"text":"바쁘게 일상에 집중하며 미련을 지워낸다","cloudType":"sunlit"},{"text":"슬픈 노래를 들으며 마음껏 슬퍼하고 추억한다","cloudType":"mist"},{"text":"새로운 사람을 만나거나 낯선 곳으로 떠난다","cloudType":"wild"},{"text":"시간이 약이라 믿으며 조용히 견뎌낸다","cloudType":"shade"}]'::jsonb, 'ko', 36);

-- Step 10: 사랑의 빛깔
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(10, '가장 사랑받고 있다고 느끼는 순간은?', '[{"text":"나와의 미래를 진지하게 이야기해줄 때","cloudType":"sunlit"},{"text":"나의 작은 습관 하나까지 기억해줄 때","cloudType":"mist"},{"text":"불꽃 튀는 눈빛으로 나를 갈구할 때","cloudType":"storm"},{"text":"말없이 안아주는 품에서 평온함을 느낄 때","cloudType":"shade"}]'::jsonb, 'ko', 37),
(10, '당신이 줄 수 있는 가장 큰 사랑은?', '[{"text":"상대의 인생을 더 나은 방향으로 이끌어주는 것","cloudType":"sunlit"},{"text":"세상 누구보다 그를 깊이 이해하고 편들어주는 것","cloudType":"mist"},{"text":"나의 모든 열정과 생명력을 쏟아붓는 것","cloudType":"storm"},{"text":"언제든 돌아와 쉴 수 있는 든든한 나무가 되는 것","cloudType":"shade"}]'::jsonb, 'ko', 38),
(10, '예상치 못한 깜짝 선물에 대한 생각은?', '[{"text":"나를 생각한 정성이 고마워 감동한다","cloudType":"sunlit"},{"text":"그의 마음이 느껴져 눈물이 날 만큼 좋다","cloudType":"mist"},{"text":"짜릿하고 즐거운 자극이 되어 기쁘다","cloudType":"wild"},{"text":"고맙지만 가끔은 부담스럽기도 하다","cloudType":"dawn"}]'::jsonb, 'ko', 39),
(10, '연인에게 꼭 듣고 싶은 한 마디는?', '[{"text":"\"네 덕분에 내 삶이 더 의미 있어졌어\"","cloudType":"sunlit"},{"text":"\"세상에서 나를 가장 잘 아는 건 너뿐이야\"","cloudType":"mist"},{"text":"\"너만 보면 심장이 멈출 것 같아\"","cloudType":"storm"},{"text":"\"너랑 있으면 마음이 정말 편안해\"","cloudType":"shade"}]'::jsonb, 'ko', 40);

-- Step 11: 혼자와 둘 사이
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(11, '친구들과의 모임에 연인을 데려가는 것은?', '[{"text":"나의 사람들을 소개하는 기쁜 일이다","cloudType":"sunlit"},{"text":"신경 쓸 게 많아져서 조금 피곤하다","cloudType":"mist"},{"text":"함께 어울려 노는 게 즐거워 자주 한다","cloudType":"wild"},{"text":"상대가 원한다면 기꺼이 동행한다","cloudType":"dawn"}]'::jsonb, 'ko', 41),
(11, '서로의 취미가 전혀 다르다면?', '[{"text":"서로의 취미를 함께 배울 수 있도록 계획한다","cloudType":"sunlit"},{"text":"취향이 안 맞으면 대화가 안 통할까 걱정된다","cloudType":"mist"},{"text":"각자 즐기면 된다, 다름이 매력이다","cloudType":"wild"},{"text":"조금씩 양보하며 같이 할 접점을 찾는다","cloudType":"shade"}]'::jsonb, 'ko', 42),
(11, '핸드폰 비밀번호를 공유하는 것에 대해?', '[{"text":"숨길 게 없다면 공유해도 상관없다","cloudType":"sunlit"},{"text":"신뢰의 문제라고 생각하지만 왠지 꺼림칙하다","cloudType":"mist"},{"text":"절대 안 된다, 개인의 프라이버시다","cloudType":"wild"},{"text":"상대가 원하면 보여줄 수 있다","cloudType":"dawn"}]'::jsonb, 'ko', 43),
(11, '연인이 나를 닮아가는 모습을 볼 때?', '[{"text":"우리가 하나가 된 것 같아 뿌듯하다","cloudType":"sunlit"},{"text":"나의 안 좋은 점까지 닮을까 봐 걱정된다","cloudType":"mist"},{"text":"나의 색깔이 옅어지는 것 같아 묘한 기분이다","cloudType":"wild"},{"text":"그저 사랑스럽고 고마운 마음이다","cloudType":"shade"}]'::jsonb, 'ko', 44);

-- Step 12: 영원히 머무를 하늘
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(12, '당신의 사랑을 날씨로 표현한다면?', '[{"text":"구름 한 점 없이 맑고 쨍한 정오","cloudType":"sunlit"},{"text":"모든 것을 포근하게 감싸는 새벽 안개","cloudType":"mist"},{"text":"번쩍이는 번개 뒤에 찾아오는 뜨거운 소나기","cloudType":"storm"},{"text":"한낮의 열기를 식혀주는 울창한 나무 그늘","cloudType":"shade"}]'::jsonb, 'ko', 45),
(12, '오랜 시간이 흐른 뒤, 우리 사이는 어떤 모습일까요?', '[{"text":"더욱 단단한 성벽이 된 신뢰의 관계","cloudType":"sunlit"},{"text":"말하지 않아도 서로를 읽어내는 깊은 공명","cloudType":"mist"},{"text":"여전히 서로를 탐닉하는 뜨거운 온도","cloudType":"storm"},{"text":"공기처럼 당연하고 평온한 안식처","cloudType":"shade"}]'::jsonb, 'ko', 46),
(12, '당신에게 완벽한 데이트의 마침표는?', '[{"text":"정성껏 준비한 이벤트와 특별한 외식","cloudType":"sunlit"},{"text":"단둘이 와인 한 잔하며 나누는 깊은 대화","cloudType":"mist"},{"text":"낯선 곳으로 떠나 만나는 예상치 못한 모험","cloudType":"wild"},{"text":"익숙한 소파에 나란히 누워 보내는 고요","cloudType":"dawn"}]'::jsonb, 'ko', 47),
(12, '결국 사랑에서 가장 본질적인 것은?', '[{"text":"함께 나아갈 방향과 흔들리지 않는 약속","cloudType":"sunlit"},{"text":"영혼의 미세한 떨림까지 맞닿는 이해","cloudType":"mist"},{"text":"심장을 뛰게 만드는 거부할 수 없는 이끌림","cloudType":"storm"},{"text":"어떤 폭풍우 속에서도 나를 지켜줄 평온","cloudType":"shade"}]'::jsonb, 'ko', 48);
