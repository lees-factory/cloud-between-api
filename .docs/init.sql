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
-- 결제 관련 테이블
-- =============================================

CREATE TABLE cloud_between.payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES cloud_between.user_profiles(id),
    order_id VARCHAR(100) UNIQUE NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.payments IS '결제 정보 관리 테이블';
COMMENT ON COLUMN cloud_between.payments.order_id IS 'PayPal 주문 ID';
COMMENT ON COLUMN cloud_between.payments.status IS '결제 상태 (PENDING, COMPLETED)';

CREATE TABLE cloud_between.payment_cancels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id UUID REFERENCES cloud_between.payments(id),
    order_id VARCHAR(100) NOT NULL,
    reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.payment_cancels IS '결제 취소 기록 테이블';

