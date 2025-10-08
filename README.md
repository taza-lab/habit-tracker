# Habit Tracker


**毎日の"できた"を積み重ねる、シンプルな習慣化サポートアプリです。**



## 概要 📖

日々の生活の中で「これを習慣にしたい！」と思っても、つい忘れてしまったり、三日坊主で終わってしまったりすることはありませんか？

このアプリケーションは、そんな課題を解決するために開発しました。毎日同じタスクをTodoリストに手入力する手間をなくし、「今日のタスク」を自動で生成。ユーザーは達成した習慣にチェックを入れるだけで、簡単に日々の頑張りを記録・可視化できます。

実務経験のないモダンな技術スタック（Go, Next.js, AWS App Runnerなど）の学習と実践を目的として、ゼロから開発しました。



## デモ 🚀

以下のURLから実際にアプリケーションを触ることができます。

**URL : [https://d3oxqyisydxdow.cloudfront.net/login](https://d3oxqyisydxdow.cloudfront.net/login)** 



## 主な機能 ✨

* **習慣の登録・管理 (CRD):** 習慣にしたい項目を自由に登録・削除できます。
* **デイリーチェックリスト:** 登録した習慣が毎朝自動でリストアップされます。
* **ポイントシステム:** 習慣を達成してチェックを入れるたびにポイントが貯まります。日々の頑張りがスコアとして蓄積されていくので、モチベーションがアップします！



## こだわった点 🛠️

### 🥇 クリーンアーキテクチャによる依存性の分離

バックエンド（Go）に**クリーンアーキテクチャ**を採用しました。
業務や学習を通じてこの設計に触れてきましたが、ゼロから設計思想を適用し構築した経験がなかったため、 **「実践し、自分のものにしたい」** と思い、技術的な挑戦として採用しました。


ビジネスロジック、データベースといった各関心事をレイヤーごとに明確に分離することで、**特定のフレームワークやデータベースへの依存を最小限に**抑えています。
これにより、以下のメリットが生まれました。

* **高いテスト容易性:** ビジネスロジックがインフラストラクチャから独立しているため、純粋なロジックのみを簡単にユニットテストできます。　
* **柔軟なコンポーネント交換:** 将来的にGinから別のWebフレームワークに移行したり、MongoDBを他のデータベースに置き換える際も、ビジネスロジックへの影響を最小限に留めることができます。
* **保守性の向上:** コードの見通しが良くなり、機能追加や修正が容易になりました。



## 技術スタック 💻

![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)
![Amazon AWS](https://img.shields.io/badge/Amazon_AWS-232F3E?style=for-the-badge&logo=amazonaws&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

| カテゴリ         | 技術                                                                                             |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| **フロントエンド** | Next.js, TypeScript, React                                                                       |
| **バックエンド** | Go, Gin                                                                                          |
| **データベース** | MongoDB                                                                                          |
| **インフラ** | AWS (App Runner, ECR, S3, CloudFront), MongoDB Atlas, Docker                                     |
| **その他** | Gemini, ChatGPT |




## 環境構築 ⚙️

このリポジトリをローカル環境で動かすための手順です。

```bash
# 1. リポジトリをクローン
$ git clone https://github.com/taza-lab/habit-tracker.git

# 2. 環境変数を設定
# backend/に .env ファイルを作成し、以下の環境変数を設定してください。
APP_SECRET_KEY=app_secret_key
JWT_SECRET_KEY=jwt_secret_key
NEXT_BASE_URL="http://localhost:3000"
DATABASE_URI=mongodb://mongodb:27017
DATABASE_NAME=habit_tracker

# frontend/に .env.local ファイルを作成し、以下の環境変数を設定してください。
NEXT_PUBLIC_API_BASE_URL='http://localhost:8080'
NEXT_PUBLIC_HABIT_DONE_POINT=3

# 3. ビルド・起動
$ docker-compose up --build -d

```

http://localhost:3000/login にアクセス！



## 今後の展望 🗺️

### ✨ 新機能の追加

- **🎁 ポイント交換機能:** 貯まったポイントをアプリ内のテーマカラーやアイコンなどと交換できるようにします。
- **🏆 実績バッチ機能:** 習慣ごとの達成率や継続日数によって実績バッチを獲得できるようにし、モチベーションをサポートします。
- **🎲 ランダムタスク機能:** 登録している習慣とは別に「犬の写真を撮る」「顔の体操をする」などのユニークなタスクを毎日ランダムで生成します。チェック時にはボーナスポイントを獲得できます。


### 🛠️ 品質の向上と開発体験の改善

- **🎨 UI/UX改善:** ユーザーからのフィードバックを元に、より直感的で使いやすいデザインに刷新します。
- **🛡️ 入力値バリデーション強化:** フロントエンドとバックエンドの両方で堅牢なバリデーションを実装し、アプリケーションの安定性とセキュリティを向上させます。
- **🚀 CI/CDパイプライン構築:** GitHub Actionsなどを利用してテストとデプロイを自動化し、開発サイクルの高速化と品質担保を実現します。
- **🌐 独自ドメイン取得:** 現在のCloudFrontのエンドポイントから、覚えやすい独自のドメインへ移行します。
- **✅ Unitテスト・Featureテストの実装:** バックエンドでビジネスロジックに対するUnitテスト、APIのFeatureテストを実装し、CI/CDに組み込みます。
