package application

import (
	"context"
	"errors"
	"testing"
)

// MockGameRepository 테스트용 모의 리포지토리
type MockGameRepository struct {
	registerPlayerFunc func(ctx context.Context, name string) (string, error)
	addMonsterFunc     func(ctx context.Context, name string, hp, reward int) (string, error)
	huntMonsterFunc    func(ctx context.Context, monsterID int64) (string, error)
}

func (m *MockGameRepository) RegisterPlayer(ctx context.Context, name string) (string, error) {
	if m.registerPlayerFunc != nil {
		return m.registerPlayerFunc(ctx, name)
	}
	return "", errors.New("not implemented")
}

func (m *MockGameRepository) AddMonster(ctx context.Context, name string, hp, reward int) (string, error) {
	if m.addMonsterFunc != nil {
		return m.addMonsterFunc(ctx, name, hp, reward)
	}
	return "", errors.New("not implemented")
}

func (m *MockGameRepository) HuntMonster(ctx context.Context, monsterID int64) (string, error) {
	if m.huntMonsterFunc != nil {
		return m.huntMonsterFunc(ctx, monsterID)
	}
	return "", errors.New("not implemented")
}

func TestGameService_RegisterPlayer(t *testing.T) {
	tests := []struct {
		name          string
		playerName    string
		mockFunc      func(ctx context.Context, name string) (string, error)
		expectedError bool
	}{
		{
			name:       "성공적인 플레이어 등록",
			playerName: "TestPlayer",
			mockFunc: func(ctx context.Context, name string) (string, error) {
				return "0x123456789", nil
			},
			expectedError: false,
		},
		{
			name:       "빈 이름으로 등록 실패",
			playerName: "",
			mockFunc: func(ctx context.Context, name string) (string, error) {
				return "", errors.New("name cannot be empty")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockGameRepository{
				registerPlayerFunc: tt.mockFunc,
			}
			
			service := NewGameService(mockRepo)
			ctx := context.Background()
			
			result, err := service.RegisterPlayer(ctx, tt.playerName)
			
			if tt.expectedError {
				if err == nil {
					t.Errorf("예상된 에러가 발생하지 않았습니다")
				}
			} else {
				if err != nil {
					t.Errorf("예상치 못한 에러: %v", err)
				}
				if result == "" {
					t.Errorf("예상된 결과가 비어있습니다")
				}
			}
		})
	}
}

func TestGameService_AddMonster(t *testing.T) {
	tests := []struct {
		name          string
		monsterName   string
		hp            int
		reward        int
		mockFunc      func(ctx context.Context, name string, hp, reward int) (string, error)
		expectedError bool
	}{
		{
			name:        "성공적인 몬스터 추가",
			monsterName: "TestMonster",
			hp:          100,
			reward:      50,
			mockFunc: func(ctx context.Context, name string, hp, reward int) (string, error) {
				return "0xmonsterhash", nil
			},
			expectedError: false,
		},
		{
			name:        "잘못된 HP로 추가 실패",
			monsterName: "TestMonster",
			hp:          -10,
			reward:      50,
			mockFunc: func(ctx context.Context, name string, hp, reward int) (string, error) {
				return "", errors.New("invalid HP")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockGameRepository{
				addMonsterFunc: tt.mockFunc,
			}
			
			service := NewGameService(mockRepo)
			ctx := context.Background()
			
			result, err := service.AddMonster(ctx, tt.monsterName, tt.hp, tt.reward)
			
			if tt.expectedError {
				if err == nil {
					t.Errorf("예상된 에러가 발생하지 않았습니다")
				}
			} else {
				if err != nil {
					t.Errorf("예상치 못한 에러: %v", err)
				}
				if result == "" {
					t.Errorf("예상된 결과가 비어있습니다")
				}
			}
		})
	}
}

func TestGameService_HuntMonster(t *testing.T) {
	tests := []struct {
		name          string
		monsterID     int64
		mockFunc      func(ctx context.Context, monsterID int64) (string, error)
		expectedError bool
	}{
		{
			name:      "성공적인 사냥",
			monsterID: 1,
			mockFunc: func(ctx context.Context, monsterID int64) (string, error) {
				return "0xabcdef123", nil
			},
			expectedError: false,
		},
		{
			name:      "존재하지 않는 몬스터 사냥 실패",
			monsterID: 999,
			mockFunc: func(ctx context.Context, monsterID int64) (string, error) {
				return "", errors.New("monster not found")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockGameRepository{
				huntMonsterFunc: tt.mockFunc,
			}
			
			service := NewGameService(mockRepo)
			ctx := context.Background()
			
			result, err := service.HuntMonster(ctx, tt.monsterID)
			
			if tt.expectedError {
				if err == nil {
					t.Errorf("예상된 에러가 발생하지 않았습니다")
				}
			} else {
				if err != nil {
					t.Errorf("예상치 못한 에러: %v", err)
				}
				if result == "" {
					t.Errorf("예상된 결과가 비어있습니다")
				}
			}
		})
	}
} 