package provider

import "context"

func (p *provider) HelloWorld(ctx context.Context) error {
	err := p.Dbms.FirestoreDb.HelloWorld(ctx)
	if err != nil {
		return err
	}
	return nil
}
