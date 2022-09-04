package usecase

import "context"

func (u *usecase) DoHelloWorld(ctx context.Context) {
	err := u.Provider.HelloWorld(ctx)

	if err != nil {
		return
	}
}
