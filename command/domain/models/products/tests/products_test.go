package products_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"commandservice/domain/models/products"
)

// 商品エンティティ、値オブジェクトテスト用エントリーポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)                      // テストが失敗した場合のハンドラを登録する
	RunSpecs(t, "domain/models/productsパッケージのテスト") // 登録されたすべてのテストを実行する
}

// ProductIdのオブジェクト生成のベンチマーク
func BenchmarkNewProductId(b *testing.B) {
	// テスト用のUUIDを生成
	const UUID string = "12345678-1234-5678-1234-567812345678"
	// ベンチマークの実行
	for i := 0; i < b.N; i++ {
		_, _ = products.NewProductId(UUID)
	}
}
