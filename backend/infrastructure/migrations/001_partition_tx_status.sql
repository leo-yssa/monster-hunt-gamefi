-- tx_status 테이블 기본 생성 (파티셔닝은 002_optimize_indexes.sql에서 처리)

-- 1. 기존 테이블이 있다면 리네임(백업)
ALTER TABLE IF EXISTS tx_status RENAME TO tx_status_backup;

-- 2. 기본 테이블 생성 (파티셔닝 없이)
CREATE TABLE tx_status (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(66) UNIQUE NOT NULL,
    action VARCHAR(32) NOT NULL,
    params JSONB,
    user_id VARCHAR(64),
    status VARCHAR(16) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. 기본 인덱스 생성
CREATE INDEX IF NOT EXISTS idx_tx_status_created_at ON tx_status (created_at);
CREATE INDEX IF NOT EXISTS idx_tx_status_status ON tx_status (status);
CREATE INDEX IF NOT EXISTS idx_tx_status_user_id ON tx_status (user_id);
