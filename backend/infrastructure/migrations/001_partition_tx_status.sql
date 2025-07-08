-- tx_status 테이블을 월별(created_at) RANGE 파티셔닝으로 변경

-- 1. 기존 테이블이 있다면 리네임(백업)
ALTER TABLE IF EXISTS tx_status RENAME TO tx_status_backup;

-- 2. 파티셔닝 테이블 생성
CREATE TABLE tx_status (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(66) UNIQUE NOT NULL,
    action VARCHAR(32) NOT NULL,
    params JSONB,
    user_id VARCHAR(64),
    status VARCHAR(16) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
) PARTITION BY RANGE (created_at);

-- 3. 2024년 7~9월 파티션 생성 예시
CREATE TABLE tx_status_2024_07 PARTITION OF tx_status
    FOR VALUES FROM ('2024-07-01') TO ('2024-08-01');
CREATE TABLE tx_status_2024_08 PARTITION OF tx_status
    FOR VALUES FROM ('2024-08-01') TO ('2024-09-01');
CREATE TABLE tx_status_2024_09 PARTITION OF tx_status
    FOR VALUES FROM ('2024-09-01') TO ('2024-10-01');
