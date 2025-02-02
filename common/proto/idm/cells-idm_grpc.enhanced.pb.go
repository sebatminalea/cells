// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: cells-idm.proto

package idm

import (
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	sync "sync"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var (
	enhancedRoleServiceServers     = make(map[string]RoleServiceEnhancedServer)
	enhancedRoleServiceServersLock = sync.RWMutex{}
)

type NamedRoleServiceServer interface {
	RoleServiceServer
	Name() string
}
type RoleServiceEnhancedServer map[string]NamedRoleServiceServer

func (m RoleServiceEnhancedServer) CreateRole(ctx context.Context, r *CreateRoleRequest) (*CreateRoleResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CreateRole should have a context")
	}
	enhancedRoleServiceServersLock.RLock()
	defer enhancedRoleServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateRole(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}

func (m RoleServiceEnhancedServer) DeleteRole(ctx context.Context, r *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeleteRole should have a context")
	}
	enhancedRoleServiceServersLock.RLock()
	defer enhancedRoleServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteRole(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func (m RoleServiceEnhancedServer) SearchRole(r *SearchRoleRequest, s RoleService_SearchRoleServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method SearchRole should have a context")
	}
	enhancedRoleServiceServersLock.RLock()
	defer enhancedRoleServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.SearchRole(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method SearchRole not implemented")
}

func (m RoleServiceEnhancedServer) StreamRole(s RoleService_StreamRoleServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method StreamRole should have a context")
	}
	enhancedRoleServiceServersLock.RLock()
	defer enhancedRoleServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StreamRole(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method StreamRole not implemented")
}

func (m RoleServiceEnhancedServer) CountRole(ctx context.Context, r *SearchRoleRequest) (*CountRoleResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CountRole should have a context")
	}
	enhancedRoleServiceServersLock.RLock()
	defer enhancedRoleServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CountRole(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CountRole not implemented")
}
func (m RoleServiceEnhancedServer) mustEmbedUnimplementedRoleServiceServer() {}
func RegisterRoleServiceEnhancedServer(s grpc.ServiceRegistrar, srv NamedRoleServiceServer) {
	enhancedRoleServiceServersLock.Lock()
	defer enhancedRoleServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedRoleServiceServers[addr]
	if !ok {
		m = RoleServiceEnhancedServer{}
		enhancedRoleServiceServers[addr] = m
		RegisterRoleServiceServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterRoleServiceEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedRoleServiceServersLock.Lock()
	defer enhancedRoleServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedRoleServiceServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedUserServiceServers     = make(map[string]UserServiceEnhancedServer)
	enhancedUserServiceServersLock = sync.RWMutex{}
)

type NamedUserServiceServer interface {
	UserServiceServer
	Name() string
}
type UserServiceEnhancedServer map[string]NamedUserServiceServer

func (m UserServiceEnhancedServer) CreateUser(ctx context.Context, r *CreateUserRequest) (*CreateUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CreateUser should have a context")
	}
	enhancedUserServiceServersLock.RLock()
	defer enhancedUserServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateUser(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (m UserServiceEnhancedServer) DeleteUser(ctx context.Context, r *DeleteUserRequest) (*DeleteUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeleteUser should have a context")
	}
	enhancedUserServiceServersLock.RLock()
	defer enhancedUserServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteUser(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func (m UserServiceEnhancedServer) BindUser(ctx context.Context, r *BindUserRequest) (*BindUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method BindUser should have a context")
	}
	enhancedUserServiceServersLock.RLock()
	defer enhancedUserServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.BindUser(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method BindUser not implemented")
}

func (m UserServiceEnhancedServer) CountUser(ctx context.Context, r *SearchUserRequest) (*CountUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CountUser should have a context")
	}
	enhancedUserServiceServersLock.RLock()
	defer enhancedUserServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CountUser(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CountUser not implemented")
}

func (m UserServiceEnhancedServer) SearchUser(r *SearchUserRequest, s UserService_SearchUserServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method SearchUser should have a context")
	}
	enhancedUserServiceServersLock.RLock()
	defer enhancedUserServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.SearchUser(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}

func (m UserServiceEnhancedServer) StreamUser(s UserService_StreamUserServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method StreamUser should have a context")
	}
	enhancedUserServiceServersLock.RLock()
	defer enhancedUserServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StreamUser(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method StreamUser not implemented")
}
func (m UserServiceEnhancedServer) mustEmbedUnimplementedUserServiceServer() {}
func RegisterUserServiceEnhancedServer(s grpc.ServiceRegistrar, srv NamedUserServiceServer) {
	enhancedUserServiceServersLock.Lock()
	defer enhancedUserServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedUserServiceServers[addr]
	if !ok {
		m = UserServiceEnhancedServer{}
		enhancedUserServiceServers[addr] = m
		RegisterUserServiceServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterUserServiceEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedUserServiceServersLock.Lock()
	defer enhancedUserServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedUserServiceServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedWorkspaceServiceServers     = make(map[string]WorkspaceServiceEnhancedServer)
	enhancedWorkspaceServiceServersLock = sync.RWMutex{}
)

type NamedWorkspaceServiceServer interface {
	WorkspaceServiceServer
	Name() string
}
type WorkspaceServiceEnhancedServer map[string]NamedWorkspaceServiceServer

func (m WorkspaceServiceEnhancedServer) CreateWorkspace(ctx context.Context, r *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CreateWorkspace should have a context")
	}
	enhancedWorkspaceServiceServersLock.RLock()
	defer enhancedWorkspaceServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateWorkspace(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}

func (m WorkspaceServiceEnhancedServer) DeleteWorkspace(ctx context.Context, r *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeleteWorkspace should have a context")
	}
	enhancedWorkspaceServiceServersLock.RLock()
	defer enhancedWorkspaceServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteWorkspace(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}

func (m WorkspaceServiceEnhancedServer) SearchWorkspace(r *SearchWorkspaceRequest, s WorkspaceService_SearchWorkspaceServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method SearchWorkspace should have a context")
	}
	enhancedWorkspaceServiceServersLock.RLock()
	defer enhancedWorkspaceServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.SearchWorkspace(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method SearchWorkspace not implemented")
}

func (m WorkspaceServiceEnhancedServer) StreamWorkspace(s WorkspaceService_StreamWorkspaceServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method StreamWorkspace should have a context")
	}
	enhancedWorkspaceServiceServersLock.RLock()
	defer enhancedWorkspaceServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StreamWorkspace(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method StreamWorkspace not implemented")
}
func (m WorkspaceServiceEnhancedServer) mustEmbedUnimplementedWorkspaceServiceServer() {}
func RegisterWorkspaceServiceEnhancedServer(s grpc.ServiceRegistrar, srv NamedWorkspaceServiceServer) {
	enhancedWorkspaceServiceServersLock.Lock()
	defer enhancedWorkspaceServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedWorkspaceServiceServers[addr]
	if !ok {
		m = WorkspaceServiceEnhancedServer{}
		enhancedWorkspaceServiceServers[addr] = m
		RegisterWorkspaceServiceServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterWorkspaceServiceEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedWorkspaceServiceServersLock.Lock()
	defer enhancedWorkspaceServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedWorkspaceServiceServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedACLServiceServers     = make(map[string]ACLServiceEnhancedServer)
	enhancedACLServiceServersLock = sync.RWMutex{}
)

type NamedACLServiceServer interface {
	ACLServiceServer
	Name() string
}
type ACLServiceEnhancedServer map[string]NamedACLServiceServer

func (m ACLServiceEnhancedServer) CreateACL(ctx context.Context, r *CreateACLRequest) (*CreateACLResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CreateACL should have a context")
	}
	enhancedACLServiceServersLock.RLock()
	defer enhancedACLServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateACL(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CreateACL not implemented")
}

func (m ACLServiceEnhancedServer) ExpireACL(ctx context.Context, r *ExpireACLRequest) (*ExpireACLResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method ExpireACL should have a context")
	}
	enhancedACLServiceServersLock.RLock()
	defer enhancedACLServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ExpireACL(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method ExpireACL not implemented")
}

func (m ACLServiceEnhancedServer) DeleteACL(ctx context.Context, r *DeleteACLRequest) (*DeleteACLResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeleteACL should have a context")
	}
	enhancedACLServiceServersLock.RLock()
	defer enhancedACLServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteACL(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteACL not implemented")
}

func (m ACLServiceEnhancedServer) SearchACL(r *SearchACLRequest, s ACLService_SearchACLServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method SearchACL should have a context")
	}
	enhancedACLServiceServersLock.RLock()
	defer enhancedACLServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.SearchACL(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method SearchACL not implemented")
}

func (m ACLServiceEnhancedServer) StreamACL(s ACLService_StreamACLServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method StreamACL should have a context")
	}
	enhancedACLServiceServersLock.RLock()
	defer enhancedACLServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StreamACL(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method StreamACL not implemented")
}
func (m ACLServiceEnhancedServer) mustEmbedUnimplementedACLServiceServer() {}
func RegisterACLServiceEnhancedServer(s grpc.ServiceRegistrar, srv NamedACLServiceServer) {
	enhancedACLServiceServersLock.Lock()
	defer enhancedACLServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedACLServiceServers[addr]
	if !ok {
		m = ACLServiceEnhancedServer{}
		enhancedACLServiceServers[addr] = m
		RegisterACLServiceServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterACLServiceEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedACLServiceServersLock.Lock()
	defer enhancedACLServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedACLServiceServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedUserMetaServiceServers     = make(map[string]UserMetaServiceEnhancedServer)
	enhancedUserMetaServiceServersLock = sync.RWMutex{}
)

type NamedUserMetaServiceServer interface {
	UserMetaServiceServer
	Name() string
}
type UserMetaServiceEnhancedServer map[string]NamedUserMetaServiceServer

func (m UserMetaServiceEnhancedServer) UpdateUserMeta(ctx context.Context, r *UpdateUserMetaRequest) (*UpdateUserMetaResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method UpdateUserMeta should have a context")
	}
	enhancedUserMetaServiceServersLock.RLock()
	defer enhancedUserMetaServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.UpdateUserMeta(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMeta not implemented")
}

func (m UserMetaServiceEnhancedServer) SearchUserMeta(r *SearchUserMetaRequest, s UserMetaService_SearchUserMetaServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method SearchUserMeta should have a context")
	}
	enhancedUserMetaServiceServersLock.RLock()
	defer enhancedUserMetaServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.SearchUserMeta(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method SearchUserMeta not implemented")
}

func (m UserMetaServiceEnhancedServer) UpdateUserMetaNamespace(ctx context.Context, r *UpdateUserMetaNamespaceRequest) (*UpdateUserMetaNamespaceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method UpdateUserMetaNamespace should have a context")
	}
	enhancedUserMetaServiceServersLock.RLock()
	defer enhancedUserMetaServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.UpdateUserMetaNamespace(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMetaNamespace not implemented")
}

func (m UserMetaServiceEnhancedServer) ListUserMetaNamespace(r *ListUserMetaNamespaceRequest, s UserMetaService_ListUserMetaNamespaceServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method ListUserMetaNamespace should have a context")
	}
	enhancedUserMetaServiceServersLock.RLock()
	defer enhancedUserMetaServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ListUserMetaNamespace(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method ListUserMetaNamespace not implemented")
}
func (m UserMetaServiceEnhancedServer) mustEmbedUnimplementedUserMetaServiceServer() {}
func RegisterUserMetaServiceEnhancedServer(s grpc.ServiceRegistrar, srv NamedUserMetaServiceServer) {
	enhancedUserMetaServiceServersLock.Lock()
	defer enhancedUserMetaServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedUserMetaServiceServers[addr]
	if !ok {
		m = UserMetaServiceEnhancedServer{}
		enhancedUserMetaServiceServers[addr] = m
		RegisterUserMetaServiceServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterUserMetaServiceEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedUserMetaServiceServersLock.Lock()
	defer enhancedUserMetaServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedUserMetaServiceServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedPolicyEngineServiceServers     = make(map[string]PolicyEngineServiceEnhancedServer)
	enhancedPolicyEngineServiceServersLock = sync.RWMutex{}
)

type NamedPolicyEngineServiceServer interface {
	PolicyEngineServiceServer
	Name() string
}
type PolicyEngineServiceEnhancedServer map[string]NamedPolicyEngineServiceServer

func (m PolicyEngineServiceEnhancedServer) IsAllowed(ctx context.Context, r *PolicyEngineRequest) (*PolicyEngineResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method IsAllowed should have a context")
	}
	enhancedPolicyEngineServiceServersLock.RLock()
	defer enhancedPolicyEngineServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.IsAllowed(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method IsAllowed not implemented")
}

func (m PolicyEngineServiceEnhancedServer) StorePolicyGroup(ctx context.Context, r *StorePolicyGroupRequest) (*StorePolicyGroupResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method StorePolicyGroup should have a context")
	}
	enhancedPolicyEngineServiceServersLock.RLock()
	defer enhancedPolicyEngineServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StorePolicyGroup(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method StorePolicyGroup not implemented")
}

func (m PolicyEngineServiceEnhancedServer) ListPolicyGroups(ctx context.Context, r *ListPolicyGroupsRequest) (*ListPolicyGroupsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method ListPolicyGroups should have a context")
	}
	enhancedPolicyEngineServiceServersLock.RLock()
	defer enhancedPolicyEngineServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ListPolicyGroups(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyGroups not implemented")
}

func (m PolicyEngineServiceEnhancedServer) DeletePolicyGroup(ctx context.Context, r *DeletePolicyGroupRequest) (*DeletePolicyGroupResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeletePolicyGroup should have a context")
	}
	enhancedPolicyEngineServiceServersLock.RLock()
	defer enhancedPolicyEngineServiceServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeletePolicyGroup(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyGroup not implemented")
}
func (m PolicyEngineServiceEnhancedServer) mustEmbedUnimplementedPolicyEngineServiceServer() {}
func RegisterPolicyEngineServiceEnhancedServer(s grpc.ServiceRegistrar, srv NamedPolicyEngineServiceServer) {
	enhancedPolicyEngineServiceServersLock.Lock()
	defer enhancedPolicyEngineServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedPolicyEngineServiceServers[addr]
	if !ok {
		m = PolicyEngineServiceEnhancedServer{}
		enhancedPolicyEngineServiceServers[addr] = m
		RegisterPolicyEngineServiceServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterPolicyEngineServiceEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedPolicyEngineServiceServersLock.Lock()
	defer enhancedPolicyEngineServiceServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedPolicyEngineServiceServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}
