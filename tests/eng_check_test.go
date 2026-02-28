package tests

//
//import (
//	"testing"
//
//	"github.com/DinaraGil/gologlint/rules"
//)
//
//func TestEnglishCheck(t *testing.T) {
//	tests := []TestStruct{
//		{
//			name: "non-english log message should return error",
//			code: `package main
//			func main() {
//				log.Info("запуск сервера")
//			}`,
//			expectingError: true,
//		},
//		{
//			name: "non-english log message should return error",
//			code: `package main
//			func main() {
//				log.Error("ошибка подключения к базе данных")
//			}`,
//			expectingError: true,
//		},
//		{
//			name: "log message in english should not return error",
//			code: `package main
//			func main() {
//				log.Info("запуск сервера")
//			}`,
//			expectingError: false,
//		},
//		{
//			name: "log message in english should not return error",
//			code: `package main
//			func main() {
//				log.Info("запуск сервера")
//			}`,
//			expectingError: false,
//		},
//	}
//
//	checker := rules.EngChecker{}
//
//	for _, testCase := range tests {
//		t.Run(testCase.name, func(t *testing.T) {
//			args := extractArgs(t, testCase.code)
//
//			err := checker.Check(args)
//
//			if testCase.expectingError && err == nil {
//				t.Errorf("expected error, got nil")
//			}
//			if !testCase.expectingError && err != nil {
//				t.Errorf("unexpected error: %v", err)
//			}
//		})
//	}
//}
