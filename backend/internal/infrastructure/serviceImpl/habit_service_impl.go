package serviceImpl

// serviceImpl規約
// ・エラーはhandlerに返すのみ。handler側でログ出力する。
// ・複数のrepositoryメソッドもしくはデータを変更するrepositoryメソッドを使用する場合はトランザクションを実行する
// ・repositoryのメソッドに渡すcontext.Contextについて
//   -> トランザクションが不要な場合はhandlerから受け取ったctxをそのまま渡す
//   -> トランザクションを実行する場合はmongo.SessionContextを渡す（mongo.SessionContextはcontext.Contextインターフェースを満たしている）

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"backend/internal/domain/common"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/model/habit"
	"backend/internal/domain/repository"
)

type habitService struct {
	client         *mongo.Client
	habitRepo      repository.HabitRepository
	dailyTrackRepo repository.DailyTrackRepository
}

func NewHabitService(client *mongo.Client, habitRepo repository.HabitRepository, dailyTrackRepo repository.DailyTrackRepository) *habitService {
	return &habitService{
		client:         client,
		habitRepo:      habitRepo,
		dailyTrackRepo: dailyTrackRepo,
	}
}

func (s *habitService) GetHabitList(ctx context.Context, userId string) ([]*habit.Habit, error) {
	habits, err := s.habitRepo.FetchAll(ctx, userId)

	if err != nil {
		return nil, err
	}

	if habits == nil {
		habits = make([]*habit.Habit, 0)
	}

	return habits, nil
}

func (s *habitService) RegisterHabit(ctx context.Context, userId string, habitName string) (*habit.Habit, error) {

	// セッションの開始
	session, err := s.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	// トランザクションの実行
	var resultHabit *habit.Habit
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {

		// 新規登録
		newHabit := habit.Habit{UserId: userId, Name: habitName}
		resultHabit, err = s.habitRepo.Register(sessionContext, &newHabit)

		if err != nil {
			return err
		}

		// 今日のdaily-trackを取得
		todayString := time.Now().Format(`2006-01-02`) // YYYY-MM-DD
		todaysTrack, err := s.dailyTrackRepo.FindDailyTrack(sessionContext, userId, todayString)
		if err != nil && !errors.Is(err, common.ErrNotFound) {
			return err
		}

		// 今日のdaily-trackがあれば作成した習慣を追加
		if !errors.Is(err, common.ErrNotFound) {

			newHabitStatus := &daily_track.HabitStatus{
				HabitId:   newHabit.Id,
				HabitName: newHabit.Name,
				IsDone:    false,
			}
			todaysTrack.HabitStatuses = append(todaysTrack.HabitStatuses, newHabitStatus)

			err = s.dailyTrackRepo.UpdateHabitStatuses(sessionContext, todaysTrack)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resultHabit, nil
}

func (s *habitService) DeleteHabit(ctx context.Context, userId string, habitId string) error {
	// セッションの開始
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	// トランザクションの実行
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {

		// 削除
		err := s.habitRepo.Delete(sessionContext, habitId)

		if err != nil {
			return err
		}

		// 今日のdaily-trackを取得
		todayString := time.Now().Format(`2006-01-02`) // YYYY-MM-DD
		todaysTrack, err := s.dailyTrackRepo.FindDailyTrack(sessionContext, userId, todayString)
		if err != nil && !errors.Is(err, common.ErrNotFound) {
			return err
		}

		if errors.Is(err, common.ErrNotFound) {
			return nil
		}

		// 今日のdaily-trackがあれば習慣を削除
		for index, habitStatus := range todaysTrack.HabitStatuses {
			if habitStatus.HabitId == habitId && !habitStatus.IsDone {
				// 削除対象の習慣が完了していなければ削除
				todaysTrack.HabitStatuses = append(todaysTrack.HabitStatuses[:index], todaysTrack.HabitStatuses[index+1:]...)

				// 永続化
				err = s.dailyTrackRepo.UpdateHabitStatuses(sessionContext, todaysTrack)
				if err != nil {
					return err
				}

				break
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
