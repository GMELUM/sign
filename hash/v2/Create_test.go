package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestCreate(t *testing.T) {
	// Определение тестовых случаев
	tests := []struct {
		name     string
		data     interface{}
		secret   string
		wantErr  bool
	}{
		{
			name:   "Simple flat data",
			data:   map[string]interface{}{"field1": "value1", "field2": 42},
			secret: "simple_secret_key",
		},
		{
			name:   "Nested map data",
			data:   map[string]interface{}{"field1": "value1", "nested": map[string]interface{}{"subfield": "subvalue"}},
			secret: "nested_secret_key",
		},
		{
			name:   "Array data",
			data:   map[string]interface{}{"field1": []int{1, 2, 3}, "field2": "array_data"},
			secret: "array_secret_key",
		},
		{
			name: "Deeply nested data",
			data: map[string]interface{}{
				"field1": "value1",
				"nested": map[string]interface{}{
					"subfield1": "subvalue1",
					"subfield2": map[string]interface{}{
						"subsubfield1": 123,
						"subsubfield2": []string{"a", "b", "c"},
					},
				},
			},
			secret: "deep_nested_secret_key",
		},
		{
			name:    "Empty data",
			data:    map[string]interface{}{},
			secret:  "empty_secret_key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызов функции Create
			result, err := Create(tt.data, tt.secret)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Create() ошибка = %v, ожидалось %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			// Десериализация результата
			var data Data
			if err := msgpack.Unmarshal(result, &data); err != nil {
				t.Fatalf("Ошибка при десериализации результата: %v", err)
			}

			// Проверка, что Data не пустое
			if len(data.Data) == 0 {
				t.Error("Поле Data не должно быть пустым")
			}

			// Проверка, что Hash не пустое
			if data.Hash == "" {
				t.Error("Поле Hash не должно быть пустым")
			}

			// Перепроверка корректности Hash
			h := hmac.New(sha256.New, []byte(tt.secret))
			h.Write(data.Data)
			expectedHash := hex.EncodeToString(h.Sum(nil))

			if data.Hash != expectedHash {
				t.Errorf("Некорректный хеш: ожидалось %s, получено %s", expectedHash, data.Hash)
			}
		})
	}
}

func BenchmarkCreate(b *testing.B) {
	// Определяем различные тестовые наборы данных
	tests := []struct {
		name   string
		data   interface{}
		secret string
	}{
		{
			name:   "Simple",
			data:   map[string]interface{}{"field1": "value1", "field2": 42},
			secret: "simple_secret_key",
		},
		{
			name:   "Nested",
			data:   map[string]interface{}{"field1": "value1", "nested": map[string]interface{}{"subfield": "subvalue"}},
			secret: "nested_secret_key",
		},
		{
			name:   "Array",
			data:   map[string]interface{}{"field1": []int{1, 2, 3}, "field2": "array_data"},
			secret: "array_secret_key",
		},
		{
			name: "DeeplyNested",
			data: map[string]interface{}{
				"field1": "value1",
				"nested": map[string]interface{}{
					"subfield1": "subvalue1",
					"subfield2": map[string]interface{}{
						"subsubfield1": 123,
						"subsubfield2": []string{"a", "b", "c"},
					},
				},
			},
			secret: "deeply_nested_secret_key",
		},
		{
			name: "Large",
			data: func() map[string]interface{} {
				largeData := make(map[string]interface{})
				for i := 0; i < 1000; i++ {
					largeData[fmt.Sprintf("field_%d", i)] = i
				}
				return largeData
			}(),
			secret: "large_data_secret_key",
		},
	}

	// Запускаем каждый тест как под-бенчмарк
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := Create(tt.data, tt.secret)
				if err != nil {
					b.Fatalf("Ошибка при выполнении Create: %v", err)
				}
			}
		})
	}
}