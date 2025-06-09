package main

import (
	"ai_trainer/internal/infrastructure"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// ログにファイル名と行番号を含める
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// .env をロード
	if err := godotenv.Load(); err != nil {
		log.Println(".env ファイルが読み込めませんでした: ", err)
	}

	// 環境変数が設定されているか確認
	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("OPENAI_API_KEY が設定されていません")
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	infrastructure.RegisterGRPCServices(server)

	reflection.Register(server)

	log.Println("Starting gRPC server on :8080")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
