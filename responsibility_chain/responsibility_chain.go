package responsibility_chain

import "fmt"

/*
 责任链模式，将多个处理器串成链，让请求在链上传递
*/

// Manager 管理者审批接口
type Manager interface {
	Review(req Request) bool
}

type Request struct {
	// 请求人
	Name string

	// 金额
	Amount int
}

// Leader 直接上级
type Leader struct {
}

// Director 总监
type Director struct {
}

// CFO 首席财务执行官
type CFO struct {
}

func (leader Leader) Review(req Request) bool {
	if req.Amount < 500 {
		fmt.Printf("leader处理了%s的%d元报销", req.Name, req.Amount)
		return true
	}
	return false
}

func (director Director) Review(req Request) bool {
	if req.Amount < 5000 {
		fmt.Printf("director处理了%s的%d元报销", req.Name, req.Amount)
		return true
	}
	return false
}

func (cfo CFO) Review(req Request) bool {
	fmt.Printf("cfo处理了%s的%d元报销", req.Name, req.Amount)
	return true
}

// HandlerChain 责任链处理器
type HandlerChain struct {
	// 责任链切片
	managers []Manager
}

// AddHandler 责任链处理器 添加处理器
func (chain *HandlerChain) AddHandler(manager Manager) {
	chain.managers = append(chain.managers, manager)
}

// AddHandlers 责任链处理器 批量添加处理器
func (chain *HandlerChain) AddHandlers(managers ...Manager) {
	chain.managers = append(chain.managers, managers...)
}

// HandlerRequest 责任链处理器 处理请求
func (chain *HandlerChain) HandlerRequest(request Request) error {
	for _, manager := range chain.managers {
		if manager.Review(request) {
			return nil
		}
	}
	return fmt.Errorf("请求未被处理")
}
