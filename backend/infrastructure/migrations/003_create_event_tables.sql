-- MonsterHuntedEvent 테이블 생성
CREATE TABLE IF NOT EXISTS monster_hunted_events (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(66) NOT NULL,
    player VARCHAR(42) NOT NULL,
    monster_id VARCHAR(64) NOT NULL,
    reward VARCHAR(64) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- PlayerRegisteredEvent 테이블 생성
CREATE TABLE IF NOT EXISTS player_registered_events (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(66) NOT NULL,
    player VARCHAR(42) NOT NULL,
    name VARCHAR(64) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
); 