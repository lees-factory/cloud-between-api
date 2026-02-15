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
COMMENT ON COLUMN cloud_between.test_questions.options IS '선택지 배열: [{"text": "...", "personaType": "..."}]';
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
-- 테스트 메타데이터 INSERT
-- =============================================

-- 스텝 데이터 (12개)
INSERT INTO cloud_between.test_steps (id, title, emoji, order_index, locale) VALUES
(1,  '사랑의 시작',    '💕', 1,  'ko'),
(2,  '감정 표현',      '💬', 2,  'ko'),
(3,  '갈등 대처',      '⚡', 3,  'ko'),
(4,  '일상 속 사랑',   '☕', 4,  'ko'),
(5,  '자유와 공간',    '🕊️', 5,  'ko'),
(6,  '미래와 계획',    '🔮', 6,  'ko'),
(7,  '질투와 소유욕',  '👀', 7,  'ko'),
(8,  '친밀감',         '💫', 8,  'ko'),
(9,  '위기 대응',      '🌊', 9,  'ko'),
(10, '사랑의 언어',    '💝', 10, 'ko'),
(11, '혼자 vs 함께',   '🎭', 11, 'ko'),
(12, '사랑의 온도',    '🌡️', 12, 'ko');

-- 질문 데이터 (48개 = 12스텝 x 4질문)

-- Step 1: 사랑의 시작
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(1, '좋아하는 사람이 생겼을 때, 당신은?',
 '[{"text":"먼저 다가가서 대화를 시작한다","personaType":"sunlit"},{"text":"상대가 먼저 오기를 기다리며 관찰한다","personaType":"mist"},{"text":"직접적으로 호감을 표현한다","personaType":"storm"},{"text":"자연스럽게 친구처럼 지낸다","personaType":"dawn"}]'::jsonb,
 'ko', 1),
(1, '처음 만난 사람에게 당신은?',
 '[{"text":"활발하게 이야기를 이끈다","personaType":"sunlit"},{"text":"듣는 편이지만 공감을 잘한다","personaType":"mist"},{"text":"강한 인상을 남긴다","personaType":"storm"},{"text":"편안한 분위기를 만든다","personaType":"dawn"}]'::jsonb,
 'ko', 2),
(1, '연애에서 당신이 가장 중요하게 생각하는 것은?',
 '[{"text":"신뢰와 미래 계획","personaType":"sunlit"},{"text":"감정적 연결과 이해","personaType":"mist"},{"text":"열정과 케미","personaType":"storm"},{"text":"편안함과 안정","personaType":"shade"}]'::jsonb,
 'ko', 3),
(1, '사랑에 빠지는 속도는?',
 '[{"text":"천천히, 확신이 들면","personaType":"sunlit"},{"text":"시간을 두고 깊어진다","personaType":"mist"},{"text":"빠르게, 강렬하게","personaType":"storm"},{"text":"자연스럽게, 알아차리지 못할 정도로","personaType":"dawn"}]'::jsonb,
 'ko', 4);

-- Step 2: 감정 표현
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(2, '사랑한다는 말을 언제 하나요?',
 '[{"text":"확신이 들면 먼저 말한다","personaType":"sunlit"},{"text":"상대가 먼저 하면 나도 한다","personaType":"mist"},{"text":"느낌이 오면 바로 한다","personaType":"storm"},{"text":"말보다 행동으로 보여준다","personaType":"shade"}]'::jsonb,
 'ko', 5),
(2, '기분이 안 좋을 때 당신은?',
 '[{"text":"이야기하고 해결책을 찾는다","personaType":"sunlit"},{"text":"혼자 있고 싶어진다","personaType":"mist"},{"text":"감정을 숨기지 않는다","personaType":"storm"},{"text":"괜찮은 척 넘어간다","personaType":"dawn"}]'::jsonb,
 'ko', 6),
(2, '연인에게 애정을 표현하는 방식은?',
 '[{"text":"\"사랑해\", \"우리\" 같은 말을 자주 한다","personaType":"sunlit"},{"text":"작은 선물이나 메시지로","personaType":"mist"},{"text":"스킨십과 강한 표현","personaType":"storm"},{"text":"함께 있어주는 것 자체로","personaType":"shade"}]'::jsonb,
 'ko', 7),
(2, '상대방이 힘들어할 때 당신은?',
 '[{"text":"조언과 방향을 제시한다","personaType":"sunlit"},{"text":"공감하고 들어준다","personaType":"mist"},{"text":"함께 화내주거나 위로한다","personaType":"storm"},{"text":"조용히 곁에 있어준다","personaType":"shade"}]'::jsonb,
 'ko', 8);

-- Step 3: 갈등 대처
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(3, '연인과 싸웠을 때 당신은?',
 '[{"text":"대화로 해결하려고 한다","personaType":"sunlit"},{"text":"시간을 두고 생각한다","personaType":"mist"},{"text":"즉시 감정을 표출한다","personaType":"storm"},{"text":"웬만하면 피한다","personaType":"dawn"}]'::jsonb,
 'ko', 9),
(3, '갈등이 생기면 당신의 첫 반응은?',
 '[{"text":"\"우리 얘기 좀 하자\"","personaType":"sunlit"},{"text":"혼자 반복해서 생각한다","personaType":"mist"},{"text":"바로 따진다","personaType":"storm"},{"text":"\"괜찮아\"라고 넘긴다","personaType":"dawn"}]'::jsonb,
 'ko', 10),
(3, '화해는 어떻게 하나요?',
 '[{"text":"내가 먼저 화해를 제안한다","personaType":"sunlit"},{"text":"상대가 먼저 오면 받아준다","personaType":"mist"},{"text":"싸운 만큼 빠르게 화해한다","personaType":"storm"},{"text":"시간이 지나면 자연스럽게","personaType":"dawn"}]'::jsonb,
 'ko', 11),
(3, '상대가 잘못했을 때 당신은?',
 '[{"text":"명확하게 지적한다","personaType":"sunlit"},{"text":"상처받지만 말하지 않는다","personaType":"mist"},{"text":"즉시 표현한다","personaType":"storm"},{"text":"넘어가려고 노력한다","personaType":"dawn"}]'::jsonb,
 'ko', 12);

-- Step 4: 일상 속 사랑
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(4, '주말에 연인과 함께 할 때 당신은?',
 '[{"text":"계획을 미리 세운다","personaType":"sunlit"},{"text":"집에서 조용히 있고 싶다","personaType":"mist"},{"text":"즉흥적으로 뭔가 한다","personaType":"wild"},{"text":"편하게 있는 게 좋다","personaType":"dawn"}]'::jsonb,
 'ko', 13),
(4, '연인의 작은 변화를 알아차리나요?',
 '[{"text":"중요한 건 놓치지 않는다","personaType":"sunlit"},{"text":"아주 작은 것도 다 느낀다","personaType":"mist"},{"text":"큰 변화만 알아챈다","personaType":"storm"},{"text":"말해주면 알아챈다","personaType":"shade"}]'::jsonb,
 'ko', 14),
(4, '연인과의 루틴이 생기면?',
 '[{"text":"좋다, 안정적이다","personaType":"sunlit"},{"text":"편하지만 가끔 답답하다","personaType":"mist"},{"text":"지루하다, 변화가 필요하다","personaType":"wild"},{"text":"아주 좋다","personaType":"dawn"}]'::jsonb,
 'ko', 15),
(4, '기념일에 대한 당신의 생각은?',
 '[{"text":"중요하다, 챙긴다","personaType":"sunlit"},{"text":"의미 있게 보내고 싶다","personaType":"mist"},{"text":"특별하게 만들고 싶다","personaType":"storm"},{"text":"함께 있으면 된다","personaType":"shade"}]'::jsonb,
 'ko', 16);

-- Step 5: 자유와 공간
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(5, '연애할 때 당신에게 필요한 것은?',
 '[{"text":"명확한 관계 정의","personaType":"sunlit"},{"text":"감정적 안정감","personaType":"mist"},{"text":"설렘과 자극","personaType":"storm"},{"text":"개인 시간","personaType":"wild"}]'::jsonb,
 'ko', 17),
(5, '연인이 혼자 시간을 원하면?',
 '[{"text":"이유를 묻는다","personaType":"sunlit"},{"text":"나도 혼자 있고 싶어진다","personaType":"mist"},{"text":"섭섭하다","personaType":"storm"},{"text":"당연하다, 존중한다","personaType":"wild"}]'::jsonb,
 'ko', 18),
(5, '갑자기 여행 가자고 하면?',
 '[{"text":"일정 확인 후 계획한다","personaType":"sunlit"},{"text":"부담스럽다","personaType":"mist"},{"text":"좋아! 바로 간다","personaType":"wild"},{"text":"생각해본다","personaType":"dawn"}]'::jsonb,
 'ko', 19),
(5, '연인과 항상 붙어있는 것에 대해?',
 '[{"text":"좋다, 함께가 좋다","personaType":"sunlit"},{"text":"가끔은 숨 막힌다","personaType":"mist"},{"text":"상황에 따라 다르다","personaType":"storm"},{"text":"각자 시간도 필요하다","personaType":"wild"}]'::jsonb,
 'ko', 20);

-- Step 6: 미래와 계획
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(6, '연애의 미래에 대해 얼마나 생각하나요?',
 '[{"text":"자주, 구체적으로","personaType":"sunlit"},{"text":"막연하게","personaType":"mist"},{"text":"지금이 중요하다","personaType":"storm"},{"text":"가끔","personaType":"shade"}]'::jsonb,
 'ko', 21),
(6, '"우리 어디까지 갈 것 같아?" 라는 질문에?',
 '[{"text":"구체적인 그림이 있다","personaType":"sunlit"},{"text":"잘 모르겠다","personaType":"mist"},{"text":"지금 행복하면 됐다","personaType":"storm"},{"text":"천천히 보자","personaType":"dawn"}]'::jsonb,
 'ko', 22),
(6, '결혼에 대해 이야기하는 것은?',
 '[{"text":"중요하다, 명확해야 한다","personaType":"sunlit"},{"text":"조심스럽다","personaType":"mist"},{"text":"너무 이르다","personaType":"wild"},{"text":"때가 되면","personaType":"shade"}]'::jsonb,
 'ko', 23),
(6, '장거리 연애를 할 수 있나요?',
 '[{"text":"계획이 있으면 가능하다","personaType":"sunlit"},{"text":"힘들 것 같다","personaType":"mist"},{"text":"감정이 식을 것 같다","personaType":"storm"},{"text":"신뢰하면 가능하다","personaType":"shade"}]'::jsonb,
 'ko', 24);

-- Step 7: 질투와 소유욕
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(7, '연인이 이성 친구와 만나면?',
 '[{"text":"괜찮지만 알고 싶다","personaType":"sunlit"},{"text":"불안하다","personaType":"mist"},{"text":"질투난다","personaType":"storm"},{"text":"신경 안 쓴다","personaType":"wild"}]'::jsonb,
 'ko', 25),
(7, '연인의 과거 연애에 대해?',
 '[{"text":"궁금하지만 묻지 않는다","personaType":"sunlit"},{"text":"알고 싶지 않다","personaType":"mist"},{"text":"궁금하다","personaType":"storm"},{"text":"과거는 과거다","personaType":"dawn"}]'::jsonb,
 'ko', 26),
(7, '연인이 나를 소개하지 않으면?',
 '[{"text":"이유를 물어본다","personaType":"sunlit"},{"text":"상처받는다","personaType":"mist"},{"text":"화난다","personaType":"storm"},{"text":"이해한다","personaType":"wild"}]'::jsonb,
 'ko', 27),
(7, '당신의 소유욕은?',
 '[{"text":"적당히 있다","personaType":"sunlit"},{"text":"많은 편이다","personaType":"mist"},{"text":"강하다","personaType":"storm"},{"text":"거의 없다","personaType":"wild"}]'::jsonb,
 'ko', 28);

-- Step 8: 친밀감
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(8, '스킨십에 대한 당신의 생각은?',
 '[{"text":"중요하다","personaType":"sunlit"},{"text":"편할 때만","personaType":"mist"},{"text":"매우 중요하다","personaType":"storm"},{"text":"있어도 되고 없어도 된다","personaType":"shade"}]'::jsonb,
 'ko', 29),
(8, '연인과 깊은 이야기를 나누는 것은?',
 '[{"text":"자주, 관계를 위해 필요하다","personaType":"sunlit"},{"text":"하고 싶지만 어렵다","personaType":"mist"},{"text":"감정이 격해질 때","personaType":"storm"},{"text":"가끔, 필요할 때","personaType":"shade"}]'::jsonb,
 'ko', 30),
(8, '연인 앞에서 당신은?',
 '[{"text":"나 자신이다","personaType":"sunlit"},{"text":"조심스럽다","personaType":"mist"},{"text":"더 솔직해진다","personaType":"storm"},{"text":"편안하다","personaType":"dawn"}]'::jsonb,
 'ko', 31),
(8, '잠들기 전 연락은?',
 '[{"text":"매일 하고 싶다","personaType":"sunlit"},{"text":"하면 좋지만 꼭은 아니다","personaType":"mist"},{"text":"당연하다","personaType":"storm"},{"text":"바쁘면 안 해도 된다","personaType":"dawn"}]'::jsonb,
 'ko', 32);

-- Step 9: 위기 대응
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(9, '관계가 흔들릴 때 당신은?',
 '[{"text":"적극적으로 해결한다","personaType":"sunlit"},{"text":"혼자 고민한다","personaType":"mist"},{"text":"감정적으로 반응한다","personaType":"storm"},{"text":"시간을 둔다","personaType":"dawn"}]'::jsonb,
 'ko', 33),
(9, '이별 위기가 오면?',
 '[{"text":"끝까지 노력한다","personaType":"sunlit"},{"text":"상처받고 물러선다","personaType":"mist"},{"text":"강하게 잡거나 빠르게 떠난다","personaType":"storm"},{"text":"담담하게 받아들인다","personaType":"shade"}]'::jsonb,
 'ko', 34),
(9, '연인이 변했다고 느끼면?',
 '[{"text":"대화를 시도한다","personaType":"sunlit"},{"text":"눈치만 본다","personaType":"mist"},{"text":"직접 묻는다","personaType":"storm"},{"text":"지켜본다","personaType":"shade"}]'::jsonb,
 'ko', 35),
(9, '신뢰가 깨지면?',
 '[{"text":"회복을 시도한다","personaType":"sunlit"},{"text":"깊이 상처받는다","personaType":"mist"},{"text":"끝이다","personaType":"storm"},{"text":"시간이 필요하다","personaType":"shade"}]'::jsonb,
 'ko', 36);

-- Step 10: 사랑의 언어
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(10, '사랑받는다고 느끼는 순간은?',
 '[{"text":"미래를 함께 그릴 때","personaType":"sunlit"},{"text":"나를 이해해줄 때","personaType":"mist"},{"text":"열정적으로 대할 때","personaType":"storm"},{"text":"곁에 있어줄 때","personaType":"shade"}]'::jsonb,
 'ko', 37),
(10, '당신이 사랑을 표현하는 방식은?',
 '[{"text":"말과 계획","personaType":"sunlit"},{"text":"공감과 배려","personaType":"mist"},{"text":"행동과 감정","personaType":"storm"},{"text":"존재와 신뢰","personaType":"shade"}]'::jsonb,
 'ko', 38),
(10, '선물을 받는 것에 대해?',
 '[{"text":"의미가 중요하다","personaType":"sunlit"},{"text":"마음이 느껴지면 좋다","personaType":"mist"},{"text":"서프라이즈가 좋다","personaType":"storm"},{"text":"부담스럽다","personaType":"dawn"}]'::jsonb,
 'ko', 39),
(10, '연인에게 가장 해주고 싶은 말은?',
 '[{"text":"\"우리 미래를 함께 만들자\"","personaType":"sunlit"},{"text":"\"나는 네가 이해해\"","personaType":"mist"},{"text":"\"너 없인 못 살아\"","personaType":"storm"},{"text":"\"내가 여기 있어\"","personaType":"shade"}]'::jsonb,
 'ko', 40);

-- Step 11: 혼자 vs 함께
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(11, '주말에 혼자 있고 싶다면?',
 '[{"text":"계획된 일이면","personaType":"sunlit"},{"text":"자주 그렇다","personaType":"mist"},{"text":"거의 없다","personaType":"storm"},{"text":"필요할 때마다","personaType":"wild"}]'::jsonb,
 'ko', 41),
(11, '연인과 취미를 공유하는 것은?',
 '[{"text":"좋다, 함께 할 수 있다","personaType":"sunlit"},{"text":"부담스럽다","personaType":"mist"},{"text":"재미있다","personaType":"storm"},{"text":"각자 해도 된다","personaType":"wild"}]'::jsonb,
 'ko', 42),
(11, '항상 연락이 닿아야 하나요?',
 '[{"text":"어느 정도는","personaType":"sunlit"},{"text":"부담스럽다","personaType":"mist"},{"text":"당연하다","personaType":"storm"},{"text":"아니다","personaType":"wild"}]'::jsonb,
 'ko', 43),
(11, '연인의 모든 걸 알고 싶나요?',
 '[{"text":"중요한 건 알고 싶다","personaType":"sunlit"},{"text":"알면 부담된다","personaType":"mist"},{"text":"다 알고 싶다","personaType":"storm"},{"text":"말해주면 듣는다","personaType":"shade"}]'::jsonb,
 'ko', 44);

-- Step 12: 사랑의 온도
INSERT INTO cloud_between.test_questions (step_id, question_text, options, locale, order_index) VALUES
(12, '당신의 사랑은?',
 '[{"text":"따뜻하고 안정적","personaType":"sunlit"},{"text":"깊고 섬세함","personaType":"mist"},{"text":"뜨겁고 강렬함","personaType":"storm"},{"text":"고요하고 단단함","personaType":"shade"}]'::jsonb,
 'ko', 45),
(12, '오래 사귄 연인과는?',
 '[{"text":"더 든든하다","personaType":"sunlit"},{"text":"더 편하다","personaType":"mist"},{"text":"가끔 지루하다","personaType":"wild"},{"text":"더 깊어진다","personaType":"shade"}]'::jsonb,
 'ko', 46),
(12, '완벽한 데이트는?',
 '[{"text":"계획된 특별한 하루","personaType":"sunlit"},{"text":"둘만의 조용한 시간","personaType":"mist"},{"text":"예상 못한 모험","personaType":"wild"},{"text":"함께 있는 평범한 시간","personaType":"dawn"}]'::jsonb,
 'ko', 47),
(12, '사랑에서 가장 중요한 건?',
 '[{"text":"신뢰와 방향성","personaType":"sunlit"},{"text":"이해와 공감","personaType":"mist"},{"text":"열정과 끌림","personaType":"storm"},{"text":"편안함과 지속성","personaType":"shade"}]'::jsonb,
 'ko', 48);
