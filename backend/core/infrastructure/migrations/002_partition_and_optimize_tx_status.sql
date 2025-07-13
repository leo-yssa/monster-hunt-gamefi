-- tx_status 테이블 파티셔닝, 고급 인덱스, 뷰, 파티션 관리 함수 등 성능 최적화

-- 1. 파티셔닝 설정 (기존 테이블을 파티셔닝으로 변경)
-- 기존 테이블 백업
CREATE TABLE IF NOT EXISTS tx_status_backup AS SELECT * FROM tx_status;

-- 기존 테이블 삭제
DROP TABLE IF EXISTS tx_status;

-- 파티셔닝 테이블 생성 (PRIMARY KEY에 created_at 포함)
CREATE TABLE tx_status (
    id SERIAL,
    tx_hash VARCHAR(66) NOT NULL,
    action VARCHAR(32) NOT NULL,
    params JSONB,
    user_id VARCHAR(64),
    status VARCHAR(16) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, created_at),
    UNIQUE (tx_hash, created_at)
) PARTITION BY RANGE (created_at);

-- 2. 파티션 생성 (오늘 기준으로 현재 월부터 3개월)
-- 현재 월: 2025년 7월
-- 다음 3개월: 2025년 7월, 8월, 9월
CREATE TABLE tx_status_2025_07 PARTITION OF tx_status
    FOR VALUES FROM ('2025-07-01') TO ('2025-08-01');
CREATE TABLE tx_status_2025_08 PARTITION OF tx_status
    FOR VALUES FROM ('2025-08-01') TO ('2025-09-01');
CREATE TABLE tx_status_2025_09 PARTITION OF tx_status
    FOR VALUES FROM ('2025-09-01') TO ('2025-10-01');

-- 3. 기존 데이터 복원 (백업이 있는 경우)
INSERT INTO tx_status SELECT * FROM tx_status_backup;
DROP TABLE IF EXISTS tx_status_backup;

-- 4. 추가 성능 최적화 인덱스
-- 상태별 빠른 조회를 위한 부분 인덱스
CREATE INDEX IF NOT EXISTS idx_tx_status_pending 
ON tx_status (created_at) 
WHERE status = 'pending';

CREATE INDEX IF NOT EXISTS idx_tx_status_success 
ON tx_status (created_at) 
WHERE status = 'success';

CREATE INDEX IF NOT EXISTS idx_tx_status_fail 
ON tx_status (created_at) 
WHERE status = 'fail';

-- 액션별 조회 최적화
CREATE INDEX IF NOT EXISTS idx_tx_status_action_created 
ON tx_status (action, created_at);

-- 사용자별 최근 트랜잭션 조회 최적화
CREATE INDEX IF NOT EXISTS idx_tx_status_user_created_desc 
ON tx_status (user_id, created_at DESC);

-- 5. 통계 정보 업데이트 (쿼리 플래너 최적화)
ANALYZE tx_status;

-- 6. 자동 파티션 생성 함수 (월별)
-- 이 함수는 미래 파티션을 자동으로 생성합니다
CREATE OR REPLACE FUNCTION create_tx_status_partition(partition_date DATE)
RETURNS VOID AS $$
DECLARE
    partition_name TEXT;
    start_date DATE;
    end_date DATE;
BEGIN
    partition_name := 'tx_status_' || to_char(partition_date, 'YYYY_MM');
    start_date := date_trunc('month', partition_date);
    end_date := start_date + interval '1 month';
    
    EXECUTE format('CREATE TABLE IF NOT EXISTS %I PARTITION OF tx_status
                    FOR VALUES FROM (%L) TO (%L)',
                    partition_name, start_date, end_date);
                    
    RAISE NOTICE 'Created partition: %', partition_name;
END;
$$ LANGUAGE plpgsql;

-- 7. 미래 3개월 파티션 자동 생성 (오늘 기준, date 캐스팅)
-- 현재: 2025년 7월 → 미래 3개월: 2025년 10월, 11월, 12월
SELECT create_tx_status_partition((CURRENT_DATE + INTERVAL '3 month')::date);  -- 2025-10
SELECT create_tx_status_partition((CURRENT_DATE + INTERVAL '4 month')::date);  -- 2025-11  
SELECT create_tx_status_partition((CURRENT_DATE + INTERVAL '5 month')::date);  -- 2025-12

-- 8. 성능 모니터링 뷰 생성
CREATE OR REPLACE VIEW tx_status_stats AS
SELECT 
    status,
    COUNT(*) as count,
    MIN(created_at) as oldest,
    MAX(created_at) as newest,
    AVG(EXTRACT(EPOCH FROM (updated_at - created_at))) as avg_processing_time_seconds
FROM tx_status 
GROUP BY status;

-- 9. 파티션별 통계 뷰
CREATE OR REPLACE VIEW partition_stats AS
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) as size,
    pg_total_relation_size(schemaname||'.'||tablename) as size_bytes
FROM pg_tables 
WHERE tablename LIKE 'tx_status_%'
ORDER BY size_bytes DESC; 