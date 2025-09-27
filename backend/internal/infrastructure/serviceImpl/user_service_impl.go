package serviceImpl

// serviceImpl規約
// ・エラーはhandlerに返すのみ。handler側でログ出力する。
// ・複数のrepositoryメソッドもしくはデータを変更するrepositoryメソッドを使用する場合はトランザクションを実行する
// ・repositoryのメソッドに渡すcontext.Contextについて
//   -> トランザクションが不要な場合はhandlerから受け取ったctxをそのまま渡す
//   -> トランザクションを実行する場合はmongo.SessionContextを渡す（mongo.SessionContextはcontext.Contextインターフェースを満たしている）

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"backend/internal/config"
	"backend/internal/domain/common"
	userModel "backend/internal/domain/model/user"
	"backend/internal/domain/repository"
)

type userService struct {
	client   *mongo.Client
	userRepo repository.UserRepository
}

func NewUserService(client *mongo.Client, userRepo repository.UserRepository) *userService {
	return &userService{
		client:   client,
		userRepo: userRepo,
	}
}

func (s *userService) SignUp(ctx context.Context, userName string, password string) (*userModel.User, error) {
	// セッションの開始
	session, err := s.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	// トランザクションの実行
	var resultUser *userModel.User
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		// 同一usernameが登録済みかどうかのチェック
		_, err := s.userRepo.FindByUserName(sessionContext, userName)

		if err != nil && err != common.ErrNotFound {
			return err
		}
		if err == nil {
			return common.ErrAlreadyExists
		}

		// 登録
		user := userModel.User{Username: userName, Password: password, Points: 0}
		resultUser, err = s.userRepo.Register(sessionContext, &user)
		if err != nil {
			return err
		}
		resultUser.Password = ""

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resultUser, nil

}

func (s *userService) Login(ctx context.Context, userName string, password string) (*userModel.User, string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	var user *userModel.User
	var tokenString string

	// ユーザー取得
	user, err := s.userRepo.FindByUserName(ctx, userName)

	if err != nil {
		return nil, tokenString, err
	}

	// パスワードチェック
	// bcrypt.CompareHashAndPasswordが保存されているハッシュ値とユーザーが入力したパスワードが一致するかを検証
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, tokenString, common.ErrPasswordMismatch
	}
	user.Password = ""

	// JWTトークンの生成
	expirationTime := time.Now().Add(config.JwtExpirationHour * time.Hour)
	claims := &userModel.Claims{
		UserId:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtSecret)
	if err != nil {
		return nil, tokenString, err
	}

	return user, tokenString, nil
}
