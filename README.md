
📝 サーバー立ち上げ手順と動作確認手順
=======================

✅ 1. 依存ライブラリのインストール
go mod tidy

✅ 2. OpenAI APIキーの設定
export OPENAI_API_KEY="your-api-key"

✅ 3. サーバーの起動
cd backend/cmd/server
go run main.go
ログ: "Starting gRPC server on :8080" が出たら成功

✅ 4. 動作確認（例: curl またはフロントエンドから）

```
localhost:8080> package trainerpf.v1

trainerpf.v1@localhost:8080> service WorkoutService

trainerpf.v1.WorkoutService@localhost:8080> call EvaluateWorkout
age (TYPE_INT32) => 23
weight (TYPE_FLOAT) => 62
<repeated> workouts::name (TYPE_STRING) => ベンチプレス
<repeated> workouts::sets::load (TYPE_FLOAT) => 30
<repeated> workouts::sets::reps (TYPE_INT32) => 10
<repeated> workouts::sets::load (TYPE_FLOAT) => 
<repeated> workouts::name (TYPE_STRING) => 
{
  "advice": "評価: ベンチプレスで30 kgを10回行うのは良い成績です。継続してトレーニングを行うことでさらなる強化が見込めます。\n\nアドバイス: 継続してトレーニングを行い、徐々に重量を増やすことで筋力が向上します。適切な栄養摂取も大切です。"
}
```
