package hash

import (
	"fmt"
	"testing"
)

func BenchmarkValidate(b *testing.B) {
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
			// Генерация данных с использованием функции Create
			data, err := Create(tt.data, tt.secret)
			if err != nil {
				b.Fatalf("Ошибка при генерации данных: %v", err)
			}

			// Запуск под-бенчмарка
			for i := 0; i < b.N; i++ {
				var result interface{}
				err := Validate(&result, data, tt.secret)
				if err != nil {
					b.Fatalf("Ошибка при выполнении Validate: %v", err)
				}
			}
		})
	}
}