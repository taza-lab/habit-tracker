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

	"go.mongodb.org/mongo-driver/mongo"

	"backend/internal/config"
	"backend/internal/domain/common"
	"backend/internal/domain/model/daily_track"
	"backend/internal/domain/model/habit"
	userModel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
)

type dailyTrackService struct {
	client         *mongo.Client
	userRepo       repository.UserRepository
	habitRepo      repository.HabitRepository
	dailyTrackRepo repository.DailyTrackRepository
}

func NewDailyTrackService(
	client *mongo.Client,
	userRepo repository.UserRepository,
	habitRepo repository.HabitRepository,
	dailyTrackRepo repository.DailyTrackRepository,
) *dailyTrackService {
	return &dailyTrackService{
		client:         client,
		userRepo:       userRepo,
		habitRepo:      habitRepo,
		dailyTrackRepo: dailyTrackRepo,
	}
}

func (s *dailyTrackService) GetDailyTrack(ctx context.Context, userId string, targetDate string) (*daily_track.DailyTrack, error) {
	// セッションの開始
	session, err := s.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	// トランザクションの実行
	var todaysTrack *daily_track.DailyTrack
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		todaysTrack, err = s.dailyTrackRepo.FindDailyTrack(sessionContext, userId, targetDate)

		// 想定外のエラーはエラーとして返す
		if err != nil && !errors.Is(err, common.ErrNotFound) {
			return err
		}

		// 指定日のデータが存在していたら抜ける
		if todaysTrack != nil {
			return nil
		}

		// 指定日のデータが存在しなかったら作成

		// 習慣一覧を取得
		var habits []*habit.Habit
		habits, err = s.habitRepo.FetchAll(sessionContext, userId)
		if err != nil {
			return err
		}

		// 習慣一覧を使ってステータスの配列を作成
		var habitStatuses []*daily_track.HabitStatus
		for _, habit := range habits {
			habitStatus := &daily_track.HabitStatus{
				HabitId:   habit.Id,
				HabitName: habit.Name,
				IsDone:    false,
			}
			habitStatuses = append(habitStatuses, habitStatus)
		}

		newDailyTrack := daily_track.DailyTrack{
			UserId:        userId,
			Date:          targetDate,
			HabitStatuses: habitStatuses,
		}

		todaysTrack, err = s.dailyTrackRepo.RegisterDailyTrack(sessionContext, &newDailyTrack)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return todaysTrack, nil

}

func (s *dailyTrackService) UpdateDoneDailyTrack(ctx context.Context, userId string, targetDate string, targetHabitId string) error {
	// セッションの開始
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	// トランザクションの実行
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		// todaysTrack 取得
		todaysTrack, err := s.dailyTrackRepo.FindDailyTrack(sessionContext, userId, targetDate)
		if err != nil {
			return err
		}

		// todaysTrack.IsDone 更新
		for _, habitStatus := range todaysTrack.HabitStatuses {
			if habitStatus.HabitId == targetHabitId {
				habitStatus.IsDone = true
			}
		}

		// dailyTrack更新して保存
		err = s.dailyTrackRepo.UpdateHabitStatuses(sessionContext, todaysTrack)
		if err != nil {
			return err
		}

		// ユーザー情報取得
		var user *userModel.User
		user, err = s.userRepo.Find(sessionContext, userId)
		if err != nil {
			return err
		}

		// point 加算
		points := user.Points + config.PointsForHabitDone
		err = s.userRepo.UpdatePoints(sessionContext, userId, points)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
