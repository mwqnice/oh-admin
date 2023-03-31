package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

//GetLinkList 获取菜单列表
func (d *Dao) GetLinkList(ctx context.Context, params *dto.GetLinkListRequest) ([]*model.Link, int64, error) {
	link := model.Link{}
	return link.List(ctx, d.engine, params)
}

//CreateLink 创建友链
func (d *Dao) CreateLink(ctx context.Context, link *model.Link) error {
	return link.Create(ctx, d.engine)
}

//GetLinkInfo 获取友链信息
func (d *Dao) GetLinkInfo(ctx context.Context, link *model.Link) (*model.Link, error) {
	return link.Get(ctx, d.engine)
}

//UpdateLink 更新友链
func (d *Dao) UpdateLink(ctx context.Context, link *model.Link) error {
	return link.Update(ctx, d.engine)
}

//DeleteLink 删除友链
func (d *Dao) DeleteLink(ctx context.Context, id int64) error {
	link := model.Link{}
	return link.Delete(ctx, d.engine, int(id))
}
