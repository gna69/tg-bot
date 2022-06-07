package auth

import (
	"context"
	"github.com/gna69/tg-bot/internal/entity"
	pb "github.com/gna69/tg-bot/proto"
)

type GroupsManager struct {
	client pb.AuthServiceClient
}

func NewGroupsManager(client pb.AuthServiceClient) *GroupsManager {
	return &GroupsManager{client: client}
}

func (gm *GroupsManager) Add(ctx context.Context, group entity.Object) error {
	_, err := gm.client.CreateGroup(ctx, &pb.Group{
		Name:    group.GetName(),
		OwnerId: int32(group.GetOwnerId()),
		Members: group.GetMembers(),
	})
	return err
}

func (gm *GroupsManager) Update(ctx context.Context, groupId entity.Object) error {
	return nil
}

func (gm *GroupsManager) Delete(ctx context.Context, groupId uint) error {
	_, err := gm.client.RemoveGroup(ctx, &pb.Group{Id: int32(groupId)})
	return err
}

func (gm *GroupsManager) Get(ctx context.Context, groupId uint, ownerId uint) (entity.Object, error) {
	return nil, nil
}

func (gm *GroupsManager) GetAll(ctx context.Context, ownerId uint) ([]entity.Object, error) {
	resp, err := gm.client.GetUserGroups(ctx, &pb.GroupsRequest{OwnerId: int32(ownerId)})
	if err != nil {
		return nil, err
	}
	return gm.toEntity(resp), nil
}

func (gm *GroupsManager) toEntity(resp *pb.GroupsResponse) []entity.Object {
	groups := make([]entity.Object, len(resp.Groups))
	for idx, group := range resp.Groups {
		groups[idx] = &entity.Group{
			Id:      uint(group.Id),
			OwnerId: uint(group.OwnerId),
			Name:    group.Name,
			Members: group.Members,
		}
	}
	return groups
}
