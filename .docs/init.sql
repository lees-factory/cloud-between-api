CREATE TABLE cloud_between.test_questions (
                                              id SERIAL PRIMARY KEY,
                                              question_text TEXT NOT NULL,
                                              options JSONB NOT NULL,
                                              locale VARCHAR(5) DEFAULT 'ko',
                                              order_index INT,
                                              created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.test_questions IS '심리 테스트 질문 및 선택지 관리 테이블';
COMMENT ON COLUMN cloud_between.test_questions.options IS '선택지 배열: [{text: "...", cloudType: "..."}] 형태';
COMMENT ON COLUMN cloud_between.test_questions.locale IS '다국어 코드 (ko, en 등)';
COMMENT ON COLUMN cloud_between.test_questions.order_index IS '질문 노출 순서';

CREATE TABLE cloud_between.cloud_profiles (
                                              type_key VARCHAR(20),
                                              locale VARCHAR(5) DEFAULT 'ko',
                                              emoji VARCHAR(10),
                                              name VARCHAR(50) NOT NULL,
                                              subtitle VARCHAR(100),
                                              keywords TEXT[],
                                              lore TEXT,
                                              strengths TEXT[],
                                              shadows TEXT[],
                                              PRIMARY KEY (type_key, locale) -- 타입과 언어의 조합을 복합키로 설정
);

COMMENT ON TABLE cloud_between.cloud_profiles IS '각 구름 유형별 상세 설명 및 특징 정보';
COMMENT ON COLUMN cloud_between.cloud_profiles.type_key IS '구름 유형 식별자 (sunlit, mist 등)';
COMMENT ON COLUMN cloud_between.cloud_profiles.keywords IS '구름의 특징 키워드 배열';
COMMENT ON COLUMN cloud_between.cloud_profiles.lore IS '구름에 담긴 이야기/설명';


CREATE TABLE cloud_between.chemistry_matrix (
                                                id SERIAL PRIMARY KEY,
                                                cloud_type_1 VARCHAR(20) NOT NULL,
                                                cloud_type_2 VARCHAR(20) NOT NULL,
                                                sky_name VARCHAR(100),
                                                phenomenon VARCHAR(50),
                                                narrative TEXT,
                                                warning TEXT,
                                                UNIQUE(cloud_type_1, cloud_type_2)
);

COMMENT ON TABLE cloud_between.chemistry_matrix IS '두 구름 유형 간의 궁합 정보 매트릭스';
COMMENT ON COLUMN cloud_between.chemistry_matrix.phenomenon IS '함께 있을 때 나타나는 기상 현상 (glow, rain 등)';

/*
CREATE TABLE cloud_between.user_profiles (
                                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                             google_id TEXT UNIQUE NOT NULL,
                                             email TEXT NOT NULL,
                                             profile_image_url TEXT,
                                             is_paid BOOLEAN DEFAULT FALSE,
                                             last_login_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                                             created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
*/

CREATE TABLE cloud_between.user_profiles (
                                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                             google_id TEXT UNIQUE, -- Nullable for standard signup
                                             email TEXT UNIQUE NOT NULL,
                                             password_hash TEXT,    -- Nullable for Google login
                                             profile_image_url TEXT,
                                             is_paid BOOLEAN DEFAULT FALSE,
                                             last_login_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                                             created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMENT ON TABLE cloud_between.user_profiles IS '사용자 기본 정보 및 결제 상태 관리 테이블';
COMMENT ON COLUMN cloud_between.user_profiles.google_id IS 'Google OAuth2에서 발급받은 고유 ID (Sub)';
COMMENT ON COLUMN cloud_between.user_profiles.email IS '사용자 이메일 (로그인 식별자)';
COMMENT ON COLUMN cloud_between.user_profiles.password_hash IS '암호화된 비밀번호';
COMMENT ON COLUMN cloud_between.user_profiles.is_paid IS '유료 컨텐츠(상세 결과) 접근 권한 여부';