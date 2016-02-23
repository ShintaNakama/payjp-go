// PAY.JP For Golang

// PAY.JP is a REST based payment API. It supports paying in each case and
// subscription paying, customer management and more.
//
// PAY.JPは、RESTをベースに構成された決済APIです。都度の支払い、定期的な支払い、
// 顧客情報の管理など、ビジネス運用における様々なことができます。
//
// - PAY.JP Getting Started Guide: https://pay.jp/docs/started
//
// - PAY.JP API Docs: https://pay.jp/docs/api/
//
// Installation
//
// You can download Pay.jp Golang SDK by the following command:
//
//   $ go get github.com/payjp/payjp-go/v1
//
// How To Use
//
// Entry point of this package is payjp.New():
//
//   pay := payjp.New("api key", nil)
//
// pay has several children like Token, Customer, Plan, Subscription etc:
//
//   customer := pay.Customer.Get("customer ID")
//
package payjp
