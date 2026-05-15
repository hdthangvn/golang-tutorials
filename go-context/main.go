package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// // Naive | Edit | Test | Explain | Document
// func placeOrderWithoutContext(orderID string) error {
// 	log.Printf("Đặt đơn hàng %s với %s đơn hàng:", orderID)

// 	time.Sleep(3 * time.Second)

// 	log.Printf("Xử lý đơn hàng %s thành công (sau 3 giây)", orderID)
// 	return nil // Thành công
// }

// // Naive | Edit | Test | Explain | Document
// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	orderID := "GO-12345"

// 	err := placeOrderWithoutContext(orderID)

// 	if err != nil {
// 		http.Error(w, "Lỗi xử lý đơn hàng", http.StatusInternalServerError) // 500 Internal Server Error
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Đặt hàng thành công"))
// }

// // OrderHandlerSelect
// // Tabline | Edit | Test | Explain | Document
// func OrderHandlerSelect(w http.ResponseWriter, r *http.Request) {
// 	orderID := "GO-12345"
// 	resultChan := make(chan error, 1)

// 	// xử lý đơn hàng trong goroutine
// 	go func() {
// 		err := placeOrderWithoutContext(orderID)
// 		resultChan <- err
// 	}()

// 	select {
// 	case err := <-resultChan:
// 		if err != nil {
// 			log.Printf("Xử lý đơn hàng %s thất bại", orderID)
// 			http.Error(w, "Lỗi xử lý đơn hàng", http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Đặt hàng thành công"))

// 	case <-time.After(2 * time.Second):
// 		log.Printf("Xử lý đơn hàng %s quá 2 giây, trả lời cho client\n", orderID)
// 		http.Error(w, "Quá thời gian xử lý, vui lòng thử lại sau", http.StatusGatewayTimeout) // 504 Gateway Timeout
// 	}
// }

// OrderHandlerWithContext
// Tabline | Edit | Test | Explain | Document
func OrderHandlerWithContext(w http.ResponseWriter, r *http.Request) {
	orderID := "GO-12345"
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	err := placeOrderWithContext(ctx, orderID)

	if err != nil {
		log.Printf("Xử lý đơn hàng %s thất bại: %s\n", orderID, err)
		http.Error(w, "Lỗi xử lý đơn hàng hoặc quá thời gian", http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Đặt hàng thành công"))
}

// placeOrderWithContext
// Tabline | Edit | Test | Explain | Document
func placeOrderWithContext(ctx context.Context, orderID string) error {
	log.Printf("Xử lý đơn hàng %s bắt đầu\n", orderID)

	select {
	case <-time.After(3 * time.Second): // giả lập xử lý 3 giây
		log.Printf("Xử lý đơn hàng %s thành công\n", orderID)
		return nil
	case <-ctx.Done():
		log.Printf("Xử lý đơn hàng %s bị hủy\n", orderID)
		return ctx.Err()
	}
}

// Naive | Edit | Test | Explain | Document
func main() {
	// // http.HandleFunc("/order", OrderHandler)
	// // http.HandleFunc("/order", OrderHandlerSelect)
	// http.HandleFunc("/order", OrderHandlerWithContext)
	// log.Print("Server đang chạy tại http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// context
	ctx := context.Background() // root -> request http
	// Tạo context có timeout 2 giây
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // 2
	defer cancel()
	orderID := "GO-12345"
	err := placeOrderWithContext(ctx, orderID)
	if err != nil {
		log.Printf("Order failed", err)
	} else {
		log.Printf("Order Success")
	}
}
